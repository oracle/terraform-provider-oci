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
	DataSafeSensitiveTypeRequiredOnlyResource = DataSafeSensitiveTypeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type", "test_sensitive_type", acctest.Required, acctest.Create, sensitiveTypeRepresentation)

	DataSafeSensitiveTypeResourceConfig = DataSafeSensitiveTypeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type", "test_sensitive_type", acctest.Optional, acctest.Update, sensitiveTypeRepresentation)

	DataSafesensitiveTypeSingularDataSourceRepresentation = map[string]interface{}{
		"sensitive_type_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_type.test_sensitive_type.id}`},
	}

	DataSafeSensitiveTypeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"entity_type":               acctest.Representation{RepType: acctest.Optional, Create: `SENSITIVE_TYPE`, Update: `SENSITIVE_TYPE`},
		//"is_common":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"sensitive_type_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_type.test_sensitive_type.id}`},
		//"sensitive_type_source":     acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: sensitiveTypeDataSourceFilterRepresentation},
	}

	sensitiveTypeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_sensitive_type.test_sensitive_type.id}`}},
	}

	sensitiveTypeRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"entity_type":     acctest.Representation{RepType: acctest.Required, Create: `SENSITIVE_TYPE`, Update: `SENSITIVE_TYPE`},
		"comment_pattern": acctest.Representation{RepType: acctest.Optional, Create: `commentPattern`, Update: `commentPattern2`},
		"data_pattern":    acctest.Representation{RepType: acctest.Optional, Create: `dataPattern`, Update: `dataPattern2`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"name_pattern":    acctest.Representation{RepType: acctest.Optional, Create: `namePattern`, Update: `namePattern2`},
		"search_type":     acctest.Representation{RepType: acctest.Optional, Create: `OR`, Update: `AND`},
		"short_name":      acctest.Representation{RepType: acctest.Optional, Create: `shortName`, Update: `shortName2`},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSensitiveTypeSystemTagsChangesRep},
	}

	ignoreSensitiveTypeSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSensitiveTypeResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_sensitive_type.test_sensitive_type"
	datasourceName := "data.oci_data_safe_sensitive_types.test_sensitive_types"
	singularDatasourceName := "data.oci_data_safe_sensitive_type.test_sensitive_type"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeSensitiveTypeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type", "test_sensitive_type", acctest.Optional, acctest.Create, sensitiveTypeRepresentation), "datasafe", "sensitiveType", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSensitiveTypeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type", "test_sensitive_type", acctest.Required, acctest.Create, sensitiveTypeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "entity_type", "SENSITIVE_TYPE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type", "test_sensitive_type", acctest.Optional, acctest.Create, sensitiveTypeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "comment_pattern", "commentPattern"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_pattern", "dataPattern"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "entity_type", "SENSITIVE_TYPE"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name_pattern", "namePattern"),
				resource.TestCheckResourceAttr(resourceName, "search_type", "OR"),
				resource.TestCheckResourceAttr(resourceName, "short_name", "shortName"),
				resource.TestCheckResourceAttrSet(resourceName, "source"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeSensitiveTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type", "test_sensitive_type", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(sensitiveTypeRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "comment_pattern", "commentPattern"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "data_pattern", "dataPattern"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "entity_type", "SENSITIVE_TYPE"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name_pattern", "namePattern"),
				resource.TestCheckResourceAttr(resourceName, "search_type", "OR"),
				resource.TestCheckResourceAttr(resourceName, "short_name", "shortName"),
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
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type", "test_sensitive_type", acctest.Optional, acctest.Update, sensitiveTypeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "comment_pattern", "commentPattern2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_pattern", "dataPattern2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "entity_type", "SENSITIVE_TYPE"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name_pattern", "namePattern2"),
				resource.TestCheckResourceAttr(resourceName, "search_type", "AND"),
				resource.TestCheckResourceAttr(resourceName, "short_name", "shortName2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_types", "test_sensitive_types", acctest.Optional, acctest.Update, DataSafeSensitiveTypeDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSensitiveTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type", "test_sensitive_type", acctest.Optional, acctest.Update, sensitiveTypeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "entity_type", "SENSITIVE_TYPE"),
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_type_id"),
				resource.TestCheckResourceAttr(datasourceName, "sensitive_type_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sensitive_type_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_type", "test_sensitive_type", acctest.Required, acctest.Create, DataSafesensitiveTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSensitiveTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_type_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_common"),
				resource.TestCheckResourceAttr(singularDatasourceName, "short_name", "shortName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypeResourceConfig,
		},
		// verify resource import
		{
			Config:                  config + DataSafeSensitiveTypeRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSensitiveTypeDestroy(s *terraform.State) error {
	fmt.Printf("testAccCheckDataSafeSensitiveTypeDestroy  ***** CALLED ***** \n")
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_sensitive_type" {
			noResourceFound = false
			request := oci_data_safe.GetSensitiveTypeRequest{}

			tmp := rs.Primary.ID
			request.SensitiveTypeId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSensitiveType(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.DiscoveryLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			// Verify that exception is for '404 not found'.
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
	if !acctest.InSweeperExcludeList("DataSafeSensitiveType") {
		resource.AddTestSweepers("DataSafeSensitiveType", &resource.Sweeper{
			Name:         "DataSafeSensitiveType",
			Dependencies: acctest.DependencyGraph["sensitiveType"],
			F:            sweepDataSafeSensitiveTypeResource,
		})
	}
}

func sweepDataSafeSensitiveTypeResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	sensitiveTypeIds, err := getDataSafeSensitiveTypeIds(compartment)
	if err != nil {
		return err
	}
	for _, sensitiveTypeId := range sensitiveTypeIds {
		if ok := acctest.SweeperDefaultResourceId[sensitiveTypeId]; !ok {
			deleteSensitiveTypeRequest := oci_data_safe.DeleteSensitiveTypeRequest{}

			deleteSensitiveTypeRequest.SensitiveTypeId = &sensitiveTypeId

			deleteSensitiveTypeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSensitiveType(context.Background(), deleteSensitiveTypeRequest)
			if error != nil {
				fmt.Printf("Error deleting SensitiveType %s %s, It is possible that the resource is already deleted. Please verify manually \n", sensitiveTypeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sensitiveTypeId, DataSafesensitiveTypesSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafesensitiveTypesSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSensitiveTypeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SensitiveTypeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSensitiveTypesRequest := oci_data_safe.ListSensitiveTypesRequest{}
	listSensitiveTypesRequest.CompartmentId = &compartmentId
	listSensitiveTypesRequest.LifecycleState = oci_data_safe.ListSensitiveTypesLifecycleStateActive
	listSensitiveTypesResponse, err := dataSafeClient.ListSensitiveTypes(context.Background(), listSensitiveTypesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SensitiveType list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sensitiveType := range listSensitiveTypesResponse.Items {
		id := *sensitiveType.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SensitiveTypeId", id)
		acctest.SweeperDefaultResourceId[*sensitiveType.DefaultMaskingFormatId] = true

	}
	return resourceIds, nil
}

func DataSafesensitiveTypesSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sensitiveTypeResponse, ok := response.Response.(oci_data_safe.GetSensitiveTypeResponse); ok {
		return sensitiveTypeResponse.GetLifecycleState() != oci_data_safe.DiscoveryLifecycleStateDeleted
	}
	return false
}

func DataSafesensitiveTypesSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSensitiveType(context.Background(), oci_data_safe.GetSensitiveTypeRequest{
		SensitiveTypeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
