// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MysqlReplicaRequiredOnlyResource = MysqlReplicaResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_replica", "test_replica", acctest.Required, acctest.Create, MysqlReplicaRepresentation)

	MysqlReplicaResourceConfig = MysqlReplicaResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_replica", "test_replica", acctest.Optional, acctest.Update, MysqlReplicaRepresentation)

	MysqlMysqlReplicaSingularDataSourceRepresentation = map[string]interface{}{
		"replica_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_replica.test_replica.id}`},
	}

	MysqlReplicaDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.configuration_id}`},
		"db_system_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"is_up_to_date":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"replica_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_replica.test_replica.id}`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlReplicaDataSourceFilterRepresentation},
	}

	MysqlReplicaDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_mysql_replica.test_replica.id}`}},
	}

	MysqlReplicaRepresentation = map[string]interface{}{
		"db_system_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_delete_protected": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForMysqlReplica},
		"replica_overrides":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlReplicaReplicaOverridesRepresentation},
	}

	MysqlReplicaReplicaOverridesRepresentation = map[string]interface{}{
		"configuration_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.configuration_id}`},
		"mysql_version":    acctest.Representation{RepType: acctest.Optional, Create: `8.0.35`, Update: `8.1.0`},
		"shape_name":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.shape_name}`},
	}

	ignoreDefinedTagsChangesForMysqlReplica = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"defined_tags"}},
	}

	mysqlDbSystemRepresentationWithReadReplica = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlHAConfigurationOCID[var.region]}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.4.64GB`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"backup_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemBackupPolicyRepresentation},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `MySQL Database Service`, Update: `description2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `DBSystem001`, Update: `displayName2`},
		"fault_domain":            acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":          acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`},
		"is_highly_available":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"maintenance":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemMaintenanceRepresentation},
		"port":                    acctest.Representation{RepType: acctest.Optional, Create: `3306`},
		"port_x":                  acctest.Representation{RepType: acctest.Optional, Create: `33306`},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForMysqlRep},
	}

	MysqlReplicaResourceDependencies = MysqlMysqlDbSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Create, mysqlDbSystemRepresentationWithReadReplica)
)

