// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	DataSafeAlertPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_user_defined":           acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"alert_policy_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_alert_policy.test_alert_policy.id}`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeAlertPolicyDataSourceFilterRepresentation},
	}
	DataSafeAlertPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_alert_policy.test_alert_policy.id}`}},
	}

	DataSafeAlertPolicyRequiredOnlyResource = DataSafeAlertPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Required, acctest.Create, DataSafeAlertPolicyRepresentation)

	DataSafeAlertPolicySingularDataSourceRepresentation = map[string]interface{}{
		"alert_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_alert_policy.test_alert_policy.id}`},
	}

	DataSafeAlertPolicyRepresentation = map[string]interface{}{
		"alert_policy_type":         acctest.Representation{RepType: acctest.Required, Create: `AUDITING`},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"severity":                  acctest.Representation{RepType: acctest.Required, Create: `CRITICAL`, Update: `HIGH`},
		"alert_policy_rule_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataSafeAlertPolicyAlertPolicyRuleDetailsRepresentation},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `Check if remote login password file is exclusive and remote login is enabled `, Update: `description2`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `Check remote login`, Update: `displayName2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreAlertPolicyRep},
	}
	DataSafeAlertPolicyAlertPolicyRuleDetailsRepresentation = map[string]interface{}{
		"expression":   acctest.Representation{RepType: acctest.Required, Create: `operation eq \"abc\"`},
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `Check if remote login password file is exclusive and remote login is enabled `},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `Check remote login`},
	}

	DataSafeAlertPolicyResourceDependencies = DefinedTagsDependencies
	ignoreAlertPolicyRep                    = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`, `system_tags`}},
	}
	DataSafeAlertPolicyResourceConfig = DataSafeAlertPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Optional, acctest.Update, DataSafeAlertPolicyRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeAlertPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAlertPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_alert_policy.test_alert_policy"

	policyId := utils.GetEnvSettingWithBlankDefault("data_safe_alert_policy_ocid")
	policyIdVariableStr := fmt.Sprintf("variable \"policy_id\" { default = \"%s\" }\n", policyId)

	datasourceName := "data.oci_data_safe_alert_policies.test_alert_policies"
	singularDatasourceName := "data.oci_data_safe_alert_policy.test_alert_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeAlertPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Optional, acctest.Create, DataSafeAlertPolicyRepresentation), "datasafe", "alertPolicy", t)

	acctest.ResourceTest(t, testAccCheckDataSafeAlertPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeAlertPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Required, acctest.Create, DataSafeAlertPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alert_policy_type", "AUDITING"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "severity", "CRITICAL"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeAlertPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeAlertPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Optional, acctest.Create, DataSafeAlertPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.0.description", "Check if remote login password file is exclusive and remote login is enabled "),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.0.display_name", "Check remote login"),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.0.expression", "operation eq \"abc\""),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_type", "AUDITING"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Check if remote login password file is exclusive and remote login is enabled "),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Check remote login"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "severity", "CRITICAL"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + policyIdVariableStr + DataSafeAlertPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeAlertPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.0.description", "Check if remote login password file is exclusive and remote login is enabled "),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.0.display_name", "Check remote login"),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.0.expression", "operation eq \"abc\""),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_type", "AUDITING"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Check if remote login password file is exclusive and remote login is enabled "),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Check remote login"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "severity", "CRITICAL"),
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
			Config: config + compartmentIdVariableStr + DataSafeAlertPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Optional, acctest.Update, DataSafeAlertPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.0.description", "Check if remote login password file is exclusive and remote login is enabled "),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.0.display_name", "Check remote login"),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_rule_details.0.expression", "operation eq \"abc\""),
				resource.TestCheckResourceAttr(resourceName, "alert_policy_type", "AUDITING"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "severity", "HIGH"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert_policies", "test_alert_policies", acctest.Optional, acctest.Update, DataSafeAlertPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAlertPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Optional, acctest.Update, DataSafeAlertPolicyRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "alert_policy_collection.#", "1"),
			),
		},
		// verify singular datasource

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Required, acctest.Create, DataSafeAlertPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + policyIdVariableStr + DataSafeAlertPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_policy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_policy_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "severity"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DataSafeAlertPolicyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"alert_policy_rule_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataSafeAlertPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_alert_policy" {
			noResourceFound = false
			request := oci_data_safe.GetAlertPolicyRequest{}

			tmp := rs.Primary.ID
			request.AlertPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetAlertPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.AlertPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeAlertPolicy") {
		resource.AddTestSweepers("DataSafeAlertPolicy", &resource.Sweeper{
			Name:         "DataSafeAlertPolicy",
			Dependencies: acctest.DependencyGraph["alertPolicy"],
			F:            sweepDataSafeAlertPolicyResource,
		})
	}
}

func sweepDataSafeAlertPolicyResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	alertPolicyIds, err := getDataSafeAlertPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, alertPolicyId := range alertPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[alertPolicyId]; !ok {
			deleteAlertPolicyRequest := oci_data_safe.DeleteAlertPolicyRequest{}

			deleteAlertPolicyRequest.AlertPolicyId = &alertPolicyId

			deleteAlertPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteAlertPolicy(context.Background(), deleteAlertPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting AlertPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", alertPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &alertPolicyId, DataSafeAlertPolicySweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeAlertPolicySweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeAlertPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AlertPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	userDefined := true
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listAlertPoliciesRequest := oci_data_safe.ListAlertPoliciesRequest{}
	listAlertPoliciesRequest.CompartmentId = &compartmentId
	listAlertPoliciesRequest.LifecycleState = oci_data_safe.ListAlertPoliciesLifecycleStateActive
	listAlertPoliciesRequest.IsUserDefined = &userDefined
	listAlertPoliciesResponse, err := dataSafeClient.ListAlertPolicies(context.Background(), listAlertPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AlertPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, alertPolicy := range listAlertPoliciesResponse.Items {
		id := *alertPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AlertPolicyId", id)
	}
	return resourceIds, nil
}

func DataSafeAlertPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if alertPolicyResponse, ok := response.Response.(oci_data_safe.GetAlertPolicyResponse); ok {
		return alertPolicyResponse.LifecycleState != oci_data_safe.AlertPolicyLifecycleStateDeleted
	}
	return false
}

func DataSafeAlertPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetAlertPolicy(context.Background(), oci_data_safe.GetAlertPolicyRequest{
		AlertPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
