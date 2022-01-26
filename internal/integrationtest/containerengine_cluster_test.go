// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/containerengine"
	oci_containerengine "github.com/oracle/oci-go-sdk/v56/containerengine"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ClusterRequiredOnlyResource = ClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, clusterRepresentation)

	clusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: []string{`CREATING`, `ACTIVE`, `FAILED`, `DELETING`, `DELETED`, `UPDATING`}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: clusterDataSourceFilterRepresentation}}
	clusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_cluster.test_cluster.id}`}},
	}

	clusterRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kubernetes_version":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-2]}`, Update: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-1]}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"vcn_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"endpoint_config":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: clusterEndpointConfigRepresentation},
		"kms_key_id":          acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"options":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: clusterOptionsRepresentation},
		"image_policy_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: clusterImagePolicyConfigRepresentation},
	}
	clusterEndpointConfigRepresentation = map[string]interface{}{
		"is_public_ip_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"nsg_ids":              acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"subnet_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	clusterImagePolicyConfigRepresentation = map[string]interface{}{
		"is_policy_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key_details":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: clusterImagePolicyConfigKeyDetailsRepresentation},
	}
	clusterOptionsRepresentation = map[string]interface{}{
		"add_ons":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: clusterOptionsAddOnsRepresentation},
		"admission_controller_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: clusterOptionsAdmissionControllerOptionsRepresentation},
		"kubernetes_network_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: clusterOptionsKubernetesNetworkConfigRepresentation},
		"service_lb_subnet_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_subnet.clusterSubnet_1.id}`, `${oci_core_subnet.clusterSubnet_2.id}`}},
	}
	clusterImagePolicyConfigKeyDetailsRepresentation = map[string]interface{}{
		"kms_key_id": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
	}
	clusterOptionsAddOnsRepresentation = map[string]interface{}{
		"is_kubernetes_dashboard_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_tiller_enabled":               acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	clusterOptionsAdmissionControllerOptionsRepresentation = map[string]interface{}{
		"is_pod_security_policy_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	clusterOptionsKubernetesNetworkConfigRepresentation = map[string]interface{}{
		"pods_cidr":     acctest.Representation{RepType: acctest.Optional, Create: `10.1.0.0/16`},
		"services_cidr": acctest.Representation{RepType: acctest.Optional, Create: `10.2.0.0/16`},
	}

	ClusterResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.21.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, clusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_cluster.test_cluster"
	datasourceName := "data.oci_containerengine_clusters.test_clusters"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, clusterRepresentation), "containerengine", "cluster", t)

	acctest.ResourceTest(t, testAccCheckContainerengineClusterDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + ClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, clusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ClusterResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Create, clusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.0.is_public_ip_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.is_policy_enabled", "false"),
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
			Config: config + compartmentIdVariableStr + ClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Update, clusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.is_policy_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "image_policy_config.0.key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "image_policy_config.0.key_details.0.kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.0.is_public_ip_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_clusters", "test_clusters", acctest.Optional, acctest.Update, clusterDataSourceRepresentation) +
				compartmentIdVariableStr + ClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Optional, acctest.Update, clusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "6"),

				resource.TestCheckResourceAttr(datasourceName, "clusters.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.available_kubernetes_upgrades.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.endpoint_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.endpoint_config.0.is_public_ip_enabled", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.endpoint_config.0.subnet_id"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.endpoints.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.image_policy_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.image_policy_config.0.is_policy_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "clusters.0.image_policy_config.0.key_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "clusters.0.image_policy_config.0.key_details.0.kms_key_id"),
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
	})
}

func testAccCheckContainerengineClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_cluster" {
			noResourceFound = false
			request := oci_containerengine.GetClusterRequest{}

			tmp := rs.Primary.ID
			request.ClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ContainerengineCluster") {
		resource.AddTestSweepers("ContainerengineCluster", &resource.Sweeper{
			Name:         "ContainerengineCluster",
			Dependencies: acctest.DependencyGraph["cluster"],
			F:            sweepContainerengineClusterResource,
		})
	}
}

func sweepContainerengineClusterResource(compartment string) error {
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()
	clusterIds, err := getClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterId := range clusterIds {
		if ok := acctest.SweeperDefaultResourceId[clusterId]; !ok {
			deleteClusterRequest := oci_containerengine.DeleteClusterRequest{}

			deleteClusterRequest.ClusterId = &clusterId

			deleteClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteCluster(context.Background(), deleteClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting Cluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &clusterId, clusterSweepWaitCondition, time.Duration(3*time.Minute),
				clusterSweepResponseFetchOperation, "containerengine", true)
		}
	}
	return nil
}

func getClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listClustersRequest := oci_containerengine.ListClustersRequest{}
	listClustersRequest.CompartmentId = &compartmentId
	listClustersRequest.LifecycleState = []containerengine.ClusterLifecycleStateEnum{oci_containerengine.ClusterLifecycleStateActive}
	listClustersResponse, err := containerEngineClient.ListClusters(context.Background(), listClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Cluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cluster := range listClustersResponse.Items {
		id := *cluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ClusterId", id)
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

func clusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerEngineClient().GetCluster(context.Background(), oci_containerengine.GetClusterRequest{
		ClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
