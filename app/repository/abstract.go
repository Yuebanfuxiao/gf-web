package repository

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"time"
)

type Abstract struct {
	request *ghttp.Request
}

// 分页
type Pagination struct {
	Page  int     `json:"page"`
	Total int     `json:"total"`
	Limit int     `json:"limit"`
	List  []g.Map `json:"list"`
}

type arModel interface {
	As(as string) arModel
	TX(tx *gdb.TX) arModel
	Master() arModel
	Slave() arModel
	LeftJoin(table ...string) arModel
	RightJoin(table ...string) arModel
	InnerJoin(table ...string) arModel
	Fields(fieldNamesOrMapStruct ...interface{}) arModel
	FieldsEx(fieldNamesOrMapStruct ...interface{}) arModel
	Option(option int) arModel
	OmitEmpty() arModel
	Filter() arModel
	Where(where interface{}, args ...interface{}) arModel
	WherePri(where interface{}, args ...interface{}) arModel
	And(where interface{}, args ...interface{}) arModel
	Or(where interface{}, args ...interface{}) arModel
	Group(groupBy string) arModel
	Order(orderBy ...string) arModel
	Limit(limit ...int) arModel
	Offset(offset int) arModel
	Page(page, limit int) arModel
	Batch(batch int) arModel
	Cache(duration time.Duration, name ...string) arModel
	Data(data ...interface{}) arModel
	Delete(where ...interface{}) (result sql.Result, err error)
}

// 处理删除结果
func (r *Abstract) handleDeleteResult(result sql.Result) (int64, error) {
	if rows, err := result.RowsAffected(); err != nil {
		return 0, err
	} else {
		return rows, nil
	}
}

// 处理保存结果
func (r *Abstract) handleSaveResult(result sql.Result) (interface{}, interface{}, error) {
	if id, err := result.LastInsertId(); err != nil {
		if rows, err := result.RowsAffected(); err != nil {
			return nil, nil, err
		} else {
			return rows, nil, nil
		}
	} else {
		return nil, gconv.Uint(id), nil
	}
}

func (r *Abstract) DoSave(model interface{}) {

}

func (r *Abstract) TransactionBegin() (*gdb.TX, error) {
	return g.DB().Begin()
}

func (r *Abstract) TransactionOver(tx *gdb.TX, err error) func() {
	return func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}
}

// 携带请求
func (r *Abstract) WithRequest(request *ghttp.Request) {
	r.request = request
}

// 处理分页参数
func (r *Abstract) handlePageArgs(args g.Map) (int, int) {
	var (
		page  int
		limit int
	)

	if page, ok := args["page"]; ok {
		page = gconv.Int(page)
	}

	if page <= 0 {
		page = 1
	}

	if limit, ok := args["limit"]; ok {
		limit = gconv.Int(limit)
	}

	if limit <= 0 {
		limit = 10
	}

	return page, limit
}

// 处理列表
func (r *Abstract) handleList(entities interface{}, fields interface{}) ([]g.Map, error) {
	var (
		err  error
		list = make([]g.Map, 0)
	)

	for _, entity := range gconv.SliceAny(entities) {
		err = gconv.Struct(entity, &fields)

		if err != nil {
			return nil, err
		}

		list = append(list, gconv.Map(fields))
	}

	return list, err
}
