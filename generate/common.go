package generate

import (
	"bytes"
	"fmt"
	"html/template"
	"sort"
	"strings"

	"orm/conf"
	"orm/def"
)

// GenElement element of sturct.元素类
type GenElement struct {
	Name  string              // Name.元素名
	Type  string              // Type.类型标记
	Notes string              // Notes.注释
	Tags  map[string][]string // tages.标记
}

// GenStruct struct of IStruct .结构体
type GenStruct struct {
	SQLBuildStr string       // Create SQL statements.创建sql语句
	TableName   string       // table_name.表名
	Name        string       // name.名字
	Notes       string       // notes.注释
	Em          []GenElement // em.元素组合
}

// GenPackage package of IPackage.包体
type GenPackage struct {
	Name        string            // name.名字
	Imports     map[string]string // Inclusion term.元素组合
	Structs     []GenStruct       // struct list .结构体组合
	FuncStrList []string          // func of template on string. 函数的最终定义
}

// SetName Setting element name.设置元素名字
func (e *GenElement) SetName(name string) {
	e.Name = name
}

// SetType Setting element type.设置元素类型
func (e *GenElement) SetType(tp string) {
	e.Type = tp
}

// SetNotes Setting element notes.设置注释
func (e *GenElement) SetNotes(notes string) {
	e.Notes = strings.Replace(notes, "\n", ",", -1)
}

// AddTag Add a tag .添加一个tag标记
func (e *GenElement) AddTag(k string, v string) {
	if e.Tags == nil {
		e.Tags = make(map[string][]string)
	}
	e.Tags[k] = append(e.Tags[k], v)
}

// Generate Get the result data.获取结果数据
func (e *GenElement) Generate() string {
	tag := ""
	if e.Tags != nil {
		var ks []string
		for k := range e.Tags {
			ks = append(ks, k)
		}
		sort.Strings(ks)

		var tags []string
		for _, v := range ks {
			tags = append(tags, fmt.Sprintf(`%v:"%v"`, v, strings.Join(e.Tags[v], ";")))
		}
		tag = fmt.Sprintf("`%v`", strings.Join(tags, " "))
	}

	var p PrintAtom
	if len(e.Notes) > 0 {
		p.Add(e.Name, e.Type, tag, "// "+e.Notes)
	} else {
		p.Add(e.Name, e.Type, tag)
	}

	return p.Generates()[0]
}

//////////////////////////////////////////////////////////////////////////////
// struct
//////////////////////////////////////////////////////////////////////////////

// SetCreatTableStr Set up SQL create statement, backup use setup create statement, backup use.设置创建语句，备份使用
func (s *GenStruct) SetCreatTableStr(sql string) {
	s.SQLBuildStr = sql
}

// SetTableName Setting the name of struct.设置struct名字
func (s *GenStruct) SetTableName(name string) {
	s.TableName = name
}

// SetStructName Setting the name of struct.设置struct名字
func (s *GenStruct) SetStructName(name string) {
	s.Name = name
}

// SetNotes set the notes.设置注释
func (s *GenStruct) SetNotes(notes string) {
	if len(notes) == 0 {
		notes = "[...]" // default of struct notes(for export ).struct 默认注释(为了导出注释)
	}

	notes = s.Name + " " + notes

	a := strings.Split(notes, "\n")
	var text []string

	for _, v := range a {
		// if len(v) > 0 {
		text = append(text, "// "+v)
		// }
	}
	s.Notes = strings.Join(text, "\r\n")
}

// AddElement Add one or more elements.添加一个/或多个元素
func (s *GenStruct) AddElement(e ...GenElement) {
	s.Em = append(s.Em, e...)
}

func (s *GenStruct) GenerateTableName() []string {
	tmpl, err := template.New("gen_tnf").Parse(def.TableName)
	if err != nil {
		panic(err)
	}
	var data struct {
		TableName  string
		StructName string
	}
	data.TableName, data.StructName = s.TableName, s.Name
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	return []string{buf.String()}
}

