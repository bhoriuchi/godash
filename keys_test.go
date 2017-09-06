package godash

import (
	"testing"
	"fmt"
)

var mapTestData = map[string]interface{}{
	"a": 1,
	"b": 2,
	"c": 3,
}

var mapTestDataKeys = [3]interface{}{ "a", "b", "c" }
var arrTestDataKeys = [3]interface{}{0, 1, 2}

func TestKeys(t *testing.T) {
	k := Keys(mapTestData)
	if fmt.Sprintf("%+v", mapTestDataKeys) != fmt.Sprintf("%+v", k) {
		t.Errorf("Expected %+v, Got %+v", mapTestDataKeys, k)
	}

	ak := Keys(mapTestDataKeys)
	if fmt.Sprintf("%+v", arrTestDataKeys) != fmt.Sprintf("%+v", ak) {
		t.Errorf("Expected %+v, Got %+v", arrTestDataKeys, ak)
	}
}