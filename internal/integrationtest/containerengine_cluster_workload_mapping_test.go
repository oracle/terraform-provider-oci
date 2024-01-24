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
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineClusterWorkloadMappingRequiredOnlyResource = ContainerengineClusterWorkloadMappingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_workload_mapping", "test_cluster_workload_mapping", acctest.Required, acctest.Create, ContainerengineClusterWorkloadMappingRepresentation)

	ContainerengineClusterWorkloadMappingResourceConfig = ContainerengineClusterWorkloadMappingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_workload_mapping", "test_cluster_workload_mapping", acctest.Optional, acctest.Update, ContainerengineClusterWorkloadMappingRepresentation)

	ContainerengineContainerengineClusterWorkloadMappingSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"workload_mapping_id": acctest.Representation{RepType: acctest.Required, Create: `${split("/", oci_containerengine_cluster_workload_mapping.test_cluster_workload_mapping.id)[3]}`},
	}

	ContainerengineContainerengineClusterWorkloadMappingDataSourceRepresentation = map[string]interface{}{
		"cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"filter":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineClusterWorkloadMappingDataSourceFilterRepresentation},
	}
	ContainerengineClusterWorkloadMappingDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${split("/", oci_containerengine_cluster_workload_mapping.test_cluster_workload_mapping.id)[3]}`}},
	}

	ContainerengineClusterWorkloadMappingRepresentation = map[string]interface{}{
		"cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"mapped_compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":             acctest.Representation{RepType: acctest.Required, Create: `namespace`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineClusterWorkloadMappingIgnoreDefinedTagsRepresentation},
	}

	//ignore Defined tag change because we have default tag in the test tenancy
	ContainerengineClusterWorkloadMappingIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	ContainerengineClusterWorkloadMappingResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create, ContainerengineClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster1`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "clusterSubnet_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.0.21.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `cluster2`}})) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		DefinedTagsDependencies
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterWorkloadMappingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterWorkloadMappingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_cluster_workload_mapping.test_cluster_workload_mapping"
	datasourceName := "data.oci_containerengine_cluster_workload_mappings.test_cluster_workload_mappings"
	singularDatasourceName := "data.oci_containerengine_cluster_workload_mapping.test_cluster_workload_mapping"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerengineClusterWorkloadMappingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_workload_mapping", "test_cluster_workload_mapping", acctest.Optional, acctest.Create, ContainerengineClusterWorkloadMappingRepresentation), "containerengine", "clusterWorkloadMapping", t)

	acctest.ResourceTest(t, testAccCheckContainerengineClusterWorkloadMappingDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterWorkloadMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_workload_mapping", "test_cluster_workload_mapping", acctest.Required, acctest.Create, ContainerengineClusterWorkloadMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "mapped_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterWorkloadMappingResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterWorkloadMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_workload_mapping", "test_cluster_workload_mapping", acctest.Optional, acctest.Create, ContainerengineClusterWorkloadMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "mapped_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "mapped_tenancy_id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
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
			Config: config + compartmentIdVariableStr + ContainerengineClusterWorkloadMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_workload_mapping", "test_cluster_workload_mapping", acctest.Optional, acctest.Update, ContainerengineClusterWorkloadMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "mapped_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "mapped_tenancy_id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_workload_mappings", "test_cluster_workload_mappings", acctest.Optional, acctest.Update, ContainerengineContainerengineClusterWorkloadMappingDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterWorkloadMappingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_workload_mapping", "test_cluster_workload_mapping", acctest.Optional, acctest.Update, ContainerengineClusterWorkloadMappingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_id"),

				resource.TestCheckResourceAttr(datasourceName, "workload_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "workload_mappings.0.cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "workload_mappings.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "workload_mappings.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "workload_mappings.0.mapped_compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "workload_mappings.0.mapped_tenancy_id"),
				resource.TestCheckResourceAttr(datasourceName, "workload_mappings.0.namespace", "namespace"),
				resource.TestCheckResourceAttrSet(datasourceName, "workload_mappings.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "workload_mappings.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_workload_mapping", "test_cluster_workload_mapping", acctest.Required, acctest.Create, ContainerengineContainerengineClusterWorkloadMappingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterWorkloadMappingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workload_mapping_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mapped_tenancy_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + ContainerengineClusterWorkloadMappingRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckContainerengineClusterWorkloadMappingDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_cluster_workload_mapping" {
			noResourceFound = false
			request := oci_containerengine.GetWorkloadMappingRequest{}

			if value, ok := rs.Primary.Attributes["cluster_id"]; ok {
				request.ClusterId = &value
			}

			if value, ok := rs.Primary.Attributes["id"]; ok {
				request.WorkloadMappingId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")

			response, err := client.GetWorkloadMapping(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_containerengine.WorkloadMappingLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ContainerengineClusterWorkloadMapping") {
		resource.AddTestSweepers("ContainerengineClusterWorkloadMapping", &resource.Sweeper{
			Name:         "ContainerengineClusterWorkloadMapping",
			Dependencies: acctest.DependencyGraph["clusterWorkloadMapping"],
			F:            sweepContainerengineClusterWorkloadMappingResource,
		})
	}
}

func sweepContainerengineClusterWorkloadMappingResource(compartment string) error {
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()
	clusterWorkloadMappingIds, err := getContainerengineClusterWorkloadMappingIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterWorkloadMappingId := range clusterWorkloadMappingIds {
		if ok := acctest.SweeperDefaultResourceId[clusterWorkloadMappingId]; !ok {
			deleteWorkloadMappingRequest := oci_containerengine.DeleteWorkloadMappingRequest{}

			deleteWorkloadMappingRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteWorkloadMapping(context.Background(), deleteWorkloadMappingRequest)
			if error != nil {
				fmt.Printf("Error deleting ClusterWorkloadMapping %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterWorkloadMappingId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &clusterWorkloadMappingId, ContainerengineClusterWorkloadMappingSweepWaitCondition, time.Duration(3*time.Minute),
				ContainerengineClusterWorkloadMappingSweepResponseFetchOperation, "containerengine", true)
		}
	}
	return nil
}

func getContainerengineClusterWorkloadMappingIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ClusterWorkloadMappingId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listWorkloadMappingsRequest := oci_containerengine.ListWorkloadMappingsRequest{}

	clusterIds, error := getClusterIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting clusterId required for ClusterWorkloadMapping resource requests \n")
	}
	for _, clusterId := range clusterIds {
		listWorkloadMappingsRequest.ClusterId = &clusterId

		listWorkloadMappingsResponse, err := containerEngineClient.ListWorkloadMappings(context.Background(), listWorkloadMappingsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting ClusterWorkloadMapping list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, clusterWorkloadMapping := range listWorkloadMappingsResponse.Items {
			id := *clusterWorkloadMapping.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ClusterWorkloadMappingId", id)
		}

	}
	return resourceIds, nil
}

func ContainerengineClusterWorkloadMappingSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if clusterWorkloadMappingResponse, ok := response.Response.(oci_containerengine.GetWorkloadMappingResponse); ok {
		return clusterWorkloadMappingResponse.LifecycleState != oci_containerengine.WorkloadMappingLifecycleStateDeleted
	}
	return false
}

func ContainerengineClusterWorkloadMappingSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerEngineClient().GetWorkloadMapping(context.Background(), oci_containerengine.GetWorkloadMappingRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
