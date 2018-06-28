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

`
	InstanceCredentialPropertyVariables = `

`
	InstanceCredentialResourceDependencies = ""
)

func TestCoreInstanceCredentialResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_instance_credentials.test_instance_credentials"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `

data "oci_core_instance_credentials" "test_instance_credentials" {
	#Required
	instance_id = "${oci_core_instance.test_instance.id}"
}
                ` + compartmentIdVariableStr + InstanceCredentialResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "password"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "username"),
				),
			},
		},
	})
}
