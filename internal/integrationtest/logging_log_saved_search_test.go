// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_logging "github.com/oracle/oci-go-sdk/v65/logging"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	LoggingLogSavedSearchRequiredOnlyResource = LoggingLogSavedSearchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_saved_search", "test_log_saved_search", acctest.Required, acctest.Create, LoggingLogSavedSearchRepresentation)

	LoggingLogSavedSearchResourceConfig = LoggingLogSavedSearchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_saved_search", "test_log_saved_search", acctest.Optional, acctest.Update, LoggingLogSavedSearchRepresentation)

	LoggingLoggingLogSavedSearchSingularDataSourceRepresentation = map[string]interface{}{
		"log_saved_search_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_saved_search.test_log_saved_search.id}`},
	}

	LoggingLoggingLogSavedSearchDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"log_saved_search_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_logging_log_saved_search.test_log_saved_search.id}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingLogSavedSearchDataSourceFilterRepresentation}}

	LoggingLogSavedSearchDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_logging_log_saved_search.test_log_saved_search.id}`}},
	}

	LoggingLogSavedSearchRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"query":          acctest.Representation{RepType: acctest.Required, Create: `query`, Update: `query2`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	LoggingLogSavedSearchResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: logging/default
func TestLoggingLogSavedSearchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingLogSavedSearchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_logging_log_saved_search.test_log_saved_search"
	datasourceName := "data.oci_logging_log_saved_searches.test_log_saved_searches"
	singularDatasourceName := "data.oci_logging_log_saved_search.test_log_saved_search"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LoggingLogSavedSearchResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_saved_search", "test_log_saved_search", acctest.Optional, acctest.Create, LoggingLogSavedSearchRepresentation), "logging", "logSavedSearch", t)

	acctest.ResourceTest(t, testAccCheckLoggingLogSavedSearchDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LoggingLogSavedSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log_saved_search", "test_log_saved_search", acctest.Required, acctest.Create, LoggingLogSavedSearchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "query", "query"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LoggingLogSavedSearchResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LoggingLogSavedSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log_saved_search", "test_log_saved_search", acctest.Optional, acctest.Create, LoggingLogSavedSearchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "query", "query"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LoggingLogSavedSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log_saved_search", "test_log_saved_search", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LoggingLogSavedSearchRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "query", "query"),

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
			Config: config + compartmentIdVariableStr + LoggingLogSavedSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log_saved_search", "test_log_saved_search", acctest.Optional, acctest.Update, LoggingLogSavedSearchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "query", "query2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_logging_log_saved_searches", "test_log_saved_searches", acctest.Optional, acctest.Update, LoggingLoggingLogSavedSearchDataSourceRepresentation) +
				compartmentIdVariableStr + LoggingLogSavedSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_log_saved_search", "test_log_saved_search", acctest.Optional, acctest.Update, LoggingLogSavedSearchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "log_saved_search_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),

				resource.TestCheckResourceAttr(datasourceName, "log_saved_search_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "log_saved_search_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_logging_log_saved_search", "test_log_saved_search", acctest.Required, acctest.Create, LoggingLoggingLogSavedSearchSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LoggingLogSavedSearchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_saved_search_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query", "query2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
			),
		},
		// verify resource import
		{
			Config:                  config + LoggingLogSavedSearchRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLoggingLogSavedSearchDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LoggingManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_logging_log_saved_search" {
			noResourceFound = false
			request := oci_logging.GetLogSavedSearchRequest{}

			tmp := rs.Primary.ID
			request.LogSavedSearchId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "logging")

			_, err := client.GetLogSavedSearch(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("LoggingLogSavedSearch") {
		resource.AddTestSweepers("LoggingLogSavedSearch", &resource.Sweeper{
			Name:         "LoggingLogSavedSearch",
			Dependencies: acctest.DependencyGraph["logSavedSearch"],
			F:            sweepLoggingLogSavedSearchResource,
		})
	}
}

func sweepLoggingLogSavedSearchResource(compartment string) error {
	loggingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).LoggingManagementClient()
	logSavedSearchIds, err := getLoggingLogSavedSearchIds(compartment)
	if err != nil {
		return err
	}
	for _, logSavedSearchId := range logSavedSearchIds {
		if ok := acctest.SweeperDefaultResourceId[logSavedSearchId]; !ok {
			deleteLogSavedSearchRequest := oci_logging.DeleteLogSavedSearchRequest{}

			deleteLogSavedSearchRequest.LogSavedSearchId = &logSavedSearchId

			deleteLogSavedSearchRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "logging")
			_, error := loggingManagementClient.DeleteLogSavedSearch(context.Background(), deleteLogSavedSearchRequest)
			if error != nil {
				fmt.Printf("Error deleting LogSavedSearch %s %s, It is possible that the resource is already deleted. Please verify manually \n", logSavedSearchId, error)
				continue
			}
		}
	}
	return nil
}

func getLoggingLogSavedSearchIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LogSavedSearchId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loggingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).LoggingManagementClient()

	listLogSavedSearchesRequest := oci_logging.ListLogSavedSearchesRequest{}
	listLogSavedSearchesRequest.CompartmentId = &compartmentId
	listLogSavedSearchesResponse, err := loggingManagementClient.ListLogSavedSearches(context.Background(), listLogSavedSearchesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting LogSavedSearch list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, logSavedSearch := range listLogSavedSearchesResponse.Items {
		id := *logSavedSearch.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LogSavedSearchId", id)
	}
	return resourceIds, nil
}
