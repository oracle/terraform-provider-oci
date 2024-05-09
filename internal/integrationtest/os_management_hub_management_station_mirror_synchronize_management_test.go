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
	OsManagementHubManagementStationMirrorSynchronizeManagementRepresentation = map[string]interface{}{
		"management_station_id": acctest.Representation{RepType: acctest.Required, Create: `${var.management_station_id}`},
		"mirror_id":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`},
	}

	OsManagementHubManagementStationMirrorSynchronizeManagementResourceDependencies = OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagementStationMirrorSynchronizeManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagementStationMirrorSynchronizeManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementStationId := utils.GetEnvSettingWithBlankDefault("management_station_ocid")
	managementStationIdVariableStr := fmt.Sprintf("variable \"management_station_id\" { default = \"%s\" }\n", managementStationId)

	resourceName := "oci_os_management_hub_management_station_mirror_synchronize_management.test_management_station_mirror_synchronize_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managementStationIdVariableStr+OsManagementHubManagementStationMirrorSynchronizeManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station_mirror_synchronize_management", "test_management_station_mirror_synchronize_management", acctest.Required, acctest.Create, OsManagementHubManagementStationMirrorSynchronizeManagementRepresentation), "osmanagementhub", "managementStationMirrorSynchronizeManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementStationIdVariableStr + OsManagementHubManagementStationMirrorSynchronizeManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station_mirror_synchronize_management", "test_management_station_mirror_synchronize_management", acctest.Required, acctest.Create, OsManagementHubManagementStationMirrorSynchronizeManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "management_station_id"),
				resource.TestCheckResourceAttrSet(resourceName, "mirror_id"),

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
