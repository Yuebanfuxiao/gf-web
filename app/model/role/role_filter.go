package role

import (
	"fmt"
	"gf-web/app/helper"
	"gf-web/app/model"
	"github.com/gogf/gf/frame/g"
	"reflect"
	"strings"
)

const (
	SortByDefault           = "default"
	SortByLatestUpdatedTime = "latest_update_time"
	SortByOldestUpdatedTime = "oldest_update_time"
)

type Filter struct {
	Args  g.Map
	Model *arModel
}

func NewFilter(args g.Map) *Filter {
	f := &Filter{}
	f.Args = args
	f.Model = Model

	f.Setup()

	model.HandleFilter()

	return f
}

// 初始化设置
func (f *Filter) Setup() {
	if sort, ok := f.Args["sort"]; !ok || sort == "" {
		f.Args["sort"] = SortByDefault
	}
}

// 过滤角色名字
func (f *Filter) Name(value interface{}) {
	switch helper.Kind(value) {
	case reflect.String:
		f.Model = f.Model.Where("name = ?", "%"+value.(string)+"%")
	}
}

// 过滤角色状态
func (f *Filter) Status(value interface{}) {
	switch helper.Kind(value) {
	case reflect.Int, reflect.Uint:
		f.Model = f.Model.Where("status = ?", value)
	case reflect.String:
		f.Model = f.Model.Where("status in (?)", strings.Split(value.(string), ","))
	case reflect.Slice, reflect.Array:
		f.Model = f.Model.Where("status in (?)", value)
	}
}

// 过滤分页
func (f *Filter) Page(value interface{}) {
	fmt.Println(value)
	f.Model = f.Model.Page(f.HandlePageArgs(f.Args))
}

// 过滤排序
func (f *Filter) Sort(value interface{}) {
	switch value {
	case SortByLatestUpdatedTime:
		f.SortByLatestUpdatedAt()
	case SortByOldestUpdatedTime:
		f.SortByOldestUpdatedAt()
	default:
		f.SortByDefault()
	}
}

// 默认排序
func (f *Filter) SortByDefault() {
	f.Model = f.Model.Order("sort_index", "desc").Order("id", "desc")
}

// 最新更新时间排序
func (f *Filter) SortByLatestUpdatedAt() {
	f.Model = f.Model.Order("updated_at", "desc")
}

// 最旧更新时间排序
func (f *Filter) SortByOldestUpdatedAt() {
	f.Model = f.Model.Order("updated_at", "asc")
}
