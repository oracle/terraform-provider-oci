// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LimitsLimitsLimitDefinitionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `custom-image-count`},
		"service_name":   acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_limits_services.test_services.services.0.name}`},
	}

	LimitsLimitDefinitionResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_limits_services", "test_services", acctest.Required, acctest.Create, LimitsLimitsServiceDataSourceRepresentation)
)

// issue-routing-tag: limits/default
func TestLimitsLimitDefinitionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsLimitDefinitionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	datasourceName := "data.oci_limits_limit_definitions.test_limit_definitions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_limit_definitions", "test_limit_definitions", acctest.Required, acctest.Create, LimitsLimitsLimitDefinitionDataSourceRepresentation) +
				compartmentIdVariableStr + LimitsLimitDefinitionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	})
}
