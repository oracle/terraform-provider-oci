// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"

	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_management_dashboard "github.com/oracle/oci-go-sdk/v65/managementdashboard"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ManagementDashboardManagementSavedSearchRequiredOnlyResource = ManagementDashboardManagementSavedSearchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_management_dashboard_management_saved_search", "test_management_saved_search", acctest.Required, acctest.Create, ManagementDashboardManagementSavedSearchRepresentation)

	ManagementDashboardManagementSavedSearchResourceConfig = ManagementDashboardManagementSavedSearchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_management_dashboard_management_saved_search", "test_management_saved_search", acctest.Optional, acctest.Update, ManagementDashboardManagementSavedSearchRepresentation)

	ManagementDashboardManagementSavedSearchSingularDataSourceRepresentation = map[string]interface{}{
		"management_saved_search_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_management_dashboard_management_saved_search.test_management_saved_search.id}`},
	}

	ManagementDashboardManagementSavedSearchDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ManagementDashboardManagementSavedSearchDataSourceFilterRepresentation}}
	ManagementDashboardManagementSavedSearchDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_management_dashboard_management_saved_search.test_management_saved_search.id}`}},
	}

	ManagementDashboardManagementSavedSearchRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_config":         acctest.Representation{RepType: acctest.Required, Create: `[{\n  \"key1\": \"key2\" \n}]`, Update: `[{\n  \"key3\": \"key4\" \n}]`},
		"description":         acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"is_oob_saved_search": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"metadata_version":    acctest.Representation{RepType: acctest.Required, Create: `2.0`, Update: `2.0`},
		"nls":                 acctest.Representation{RepType: acctest.Required, Create: `{\n  \"key1\": \"key2\",\n      \"key3\": \"key4\" \n}`},
		"provider_id":         acctest.Representation{RepType: acctest.Required, Create: `management-dashboard`},
		"provider_name":       acctest.Representation{RepType: acctest.Required, Create: `providerName`, Update: `providerName2`},
		"provider_version":    acctest.Representation{RepType: acctest.Required, Create: `providerVersion`, Update: `providerVersion2`},
		"screen_image":        acctest.Representation{RepType: acctest.Required, Create: `screenImage`, Update: `screenImage2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `SEARCH_SHOW_IN_DASHBOARD`, Update: `SEARCH_DONT_SHOW_IN_DASHBOARD`},
		"ui_config":           acctest.Representation{RepType: acctest.Required, Create: `{\n  \"key1\": \"key2\",\n      \"key3\": \"key4\" \n}`},
		"widget_template":     acctest.Representation{RepType: acctest.Required, Create: `widgetTemplate`, Update: `widgetTemplate2`},
		"widget_vm":           acctest.Representation{RepType: acctest.Required, Create: `widgetVM`, Update: `widgetVM2`},
		// "defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"drilldown_config": acctest.Representation{RepType: acctest.Required, Create: `[{\n  \"key1\": \"key2\" \n}]`, Update: `[{\n  \"key3\": \"key4\" \n}]`},
		// "freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"parameters_config": acctest.Representation{RepType: acctest.Required, Create: `[{\n  \"key1\": \"key2\" \n}]`, Update: `[{\n  \"key3\": \"key4\" \n}]`},
	}

	ManagementDashboardManagementSavedSearchResourceDependencies = ""
)

// issue-routing-tag: management_dashboard/default
func TestManagementDashboardManagementSavedSearchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementDashboardManagementSavedSearchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_management_dashboard_management_saved_search.test_management_saved_search"
	datasourceName := "data.oci_management_dashboard_management_saved_searches.test_management_saved_searches"
	singularDatasourceName := "data.oci_management_dashboard_management_saved_search.test_management_saved_search"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ManagementDashboardManagementSavedSearchResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_management_dashboard_management_saved_search", "test_management_saved_search", acctest.Required, acctest.Create, ManagementDashboardManagementSavedSearchRepresentation), "managementdashboard", "managementSavedSearch", t)

	acctest.ResourceTest(t, testAccCheckManagementDashboardManagementSavedSearchDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ManagementDashboardManagementSavedSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_dashboard_management_saved_search", "test_management_saved_search", acctest.Required, acctest.Create, ManagementDashboardManagementSavedSearchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_config", "[{\"key1\":\"key2\"}]"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "is_oob_saved_search", "false"),
				resource.TestCheckResourceAttr(resourceName, "metadata_version", "2.0"),
				resource.TestCheckResourceAttr(resourceName, "nls", "{\"key1\":\"key2\",\"key3\":\"key4\"}"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_name", "providerName"),
				resource.TestCheckResourceAttr(resourceName, "provider_version", "providerVersion"),
				resource.TestCheckResourceAttr(resourceName, "screen_image", "screenImage"),
				resource.TestCheckResourceAttr(resourceName, "type", "SEARCH_SHOW_IN_DASHBOARD"),
				resource.TestCheckResourceAttr(resourceName, "ui_config", "{\"key1\":\"key2\",\"key3\":\"key4\"}"),
				resource.TestCheckResourceAttr(resourceName, "widget_template", "widgetTemplate"),
				resource.TestCheckResourceAttr(resourceName, "widget_vm", "widgetVM"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ManagementDashboardManagementSavedSearchResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ManagementDashboardManagementSavedSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_dashboard_management_saved_search", "test_management_saved_search", acctest.Required, acctest.Create, ManagementDashboardManagementSavedSearchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "data_config", "[{\"key1\":\"key2\"}]"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "drilldown_config", "[{\"key1\":\"key2\"}]"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				// resource.TestCheckResourceAttr(resourceName, "id", "id"),
				resource.TestCheckResourceAttr(resourceName, "is_oob_saved_search", "false"),
				resource.TestCheckResourceAttr(resourceName, "metadata_version", "2.0"),
				resource.TestCheckResourceAttr(resourceName, "nls", "{\"key1\":\"key2\",\"key3\":\"key4\"}"),
				resource.TestCheckResourceAttr(resourceName, "parameters_config", "[{\"key1\":\"key2\"}]"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_name", "providerName"),
				resource.TestCheckResourceAttr(resourceName, "provider_version", "providerVersion"),
				resource.TestCheckResourceAttr(resourceName, "screen_image", "screenImage"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "SEARCH_SHOW_IN_DASHBOARD"),
				resource.TestCheckResourceAttr(resourceName, "ui_config", "{\"key1\":\"key2\",\"key3\":\"key4\"}"),
				resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
				resource.TestCheckResourceAttr(resourceName, "widget_template", "widgetTemplate"),
				resource.TestCheckResourceAttr(resourceName, "widget_vm", "widgetVM"),

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

		// // verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ManagementDashboardManagementSavedSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_dashboard_management_saved_search", "test_management_saved_search", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ManagementDashboardManagementSavedSearchRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "data_config", "[{\"key1\":\"key2\"}]"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "drilldown_config", "[{\"key1\":\"key2\"}]"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				// resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_oob_saved_search", "false"),
				resource.TestCheckResourceAttr(resourceName, "metadata_version", "2.0"),
				resource.TestCheckResourceAttr(resourceName, "nls", "{\"key1\":\"key2\",\"key3\":\"key4\"}"),
				resource.TestCheckResourceAttr(resourceName, "parameters_config", "[{\"key1\":\"key2\"}]"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_name", "providerName"),
				resource.TestCheckResourceAttr(resourceName, "provider_version", "providerVersion"),
				resource.TestCheckResourceAttr(resourceName, "screen_image", "screenImage"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "SEARCH_SHOW_IN_DASHBOARD"),
				resource.TestCheckResourceAttr(resourceName, "ui_config", "{\"key1\":\"key2\",\"key3\":\"key4\"}"),
				resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
				resource.TestCheckResourceAttr(resourceName, "widget_template", "widgetTemplate"),
				resource.TestCheckResourceAttr(resourceName, "widget_vm", "widgetVM"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// // verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ManagementDashboardManagementSavedSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_dashboard_management_saved_search", "test_management_saved_search", acctest.Optional, acctest.Update, ManagementDashboardManagementSavedSearchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "data_config", "[{\"key3\":\"key4\"}]"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "drilldown_config", "[{\"key3\":\"key4\"}]"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				// resource.TestCheckResourceAttr(resourceName, "id", "id"),
				resource.TestCheckResourceAttr(resourceName, "is_oob_saved_search", "true"),
				resource.TestCheckResourceAttr(resourceName, "metadata_version", "2.0"),
				resource.TestCheckResourceAttr(resourceName, "nls", "{\"key1\":\"key2\",\"key3\":\"key4\"}"),
				resource.TestCheckResourceAttr(resourceName, "parameters_config", "[{\"key3\":\"key4\"}]"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_name", "providerName2"),
				resource.TestCheckResourceAttr(resourceName, "provider_version", "providerVersion2"),
				resource.TestCheckResourceAttr(resourceName, "screen_image", "screenImage2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "SEARCH_DONT_SHOW_IN_DASHBOARD"),
				resource.TestCheckResourceAttr(resourceName, "ui_config", "{\"key1\":\"key2\",\"key3\":\"key4\"}"),
				resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
				resource.TestCheckResourceAttr(resourceName, "widget_template", "widgetTemplate2"),
				resource.TestCheckResourceAttr(resourceName, "widget_vm", "widgetVM2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// // verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_dashboard_management_saved_searches", "test_management_saved_searches", acctest.Optional, acctest.Update, ManagementDashboardManagementSavedSearchDataSourceRepresentation) +
				compartmentIdVariableStr + ManagementDashboardManagementSavedSearchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_dashboard_management_saved_search", "test_management_saved_search", acctest.Optional, acctest.Update, ManagementDashboardManagementSavedSearchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "management_saved_search_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "management_saved_search_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_dashboard_management_saved_search", "test_management_saved_search", acctest.Required, acctest.Create, ManagementDashboardManagementSavedSearchSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagementDashboardManagementSavedSearchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_saved_search_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_config", "[{\"key3\":\"key4\"}]"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "drilldown_config", "[{\"key3\":\"key4\"}]"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_oob_saved_search", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata_version", "2.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nls", "{\"key1\":\"key2\",\"key3\":\"key4\"}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters_config", "[{\"key3\":\"key4\"}]"),
				resource.TestCheckResourceAttr(singularDatasourceName, "provider_name", "providerName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "provider_version", "providerVersion2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "screen_image", "screenImage2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "SEARCH_DONT_SHOW_IN_DASHBOARD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ui_config", "{\"key1\":\"key2\",\"key3\":\"key4\"}"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "updated_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "widget_template", "widgetTemplate2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "widget_vm", "widgetVM2"),
			),
		},
		// verify resource import
		{
			Config:                  config + ManagementDashboardManagementSavedSearchRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckManagementDashboardManagementSavedSearchDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DashxApisClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_management_dashboard_management_saved_search" {
			noResourceFound = false
			request := oci_management_dashboard.GetManagementSavedSearchRequest{}

			tmp := rs.Primary.ID
			request.ManagementSavedSearchId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_dashboard")

			_, err := client.GetManagementSavedSearch(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("ManagementDashboardManagementSavedSearch") {
		resource.AddTestSweepers("ManagementDashboardManagementSavedSearch", &resource.Sweeper{
			Name:         "ManagementDashboardManagementSavedSearch",
			Dependencies: acctest.DependencyGraph["managementSavedSearch"],
			F:            sweepManagementDashboardManagementSavedSearchResource,
		})
	}
}

func sweepManagementDashboardManagementSavedSearchResource(compartment string) error {
	dashxApisClient := acctest.GetTestClients(&schema.ResourceData{}).DashxApisClient()
	managementSavedSearchIds, err := getManagementDashboardManagementSavedSearchIds(compartment)
	if err != nil {
		return err
	}
	for _, managementSavedSearchId := range managementSavedSearchIds {
		if ok := acctest.SweeperDefaultResourceId[managementSavedSearchId]; !ok {
			deleteManagementSavedSearchRequest := oci_management_dashboard.DeleteManagementSavedSearchRequest{}

			deleteManagementSavedSearchRequest.ManagementSavedSearchId = &managementSavedSearchId

			deleteManagementSavedSearchRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_dashboard")
			_, error := dashxApisClient.DeleteManagementSavedSearch(context.Background(), deleteManagementSavedSearchRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagementSavedSearch %s %s, It is possible that the resource is already deleted. Please verify manually \n", managementSavedSearchId, error)
				continue
			}
		}
	}
	return nil
}

func getManagementDashboardManagementSavedSearchIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagementSavedSearchId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dashxApisClient := acctest.GetTestClients(&schema.ResourceData{}).DashxApisClient()

	listManagementSavedSearchesRequest := oci_management_dashboard.ListManagementSavedSearchesRequest{}
	listManagementSavedSearchesRequest.CompartmentId = &compartmentId
	listManagementSavedSearchesResponse, err := dashxApisClient.ListManagementSavedSearches(context.Background(), listManagementSavedSearchesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagementSavedSearch list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managementSavedSearch := range listManagementSavedSearchesResponse.Items {
		id := *managementSavedSearch.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagementSavedSearchId", id)
	}
	return resourceIds, nil
}
