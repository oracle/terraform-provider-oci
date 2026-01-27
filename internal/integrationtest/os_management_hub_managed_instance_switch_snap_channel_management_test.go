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
	OsManagementHubManagedInstanceSwitchSnapChannelManagementRequiredOnlyResource = OsManagementHubManagedInstanceSwitchSnapChannelManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_switch_snap_channel_management", "test_managed_instance_switch_snap_channel_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceSwitchSnapChannelManagementRepresentation)

	OsManagementHubManagedInstanceInstallSnapsTestSpeedCliManagementRepresentation = map[string]interface{}{
		"managed_instance_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
		"snap_details":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubManagedInstanceInstallSnapsSpeedTestCliManagementSnapDetailsRepresentation},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceInstallSnapsManagementWorkRequestDetailsRepresentation},
	}

	OsManagementHubManagedInstanceInstallSnapsSpeedTestCliManagementSnapDetailsRepresentation = map[string]interface{}{
		"name":    acctest.Representation{RepType: acctest.Required, Create: `speedtest-cli`},
		"channel": acctest.Representation{RepType: acctest.Required, Create: `stable`},
	}

	OsManagementHubManagedInstanceSwitchSnapChannelManagementRepresentation = map[string]interface{}{
		"managed_instance_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
		"snap_details":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubManagedInstanceSwitchSnapChannelManagementSnapDetailsRepresentation},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceSwitchSnapChannelManagementWorkRequestDetailsRepresentation},
		"depends_on":           acctest.Representation{RepType: acctest.Required, Create: []string{"oci_os_management_hub_managed_instance_install_snaps_management.test_managed_instance_install_snaps_management"}},
	}
	OsManagementHubManagedInstanceSwitchSnapChannelManagementSnapDetailsRepresentation = map[string]interface{}{
		"name":    acctest.Representation{RepType: acctest.Required, Create: `speedtest-cli`},
		"channel": acctest.Representation{RepType: acctest.Required, Create: `beta`},
	}
	OsManagementHubManagedInstanceSwitchSnapChannelManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubManagedInstanceSwitchSnapChannelManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubUbuntuManagedInstanceRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceSwitchSnapChannelManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceSwitchSnapChannelManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_managed_instance_switch_snap_channel_management.test_managed_instance_switch_snap_channel_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubManagedInstanceSwitchSnapChannelManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_switch_snap_channel_management", "test_managed_instance_switch_snap_channel_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceSwitchSnapChannelManagementRepresentation), "osmanagementhub", "managedInstanceSwitchSnapChannelManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceSwitchSnapChannelManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_switch_snap_channel_management", "test_managed_instance_switch_snap_channel_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceSwitchSnapChannelManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_snaps_management", "test_managed_instance_install_snaps_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceInstallSnapsTestSpeedCliManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceSwitchSnapChannelManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceSwitchSnapChannelManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_switch_snap_channel_management", "test_managed_instance_switch_snap_channel_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceSwitchSnapChannelManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_snaps_management", "test_managed_instance_install_snaps_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceInstallSnapsTestSpeedCliManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "snap_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "snap_details.0.channel", "beta"),
				resource.TestCheckResourceAttr(resourceName, "snap_details.0.name", "speedtest-cli"),
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
