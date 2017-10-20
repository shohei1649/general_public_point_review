package test

import (
	"testing"
)

func TestCreateOccupyDefaultPostData(t *testing.T) {

	var defaultPostJSON = CreateDefaultPostJSON()

	t.Log(defaultPostJSON)

	var defaultPostJSONBase64Str = CreateBase64PostJSON(defaultPostJSON)

	t.Log(defaultPostJSONBase64Str)

}
