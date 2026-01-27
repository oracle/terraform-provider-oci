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
	OsManagementHubManagedInstanceRemoveSnapsManagementRequiredOnlyResource = OsManagementHubManagedInstanceRemoveSnapsManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_remove_snaps_management", "test_managed_instance_remove_snaps_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceRemoveSnapsManagementRepresentation)

	OsManagementHubManagedInstanceRemoveSnapsManagementRepresentation = map[string]interface{}{
		"managed_instance_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
		"snap_details":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubManagedInstanceRemoveSnapsManagementSnapDetailsRepresentation},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceRemoveSnapsManagementWorkRequestDetailsRepresentation},
		"depends_on":           acctest.Representation{RepType: acctest.Required, Create: []string{"oci_os_management_hub_managed_instance_install_snaps_management.test_managed_instance_install_snaps_management"}},
	}
	OsManagementHubManagedInstanceRemoveSnapsManagementSnapDetailsRepresentation = map[string]interface{}{
		"name":     acctest.Representation{RepType: acctest.Required, Create: `hello-world`},
		"revision": acctest.Representation{RepType: acctest.Optional, Create: `27`},
	}
	OsManagementHubManagedInstanceRemoveSnapsManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubManagedInstanceRemoveSnapsManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubUbuntuManagedInstanceRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceRemoveSnapsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceRemoveSnapsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_managed_instance_remove_snaps_management.test_managed_instance_remove_snaps_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubManagedInstanceRemoveSnapsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_remove_snaps_management", "test_managed_instance_remove_snaps_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceRemoveSnapsManagementRepresentation), "osmanagementhub", "managedInstanceRemoveSnapsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceRemoveSnapsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_remove_snaps_management", "test_managed_instance_remove_snaps_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceRemoveSnapsManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_snaps_management", "test_managed_instance_install_snaps_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceInstallSnapsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "snap_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "snap_details.0.name", "hello-world"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceRemoveSnapsManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceRemoveSnapsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_remove_snaps_management", "test_managed_instance_remove_snaps_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceRemoveSnapsManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_snaps_management", "test_managed_instance_install_snaps_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceInstallSnapsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "snap_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "snap_details.0.name", "hello-world"),
				resource.TestCheckResourceAttr(resourceName, "snap_details.0.revision", "27"),
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
