// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ComputeinstanceagentComputeinstanceagentInstanceAvailablePluginDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"os_name":        acctest.Representation{RepType: acctest.Required, Create: `Oracle Linux`},
		"os_version":     acctest.Representation{RepType: acctest.Required, Create: `7.8`},
	}

	ComputeinstanceagentInstanceAvailablePluginResourceConfig = ""
)

// issue-routing-tag: computeinstanceagent/default
func TestComputeinstanceagentInstanceAvailablePluginResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestComputeinstanceagentInstanceAvailablePluginResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_computeinstanceagent_instance_available_plugins.test_instance_available_plugins"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_computeinstanceagent_instance_available_plugins", "test_instance_available_plugins", acctest.Required, acctest.Create, ComputeinstanceagentComputeinstanceagentInstanceAvailablePluginDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeinstanceagentInstanceAvailablePluginResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "os_name", "Oracle Linux"),
				resource.TestCheckResourceAttr(datasourceName, "os_version", "7.8"),

				resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.0.is_enabled_by_default"),
				resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.0.is_supported"),
				resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.0.summary"),
			),
		},
	})
}
