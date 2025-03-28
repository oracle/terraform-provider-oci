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
	OsManagementHubLifecycleStageAttachManagedInstancesManagementRepresentation = map[string]interface{}{
		"lifecycle_stage_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id}`},
		"managed_instance_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubLifecycleStageAttachManagedInstancesManagementManagedInstanceDetailsRepresentation},
	}
	OsManagementHubLCStageDetachManagedInstancesManagementRepresentation = map[string]interface{}{
		"lifecycle_stage_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id}`},
		"managed_instance_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubLifecycleStageDetachManagedInstancesManagementManagedInstanceDetailsRepresentation},
		"depends_on":               acctest.Representation{RepType: acctest.Required, Create: []string{`oci_os_management_hub_lifecycle_stage_attach_managed_instances_management.test_lifecycle_stage_attach_managed_instances_management`}},
	}
	OsManagementHubLifecycleStageAttachManagedInstancesManagementManagedInstanceDetailsRepresentation = map[string]interface{}{
		"managed_instances":    acctest.Representation{RepType: acctest.Required, Create: []string{utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_ocid")}},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubLifecycleStageAttachManagedInstancesManagementManagedInstanceDetailsWorkRequestDetailsRepresentation},
	}
	OsManagementHubLifecycleStageAttachManagedInstancesManagementManagedInstanceDetailsWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubLifecycleStageAttachManagedInstancesManagementResourceDependencies = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubLifecycleStageAttachManagedInstancesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubLifecycleStageAttachManagedInstancesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_lifecycle_stage_attach_managed_instances_management.test_lifecycle_stage_attach_managed_instances_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubLifecycleStageAttachManagedInstancesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_attach_managed_instances_management", "test_lifecycle_stage_attach_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubLifecycleStageAttachManagedInstancesManagementRepresentation), "osmanagementhub", "lifecycleStageAttachManagedInstancesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleStageAttachManagedInstancesManagementResourceDependencies + OsManagementHubLifecycleEnvironmentResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_detach_managed_instances_management", "test_lifecycle_stage_detach_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubLCStageDetachManagedInstancesManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage_attach_managed_instances_management", "test_lifecycle_stage_attach_managed_instances_management", acctest.Optional, acctest.Create, OsManagementHubLifecycleStageAttachManagedInstancesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_details.0.managed_instances.#", "1"),

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
