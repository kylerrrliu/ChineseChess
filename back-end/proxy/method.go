package proxy

import (
	"fmt"
	"github.com/go-errors/errors"
	"reflect"
)

type method struct {
	value reflect.Value
}

func (m *method) invoke(args ...interface{}) (res []interface{}, err error) {
	defer func() {
		if p := recover(); p != nil {
			switch err.(type) {
			case *errors.Error:
				fmt.Println(fmt.Sprintf("Method invocation panicked: %s", err.(*errors.Error).ErrorStack()))
			default:
				fmt.Println(fmt.Sprintf("Method invocation panicked: %s", errors.Wrap(err, 1).ErrorStack()))
			}

			err = errors.New(p)
		}
	}()

	params := make([]reflect.Value, 0)
	if args != nil {
		for i := 0; i < len(args); i++ {
			params = append(params, reflect.ValueOf(args[i]))
		}
	}
	res = make([]interface{}, 0)
	ret := m.value.Call(params)
	if ret != nil && len(ret) > 0 {
		for i := 0; i < len(ret); i++ {
			switch ret[i].Interface().(type) {
			case error:
				err = ret[i].Interface().(error)
				break
			default:
				res = append(res, ret[i].Interface())
				break
			}
		}
	}
	return
}
