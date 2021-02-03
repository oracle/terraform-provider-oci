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
	"github.com/oracle/oci-go-sdk/v35/common"
	oci_database "github.com/oracle/oci-go-sdk/v35/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousVmClusterRequiredOnlyResource = AutonomousVmClusterResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Required, Create, autonomousVmClusterRepresentation)

	AutonomousVmClusterResourceConfig = AutonomousVmClusterResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Optional, Update, autonomousVmClusterRepresentation)

	autonomousVmClusterSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id": Representation{repType: Required, create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
	}

	autonomousVmClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":              Representation{repType: Optional, create: `autonomousVmCluster`},
		"exadata_infrastructure_id": Representation{repType: Optional, create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"state":                     Representation{repType: Optional, create: `AVAILABLE`},
		"filter":                    RepresentationGroup{Required, autonomousVmClusterDataSourceFilterRepresentation}}
	autonomousVmClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`}},
	}

	autonomousVmClusterRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":              Representation{repType: Required, create: `autonomousVmCluster`},
		"exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"vm_cluster_network_id":     Representation{repType: Required, create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":   Representation{repType: Optional, create: `false`},
		"license_model":             Representation{repType: Optional, create: `LICENSE_INCLUDED`},
		"time_zone":                 Representation{repType: Optional, create: `US/Pacific`},
	}

	AutonomousVmClusterResourceDependencies = generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create,
		representationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, map[string]interface{}{"activation_file": Representation{repType: Required, create: activationFilePath}})) +
		generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Required, Create,
			representationCopyWithNewProperties(vmClusterNetworkRepresentation, map[string]interface{}{"validate_vm_cluster_network": Representation{repType: Required, create: "true"}})) +
		DefinedTagsDependencies
)

func TestDatabaseAutonomousVmClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster"
	datasourceName := "data.oci_database_autonomous_vm_clusters.test_autonomous_vm_clusters"
	singularDatasourceName := "data.oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseAutonomousVmClusterDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousVmClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Required, Create, autonomousVmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "autonomousVmCluster"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousVmClusterResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousVmClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Optional, Create, autonomousVmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "autonomousVmCluster"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
					resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

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

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousVmClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Optional, Create,
						representationCopyWithNewProperties(autonomousVmClusterRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "autonomousVmCluster"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
					resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutonomousVmClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Optional, Update, autonomousVmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "autonomousVmCluster"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
					resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_vm_clusters", "test_autonomous_vm_clusters", Optional, Update, autonomousVmClusterDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousVmClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Optional, Update, autonomousVmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "autonomousVmCluster"),
					resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.available_cpus"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.available_data_storage_size_in_tbs"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.cpus_enabled"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.data_storage_size_in_tbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.db_node_storage_size_in_gbs"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.display_name", "autonomousVmCluster"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.is_local_backup_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.memory_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_vm_clusters.0.time_zone", "US/Pacific"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_clusters.0.vm_cluster_network_id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Required, Create, autonomousVmClusterSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousVmClusterResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_vm_cluster_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "available_cpus"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "available_data_storage_size_in_tbs"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cpus_enabled"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_storage_size_in_tbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_storage_size_in_gbs"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "autonomousVmCluster"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_local_backup_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "US/Pacific"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + AutonomousVmClusterResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckDatabaseAutonomousVmClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_vm_cluster" {
			noResourceFound = false
			request := oci_database.GetAutonomousVmClusterRequest{}

			tmp := rs.Primary.ID
			request.AutonomousVmClusterId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetAutonomousVmCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AutonomousVmClusterLifecycleStateTerminated): true,
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
	if !inSweeperExcludeList("DatabaseAutonomousVmCluster") {
		resource.AddTestSweepers("DatabaseAutonomousVmCluster", &resource.Sweeper{
			Name:         "DatabaseAutonomousVmCluster",
			Dependencies: DependencyGraph["autonomousVmCluster"],
			F:            sweepDatabaseAutonomousVmClusterResource,
		})
	}
}

func sweepDatabaseAutonomousVmClusterResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	autonomousVmClusterIds, err := getAutonomousVmClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, autonomousVmClusterId := range autonomousVmClusterIds {
		if ok := SweeperDefaultResourceId[autonomousVmClusterId]; !ok {
			deleteAutonomousVmClusterRequest := oci_database.DeleteAutonomousVmClusterRequest{}

			deleteAutonomousVmClusterRequest.AutonomousVmClusterId = &autonomousVmClusterId

			deleteAutonomousVmClusterRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteAutonomousVmCluster(context.Background(), deleteAutonomousVmClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousVmCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", autonomousVmClusterId, error)
				continue
			}
			waitTillCondition(testAccProvider, &autonomousVmClusterId, autonomousVmClusterSweepWaitCondition, time.Duration(3*time.Minute),
				autonomousVmClusterSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getAutonomousVmClusterIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "AutonomousVmClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listAutonomousVmClustersRequest := oci_database.ListAutonomousVmClustersRequest{}
	listAutonomousVmClustersRequest.CompartmentId = &compartmentId
	listAutonomousVmClustersRequest.LifecycleState = oci_database.AutonomousVmClusterSummaryLifecycleStateAvailable
	listAutonomousVmClustersResponse, err := databaseClient.ListAutonomousVmClusters(context.Background(), listAutonomousVmClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AutonomousVmCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, autonomousVmCluster := range listAutonomousVmClustersResponse.Items {
		id := *autonomousVmCluster.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "AutonomousVmClusterId", id)
	}
	return resourceIds, nil
}

func autonomousVmClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousVmClusterResponse, ok := response.Response.(oci_database.GetAutonomousVmClusterResponse); ok {
		return autonomousVmClusterResponse.LifecycleState != oci_database.AutonomousVmClusterLifecycleStateTerminated
	}
	return false
}

func autonomousVmClusterSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetAutonomousVmCluster(context.Background(), oci_database.GetAutonomousVmClusterRequest{
		AutonomousVmClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
