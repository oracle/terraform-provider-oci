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
	// Oracle Linux 8 Application Stream (x86_64)
	softwareSourceOCID                                                          = utils.GetEnvSettingWithBlankDefault("software_source_ocid")
	OsmanagementOsmanagementSoftwareSourceStreamProfileDataSourceRepresentation = map[string]interface{}{
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: softwareSourceOCID},
		"module_name":        acctest.Representation{RepType: acctest.Optional, Create: `subversion`},
		"profile_name":       acctest.Representation{RepType: acctest.Optional, Create: `common`},
		"stream_name":        acctest.Representation{RepType: acctest.Optional, Create: `1.10`},
	}
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementSoftwareSourceStreamProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementSoftwareSourceStreamProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_osmanagement_software_source_stream_profiles.test_software_source_stream_profiles"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_software_source_stream_profiles", "test_software_source_stream_profiles", acctest.Optional, acctest.Create, OsmanagementOsmanagementSoftwareSourceStreamProfileDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "module_name", "subversion"),
				resource.TestCheckResourceAttrSet(datasourceName, "profile_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "stream_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_profiles.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_profiles.0.module_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_profiles.0.profile_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "module_stream_profiles.0.stream_name"),
			),
		},
	})
}
