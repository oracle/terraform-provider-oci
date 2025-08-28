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
	DataSafeSecurityPolicyDeploymentRequiredOnlyResource = DataSafeSecurityPolicyDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Required, acctest.Create, DataSafeSecurityPolicyDeploymentRepresentation)

	DataSafeSecurityPolicyDeploymentResourceConfig = DataSafeSecurityPolicyDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentRepresentation)

	DataSafeSecurityPolicyDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"security_policy_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_policy_deployment.test_security_policy_deployment.id}`},
	}

	DataSafeSecurityPolicyDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName2`},
	}

	DataSafeSecurityPolicyDeploymentDeployRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"security_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_id}`},
		"target_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"target_type":        acctest.Representation{RepType: acctest.Required, Create: `TARGET_DATABASE`},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreSecurityPolicyDeploymentSystemTagsChangesRep},
		"deploy_trigger":     acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	DataSafeSecurityPolicyDeploymentRefreshRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"security_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_id}`},
		"target_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"target_type":        acctest.Representation{RepType: acctest.Required, Create: `TARGET_DATABASE`},
		"deploy_trigger":     acctest.Representation{RepType: acctest.Optional, Create: '1', Update: `1`},
		"refresh_trigger":    acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	DataSafeSecurityPolicyDeploymentRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"security_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_id}`},
		"target_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"target_type":        acctest.Representation{RepType: acctest.Required, Create: `TARGET_DATABASE`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreSecurityPolicyDeploymentSystemTagsChangesRep},
	}

	IgnoreSecurityPolicyDeploymentSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSecurityPolicyDeploymentResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyDeploymentResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the target ocid & security policy ocid are hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSecurityPolicyDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	securityPolicyId := utils.GetEnvSettingWithBlankDefault("security_policy_ocid")
	securityPolicyIdVariableStr := fmt.Sprintf("variable \"security_policy_id\" { default = \"%s\" }\n", securityPolicyId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_security_policy_deployment.test_security_policy_deployment"
	datasourceName := "data.oci_data_safe_security_policy_deployments.test_security_policy_deployments"
	singularDatasourceName := "data.oci_data_safe_security_policy_deployment.test_security_policy_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.

	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeSecurityPolicyDeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Create, DataSafeSecurityPolicyDeploymentRepresentation), "datasafe", "securityPolicyDeployment", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSecurityPolicyDeploymentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Required, acctest.Create, DataSafeSecurityPolicyDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//verify Delete
		{
			Config: config + compartmentIdVariableStr + DataSafeSecurityPolicyDeploymentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Create, DataSafeSecurityPolicyDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + securityPolicyIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyDeploymentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE"),
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
			Config: config + compartmentIdVariableStr + DataSafeSecurityPolicyDeploymentResourceDependencies + securityPolicyIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
		// verify deploy
		{
			Config: config + compartmentIdVariableStr + DataSafeSecurityPolicyDeploymentResourceDependencies + securityPolicyIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentDeployRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be deployed.")
					}
					return err
				},
			),
		},
		// verify refresh
		{
			Config: config + compartmentIdVariableStr + DataSafeSecurityPolicyDeploymentResourceDependencies + securityPolicyIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentRefreshRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be refreshed.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_deployments", "test_security_policy_deployments", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "security_policy_deployment_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Required, acctest.Create, DataSafeSecurityPolicyDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyDeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_deployment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + DataSafeSecurityPolicyDeploymentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`security_policy_deployment_id`},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSecurityPolicyDeploymentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_security_policy_deployment" {
			noResourceFound = false
			request := oci_data_safe.GetSecurityPolicyDeploymentRequest{}

			tmp := rs.Primary.ID
			request.SecurityPolicyDeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSecurityPolicyDeployment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SecurityPolicyDeploymentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeSecurityPolicyDeployment") {
		resource.AddTestSweepers("DataSafeSecurityPolicyDeployment", &resource.Sweeper{
			Name:         "DataSafeSecurityPolicyDeployment",
			Dependencies: acctest.DependencyGraph["securityPolicyDeployment"],
			F:            sweepDataSafeSecurityPolicyDeploymentResource,
		})
	}
}

func sweepDataSafeSecurityPolicyDeploymentResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	securityPolicyDeploymentIds, err := getDataSafeSecurityPolicyDeploymentIds(compartment)
	if err != nil {
		return err
	}
	for _, securityPolicyDeploymentId := range securityPolicyDeploymentIds {
		if ok := acctest.SweeperDefaultResourceId[securityPolicyDeploymentId]; !ok {
			deleteSecurityPolicyDeploymentRequest := oci_data_safe.DeleteSecurityPolicyDeploymentRequest{}

			deleteSecurityPolicyDeploymentRequest.SecurityPolicyDeploymentId = &securityPolicyDeploymentId

			deleteSecurityPolicyDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSecurityPolicyDeployment(context.Background(), deleteSecurityPolicyDeploymentRequest)
			if error != nil {
				fmt.Printf("Error deleting SecurityPolicyDeployment %s %s, It is possible that the resource is already deleted. Please verify manually \n", securityPolicyDeploymentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &securityPolicyDeploymentId, DataSafeSecurityPolicyDeploymentSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSecurityPolicyDeploymentSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSecurityPolicyDeploymentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SecurityPolicyDeploymentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSecurityPolicyDeploymentsRequest := oci_data_safe.ListSecurityPolicyDeploymentsRequest{}
	listSecurityPolicyDeploymentsRequest.CompartmentId = &compartmentId
	listSecurityPolicyDeploymentsRequest.LifecycleState = oci_data_safe.ListSecurityPolicyDeploymentsLifecycleStateNeedsAttention
	listSecurityPolicyDeploymentsResponse, err := dataSafeClient.ListSecurityPolicyDeployments(context.Background(), listSecurityPolicyDeploymentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SecurityPolicyDeployment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, securityPolicyDeployment := range listSecurityPolicyDeploymentsResponse.Items {
		id := *securityPolicyDeployment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SecurityPolicyDeploymentId", id)
	}
	return resourceIds, nil
}

func DataSafeSecurityPolicyDeploymentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if securityPolicyDeploymentResponse, ok := response.Response.(oci_data_safe.GetSecurityPolicyDeploymentResponse); ok {
		return securityPolicyDeploymentResponse.LifecycleState != oci_data_safe.SecurityPolicyDeploymentLifecycleStateDeleted
	}
	return false
}

func DataSafeSecurityPolicyDeploymentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSecurityPolicyDeployment(context.Background(), oci_data_safe.GetSecurityPolicyDeploymentRequest{
		SecurityPolicyDeploymentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
