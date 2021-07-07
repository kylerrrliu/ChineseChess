package proxy

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/go-errors/errors"
)

type invocationHandler interface {
	invoke(proxy *Proxy, method *method, args []interface{}, methodName string) ([]interface{}, error)
}

// MetricsProxy is object proxy that provides metrics for timing and logs error
type MetricsProxy struct {
}

func (mp *MetricsProxy) invoke(proxy *Proxy, method *method, args []interface{}, methodName string) ([]interface{}, error) {
	objectName := reflect.TypeOf(proxy.Target).Elem().Name()

	startTime := time.Now()
	log.Printf(fmt.Sprintf("================%s================", startTime.UTC().String()))
	defer func() {
		latency := float64(time.Since(startTime).Nanoseconds()) / 1e6
		log.Printf(fmt.Sprintf("%s.%s: %fms", objectName, methodName, latency))
	}()
	result, err := method.invoke(args...)
	if err != nil {
		errData := make(map[string]interface{})
		errData["object_name"] = objectName
		errData["method_name"] = methodName
		var errMsg string
		switch err.(type) {
		case *errors.Error:
			errMsg = err.(*errors.Error).ErrorStack()
		default:
			errMsg = errors.Wrap(err, 2).ErrorStack()
		}
		fmt.Println(fmt.Sprintf("%s.%s error: %s", objectName, methodName, errMsg))
		errData["error"] = errMsg

	}

	return result, err
}
