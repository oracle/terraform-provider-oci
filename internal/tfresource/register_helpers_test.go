package tfresource

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

func TestUnitRegisterHelper(t *testing.T) {
	testSchema := &schema.Resource{}
	type args struct {
		name           string
		resourceSchema *schema.Resource
	}
	tests := []struct {
		name string
		args args
		want *schema.Resource
	}{
		{
			name: "Test Register Helpers",
			args: args{
				name:           "ai_vision",
				resourceSchema: testSchema,
			},
			want: testSchema,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running %s", tt.name)
			RegisterResource(tt.args.name, tt.args.resourceSchema)
			if !reflect.DeepEqual(globalvar.OciResources[tt.args.name], tt.want) {
				t.Errorf("Error registering OCI resources")
			}
			RegisterDatasource(tt.args.name, tt.args.resourceSchema)
			if !reflect.DeepEqual(globalvar.OciDatasources[tt.args.name], tt.want) {
				t.Errorf("Error registering OCI datasources")
			}
		})
	}
}
