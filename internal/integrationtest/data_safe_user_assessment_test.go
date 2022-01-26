// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v56/datasafe"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	UserAssessmentRequiredOnlyResource = UserAssessmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, userAssessmentRepresentation)

	UserAssessmentResourceConfig = UserAssessmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Optional, acctest.Update, userAssessmentRepresentation)

	userAssessmentSingularDataSourceRepresentation = map[string]interface{}{
		"user_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
	}

	userAssessmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_database.test_target_database.id}`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `SAVED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: userAssessmentDataSourceFilterRepresentation}}
	userAssessmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_user_assessment.test_user_assessment.id}`}},
	}

	userAssessmentRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_database.test_target_database.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesUserAssessmentRepresentation},
	}

	userAssessmentChangeCompartmentRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_database.test_target_database.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"schedule":       acctest.Representation{RepType: acctest.Optional, Create: `v1; 00 30 15 * *`, Update: `v1; 00 30 20 * *`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesUserAssessmentRepresentation},
	}

	ignoreChangesUserAssessmentRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	UserAssessmentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Required, acctest.Create, targetDatabaseRepresentation) +
		DefinedTagsDependencies
)

//issue-routing-tag: data_safe/default
func TestDataSafeUserAssessmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUserAssessmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_user_assessment.test_user_assessment"
	datasourceName := "data.oci_data_safe_user_assessments.test_user_assessments"
	singularDatasourceName := "data.oci_data_safe_user_assessment.test_user_assessment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+UserAssessmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Optional, acctest.Create, userAssessmentRepresentation), "datasafe", "userassessment", t)

	acctest.ResourceTest(t, testAccCheckDataSafeUserAssessmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + UserAssessmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, userAssessmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + UserAssessmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + UserAssessmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Optional, acctest.Create, userAssessmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + UserAssessmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(userAssessmentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

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
			Config: config + compartmentIdVariableStr + UserAssessmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Update, userAssessmentChangeCompartmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessments", "test_user_assessments", acctest.Optional, acctest.Update, userAssessmentDataSourceRepresentation) +
				compartmentIdVariableStr + UserAssessmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Optional, acctest.Update, userAssessmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttr(datasourceName, "type", "SAVED"),

				resource.TestCheckResourceAttr(datasourceName, "user_assessments.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "user_assessments.0.compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "user_assessments.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessments.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessments.0.is_baseline"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessments.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessments.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessments.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessments.0.triggered_by"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessments.0.type"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, userAssessmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + UserAssessmentRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_assessment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_baseline"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "triggered_by"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + UserAssessmentResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"target_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataSafeUserAssessmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_user_assessment" {
			noResourceFound = false
			request := oci_data_safe.GetUserAssessmentRequest{}

			tmp := rs.Primary.ID
			request.UserAssessmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetUserAssessment(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataSafeUserAssessment") {
		resource.AddTestSweepers("DataSafeUserAssessment", &resource.Sweeper{
			Name:         "DataSafeUserAssessment",
			Dependencies: acctest.DependencyGraph["userassessment"],
			F:            sweepDataSafeUserAssessmentResource,
		})
	}
}

func sweepDataSafeUserAssessmentResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	userAssessmentIds, err := getUserAssessmentIds(compartment)
	if err != nil {
		return err
	}
	for _, userAssessmentId := range userAssessmentIds {
		if ok := acctest.SweeperDefaultResourceId[userAssessmentId]; !ok {
			deleteUserAssessmentRequest := oci_data_safe.DeleteUserAssessmentRequest{}

			deleteUserAssessmentRequest.UserAssessmentId = &userAssessmentId

			deleteUserAssessmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteUserAssessment(context.Background(), deleteUserAssessmentRequest)
			if error != nil {
				fmt.Printf("Error deleting UserAssessment %s %s, It is possible that the resource is already deleted. Please verify manually \n", userAssessmentId, error)
				continue
			}
		}
	}
	return nil
}

func getUserAssessmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "UserAssessmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listUserAssessmentsRequest := oci_data_safe.ListUserAssessmentsRequest{}
	listUserAssessmentsRequest.CompartmentId = &compartmentId
	listUserAssessmentsResponse, err := dataSafeClient.ListUserAssessments(context.Background(), listUserAssessmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting UserAssessment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, userAssessment := range listUserAssessmentsResponse.Items {
		id := *userAssessment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "UserAssessmentId", id)
	}
	return resourceIds, nil
}
