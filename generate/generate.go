package generate

import (
	"fmt"
	"io"
	"log"
	"orm/conf"
	"orm/db"
	"orm/def"
	"os"
	"os/exec"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

func Genertate(tableNames ...string) {
	tableNamesStr := ""
	for _, name := range tableNames {
		if tableNamesStr != "" {
			tableNamesStr += ","
		}
		tableNamesStr += "'" + name + "'"
	}

	//生成所有表信息
	tables := getTables(tableNamesStr)
	for _, tab := range tables {
		outStruct(generateModel(tab), conf.Conf.ModelPath+tab.Name+".go")
	}
	// dbPath := conf.Conf.ModelPath + conf.Conf.Db.DbName + ".go"
	// outStruct(GenerateStruct(tables), dbPath)
}

// titleCase title case.首字母小写
func lowerCase(name string) string {
	vv := []rune(name)
	if len(vv) > 0 {
		if bool(vv[0] >= 'A' && vv[0] <= 'Z') { // title case.首字母小写
			vv[0] += 32
		}
	}

	return string(vv)
}

//获取表信息
func getTables(tableNames string) []def.Table {
	db := db.DB("openapi")
	var tables []def.Table
	if tableNames == "" {
		db.Raw("SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES WHERE table_schema='" + conf.Conf.Db.DbName + "';").Find(&tables)
	} else {
		db.Raw("SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES WHERE TABLE_NAME IN (" + tableNames + ") AND table_schema='" + conf.Conf.Db.DbName + "';").Find(&tables)
	}
	return tables
}

//获取所有字段信息
func getFields(tableName string) []def.Field {
	db := db.DB("openapi")
	var fields []def.Field
	db.Raw("show FULL COLUMNS from " + tableName + ";").Find(&fields)
	return fields
}

func outStruct(content, filename string) {
	var f *os.File
	var err error
	if checkFileIsExist(filename) {
		if !conf.Conf.ModelReplace {
			fmt.Println(filename + " 已存在，需删除才能重新生成...")
			return
		}
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666) //打开文件
		if err != nil {
			panic(err)
		}
	} else {
		f, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()
	_, err = io.WriteString(f, content)
	if err != nil {
		panic(err)
	} else {
		exec.Command("goimports", "-l", "-w", filename).Output()
		log.Println(filename + " 已生成...")
	}
}

//检查文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// GetPackage gen sturct on table
func generateModel(tab def.Table) string {
	var pkg GenPackage
	pkg.SetPackage(conf.Conf.ModelPackage) //package name
	addTable(tab, &pkg)
	return pkg.Generate()
}

func generateStruct(tables []def.Table) string {
	var pkg GenPackage
	pkg.SetPackage(conf.Conf.ModelPackage) //package name
	for _, tab := range tables {
		addTable(tab, &pkg)
	}
	return pkg.Generate()
}

func addTable(tab def.Table, pkg *GenPackage) {
	var sct GenStruct
	sct.SetTableName(tab.Name)
	sct.SetStructName(generator.CamelCase(tab.Name)) // Big hump.大驼峰
	sct.SetNotes(tab.Comment)
	fields := getFields(tab.Name)
	sct.AddElement(genTableElement(fields, true)...) // build element.构造元素
	pkg.AddStruct(sct)
}

// genTableElement Get table columns and comments.获取表列及注释
func genTableElement(cols []def.Field, isOrm bool) (el []GenElement) {
	_tagGorm := "gorm"
	_tagJSON := "json"

	for _, v := range cols {
		var tmp GenElement
		if strings.EqualFold(v.Type, "gorm.Model") { // gorm model
			tmp.SetType(v.Type) //
		} else {
			tmp.SetName(generator.CamelCase(v.Field))
			tmp.SetNotes(v.Comment)
			tmp.SetType(def.GetTypeName(v.Type))
			if v.Key == "PRI" {
				tmp.AddTag(_tagGorm, "primary_key")
			}
		}

		if len(v.Field) > 0 {
			// not simple output
			if isOrm {
				tmp.AddTag(_tagGorm, "column:"+v.Field)
			}

			// json tag
			if true {
				if strings.EqualFold(v.Field, "id") {
					tmp.AddTag(_tagJSON, "-")
				} else {
					tmp.AddTag(_tagJSON, lowerCase(generator.CamelCase(v.Field)))
				}
			}
		}
		el = append(el, tmp)
	}

	return
}
