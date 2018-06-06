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

	AvailabilityDomainConfig = `
data "oci_identity_availability_domains" "test_availability_domains" {
	compartment_id = "${var.compartment_id}"
}
`
)

func TestIdentityAvailabilityDomainResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

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
                ` + compartmentIdVariableStr + AvailabilityDomainResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "availability_domains.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domains.0.name"),
				),
			},
		},
	})
}
