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
	OsManagementHubSoftwareSourceModuleStreamSingularDataSourceRepresentation = map[string]interface{}{
		"module_name":        acctest.Representation{RepType: acctest.Required, Create: `php`},
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`},
		"stream_name":        acctest.Representation{RepType: acctest.Required, Create: `8.0`},
	}

	OsManagementHubSoftwareSourceModuleStreamDataSourceRepresentation = map[string]interface{}{
		"software_source_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`},
		"is_latest":            acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"module_name":          acctest.Representation{RepType: acctest.Optional, Create: `php`},
		"module_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `php`},
		"name":                 acctest.Representation{RepType: acctest.Optional, Create: `8.0`},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourceModuleStreamResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourceModuleStreamResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_software_source_module_streams.test_software_source_module_streams"
	singularDatasourceName := "data.oci_os_management_hub_software_source_module_stream.test_software_source_module_stream"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_module_streams", "test_software_source_module_streams", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceModuleStreamDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "is_latest", "true"),
				resource.TestCheckResourceAttr(datasourceName, "module_name", "php"),
				resource.TestCheckResourceAttr(datasourceName, "module_name_contains", "php"),
				resource.TestCheckResourceAttr(datasourceName, "name", "8.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source_module_stream", "test_software_source_module_stream", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceModuleStreamSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "module_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_source_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_name"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "arch_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_default"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "module_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "packages.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profiles.#"),
			),
		},
	})
}
