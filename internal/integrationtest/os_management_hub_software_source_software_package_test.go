// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	OsManagementHubSoftwareSourceSoftwarePackageSingularDataSourceRepresentation = map[string]interface{}{
		"software_package_name": acctest.Representation{RepType: acctest.Required, Create: `zsh-5.5.1-10.el8.x86_64.rpm`},
		"software_source_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`},
	}

	OsManagementHubSoftwareSourceSoftwarePackageDataSourceRepresentation = map[string]interface{}{
		"software_source_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `zsh`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `zsh`},
		"is_latest":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourceSoftwarePackageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourceSoftwarePackageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_software_source_software_packages.test_software_source_software_packages"
	singularDatasourceName := "data.oci_os_management_hub_software_source_software_package.test_software_source_software_package"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_software_packages", "test_software_source_software_packages", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceSoftwarePackageDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "zsh"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "zsh"),
				resource.TestCheckResourceAttr(datasourceName, "is_latest", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "software_package_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_software_package", "test_software_source_software_package", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceSoftwarePackageSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "software_package_name", "zsh-5.5.1-10.el8.x86_64.rpm"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_source_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "architecture"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "checksum"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "checksum_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dependencies.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "files.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_latest"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_modified_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_sources.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
	})
}
