// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

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
	DataSafeTargetDatabaseGroupRequiredOnlyResource = DataSafeTargetDatabaseGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_group", "test_target_database_group", acctest.Required, acctest.Create, DataSafeTargetDatabaseGroupRepresentation)

	DataSafeTargetDatabaseGroupResourceConfig = DataSafeTargetDatabaseGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_group", "test_target_database_group", acctest.Optional, acctest.Update, DataSafeTargetDatabaseGroupRepresentation)

	DataSafeTargetDatabaseGroupSingularDataSourceRepresentation = map[string]interface{}{
		"target_database_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_database_group.test_target_database_group.id}`},
	}

	DataSafeTargetDatabaseGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                          acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `TestGroup`, Update: `displayName2`},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"target_database_group_id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.target_database_id}`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2028-01-01T00:00:00.000Z`},
	}
	DataSafeTargetDatabaseGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `targetDatabaseOcid`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.target_database_id}`}},
	}

	DataSafeTargetDatabaseGroupRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `TestGroup`, Update: `displayName2`},
		"matching_criteria": acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeTargetDatabaseGroupMatchingCriteriaRepresentation},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `Description of the target database group (optional).`, Update: `description2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreTargetDatabaseGroupSystemTagsChangesRep},
	}
	DataSafeTargetDatabaseGroupMatchingCriteriaRepresentation = map[string]interface{}{
		"include": acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeTargetDatabaseGroupMatchingCriteriaIncludeRepresentation},
	}
	DataSafeTargetDatabaseGroupMatchingCriteriaIncludeRepresentation = map[string]interface{}{
		"compartments":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataSafeTargetDatabaseGroupMatchingCriteriaIncludeCompartmentsRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"target_database_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.target_database_id}`}, Update: []string{`${var.target_database_id2}`}},
	}
	DataSafeTargetDatabaseGroupMatchingCriteriaExcludeRepresentation = map[string]interface{}{
		"target_database_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.target_database_id2}`}, Update: []string{`${var.target_database_id}`}},
	}
	DataSafeTargetDatabaseGroupMatchingCriteriaIncludeCompartmentsRepresentation = map[string]interface{}{
		"id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id}`},
		"is_include_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
	}

	DataSafeTargetDatabaseGroupResourceDependencies = DefinedTagsDependencies

	ignoreTargetDatabaseGroupSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeTargetDatabaseGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeTargetDatabaseGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetDatabaseId := utils.GetEnvSettingWithBlankDefault("target_database_id")
	targetDatabaseIdVariableStr := fmt.Sprintf("variable \"target_database_id\" { default = \"%s\" }\n", targetDatabaseId)

	targetDatabaseId2 := utils.GetEnvSettingWithBlankDefault("target_database_id2")
	targetDatabaseIdVariableStr2 := fmt.Sprintf("variable \"target_database_id2\" { default = \"%s\" }\n", targetDatabaseId2)

	resourceName := "oci_data_safe_target_database_group.test_target_database_group"
	datasourceName := "data.oci_data_safe_target_database_groups.test_target_database_groups"
	singularDatasourceName := "data.oci_data_safe_target_database_group.test_target_database_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeTargetDatabaseGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_group", "test_target_database_group", acctest.Optional, acctest.Create, DataSafeTargetDatabaseGroupRepresentation), "datasafe", "targetDatabaseGroup", t)

	acctest.ResourceTest(t, testAccCheckDataSafeTargetDatabaseGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeTargetDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_group", "test_target_database_group", acctest.Required, acctest.Create, DataSafeTargetDatabaseGroupRepresentation) + targetDatabaseIdVariableStr + targetDatabaseIdVariableStr2,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestGroup"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeTargetDatabaseGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeTargetDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_group", "test_target_database_group", acctest.Optional, acctest.Create, DataSafeTargetDatabaseGroupRepresentation) + targetDatabaseIdVariableStr + targetDatabaseIdVariableStr2,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Description of the target database group (optional)."),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestGroup"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.exclude.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.exclude.0.target_database_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.compartments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.compartments.0.id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.compartments.0.is_include_subtree", "true"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.target_database_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "membership_count"),
				resource.TestCheckResourceAttrSet(resourceName, "membership_update_time"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeTargetDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_group", "test_target_database_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeTargetDatabaseGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})) + targetDatabaseIdVariableStr + targetDatabaseIdVariableStr2,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Description of the target database group (optional)."),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestGroup"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.exclude.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.exclude.0.target_database_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.compartments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.compartments.0.id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.compartments.0.is_include_subtree", "true"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.target_database_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "membership_count"),
				resource.TestCheckResourceAttrSet(resourceName, "membership_update_time"),
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
			Config: config + compartmentIdVariableStr + DataSafeTargetDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_group", "test_target_database_group", acctest.Optional, acctest.Update, DataSafeTargetDatabaseGroupRepresentation) +
				targetDatabaseIdVariableStr + targetDatabaseIdVariableStr2,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.exclude.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.exclude.0.target_database_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.compartments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.compartments.0.id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.compartments.0.is_include_subtree", "true"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "matching_criteria.0.include.0.target_database_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "membership_count"),
				resource.TestCheckResourceAttrSet(resourceName, "membership_update_time"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_database_groups", "test_target_database_groups", acctest.Optional, acctest.Update, DataSafeTargetDatabaseGroupDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeTargetDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_group", "test_target_database_group", acctest.Optional, acctest.Update, DataSafeTargetDatabaseGroupRepresentation) + targetDatabaseIdVariableStr + targetDatabaseIdVariableStr2,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_database_group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "target_database_group_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "target_database_group_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_database_group", "test_target_database_group", acctest.Required, acctest.Create, DataSafeTargetDatabaseGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeTargetDatabaseGroupResourceConfig + targetDatabaseIdVariableStr + targetDatabaseIdVariableStr2,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_database_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_criteria.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_criteria.0.exclude.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_criteria.0.exclude.0.target_database_ids.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_criteria.0.include.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_criteria.0.include.0.compartments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_criteria.0.include.0.compartments.0.id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_criteria.0.include.0.compartments.0.is_include_subtree", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_criteria.0.include.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "matching_criteria.0.include.0.target_database_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "membership_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "membership_update_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeTargetDatabaseGroupRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeTargetDatabaseGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_target_database_group" {
			noResourceFound = false
			request := oci_data_safe.GetTargetDatabaseGroupRequest{}

			tmp := rs.Primary.ID
			request.TargetDatabaseGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetTargetDatabaseGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.TargetDatabaseGroupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeTargetDatabaseGroup") {
		resource.AddTestSweepers("DataSafeTargetDatabaseGroup", &resource.Sweeper{
			Name:         "DataSafeTargetDatabaseGroup",
			Dependencies: acctest.DependencyGraph["targetDatabaseGroup"],
			F:            sweepDataSafeTargetDatabaseGroupResource,
		})
	}
}

func sweepDataSafeTargetDatabaseGroupResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	targetDatabaseGroupIds, err := getDataSafeTargetDatabaseGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, targetDatabaseGroupId := range targetDatabaseGroupIds {
		if ok := acctest.SweeperDefaultResourceId[targetDatabaseGroupId]; !ok {
			deleteTargetDatabaseGroupRequest := oci_data_safe.DeleteTargetDatabaseGroupRequest{}

			deleteTargetDatabaseGroupRequest.TargetDatabaseGroupId = &targetDatabaseGroupId

			deleteTargetDatabaseGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteTargetDatabaseGroup(context.Background(), deleteTargetDatabaseGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting TargetDatabaseGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", targetDatabaseGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &targetDatabaseGroupId, DataSafeTargetDatabaseGroupSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeTargetDatabaseGroupSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeTargetDatabaseGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TargetDatabaseGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listTargetDatabaseGroupsRequest := oci_data_safe.ListTargetDatabaseGroupsRequest{}
	listTargetDatabaseGroupsRequest.CompartmentId = &compartmentId
	listTargetDatabaseGroupsRequest.LifecycleState = oci_data_safe.ListTargetDatabaseGroupsLifecycleStateActive
	listTargetDatabaseGroupsResponse, err := dataSafeClient.ListTargetDatabaseGroups(context.Background(), listTargetDatabaseGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting TargetDatabaseGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, targetDatabaseGroup := range listTargetDatabaseGroupsResponse.Items {
		id := *targetDatabaseGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TargetDatabaseGroupId", id)
	}
	return resourceIds, nil
}

func DataSafeTargetDatabaseGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if targetDatabaseGroupResponse, ok := response.Response.(oci_data_safe.GetTargetDatabaseGroupResponse); ok {
		return targetDatabaseGroupResponse.LifecycleState != oci_data_safe.TargetDatabaseGroupLifecycleStateDeleted
	}
	return false
}

func DataSafeTargetDatabaseGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetTargetDatabaseGroup(context.Background(), oci_data_safe.GetTargetDatabaseGroupRequest{
		TargetDatabaseGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
