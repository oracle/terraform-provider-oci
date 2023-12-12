// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafeLibraryMaskingFormatRequiredOnlyResource = DataSafeLibraryMaskingFormatResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_library_masking_format", "test_library_masking_format", acctest.Required, acctest.Create, libraryMaskingFormatRepresentation)

	DataSafeLibraryMaskingFormatResourceConfig = DataSafeLibraryMaskingFormatResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_library_masking_format", "test_library_masking_format", acctest.Optional, acctest.Update, libraryMaskingFormatRepresentation)

	DataSafelibraryMaskingFormatSingularDataSourceRepresentation = map[string]interface{}{
		"library_masking_format_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_library_masking_format.test_library_masking_format.id}`},
	}

	DataSafelibraryMaskingFormatDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: libraryMaskingFormatDataSourceFilterRepresentation},
	}
	libraryMaskingFormatDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_library_masking_format.test_library_masking_format.id}`}},
	}

	libraryMaskingFormatRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"format_entries": acctest.RepresentationGroup{RepType: acctest.Required, Group: libraryMaskingFormatFormatEntriesRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	libraryMaskingFormatFormatEntriesRepresentation = map[string]interface{}{
		"type":         acctest.Representation{RepType: acctest.Required, Create: `DELETE_ROWS`, Update: `FIXED_STRING`},
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"fixed_string": acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `fixedString2`},
	}

	DataSafeLibraryMaskingFormatFormatEntriesRepresentation = map[string]interface{}{
		"type":                      acctest.Representation{RepType: acctest.Required, Create: `DELETE_ROWS`, Update: `DETERMINISTIC_SUBSTITUTION`},
		"column_name":               acctest.Representation{RepType: acctest.Optional, Create: `columnName`, Update: `columnName2`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"end_date":                  acctest.Representation{RepType: acctest.Optional, Create: `endDate`, Update: `endDate2`},
		"end_length":                acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"end_value":                 acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.1`},
		"fixed_number":              acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.1`},
		"fixed_string":              acctest.Representation{RepType: acctest.Optional, Create: `fixedString`, Update: `fixedString2`},
		"grouping_columns":          acctest.Representation{RepType: acctest.Optional, Create: []string{`groupingColumns`}, Update: []string{`groupingColumns2`}},
		"length":                    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"library_masking_format_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_library_masking_format.test_library_masking_format.id}`},
		"pattern":                   acctest.Representation{RepType: acctest.Optional, Create: `pattern`, Update: `pattern2`},
		"post_processing_function":  acctest.Representation{RepType: acctest.Optional, Create: `postProcessingFunction`, Update: `postProcessingFunction2`},
		"random_list":               acctest.Representation{RepType: acctest.Optional, Create: []string{`randomList`}, Update: []string{`randomList2`}},
		"regular_expression":        acctest.Representation{RepType: acctest.Optional, Create: `regularExpression`, Update: `regularExpression2`},
		"replace_with":              acctest.Representation{RepType: acctest.Optional, Create: `replaceWith`, Update: `replaceWith2`},
		"schema_name":               acctest.Representation{RepType: acctest.Optional, Create: `schemaName`, Update: `schemaName2`},
		"sql_expression":            acctest.Representation{RepType: acctest.Optional, Create: `sqlExpression`, Update: `sqlExpression2`},
		"start_date":                acctest.Representation{RepType: acctest.Optional, Create: `startDate`, Update: `startDate2`},
		"start_length":              acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"start_position":            acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"start_value":               acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.1`},
		"table_name":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_nosql_table.test_table.name}`},
		"user_defined_function":     acctest.Representation{RepType: acctest.Optional, Create: `userDefinedFunction`, Update: `userDefinedFunction2`},
	}

	DataSafeLibraryMaskingFormatResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeLibraryMaskingFormatResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeLibraryMaskingFormatResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_library_masking_format.test_library_masking_format"
	datasourceName := "data.oci_data_safe_library_masking_formats.test_library_masking_formats"
	singularDatasourceName := "oci_data_safe_library_masking_format.test_library_masking_format"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		DataSafeLibraryMaskingFormatResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_library_masking_format", "test_library_masking_format", acctest.Optional, acctest.Create, libraryMaskingFormatRepresentation), "datasafe", "libraryMaskingFormat", t)

	acctest.ResourceTest(t, testAccCheckDataSafeLibraryMaskingFormatDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeLibraryMaskingFormatResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_library_masking_format", "test_library_masking_format", acctest.Required, acctest.Create, libraryMaskingFormatRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "format_entries.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "format_entries.0.type", "DELETE_ROWS"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeLibraryMaskingFormatResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeLibraryMaskingFormatResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_library_masking_format", "test_library_masking_format", acctest.Optional, acctest.Create, libraryMaskingFormatRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "format_entries.0.type", "DELETE_ROWS"),
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

		//verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeLibraryMaskingFormatResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_library_masking_format", "test_library_masking_format", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(libraryMaskingFormatRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "format_entries.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "format_entries.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "format_entries.0.type", "DELETE_ROWS"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "source"),
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
			Config: config + compartmentIdVariableStr + DataSafeLibraryMaskingFormatResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_library_masking_format", "test_library_masking_format", acctest.Optional, acctest.Update, libraryMaskingFormatRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "format_entries.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "format_entries.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "format_entries.0.fixed_string", "fixedString2"),
				resource.TestCheckResourceAttr(resourceName, "format_entries.0.type", "FIXED_STRING"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "source"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_library_masking_formats", "test_library_masking_formats", acctest.Optional, acctest.Update, DataSafelibraryMaskingFormatDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeLibraryMaskingFormatResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_library_masking_format", "test_library_masking_format", acctest.Optional, acctest.Update, libraryMaskingFormatRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "library_masking_format_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "library_masking_format_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_library_masking_format", "test_library_masking_format", acctest.Required, acctest.Create, DataSafelibraryMaskingFormatSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeLibraryMaskingFormatResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "format_entries.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "format_entries.0.fixed_string", "fixedString2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sensitive_type_ids.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DataSafeLibraryMaskingFormatResourceConfig,
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

func testAccCheckDataSafeLibraryMaskingFormatDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_library_masking_format" {
			noResourceFound = false
			request := oci_data_safe.GetLibraryMaskingFormatRequest{}

			tmp := rs.Primary.ID
			request.LibraryMaskingFormatId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetLibraryMaskingFormat(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.MaskingLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeLibraryMaskingFormat") {
		resource.AddTestSweepers("DataSafeLibraryMaskingFormat", &resource.Sweeper{
			Name:         "DataSafeLibraryMaskingFormat",
			Dependencies: acctest.DependencyGraph["libraryMaskingFormat"],
			F:            sweepDataSafeLibraryMaskingFormatResource,
		})
	}
}

func sweepDataSafeLibraryMaskingFormatResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	libraryMaskingFormatIds, err := getDataSafeLibraryMaskingFormatIds(compartment)
	if err != nil {
		return err
	}
	for _, libraryMaskingFormatId := range libraryMaskingFormatIds {
		if ok := acctest.SweeperDefaultResourceId[libraryMaskingFormatId]; !ok {
			deleteLibraryMaskingFormatRequest := oci_data_safe.DeleteLibraryMaskingFormatRequest{}

			deleteLibraryMaskingFormatRequest.LibraryMaskingFormatId = &libraryMaskingFormatId

			deleteLibraryMaskingFormatRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteLibraryMaskingFormat(context.Background(), deleteLibraryMaskingFormatRequest)
			if error != nil {
				fmt.Printf("Error deleting LibraryMaskingFormat %s %s, It is possible that the resource is already deleted. Please verify manually \n", libraryMaskingFormatId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &libraryMaskingFormatId, DataSafelibraryMaskingFormatsSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafelibraryMaskingFormatsSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeLibraryMaskingFormatIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LibraryMaskingFormatId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listLibraryMaskingFormatsRequest := oci_data_safe.ListLibraryMaskingFormatsRequest{}
	listLibraryMaskingFormatsRequest.CompartmentId = &compartmentId
	listLibraryMaskingFormatsRequest.LifecycleState = oci_data_safe.ListLibraryMaskingFormatsLifecycleStateActive
	listLibraryMaskingFormatsResponse, err := dataSafeClient.ListLibraryMaskingFormats(context.Background(), listLibraryMaskingFormatsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting LibraryMaskingFormat list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, libraryMaskingFormat := range listLibraryMaskingFormatsResponse.Items {
		id := *libraryMaskingFormat.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LibraryMaskingFormatId", id)
	}
	return resourceIds, nil
}

func DataSafelibraryMaskingFormatsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if libraryMaskingFormatResponse, ok := response.Response.(oci_data_safe.GetLibraryMaskingFormatResponse); ok {
		return libraryMaskingFormatResponse.LifecycleState != oci_data_safe.MaskingLifecycleStateDeleted
	}
	return false
}

func DataSafelibraryMaskingFormatsSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetLibraryMaskingFormat(context.Background(), oci_data_safe.GetLibraryMaskingFormatRequest{
		LibraryMaskingFormatId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
