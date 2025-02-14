// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubLifecycleStageRebootManagementRequiredOnlyResource = OsManagementHubLifecycleStageRebootManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_reboot_management", "test_lifecycle_stage_reboot_management", acctest.Required, acctest.Create, OsManagementHubLifecycleStageRebootManagementRepresentation)

	OsManagementHubLifecycleStageRebootManagementRepresentation = map[string]interface{}{
		"lifecycle_stage_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id}`},
		"reboot_timeout_in_mins": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"work_request_details":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubLifecycleStageRebootManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubLifecycleStageRebootManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubLifecycleStageRebootManagementResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_stages", "test_lifecycle_stages", acctest.Required, acctest.Create, OsManagementHubLifecycleStageDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage", "test_lifecycle_stage", acctest.Required, acctest.Create, OsManagementHubLifecycleStageSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Required, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubLifecycleStageRebootManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubLifecycleStageRebootManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_lifecycle_stage_reboot_management.test_lifecycle_stage_reboot_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubLifecycleStageRebootManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_reboot_management", "test_lifecycle_stage_reboot_management", acctest.Optional, acctest.Create, OsManagementHubLifecycleStageRebootManagementRepresentation), "osmanagementhub", "lifecycleStageRebootManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStageRebootManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_reboot_management", "test_lifecycle_stage_reboot_management", acctest.Required, acctest.Create, OsManagementHubLifecycleStageRebootManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStageRebootManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStageRebootManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_reboot_management", "test_lifecycle_stage_reboot_management", acctest.Optional, acctest.Create, OsManagementHubLifecycleStageRebootManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttr(resourceName, "reboot_timeout_in_mins", "10"),
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
	})
}
