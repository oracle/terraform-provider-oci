// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	managementAgentGetAutoUpgradableConfigSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	ManagementAgentGetAutoUpgradableConfigResourceConfig = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentGetAutoUpgradableConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentGetAutoUpgradableConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_management_agent_management_agent_get_auto_upgradable_config.test_management_agent_get_auto_upgradable_config"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_get_auto_upgradable_config", "test_management_agent_get_auto_upgradable_config", Required, Create, managementAgentGetAutoUpgradableConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagementAgentGetAutoUpgradableConfigResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_agent_auto_upgradable"),
			),
		},
	})
}
