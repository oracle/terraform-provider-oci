// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v46/common"
	oci_database "github.com/oracle/oci-go-sdk/v46/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PluggableDatabaseRequiredOnlyResource = PluggableDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Required, Create, pluggableDatabaseRepresentation)

	PluggableDatabaseResourceConfig = PluggableDatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Optional, Update, pluggableDatabaseRepresentation)

	pluggableDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"pluggable_database_id": Representation{repType: Required, create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
	}
	ignoreChangesPdbepresentation = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}
	pluggableDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"pdb_name":       Representation{repType: Optional, create: `SalesPdb`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, pluggableDatabaseDataSourceFilterRepresentation}}
	pluggableDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_pluggable_database.test_pluggable_database.id}`}},
	}

	pluggableDatabaseRepresentation = map[string]interface{}{
		"container_database_id": Representation{repType: Required, create: `${data.oci_database_database.t.id}`},
		"pdb_admin_password":    Representation{repType: Required, create: `BEstrO0ng_#11`},
		"pdb_name":              Representation{repType: Required, create: `SalesPdb`},
		"tde_wallet_password":   Representation{repType: Required, create: `BEstrO0ng_#11`},
		"defined_tags":          Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":         Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"lifecycle":             RepresentationGroup{Required, ignoreChangesLBRepresentation},
		"depends_on":            Representation{repType: Required, create: []string{"oci_database_db_system.t"}},
	}

	ResourcePluggableDatabaseBaseConfig = `

	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.1.0.0/16"
		display_name = "-tf-vcn"
		dns_label = "tfvcn"
	}

	resource "oci_core_route_table" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		route_rules {
			cidr_block = "0.0.0.0/0"
			network_entity_id = "${oci_core_internet_gateway.t.id}"
		}
	}
	resource "oci_core_internet_gateway" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		display_name = "-tf-internet-gateway"
	}

	resource "oci_core_subnet" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		cidr_block          = "10.1.20.0/24"
		display_name        = "TFSubnet1"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_route_table.t.id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dns_label           = "tfsubnet"
	}
	resource "oci_core_subnet" "t2" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		cidr_block          = "10.1.21.0/24"
		display_name        = "TFSubnet2"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		route_table_id      = "${oci_core_route_table.t.id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dns_label           = "tfsubnet2"
	}
	resource "oci_core_network_security_group" "test_network_security_group" {
         compartment_id  = "${var.compartment_id}"
		 vcn_id            = "${oci_core_virtual_network.t.id}"
         display_name      =  "displayName"
    }

	resource "oci_core_network_security_group" "test_network_security_group2" {
		compartment_id = "${var.compartment_id}"
		vcn_id            = "${oci_core_virtual_network.t.id}"
	}`

	dbSystemForPluggableDbRepresentation = `
		resource "oci_database_db_system" "t" {
			compartment_id = "${var.compartment_id}"
			subnet_id = "${oci_core_subnet.t.id}"
			database_edition = "ENTERPRISE_EDITION"
			availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
			disk_redundancy = "NORMAL"
			shape = "VM.Standard2.4"
			ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
			display_name = "-tf-dbSystem-001"
			domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
			hostname = "myOracleDB" // this will be lowercased server side
			data_storage_size_in_gb = "256"
			license_model = "LICENSE_INCLUDED"
			node_count = "1"
			fault_domains = ["FAULT-DOMAIN-1"]
			db_home {
				db_version = "19.11.0.0"
				display_name = "-tf-db-home"
				database {
					admin_password = "BEstrO0ng_#11"
					db_name = "aTFdb"
					character_set = "AL32UTF8"
					defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
					freeform_tags = {"Department" = "Finance"}
					ncharacter_set = "AL16UTF16"
					db_workload = "OLTP"
					pdb_name = "pdbName"
				}
			}
			db_system_options {
				storage_management = "LVM"
			}
			defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
			freeform_tags = {"Department" = "Finance"}
			nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
			lifecycle {
				ignore_changes = [
					db_home.0.db_version,
					defined_tags,
					db_home.0.database.0.defined_tags,
				]
			}
		}
		data "oci_database_db_systems" "t" {
			compartment_id = "${var.compartment_id}"
			filter {
				name   = "id"
				values = ["${oci_database_db_system.t.id}"]
			}
		}
		data "oci_database_db_homes" "t" {
			compartment_id = "${var.compartment_id}"
			db_system_id = "${oci_database_db_system.t.id}"
			filter {
				name   = "db_system_id"
				values = ["${oci_database_db_system.t.id}"]
			}
		}
		data "oci_database_db_home" "t" {
			db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
		}
		data "oci_database_databases" "t" {
			compartment_id = "${var.compartment_id}"
			db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
			filter {
				name   = "db_name"
				values = ["${oci_database_db_system.t.db_home.0.database.0.db_name}"]
			}
		}
		data "oci_database_database" "t" {
			  database_id = "${data.oci_database_databases.t.databases.0.id}"
		}`

	PluggableDatabaseResourceDependencies = AvailabilityDomainConfig + ResourcePluggableDatabaseBaseConfig +
		DefinedTagsDependencies + dbSystemForPluggableDbRepresentation
)

