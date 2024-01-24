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
	DatabaseExternalcontainerdatabasesStackMonitoringRepresentation = map[string]interface{}{
		"external_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"external_database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"enable_stack_monitoring":        acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	DatabaseExternalcontainerdatabasesStackMonitoringResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseExternalcontainerdatabasesStackMonitoringResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalcontainerdatabasesStackMonitoringResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	agentId := utils.GetEnvSettingWithBlankDefault("connector_agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)

	resourceName := "oci_database_externalcontainerdatabases_stack_monitoring.test_externalcontainerdatabases_stack_monitoring"
	resourceCDB := "oci_database_external_container_database.test_external_container_database"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+DatabaseExternalcontainerdatabasesStackMonitoringResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_externalcontainerdatabases_stack_monitoring", "test_externalcontainerdatabases_stack_monitoring", acctest.Required, acctest.Create, DatabaseExternalcontainerdatabasesStackMonitoringRepresentation), "database", "externalContainerDatabasesStackMonitoring", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalcontainerdatabasesStackMonitoringResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_externalcontainerdatabases_stack_monitoring", "test_externalcontainerdatabases_stack_monitoring", acctest.Required, acctest.Create, DatabaseExternalcontainerdatabasesStackMonitoringRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},

		// verify Enabled
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalcontainerdatabasesStackMonitoringResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_externalcontainerdatabases_stack_monitoring", "test_externalcontainerdatabases_stack_monitoring", acctest.Required, acctest.Create, DatabaseExternalcontainerdatabasesStackMonitoringRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceCDB, "stack_monitoring_config.0.stack_monitoring_status", "ENABLED"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalcontainerdatabasesStackMonitoringResourceDependencies,
		},
	})
}
