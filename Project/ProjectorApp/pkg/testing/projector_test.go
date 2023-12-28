package projector_test

import (
	"projector/pkg/projector"
	"testing"
)


func getData() *projector.Data{
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/":{
				"foo" : "bar1",
				"fem" : "is_great",
			},
			"/foo":{
				"foo": "bar2",
			},
			"/foo/bar":{
				"foo":"bar2",
			},
		},
	}
}

func getProjector(pwd string, data *projector.Data) *projector.Projector{
	return projector.CreateProjector(
			&projector.Config{
				Args: []string{},
				Operation: projector.Print,
				Pwd: pwd,
				Config: "Hello, Fem",
			},data)

}

func test(t *testing.T, proj *projector.Projector, key, value string){
	v, ok := proj.GetValue(key)
	if !ok {
		t.Error("Something")
	}
	if value != v {
		t.Errorf("Expected to find bar 3 but received %v", value)
	}

}

func TestGetValue(t *testing.T){
	data := getData()
	proj := getProjector("/foo/bar", data)
	test(t, proj, "foo", "bar2")
}
