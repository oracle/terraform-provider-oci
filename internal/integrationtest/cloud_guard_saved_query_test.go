// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CloudGuardSavedQueryRequiredOnlyResource = CloudGuardSavedQueryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_saved_query", "test_saved_query", acctest.Required, acctest.Create, CloudGuardSavedQueryRepresentation)

	CloudGuardSavedQueryResourceConfig = CloudGuardSavedQueryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_saved_query", "test_saved_query", acctest.Optional, acctest.Update, CloudGuardSavedQueryRepresentation)

	CloudGuardSavedQuerySingularDataSourceRepresentation = map[string]interface{}{
		"saved_query_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_saved_query.test_saved_query.id}`},
	}

	CloudGuardSavedQueryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudGuardSavedQueryDataSourceFilterRepresentation}}
	CloudGuardSavedQueryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_guard_saved_query.test_saved_query.id}`}},
	}

	CloudGuardSavedQueryRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayNameUp`},
		"query":          acctest.Representation{RepType: acctest.Required, Create: `select pid from processes`, Update: `select pid from processes`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"bar-key": "value"}},
	}

	CloudGuardSavedQueryResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardSavedQueryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardSavedQueryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_guard_saved_query.test_saved_query"
	datasourceName := "data.oci_cloud_guard_saved_queries.test_saved_queries"
	singularDatasourceName := "data.oci_cloud_guard_saved_query.test_saved_query"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudGuardSavedQueryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_saved_query", "test_saved_query", acctest.Optional, acctest.Create, CloudGuardSavedQueryRepresentation), "cloudguard", "savedQuery", t)

	acctest.ResourceTest(t, testAccCheckCloudGuardSavedQueryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardSavedQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_saved_query", "test_saved_query", acctest.Required, acctest.Create, CloudGuardSavedQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "query", "select pid from processes"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardSavedQueryResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudGuardSavedQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_saved_query", "test_saved_query", acctest.Optional, acctest.Create, CloudGuardSavedQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "query", "select pid from processes"),

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
			Config: config + compartmentIdUVariableStr + CloudGuardSavedQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_saved_query", "test_saved_query", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudGuardSavedQueryRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "query", "select pid from processes"),

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
			Config: config + compartmentIdVariableStr + CloudGuardSavedQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_saved_query", "test_saved_query", acctest.Optional, acctest.Update, CloudGuardSavedQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameUp"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "query", "select pid from processes"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_saved_queries", "test_saved_queries", acctest.Optional, acctest.Update, CloudGuardSavedQueryDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardSavedQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_saved_query", "test_saved_query", acctest.Optional, acctest.Update, CloudGuardSavedQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "saved_query_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "saved_query_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_saved_query", "test_saved_query", acctest.Required, acctest.Create, CloudGuardSavedQuerySingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardSavedQueryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "saved_query_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query", "select pid from processes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudGuardSavedQueryRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudGuardSavedQueryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_saved_query" {
			noResourceFound = false
			request := oci_cloud_guard.GetSavedQueryRequest{}

			tmp := rs.Primary.ID
			request.SavedQueryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")

			response, err := client.GetSavedQuery(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_guard.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudGuardSavedQuery") {
		resource.AddTestSweepers("CloudGuardSavedQuery", &resource.Sweeper{
			Name:         "CloudGuardSavedQuery",
			Dependencies: acctest.DependencyGraph["savedQuery"],
			F:            sweepCloudGuardSavedQueryResource,
		})
	}
}

func sweepCloudGuardSavedQueryResource(compartment string) error {
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()
	savedQueryIds, err := getCloudGuardSavedQueryIds(compartment)
	if err != nil {
		return err
	}
	for _, savedQueryId := range savedQueryIds {
		if ok := acctest.SweeperDefaultResourceId[savedQueryId]; !ok {
			deleteSavedQueryRequest := oci_cloud_guard.DeleteSavedQueryRequest{}

			deleteSavedQueryRequest.SavedQueryId = &savedQueryId

			deleteSavedQueryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteSavedQuery(context.Background(), deleteSavedQueryRequest)
			if error != nil {
				fmt.Printf("Error deleting SavedQuery %s %s, It is possible that the resource is already deleted. Please verify manually \n", savedQueryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &savedQueryId, CloudGuardSavedQuerySweepWaitCondition, time.Duration(3*time.Minute),
				CloudGuardSavedQuerySweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getCloudGuardSavedQueryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SavedQueryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()

	listSavedQueriesRequest := oci_cloud_guard.ListSavedQueriesRequest{}
	listSavedQueriesRequest.CompartmentId = &compartmentId
	listSavedQueriesResponse, err := cloudGuardClient.ListSavedQueries(context.Background(), listSavedQueriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SavedQuery list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, savedQuery := range listSavedQueriesResponse.Items {
		id := *savedQuery.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SavedQueryId", id)
	}
	return resourceIds, nil
}

func CloudGuardSavedQuerySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if savedQueryResponse, ok := response.Response.(oci_cloud_guard.GetSavedQueryResponse); ok {
		return savedQueryResponse.LifecycleState != oci_cloud_guard.LifecycleStateDeleted
	}
	return false
}

func CloudGuardSavedQuerySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CloudGuardClient().GetSavedQuery(context.Background(), oci_cloud_guard.GetSavedQueryRequest{
		SavedQueryId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
