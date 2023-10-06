// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	ContainerengineClusterNamespaceRequiredOnlyResource = ContainerengineClusterNamespaceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace", "test_cluster_namespace", acctest.Required, acctest.Create, ContainerengineClusterNamespaceRepresentation)

	ContainerengineClusterNamespaceResourceConfig = ContainerengineClusterNamespaceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace", "test_cluster_namespace", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceRepresentation)

	ContainerengineClusterNamespaceSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_namespace_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster_namespace.test_cluster_namespace.id}`},
	}

	ContainerengineClusterNamespaceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster_namespace.test_cluster_namespace.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineClusterNamespaceDataSourceFilterRepresentation}}
	ContainerengineClusterNamespaceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_cluster_namespace.test_cluster_namespace.id}`}},
	}

	ContainerengineClusterNamespaceRepresentation = map[string]interface{}{
		"cluster_namespace_profile_version_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster_namespace_profile_version.test_cluster_namespace_profile_version.id}`},
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                                 acctest.Representation{RepType: acctest.Required, Create: `name`},
		"description":                          acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"freeform_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"namespace_annotations":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterNamespaceNamespaceAnnotationsRepresentation},
		"namespace_labels":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterNamespaceNamespaceLabelsRepresentation},
		"depends_on":                           acctest.Representation{RepType: acctest.Required, Create: []string{`oci_containerengine_cluster.test_cluster`}},
	}
	ContainerengineClusterNamespaceNamespaceAnnotationsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `allowed-annotation-1`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `allowed-annotation-value-1`},
	}
	ContainerengineClusterNamespaceNamespaceLabelsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `allowed-label`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `allowed-label-value`},
	}

	ContainerengineClusterNamespaceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_version", "test_cluster_namespace_profile_version", acctest.Required, acctest.Create, ContainerengineClusterNamespaceProfileVersionRepresentation) +
		ContainerengineClusterNamespaceProfileVersionResourceDependencies
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterNamespaceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterNamespaceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_containerengine_cluster_namespace.test_cluster_namespace"
	datasourceName := "data.oci_containerengine_cluster_namespaces.test_cluster_namespaces"
	singularDatasourceName := "data.oci_containerengine_cluster_namespace.test_cluster_namespace"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerengineClusterNamespaceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace", "test_cluster_namespace", acctest.Optional, acctest.Create, ContainerengineClusterNamespaceRepresentation), "containerengine", "clusterNamespace", t)

	acctest.ResourceTest(t, testAccCheckContainerengineClusterNamespaceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace", "test_cluster_namespace", acctest.Required, acctest.Create, ContainerengineClusterNamespaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_version_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace", "test_cluster_namespace", acctest.Optional, acctest.Create, ContainerengineClusterNamespaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_version_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "namespace_annotations", map[string]string{
					"key":   "allowed-annotation-1",
					"value": "allowed-annotation-value-1",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "namespace_labels", map[string]string{
					"key":   "allowed-label",
					"value": "allowed-label-value",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ContainerengineClusterNamespaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace", "test_cluster_namespace", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ContainerengineClusterNamespaceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_version_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "namespace_annotations", map[string]string{
					"key":   "allowed-annotation-1",
					"value": "allowed-annotation-value-1",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "namespace_labels", map[string]string{
					"key":   "allowed-label",
					"value": "allowed-label-value",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace", "test_cluster_namespace", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_version_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "namespace_annotations", map[string]string{
					"key":   "allowed-annotation-1",
					"value": "allowed-annotation-value-1",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "namespace_labels", map[string]string{
					"key":   "allowed-label",
					"value": "allowed-label-value",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_namespaces", "test_cluster_namespaces", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterNamespaceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace", "test_cluster_namespace", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "cluster_namespace_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_namespace_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_namespace", "test_cluster_namespace", acctest.Required, acctest.Create, ContainerengineClusterNamespaceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterNamespaceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_namespace_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "namespace_annotations", map[string]string{
					"key":   "allowed-annotation-1",
					"value": "allowed-annotation-value-1",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "namespace_labels", map[string]string{
					"key":   "allowed-label",
					"value": "allowed-label-value",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ContainerengineClusterNamespaceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckContainerengineClusterNamespaceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_cluster_namespace" {
			noResourceFound = false
			request := oci_containerengine.GetClusterNamespaceRequest{}

			tmp := rs.Primary.ID
			request.ClusterNamespaceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")

			response, err := client.GetClusterNamespace(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_containerengine.ClusterNamespaceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ContainerengineClusterNamespace") {
		resource.AddTestSweepers("ContainerengineClusterNamespace", &resource.Sweeper{
			Name:         "ContainerengineClusterNamespace",
			Dependencies: acctest.DependencyGraph["clusterNamespace"],
			F:            sweepContainerengineClusterNamespaceResource,
		})
	}
}

func sweepContainerengineClusterNamespaceResource(compartment string) error {
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()
	clusterNamespaceIds, err := getContainerengineClusterNamespaceIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterNamespaceId := range clusterNamespaceIds {
		if ok := acctest.SweeperDefaultResourceId[clusterNamespaceId]; !ok {
			deleteClusterNamespaceRequest := oci_containerengine.DeleteClusterNamespaceRequest{}

			deleteClusterNamespaceRequest.ClusterNamespaceId = &clusterNamespaceId

			deleteClusterNamespaceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteClusterNamespace(context.Background(), deleteClusterNamespaceRequest)
			if error != nil {
				fmt.Printf("Error deleting ClusterNamespace %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterNamespaceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &clusterNamespaceId, ContainerengineClusterNamespaceSweepWaitCondition, time.Duration(3*time.Minute),
				ContainerengineClusterNamespaceSweepResponseFetchOperation, "containerengine", true)
		}
	}
	return nil
}

func getContainerengineClusterNamespaceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ClusterNamespaceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listClusterNamespacesRequest := oci_containerengine.ListClusterNamespacesRequest{}
	listClusterNamespacesRequest.CompartmentId = &compartmentId
	listClusterNamespacesRequest.LifecycleState = oci_containerengine.ClusterNamespaceLifecycleStateActive
	listClusterNamespacesResponse, err := containerEngineClient.ListClusterNamespaces(context.Background(), listClusterNamespacesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ClusterNamespace list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, clusterNamespace := range listClusterNamespacesResponse.Items {
		id := *clusterNamespace.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ClusterNamespaceId", id)
	}
	return resourceIds, nil
}

func ContainerengineClusterNamespaceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if clusterNamespaceResponse, ok := response.Response.(oci_containerengine.GetClusterNamespaceResponse); ok {
		return clusterNamespaceResponse.LifecycleState != oci_containerengine.ClusterNamespaceLifecycleStateDeleted
	}
	return false
}

func ContainerengineClusterNamespaceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerEngineClient().GetClusterNamespace(context.Background(), oci_containerengine.GetClusterNamespaceRequest{
		ClusterNamespaceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
