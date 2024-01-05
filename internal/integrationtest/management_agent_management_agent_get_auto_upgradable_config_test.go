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
	managementAgentGetAutoUpgradableConfigSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	ManagementAgentGetAutoUpgradableConfigResourceConfig = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentGetAutoUpgradableConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentGetAutoUpgradableConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_management_agent_management_agent_get_auto_upgradable_config.test_management_agent_get_auto_upgradable_config"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_get_auto_upgradable_config", "test_management_agent_get_auto_upgradable_config", acctest.Required, acctest.Create, managementAgentGetAutoUpgradableConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagementAgentGetAutoUpgradableConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_agent_auto_upgradable"),
			),
		},
	})
}