func (s *GenStruct) GenerateCreate() []string {
	tmpl, err := template.New("gen_create").Parse(def.Create)
	if err != nil {
		panic(err)
	}
	var data struct {
		DataBase   string
		StructName string
	}
	data.DataBase, data.StructName = conf.Conf.Db.DbName, s.Name
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	return []string{buf.String()}
}

func (s *GenStruct) GenerateUpdate() []string {
	tmpl, err := template.New("gen_update").Parse(def.Update)
	if err != nil {
		panic(err)
	}
	var data struct {
		DataBase   string
		StructName string
	}
	data.DataBase, data.StructName = conf.Conf.Db.DbName, s.Name
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	return []string{buf.String()}
}

func (s *GenStruct) GenerateGetById() []string {
	tmpl, err := template.New("gen_getbyid").Parse(def.GetById)
	if err != nil {
		panic(err)
	}
	var data struct {
		DataBase   string
		StructName string
	}
	data.DataBase, data.StructName = conf.Conf.Db.DbName, s.Name
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	return []string{buf.String()}
}

// Generates Get the result data.获取结果数据
func (s *GenStruct) Generates() []string {
	var p PrintAtom
	// if config.GetIsOutSQL() {
	// 	p.Add("/******sql******")
	// 	p.Add(s.SQLBuildStr)
	// 	p.Add("******sql******/")
	// }
	p.Add(s.Notes)
	p.Add("type", s.Name, "struct {")
	mp := make(map[string]bool, len(s.Em))
	for _, v := range s.Em {
		if !mp[v.Name] {
			mp[v.Name] = true
			p.Add(v.Generate())
		}
	}
	p.Add("}")

	return p.Generates()
}

//////////////////////////////////////////////////////////////////////////////
// package
//////////////////////////////////////////////////////////////////////////////

// SetPackage Defining package names.定义包名
func (p *GenPackage) SetPackage(pname string) {
	p.Name = pname
}

// AddImport Add import by type.通过类型添加import
func (p *GenPackage) AddImport(imp string) {
	if p.Imports == nil {
		p.Imports = make(map[string]string)
	}
	p.Imports[imp] = imp
}

// AddStruct Add a structure.添加一个结构体
func (p *GenPackage) AddStruct(st GenStruct) {
	p.Structs = append(p.Structs, st)
}

// Generate Get the result data.获取结果数据
func (p *GenPackage) Generate() string {
	p.genimport() // auto add import .补充 import

	var pa PrintAtom
	pa.Add("package", p.Name)
	// add import
	if p.Imports != nil {
		pa.Add("import (")
		for _, v := range p.Imports {
			pa.Add(v)
		}
		pa.Add(")")
	}
	// -----------end
	// add struct
	for _, v := range p.Structs {
		for _, v1 := range v.Generates() {
			pa.Add(v1)
		}

		if true { // add table name func
			for _, v1 := range v.GenerateTableName() {
				pa.Add(v1)
			}
		}
		if true { // add table name func
			for _, v1 := range v.GenerateCreate() {
				pa.Add(v1)
			}
		}
		if true { // add table name func
			for _, v1 := range v.GenerateUpdate() {
				pa.Add(v1)
			}
		}
		if true { // add table name func
			for _, v1 := range v.GenerateGetById() {
				pa.Add(v1)
			}
		}
	}
	// -----------end

	// add func
	for _, v := range p.FuncStrList {
		pa.Add(v)
	}
	// -----------end

	// output.输出
	strOut := ""
	for _, v := range pa.Generates() {
		strOut += v + "\n"
	}

	return strOut
}

// AddFuncStr add func coding string.添加函数串
func (p *GenPackage) AddFuncStr(src string) {
	p.FuncStrList = append(p.FuncStrList, src)
}

// compensate and import .获取结果数据
func (p *GenPackage) genimport() {
	for _, v := range p.Structs {
		for _, v1 := range v.Em {
			if v2, ok := def.EImportsHead[v1.Type]; ok {
				if len(v2) > 0 {
					p.AddImport(v2)
				}
			}
		}
	}
}
