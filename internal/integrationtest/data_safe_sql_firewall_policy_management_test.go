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
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSqlFirewallPolicyManagementRepresentation = map[string]interface{}{
		//"sql_firewall_policy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.sql_firewall_policy_id}`},
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"db_user_name":                acctest.Representation{RepType: acctest.Required, Create: `${var.sql_firewall_policy_dbUserName}`},
		"allowed_client_ips":          acctest.Representation{RepType: acctest.Optional, Create: []string{`10.239.175.91`}, Update: []string{`10.239.175.92`}},
		"allowed_client_os_usernames": acctest.Representation{RepType: acctest.Optional, Create: []string{`allowedClientOsUsernames`}, Update: []string{`allowedClientOsUsernames2`}},
		"allowed_client_programs":     acctest.Representation{RepType: acctest.Optional, Create: []string{`allowedClientPrograms`}, Update: []string{`allowedClientPrograms2`}},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `sample SQL firewall policy`, Update: `description2`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `samplePolicy`, Update: `displayName2`},
		"enforcement_scope":           acctest.Representation{RepType: acctest.Optional, Create: `ENFORCE_CONTEXT`, Update: `ENFORCE_SQL`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                      acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"violation_action":            acctest.Representation{RepType: acctest.Optional, Create: `BLOCK`, Update: `OBSERVE`},
		"violation_audit":             acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreSqlFirewallPolicyManagementSystemTagsChangesRep},
	}

	IgnoreSqlFirewallPolicyManagementSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSqlFirewallPolicyManagementResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSqlFirewallPolicyManagementResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the target ocid, policy ocid and dbUserName are hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSqlFirewallPolicyManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	policyId := utils.GetEnvSettingWithBlankDefault("sql_firewall_policy_ocid")
	policyIdVariableStr := fmt.Sprintf("variable \"sql_firewall_policy_id\" { default = \"%s\" }\n", policyId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	dbUserName := utils.GetEnvSettingWithBlankDefault("sql_firewall_policy_user")
	dbUserNameVariableStr := fmt.Sprintf("variable \"sql_firewall_policy_dbUserName\" { default = \"%s\" }\n", dbUserName)
	fmt.Println("dbUserNameVariableStr", dbUserNameVariableStr)

	resourceName := "oci_data_safe_sql_firewall_policy_management.test_sql_firewall_policy_management"

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + policyIdVariableStr + dbUserNameVariableStr + targetIdVariableStr + DataSafeSqlFirewallPolicyManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sql_firewall_policy_management", "test_sql_firewall_policy_management", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeSqlFirewallPolicyManagementRepresentation, map[string]interface{}{
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
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "violation_action", "OBSERVE"),
				resource.TestCheckResourceAttr(resourceName, "violation_audit", "DISABLED"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataSafeSqlFirewallPolicyManagement") {
		resource.AddTestSweepers("DataSafeSqlFirewallPolicyManagement", &resource.Sweeper{
			Name:         "DataSafeSqlFirewallPolicyManagement",
			Dependencies: acctest.DependencyGraph["sqlFirewallPolicy"],
			F:            sweepDataSafeSqlFirewallPolicyManagementResource,
		})
	}
}

func sweepDataSafeSqlFirewallPolicyManagementResource(compartment string) error {
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &sqlFirewallPolicyId, DataSafeSqlFirewallPolicyManagementSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSqlFirewallPolicyManagementSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSqlFirewallPolicyManagementIds(compartment string) ([]string, error) {
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

func DataSafeSqlFirewallPolicyManagementSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sqlFirewallPolicyResponse, ok := response.Response.(oci_data_safe.GetSqlFirewallPolicyResponse); ok {
		return sqlFirewallPolicyResponse.LifecycleState != oci_data_safe.SqlFirewallPolicyLifecycleStateDeleted
	}
	return false
}

func DataSafeSqlFirewallPolicyManagementSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSqlFirewallPolicy(context.Background(), oci_data_safe.GetSqlFirewallPolicyRequest{
		SqlFirewallPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
