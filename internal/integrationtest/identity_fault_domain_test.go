// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityIdentityFaultDomainDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	IdentityFaultDomainResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: identity/default
func TestIdentityFaultDomainResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityFaultDomainResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_fault_domains.test_fault_domains"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_fault_domains", "test_fault_domains", acctest.Required, acctest.Create, IdentityIdentityFaultDomainDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityFaultDomainResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	})
}
