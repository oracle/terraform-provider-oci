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
resource "oci_identity_availability_domain" "test_availability_domain" {
}
`
	AvailabilityDomainPropertyVariables = `

`
	AvailabilityDomainResourceDependencies = ""
)

func TestIdentityAvailabilityDomainResource_basic(t *testing.T) {
	t.Skip("Creating availability domain resource is not supported. Data source test is covered by legacy test. Need to enable data source only tests in generator. https://jira.aka.lgl.grungy.us/browse/ORCH-708")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_identity_availability_domain.test_availability_domain"
	datasourceName := "data.oci_identity_availability_domains.test_availability_domains"

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
				Config:            config + AvailabilityDomainPropertyVariables + compartmentIdVariableStr + AvailabilityDomainResourceConfig,
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

                ` + compartmentIdVariableStr2 + AvailabilityDomainResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

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

data "oci_identity_availability_domains" "test_availability_domains" {
	#Required
	compartment_id = "${var.compartment_id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_availability_domain.test_availability_domain.id}"]
    }
}
                ` + compartmentIdVariableStr2 + AvailabilityDomainResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),

					resource.TestCheckResourceAttr(datasourceName, "availability_domains.#", "1"),
				),
			},
		},
	})
}