// issue-routing-tag: mysql/default
func TestMysqlReplicaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlReplicaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_replica.test_replica"
	dbSystemResourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"
	datasourceName := "data.oci_mysql_replicas.test_replicas"
	singularDatasourceName := "data.oci_mysql_replica.test_replica"

	var resId, resId2, resId3 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MysqlReplicaResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_replica", "test_replica", acctest.Optional, acctest.Create, MysqlReplicaRepresentation), "mysql", "replica", t)

	acctest.ResourceTest(t, testAccCheckMysqlReplicaDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MysqlReplicaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_replica", "test_replica", acctest.Required, acctest.Create, MysqlReplicaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
				func(s *terraform.State) (err error) {
					resId3, err = acctest.FromInstanceState(s, dbSystemResourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MysqlReplicaResourceDependencies,
		},
		// verify Create with optionals
		{
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId3, dbSystemAvailableWaitCondition, time.Duration(10*time.Minute),
				dbSystemFetchOperation, "mysql", false),
			Config: config + compartmentIdVariableStr + MysqlReplicaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_replica", "test_replica", acctest.Optional, acctest.Create, MysqlReplicaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttr(resourceName, "is_delete_protected", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "mysql_version"),
				resource.TestCheckResourceAttrSet(resourceName, "port"),
				resource.TestCheckResourceAttrSet(resourceName, "port_x"),
				resource.TestCheckResourceAttr(resourceName, "replica_overrides.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "replica_overrides.0.configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "replica_overrides.0.mysql_version", "8.0.35"),
				resource.TestCheckResourceAttrSet(resourceName, "replica_overrides.0.shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId3, dbSystemAvailableWaitCondition, time.Duration(10*time.Minute),
				dbSystemFetchOperation, "mysql", false),
			Config: config + compartmentIdVariableStr + MysqlReplicaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_replica", "test_replica", acctest.Optional, acctest.Update, MysqlReplicaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttr(resourceName, "is_delete_protected", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "mysql_version"),
				resource.TestCheckResourceAttrSet(resourceName, "port"),
				resource.TestCheckResourceAttrSet(resourceName, "port_x"),
				resource.TestCheckResourceAttr(resourceName, "replica_overrides.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "replica_overrides.0.configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "replica_overrides.0.mysql_version", "8.1.0"),
				resource.TestCheckResourceAttrSet(resourceName, "replica_overrides.0.shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_replicas", "test_replicas", acctest.Optional, acctest.Update, MysqlReplicaDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlReplicaResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_replica", "test_replica", acctest.Optional, acctest.Update, MysqlReplicaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "configuration_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_up_to_date", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "replica_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "replicas.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.configuration_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "replicas.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "replicas.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.fault_domain"),
				resource.TestCheckResourceAttr(datasourceName, "replicas.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.ip_address"),
				resource.TestCheckResourceAttr(datasourceName, "replicas.0.is_delete_protected", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.mysql_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.port"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.port_x"),
				resource.TestCheckResourceAttr(datasourceName, "replicas.0.replica_overrides.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.replica_overrides.0.configuration_id"),
				resource.TestCheckResourceAttr(datasourceName, "replicas.0.replica_overrides.0.mysql_version", "8.1.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.replica_overrides.0.shape_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.shape_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "replicas.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_replica", "test_replica", acctest.Required, acctest.Create, MysqlMysqlReplicaSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlReplicaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replica_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fault_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_delete_protected", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mysql_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "port"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "port_x"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replica_overrides.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replica_overrides.0.mysql_version", "8.1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "secure_connections.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + MysqlReplicaRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMysqlReplicaDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ReplicasClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_replica" {
			noResourceFound = false
			request := oci_mysql.GetReplicaRequest{}

			tmp := rs.Primary.ID
			request.ReplicaId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")

			response, err := client.GetReplica(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.ReplicaLifecycleStateDeleted): true,
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("MysqlReplica") {
		resource.AddTestSweepers("MysqlReplica", &resource.Sweeper{
			Name:         "MysqlReplica",
			Dependencies: acctest.DependencyGraph["replica"],
			F:            sweepMysqlReplicaResource,
		})
	}
}

func sweepMysqlReplicaResource(compartment string) error {
	replicasClient := acctest.GetTestClients(&schema.ResourceData{}).ReplicasClient()
	replicaIds, err := getMysqlReplicaIds(compartment)
	if err != nil {
		return err
	}
	for _, replicaId := range replicaIds {
		if ok := acctest.SweeperDefaultResourceId[replicaId]; !ok {
			deleteReplicaRequest := oci_mysql.DeleteReplicaRequest{}

			deleteReplicaRequest.ReplicaId = &replicaId

			deleteReplicaRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")
			_, error := replicasClient.DeleteReplica(context.Background(), deleteReplicaRequest)
			if error != nil {
				fmt.Printf("Error deleting Replica %s %s, It is possible that the resource is already deleted. Please verify manually \n", replicaId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &replicaId, MysqlReplicaSweepWaitCondition, time.Duration(3*time.Minute),
				MysqlReplicaSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func getMysqlReplicaIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ReplicaId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	replicasClient := acctest.GetTestClients(&schema.ResourceData{}).ReplicasClient()

	listReplicasRequest := oci_mysql.ListReplicasRequest{}
	listReplicasRequest.CompartmentId = &compartmentId
	listReplicasRequest.LifecycleState = oci_mysql.ReplicaSummaryLifecycleStateActive
	listReplicasResponse, err := replicasClient.ListReplicas(context.Background(), listReplicasRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Replica list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, replica := range listReplicasResponse.Items {
		id := *replica.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ReplicaId", id)
	}
	return resourceIds, nil
}

func MysqlReplicaSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if replicaResponse, ok := response.Response.(oci_mysql.GetReplicaResponse); ok {
		return replicaResponse.LifecycleState != oci_mysql.ReplicaLifecycleStateDeleted
	}
	return false
}

func MysqlReplicaSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ReplicasClient().GetReplica(context.Background(), oci_mysql.GetReplicaRequest{
		ReplicaId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func dbSystemAvailableWaitCondition(response common.OCIOperationResponse) bool {
	if dbSystemResponse, ok := response.Response.(oci_mysql.GetDbSystemResponse); ok {
		return dbSystemResponse.LifecycleState != oci_mysql.DbSystemLifecycleStateActive
	}
	return false
}

func dbSystemFetchOperation(client *tf_client.OracleClients, dbSystemId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbSystemClient().GetDbSystem(context.Background(), oci_mysql.GetDbSystemRequest{
		DbSystemId: dbSystemId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
