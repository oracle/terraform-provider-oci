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
	VmClusterRequiredOnlyResource = VmClusterResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRepresentation)

	VmClusterResourceConfig = VmClusterResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Update, vmClusterRepresentation)

	vmClusterSingularDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": Representation{repType: Required, create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	vmClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":              Representation{repType: Optional, create: `vmCluster`},
		"exadata_infrastructure_id": Representation{repType: Optional, create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"state":                     Representation{repType: Optional, create: `AVAILABLE`},
		"filter":                    RepresentationGroup{Required, vmClusterDataSourceFilterRepresentation}}
	vmClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_vm_cluster.test_vm_cluster.id}`}},
	}

	vmClusterRepresentation = map[string]interface{}{
		"compartment_id":              Representation{repType: Required, create: `${var.compartment_id}`},
		"cpu_core_count":              Representation{repType: Required, create: `4`, update: `2`},
		"display_name":                Representation{repType: Required, create: `vmCluster`},
		"exadata_infrastructure_id":   Representation{repType: Required, create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"gi_version":                  Representation{repType: Required, create: `19.1.0.0`},
		"ssh_public_keys":             Representation{repType: Required, create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"vm_cluster_network_id":       Representation{repType: Required, create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"defined_tags":                Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":     Representation{repType: Optional, create: `false`},
		"is_sparse_diskgroup_enabled": Representation{repType: Optional, create: `false`},
		"license_model":               Representation{repType: Optional, create: `LICENSE_INCLUDED`},
		"time_zone":                   Representation{repType: Optional, create: `US/Pacific`},
	}

	VmClusterResourceDependencies = VmClusterNetworkValidatedResourceConfig
)

func TestDatabaseVmClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_vm_cluster.test_vm_cluster"
	datasourceName := "data.oci_database_vm_clusters.test_vm_clusters"
	singularDatasourceName := "data.oci_database_vm_cluster.test_vm_cluster"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseVmClusterDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VmClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "gi_version", "19.1.0.0"),
					resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VmClusterResourceDependencies,
			},
			//verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VmClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Create, vmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "gi_version", "19.1.0.0"),
					resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
					resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_network_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VmClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Create,
						representationCopyWithNewProperties(vmClusterRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "4"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "gi_version", "19.1.0.0"),
					resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
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
				Config: config + compartmentIdVariableStr + VmClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Update, vmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "vmCluster"),
					resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "gi_version", "19.1.0.0"),
					resource.TestCheckResourceAttr(resourceName, "is_local_backup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_sparse_diskgroup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
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
					generateDataSourceFromRepresentationMap("oci_database_vm_clusters", "test_vm_clusters", Optional, Update, vmClusterDataSourceRepresentation) +
					compartmentIdVariableStr + VmClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Optional, Update, vmClusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "vmCluster"),
					resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.cpus_enabled"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.data_storage_size_in_tbs"),
					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.display_name", "vmCluster"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.gi_version", "19.1.0.0"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.is_local_backup_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.is_sparse_diskgroup_enabled", "false"),
					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.shape"),
					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.ssh_public_keys.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "vm_clusters.0.time_zone", "US/Pacific"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_clusters.0.vm_cluster_network_id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VmClusterResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cpus_enabled"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_storage_size_in_tbs"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "vmCluster"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "gi_version", "19.1.0.0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_local_backup_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_sparse_diskgroup_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ssh_public_keys.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "US/Pacific"),
				),
			},
		},
	})
}

func testAccCheckDatabaseVmClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_vm_cluster" {
			noResourceFound = false
			request := oci_database.GetVmClusterRequest{}

			tmp := rs.Primary.ID
			request.VmClusterId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseVmCluster") {
		resource.AddTestSweepers("DatabaseVmCluster", &resource.Sweeper{
			Name:         "DatabaseVmCluster",
			Dependencies: DependencyGraph["vmCluster"],
			F:            sweepDatabaseVmClusterResource,
		})
	}
}

func sweepDatabaseVmClusterResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient
	vmClusterIds, err := getVmClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, vmClusterId := range vmClusterIds {
		if ok := SweeperDefaultResourceId[vmClusterId]; !ok {
			deleteVmClusterRequest := oci_database.DeleteVmClusterRequest{}

			deleteVmClusterRequest.VmClusterId = &vmClusterId

			deleteVmClusterRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteVmCluster(context.Background(), deleteVmClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting VmCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", vmClusterId, error)
				continue
			}
			waitTillCondition(testAccProvider, &vmClusterId, vmClusterSweepWaitCondition, time.Duration(3*time.Minute),
				vmClusterSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getVmClusterIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VmClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "VmClusterId", id)
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

func vmClusterSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient.GetVmCluster(context.Background(), oci_database.GetVmClusterRequest{
		VmClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
