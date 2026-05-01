package core

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func buildCoreSubnetResourceCrudForSetDataTest(t *testing.T, subnet oci_core.Subnet) *CoreSubnetResourceCrud {
	t.Helper()

	d := schema.TestResourceDataRaw(t, CoreSubnetResource().Schema, map[string]interface{}{})

	return &CoreSubnetResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: d},
		Res:      &subnet,
	}
}

func buildCoreSubnetResourceCrudForUpdateRequestTest(t *testing.T, raw map[string]interface{}) *CoreSubnetResourceCrud {
	t.Helper()

	d := schema.TestResourceDataRaw(t, CoreSubnetResource().Schema, raw)
	d.SetId("ocid1.subnet.oc1..exampleuniqueID")

	return &CoreSubnetResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: d},
	}
}

func buildCoreSubnetResourceCrudForExplicitDiffTest(t *testing.T, stateAttrs map[string]string, diffAttrs map[string]*terraform.ResourceAttrDiff) *CoreSubnetResourceCrud {
	t.Helper()

	if stateAttrs == nil {
		stateAttrs = map[string]string{}
	}

	schemaMap := schema.InternalMap(CoreSubnetResource().Schema)
	data, err := schemaMap.Data(&terraform.InstanceState{
		ID:         "ocid1.subnet.oc1..exampleuniqueID",
		Attributes: stateAttrs,
	}, &terraform.InstanceDiff{
		Attributes: diffAttrs,
	})
	if err != nil {
		t.Fatalf("schemaMap.Data() error = %v", err)
	}

	return &CoreSubnetResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: data},
	}
}

func TestCoreSubnetResourceCrudSetData(t *testing.T) {
	empty := ""
	singular := "2607:9b80:9a0c:ac00::/64"
	whiteSpace := " "
	singularWithSpace := singular + whiteSpace
	multipleWithSpace := "    \t   "
	tests := []struct {
		name             string
		subnet           oci_core.Subnet
		expectedIpv6Cidr string
	}{
		{
			name: "Set ipv6cidr_block from last ipv6cidr_blocks entry",
			subnet: oci_core.Subnet{
				Ipv6CidrBlocks: []string{
					"2607:9b80:9a0c:ac00:0000:0000:0000:0000/64",
					"fc00:1000:0000:0000:0000:0000:0000:0000/64",
				},
			},
			expectedIpv6Cidr: "fc00:1000:0000:0000:0000:0000:0000:0000/64",
		},
		{
			name: "Set ipv6cidr_block from last ipv6cidr_blocks entry when scalar is not in list",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock: &singular,
				Ipv6CidrBlocks: []string{
					"fc00:1000:0000:0000:0000:0000:0000:0000/64",
					"fc00:1001:0000:0000:0000:0000:0000:0000/64",
				},
			},
			expectedIpv6Cidr: "fc00:1001:0000:0000:0000:0000:0000:0000/64",
		},
		{
			name: "Keep ipv6cidr_block when it is present in ipv6cidr_blocks",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock: &singular,
				Ipv6CidrBlocks: []string{
					"fc00:1000:0000:0000:0000:0000:0000:0000/64",
					singular,
				},
			},
			expectedIpv6Cidr: singular,
		},
		{
			name: "Keep ipv6cidr_block when equivalent value is present in ipv6cidr_blocks",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock: &singular,
				Ipv6CidrBlocks: []string{
					"fc00:1000:0000:0000:0000:0000:0000:0000/64",
					"2607:9b80:9a0c:ac00:0000:0000:0000:0000/64",
				},
			},
			expectedIpv6Cidr: singular,
		},
		{
			name: "Keep ipv6cidr_block when ipv6cidr_blocks is empty",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock:  &singular,
				Ipv6CidrBlocks: []string{},
			},
			expectedIpv6Cidr: singular,
		},
		{
			name: "Keep ipv6cidr_block when ipv6cidr_blocks nil",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock:  &singular,
				Ipv6CidrBlocks: nil,
			},
			expectedIpv6Cidr: singular,
		},
		{
			name: "Match empty string when ipv6cidr_block is empty and ipv6cidr_blocks is nil",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock:  &empty,
				Ipv6CidrBlocks: nil,
			},
			expectedIpv6Cidr: "",
		},
		{
			name: "Trim spaces when comparing values - Set ipv6cidr_block from last ipv6cidr_blocks entry when scalar is not in list",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock: &singularWithSpace,
				Ipv6CidrBlocks: []string{
					"fc00:1000:0000:0000:0000:0000:0000:0000/64",
					"fc00:1001:0000:0000:0000:0000:0000:0000/64",
				},
			},
			expectedIpv6Cidr: "fc00:1001:0000:0000:0000:0000:0000:0000/64",
		},
		{
			name: "Trim spaces when comparing values - Prefer ipv6cidr_blocks when ipv6cidr_block is just a space",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock: &whiteSpace,
				Ipv6CidrBlocks: []string{
					"fc00:1000:0000:0000:0000:0000:0000:0000/64",
					"fc00:1001:0000:0000:0000:0000:0000:0000/64",
				},
			},
			expectedIpv6Cidr: "fc00:1001:0000:0000:0000:0000:0000:0000/64",
		},
		{
			name: "Trim spaces when comparing values - Prefer ipv6cidr_blocks when ipv6cidr_block is a string of white spaces",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock: &multipleWithSpace,
				Ipv6CidrBlocks: []string{
					"fc00:1000:0000:0000:0000:0000:0000:0000/64",
					"fc00:1001:0000:0000:0000:0000:0000:0000/64",
				},
			},
			expectedIpv6Cidr: "fc00:1001:0000:0000:0000:0000:0000:0000/64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crud := buildCoreSubnetResourceCrudForSetDataTest(t, tt.subnet)

			if err := crud.SetData(); err != nil {
				t.Fatalf("SetData() returned error: %v", err)
			}

			if got := crud.D.Get("ipv6cidr_block").(string); got != tt.expectedIpv6Cidr {
				t.Fatalf("expected ipv6cidr_block to be %q, got %q", tt.expectedIpv6Cidr, got)
			}
		})
	}
}

