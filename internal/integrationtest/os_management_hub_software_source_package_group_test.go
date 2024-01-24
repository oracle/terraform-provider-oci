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
	OsManagementHubSoftwareSourcePackageGroupSingularDataSourceRepresentation = map[string]interface{}{
		"package_group_id":   acctest.Representation{RepType: acctest.Required, Create: `base`},
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`},
	}

	OsManagementHubSoftwareSourcePackageGroupDataSourceRepresentation = map[string]interface{}{
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`},
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"group_type":         acctest.Representation{RepType: acctest.Optional, Create: []string{`GROUP`}},
		"name":               acctest.Representation{RepType: acctest.Optional, Create: `Base`},
		"name_contains":      acctest.Representation{RepType: acctest.Optional, Create: `Base`},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourcePackageGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourcePackageGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_software_source_package_groups.test_software_source_package_groups"
	singularDatasourceName := "data.oci_os_management_hub_software_source_package_group.test_software_source_package_group"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_package_groups", "test_software_source_package_groups", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourcePackageGroupDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "group_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "name", "Base"),
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "Base"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "package_group_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_package_group", "test_software_source_package_group", acctest.Required, acctest.Create, OsManagementHubSoftwareSourcePackageGroupSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_group_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_source_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_default"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_user_visible"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "packages.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repositories.#"),
			),
		},
	})
}
