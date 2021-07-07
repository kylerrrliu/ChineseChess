package proxy

import (
	"reflect"
)

// Proxy is object proxy
type Proxy struct {
	Target  interface{}
	methods map[string]*method
	handler invocationHandler
}

// NewProxy -
func NewProxy(target interface{}, handler invocationHandler) *Proxy {
	typ := reflect.TypeOf(target)
	value := reflect.ValueOf(target)
	methods := make(map[string]*method, 0)
	for i := 0; i < value.NumMethod(); i++ {
		methodValue := value.Method(i)
		methods[typ.Method(i).Name] = &method{value: methodValue}
	}
	return &Proxy{target, methods, handler}
}

// Invoke -
func (p *Proxy) Invoke(name string, args ...interface{}) ([]interface{}, error) {
	return p.handler.invoke(p, p.methods[name], args, name)
}
