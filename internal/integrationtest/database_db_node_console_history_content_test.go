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
	DatabaseDbNodeConsoleHistoryContentSingularDataSourceRepresentation = map[string]interface{}{
		"db_node_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.db_node_id}`},
		"console_history_id": acctest.Representation{RepType: acctest.Required, Create: `${var.console_history_id}`},
	}

	DatabaseConsoleHistoryConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_history", "test_console_history", acctest.Optional, acctest.Update, DbNodeConsoleHistoryRepresentation) +
		DatabaseDbNodeConsoleHistoryContentResourceConfig

	DbNodeConsoleHistoryRepresentation = map[string]interface{}{
		"db_node_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_node.test_db_node.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `console-history-20221202-1943`, Update: `displayName2`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseDbNodeConsoleHistoryContentResourceConfig = DatabaseVmClusterRequiredOnlyResource +
		AvailabilityDomainConfig + `
		  data "oci_database_db_nodes" "test_db_nodes" {
		     compartment_id = "${var.compartment_id}"
		     vm_cluster_id = "${oci_database_vm_cluster.test_vm_cluster.id}"
		  }
		  data "oci_database_db_node" "test_db_node" {
		     db_node_id = "${data.oci_database_db_nodes.test_db_nodes.db_nodes.0.id}"
		  }`
)

// issue-routing-tag: database/default
func TestDatabaseDbNodeConsoleHistoryContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbNodeConsoleHistoryContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	dbNodeId := utils.GetEnvSettingWithBlankDefault("db_node_ocid")
	consoleHistoryId := utils.GetEnvSettingWithBlankDefault("console_ocid")
	dbNodeIdVariableStr := fmt.Sprintf("variable \"db_node_id\" { default = \"%s\" }\n", dbNodeId)
	consoleHistoryIdStr := fmt.Sprintf("variable \"console_history_id\" { default = \"%s\" }\n", consoleHistoryId)

	singularDatasourceName := "data.oci_database_db_node_console_history_content.test_db_node_console_history_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + consoleHistoryIdStr + dbNodeIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node_console_history_content", "test_db_node_console_history_content", acctest.Required, acctest.Create, DatabaseDbNodeConsoleHistoryContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDbNodeConsoleHistoryContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "console_history_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_id"),
			),
		},
	})
}
