// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OcvpDatastoreClusterRequiredOnlyResource = OcvpDatastoreClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore_cluster", "test_datastore_cluster", acctest.Required, acctest.Create, OcvpDatastoreClusterRepresentation)

	OcvpDatastoreClusterResourceConfig = OcvpDatastoreClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore_cluster", "test_datastore_cluster", acctest.Optional, acctest.Update, OcvpDatastoreClusterRepresentation)

	OcvpDatastoreClusterSingularDataSourceRepresentation = map[string]interface{}{
		"datastore_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_datastore_cluster.test_datastore_cluster.id}`},
	}

	OcvpDatastoreClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"datastore_cluster_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_datastore_cluster.test_datastore_cluster.id}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `Active`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpDatastoreClusterDataSourceFilterRepresentation}}
	OcvpDatastoreClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ocvp_datastore_cluster.test_datastore_cluster.id}`}},
	}

	OcvpDatastoreClusterRepresentation = map[string]interface{}{
		"availability_domain":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"datastore_cluster_type": acctest.Representation{RepType: acctest.Required, Create: `MANAGEMENT`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"datastore_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_ocvp_datastore.test_datastore_1.id}`}, Update: []string{`${oci_ocvp_datastore.test_datastore_1.id}`, `${oci_ocvp_datastore.test_datastore_2.id}`}},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
	}

	OcvpDatastoreClusterResourceDependencies = OcvpDatastoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore", "test_datastore_1", acctest.Required, acctest.Create,
			OcvpDatastoreRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore", "test_datastore_2", acctest.Required, acctest.Create,
			acctest.GetUpdatedRepresentationCopy("block_volume_ids", acctest.Representation{
				RepType: acctest.Required, Create: []string{`${oci_core_volume.test_volumes[0].id}`, `${oci_core_volume.test_volumes[1].id}`},
			}, OcvpDatastoreRepresentation))
	/*
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore_cluster", "test_datastore_cluster", acctest.Required, acctest.Create, OcvpDatastoreClusterRepresentation)
	*/
)

// issue-routing-tag: ocvp/default
func TestOcvpDatastoreClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpDatastoreClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ocvp_datastore_cluster.test_datastore_cluster"
	datasourceName := "data.oci_ocvp_datastore_clusters.test_datastore_clusters"
	singularDatasourceName := "data.oci_ocvp_datastore_cluster.test_datastore_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OcvpDatastoreClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore_cluster", "test_datastore_cluster", acctest.Optional, acctest.Create, OcvpDatastoreClusterRepresentation), "ocvp", "datastoreCluster", t)

	acctest.ResourceTest(t, testAccCheckOcvpDatastoreClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OcvpDatastoreClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore_cluster", "test_datastore_cluster", acctest.Required, acctest.Create, OcvpDatastoreClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "datastore_cluster_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OcvpDatastoreClusterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OcvpDatastoreClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore_cluster", "test_datastore_cluster", acctest.Optional, acctest.Create, OcvpDatastoreClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "datastore_cluster_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "datastore_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OcvpDatastoreClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore_cluster", "test_datastore_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OcvpDatastoreClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "datastore_cluster_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "datastore_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + OcvpDatastoreClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore_cluster", "test_datastore_cluster", acctest.Optional, acctest.Update, OcvpDatastoreClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "datastore_cluster_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(resourceName, "datastore_ids.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_datastore_clusters", "test_datastore_clusters", acctest.Optional, acctest.Update, OcvpDatastoreClusterDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpDatastoreClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore_cluster", "test_datastore_cluster", acctest.Optional, acctest.Update, OcvpDatastoreClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "datastore_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "Active"),

				resource.TestCheckResourceAttr(datasourceName, "datastore_cluster_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "datastore_cluster_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_datastore_cluster", "test_datastore_cluster", acctest.Required, acctest.Create, OcvpDatastoreClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpDatastoreClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "datastore_cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "datastore_cluster_type", "MANAGEMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datastore_ids.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "esxi_host_ids.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OcvpDatastoreClusterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOcvpDatastoreClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatastoreClusterClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_datastore_cluster" {
			noResourceFound = false
			request := oci_ocvp.GetDatastoreClusterRequest{}

			tmp := rs.Primary.ID
			request.DatastoreClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")

			response, err := client.GetDatastoreCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ocvp.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("OcvpDatastoreCluster") {
		resource.AddTestSweepers("OcvpDatastoreCluster", &resource.Sweeper{
			Name:         "OcvpDatastoreCluster",
			Dependencies: acctest.DependencyGraph["datastoreCluster"],
			F:            sweepOcvpDatastoreClusterResource,
		})
	}
}

func sweepOcvpDatastoreClusterResource(compartment string) error {
	datastoreClusterClient := acctest.GetTestClients(&schema.ResourceData{}).DatastoreClusterClient()
	datastoreClusterIds, err := getOcvpDatastoreClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, datastoreClusterId := range datastoreClusterIds {
		if ok := acctest.SweeperDefaultResourceId[datastoreClusterId]; !ok {
			deleteDatastoreClusterRequest := oci_ocvp.DeleteDatastoreClusterRequest{}

			deleteDatastoreClusterRequest.DatastoreClusterId = &datastoreClusterId

			deleteDatastoreClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")
			_, error := datastoreClusterClient.DeleteDatastoreCluster(context.Background(), deleteDatastoreClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting DatastoreCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", datastoreClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &datastoreClusterId, OcvpDatastoreClusterSweepWaitCondition, time.Duration(3*time.Minute),
				OcvpDatastoreClusterSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getOcvpDatastoreClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatastoreClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	datastoreClusterClient := acctest.GetTestClients(&schema.ResourceData{}).DatastoreClusterClient()

	listDatastoreClustersRequest := oci_ocvp.ListDatastoreClustersRequest{}
	listDatastoreClustersRequest.CompartmentId = &compartmentId
	listDatastoreClustersRequest.LifecycleState = oci_ocvp.ListDatastoreClustersLifecycleStateActive
	listDatastoreClustersResponse, err := datastoreClusterClient.ListDatastoreClusters(context.Background(), listDatastoreClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatastoreCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, datastoreCluster := range listDatastoreClustersResponse.Items {
		id := *datastoreCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatastoreClusterId", id)
	}
	return resourceIds, nil
}

func OcvpDatastoreClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if datastoreClusterResponse, ok := response.Response.(oci_ocvp.GetDatastoreClusterResponse); ok {
		return datastoreClusterResponse.LifecycleState != oci_ocvp.LifecycleStatesDeleted
	}
	return false
}

func OcvpDatastoreClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatastoreClusterClient().GetDatastoreCluster(context.Background(), oci_ocvp.GetDatastoreClusterRequest{
		DatastoreClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
