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
	OsManagementHubSoftwarePackageSingularDataSourceRepresentation = map[string]interface{}{
		"software_package_name": acctest.Representation{RepType: acctest.Required, Create: `ModemManager-glib-devel-1.10.4-1.el8.x86_64.rpm`},
	}

	OsManagementHubSoftwarePackageDataSourceRepresentation = map[string]interface{}{
		"architecture":          acctest.Representation{RepType: acctest.Required, Create: `X86_64`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"display_name_contains": acctest.Representation{RepType: acctest.Required, Create: `ModemManager`},
		"is_latest":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"os_family":             acctest.Representation{RepType: acctest.Required, Create: `ORACLE_LINUX_8`},
		"version":               acctest.Representation{RepType: acctest.Required, Create: `1.10.4-1.el8`},
	}

	OsManagementHubSoftwarePackageResourceConfig = OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwarePackageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwarePackageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_software_packages.test_software_packages"
	singularDatasourceName := "data.oci_os_management_hub_software_package.test_software_package"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_packages", "test_software_packages", acctest.Required, acctest.Create, OsManagementHubSoftwarePackageDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubSoftwarePackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "architecture", "X86_64"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "ModemManager"),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(datasourceName, "version", "1.10.4-1.el8"),

				resource.TestCheckResourceAttrSet(datasourceName, "software_package_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_package", "test_software_package", acctest.Required, acctest.Create, OsManagementHubSoftwarePackageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubSoftwarePackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_package_name"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "architecture"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "checksum"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "checksum_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dependencies.#", "14"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "files.#", "182"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_latest"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_modified_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),

				// TODO - API bug returns null osFamilies and/or softwareSources
				//resource.TestCheckResourceAttr(singularDatasourceName, "os_families.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "software_sources.#", "1"),
			),
		},
	})
}
