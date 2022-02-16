// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VmClusterRequiredOnlyResource = VmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, vmClusterRepresentation)

	VmClusterResourceConfig = VmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Update, vmClusterRepresentation)

	vmClusterSingularDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	vmClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `vmCluster`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: vmClusterDataSourceFilterRepresentation}}
	vmClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_vm_cluster.test_vm_cluster.id}`}},
	}

	vmClusterRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":              acctest.Representation{RepType: acctest.Required, Create: `4`, Update: `6`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `vmCluster`},
		"exadata_infrastructure_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"gi_version":                  acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0.0`},
		"ssh_public_keys":             acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"vm_cluster_network_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"data_storage_size_in_tbs":    acctest.Representation{RepType: acctest.Optional, Create: `84`, Update: `86`},
		"db_node_storage_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `120`, Update: `160`},
		"db_servers":                  acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`, `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`}},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_sparse_diskgroup_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":               acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"memory_size_in_gbs":          acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `90`},
		"time_zone":                   acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
	}

	VmClusterResourceDependencies = VmClusterNetworkValidatedResourceConfig
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_vm_cluster.test_vm_cluster"
	datasourceName := "data.oci_database_vm_clusters.test_vm_clusters"
	singularDatasourceName := "data.oci_database_vm_cluster.test_vm_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+VmClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Create, vmClusterRepresentation), "database", "vmCluster", t)

	acctest.ResourceTest(t, testAccCheckDatabaseVmClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VmClusterResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, dbServerDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, vmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VmClusterResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, dbServerDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Create, vmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "84"),
				resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "120"),
				resource.TestCheckResourceAttr(resourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "60"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VmClusterResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, dbServerDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(vmClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "84"),
				resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "120"),
				resource.TestCheckResourceAttr(resourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "60"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + VmClusterResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, dbServerDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Update, vmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "6"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "86"),
				resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "160"),
				resource.TestCheckResourceAttr(resourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "90"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_clusters", "test_vm_clusters", acctest.Optional, acctest.Update, vmClusterDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, dbServerDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Update, vmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.cpus_enabled"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.data_storage_size_in_tbs", "86"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.db_node_storage_size_in_gbs", "160"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.db_servers.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.display_name", "vmCluster"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.memory_size_in_gbs", "90"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.shape"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.time_zone", "US/Pacific"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.vm_cluster_network_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, vmClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, dbServerDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Update, vmClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpus_enabled"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_storage_size_in_tbs", "86"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_node_storage_size_in_gbs", "160"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_servers.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "vmCluster"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gi_version", "19.0.0.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_local_backup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_sparse_diskgroup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "memory_size_in_gbs", "90"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "US/Pacific"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, dbServerDataSourceRepresentation) +
				VmClusterResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cpu_core_count",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseVmClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_vm_cluster" {
			noResourceFound = false
			request := oci_database.GetVmClusterRequest{}

			tmp := rs.Primary.ID
			request.VmClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetVmCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.VmClusterLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseVmCluster") {
		resource.AddTestSweepers("DatabaseVmCluster", &resource.Sweeper{
			Name:         "DatabaseVmCluster",
			Dependencies: acctest.DependencyGraph["vmCluster"],
			F:            sweepDatabaseVmClusterResource,
		})
	}
}

func sweepDatabaseVmClusterResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	vmClusterIds, err := getVmClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, vmClusterId := range vmClusterIds {
		if ok := acctest.SweeperDefaultResourceId[vmClusterId]; !ok {
			deleteVmClusterRequest := oci_database.DeleteVmClusterRequest{}

			deleteVmClusterRequest.VmClusterId = &vmClusterId

			deleteVmClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteVmCluster(context.Background(), deleteVmClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting VmCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", vmClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vmClusterId, vmClusterSweepWaitCondition, time.Duration(3*time.Minute),
				vmClusterSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getVmClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VmClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listVmClustersRequest := oci_database.ListVmClustersRequest{}
	listVmClustersRequest.CompartmentId = &compartmentId
	listVmClustersRequest.LifecycleState = oci_database.VmClusterSummaryLifecycleStateAvailable
	listVmClustersResponse, err := databaseClient.ListVmClusters(context.Background(), listVmClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VmCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vmCluster := range listVmClustersResponse.Items {
		id := *vmCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VmClusterId", id)
	}
	return resourceIds, nil
}

func vmClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vmClusterResponse, ok := response.Response.(oci_database.GetVmClusterResponse); ok {
		return vmClusterResponse.LifecycleState != oci_database.VmClusterLifecycleStateTerminated
	}
	return false
}

func vmClusterSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetVmCluster(context.Background(), oci_database.GetVmClusterRequest{
		VmClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
