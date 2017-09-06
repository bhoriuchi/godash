package godash

import (
	"reflect"
	"log"
)

func Intersection(a interface{}, b interface{}) []interface{} {
	res := make([]interface{}, 0)
	kindA := reflect.TypeOf(a).Kind()
	kindB := reflect.TypeOf(b).Kind()

	if (kindA != reflect.Array && kindA != reflect.Slice) || (kindB != reflect.Array && kindB != reflect.Slice) {
		return res
	}

	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)

	for i := 0; i < valA.Len(); i++ {
		val := valA.Index(i)
		for j := 0; j < valB.Len(); j++ {
			log.Printf("%T:%v, == %T:%v", val.Kind(), val, valB.Index(j).Kind(), valB.Index(j))
			if val == valB.Index(j) {
			// if fmt.Sprintf("%+v", val) == fmt.Sprintf("%+v", valB.Index(j)) {
				res = append(res, val)
				break
			}
		}
	}
	return res
}
