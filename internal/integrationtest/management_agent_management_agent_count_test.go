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
	managementAgentCountSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"group_by":       acctest.Representation{RepType: acctest.Required, Create: []string{`version`}},
		"has_plugins":    acctest.Representation{RepType: acctest.Required, Create: `true`},
		"install_type":   acctest.Representation{RepType: acctest.Required, Create: `AGENT`},
	}

	ManagementAgentCountResourceConfig = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentCountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentCountResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_management_agent_management_agent_count.test_management_agent_count"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_count", "test_management_agent_count", acctest.Required, acctest.Create, managementAgentCountSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagementAgentCountResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "has_plugins", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "install_type", "AGENT"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
