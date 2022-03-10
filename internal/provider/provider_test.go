package provider

import (
	"reflect"
	"testing"
)

type mockResourceData struct {
	state string
}

func (d *mockResourceData) GetOkExists(_ string) (interface{}, bool) {
	if d.state == "1" {
		return []interface{}{}, false
	}
	return []interface{}{"abc, xyz"}, true
}

func TestUnitIgnoreTags(t *testing.T) {
	type args struct {
		d *mockResourceData
	}
	type testFormat struct {
		name   string
		args   args
		output []string
	}
	tests := []testFormat{
		{
			name:   "Test ignoreDefinedTags is set",
			args:   args{d: &mockResourceData{}},
			output: []string{"abc", "xyz"},
		},
		{
			name:   "Test ignoreDefinedTags is not set",
			args:   args{d: &mockResourceData{state: "1"}},
			output: nil,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if res := IgnoreDefinedTags(test.args.d); reflect.DeepEqual(res, test.output) {
			if res != nil || test.output != nil {
				t.Errorf("Output array - %q which is not equal to expected array - %q", res, test.output)
			}
		}
	}

}
