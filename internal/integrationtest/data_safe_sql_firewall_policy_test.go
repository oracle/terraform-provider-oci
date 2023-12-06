// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
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
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSqlFirewallPolicyRequiredOnlyResource = DataSafeSqlFirewallPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_firewall_policy", "test_sql_firewall_policy", acctest.Required, acctest.Create, DataSafeSqlFirewallPolicyRepresentation)

	DataSafeSqlFirewallPolicyResourceConfig = DataSafeSqlFirewallPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_firewall_policy", "test_sql_firewall_policy", acctest.Optional, acctest.Update, DataSafeSqlFirewallPolicyRepresentation)

	DataSafeSqlFirewallPolicySingularDataSourceRepresentation = map[string]interface{}{
		"sql_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_firewall_policy_id}`},
	}

	DataSafeSqlFirewallPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	DataSafeSqlFirewallPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_sql_firewall_policy.test_sql_firewall_policy.id}`}},
	}

	DataSafeSqlFirewallPolicyRepresentation = map[string]interface{}{
		"sql_firewall_policy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.sql_firewall_policy_id}`},
		"allowed_client_ips":          acctest.Representation{RepType: acctest.Optional, Create: []string{`10.239.175.91`}, Update: []string{`10.239.175.92`}},
		"allowed_client_os_usernames": acctest.Representation{RepType: acctest.Optional, Create: []string{`allowedClientOsUsernames`}, Update: []string{`allowedClientOsUsernames2`}},
		"allowed_client_programs":     acctest.Representation{RepType: acctest.Optional, Create: []string{`allowedClientPrograms`}, Update: []string{`allowedClientPrograms2`}},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `sample SQL firewall policy`, Update: `description2`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `samplePolicy`, Update: `displayName2`},
		"enforcement_scope":           acctest.Representation{RepType: acctest.Optional, Create: `ENFORCE_CONTEXT`, Update: `ENFORCE_SQL`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                      acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `ENABLED`},
		"violation_action":            acctest.Representation{RepType: acctest.Optional, Create: `BLOCK`, Update: `OBSERVE`},
		"violation_audit":             acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreSqlFirewallPolicySystemTagsChangesRep},
	}

	IgnoreSqlFirewallPolicySystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSqlFirewallPolicyListRepresentation = map[string]interface{}{
		"sql_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_firewall_policy_id}`},
		"status":                 acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`}, "description": acctest.Representation{RepType: acctest.Optional, Create: `updated-description`, Update: `description2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreSqlFirewallPolicySystemTagsChangesRep},
	}

	DataSafeSqlFirewallPolicyChangeCompartmentRepresentation = map[string]interface{}{
		"sql_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_firewall_policy_id}`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `sample SQL firewall policy`, Update: `description2`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `samplePolicy`, Update: `displayName2`},
		"lifecycle":              acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreSqlFirewallPolicySystemTagsChangesRep},
	}

	DataSafeSqlFirewallPolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSqlFirewallPolicyResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the policy ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSqlFirewallPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	fmt.Println("compartmentIdUVariableStr", compartmentIdUVariableStr)

	policyId := utils.GetEnvSettingWithBlankDefault("sql_firewall_policy_ocid")
	policyIdVariableStr := fmt.Sprintf("variable \"sql_firewall_policy_id\" { default = \"%s\" }\n", policyId)

	resourceName := "oci_data_safe_sql_firewall_policy.test_sql_firewall_policy"
	datasourceName := "data.oci_data_safe_sql_firewall_policies.test_sql_firewall_policies"
	singularDatasourceName := "data.oci_data_safe_sql_firewall_policy.test_sql_firewall_policy"

	var resId, resId2 string
	resId = policyId

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify change compartment
		{
			Config: config + policyIdVariableStr + compartmentIdUVariableStr + DataSafeSecurityPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_firewall_policy", "test_sql_firewall_policy", acctest.Optional, acctest.Update, DataSafeSqlFirewallPolicyChangeCompartmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + policyIdVariableStr + DataSafeSqlFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_firewall_policy", "test_sql_firewall_policy", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeSqlFirewallPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allowed_client_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_client_os_usernames.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "allowed_client_programs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_user_name"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "enforcement_scope", "ENFORCE_SQL"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "violation_action", "OBSERVE"),
				resource.TestCheckResourceAttr(resourceName, "violation_audit", "DISABLED"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_firewall_policies", "test_sql_firewall_policies", acctest.Optional, acctest.Update, DataSafeSqlFirewallPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + policyIdVariableStr + DataSafeSqlFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_firewall_policy", "test_sql_firewall_policy", acctest.Optional, acctest.Update, DataSafeSqlFirewallPolicyListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_firewall_policy", "test_sql_firewall_policy", acctest.Required, acctest.Create, DataSafeSqlFirewallPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + policyIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_firewall_policy_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_client_ips.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_client_os_usernames.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_client_programs.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "enforcement_scope", "ENFORCE_SQL"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_level"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "violation_action", "OBSERVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "violation_audit", "DISABLED"),
			),
		},
	})
}

func testAccCheckDataSafeSqlFirewallPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_sql_firewall_policy" {
			noResourceFound = false
			request := oci_data_safe.GetSqlFirewallPolicyRequest{}

			tmp := rs.Primary.ID
			request.SqlFirewallPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSqlFirewallPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SqlFirewallPolicyLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				// resource lifecycle state is in expected deleted lifecycle states. continue with next one.
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
	if !acctest.InSweeperExcludeList("DataSafeSqlFirewallPolicy") {
		resource.AddTestSweepers("DataSafeSqlFirewallPolicy", &resource.Sweeper{
			Name:         "DataSafeSqlFirewallPolicy",
			Dependencies: acctest.DependencyGraph["sqlFirewallPolicy"],
			F:            sweepDataSafeSqlFirewallPolicyResource,
		})
	}
}

func sweepDataSafeSqlFirewallPolicyResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	sqlFirewallPolicyIds, err := getDataSafeSqlFirewallPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, sqlFirewallPolicyId := range sqlFirewallPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[sqlFirewallPolicyId]; !ok {
			deleteSqlFirewallPolicyRequest := oci_data_safe.DeleteSqlFirewallPolicyRequest{}

			deleteSqlFirewallPolicyRequest.SqlFirewallPolicyId = &sqlFirewallPolicyId

			deleteSqlFirewallPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSqlFirewallPolicy(context.Background(), deleteSqlFirewallPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting SqlFirewallPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", sqlFirewallPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sqlFirewallPolicyId, DataSafeSqlFirewallPolicySweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSqlFirewallPolicySweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSqlFirewallPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SqlFirewallPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSqlFirewallPoliciesRequest := oci_data_safe.ListSqlFirewallPoliciesRequest{}
	listSqlFirewallPoliciesRequest.CompartmentId = &compartmentId
	listSqlFirewallPoliciesRequest.LifecycleState = oci_data_safe.ListSqlFirewallPoliciesLifecycleStateActive
	listSqlFirewallPoliciesResponse, err := dataSafeClient.ListSqlFirewallPolicies(context.Background(), listSqlFirewallPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SqlFirewallPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sqlFirewallPolicy := range listSqlFirewallPoliciesResponse.Items {
		id := *sqlFirewallPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SqlFirewallPolicyId", id)
	}
	return resourceIds, nil
}

func DataSafeSqlFirewallPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sqlFirewallPolicyResponse, ok := response.Response.(oci_data_safe.GetSqlFirewallPolicyResponse); ok {
		return sqlFirewallPolicyResponse.LifecycleState != oci_data_safe.SqlFirewallPolicyLifecycleStateDeleted
	}
	return false
}

func DataSafeSqlFirewallPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSqlFirewallPolicy(context.Background(), oci_data_safe.GetSqlFirewallPolicyRequest{
		SqlFirewallPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
