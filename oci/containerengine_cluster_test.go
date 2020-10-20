// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v27/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v27/containerengine"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ClusterRequiredOnlyResource = ClusterResourceDependencies +
		generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Required, Create, clusterRepresentation)

	ClusterResourceConfig = ClusterResourceDependencies +
		generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Optional, Create, clusterRepresentation)

	clusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"name":           Representation{repType: Optional, create: `name`, update: `name2`},
		"state":          Representation{repType: Optional, create: []string{`CREATING`, `ACTIVE`, `FAILED`, `DELETING`, `DELETED`, `UPDATING`}},
		"filter":         RepresentationGroup{Required, clusterDataSourceFilterRepresentation}}
	clusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_containerengine_cluster.test_cluster.id}`}},
	}

	clusterRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"kubernetes_version": Representation{repType: Required, create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-1]}`},
		"name":               Representation{repType: Required, create: `name`, update: `name2`},
		"vcn_id":             Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
		"kms_key_id":         Representation{repType: Optional, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"options":            RepresentationGroup{Optional, clusterOptionsRepresentation},
	}
	clusterOptionsRepresentation = map[string]interface{}{
		"add_ons":                      RepresentationGroup{Optional, clusterOptionsAddOnsRepresentation},
		"admission_controller_options": RepresentationGroup{Optional, clusterOptionsAdmissionControllerOptionsRepresentation},
		"kubernetes_network_config":    RepresentationGroup{Optional, clusterOptionsKubernetesNetworkConfigRepresentation},
		"service_lb_subnet_ids":        Representation{repType: Optional, create: []string{`${oci_core_subnet.clusterSubnet_1.id}`, `${oci_core_subnet.clusterSubnet_2.id}`}},
	}
	clusterOptionsAddOnsRepresentation = map[string]interface{}{
		"is_kubernetes_dashboard_enabled": Representation{repType: Optional, create: `true`},
		"is_tiller_enabled":               Representation{repType: Optional, create: `true`},
	}
	clusterOptionsAdmissionControllerOptionsRepresentation = map[string]interface{}{
		"is_pod_security_policy_enabled": Representation{repType: Optional, create: `false`, update: `true`},
	}
	clusterOptionsKubernetesNetworkConfigRepresentation = map[string]interface{}{
		"pods_cidr":     Representation{repType: Optional, create: `10.1.0.0/16`},
		"services_cidr": Representation{repType: Optional, create: `10.2.0.0/16`},
	}

	ClusterResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{repType: Required, create: `10.0.20.0/24`}, "dns_label": Representation{repType: Required, create: `cluster1`}})) +
		generateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{repType: Required, create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{repType: Required, create: `10.0.21.0/24`}, "dns_label": Representation{repType: Required, create: `cluster2`}})) +
		AvailabilityDomainConfig +
		generateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", Required, Create, clusterOptionSingularDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		KeyResourceDependencyConfig
)

func TestContainerengineClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_cluster.test_cluster"
	datasourceName := "data.oci_containerengine_clusters.test_clusters"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckContainerengineClusterDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Required, Create, clusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ClusterResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Optional, Create, clusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_kubernetes_dashboard_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_tiller_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.0.is_pod_security_policy_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.pods_cidr", "10.1.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.services_cidr", "10.2.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "options.0.service_lb_subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + ClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Optional, Update, clusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_kubernetes_dashboard_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.add_ons.0.is_tiller_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.admission_controller_options.0.is_pod_security_policy_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.pods_cidr", "10.1.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "options.0.kubernetes_network_config.0.services_cidr", "10.2.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "options.0.service_lb_subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
					generateDataSourceFromRepresentationMap("oci_containerengine_clusters", "test_clusters", Optional, Update, clusterDataSourceRepresentation) +
					compartmentIdVariableStr + ClusterResourceDependencies +
					generateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Optional, Update, clusterRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", "name2"),

					resource.TestCheckResourceAttr(datasourceName, "clusters.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.endpoints.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.kubernetes_version"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.metadata.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.name", "name2"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.add_ons.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.add_ons.0.is_kubernetes_dashboard_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.add_ons.0.is_tiller_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.admission_controller_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.admission_controller_options.0.is_pod_security_policy_enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.kubernetes_network_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.kubernetes_network_config.0.pods_cidr", "10.1.0.0/16"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.kubernetes_network_config.0.services_cidr", "10.2.0.0/16"),
					resource.TestCheckResourceAttr(datasourceName, "clusters.0.options.0.service_lb_subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.vcn_id"),
				),
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

func testAccCheckContainerengineClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).containerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_cluster" {
			noResourceFound = false
			request := oci_containerengine.GetClusterRequest{}

			tmp := rs.Primary.ID
			request.ClusterId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "containerengine")

			response, err := client.GetCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_containerengine.ClusterLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("ContainerengineCluster") {
		resource.AddTestSweepers("ContainerengineCluster", &resource.Sweeper{
			Name:         "ContainerengineCluster",
			Dependencies: DependencyGraph["cluster"],
			F:            sweepContainerengineClusterResource,
		})
	}
}

func sweepContainerengineClusterResource(compartment string) error {
	containerEngineClient := GetTestClients(&schema.ResourceData{}).containerEngineClient()
	clusterIds, err := getClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterId := range clusterIds {
		if ok := SweeperDefaultResourceId[clusterId]; !ok {
			deleteClusterRequest := oci_containerengine.DeleteClusterRequest{}

			deleteClusterRequest.ClusterId = &clusterId

			deleteClusterRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteCluster(context.Background(), deleteClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting Cluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterId, error)
				continue
			}
			waitTillCondition(testAccProvider, &clusterId, clusterSweepWaitCondition, time.Duration(3*time.Minute),
				clusterSweepResponseFetchOperation, "containerengine", true)
		}
	}
	return nil
}

func getClusterIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := GetTestClients(&schema.ResourceData{}).containerEngineClient()

	listClustersRequest := oci_containerengine.ListClustersRequest{}
	listClustersRequest.CompartmentId = &compartmentId
	listClustersResponse, err := containerEngineClient.ListClusters(context.Background(), listClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Cluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cluster := range listClustersResponse.Items {
		id := *cluster.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ClusterId", id)
	}
	return resourceIds, nil
}

func clusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if clusterResponse, ok := response.Response.(oci_containerengine.GetClusterResponse); ok {
		return clusterResponse.LifecycleState != oci_containerengine.ClusterLifecycleStateDeleted
	}
	return false
}

func clusterSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.containerEngineClient().GetCluster(context.Background(), oci_containerengine.GetClusterRequest{
		ClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
