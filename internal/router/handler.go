package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/middleware"
	"reflect"
)

func h(fn any) gin.HandlerFunc {
	f := reflect.ValueOf(fn)
	if f.Kind() != reflect.Func {
		panic("不是正确的func类型")
	}

	ft := f.Type()
	fc := ft.NumIn()

	return func(c *gin.Context) {
		var params []reflect.Value
		for i := 0; i < fc; i++ {
			fi := ft.In(i)
			if fi.Kind() == reflect.Struct {
				req := reflect.New(fi)
				err := c.ShouldBindJSON(req.Interface())
				if err != nil {
					panic(err)
				}
				params = append(params, req.Elem())
			} else if fi.AssignableTo(reflect.TypeOf(&gin.Context{})) {
				params = append(params, reflect.ValueOf(c))
			}
		}
		// 解析返回
		rt := reflect.ValueOf(fn).Call(params)
		if len(rt) > 0 {
			result := rt[0].Interface()
			c.Set(middleware.HandlerResultKey, result)
		}
	}
}
