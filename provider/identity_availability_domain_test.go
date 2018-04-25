// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	AvailabilityDomainResourceConfig = AvailabilityDomainResourceDependencies + `

`
	AvailabilityDomainPropertyVariables = `

`
	AvailabilityDomainResourceDependencies = ""
)

func TestIdentityAvailabilityDomainResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	datasourceName := "data.oci_identity_availability_domains.test_availability_domains"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_identity_availability_domains" "test_availability_domains" {
	#Required
	compartment_id = "${var.compartment_id}"
}
                ` + compartmentIdVariableStr2 + AvailabilityDomainResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),

					resource.TestCheckResourceAttrSet(datasourceName, "availability_domains.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domains.0.name"),
				),
			},
		},
	})
}
