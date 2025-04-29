// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
	OcvpDatastoreRequiredOnlyResource = OcvpDatastoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore", "test_datastore", acctest.Required, acctest.Create, OcvpDatastoreRepresentation)

	OcvpDatastoreResourceConfig = OcvpDatastoreResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore", "test_datastore", acctest.Optional, acctest.Update, OcvpDatastoreRepresentationForUpdate)

	OcvpDatastoreSingularDataSourceRepresentation = map[string]interface{}{
		"datastore_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_datastore.test_datastore.id}`},
	}

	OcvpDatastoreDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"datastore_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_datastore.test_datastore.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `Active`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpDatastoreDataSourceFilterRepresentation}}
	OcvpDatastoreDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ocvp_datastore.test_datastore.id}`}},
	}

	OcvpDatastoreRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"block_volume_ids":    acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_volume.test_volumes[0].id}`}},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
	}

	OcvpDatastoreRepresentationForUpdate = acctest.GetUpdatedRepresentationCopy("block_volume_ids", acctest.Representation{
		RepType: acctest.Required, Create: []string{`${oci_core_volume.test_volumes[0].id}`, `${oci_core_volume.test_volumes[1].id}`},
	}, OcvpDatastoreRepresentation)

	OcvpDatastoreResourceDependencies = AvailabilityDomainConfig + DefinedTagsDependencies + `
resource "oci_core_volume" "test_volumes" {
  count               = 2
  display_name		  = "test_volume_${count.index}"
  availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
  compartment_id      = var.compartment_id
  size_in_gbs         = 50
}
`
)

// issue-routing-tag: ocvp/default
func TestOcvpDatastoreResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpDatastoreResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ocvp_datastore.test_datastore"
	datasourceName := "data.oci_ocvp_datastores.test_datastores"
	singularDatasourceName := "data.oci_ocvp_datastore.test_datastore"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OcvpDatastoreResourceDependencies, "ocvp", "datastore", t)

	acctest.ResourceTest(t, testAccCheckOcvpDatastoreDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OcvpDatastoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore", "test_datastore", acctest.Required, acctest.Create, OcvpDatastoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "block_volume_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_ids.0"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "block_volume_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OcvpDatastoreResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OcvpDatastoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore", "test_datastore", acctest.Optional, acctest.Create, OcvpDatastoreRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "block_volume_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_ids.0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "block_volume_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_details.0.id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OcvpDatastoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore", "test_datastore", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OcvpDatastoreRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "block_volume_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_ids.0"),
				resource.TestCheckResourceAttr(resourceName, "block_volume_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
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
			Config: config + compartmentIdVariableStr + OcvpDatastoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore", "test_datastore", acctest.Optional, acctest.Update, OcvpDatastoreRepresentationForUpdate),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "block_volume_ids.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_ids.1"),
				resource.TestCheckResourceAttr(resourceName, "block_volume_details.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_details.1.id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_datastores", "test_datastores", acctest.Optional, acctest.Update, OcvpDatastoreDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpDatastoreResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_datastore", "test_datastore", acctest.Optional, acctest.Update, OcvpDatastoreRepresentationForUpdate),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "datastore_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "Active"),

				resource.TestCheckResourceAttr(datasourceName, "datastore_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "datastore_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_datastore", "test_datastore", acctest.Required, acctest.Create, OcvpDatastoreSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpDatastoreResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "datastore_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "block_volume_ids.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "block_volume_details.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "block_volume_details.1.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "capacity_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OcvpDatastoreRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOcvpDatastoreDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatastoreClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_datastore" {
			noResourceFound = false
			request := oci_ocvp.GetDatastoreRequest{}

			tmp := rs.Primary.ID
			request.DatastoreId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")

			response, err := client.GetDatastore(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("OcvpDatastore") {
		resource.AddTestSweepers("OcvpDatastore", &resource.Sweeper{
			Name:         "OcvpDatastore",
			Dependencies: acctest.DependencyGraph["datastore"],
			F:            sweepOcvpDatastoreResource,
		})
	}
}

func sweepOcvpDatastoreResource(compartment string) error {
	datastoreClient := acctest.GetTestClients(&schema.ResourceData{}).DatastoreClient()
	datastoreIds, err := getOcvpDatastoreIds(compartment)
	if err != nil {
		return err
	}
	for _, datastoreId := range datastoreIds {
		if ok := acctest.SweeperDefaultResourceId[datastoreId]; !ok {
			deleteDatastoreRequest := oci_ocvp.DeleteDatastoreRequest{}

			deleteDatastoreRequest.DatastoreId = &datastoreId

			deleteDatastoreRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")
			_, error := datastoreClient.DeleteDatastore(context.Background(), deleteDatastoreRequest)
			if error != nil {
				fmt.Printf("Error deleting Datastore %s %s, It is possible that the resource is already deleted. Please verify manually \n", datastoreId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &datastoreId, OcvpDatastoreSweepWaitCondition, time.Duration(3*time.Minute),
				OcvpDatastoreSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getOcvpDatastoreIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatastoreId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	datastoreClient := acctest.GetTestClients(&schema.ResourceData{}).DatastoreClient()

	listDatastoresRequest := oci_ocvp.ListDatastoresRequest{}
	listDatastoresRequest.CompartmentId = &compartmentId
	listDatastoresRequest.LifecycleState = oci_ocvp.ListDatastoresLifecycleStateActive
	listDatastoresResponse, err := datastoreClient.ListDatastores(context.Background(), listDatastoresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Datastore list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, datastore := range listDatastoresResponse.Items {
		id := *datastore.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatastoreId", id)
	}
	return resourceIds, nil
}

func OcvpDatastoreSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if datastoreResponse, ok := response.Response.(oci_ocvp.GetDatastoreResponse); ok {
		return datastoreResponse.LifecycleState != oci_ocvp.LifecycleStatesDeleted
	}
	return false
}

func OcvpDatastoreSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatastoreClient().GetDatastore(context.Background(), oci_ocvp.GetDatastoreRequest{
		DatastoreId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
