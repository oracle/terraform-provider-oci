// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	DrgRequiredOnlyResource = DrgRequiredOnlyResourceDependencies + `
resource "oci_core_drg" "test_drg" {
	#Required
	compartment_id = "${var.compartment_id}"
}
`

	DrgResourceConfig = DrgResourceDependencies + `
resource "oci_core_drg" "test_drg" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.drg_defined_tags_value}")}"
	display_name = "${var.drg_display_name}"
	freeform_tags = "${var.drg_freeform_tags}"
}
`
	DrgPropertyVariables = `
variable "drg_defined_tags_value" { default = "value" }
variable "drg_display_name" { default = "MyDrg" }
variable "drg_freeform_tags" { default = {"Department"= "Finance"} }

`
	DrgRequiredOnlyResourceDependencies = ``
	DrgResourceDependencies             = DefinedTagsDependencies
)

func TestCoreDrgResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg.test_drg"
	datasourceName := "data.oci_core_drgs.test_drgs"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreDrgDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + DrgPropertyVariables + compartmentIdVariableStr + DrgRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DrgResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + DrgPropertyVariables + compartmentIdVariableStr + DrgResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyDrg"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "drg_defined_tags_value" { default = "updatedValue" }
variable "drg_display_name" { default = "displayName2" }
variable "drg_freeform_tags" { default = {"Department"= "Accounting"} }

                ` + compartmentIdVariableStr + DrgResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
variable "drg_defined_tags_value" { default = "updatedValue" }
variable "drg_display_name" { default = "displayName2" }
variable "drg_freeform_tags" { default = {"Department"= "Accounting"} }

data "oci_core_drgs" "test_drgs" {
	#Required
	compartment_id = "${var.compartment_id}"

    filter {
    	name = "id"
    	values = ["${oci_core_drg.test_drg.id}"]
    }
}
                ` + compartmentIdVariableStr + DrgResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "drgs.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "drgs.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "drgs.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "drgs.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "drgs.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "drgs.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drgs.0.state"),
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

func testAccCheckCoreDrgDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_drg" {
			noResourceFound = false
			request := oci_core.GetDrgRequest{}

			tmp := rs.Primary.ID
			request.DrgId = &tmp

			response, err := client.GetDrg(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.DrgLifecycleStateTerminated): true,
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
