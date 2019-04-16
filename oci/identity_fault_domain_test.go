// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	faultDomainDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
	}

	FaultDomainResourceConfig = ""
)

func TestIdentityFaultDomainResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityFaultDomainResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_fault_domains.test_fault_domains"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + AvailabilityDomainConfig +
					generateDataSourceFromRepresentationMap("oci_identity_fault_domains", "test_fault_domains", Required, Create, faultDomainDataSourceRepresentation) +
					compartmentIdVariableStr + FaultDomainResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(datasourceName, "availability_domain", regexp.MustCompile(`\w+-AD-\d+`)),
					resource.TestMatchResourceAttr(datasourceName, "compartment_id", regexp.MustCompile(`.*?(tenancy|compartment).*?`)),

					resource.TestCheckResourceAttr(datasourceName, "fault_domains.#", "3"),
					resource.TestCheckResourceAttrSet(datasourceName, "fault_domains.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "fault_domains.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "fault_domains.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "fault_domains.0.name"),
					resource.TestMatchResourceAttr(datasourceName, "fault_domains.0.availability_domain", regexp.MustCompile(`\w+-AD-\d+`)),
					resource.TestMatchResourceAttr(datasourceName, "fault_domains.0.compartment_id", regexp.MustCompile(`.*?(tenancy|compartment).*?`)),
					resource.TestMatchResourceAttr(datasourceName, "fault_domains.0.id", regexp.MustCompile(`.*?faultdomain.*?`)),
					resource.TestCheckResourceAttr(datasourceName, "fault_domains.0.name", "FAULT-DOMAIN-1"),
					resource.TestCheckResourceAttr(datasourceName, "fault_domains.1.name", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(datasourceName, "fault_domains.2.name", "FAULT-DOMAIN-3"),
				),
			},
		},
	})
}
