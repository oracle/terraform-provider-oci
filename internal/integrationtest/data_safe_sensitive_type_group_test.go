// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
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
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSensitiveTypeGroupRequiredOnlyResource = DataSafeSensitiveTypeGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group", "test_sensitive_type_group", acctest.Required, acctest.Create, DataSafeSensitiveTypeGroupRepresentation)

	DataSafeSensitiveTypeGroupResourceConfig = DataSafeSensitiveTypeGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group", "test_sensitive_type_group", acctest.Optional, acctest.Update, DataSafeSensitiveTypeGroupRepresentation)

	DataSafeSensitiveTypeGroupSingularDataSourceRepresentation = map[string]interface{}{
		"sensitive_type_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_type_group.test_sensitive_type_group.id}`},
	}

	DataSafeSensitiveTypeGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                          acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"sensitive_type_group_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_type_group.test_sensitive_type_group.id}`},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSensitiveTypeGroupDataSourceFilterRepresentation}}
	DataSafeSensitiveTypeGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_sensitive_type_group.test_sensitive_type_group.id}`}},
	}

	DataSafeSensitiveTypeGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSensitiveTypeGroupSystemTagsChangesRep},
	}

	ignoreSensitiveTypeGroupSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSensitiveTypeGroupResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveTypeGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveTypeGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_sensitive_type_group.test_sensitive_type_group"
	datasourceName := "data.oci_data_safe_sensitive_type_groups.test_sensitive_type_groups"
	singularDatasourceName := "data.oci_data_safe_sensitive_type_group.test_sensitive_type_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeSensitiveTypeGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group", "test_sensitive_type_group", acctest.Optional, acctest.Create, DataSafeSensitiveTypeGroupRepresentation), "datasafe", "sensitiveTypeGroup", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSensitiveTypeGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypeGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group", "test_sensitive_type_group", acctest.Required, acctest.Create, DataSafeSensitiveTypeGroupRepresentation),
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
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypeGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypeGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group", "test_sensitive_type_group", acctest.Optional, acctest.Create, DataSafeSensitiveTypeGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_type_count"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeSensitiveTypeGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group", "test_sensitive_type_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeSensitiveTypeGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_type_count"),
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
			Config: config + compartmentIdVariableStr + DataSafeSensitiveTypeGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group", "test_sensitive_type_group", acctest.Optional, acctest.Update, DataSafeSensitiveTypeGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_type_count"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_type_groups", "test_sensitive_type_groups", acctest.Optional, acctest.Update, DataSafeSensitiveTypeGroupDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSensitiveTypeGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_type_group", "test_sensitive_type_group", acctest.Optional, acctest.Update, DataSafeSensitiveTypeGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_type_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),

				resource.TestCheckResourceAttr(datasourceName, "sensitive_type_group_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sensitive_type_group_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_type_group", "test_sensitive_type_group", acctest.Required, acctest.Create, DataSafeSensitiveTypeGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSensitiveTypeGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_type_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_type_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeSensitiveTypeGroupRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSensitiveTypeGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_sensitive_type_group" {
			noResourceFound = false
			request := oci_data_safe.GetSensitiveTypeGroupRequest{}

			tmp := rs.Primary.ID
			request.SensitiveTypeGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSensitiveTypeGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SensitiveTypeGroupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeSensitiveTypeGroup") {
		resource.AddTestSweepers("DataSafeSensitiveTypeGroup", &resource.Sweeper{
			Name:         "DataSafeSensitiveTypeGroup",
			Dependencies: acctest.DependencyGraph["sensitiveTypeGroup"],
			F:            sweepDataSafeSensitiveTypeGroupResource,
		})
	}
}

func sweepDataSafeSensitiveTypeGroupResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	sensitiveTypeGroupIds, err := getDataSafeSensitiveTypeGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, sensitiveTypeGroupId := range sensitiveTypeGroupIds {
		if ok := acctest.SweeperDefaultResourceId[sensitiveTypeGroupId]; !ok {
			deleteSensitiveTypeGroupRequest := oci_data_safe.DeleteSensitiveTypeGroupRequest{}

			deleteSensitiveTypeGroupRequest.SensitiveTypeGroupId = &sensitiveTypeGroupId

			deleteSensitiveTypeGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSensitiveTypeGroup(context.Background(), deleteSensitiveTypeGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting SensitiveTypeGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", sensitiveTypeGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sensitiveTypeGroupId, DataSafeSensitiveTypeGroupSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSensitiveTypeGroupSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSensitiveTypeGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SensitiveTypeGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSensitiveTypeGroupsRequest := oci_data_safe.ListSensitiveTypeGroupsRequest{}
	listSensitiveTypeGroupsRequest.CompartmentId = &compartmentId
	listSensitiveTypeGroupsRequest.LifecycleState = oci_data_safe.ListSensitiveTypeGroupsLifecycleStateActive
	listSensitiveTypeGroupsResponse, err := dataSafeClient.ListSensitiveTypeGroups(context.Background(), listSensitiveTypeGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SensitiveTypeGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sensitiveTypeGroup := range listSensitiveTypeGroupsResponse.Items {
		id := *sensitiveTypeGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SensitiveTypeGroupId", id)
	}
	return resourceIds, nil
}

func DataSafeSensitiveTypeGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sensitiveTypeGroupResponse, ok := response.Response.(oci_data_safe.GetSensitiveTypeGroupResponse); ok {
		return sensitiveTypeGroupResponse.LifecycleState != oci_data_safe.SensitiveTypeGroupLifecycleStateDeleted
	}
	return false
}

func DataSafeSensitiveTypeGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSensitiveTypeGroup(context.Background(), oci_data_safe.GetSensitiveTypeGroupRequest{
		SensitiveTypeGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
