package core

import (
	"context"
	"reflect"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// Define schema
func MockedCoreVcnResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"byoipv6cidr_details": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
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
			want: []string{},
		},
		{
			name: "Test valid blocks and details",
			fields: fields{
				BaseCrud: buildByoipv6DetailsBaseCrud(t, []interface{}{"2607:f590:0000:0200::/64"}),
			},
			args: args{byoipv6cidrBlocks: []interface{}{"2607:f590:2::/48", "2607:f590:3::/48", "2607:f590:0000:0200::/80"}},
			want: []string{"2607:f590:2::/48", "2607:f590:3::/48", "2607:f590:0000:0200::/80"},
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
		{
			name: "Test stale details removed while matching detail retained",
			fields: fields{
				BaseCrud: buildByoipv6DetailsBaseCrud(t, []interface{}{"2001:db8:1:2::/80", "2607:f590:0000:0200::/64"}),
			},
			args: args{byoipv6cidrBlocks: []interface{}{"2001:0db8:0001:0002:0000:0000:0000:0000/80", "2607:f590:0000:2000::/64"}},
			want: []string{"2001:db8:1:2::/80", "2607:f590:0000:2000::/64"},
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

func buildByoipv6DetailsDiff(
	t *testing.T,
	customizeDiff schema.CustomizeDiffFunc,
	stateRaw map[string]interface{},
	configRaw map[string]interface{},
) *terraform.InstanceDiff {
	t.Helper()

	resource := &schema.Resource{
		Schema:        MockedCoreVcnResourceSchema(),
		CustomizeDiff: customizeDiff,
	}

	var state *terraform.InstanceState
	if stateRaw != nil {
		data := schema.TestResourceDataRaw(t, resource.Schema, stateRaw)
		data.SetId("ocid1.vcn.oc1..exampleuniqueID")
		state = data.State()
	}

	diff, err := resource.Diff(context.Background(), state, terraform.NewResourceConfigRaw(configRaw), nil)
	if err != nil {
		t.Fatalf("Resource.Diff() error = %v", err)
	}

	return diff
}

func diffHasByoipv6CidrDetailsChanges(diff *terraform.InstanceDiff) bool {
	if diff == nil {
		return false
	}

	for key := range diff.Attributes {
		if len(key) >= len("byoipv6cidr_details") && key[:len("byoipv6cidr_details")] == "byoipv6cidr_details" {
			return true
		}
	}

	return false
}

func diffHasByoipv6CidrDetailsIndexChanges(diff *terraform.InstanceDiff, index int) bool {
	if diff == nil {
		return false
	}

	prefix := "byoipv6cidr_details." + strconv.Itoa(index) + "."
	for key := range diff.Attributes {
		if len(key) >= len(prefix) && key[:len(prefix)] == prefix {
			return true
		}
	}

	return false
}

func TestSuppressMatchingByoipv6CidrDetailsDiff(t *testing.T) {
	type fields struct {
		stateRaw map[string]interface{}
	}

	type args struct {
		configRaw map[string]interface{}
	}

	tests := []struct {
		name                string
		fields              fields
		args                args
		wantBaselineChanges bool
		wantSuppressed      bool
	}{
		{
			name: "suppresses canonically matching details on existing resources",
			fields: fields{
				stateRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2001:db8:1:2::/80",
							"byoipv6range_id": "(known_after_apply)",
						},
					},
				},
			},
			args: args{
				configRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2001:0db8:0001:0002:0000:0000:0000:0000/80",
							"byoipv6range_id": "ocid1.byoipv6range.oc1..example",
						},
					},
				},
			},
			wantBaselineChanges: true,
			wantSuppressed:      true,
		},
		{
			name: "does not suppresses when old config is nil",
			fields: fields{
				stateRaw: nil,
			},
			args: args{
				configRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2001:0db8:0001:0002:0000:0000:0000:0000/80",
							"byoipv6range_id": "ocid1.byoipv6range.oc1..example",
						},
					},
				},
			},
			wantBaselineChanges: true,
			wantSuppressed:      false,
		},
		{
			name: "suppresses exactly matching details on existing resources",
			fields: fields{
				stateRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2001:0db8:0001:0002:0000:0000:0000:0000/80",
							"byoipv6range_id": "ocid1.byoipv6range.oc1..example",
						},
					},
				},
			},
			args: args{
				configRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2001:0db8:0001:0002:0000:0000:0000:0000/80",
							"byoipv6range_id": "ocid1.byoipv6range.oc1..example",
						},
					},
				},
			},
			wantBaselineChanges: false,
			wantSuppressed:      true,
		},
		{
			name: "suppresses diff when new byoipv6 config is nil",
			fields: fields{
				stateRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2001:db8:1:2::/80",
							"byoipv6range_id": "(known_after_apply)",
						},
					},
				},
			},
			args: args{
				configRaw: nil,
			},
			wantBaselineChanges: false,
			wantSuppressed:      true,
		},
		{
			name:   "does not suppress on create when there is no existing state",
			fields: fields{},
			args: args{
				configRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2607:f590:0000:2000::/64",
							"byoipv6range_id": "ocid1.byoipv6range.oc1..example",
						},
					},
				},
			},
			wantBaselineChanges: true,
			wantSuppressed:      false,
		},
		{
			name: "does not suppress genuinely new byoipv6 details",
			fields: fields{
				stateRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2607:f590:0000:0200::/64",
							"byoipv6range_id": "(known_after_apply)",
						},
					},
				},
			},
			args: args{
				configRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2607:f590:0000:2000::/64",
							"byoipv6range_id": "ocid1.byoipv6range.oc1..example",
						},
					},
				},
			},
			wantBaselineChanges: true,
			wantSuppressed:      false,
		},
		{
			name: "reorders middle-entry drift into append-style add",
			fields: fields{
				stateRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2607:f590:0000:0200::/64",
							"byoipv6range_id": "range-a",
						},
						map[string]interface{}{
							"ipv6cidr_block":  "2607:f590:0000:2200::/64",
							"byoipv6range_id": "range-c",
						},
					},
				},
			},
			args: args{
				configRaw: map[string]interface{}{
					"byoipv6cidr_details": []interface{}{
						map[string]interface{}{
							"ipv6cidr_block":  "2607:f590:0000:0200::/64",
							"byoipv6range_id": "range-a",
						},
						map[string]interface{}{
							"ipv6cidr_block":  "2607:f590:0000:2000::/64",
							"byoipv6range_id": "range-b",
						},
						map[string]interface{}{
							"ipv6cidr_block":  "2607:f590:0000:2200::/64",
							"byoipv6range_id": "range-c",
						},
					},
				},
			},
			wantBaselineChanges: true,
			wantSuppressed:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baselineDiff := buildByoipv6DetailsDiff(t, nil, tt.fields.stateRaw, tt.args.configRaw)
			if got := diffHasByoipv6CidrDetailsChanges(baselineDiff); got != tt.wantBaselineChanges {
				t.Fatalf("baseline diff changes = %v, want %v; diff = %#v", got, tt.wantBaselineChanges, baselineDiff.Attributes)
			}

			suppressedDiff := buildByoipv6DetailsDiff(t, suppressMatchingByoipv6CidrDetailsDiff, tt.fields.stateRaw, tt.args.configRaw)
			if got := !diffHasByoipv6CidrDetailsChanges(suppressedDiff); got != tt.wantSuppressed {
				t.Fatalf("suppressed diff state = %v, want %v; diff = %#v", got, tt.wantSuppressed, suppressedDiff.Attributes)
			}

			if tt.name == "reorders middle-entry drift into append-style add" {
				if !diffHasByoipv6CidrDetailsIndexChanges(baselineDiff, 1) {
					t.Fatalf("baseline diff should show a middle-entry change; diff = %#v", baselineDiff.Attributes)
				}
				if diffHasByoipv6CidrDetailsIndexChanges(suppressedDiff, 1) {
					t.Fatalf("suppressed diff should not show a middle-entry change; diff = %#v", suppressedDiff.Attributes)
				}
				if !diffHasByoipv6CidrDetailsIndexChanges(suppressedDiff, 2) {
					t.Fatalf("suppressed diff should preserve only the append-style add; diff = %#v", suppressedDiff.Attributes)
				}
			}
		})
	}
}

