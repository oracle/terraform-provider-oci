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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSensitiveTypesExportRequiredOnlyResource = DataSafeSensitiveTypesExportResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_types_export", "test_sensitive_types_export", acctest.Required, acctest.Create, DataSafeSensitiveTypesExportRepresentation)

	DataSafeSensitiveTypesExportResourceConfig = DataSafeSensitiveTypesExportResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_types_export", "test_sensitive_types_export", acctest.Optional, acctest.Update, DataSafeSensitiveTypesExportRepresentation)

	DataSafeSensitiveTypesExportSingularDataSourceRepresentation = map[string]interface{}{
		"sensitive_types_export_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_types_export.test_sensitive_types_export.id}`},
	}

	DataSafeSensitiveTypesExportDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                          acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"sensitive_types_export_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_types_export.test_sensitive_types_export.id}`},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSensitiveTypesExportDataSourceFilterRepresentation}}
	DataSafeSensitiveTypesExportDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_sensitive_types_export.test_sensitive_types_export.id}`}},
	}

	DataSafeSensitiveTypesExportRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_include_all_sensitive_types": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSensitiveTypesExportSystemTagsChangesRep},
	}

	DataSafeSensitiveTypesExportResourceDependencies = DefinedTagsDependencies

	ignoreSensitiveTypesExportSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveTypesExportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveTypesExportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_sensitive_types_export.test_sensitive_types_export"
	datasourceName := "data.oci_data_safe_sensitive_types_exports.test_sensitive_types_exports"
	singularDatasourceName := "data.oci_data_safe_sensitive_types_export.test_sensitive_types_export"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeSensitiveTypesExportResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_types_export", "test_sensitive_types_export", acctest.Optional, acctest.Create, DataSafeSensitiveTypesExportRepresentation), "datasafe", "sensitiveTypesExport", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSensitiveTypesExportDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypesExportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_types_export", "test_sensitive_types_export", acctest.Required, acctest.Create, DataSafeSensitiveTypesExportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypesExportResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypesExportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_types_export", "test_sensitive_types_export", acctest.Optional, acctest.Create, DataSafeSensitiveTypesExportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_include_all_sensitive_types", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_type_ids_for_export.#"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeSensitiveTypesExportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_types_export", "test_sensitive_types_export", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeSensitiveTypesExportRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_include_all_sensitive_types", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_type_ids_for_export.#"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypesExportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_types_export", "test_sensitive_types_export", acctest.Optional, acctest.Update, DataSafeSensitiveTypesExportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_include_all_sensitive_types", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_type_ids_for_export.#"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_types_exports", "test_sensitive_types_exports", acctest.Optional, acctest.Update, DataSafeSensitiveTypesExportDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSensitiveTypesExportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_types_export", "test_sensitive_types_export", acctest.Optional, acctest.Update, DataSafeSensitiveTypesExportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_types_export_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),

				resource.TestCheckResourceAttr(datasourceName, "sensitive_types_export_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sensitive_types_export_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_types_export", "test_sensitive_types_export", acctest.Required, acctest.Create, DataSafeSensitiveTypesExportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSensitiveTypesExportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_types_export_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_include_all_sensitive_types", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_type_ids_for_export.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeSensitiveTypesExportRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSensitiveTypesExportDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_sensitive_types_export" {
			noResourceFound = false
			request := oci_data_safe.GetSensitiveTypesExportRequest{}

			tmp := rs.Primary.ID
			request.SensitiveTypesExportId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSensitiveTypesExport(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SensitiveTypesExportLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeSensitiveTypesExport") {
		resource.AddTestSweepers("DataSafeSensitiveTypesExport", &resource.Sweeper{
			Name:         "DataSafeSensitiveTypesExport",
			Dependencies: acctest.DependencyGraph["sensitiveTypesExport"],
			F:            sweepDataSafeSensitiveTypesExportResource,
		})
	}
}

func sweepDataSafeSensitiveTypesExportResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	sensitiveTypesExportIds, err := getDataSafeSensitiveTypesExportIds(compartment)
	if err != nil {
		return err
	}
	for _, sensitiveTypesExportId := range sensitiveTypesExportIds {
		if ok := acctest.SweeperDefaultResourceId[sensitiveTypesExportId]; !ok {
			deleteSensitiveTypesExportRequest := oci_data_safe.DeleteSensitiveTypesExportRequest{}

			deleteSensitiveTypesExportRequest.SensitiveTypesExportId = &sensitiveTypesExportId

			deleteSensitiveTypesExportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSensitiveTypesExport(context.Background(), deleteSensitiveTypesExportRequest)
			if error != nil {
				fmt.Printf("Error deleting SensitiveTypesExport %s %s, It is possible that the resource is already deleted. Please verify manually \n", sensitiveTypesExportId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sensitiveTypesExportId, DataSafeSensitiveTypesExportSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSensitiveTypesExportSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSensitiveTypesExportIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SensitiveTypesExportId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSensitiveTypesExportsRequest := oci_data_safe.ListSensitiveTypesExportsRequest{}
	listSensitiveTypesExportsRequest.CompartmentId = &compartmentId
	listSensitiveTypesExportsRequest.LifecycleState = oci_data_safe.ListSensitiveTypesExportsLifecycleStateActive
	listSensitiveTypesExportsResponse, err := dataSafeClient.ListSensitiveTypesExports(context.Background(), listSensitiveTypesExportsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SensitiveTypesExport list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sensitiveTypesExport := range listSensitiveTypesExportsResponse.Items {
		id := *sensitiveTypesExport.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SensitiveTypesExportId", id)
	}
	return resourceIds, nil
}

func DataSafeSensitiveTypesExportSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sensitiveTypesExportResponse, ok := response.Response.(oci_data_safe.GetSensitiveTypesExportResponse); ok {
		return sensitiveTypesExportResponse.LifecycleState != oci_data_safe.SensitiveTypesExportLifecycleStateDeleted
	}
	return false
}

func DataSafeSensitiveTypesExportSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSensitiveTypesExport(context.Background(), oci_data_safe.GetSensitiveTypesExportRequest{
		SensitiveTypesExportId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
