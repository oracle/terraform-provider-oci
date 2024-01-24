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
	FusionAppsFusionEnvironmentDataMaskingActivityRequiredOnlyResource = FusionAppsFusionEnvironmentDataMaskingActivityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_data_masking_activity", "test_fusion_environment_data_masking_activity", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentDataMaskingActivityRepresentation)

	FusionAppsFusionEnvironmentDataMaskingActivityResourceConfig = FusionAppsFusionEnvironmentDataMaskingActivityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_data_masking_activity", "test_fusion_environment_data_masking_activity", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentDataMaskingActivityRepresentation)

	FusionAppsFusionAppsFusionEnvironmentDataMaskingActivitySingularDataSourceRepresentation = map[string]interface{}{
		"data_masking_activity_id": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"fusion_environment_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
	}

	FusionAppsFusionAppsFusionEnvironmentDataMaskingActivityDataSourceRepresentation = map[string]interface{}{
		"fusion_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: FusionAppsFusionEnvironmentDataMaskingActivityDataSourceFilterRepresentation}}
	FusionAppsFusionEnvironmentDataMaskingActivityDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fusion_apps_fusion_environment_data_masking_activity.test_fusion_environment_data_masking_activity.id}`}},
	}

	FusionAppsFusionEnvironmentDataMaskingActivityRepresentation = map[string]interface{}{
		"fusion_environment_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_fusion_apps_fusion_environment.test_fusion_environment.id}`},
		"is_resume_data_masking": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	FusionAppsFusionEnvironmentDataMaskingActivityResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_family", "test_fusion_environment_family", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentFamilyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment", "test_fusion_environment", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentRepresentation)
)

// issue-routing-tag: fusion_apps/default
func TestFusionAppsFusionEnvironmentDataMaskingActivityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFusionAppsFusionEnvironmentDataMaskingActivityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fusion_apps_fusion_environment_data_masking_activity.test_fusion_environment_data_masking_activity"
	datasourceName := "data.oci_fusion_apps_fusion_environment_data_masking_activities.test_fusion_environment_data_masking_activities"
	singularDatasourceName := "data.oci_fusion_apps_fusion_environment_data_masking_activity.test_fusion_environment_data_masking_activity"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FusionAppsFusionEnvironmentDataMaskingActivityResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_data_masking_activity", "test_fusion_environment_data_masking_activity", acctest.Optional, acctest.Create, FusionAppsFusionEnvironmentDataMaskingActivityRepresentation), "fusionapps", "fusionEnvironmentDataMaskingActivity", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentDataMaskingActivityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_data_masking_activity", "test_fusion_environment_data_masking_activity", acctest.Required, acctest.Create, FusionAppsFusionEnvironmentDataMaskingActivityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "fusion_environment_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentDataMaskingActivityResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FusionAppsFusionEnvironmentDataMaskingActivityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_data_masking_activity", "test_fusion_environment_data_masking_activity", acctest.Optional, acctest.Create, FusionAppsFusionEnvironmentDataMaskingActivityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_resume_data_masking", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_masking_finish"),
				resource.TestCheckResourceAttrSet(resourceName, "time_masking_start"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_data_masking_activities", "test_fusion_environment_data_masking_activities", acctest.Optional, acctest.Update, FusionAppsFusionAppsFusionEnvironmentDataMaskingActivityDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentDataMaskingActivityResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fusion_apps_fusion_environment_data_masking_activity", "test_fusion_environment_data_masking_activity", acctest.Optional, acctest.Update, FusionAppsFusionEnvironmentDataMaskingActivityRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fusion_environment_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "data_masking_activity_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "data_masking_activity_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fusion_apps_fusion_environment_data_masking_activity", "test_fusion_environment_data_masking_activity", acctest.Required, acctest.Create, FusionAppsFusionAppsFusionEnvironmentDataMaskingActivitySingularDataSourceRepresentation) +
				compartmentIdVariableStr + FusionAppsFusionEnvironmentDataMaskingActivityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_masking_activity_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fusion_environment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_masking_finish"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_masking_start"),
			),
		},
		// verify resource import
		{
			Config:            config + FusionAppsFusionEnvironmentDataMaskingActivityRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_resume_data_masking",
			},
			ResourceName: resourceName,
		},
	})
}
