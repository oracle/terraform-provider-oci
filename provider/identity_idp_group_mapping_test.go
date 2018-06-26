// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	IdpGroupMappingResourceConfig = IdpGroupMappingResourceDependencies + `
resource "oci_identity_idp_group_mapping" "test_idp_group_mapping" {
	#Required
	group_id = "${oci_identity_group.test_group.id}"
	identity_provider_id = "${oci_identity_identity_provider.test_identity_provider.id}"
	idp_group_name = "${var.idp_group_mapping_idp_group_name}"
}
`
	IdpGroupMappingPropertyVariables = `
variable "idp_group_mapping_idp_group_name" { default = "idpGroupName" }

`
	IdpGroupMappingResourceDependencies = GroupPropertyVariables + GroupResourceConfig + IdentityProviderPropertyVariables + IdentityProviderRequiredOnlyResource
)

func TestIdentityIdpGroupMappingResource_basic(t *testing.T) {
	metadataFile := getEnvSetting("identity_provider_metadata_file", "")
	if metadataFile == "" {
		t.Skip("Skipping generated test for now as it has a dependency on federation metadata file")
	}

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_idp_group_mapping.test_idp_group_mapping"
	datasourceName := "data.oci_identity_idp_group_mappings.test_idp_group_mappings"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + IdpGroupMappingPropertyVariables + compartmentIdVariableStr + IdpGroupMappingResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "identity_provider_id"),
					resource.TestCheckResourceAttr(resourceName, "idp_group_name", "idpGroupName"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "idp_group_mapping_idp_group_name" { default = "idpGroupName2" }

                ` + compartmentIdVariableStr + IdpGroupMappingResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "identity_provider_id"),
					resource.TestCheckResourceAttr(resourceName, "idp_group_name", "idpGroupName2"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "idp_group_mapping_idp_group_name" { default = "idpGroupName2" }

data "oci_identity_idp_group_mappings" "test_idp_group_mappings" {
	#Required
	identity_provider_id = "${oci_identity_identity_provider.test_identity_provider.id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_idp_group_mapping.test_idp_group_mapping.id}"]
    }
}
                ` + compartmentIdVariableStr + IdpGroupMappingResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					//resource.TestCheckResourceAttrSet(datasourceName, "group_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "identity_provider_id"),

					resource.TestCheckResourceAttr(datasourceName, "idp_group_mappings.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.group_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.identity_provider_id"),
					resource.TestCheckResourceAttr(datasourceName, "idp_group_mappings.0.idp_group_name", "idpGroupName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "idp_group_mappings.0.time_created"),
				),
			},
		},
	})
}
