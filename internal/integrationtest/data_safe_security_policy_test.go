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
	DataSafeSecurityPolicyRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Required, acctest.Create, DataSafeSecurityPolicyRepresentation)

	DataSafeSecurityPolicyResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Update, DataSafeSecurityPolicyRepresentation)

	DataSafeSecurityPolicySingularDataSourceRepresentation = map[string]interface{}{
		"security_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_policy.test_security_policy.id}`},
	}

	DataSafeSecurityPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"security_policy_type":      acctest.Representation{RepType: acctest.Optional, Create: `DATASAFE_MANAGED`},
	}

	DataSafeSecurityPolicyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		//"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSecurityPolicyTagsChangesRep},
	}

	ignoreSecurityPolicyTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSecurityPolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_security_policy.test_security_policy"
	datasourceName := "data.oci_data_safe_security_policies.test_security_policies"
	singularDatasourceName := "data.oci_data_safe_security_policy.test_security_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Create, DataSafeSecurityPolicyRepresentation), "datasafe", "securityPolicy", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSecurityPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Required, acctest.Create, DataSafeSecurityPolicyRepresentation),
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
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Create, DataSafeSecurityPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Update, DataSafeSecurityPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policies", "test_security_policies", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Update, DataSafeSecurityPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "security_policy_type", "DATASAFE_MANAGED"),

				resource.TestCheckResourceAttr(datasourceName, "security_policy_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Required, acctest.Create, DataSafeSecurityPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSecurityPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Create, DataSafeSecurityPolicyRepresentation),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`security_policy_id`},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSecurityPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_security_policy" {
			noResourceFound = false
			request := oci_data_safe.GetSecurityPolicyRequest{}

			tmp := rs.Primary.ID
			request.SecurityPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSecurityPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SecurityPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeSecurityPolicy") {
		resource.AddTestSweepers("DataSafeSecurityPolicy", &resource.Sweeper{
			Name:         "DataSafeSecurityPolicy",
			Dependencies: acctest.DependencyGraph["securityPolicy"],
			F:            sweepDataSafeSecurityPolicyResource,
		})
	}
}

func sweepDataSafeSecurityPolicyResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	securityPolicyIds, err := getDataSafeSecurityPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, securityPolicyId := range securityPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[securityPolicyId]; !ok {
			deleteSecurityPolicyRequest := oci_data_safe.DeleteSecurityPolicyRequest{}

			deleteSecurityPolicyRequest.SecurityPolicyId = &securityPolicyId

			deleteSecurityPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSecurityPolicy(context.Background(), deleteSecurityPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &securityPolicyId, DataSafeSecurityPolicySweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSecurityPolicySweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSecurityPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SecurityPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSecurityPoliciesRequest := oci_data_safe.ListSecurityPoliciesRequest{}
	listSecurityPoliciesRequest.CompartmentId = &compartmentId
	listSecurityPoliciesRequest.LifecycleState = oci_data_safe.ListSecurityPoliciesLifecycleStateActive
	listSecurityPoliciesResponse, err := dataSafeClient.ListSecurityPolicies(context.Background(), listSecurityPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityPolicy := range listSecurityPoliciesResponse.Items {
		id := *securityPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityPolicyId", id)
	}
	return resourceIds, nil
}

func DataSafeSecurityPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if securityPolicyResponse, ok := response.Response.(oci_data_safe.GetSecurityPolicyResponse); ok {
		return securityPolicyResponse.LifecycleState != oci_data_safe.SecurityPolicyLifecycleStateDeleted
	}
	return false
}

func DataSafeSecurityPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSecurityPolicy(context.Background(), oci_data_safe.GetSecurityPolicyRequest{
		SecurityPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
