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
	FusionAppsFusionAppsFusionEnvironmentFamilyLimitsAndUsageSingularDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_family_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id}`},
	}

	FusionAppsFusionEnvironmentFamilyLimitsAndUsageResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation)
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentFamilyLimitsAndUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentFamilyLimitsAndUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_fusion_apps_fusion_environment_family_limits_and_usage.test_fusion_environment_family_limits_and_usage"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family_limits_and_usage", "test_fusion_environment_family_limits_and_usage", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentFamilyLimitsAndUsageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentFamilyLimitsAndUsageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_family_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "development_limit_and_usage.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "production_limit_and_usage.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_limit_and_usage.#", "1"),
			),
		},
	})
}
