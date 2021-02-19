package model

import (
	"database/sql"
	"gf-web/app/helper"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"log"
	"reflect"
	"time"
)

type ArModel interface {
	As(as string) *ArModel
	TX(tx *gdb.TX) *ArModel
	Master() *ArModel
	Slave() *ArModel
	LeftJoin(table ...string) *ArModel
	RightJoin(table ...string) *ArModel
	InnerJoin(table ...string) *ArModel
	Fields(fieldNamesOrMapStruct ...interface{}) *ArModel
	FieldsEx(fieldNamesOrMapStruct ...interface{}) *ArModel
	Option(option int) *ArModel
	OmitEmpty() *ArModel
	Filter() *ArModel
	Where(where interface{}, args ...interface{}) *ArModel
	WherePri(where interface{}, args ...interface{}) *ArModel
	And(where interface{}, args ...interface{}) *ArModel
	Or(where interface{}, args ...interface{}) *ArModel
	Group(groupBy string) *ArModel
	Order(orderBy ...string) *ArModel
	Limit(limit ...int) *ArModel
	Offset(offset int) *ArModel
	Page(page, limit int) *ArModel
	Batch(batch int) *ArModel
	Cache(duration time.Duration, name ...string) *ArModel
	Data(data ...interface{}) *ArModel
	Delete(where ...interface{}) (result sql.Result, err error)
}

type IFilter interface {
	Setup()

}

type Abstract struct {
}

// 处理参数过滤
func HandleFilter(filter interface{}) {
	var (
		fv reflect.Value
		mv reflect.Value
	)

	fv = reflect.ValueOf(filter)

	defer func() {
		if err := recover(); err != nil {
			log.Printf("recover:%v", err)
		}
	}()

	for name, value := range filter. {
		mv = fv.MethodByName(helper.CaseToCamel(name))

		if mv.Kind() != reflect.Invalid {
			mv.Call([]reflect.Value{reflect.ValueOf(value)})
		}
	}
}

// 处理分页参数
func HandlePageArgs(args g.Map) (page int, limit int) {
	if p, ok := args["page"]; ok {
		page = gconv.Int(p)
	} else {
		page = 1
	}

	if l, ok := args["limit"]; ok {
		limit = gconv.Int(l)
	} else {
		limit = gconv.Int(g.Cfg().Get("pagination.DefaultLimit"))
	}

	return
}
