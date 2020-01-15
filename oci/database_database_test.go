// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	exaVcnRepresentation = map[string]interface{}{
		"cidr_block":     Representation{repType: Required, create: `10.1.0.0/16`},
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `-tf-vcn`},
		"dns_label":      Representation{repType: Optional, create: `tfvcn`},
	}

	exaSecurityListRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"vcn_id":                 Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"display_name":           Representation{repType: Optional, create: `ExadataSecurityList`},
		"egress_security_rules":  []RepresentationGroup{{Required, exaSecurityListEgressSecurityRulesICMPRepresentation}, {Optional, exaSecurityListEgressSecurityRulesTCPRepresentation}},
		"ingress_security_rules": []RepresentationGroup{{Required, exaSecurityListIngressSecurityRulesICMPRepresentation}, {Optional, exaSecurityListIngressSecurityRulesTCPRepresentation}},
	}

	exaSecurityListIngressSecurityRulesICMPRepresentation = map[string]interface{}{
		"protocol": Representation{repType: Required, create: `1`},
		"source":   Representation{repType: Required, create: `10.1.22.0/24`},
	}
	exaSecurityListIngressSecurityRulesTCPRepresentation = map[string]interface{}{
		"protocol": Representation{repType: Required, create: `6`},
		"source":   Representation{repType: Required, create: `10.1.22.0/24`},
	}
	exaSecurityListEgressSecurityRulesICMPRepresentation = map[string]interface{}{
		"protocol":    Representation{repType: Required, create: `1`},
		"destination": Representation{repType: Required, create: `10.1.22.0/24`},
	}
	exaSecurityListEgressSecurityRulesTCPRepresentation = map[string]interface{}{
		"protocol":    Representation{repType: Required, create: `6`},
		"destination": Representation{repType: Required, create: `10.1.22.0/24`},
	}

	exaSubnetRepresentation = map[string]interface{}{
		"cidr_block":          Representation{repType: Required, create: `10.1.22.0/24`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"vcn_id":              Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain": Representation{repType: Optional, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name}")}`},
		"dhcp_options_id":     Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"display_name":        Representation{repType: Optional, create: `ExadataSubnet`},
		"dns_label":           Representation{repType: Optional, create: `subnetexadata1`},
		"route_table_id":      Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.default_route_table_id}`},
		"security_list_ids":   Representation{repType: Optional, create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`, `${oci_core_security_list.exadata_shapes_security_list.id}`}},
	}
	exaBackupSubnetRepresentation = map[string]interface{}{
		"cidr_block":          Representation{repType: Required, create: `10.1.23.0/24`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"vcn_id":              Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain": Representation{repType: Optional, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name}")}`},
		"dhcp_options_id":     Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`},
		"display_name":        Representation{repType: Optional, create: `ExadataBackupSubnet`},
		"dns_label":           Representation{repType: Optional, create: `subnetexadata2`},
		"route_table_id":      Representation{repType: Optional, create: `${oci_core_vcn.test_vcn.default_route_table_id}`},
		"security_list_ids":   Representation{repType: Optional, create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}},
	}

	exadbSystemRepresentation = map[string]interface{}{
		"availability_domain":     Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.2.name}")}`},
		"backup_subnet_id":        Representation{repType: Required, create: `${oci_core_subnet.exadata_backup_subnet.id}`},
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"database_edition":        Representation{repType: Required, create: `ENTERPRISE_EDITION_EXTREME_PERFORMANCE`},
		"db_home":                 RepresentationGroup{Required, exadbSystemDbHomeRepresentation},
		"hostname":                Representation{repType: Required, create: `myOracleDB`},
		"shape":                   Representation{repType: Required, create: `Exadata.Quarter1.84`},
		"ssh_public_keys":         Representation{repType: Required, create: []string{`ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`}},
		"subnet_id":               Representation{repType: Required, create: `${oci_core_subnet.exadata_subnet.id}`},
		"cpu_core_count":          Representation{repType: Optional, create: `22`},
		"data_storage_size_in_gb": Representation{repType: Optional, create: `256`},
		"disk_redundancy":         Representation{repType: Optional, create: `NORMAL`},
		"display_name":            Representation{repType: Optional, create: `tfDbSystemTestExadata`},
		"domain":                  Representation{repType: Optional, create: `${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_vcn.test_vcn.dns_label}.oraclevcn.com`},
		"license_model":           Representation{repType: Optional, create: `LICENSE_INCLUDED`},
		"node_count":              Representation{repType: Optional, create: `1`},
	}
	exadbSystemDbHomeRepresentation = map[string]interface{}{
		"database":     RepresentationGroup{Required, exadbSystemDbHomeDatabaseRepresentation},
		"db_version":   Representation{repType: Optional, create: `12.1.0.2`},
		"display_name": Representation{repType: Optional, create: `dbHome1`},
	}
	exadbSystemDbHomeDatabaseRepresentation = map[string]interface{}{
		"admin_password": Representation{repType: Required, create: `BEstrO0ng_#11`},
		"db_name":        Representation{repType: Optional, create: `tfDbName`},
	}

	ExaBaseDependencies = generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, exaVcnRepresentation) +
		generateResourceFromRepresentationMap("oci_core_security_list", "exadata_shapes_security_list", Optional, Create, exaSecurityListRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "exadata_subnet", Optional, Create, exaSubnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "exadata_backup_subnet", Optional, Create, exaBackupSubnetRepresentation) +
		generateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", Optional, Create, exadbSystemRepresentation)

	DatabaseRequiredOnlyResource = DatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseRepresentation)

	DatabaseResourceConfig = DatabaseResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Update, databaseRepresentation)

	databaseSingularDataSourceRepresentation = map[string]interface{}{
		"database_id": Representation{repType: Required, create: `${oci_database_database.test_database.id}`},
	}

	databaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"db_home_id":     Representation{repType: Optional, create: `${oci_database_db_home.test_db_home.id}`},
		"db_name":        Representation{repType: Optional, create: `myTestDb`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, databaseDataSourceFilterRepresentation}}
	databaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_database.test_database.id}`}},
	}

	databaseRepresentation = map[string]interface{}{
		"database":   RepresentationGroup{Required, databaseDatabaseRepresentation},
		"db_home_id": Representation{repType: Required, create: `${oci_database_db_home.test_db_home.id}`},
		"source":     Representation{repType: Required, create: `NONE`},
		"db_version": Representation{repType: Optional, create: `12.1.0.2`},
	}
	databaseDatabaseRepresentation = map[string]interface{}{
		"admin_password":   Representation{repType: Required, create: `BEstrO0ng_#11`},
		"db_name":          Representation{repType: Required, create: `myTestDb`},
		"character_set":    Representation{repType: Optional, create: `AL32UTF8`},
		"db_backup_config": RepresentationGroup{Optional, databaseDatabaseDbBackupConfigRepresentation},
		"db_unique_name":   Representation{repType: Optional, create: `myTestDb_12`},
		"db_workload":      Representation{repType: Optional, create: `OLTP`},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"freeformTags": "freeformTags"}, update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":   Representation{repType: Optional, create: `AL16UTF16`},
		"pdb_name":         Representation{repType: Optional, create: `pdbName`},
	}
	databaseDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":     Representation{repType: Optional, create: `true`},
		"auto_backup_window":      Representation{repType: Optional, create: `SLOT_TWO`, update: `SLOT_THREE`},
		"recovery_window_in_days": Representation{repType: Optional, create: `10`, update: `30`},
	}
	databaseDatabaseDbBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type": Representation{repType: Required, create: `NFS`},
		"id":   Representation{repType: Optional, create: `${oci_database_backup_destination.test_backup_destination.id}`},
	}

	DatabaseResourceDependencies = ExaBaseDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
		generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", Required, Create, dbHomeRepresentationSourceNone)
)

func TestDatabaseDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_database.test_database"
	datasourceName := "data.oci_database_databases.test_databases"
	singularDatasourceName := "data.oci_database_database.test_database"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
					resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
					resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DatabaseResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Create, databaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_TWO"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName, "db_name", "myTestDb"),
					resource.TestCheckResourceAttr(resourceName, "db_unique_name", "myTestDb_12"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName, "pdb_name", "pdbName"),
					resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
					resource.TestCheckResourceAttrSet(resourceName, "db_name"),
					resource.TestCheckResourceAttrSet(resourceName, "db_unique_name"),
					resource.TestCheckResourceAttr(resourceName, "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Update, databaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_enabled", "true"),
					//resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_THREE"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "30"),
					resource.TestCheckResourceAttr(resourceName, "db_name", "myTestDb"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName, "pdb_name", "pdbName"),
					resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
					resource.TestCheckResourceAttrSet(resourceName, "db_name"),
					resource.TestCheckResourceAttrSet(resourceName, "db_unique_name"),
					resource.TestCheckResourceAttr(resourceName, "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
					generateDataSourceFromRepresentationMap("oci_database_databases", "test_databases", Optional, Update, databaseDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Update, databaseRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "db_home_id"),
					resource.TestCheckResourceAttr(datasourceName, "db_name", "myTestDb"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.character_set"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "databases.0.db_backup_config.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_home_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_unique_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_workload"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.ncharacter_set"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.pdb_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "character_set"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_unique_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_workload"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ncharacter_set"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "pdb_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}

func testAccCheckDatabaseDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_database" {
			noResourceFound = false
			request := oci_database.GetDatabaseRequest{}

			tmp := rs.Primary.ID
			request.DatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.DatabaseLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("DatabaseDatabase") {
		resource.AddTestSweepers("DatabaseDatabase", &resource.Sweeper{
			Name:         "DatabaseDatabase",
			Dependencies: DependencyGraph["database"],
			F:            sweepDatabaseDatabaseResource,
		})
	}
}

func sweepDatabaseDatabaseResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient
	databaseIds, err := getDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseId := range databaseIds {
		if ok := SweeperDefaultResourceId[databaseId]; !ok {
			deleteDatabaseRequest := oci_database.DeleteDatabaseRequest{}

			deleteDatabaseRequest.DatabaseId = &databaseId

			deleteDatabaseRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDatabase(context.Background(), deleteDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting Database %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseId, error)
				continue
			}
			waitTillCondition(testAccProvider, &databaseId, databaseSweepWaitCondition, time.Duration(3*time.Minute),
				databaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient

	listDatabasesRequest := oci_database.ListDatabasesRequest{}
	listDatabasesRequest.CompartmentId = &compartmentId
	listDatabasesRequest.LifecycleState = oci_database.DatabaseSummaryLifecycleStateAvailable
	listDatabasesResponse, err := databaseClient.ListDatabases(context.Background(), listDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Database list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, database := range listDatabasesResponse.Items {
		id := *database.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseId", id)
	}
	return resourceIds, nil
}

func databaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseResponse, ok := response.Response.(oci_database.GetDatabaseResponse); ok {
		return databaseResponse.LifecycleState != oci_database.DatabaseLifecycleStateTerminated
	}
	return false
}

func databaseSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient.GetDatabase(context.Background(), oci_database.GetDatabaseRequest{
		DatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
