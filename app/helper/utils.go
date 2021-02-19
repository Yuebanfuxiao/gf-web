package helper

import (
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"strings"
)

// 加密字符串
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// 密码对比
func PasswordCompare(secret string, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(secret), []byte(password)); err != nil {
		return false, err
	}

	return true, nil
}

// 下划线转驼峰
func CaseToCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 获取值的类型
func Kind(value interface{}) reflect.Kind {
	var (
		rv   = reflect.ValueOf(value)
		kind = rv.Kind()
	)

	if kind == reflect.Ptr {
		kind = rv.Elem().Kind()
	}

	return kind
}