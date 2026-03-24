package core

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// Define schema
func MockedCoreVcnResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"byoipv6cidr_details": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ipv6cidr_block":  {Type: schema.TypeString, Required: true},
					"byoipv6range_id": {Type: schema.TypeString, Required: true},
				},
			},
		},
	}
}

func buildByoipv6DetailsBaseCrud(t *testing.T, byoipv6Details []interface{}) tfresource.BaseCrud {
	d := schema.TestResourceDataRaw(t, MockedCoreVcnResourceSchema(), map[string]interface{}{})
	original := make([]interface{}, len(byoipv6Details))

	for _, element := range byoipv6Details {
		original = append(original, map[string]interface{}{"ipv6cidr_block": element, "byoipv6range_id": "abc"})
	}
	d.Set("byoipv6cidr_details", original)
	return tfresource.BaseCrud{D: d}
}

func TestCoreVcnResourceCrud_setBYOIPv6Details(t *testing.T) {
	type fields struct {
		BaseCrud tfresource.BaseCrud
	}

	type args struct {
		byoipv6cidrBlocks []interface{}
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "Test empty blocks and details",
			fields: fields{
				BaseCrud: buildByoipv6DetailsBaseCrud(t, []interface{}{}),
			},
			args: args{byoipv6cidrBlocks: []interface{}{}},
			want: []string{},
		},
		{
			name: "Test empty details and valid blocks",
			fields: fields{
				BaseCrud: buildByoipv6DetailsBaseCrud(t, []interface{}{}),
			},
			args: args{byoipv6cidrBlocks: []interface{}{"2607:f590:2::/48", "2607:f590:3::/48"}},
			want: []string{"2607:f590:2::/48", "2607:f590:3::/48"},
		},
		{
			name: "Test empty blocks and valid details",
			fields: fields{
				BaseCrud: buildByoipv6DetailsBaseCrud(t, []interface{}{"2607:f590:0000:0200::/64"}),
			},
			args: args{byoipv6cidrBlocks: []interface{}{}},
			want: []string{"2607:f590:0000:0200::/64"},
		},
		{
			name: "Test valid blocks and details",
			fields: fields{
				BaseCrud: buildByoipv6DetailsBaseCrud(t, []interface{}{"2607:f590:0000:0200::/64"}),
			},
			args: args{byoipv6cidrBlocks: []interface{}{"2607:f590:2::/48", "2607:f590:3::/48", "2607:f590:0000:0200::/80"}},
			want: []string{"2607:f590:0000:0200::/64", "2607:f590:2::/48", "2607:f590:3::/48", "2607:f590:0000:0200::/80"},
		},
		{
			name: "Test overlap between blocks and details",
			fields: fields{
				BaseCrud: buildByoipv6DetailsBaseCrud(t, []interface{}{"2001:db8:1:2::/80"}),
			},
			args: args{byoipv6cidrBlocks: []interface{}{"2607:f590:2::/48", "2607:f590:3::/48", "2001:0db8:0001:0002:0000:0000:0000:0000/80"}},
			want: []string{"2001:db8:1:2::/80", "2607:f590:2::/48", "2607:f590:3::/48"},
		},
		{
			name: "Test non-64 blocks and details",
			fields: fields{
				BaseCrud: buildByoipv6DetailsBaseCrud(t, []interface{}{"2001:db8:1:2::/80"}),
			},
			args: args{byoipv6cidrBlocks: []interface{}{"2001:db8:1:2::/80", "2001:db8:1:2::/90"}},
			want: []string{"2001:db8:1:2::/80", "2001:db8:1:2::/90"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CoreVcnResourceCrud{
				BaseCrud: tt.fields.BaseCrud,
			}
			s.setBYOIPv6Details(tt.args.byoipv6cidrBlocks)
			expectedBlocks := computeIPv6BlocksFromBYOIPv6Details(s.D.Get("byoipv6cidr_details"))
			if !reflect.DeepEqual(expectedBlocks, tt.want) {
				t.Errorf("setBYOIPv6Details() = %v, want %v", expectedBlocks, tt.want)
			}
		})
	}
}
