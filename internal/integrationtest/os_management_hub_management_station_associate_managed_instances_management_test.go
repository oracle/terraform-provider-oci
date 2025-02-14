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
	OsManagementHubManagementStationAssociateManagedInstancesManagementRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station_associate_managed_instances_management", "test_management_station_associate_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubManagementStationAssociateManagedInstancesManagementRepresentation)

	OsManagementHubManagementStationAssociateManagedInstancesManagementRepresentation = map[string]interface{}{
		"managed_instances":     acctest.Representation{RepType: acctest.Required, Create: []string{utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_onprem_ocid")}},
		"management_station_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("osmh_management_station_ocid")},
		"work_request_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagementStationAssociateManagedInstancesManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubManagementStationAssociateManagedInstancesManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagementStationAssociateManagedInstancesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagementStationAssociateManagedInstancesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_management_station_associate_managed_instances_management.test_management_station_associate_managed_instances_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station_associate_managed_instances_management", "test_management_station_associate_managed_instances_management", acctest.Optional, acctest.Create, OsManagementHubManagementStationAssociateManagedInstancesManagementRepresentation), "osmanagementhub", "managementStationAssociateManagedInstancesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station_associate_managed_instances_management", "test_management_station_associate_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubManagementStationAssociateManagedInstancesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "managed_instances.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "management_station_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station_associate_managed_instances_management", "test_management_station_associate_managed_instances_management", acctest.Optional, acctest.Create, OsManagementHubManagementStationAssociateManagedInstancesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "managed_instances.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "management_station_id"),
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
