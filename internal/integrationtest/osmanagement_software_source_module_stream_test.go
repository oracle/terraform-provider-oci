// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"

	"terraform-provider-oci/internal/utils"
)

var (
	softwareSourceModuleStreamSingularDataSourceRepresentation = map[string]interface{}{
		"module_name":        acctest.Representation{RepType: acctest.Required, Create: `subversion`},
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: softwareSourceOCID},
		"stream_name":        acctest.Representation{RepType: acctest.Required, Create: `1.10`},
	}
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementSoftwareSourceModuleStreamResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementSoftwareSourceModuleStreamResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_osmanagement_software_source_module_stream.test_software_source_module_stream"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_software_source_module_stream", "test_software_source_module_stream", acctest.Required, acctest.Create, softwareSourceModuleStreamSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "module_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_source_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_name"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "architecture"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_default"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "module_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "packages.#", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "profiles.#", "2"),
			),
		},
	})
}
