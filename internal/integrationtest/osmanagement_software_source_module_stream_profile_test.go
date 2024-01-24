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
	OsmanagementOsmanagementSoftwareSourceModuleStreamProfileSingularDataSourceRepresentation = map[string]interface{}{
		"module_name":        acctest.Representation{RepType: acctest.Required, Create: `subversion`},
		"profile_name":       acctest.Representation{RepType: acctest.Required, Create: `common`},
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: softwareSourceOCID},
		"stream_name":        acctest.Representation{RepType: acctest.Required, Create: `1.10`},
	}
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementSoftwareSourceModuleStreamProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementSoftwareSourceModuleStreamProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_osmanagement_software_source_module_stream_profile.test_software_source_module_stream_profile"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_software_source_module_stream_profile", "test_software_source_module_stream_profile", acctest.Required, acctest.Create, OsmanagementOsmanagementSoftwareSourceModuleStreamProfileSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "module_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_source_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_name"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_default"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "module_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "packages.#"),
			),
		},
	})
}
