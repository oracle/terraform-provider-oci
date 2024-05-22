// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
	CloudGuardAdhocQueryRequiredOnlyResource = CloudGuardAdhocQueryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_adhoc_query", "test_adhoc_query", acctest.Required, acctest.Create, CloudGuardAdhocQueryRepresentation)

	CloudGuardAdhocQueryResourceConfig = CloudGuardAdhocQueryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_adhoc_query", "test_adhoc_query", acctest.Optional, acctest.Update, CloudGuardAdhocQueryRepresentation)

	CloudGuardAdhocQuerySingularDataSourceRepresentation = map[string]interface{}{
		"adhoc_query_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_adhoc_query.test_adhoc_query.id}`},
	}

	CloudGuardAdhocQueryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"access_level":                    acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree":       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"time_ended_filter_query_param":   acctest.Representation{RepType: acctest.Optional, Create: time.Now().Add((-2) * time.Hour)},
		"time_started_filter_query_param": acctest.Representation{RepType: acctest.Optional, Create: time.Now()},
		"filter":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudGuardAdhocQueryDataSourceFilterRepresentation}}
	CloudGuardAdhocQueryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_guard_adhoc_query.test_adhoc_query.id}`}},
	}

	CloudGuardAdhocQueryRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"adhoc_query_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudGuardAdhocQueryAdhocQueryDetailsRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
	CloudGuardAdhocQueryAdhocQueryDetailsRepresentation = map[string]interface{}{
		"query":                 acctest.Representation{RepType: acctest.Required, Create: `select pid from processes`},
		"adhoc_query_resources": acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudGuardAdhocQueryAdhocQueryDetailsAdhocQueryResourcesRepresentation},
	}
	CloudGuardAdhocQueryAdhocQueryDetailsAdhocQueryResourcesRepresentation = map[string]interface{}{
		"region":        acctest.Representation{RepType: acctest.Optional, Create: `us-phoenix-1`},
		"resource_ids":  acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.tenancy_ocid}`}},
		"resource_type": acctest.Representation{RepType: acctest.Optional, Create: `TENANCY`},
	}

	CloudGuardAdhocQueryResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardAdhocQueryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardAdhocQueryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_guard_adhoc_query.test_adhoc_query"
	datasourceName := "data.oci_cloud_guard_adhoc_queries.test_adhoc_queries"
	singularDatasourceName := "data.oci_cloud_guard_adhoc_query.test_adhoc_query"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudGuardAdhocQueryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_adhoc_query", "test_adhoc_query", acctest.Optional, acctest.Create, CloudGuardAdhocQueryRepresentation), "cloudguard", "adhocQuery", t)

	acctest.ResourceTest(t, testAccCheckCloudGuardAdhocQueryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardAdhocQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_adhoc_query", "test_adhoc_query", acctest.Required, acctest.Create, CloudGuardAdhocQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "adhoc_query_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "adhoc_query_details.0.adhoc_query_resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "adhoc_query_details.0.query", "select pid from processes"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardAdhocQueryResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudGuardAdhocQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_adhoc_query", "test_adhoc_query", acctest.Optional, acctest.Create, CloudGuardAdhocQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "adhoc_query_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "adhoc_query_details.0.adhoc_query_resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "adhoc_query_details.0.adhoc_query_resources.0.region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "adhoc_query_details.0.adhoc_query_resources.0.resource_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "adhoc_query_details.0.adhoc_query_resources.0.resource_type", "TENANCY"),
				resource.TestCheckResourceAttr(resourceName, "adhoc_query_details.0.query", "select pid from processes"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_adhoc_queries", "test_adhoc_queries", acctest.Optional, acctest.Update, CloudGuardAdhocQueryDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardAdhocQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_adhoc_query", "test_adhoc_query", acctest.Optional, acctest.Update, CloudGuardAdhocQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "adhoc_query_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "adhoc_query_collection.0.items.#", "1"),
				/*resource.TestCheckResourceAttrSet(datasourceName, "time_ended_filter_query_param"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_started_filter_query_param"),
				resource.TestCheckResourceAttr(datasourceName, "adhoc_query_status", "CREATING"),*/
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_adhoc_query", "test_adhoc_query", acctest.Required, acctest.Create, CloudGuardAdhocQuerySingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardAdhocQueryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "adhoc_query_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "adhoc_query_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "adhoc_query_details.0.adhoc_query_resources.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "adhoc_query_details.0.adhoc_query_resources.0.region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "adhoc_query_details.0.adhoc_query_resources.0.resource_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "adhoc_query_details.0.adhoc_query_resources.0.resource_type", "TENANCY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "adhoc_query_details.0.query", "select pid from processes"),
				resource.TestCheckResourceAttr(singularDatasourceName, "adhoc_query_regional_details.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudGuardAdhocQueryRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudGuardAdhocQueryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_adhoc_query" {
			noResourceFound = false
			request := oci_cloud_guard.GetAdhocQueryRequest{}

			tmp := rs.Primary.ID
			request.AdhocQueryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")

			response, err := client.GetAdhocQuery(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("CloudGuardAdhocQuery") {
		resource.AddTestSweepers("CloudGuardAdhocQuery", &resource.Sweeper{
			Name:         "CloudGuardAdhocQuery",
			Dependencies: acctest.DependencyGraph["adhocQuery"],
			F:            sweepCloudGuardAdhocQueryResource,
		})
	}
}

func sweepCloudGuardAdhocQueryResource(compartment string) error {
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()
	adhocQueryIds, err := getCloudGuardAdhocQueryIds(compartment)
	if err != nil {
		return err
	}
	for _, adhocQueryId := range adhocQueryIds {
		if ok := acctest.SweeperDefaultResourceId[adhocQueryId]; !ok {
			deleteAdhocQueryRequest := oci_cloud_guard.DeleteAdhocQueryRequest{}

			deleteAdhocQueryRequest.AdhocQueryId = &adhocQueryId

			deleteAdhocQueryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteAdhocQuery(context.Background(), deleteAdhocQueryRequest)
			if error != nil {
				fmt.Printf("Error deleting AdhocQuery %s %s, It is possible that the resource is already deleted. Please verify manually \n", adhocQueryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &adhocQueryId, CloudGuardAdhocQuerySweepWaitCondition, time.Duration(3*time.Minute),
				CloudGuardAdhocQuerySweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getCloudGuardAdhocQueryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AdhocQueryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()

	listAdhocQueriesRequest := oci_cloud_guard.ListAdhocQueriesRequest{}
	listAdhocQueriesRequest.CompartmentId = &compartmentId
	listAdhocQueriesResponse, err := cloudGuardClient.ListAdhocQueries(context.Background(), listAdhocQueriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AdhocQuery list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, adhocQuery := range listAdhocQueriesResponse.Items {
		id := *adhocQuery.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AdhocQueryId", id)
	}
	return resourceIds, nil
}

func CloudGuardAdhocQuerySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if adhocQueryResponse, ok := response.Response.(oci_cloud_guard.GetAdhocQueryResponse); ok {
		return adhocQueryResponse.LifecycleState != oci_cloud_guard.LifecycleStateDeleted
	}
	return false
}

func CloudGuardAdhocQuerySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CloudGuardClient().GetAdhocQuery(context.Background(), oci_cloud_guard.GetAdhocQueryRequest{
		AdhocQueryId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
