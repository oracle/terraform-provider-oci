// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"regexp"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	FaultDomainResourceConfig = FaultDomainResourceDependencies + `

`
	FaultDomainPropertyVariables = `

`
	FaultDomainResourceDependencies = AvailabilityDomainConfig
)

func TestIdentityFaultDomainResource_basic(t *testing.T) {
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
				Config: config + `

data "oci_identity_fault_domains" "test_fault_domains" {
	#Required
	availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
	compartment_id = "${var.compartment_id}"
}
                ` + compartmentIdVariableStr + FaultDomainResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(datasourceName, "availability_domain", regexp.MustCompile(`\w+-AD-\d+`)),
					resource.TestMatchResourceAttr(datasourceName, "compartment_id", regexp.MustCompile(`.*?(tenancy|compartment).*?`)),
					resource.TestCheckResourceAttr(datasourceName, "fault_domains.#", "3"), // more could be added in the future
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
