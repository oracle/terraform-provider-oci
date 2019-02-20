// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	availabilityDomainDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
	}

	AvailabilityDomainResourceConfig = ""

	AvailabilityDomainConfig = generateDataSourceFromRepresentationMap("oci_identity_availability_domains", "test_availability_domains", Required, Create, availabilityDomainDataSourceRepresentation)
)

func TestIdentityAvailabilityDomainResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	datasourceName := "data.oci_identity_availability_domains.test_availability_domains"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_availability_domains", "test_availability_domains", Required, Create, availabilityDomainDataSourceRepresentation) +
					compartmentIdVariableStr + AvailabilityDomainResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttrSet(datasourceName, "availability_domains.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domains.0.name"),
				),
			},
		},
	})
}
