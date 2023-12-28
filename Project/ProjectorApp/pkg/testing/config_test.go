package projector_test

import (
	"projector/pkg/projector"
	"reflect"
	"testing"
)

func getOpts(args []string) *projector.Opts{
	opts:= &projector.Opts{
		Args: args,
		Config: "",
		Pwd: "",
	}
	return opts
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation projector.Operation){
	opts := getOpts(args)
	config, err:= projector.NewConfig(opts)

	if err!=nil{
		t.Errorf("expected to gett no error %v", err)
	}
	if !reflect.DeepEqual(expectedArgs, config.Args){
		t.Errorf("expected args to be an %+v but got +%v", expectedArgs, config.Args)
	}
	if config.Operation != operation{
		t.Errorf("expected operation was %v but got %v", operation, config.Operation)
	}

}
func TestConfigPrint(t* testing.T){
	testConfig(t, []string{},[]string{}, projector.Print)
}


func TestConfigPrintkey(t* testing.T){
	testConfig(t, []string{"foo"}, []string{"foo"}, projector.Print)
}


func TestConfigAddKeyValue(t* testing.T){
	testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, projector.Add)
}


func TestConfigRemoveKey(t* testing.T){
	testConfig(t, []string{"rm","bar"}, []string{"bar"}, projector.Remove)
}
