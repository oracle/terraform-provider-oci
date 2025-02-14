// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubSoftwareSourceAvailableSoftwarePackageDataSourceRepresentation1 = map[string]interface{}{
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_software_source.test_software_source.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `ModemManager`},
		"is_latest":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	OsManagementHubSoftwareSourceAvailableSoftwarePackageDataSourceRepresentation2 = map[string]interface{}{
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_software_source.test_software_source.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `ModemManager`},
	}

	OsManagementHubSoftwareSourceAvailableSoftwarePackageDataSourceRepresentation3 = map[string]interface{}{
		"software_source_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_software_source.test_software_source.id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `ModemManager`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `ModemManager`},
		"is_latest":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	OsManagementHubSoftwareSourceAvailableSoftwarePackageDataSourceRepresentation4 = map[string]interface{}{
		"software_source_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_software_source.test_software_source.id}`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `NetworkManager-ad`},
		"is_latest":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	OsManagementHubSoftwareSourceAvailableSoftwarePackageResourceConfig = OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourceAvailableSoftwarePackageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourceAvailableSoftwarePackageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_software_source_available_software_packages.test_software_source_available_software_packages"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_available_software_packages", "test_software_source_available_software_packages", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceAvailableSoftwarePackageDataSourceRepresentation1) +
				compartmentIdVariableStr + OsManagementHubSoftwareSourceAvailableSoftwarePackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "ModemManager"),
				resource.TestCheckResourceAttr(datasourceName, "is_latest", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "software_package_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "software_package_collection.0.items.#", "1"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_available_software_packages", "test_software_source_available_software_packages", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceAvailableSoftwarePackageDataSourceRepresentation2) +
				compartmentIdVariableStr + OsManagementHubSoftwareSourceAvailableSoftwarePackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "ModemManager"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "software_package_collection.#"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_available_software_packages", "test_software_source_available_software_packages", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceAvailableSoftwarePackageDataSourceRepresentation3) +
				compartmentIdVariableStr + OsManagementHubSoftwareSourceAvailableSoftwarePackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "ModemManager"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "ModemManager"),
				resource.TestCheckResourceAttr(datasourceName, "is_latest", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "software_package_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "software_package_collection.0.items.#", "1"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_available_software_packages", "test_software_source_available_software_packages", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceAvailableSoftwarePackageDataSourceRepresentation4) +
				compartmentIdVariableStr + OsManagementHubSoftwareSourceAvailableSoftwarePackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "NetworkManager-ad"),
				resource.TestCheckResourceAttr(datasourceName, "is_latest", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "software_package_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "software_package_collection.0.items.#", "1"),
			),
		},
	})
}
