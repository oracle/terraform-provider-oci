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
	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRequiredOnlyResource = OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_promote_software_source_management", "test_lifecycle_stage_promote_software_source_management", acctest.Required, acctest.Create, OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation)

	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation = map[string]interface{}{
		"lifecycle_stage_id":   acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("osmh_promote_software_source_test_lifecycle_stage_1")},
		"software_source_id":   acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("osmh_test_versioned_custom_software_source_2")},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubLifecycleStagePromoteSoftwareSourceManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation2 = map[string]interface{}{
		"lifecycle_stage_id":   acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("osmh_promote_software_source_test_lifecycle_stage_2")},
		"software_source_id":   acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("osmh_test_versioned_custom_software_source_2")},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubLifecycleStagePromoteSoftwareSourceManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation3 = map[string]interface{}{
		"lifecycle_stage_id":   acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("osmh_promote_software_source_test_lifecycle_stage_1")},
		"software_source_id":   acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("osmh_test_versioned_custom_software_source_1")},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubLifecycleStagePromoteSoftwareSourceManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation4 = map[string]interface{}{
		"lifecycle_stage_id":   acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("osmh_promote_software_source_test_lifecycle_stage_2")},
		"software_source_id":   acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("osmh_test_versioned_custom_software_source_1")},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubLifecycleStagePromoteSoftwareSourceManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubLifecycleStagePromoteSoftwareSourceManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubLifecycleStagePromoteSoftwareSourceManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_lifecycle_stage_promote_software_source_management.test_lifecycle_stage_promote_software_source_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_promote_software_source_management", "test_lifecycle_stage_promote_software_source_management", acctest.Optional, acctest.Create, OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation), "osmanagementhub", "lifecycleStagePromoteSoftwareSourceManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// initial state before running test: stage1 (VCSS1), stage2 (VCSS1)
		// promote VCSS2 to stage1 (rank 1)
		// after promotion: stage1 (VCSS2), stage2 (VCSS1)
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_promote_software_source_management", "test_lifecycle_stage_promote_software_source_management", acctest.Optional, acctest.Create, OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_id"),
				resource.TestCheckResourceAttr(resourceName, "software_source_id", utils.GetEnvSettingWithBlankDefault("osmh_test_versioned_custom_software_source_2")),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.display_name", "displayName"),

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
		// promote VCSS2 to stage2 (rank 2)
		// after promotion: stage1 (VCSS2), stage2 (VCSS2)
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_promote_software_source_management", "test_lifecycle_stage_promote_software_source_management", acctest.Optional, acctest.Create, OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_id"),
				resource.TestCheckResourceAttr(resourceName, "software_source_id", utils.GetEnvSettingWithBlankDefault("osmh_test_versioned_custom_software_source_2")),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.display_name", "displayName"),
			),
		},
		// promote VCSS1 back to stage1 (rank 1)
		// after promotion: stage1 (VCSS1), stage2 (VCSS2)
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_promote_software_source_management", "test_lifecycle_stage_promote_software_source_management", acctest.Optional, acctest.Create, OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_id"),
				resource.TestCheckResourceAttr(resourceName, "software_source_id", utils.GetEnvSettingWithBlankDefault("osmh_test_versioned_custom_software_source_1")),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.display_name", "displayName"),
			),
		},
		// promote VCSS1 back to stage2 (rank 2)
		// after promotion should restore the status before running the test: stage1 (VCSS1), stage2 (VCSS1)
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_promote_software_source_management", "test_lifecycle_stage_promote_software_source_management", acctest.Optional, acctest.Create, OsManagementHubLifecycleStagePromoteSoftwareSourceManagementRepresentation4),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_id"),
				resource.TestCheckResourceAttr(resourceName, "software_source_id", utils.GetEnvSettingWithBlankDefault("osmh_test_versioned_custom_software_source_1")),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.display_name", "displayName"),
			),
		},
	})
}
