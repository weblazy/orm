package generate

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

// interval.间隔
var _interval = "\t"

// PrintAtom . atom print .原始打印
type PrintAtom struct {
	lines []string
}

// Add add one to print.打印
func (p *PrintAtom) Add(str ...interface{}) {
	var tmp string
	for _, v := range str {
		tmp += AsString(v) + _interval
	}
	p.lines = append(p.lines, tmp)
}

// Generates Get the generated list.获取生成列表
func (p *PrintAtom) Generates() []string {
	return p.lines
}

func AsString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	case bool:
		return strconv.FormatBool(v)
	default:
		{
			b, _ := json.Marshal(v)
			log.Println(string(b))
			return string(b)
		}
	}
	return fmt.Sprintf("%v", src)
}
