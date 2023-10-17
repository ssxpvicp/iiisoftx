package _case

import (
	"errors"
	"fmt"
	"reflect"
)

// ReflectCase 反射
func ReflectCase() {
	type user struct {
		ID    int64
		Name  string
		Hobby []string
	}
	type outUser struct {
		ID    int64
		Name  string
		Hobby []string
	}
	U := user{ID: 1, Name: "nick", Hobby: []string{"篮球", "羽毛球"}}
	out := outUser{}
	res := copys(&out, U)
	fmt.Println(res, out)

	listUser := []user{
		{ID: 1, Name: "nick", Hobby: []string{"篮球", "羽毛球"}},
		{ID: 2, Name: "nick1", Hobby: []string{"篮球2", "羽毛球"}},
		{ID: 3, Name: "nick2", Hobby: []string{"篮球3", "羽毛球"}},
	}
	list := sliceColumn(listUser, "Hobby")
	fmt.Println(list)
}
func copys(dest interface{}, source interface{}) error {
	sType := reflect.TypeOf(source)
	sValue := reflect.ValueOf(source)
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		sValue = sValue.Elem()
	}
	dType := reflect.TypeOf(dest)
	dValue := reflect.ValueOf(dest)
	if dType.Kind() != reflect.Ptr {
		return errors.New("目标对象必须为struct指针类型")
	}
	dType = dType.Elem()
	dValue = dValue.Elem()
	if sValue.Kind() != reflect.Struct {
		return errors.New("源对象必须为struct或struct的指针")
	}
	if dValue.Kind() != reflect.Struct {
		return errors.New("目录对象必须为struct的指针")
	}
	destObj := reflect.New(dType)
	for i := 0; i < dType.NumField(); i++ {
		destField := dType.Field(i)
		if sourceField, ok := sType.FieldByName(destField.Name); ok {
			if destField.Type != sourceField.Type {
				continue
			}
			value := sValue.FieldByName(destField.Name)
			destObj.Elem().FieldByName(destField.Name).Set(value)
		}
	}
	dValue.Set(destObj.Elem())
	return nil
}
func sliceColumn(slice interface{}, column string) interface{} {
	t := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		val := v.FieldByName(column)
		return val.Interface()
	}
	if v.Kind() != reflect.Slice {
		return nil
	}
	t = t.Elem()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	f, _ := t.FieldByName(column)
	sliceType := reflect.SliceOf(f.Type)
	s := reflect.MakeSlice(sliceType, 0, 0)
	for i := 0; i < v.Len(); i++ {
		o := v.Index(i)
		if o.Kind() == reflect.Struct {
			val := o.FieldByName(column)
			s = reflect.Append(s, val)
		}
	}
	return s.Interface()
}
