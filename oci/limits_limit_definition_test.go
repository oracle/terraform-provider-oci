// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	limitDefinitionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"name":           Representation{repType: Optional, create: `custom-image-count`},
		"service_name":   Representation{repType: Optional, create: `${data.oci_limits_services.test_services.services.0.name}`},
	}

	LimitDefinitionResourceConfig = generateDataSourceFromRepresentationMap("oci_limits_services", "test_services", Required, Create, limitsServiceDataSourceRepresentation)
)

func TestLimitsLimitDefinitionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsLimitDefinitionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	datasourceName := "data.oci_limits_limit_definitions.test_limit_definitions"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_limits_limit_definitions", "test_limit_definitions", Required, Create, limitDefinitionDataSourceRepresentation) +
					compartmentIdVariableStr + LimitDefinitionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.are_quotas_supported"),
					resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_deprecated"),
					resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_dynamic"),
					resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_eligible_for_limit_increase"),
					resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_resource_availability_supported"),
					resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.scope_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.service_name"),
				),
			},
		},
	})
}
