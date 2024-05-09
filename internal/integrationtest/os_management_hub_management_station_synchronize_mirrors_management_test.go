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
	OsManagementHubManagementStationSynchronizeMirrorsManagementRepresentation = map[string]interface{}{
		"management_station_id": acctest.Representation{RepType: acctest.Required, Create: `${var.management_station_id}`},
		"software_source_list":  acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`}},
	}

	OsManagementHubManagementStationSynchronizeMirrorsManagementResourceDependencies = OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagementStationSynchronizeMirrorsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagementStationSynchronizeMirrorsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementStationId := utils.GetEnvSettingWithBlankDefault("management_station_ocid")
	managementStationIdVariableStr := fmt.Sprintf("variable \"management_station_id\" { default = \"%s\" }\n", managementStationId)

	resourceName := "oci_os_management_hub_management_station_synchronize_mirrors_management.test_management_station_synchronize_mirrors_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managementStationIdVariableStr+OsManagementHubManagementStationSynchronizeMirrorsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station_synchronize_mirrors_management", "test_management_station_synchronize_mirrors_management", acctest.Required, acctest.Create, OsManagementHubManagementStationSynchronizeMirrorsManagementRepresentation), "osmanagementhub", "managementStationSynchronizeMirrorsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementStationIdVariableStr + OsManagementHubManagementStationSynchronizeMirrorsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station_synchronize_mirrors_management", "test_management_station_synchronize_mirrors_management", acctest.Required, acctest.Create, OsManagementHubManagementStationSynchronizeMirrorsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "management_station_id"),
				resource.TestCheckResourceAttr(resourceName, "software_source_list.#", "1"),

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
