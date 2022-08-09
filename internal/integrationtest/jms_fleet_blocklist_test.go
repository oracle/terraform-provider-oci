// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	JmsJmsFleetBlocklistDataSourceRepresentation = map[string]interface{}{
		"fleet_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
		"operation": acctest.Representation{RepType: acctest.Optional, Create: `DELETE_JAVA_INSTALLATION`},
	}

	fleetForBlocklistRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `Created Fleet for Blocklist`},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `Created Fleet for Blocklist`},
		"is_advanced_features_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"inventory_log":                acctest.RepresentationGroup{RepType: acctest.Required, Group: JmsFleetInventoryLogRepresentation},
		"operation_log":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetOperationLogRepresentation},
	}

	JmsJmsFleetBlocklistResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Required, acctest.Create, fleetForBlocklistRepresentation)
)

// issue-routing-tag: jms/default
func TestJmsFleetBlocklistResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetBlocklistResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	inventoryLogGroupId := utils.GetEnvSettingWithBlankDefault("inventory_log_group_ocid_for_create")
	inventoryLogGroupIdVariableStr := fmt.Sprintf("variable \"inventory_log_group_id_for_create\" { default = \"%s\" }\n", inventoryLogGroupId)

	operationLogGroupId := utils.GetEnvSettingWithBlankDefault("operation_log_group_ocid_for_create")
	operationLogGroupIdVariableStr := fmt.Sprintf("variable \"operation_log_group_id_for_create\" { default = \"%s\" }\n", operationLogGroupId)

	inventoryLogId := utils.GetEnvSettingWithBlankDefault("inventory_log_ocid_for_create")
	inventoryLogIdVariableStr := fmt.Sprintf("variable \"inventory_log_id_for_create\" { default = \"%s\" }\n", inventoryLogId)

	operationLogId := utils.GetEnvSettingWithBlankDefault("operation_log_ocid_for_create")
	operationLogIdVariableStr := fmt.Sprintf("variable \"operation_log_id_for_create\" { default = \"%s\" }\n", operationLogId)

	datasourceName := "data.oci_jms_fleet_blocklists.test_fleet_blocklists"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+
		compartmentIdVariableStr+
		inventoryLogGroupIdVariableStr+
		inventoryLogIdVariableStr+
		operationLogGroupIdVariableStr+
		operationLogIdVariableStr, "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_fleet_blocklists", "test_fleet_blocklists", acctest.Optional, acctest.Update, JmsJmsFleetBlocklistDataSourceRepresentation) +
				compartmentIdVariableStr +
				inventoryLogGroupIdVariableStr +
				inventoryLogIdVariableStr +
				operationLogGroupIdVariableStr +
				operationLogIdVariableStr +
				JmsJmsFleetBlocklistResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "operation", "DELETE_JAVA_INSTALLATION"),

				resource.TestCheckResourceAttr(datasourceName, "items.#", "0"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("JmsFleetBlocklist") {
		resource.AddTestSweepers("JmsFleetBlocklist", &resource.Sweeper{
			Name:         "JmsFleetBlocklist",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepJmsFleetResource,
		})
	}
}
