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
	FusionAppsFusionAppsFusionEnvironmentTimeAvailableForRefreshSingularDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
	}

	FusionAppsFusionAppsFusionEnvironmentTimeAvailableForRefreshDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
	}

	FusionAppsFusionEnvironmentTimeAvailableForRefreshResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentRepresentation)
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentTimeAvailableForRefreshResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentTimeAvailableForRefreshResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_fusion_apps_fusion_environment_time_available_for_refreshs.test_fusion_environment_time_available_for_refreshs"
	singularDatasourceName := "data.oci_fusion_apps_fusion_environment_time_available_for_refresh.test_fusion_environment_time_available_for_refresh"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_time_available_for_refreshs", "test_fusion_environment_time_available_for_refreshs", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentTimeAvailableForRefreshDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentTimeAvailableForRefreshResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fusion_environment_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "time_available_for_refresh_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "time_available_for_refresh_collection.0.items.#", "84"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_time_available_for_refresh", "test_fusion_environment_time_available_for_refresh", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentTimeAvailableForRefreshSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentTimeAvailableForRefreshResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "84"),
			),
		},
	})
}