func TestSuppressMatchingByoipv6CidrDetailsDiff_AfterRefreshDriftRemoval(t *testing.T) {
	resource := &schema.Resource{
		Schema:        MockedCoreVcnResourceSchema(),
		CustomizeDiff: suppressMatchingByoipv6CidrDetailsDiff,
	}

	refreshedData := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"byoipv6cidr_details": []interface{}{
			map[string]interface{}{
				"ipv6cidr_block":  "2607:f590:0000:0200::/64",
				"byoipv6range_id": "range-a",
			},
			map[string]interface{}{
				"ipv6cidr_block":  "2607:f590:0000:2000::/64",
				"byoipv6range_id": "range-b",
			},
			map[string]interface{}{
				"ipv6cidr_block":  "2607:f590:0000:2200::/64",
				"byoipv6range_id": "range-c",
			},
		},
	})
	refreshedData.SetId("ocid1.vcn.oc1..exampleuniqueID")

	crud := &CoreVcnResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: refreshedData},
	}

	// Simulate external removal of the middle block followed by provider refresh.
	crud.setBYOIPv6Details([]interface{}{
		"2607:f590:0000:0200::/64",
		"2607:f590:0000:2200::/64",
	})

	refreshedState := refreshedData.State()

	diff, err := resource.Diff(context.Background(), refreshedState, terraform.NewResourceConfigRaw(map[string]interface{}{
		"byoipv6cidr_details": []interface{}{
			map[string]interface{}{
				"ipv6cidr_block":  "2607:f590:0000:0200::/64",
				"byoipv6range_id": "range-a",
			},
			map[string]interface{}{
				"ipv6cidr_block":  "2607:f590:0000:2000::/64",
				"byoipv6range_id": "range-b",
			},
			map[string]interface{}{
				"ipv6cidr_block":  "2607:f590:0000:2200::/64",
				"byoipv6range_id": "range-c",
			},
		},
	}), nil)
	if err != nil {
		t.Fatalf("Resource.Diff() error = %v", err)
	}

	if diffHasByoipv6CidrDetailsIndexChanges(diff, 1) {
		t.Fatalf("refreshed diff should not show a middle-entry change after suppression; diff = %#v", diff.Attributes)
	}
	if !diffHasByoipv6CidrDetailsIndexChanges(diff, 2) {
		t.Fatalf("refreshed diff should preserve only the append-style add; diff = %#v", diff.Attributes)
	}
}
