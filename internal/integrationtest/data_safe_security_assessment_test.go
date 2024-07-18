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
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafeSecurityAssessmentRequiredOnlyResource = DataSafeSecurityAssessmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", acctest.Required, acctest.Create, securityAssessmentRepresentation)

	SecurityAssessmentResourceConfig = DataSafeSecurityAssessmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", acctest.Optional, acctest.Update, securityAssessmentRepresentation)

	DataSafesecurityAssessmentSingularDataSourceRepresentation = map[string]interface{}{
		"security_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_assessment.test_security_assessment.id}`},
	}

	securityAssessmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `SAVED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: securityAssessmentDataSourceFilterRepresentation}}
	securityAssessmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_security_assessment.test_security_assessment.id}`}},
	}

	securityAssessmentRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `EBS assessment`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesSecurityAssessmentRepresentation},
	}

	securityAssessmentChangeCompartmentRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `EBS assessment`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"schedule":       acctest.Representation{RepType: acctest.Optional, Create: `schedule`, Update: `schedule2`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesSecurityAssessmentRepresentation},
	}

	ignoreChangesSecurityAssessmentRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DataSafeSecurityAssessmentResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	autonomousDatabaseId := utils.GetEnvSettingWithBlankDefault("autonomous_database_id")
	autonomousDbIdVariableStr := fmt.Sprintf("variable \"autonomous_db_id\" { default = \"%s\" }\n", autonomousDatabaseId)

	resourceName := "oci_data_safe_security_assessment.test_security_assessment"
	datasourceName := "data.oci_data_safe_security_assessments.test_security_assessments"
	singularDatasourceName := "data.oci_data_safe_security_assessment.test_security_assessment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeSecurityAssessmentResourceDependencies+targetIdVariableStr+autonomousDbIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", acctest.Optional, acctest.Create, securityAssessmentRepresentation), "datasafe", "securityAssessment", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSecurityAssessmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSecurityAssessmentResourceDependencies + targetIdVariableStr + autonomousDbIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", acctest.Required, acctest.Create, securityAssessmentRepresentation),
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
			Config: config + compartmentIdVariableStr + DataSafeSecurityAssessmentResourceDependencies + targetIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeSecurityAssessmentResourceDependencies + targetIdVariableStr + autonomousDbIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", acctest.Optional, acctest.Create, securityAssessmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EBS assessment"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_assessment_scheduled", "false"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeSecurityAssessmentResourceDependencies + targetIdVariableStr + autonomousDbIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(securityAssessmentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EBS assessment"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_assessment_scheduled", "true"),
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
			Config: config + compartmentIdVariableStr + DataSafeSecurityAssessmentResourceDependencies + targetIdVariableStr + autonomousDbIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", acctest.Required, acctest.Update, securityAssessmentChangeCompartmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_assessment_scheduled", "true"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessments", "test_security_assessments", acctest.Optional, acctest.Update, securityAssessmentDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSecurityAssessmentResourceDependencies + targetIdVariableStr + autonomousDbIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", acctest.Optional, acctest.Update, securityAssessmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttr(datasourceName, "type", "SAVED"),

				resource.TestCheckResourceAttr(datasourceName, "security_assessments.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_assessments.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "security_assessments.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.is_baseline"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "security_assessments.0.statistics.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.time_last_assessed"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.triggered_by"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessments.0.type"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment", acctest.Required, acctest.Create, DataSafesecurityAssessmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSecurityAssessmentRequiredOnlyResource + targetIdVariableStr + autonomousDbIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_assessment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_assessment_scheduled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_baseline"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "statistics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "statistics.0.pass.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "statistics.0.deferred.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_assessed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "triggered_by"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + SecurityAssessmentResourceConfig + targetIdVariableStr + autonomousDbIdVariableStr,
		},
		// verify resource import
		{
			Config:            config + DataSafeSecurityAssessmentRequiredOnlyResource + targetIdVariableStr + autonomousDbIdVariableStr,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"target_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataSafeSecurityAssessmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_security_assessment" {
			noResourceFound = false
			request := oci_data_safe.GetSecurityAssessmentRequest{}

			tmp := rs.Primary.ID
			request.SecurityAssessmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSecurityAssessment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SecurityAssessmentLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is expected to be in deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
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
	if !acctest.InSweeperExcludeList("DataSafeSecurityAssessment") {
		resource.AddTestSweepers("DataSafeSecurityAssessment", &resource.Sweeper{
			Name:         "DataSafeSecurityAssessment",
			Dependencies: acctest.DependencyGraph["securityassessment"],
			F:            sweepDataSafeSecurityAssessmentResource,
		})
	}
}

func sweepDataSafeSecurityAssessmentResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	securityAssessmentIds, err := getDataSafeSecurityAssessmentIds(compartment)
	if err != nil {
		return err
	}
	for _, securityAssessmentId := range securityAssessmentIds {
		if ok := acctest.SweeperDefaultResourceId[securityAssessmentId]; !ok {
			deleteSecurityAssessmentRequest := oci_data_safe.DeleteSecurityAssessmentRequest{}

			deleteSecurityAssessmentRequest.SecurityAssessmentId = &securityAssessmentId

			deleteSecurityAssessmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSecurityAssessment(context.Background(), deleteSecurityAssessmentRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityAssessment %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityAssessmentId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeSecurityAssessmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SecurityAssessmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSecurityAssessmentsRequest := oci_data_safe.ListSecurityAssessmentsRequest{}
	listSecurityAssessmentsRequest.CompartmentId = &compartmentId
	listSecurityAssessmentsResponse, err := dataSafeClient.ListSecurityAssessments(context.Background(), listSecurityAssessmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityAssessment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityAssessment := range listSecurityAssessmentsResponse.Items {
		id := *securityAssessment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityAssessmentId", id)
	}
	return resourceIds, nil
}
