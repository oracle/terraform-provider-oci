// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	InstanceCredentialResourceConfig = InstanceCredentialResourceDependencies + `
resource "oci_core_instance_credential" "test_instance_credential" {
}
`
	InstanceCredentialPropertyVariables = `
variable "instance_credential_instance_id" { default = "instanceId" }

`
	InstanceCredentialResourceDependencies = ""
)

func TestCoreInstanceCredentialResource_basic(t *testing.T) {
	t.Skip("InstanceCredentials is a data source only artifact. Data source test is covered by legacy test. Need to enable data source only tests in generator. https://jira.aka.lgl.grungy.us/browse/ORCH-708")

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_instance_credential.test_instance_credential"
	datasourceName := "data.oci_core_instance_credentials.test_instance_credentials"

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
				Config:            config + InstanceCredentialPropertyVariables + compartmentIdVariableStr + InstanceCredentialResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `
variable "instance_credential_instance_id" { default = "instanceId2" }

                ` + compartmentIdVariableStr2 + InstanceCredentialResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "password"),
					resource.TestCheckResourceAttrSet(resourceName, "username"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "instance_credential_instance_id" { default = "instanceId2" }

data "oci_core_instance_credentials" "test_instance_credentials" {
	#Required
	instance_id = "${var.instance_credential_instance_id}"

    filter {
    	name = "id"
    	values = ["${oci_core_instance_credential.test_instance_credential.id}"]
    }
}
                ` + compartmentIdVariableStr2 + InstanceCredentialResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "instance_id", "instanceId2"),

					resource.TestCheckResourceAttr(datasourceName, "instance_credentials.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_credentials.0.password"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_credentials.0.username"),
				),
			},
		},
	})
}