func TestIpv4CidrOneEditAway(t *testing.T) {
	tests := []struct {
		name      string
		oldBlocks []string
		newBlocks []string
		want      []ipv4CidrEdit
		wantErr   bool
	}{
		{
			name:      "single modify",
			oldBlocks: []string{"10.0.0.0/24"},
			newBlocks: []string{"10.0.0.0/25"},
			want:      []ipv4CidrEdit{{operation: "modify", oldCidr: "10.0.0.0/24", newCidr: "10.0.0.0/25"}},
		},
		{
			name:      "single add",
			oldBlocks: []string{"10.0.0.0/24"},
			newBlocks: []string{"10.0.0.0/24", "10.0.1.0/24"},
			want:      []ipv4CidrEdit{{operation: "add", oldCidr: "", newCidr: "10.0.1.0/24"}},
		},
		{
			name:      "duplicate invalid",
			oldBlocks: []string{"10.0.0.0/24"},
			newBlocks: []string{"10.0.1.0/24", "10.0.1.0/24"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ipv4CidrOneEditAway(tt.oldBlocks, tt.newBlocks)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ipv4CidrOneEditAway() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("ipv4CidrOneEditAway() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestIpv6CidrOneEditAway(t *testing.T) {
	tests := []struct {
		name      string
		oldBlocks []string
		newBlocks []string
		wantOK    bool
		wantOp    string
		wantCidr  string
	}{
		{
			name:      "single add",
			oldBlocks: []string{"2001:db8::/64"},
			newBlocks: []string{"2001:db8::/64", "2001:db8:1::/64"},
			wantOK:    true,
			wantOp:    "add",
			wantCidr:  "2001:db8:1::/64",
		},
		{
			name:      "single remove",
			oldBlocks: []string{"2001:db8::/64", "2001:db8:1::/64"},
			newBlocks: []string{"2001:db8::/64"},
			wantOK:    true,
			wantOp:    "remove",
			wantCidr:  "2001:db8:1::/64",
		},
		{
			name:      "single replacement is not supported by legacy helper",
			oldBlocks: []string{"2001:db8::/64"},
			newBlocks: []string{"2001:db8:1::/64"},
			wantOK:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOK, gotOp, gotCidr := ipv6CidrOneEditAway(tt.oldBlocks, tt.newBlocks)
			if gotOK != tt.wantOK || gotOp != tt.wantOp || gotCidr != tt.wantCidr {
				t.Fatalf("ipv6CidrOneEditAway() = (%v, %q, %q), want (%v, %q, %q)", gotOK, gotOp, gotCidr, tt.wantOK, tt.wantOp, tt.wantCidr)
			}
		})
	}
}

func TestCoreSubnetResourceCrudBuildUpdateSubnetRequest(t *testing.T) {
	t.Run("uses explicit diff values for cidr_block and other mutable fields", func(t *testing.T) {
		crud := buildCoreSubnetResourceCrudForExplicitDiffTest(t,
			map[string]string{
				"cidr_block":     "10.0.0.0/24",
				"display_name":   "old-name",
				"route_table_id": "ocid1.routetable.oc1..old",
			},
			map[string]*terraform.ResourceAttrDiff{
				"cidr_block": {
					Old: "10.0.0.0/24",
					New: "10.0.1.0/24",
				},
				"display_name": {
					Old: "old-name",
					New: "new-name",
				},
				"route_table_id": {
					Old: "ocid1.routetable.oc1..old",
					New: "",
				},
			},
		)

		request, shouldUpdateSubnet, err := crud.buildUpdateSubnetRequest(subnetIpv6PatchChangeSet{}, false)
		if err != nil {
			t.Fatalf("buildUpdateSubnetRequest() error = %v", err)
		}
		if !shouldUpdateSubnet {
			t.Fatalf("buildUpdateSubnetRequest() shouldUpdateSubnet = false, want true")
		}
		if request.CidrBlock == nil || *request.CidrBlock != "10.0.1.0/24" {
			t.Fatalf("request.CidrBlock = %v, want 10.0.1.0/24", request.CidrBlock)
		}
		if request.DisplayName == nil || *request.DisplayName != "new-name" {
			t.Fatalf("request.DisplayName = %v, want new-name", request.DisplayName)
		}
		if request.RouteTableId == nil || *request.RouteTableId != "" {
			t.Fatalf("request.RouteTableId = %v, want empty string pointer", request.RouteTableId)
		}
		if request.SubnetId == nil || *request.SubnetId != crud.D.Id() {
			t.Fatalf("request.SubnetId = %v, want %q", request.SubnetId, crud.D.Id())
		}
	})

	t.Run("includes ipv6cidr_block only when update subnet still owns the add", func(t *testing.T) {
		crud := buildCoreSubnetResourceCrudForUpdateRequestTest(t, map[string]interface{}{
			"compartment_id": "ocid1.compartment.oc1..exampleuniqueID",
			"vcn_id":         "ocid1.vcn.oc1..exampleuniqueID",
		})
		changeSet := subnetIpv6PatchChangeSet{
			ipv6CidrBlockChanged: true,
			oldIpv6CidrBlock:     "",
			newIpv6CidrBlock:     "2001:db8::/64",
		}

		request, shouldUpdateSubnet, err := crud.buildUpdateSubnetRequest(changeSet, false)
		if err != nil {
			t.Fatalf("buildUpdateSubnetRequest() error = %v", err)
		}
		if !shouldUpdateSubnet {
			t.Fatalf("buildUpdateSubnetRequest() shouldUpdateSubnet = false, want true")
		}
		if request.Ipv6CidrBlock == nil || *request.Ipv6CidrBlock != "2001:db8::/64" {
			t.Fatalf("request.Ipv6CidrBlock = %v, want 2001:db8::/64", request.Ipv6CidrBlock)
		}

		request, shouldUpdateSubnet, err = crud.buildUpdateSubnetRequest(changeSet, true)
		if err != nil {
			t.Fatalf("buildUpdateSubnetRequest() error = %v", err)
		}
		if shouldUpdateSubnet {
			t.Fatalf("buildUpdateSubnetRequest() shouldUpdateSubnet = true, want false when patch owns ipv6 changes")
		}
		if request.Ipv6CidrBlock != nil {
			t.Fatalf("request.Ipv6CidrBlock = %v, want nil when patch owns ipv6 changes", request.Ipv6CidrBlock)
		}
	})
}
