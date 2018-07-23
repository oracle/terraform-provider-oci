// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

const (
	GroupRequiredOnlyResource = GroupRequiredOnlyResourceDependencies + `
resource "oci_identity_group" "test_group" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.group_description}"
	name = "${var.group_name}"
}
`

	GroupResourceConfig = GroupResourceDependencies + `
resource "oci_identity_group" "test_group" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.group_description}"
	name = "${var.group_name}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.group_defined_tags_value}")}"
	freeform_tags = "${var.group_freeform_tags}"
}
`
	GroupPropertyVariables = `
variable "group_defined_tags_value" { default = "value" }
variable "group_description" { default = "Group for network administrators" }
variable "group_freeform_tags" { default = {"Department"= "Finance"} }
variable "group_name" { default = "NetworkAdmins" }

`
	GroupRequiredOnlyResourceDependencies = ``
	GroupResourceDependencies             = DefinedTagsDependencies
)

func TestIdentityGroupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getRequiredEnvSetting("tenancy_ocid")

	resourceName := "oci_identity_group.test_group"
	datasourceName := "data.oci_identity_groups.test_groups"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityGroupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + GroupPropertyVariables + compartmentIdVariableStr + GroupRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "Group for network administrators"),
					resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + GroupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + GroupPropertyVariables + compartmentIdVariableStr + GroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "Group for network administrators"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "group_defined_tags_value" { default = "updatedValue" }
variable "group_description" { default = "description2" }
variable "group_freeform_tags" { default = {"Department"= "Accounting"} }
variable "group_name" { default = "NetworkAdmins" }

                ` + compartmentIdVariableStr + GroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),
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
variable "group_defined_tags_value" { default = "updatedValue" }
variable "group_description" { default = "description2" }
variable "group_freeform_tags" { default = {"Department"= "Accounting"} }
variable "group_name" { default = "NetworkAdmins" }

data "oci_identity_groups" "test_groups" {
	#Required
	compartment_id = "${var.tenancy_ocid}"

    filter {
    	name = "id"
    	values = ["${oci_identity_group.test_group.id}"]
    }
}
                ` + compartmentIdVariableStr + GroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttr(datasourceName, "groups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "groups.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.name", "NetworkAdmins"),
					resource.TestCheckResourceAttrSet(datasourceName, "groups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "groups.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckIdentityGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_group" {
			noResourceFound = false
			request := oci_identity.GetGroupRequest{}

			tmp := rs.Primary.ID
			request.GroupId = &tmp

			response, err := client.GetGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.GroupLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
