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
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RecoveryProtectionPolicyRequiredOnlyResource = RecoveryProtectionPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Required, acctest.Create, RecoveryProtectionPolicyRepresentation)

	RecoveryProtectionPolicyResourceConfig = RecoveryProtectionPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Optional, acctest.Update, RecoveryProtectionPolicyRepresentation)

	RecoveryprotectionPolicySingularDataSourceRepresentation = map[string]interface{}{
		"protection_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_recovery_protection_policy.test_protection_policy.id}`},
	}

	RecoveryprotectionPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"owner":                acctest.Representation{RepType: acctest.Optional, Create: `customer`},
		"protection_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_recovery_protection_policy.test_protection_policy.id}`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: RecoveryProtectionPolicyDataSourceFilterRepresentation}}
	RecoveryProtectionPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_recovery_protection_policy.test_protection_policy.id}`}},
	}

	RecoveryProtectionPolicyRepresentation = map[string]interface{}{
		"backup_retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: recoveryIgnoreDefinedTagsRepresentation},
	}

	RecoveryProtectionPolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: recovery/default
func TestRecoveryProtectionPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRecoveryProtectionPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_recovery_protection_policy.test_protection_policy"
	datasourceName := "data.oci_recovery_protection_policies.test_protection_policies"
	singularDatasourceName := "data.oci_recovery_protection_policy.test_protection_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RecoveryProtectionPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Optional, acctest.Create, RecoveryProtectionPolicyRepresentation), "recovery", "protectionPolicy", t)

	acctest.ResourceTest(t, testAccCheckRecoveryProtectionPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RecoveryProtectionPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Required, acctest.Create, RecoveryProtectionPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_retention_period_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RecoveryProtectionPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RecoveryProtectionPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Optional, acctest.Create, RecoveryProtectionPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_retention_period_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_predefined_policy"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RecoveryProtectionPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(RecoveryProtectionPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_retention_period_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_predefined_policy"),

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
			Config: config + compartmentIdVariableStr + RecoveryProtectionPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Optional, acctest.Update, RecoveryProtectionPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_retention_period_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_predefined_policy"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_protection_policies", "test_protection_policies", acctest.Optional, acctest.Update, RecoveryprotectionPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryProtectionPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Optional, acctest.Update, RecoveryProtectionPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "owner", "customer"),
				resource.TestCheckResourceAttrSet(datasourceName, "protection_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "protection_policy_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "protection_policy_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_protection_policy", "test_protection_policy", acctest.Required, acctest.Create, RecoveryprotectionPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryProtectionPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "protection_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "backup_retention_period_in_days", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_predefined_policy"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + RecoveryProtectionPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckRecoveryProtectionPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseRecoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_recovery_protection_policy" {
			noResourceFound = false
			request := oci_recovery.GetProtectionPolicyRequest{}

			tmp := rs.Primary.ID
			request.ProtectionPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "recovery")

			response, err := client.GetProtectionPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_recovery.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("RecoveryProtectionPolicy") {
		resource.AddTestSweepers("RecoveryProtectionPolicy", &resource.Sweeper{
			Name:         "RecoveryProtectionPolicy",
			Dependencies: acctest.DependencyGraph["protectionPolicy"],
			F:            sweepRecoveryProtectionPolicyResource,
		})
	}
}

func sweepRecoveryProtectionPolicyResource(compartment string) error {
	databaseRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseRecoveryClient()
	protectionPolicyIds, err := getRecoveryProtectionPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, protectionPolicyId := range protectionPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[protectionPolicyId]; !ok {
			deleteProtectionPolicyRequest := oci_recovery.DeleteProtectionPolicyRequest{}

			deleteProtectionPolicyRequest.ProtectionPolicyId = &protectionPolicyId

			deleteProtectionPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "recovery")
			_, error := databaseRecoveryClient.DeleteProtectionPolicy(context.Background(), deleteProtectionPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting ProtectionPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", protectionPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &protectionPolicyId, RecoveryProtectionPolicySweepWaitCondition, time.Duration(3*time.Minute),
				RecoveryProtectionPolicySweepResponseFetchOperation, "recovery", true)
		}
	}
	return nil
}

func getRecoveryProtectionPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProtectionPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseRecoveryClient()

	listProtectionPoliciesRequest := oci_recovery.ListProtectionPoliciesRequest{}
	listProtectionPoliciesRequest.CompartmentId = &compartmentId
	listProtectionPoliciesRequest.LifecycleState = oci_recovery.ListProtectionPoliciesLifecycleStateActive
	listProtectionPoliciesResponse, err := databaseRecoveryClient.ListProtectionPolicies(context.Background(), listProtectionPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ProtectionPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, protectionPolicy := range listProtectionPoliciesResponse.Items {
		id := *protectionPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProtectionPolicyId", id)
	}
	return resourceIds, nil
}

func RecoveryProtectionPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if protectionPolicyResponse, ok := response.Response.(oci_recovery.GetProtectionPolicyResponse); ok {
		return protectionPolicyResponse.LifecycleState != oci_recovery.LifecycleStateDeleted
	}
	return false
}

func RecoveryProtectionPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseRecoveryClient().GetProtectionPolicy(context.Background(), oci_recovery.GetProtectionPolicyRequest{
		ProtectionPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
