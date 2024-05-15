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
	OsManagementHubManagedInstanceInstallWindowsUpdatesManagementRepresentation = map[string]interface{}{
		"managed_instance_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
		"windows_update_name":  acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_os_management_hub_windows_updates.test_windows_updates.name}`},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceInstallWindowsUpdatesManagementWorkRequestDetailsRepresentation},
	}

	OsManagementHubManagedInstanceInstallWindowsUpdatesUpdateTypeManagementRepresentation = map[string]interface{}{
		"managed_instance_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
		"windows_update_types": acctest.Representation{RepType: acctest.Optional, Create: []string{`OTHER`}},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceInstallWindowsUpdatesManagementWorkRequestDetailsRepresentation},
	}

	OsManagementHubManagedInstanceInstallWindowsUpdatesManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubManagedInstanceInstallWindowsUpdatesManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceWindowsRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_windows_updates", "test_windows_updates", acctest.Optional, acctest.Create, OsManagementHubWindowsUpdateToInstallDataSourceRepresentation)

	OsManagementHubManagedInstanceInstallWindowsUpdatesUpdateTypeManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceWindowsRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceInstallWindowsUpdatesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceInstallWindowsUpdatesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//resourceName := "oci_os_management_hub_managed_instance_install_windows_updates_management.test_managed_instance_install_windows_updates_management"
	//resourceNameUpdateType := "oci_os_management_hub_managed_instance_install_windows_updates_management.test_managed_instance_install_windows_updates_management_update_type"

	//var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubManagedInstanceInstallWindowsUpdatesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_windows_updates_management", "test_managed_instance_install_windows_updates_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceInstallWindowsUpdatesManagementRepresentation), "osmanagementhub", "managedInstanceInstallWindowsUpdatesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with update name optionals
		// This test case was verified to work. As we need a new Windows Update each time we run this test case, disable it for the workflow.
		//{
		//	Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceInstallWindowsUpdatesManagementResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_windows_updates_management", "test_managed_instance_install_windows_updates_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceInstallWindowsUpdatesManagementRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
		//		resource.TestCheckResourceAttr(resourceName, "windows_update_name.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "work_request_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "work_request_details.0.description", "description"),
		//		resource.TestCheckResourceAttr(resourceName, "work_request_details.0.display_name", "displayName"),
		//
		//		func(s *terraform.State) (err error) {
		//			resId, err = acctest.FromInstanceState(s, resourceName, "id")
		//			if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
		//				if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
		//					return errExport
		//				}
		//			}
		//			return err
		//		},
		//	),
		//},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceInstallWindowsUpdatesManagementResourceDependencies,
		},
		// verify Create with update type optionals
		// This test case was verified to work. As we need a new Windows Update each time we run this test case, disable it for the workflow.
		//{
		//	Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceInstallWindowsUpdatesUpdateTypeManagementResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_windows_updates_management", "test_managed_instance_install_windows_updates_management_update_type", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceInstallWindowsUpdatesUpdateTypeManagementRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttrSet(resourceNameUpdateType, "managed_instance_id"),
		//		resource.TestCheckResourceAttr(resourceNameUpdateType, "windows_update_types.#", "1"),
		//		resource.TestCheckResourceAttr(resourceNameUpdateType, "work_request_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceNameUpdateType, "work_request_details.0.description", "description"),
		//		resource.TestCheckResourceAttr(resourceNameUpdateType, "work_request_details.0.display_name", "displayName"),
		//
		//		func(s *terraform.State) (err error) {
		//			resId, err = acctest.FromInstanceState(s, resourceNameUpdateType, "id")
		//			if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
		//				if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceNameUpdateType); errExport != nil {
		//					return errExport
		//				}
		//			}
		//			return err
		//		},
		//	),
		//},
	})
}
