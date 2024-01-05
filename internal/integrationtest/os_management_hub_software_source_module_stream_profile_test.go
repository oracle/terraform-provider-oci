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
	OsManagementHubSoftwareSourceModuleStreamProfileSingularDataSourceRepresentation = map[string]interface{}{
		"module_name":        acctest.Representation{RepType: acctest.Required, Create: `php`},
		"profile_name":       acctest.Representation{RepType: acctest.Required, Create: `common`},
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`},
		"stream_name":        acctest.Representation{RepType: acctest.Required, Create: `8.0`},
	}

	OsManagementHubSoftwareSourceModuleStreamProfileDataSourceRepresentation = map[string]interface{}{
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`},
		"module_name":        acctest.Representation{RepType: acctest.Optional, Create: `php`},
		"name":               acctest.Representation{RepType: acctest.Optional, Create: `common`},
		"stream_name":        acctest.Representation{RepType: acctest.Optional, Create: `8.0`},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourceModuleStreamProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourceModuleStreamProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_software_source_module_stream_profiles.test_software_source_module_stream_profiles"
	singularDatasourceName := "data.oci_os_management_hub_software_source_module_stream_profile.test_software_source_module_stream_profile"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_module_stream_profiles", "test_software_source_module_stream_profiles", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceModuleStreamProfileDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "module_name", "php"),
				resource.TestCheckResourceAttr(datasourceName, "name", "common"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "stream_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_profile_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_module_stream_profile", "test_software_source_module_stream_profile", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceModuleStreamProfileSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "module_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_source_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_name"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_default"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "module_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "packages.#"),
			),
		},
	})
}
