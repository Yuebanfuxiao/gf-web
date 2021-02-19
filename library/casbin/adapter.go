package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"runtime"
)

const (
	defaultGroupName = "default"
	defaultTableName = "casbin_policy"
)

type Rule struct {
	PType string `json:"ptype"`
	V0    string `json:"v0"`
	V1    string `json:"v1"`
	V2    string `json:"v2"`
	V3    string `json:"v3"`
	V4    string `json:"v4"`
	V5    string `json:"v5"`
}

// Adapter represents the gdb adapter for policy storage.
type Adapter struct {
	GroupName string
	TableName string
	db        gdb.DB
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
	// 注意不用的时候不需要使用Close方法关闭数据库连接(并且gdb也没有提供Close方法)，
	// 数据库引擎底层采用了链接池设计，当链接不再使用时会自动关闭
	a.db = nil
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(groupName string, tableName string) (*Adapter, error) {
	return NewAdapterFromOptions(&Adapter{
		GroupName: groupName,
		TableName: tableName,
	})
}

// NewAdapterFromOptions is the constructor for Adapter with existed connection
func NewAdapterFromOptions(adapter *Adapter) (*Adapter, error) {
	if adapter.TableName == "" {
		adapter.TableName = defaultTableName
	}

	if adapter.GroupName == "" {
		adapter.GroupName = defaultGroupName
	}

	if adapter.db == nil {
		err := adapter.open()
		if err != nil {
			return nil, err
		}

		runtime.SetFinalizer(adapter, finalizer)
	}

	return adapter, nil
}

func (a *Adapter) open() error {
	a.db = g.DB(a.GroupName)

	return a.createTable()
}

func (a *Adapter) close() error {
	// 注意不用的时候不需要使用Close方法关闭数据库连接(并且gdb也没有提供Close方法)，
	// 数据库引擎底层采用了链接池设计，当链接不再使用时会自动关闭
	a.db = nil
	return nil
}

func (a *Adapter) createTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS %s (
			ptype VARCHAR(10) NOT NULL DEFAULT '' COMMENT '',
			v0 VARCHAR(256) NOT NULL DEFAULT '' COMMENT '',
			v1 VARCHAR(256) NOT NULL DEFAULT '' COMMENT '',
			v2 VARCHAR(256) NOT NULL DEFAULT '' COMMENT '',
			v3 VARCHAR(256) NOT NULL DEFAULT '' COMMENT '',
			v4 VARCHAR(256) NOT NULL DEFAULT '' COMMENT '',
			v5 VARCHAR(256) NOT NULL DEFAULT '' COMMENT ''
		) ENGINE = InnoDB COMMENT = '权限策略表';
	`
	_, err := a.db.Exec(fmt.Sprintf(sql, a.TableName))

	return err
}

func (a *Adapter) dropTable() error {
	_, err := a.db.Exec(fmt.Sprintf("DROP TABLE %s", a.TableName))
	return err
}

func loadPolicyLine(line Rule, model model.Model) {
	lineText := line.PType
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}

	persist.LoadPolicyLine(lineText, model)
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	var lines []Rule

	if err := a.db.Table(a.TableName).Scan(&lines); err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}

	return nil
}

func savePolicyLine(ptype string, rule []string) Rule {
	line := Rule{}

	line.PType = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}

	return line
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {
	err := a.dropTable()
	if err != nil {
		return err
	}
	err = a.createTable()
	if err != nil {
		return err
	}

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			_, err := a.db.Table(a.TableName).Data(&line).Insert()
			if err != nil {
				return err
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			_, err := a.db.Table(a.TableName).Data(&line).Insert()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := a.db.Table(a.TableName).Data(&line).Insert()
	return err
}

// AddPolicies batch add policy rule to the storage.
func (a *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	var lines []Rule
	for _, rule := range rules {
		lines = append(lines, savePolicyLine(ptype, rule))
	}

	_, err := a.db.Table(a.TableName).Data(&lines).Insert()
	return err
}

// RemovePolicies batch removes policy rule from the storage.
func (a *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	var (
		db = a.db.Table(a.TableName)
	)

	for _, rule := range rules {
		line := savePolicyLine(ptype, rule)
		sql := ""
		val := make([]interface{}, 0)

		sql = "(ptype = ?"
		val = append(val, ptype)
		if line.V0 != "" {
			sql += " and v0 = ?"
			val = append(val, line.V0)
		}
		if line.V1 != "" {
			sql += " and v1 = ?"
			val = append(val, line.V1)
		}
		if line.V2 != "" {
			sql += " and v2 = ?"
			val = append(val, line.V2)
		}
		if line.V3 != "" {
			sql += " and v3 = ?"
			val = append(val, line.V3)
		}
		if line.V4 != "" {
			sql += " and v4 = ?"
			val = append(val, line.V4)
		}
		if line.V5 != "" {
			sql += " and v5 = ?"
			val = append(val, line.V5)
		}
		sql += ")"

		db.Or(sql, val...)
	}

	_, err := db.Delete()
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	err := rawDelete(a, line)
	return err
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := Rule{}

	line.PType = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}
	err := rawDelete(a, line)
	return err
}

func rawDelete(a *Adapter, line Rule) error {
	db := a.db.Table(a.TableName)

	db.Where("ptype = ?", line.PType)
	if line.V0 != "" {
		db.Where("v0 = ?", line.V0)
	}
	if line.V1 != "" {
		db.Where("v1 = ?", line.V1)
	}
	if line.V2 != "" {
		db.Where("v2 = ?", line.V2)
	}
	if line.V3 != "" {
		db.Where("v3 = ?", line.V3)
	}
	if line.V4 != "" {
		db.Where("v4 = ?", line.V4)
	}
	if line.V5 != "" {
		db.Where("v5 = ?", line.V5)
	}

	_, err := db.Delete()
	return err
}
