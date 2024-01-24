// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	tf_database "github.com/oracle/terraform-provider-oci/internal/service/database"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"strconv"
	"testing"
)

var (
	DbNodeConsoleConnectionRequiredOnlyResourceExaCC = DatabaseDbNodeConsoleConnectionResourcesExaCCDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Required, acctest.Create, DatabaseDbNodeConsoleConnectionRepresentationExaCC)

	DatabaseDbNodeConsoleConnectionResourceConfigExaCC = DatabaseDbNodeConsoleConnectionResourcesExaCCDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Optional, acctest.Update, DatabaseDbNodeConsoleConnectionRepresentationExaCC)

	DatabaseDatabaseDbNodeConsoleConnectionSingularDataSourceRepresentationExaCC = map[string]interface{}{
		"db_node_id": acctest.Representation{RepType: acctest.Required, Create: `${var.db_node_id}`},
		"id":         acctest.Representation{RepType: acctest.Required, Create: `${var.console_connection_id}`},
	}

	DatabaseDatabaseVmClusterSingularDataSourceRepresentationExacc = map[string]interface{}{
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	DatabaseDatabaseDbNodeConsoleConnectionDataSourceRepresentationExaCC = map[string]interface{}{
		"db_node_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_node.test_db_node.id}`},
		"filter":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDbNodeConsoleConnectionDataSourceFilterRepresentationExaCC}}
	DatabaseDbNodeConsoleConnectionDataSourceFilterRepresentationExaCC = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `db_node_id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_database_db_node.test_db_node.id}`}},
	}

	DatabaseDbNodeConsoleConnectionRepresentationExaCC = map[string]interface{}{
		"db_node_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_node.test_db_node.id}`},
		"public_key":    acctest.Representation{RepType: acctest.Required, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCldnh+UpWXKNlkRTamqCnmihMMLZ5UofvO5Thxr+dpp427HcRiZaJnrnpbrh+T5c+I+bTE9JQh4Ydsk9hEg1LWXIqCYPTDifaWneJ6o+xpbjW8benndWDR6y3XJ2yLdXOJ9S8HWMXQLkAGBEkxsqNAtdlt2U3RjRbO+4g+PdNBLycZ8vXwaD+GOxq08GtxQTQuzQMWUjt56TVH7OG6v/1nEHTQ5meYHtTP4cx1YqqOQBWyZG4Ikq9Ej8YegmP3Tbf9SRRcrXG0qASNIevofGsmhvhvgXhNtENJpEvf10aJyXmRDUAfUwGSFCKyHVEiCX9MzPLzUlvSbw2Ls1NVkZ3v mytest.vpn.oracle.com`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseDbNodeConsoleConnectionResourcesExaCCDependencies = AvailabilityDomainConfig + DatabaseVmClusterRequiredOnlyResource + `

     data "oci_database_db_nodes" "test_db_nodes" {
        compartment_id = "${var.compartment_id}"
        vm_cluster_id = "${oci_database_vm_cluster.test_vm_cluster.id}"
     }
     data "oci_database_db_node" "test_db_node" {
        db_node_id = "${data.oci_database_db_nodes.test_db_nodes.db_nodes.0.id}"
     }`
)

// issue-routing-tag: database/default
func TestDatabaseDbNodeConsoleConnectionResource_basic_exacc_only(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbNodeConsoleConnectionResource_basic_exacc_only")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	dbNodeId := utils.GetEnvSettingWithBlankDefault("db_node_ocid")
	consoleConnectionId := utils.GetEnvSettingWithBlankDefault("console_ocid")
	consoleId := utils.GetEnvSettingWithBlankDefault("console_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	consoleIdVariableStr := fmt.Sprintf("variable \"console_id\" { default = \"%s\" }\n", consoleId)
	dbNodeIdVariableStr := fmt.Sprintf("variable \"db_node_id\" { default = \"%s\" }\n", dbNodeId)
	consoleConnectionIdStr := fmt.Sprintf("variable \"console_connection_id\" { default = \"%s\" }\n", consoleConnectionId)

	resourceName := "oci_database_db_node_console_connection.test_db_node_console_connection"
	datasourceName := "data.oci_database_db_node_console_connections.test_db_node_console_connections"
	singularDatasourceName := "data.oci_database_db_node_console_connection.test_db_node_console_connection"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseDbNodeConsoleConnectionResourcesExaCCDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Optional, acctest.Create, DatabaseDbNodeConsoleConnectionRepresentationExaCC), "database", "dbNodeConsoleConnection", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeConsoleConnectionResourcesExaCCDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Required, acctest.Create, DatabaseDbNodeConsoleConnectionRepresentationExaCC),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
				resource.TestCheckResourceAttr(resourceName, "public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCldnh+UpWXKNlkRTamqCnmihMMLZ5UofvO5Thxr+dpp427HcRiZaJnrnpbrh+T5c+I+bTE9JQh4Ydsk9hEg1LWXIqCYPTDifaWneJ6o+xpbjW8benndWDR6y3XJ2yLdXOJ9S8HWMXQLkAGBEkxsqNAtdlt2U3RjRbO+4g+PdNBLycZ8vXwaD+GOxq08GtxQTQuzQMWUjt56TVH7OG6v/1nEHTQ5meYHtTP4cx1YqqOQBWyZG4Ikq9Ej8YegmP3Tbf9SRRcrXG0qASNIevofGsmhvhvgXhNtENJpEvf10aJyXmRDUAfUwGSFCKyHVEiCX9MzPLzUlvSbw2Ls1NVkZ3v mytest.vpn.oracle.com"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeConsoleConnectionResourcesExaCCDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeConsoleConnectionResourcesExaCCDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Optional, acctest.Create, DatabaseDbNodeConsoleConnectionRepresentationExaCC),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_string"),
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
				resource.TestCheckResourceAttrSet(resourceName, "fingerprint"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCldnh+UpWXKNlkRTamqCnmihMMLZ5UofvO5Thxr+dpp427HcRiZaJnrnpbrh+T5c+I+bTE9JQh4Ydsk9hEg1LWXIqCYPTDifaWneJ6o+xpbjW8benndWDR6y3XJ2yLdXOJ9S8HWMXQLkAGBEkxsqNAtdlt2U3RjRbO+4g+PdNBLycZ8vXwaD+GOxq08GtxQTQuzQMWUjt56TVH7OG6v/1nEHTQ5meYHtTP4cx1YqqOQBWyZG4Ikq9Ej8YegmP3Tbf9SRRcrXG0qASNIevofGsmhvhvgXhNtENJpEvf10aJyXmRDUAfUwGSFCKyHVEiCX9MzPLzUlvSbw2Ls1NVkZ3v mytest.vpn.oracle.com"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeConsoleConnectionResourcesExaCCDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Optional, acctest.Update, DatabaseDbNodeConsoleConnectionRepresentationExaCC),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_string"),
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
				resource.TestCheckResourceAttrSet(resourceName, "fingerprint"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "public_key", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCldnh+UpWXKNlkRTamqCnmihMMLZ5UofvO5Thxr+dpp427HcRiZaJnrnpbrh+T5c+I+bTE9JQh4Ydsk9hEg1LWXIqCYPTDifaWneJ6o+xpbjW8benndWDR6y3XJ2yLdXOJ9S8HWMXQLkAGBEkxsqNAtdlt2U3RjRbO+4g+PdNBLycZ8vXwaD+GOxq08GtxQTQuzQMWUjt56TVH7OG6v/1nEHTQ5meYHtTP4cx1YqqOQBWyZG4Ikq9Ej8YegmP3Tbf9SRRcrXG0qASNIevofGsmhvhvgXhNtENJpEvf10aJyXmRDUAfUwGSFCKyHVEiCX9MzPLzUlvSbw2Ls1NVkZ3v mytest.vpn.oracle.com"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					_, id, _ := tf_database.ParseDbNodeConsoleConnectionCompositeId(resId)
					consoleIdVariableStr = fmt.Sprintf("variable \"console_id\" { default = \"%s\" }\n", id)
					print("THe testing id of connection is:" + id)
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + consoleIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node_console_connections", "test_db_node_console_connections", acctest.Optional, acctest.Update, DatabaseDatabaseDbNodeConsoleConnectionDataSourceRepresentationExaCC) +
				compartmentIdVariableStr + DatabaseDbNodeConsoleConnectionResourcesExaCCDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Optional, acctest.Update, DatabaseDbNodeConsoleConnectionRepresentationExaCC),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_node_id"),

				resource.TestCheckResourceAttr(datasourceName, "console_connections.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.connection_string"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.db_node_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.fingerprint"),
				resource.TestCheckResourceAttr(datasourceName, "console_connections.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.service_host_key_fingerprint"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.state"),
			),
		},
		// verify singular datasource
		{
			Config: config + dbNodeIdVariableStr + consoleConnectionIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Required, acctest.Create, DatabaseDatabaseDbNodeConsoleConnectionSingularDataSourceRepresentationExaCC) +
				compartmentIdVariableStr + DatabaseDbNodeConsoleConnectionResourceConfigExaCC,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_string"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fingerprint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_host_key_fingerprint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// verify resource import
		{
			Config:            config + DbNodeConsoleConnectionRequiredOnlyResourceExaCC,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"public_key",
			},
			ResourceName: resourceName,
		},
	})
}
