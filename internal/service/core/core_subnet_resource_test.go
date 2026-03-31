package core

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			name: "Set ipv6cidr_block from first ipv6cidr_blocks entry",
			subnet: oci_core.Subnet{
				Ipv6CidrBlocks: []string{
					"2607:9b80:9a0c:ac00:0000:0000:0000:0000/64",
					"fc00:1000:0000:0000:0000:0000:0000:0000/64",
				},
			},
			expectedIpv6Cidr: "2607:9b80:9a0c:ac00:0000:0000:0000:0000/64",
		},
		{
			name: "Prefer ipv6cidr_blocks over ipv6cidr_block",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock: &singular,
				Ipv6CidrBlocks: []string{
					"fc00:1000:0000:0000:0000:0000:0000:0000/64",
					"fc00:1001:0000:0000:0000:0000:0000:0000/64",
				},
			},
			expectedIpv6Cidr: "2607:9b80:9a0c:ac00::/64",
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
			name: "Trim spaces when comparing values - Prefer ipv6cidr_block when it has a value",
			subnet: oci_core.Subnet{
				Ipv6CidrBlock: &singularWithSpace,
				Ipv6CidrBlocks: []string{
					"fc00:1000:0000:0000:0000:0000:0000:0000/64",
					"fc00:1001:0000:0000:0000:0000:0000:0000/64",
				},
			},
			expectedIpv6Cidr: singularWithSpace,
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
			expectedIpv6Cidr: "fc00:1000:0000:0000:0000:0000:0000:0000/64",
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
			expectedIpv6Cidr: "fc00:1000:0000:0000:0000:0000:0000:0000/64",
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
