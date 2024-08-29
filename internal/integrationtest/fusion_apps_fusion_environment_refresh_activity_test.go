// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FusionAppsFusionEnvironmentRefreshActivityResourceConfig = FusionAppsFusionEnvironmentRefreshActivityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_refresh_activity", "test_fusion_environment_refresh_activity", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentRefreshActivityRepresentation)

	FusionAppsFusionAppsFusionEnvironmentRefreshActivitySingularDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"refresh_activity_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment_refresh_activity.test_fusion_environment_refresh_activity.id}`},
	}

	FusionAppsFusionAppsFusionEnvironmentRefreshActivityDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`}}

	FusionAppsFusionEnvironmentRefreshActivityRepresentation = map[string]interface{}{
		"fusion_environment_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"source_fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"is_data_masking_opted":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	FusionAppsFusionEnvironmentRefreshActivityResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentRepresentation)
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentRefreshActivityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentRefreshActivityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fusion_apps_fusion_environment_refresh_activity.test_fusion_environment_refresh_activity"
	datasourceName := "data.oci_fusion_apps_fusion_environment_refresh_activities.test_fusion_environment_refresh_activities"
	singularDatasourceName := "data.oci_fusion_apps_fusion_environment_refresh_activity.test_fusion_environment_refresh_activity"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FusionAppsFusionEnvironmentRefreshActivityResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_refresh_activity", "test_fusion_environment_refresh_activity", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentRefreshActivityRepresentation), "fusionapps", "fusionEnvironmentRefreshActivity", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentRefreshActivityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_refresh_activity", "test_fusion_environment_refresh_activity", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentRefreshActivityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_fusion_environment_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_refresh_activities", "test_fusion_environment_refresh_activities", acctest.Optional, acctest.Update, FusionAppsFusionAppsFusionEnvironmentRefreshActivityDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentRefreshActivityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_refresh_activity", "test_fusion_environment_refresh_activity", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentRefreshActivityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fusion_environment_id"),

				resource.TestCheckResourceAttr(datasourceName, "refresh_activity_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "refresh_activity_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_refresh_activity", "test_fusion_environment_refresh_activity", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentRefreshActivitySingularDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentRefreshActivityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "refresh_activity_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_data_masking_opted", "false"),
				// if there's no refresh issue, this field will be null, so commented out this line
				// resource.TestCheckResourceAttr(singularDatasourceName, "refresh_issue_details_list.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_availability"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_expected_finish"),
				// Refresh resource is a scheduled activity, the create API is to create the activity.
				// This field will only be set when the actual refresh finished, commented out
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),

				// This field will only be set when the actual refresh finished, commented out
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_restoration_point"),

				// This field will only be set when the actual refresh finished, commented out
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
