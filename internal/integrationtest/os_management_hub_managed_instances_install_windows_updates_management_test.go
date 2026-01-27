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
	OsManagementHubManagedInstancesInstallWindowsUpdatesManagementRequiredOnlyResource = OsManagementHubManagedInstancesInstallWindowsUpdatesManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instances_install_windows_updates_management", "test_managed_instances_install_windows_updates_management", acctest.Required, acctest.Create, OsManagementHubManagedInstancesInstallWindowsUpdatesManagementRepresentation)

	OsManagementHubManagedInstancesInstallWindowsUpdatesManagementRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"windows_update_types": acctest.Representation{RepType: acctest.Required, Create: []string{`SECURITY`}},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstancesInstallWindowsUpdatesManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubManagedInstancesInstallWindowsUpdatesManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubManagedInstancesInstallWindowsUpdatesManagementResourceDependencies = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstancesInstallWindowsUpdatesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstancesInstallWindowsUpdatesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_managed_instances_install_windows_updates_management.test_managed_instances_install_windows_updates_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubManagedInstancesInstallWindowsUpdatesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instances_install_windows_updates_management", "test_managed_instances_install_windows_updates_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstancesInstallWindowsUpdatesManagementRepresentation), "osmanagementhub", "managedInstancesInstallWindowsUpdatesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstancesInstallWindowsUpdatesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instances_install_windows_updates_management", "test_managed_instances_install_windows_updates_management", acctest.Required, acctest.Create, OsManagementHubManagedInstancesInstallWindowsUpdatesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstancesInstallWindowsUpdatesManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstancesInstallWindowsUpdatesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instances_install_windows_updates_management", "test_managed_instances_install_windows_updates_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstancesInstallWindowsUpdatesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "windows_update_types.#", "1"),
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
