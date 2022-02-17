package tfresource

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestUnitDefinedTagsToMap(t *testing.T) {
	desiredValue := map[string]interface{}{
		"testexamples-tag-namespace.tf-example-tag-2": "awesome-app-server",
	}
	val := map[string]interface{}{
		"tf-example-tag-2": "awesome-app-server",
	}
	argument := map[string]map[string]interface{}{
		"testexamples-tag-namespace": val,
	}
	type args struct {
		definedTags map[string]map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test returned value is as expected",
			args: args{
				definedTags: argument,
			},
			want:    desiredValue,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefinedTagsToMap(tt.args.definedTags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefinedTagsToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitMapToDefinedTags(t *testing.T) {
	argument := map[string]interface{}{
		"testexamples-tag-namespace.tf-example-tag-2": "awesome-app-server",
	}
	wrongArgument := map[string]interface{}{
		"testexamples-tag-name.space.tf-example-tag-2": "awesome-app-server",
	}
	val := map[string]interface{}{
		"tf-example-tag-2": "awesome-app-server",
	}
	desiredValue := map[string]map[string]interface{}{
		"testexamples-tag-namespace": val,
	}
	type args struct {
		rawMap map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test returned value is as expected",
			args: args{
				rawMap: argument,
			},
			want:    desiredValue,
			wantErr: false,
		},
		{
			name: "Test returned value is not as expected",
			args: args{
				rawMap: wrongArgument,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapToDefinedTags(tt.args.rawMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToDefinedTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToDefinedTags() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitMapToSystemTags(t *testing.T) {
	argument := map[string]interface{}{
		"testexamples-tag-namespace.tf-example-tag-2": "awesome-app-server",
	}
	wrongArgument := map[string]interface{}{
		"testexamples-tag-name.space.tf-example-tag-2": "awesome-app-server",
	}
	val := map[string]interface{}{
		"tf-example-tag-2": "awesome-app-server",
	}
	desiredValue := map[string]map[string]interface{}{
		"testexamples-tag-namespace": val,
	}
	type args struct {
		rawMap map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test returned value is as expected",
			args: args{
				rawMap: argument,
			},
			want:    desiredValue,
			wantErr: false,
		},
		{
			name: "Test returned value is not as expected",
			args: args{
				rawMap: wrongArgument,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapToSystemTags(tt.args.rawMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToSystemTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToSystemTags() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitSystemTagsToMap(t *testing.T) {
	desiredValue := map[string]interface{}{
		"testexamples-tag-namespace.tf-example-tag-2": "awesome-app-server",
	}
	val := map[string]interface{}{
		"tf-example-tag-2": "awesome-app-server",
	}
	argument := map[string]map[string]interface{}{
		"testexamples-tag-namespace": val,
	}
	type args struct {
		systemTags map[string]map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test returned value is as expected",
			args: args{
				systemTags: argument,
			},
			want:    desiredValue,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SystemTagsToMap(tt.args.systemTags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SystemTagsToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitToLowerCaseKeyMap(t *testing.T) {
	desiredValue1 := map[string]interface{}{
		"testexamples-tag-namespace.tf-example-tag-2": "awesome-app-server",
	}
	argument1 := map[string]interface{}{
		"Testexamples-tag-namespace.Tf-example-tag-2": "awesome-app-server",
	}
	desiredValue2 := map[string]interface{}{
		"testexamples-tag-namespace.tf-example-tag-2": "Awesome-app-server",
	}
	argument2 := map[string]interface{}{
		"Testexamples-tag-namespace.Tf-example-tag-2": "Awesome-app-server",
	}
	type args struct {
		original map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test returned value is as expected",
			args: args{
				original: argument1,
			},
			want:    desiredValue1,
			wantErr: false,
		},
		{
			name: "Test returned value is as expected 2",
			args: args{
				original: argument2,
			},
			want:    desiredValue2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLowerCaseKeyMap(tt.args.original); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLowerCaseKeyMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitDefinedTagsDiffSuppressFunction(t *testing.T) {
	type args struct {
		key string
		old string
		new string
		d   *schema.ResourceData
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test returned value is as expected",
			args: args{
				key: "create_vnic_details.0.defined_tags.mynamespace.mykey",
				old: "abc",
				new: "def",
				d:   nil,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefinedTagsDiffSuppressFunction(tt.args.key, tt.args.old, tt.args.new, tt.args.d); got != tt.want {
				t.Errorf("DefinedTagsDiffSuppressFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
