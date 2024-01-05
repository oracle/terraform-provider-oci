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
	FusionAppsFusionAppsFusionEnvironmentScheduledActivitySingularDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"scheduled_activity_id": acctest.Representation{RepType: acctest.Required, Create: `{}`},
	}

	FusionAppsFusionAppsFusionEnvironmentScheduledActivityDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"run_cycle":             acctest.Representation{RepType: acctest.Optional, Create: `QUARTERLY`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"time_expected_finish_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `timeExpectedFinishLessThanOrEqualTo`},
		"time_scheduled_start_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeScheduledStartGreaterThanOrEqualTo`},
	}

	FusionAppsFusionEnvironmentScheduledActivityResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentRepresentation)
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentScheduledActivityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentScheduledActivityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_fusion_apps_fusion_environment_scheduled_activities.test_fusion_environment_scheduled_activities"
	singularDatasourceName := "data.oci_fusion_apps_fusion_environment_scheduled_activity.test_fusion_environment_scheduled_activity"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_scheduled_activities", "test_fusion_environment_scheduled_activities", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentScheduledActivityDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentScheduledActivityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttr(datasourceName, "run_cycle", "QUARTERLY"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_expected_finish_less_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_scheduled_start_greater_than_or_equal_to"),

				resource.TestCheckResourceAttrSet(datasourceName, "scheduled_activity_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_scheduled_activity", "test_fusion_environment_scheduled_activity", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentScheduledActivitySingularDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentScheduledActivityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_activity_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "actions.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "delay_in_hours"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_cycle"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_availability"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_expected_finish"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_scheduled_start"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