// issue-routing-tag: database/default
func TestDatabasePluggableDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabasePluggableDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_pluggable_database.test_pluggable_database"
	datasourceName := "data.oci_database_pluggable_databases.test_pluggable_databases"
	singularDatasourceName := "data.oci_database_pluggable_database.test_pluggable_database"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+PluggableDatabaseResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Optional, Create, pluggableDatabaseRepresentation), "database", "pluggableDatabase", t)

	ResourceTest(t, testAccCheckDatabasePluggableDatabaseDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + PluggableDatabaseResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Required, Create, pluggableDatabaseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),
				resource.TestCheckResourceAttr(resourceName, "tde_wallet_password", "BEstrO0ng_#11"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + PluggableDatabaseResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + PluggableDatabaseResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Optional, Create, pluggableDatabaseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),
				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "tde_wallet_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + PluggableDatabaseResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Optional, Update, pluggableDatabaseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "open_mode"),
				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "tde_wallet_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_pluggable_databases", "test_pluggable_databases", Optional, Update, pluggableDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + PluggableDatabaseResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Optional, Update, pluggableDatabaseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "pdb_name", "SalesPdb"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.0.connection_strings.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.open_mode"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.0.pdb_name", "SalesPdb"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Required, Update, pluggableDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PluggableDatabaseResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pluggable_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "open_mode"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pdb_name", "SalesPdb"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},

		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + PluggableDatabaseResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"pdb_admin_password",
				"tde_wallet_password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabasePluggableDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_pluggable_database" {
			noResourceFound = false
			request := oci_database.GetPluggableDatabaseRequest{}

			tmp := rs.Primary.ID
			request.PluggableDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetPluggableDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.PluggableDatabaseLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabasePluggableDatabase") {
		resource.AddTestSweepers("DatabasePluggableDatabase", &resource.Sweeper{
			Name:         "DatabasePluggableDatabase",
			Dependencies: DependencyGraph["pluggableDatabase"],
			F:            sweepDatabasePluggableDatabaseResource,
		})
	}
}

func sweepDatabasePluggableDatabaseResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	pluggableDatabaseIds, err := getPluggableDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, pluggableDatabaseId := range pluggableDatabaseIds {
		if ok := SweeperDefaultResourceId[pluggableDatabaseId]; !ok {
			deletePluggableDatabaseRequest := oci_database.DeletePluggableDatabaseRequest{}

			deletePluggableDatabaseRequest.PluggableDatabaseId = &pluggableDatabaseId

			deletePluggableDatabaseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeletePluggableDatabase(context.Background(), deletePluggableDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting PluggableDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", pluggableDatabaseId, error)
				continue
			}
			waitTillCondition(testAccProvider, &pluggableDatabaseId, pluggableDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				pluggableDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getPluggableDatabaseIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "PluggableDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listPluggableDatabasesRequest := oci_database.ListPluggableDatabasesRequest{}
	listPluggableDatabasesRequest.CompartmentId = &compartmentId
	listPluggableDatabasesRequest.LifecycleState = oci_database.PluggableDatabaseSummaryLifecycleStateAvailable
	listPluggableDatabasesResponse, err := databaseClient.ListPluggableDatabases(context.Background(), listPluggableDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PluggableDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, pluggableDatabase := range listPluggableDatabasesResponse.Items {
		id := *pluggableDatabase.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "PluggableDatabaseId", id)
	}
	return resourceIds, nil
}

func pluggableDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if pluggableDatabaseResponse, ok := response.Response.(oci_database.GetPluggableDatabaseResponse); ok {
		return pluggableDatabaseResponse.LifecycleState != oci_database.PluggableDatabaseLifecycleStateTerminated
	}
	return false
}

func pluggableDatabaseSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetPluggableDatabase(context.Background(), oci_database.GetPluggableDatabaseRequest{
		PluggableDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
