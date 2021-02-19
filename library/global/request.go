package global

import (
	"errors"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"reflect"
	"strings"
)

type Rules []struct {
	Field   string
	Bind    string
	Rule    string
	Message string
}

type IRequest interface {
	Rules() Rules
}

type Request struct {
	R *ghttp.Request
}

// 新建请求
func NewRequest(r *ghttp.Request) *Request {
	return &Request{
		R: r,
	}
}

// 接收请求
func (l *Request) Accept(req IRequest, args ...interface{}) error {
	var (
		rv       reflect.Value
		rk       reflect.Kind
		rules    map[string]string
		messages map[string]interface{}
		data     map[string]interface{}
		params   = l.R.GetMap()
	)

	rules, messages = l.handleValidateRules(req)

	if err := gvalid.CheckMap(params, rules, messages); err != nil {
		return l.handleValidateError(req, err)
	}

	if err := l.R.Parse(req); err != nil {
		return err
	}

	data = gconv.MapDeep(req)

	for _, arg := range args {
		if v, ok := arg.(reflect.Value); ok {
			rv = v
			arg = v.Interface()
		} else {
			rv = reflect.ValueOf(arg)
		}

		rk = rv.Kind()

		if rk == reflect.Ptr {
			rv = rv.Elem()
			rk = rv.Kind()
		}

		switch rk {
		case reflect.Struct:
			if err := gconv.Struct(req, arg); err != nil {
				return errors.New("data parse failed")
			}
		case reflect.Map:
			for k, v := range data {
				if _, ok := params[k]; ok {
					arg.(map[string]interface{})[k] = v
				}
			}
		}
	}

	return nil
}

// 处理验证规则
func (l *Request) handleValidateRules(req IRequest) (map[string]string, map[string]interface{}) {
	tmpRules := map[string][]string{}

	tmpMessages := map[string]map[string]string{}

	rules := map[string]string{}

	messages := map[string]interface{}{}

	for _, rule := range req.Rules() {
		if len(tmpRules[rule.Field]) == 0 {
			tmpRules[rule.Field] = []string{}
		}

		tmpRules[rule.Field] = append(tmpRules[rule.Field], rule.Rule)

		if len(tmpMessages[rule.Field]) == 0 {
			tmpMessages[rule.Field] = map[string]string{}
		}

		tmpMessages[rule.Field][strings.Split(rule.Rule, ":")[0]] = rule.Message
	}

	for f, r := range tmpRules {
		rules[f] = strings.Join(r, "|")
	}

	for f, m := range tmpMessages {
		messages[f] = (interface{})(m)
	}

	return rules, messages
}

// 处理验证错误
func (l *Request) handleValidateError(req IRequest, err *gvalid.Error) error {
	maps := err.Maps()

	for _, rule := range req.Rules() {
		if maps[rule.Field][strings.Split(rule.Rule, ":")[0]] != "" {
			return errors.New(maps[rule.Field][strings.Split(rule.Rule, ":")[0]])
		}
	}

	return nil
}
