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
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DelegateAccessControlDelegationSubscriptionRequiredOnlyResource = DelegateAccessControlDelegationSubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_subscription", "test_delegation_subscription", acctest.Required, acctest.Create, DelegateAccessControlDelegationSubscriptionRepresentation)

	DelegateAccessControlDelegationSubscriptionResourceConfig = DelegateAccessControlDelegationSubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_subscription", "test_delegation_subscription", acctest.Optional, acctest.Update, DelegateAccessControlDelegationSubscriptionRepresentation)

	DelegateAccessControlDelegationSubscriptionSingularDataSourceRepresentation = map[string]interface{}{
		"delegation_subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_delegate_access_control_delegation_subscription.test_delegation_subscription.id}`},
	}

	DelegateAccessControlDelegationSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DelegateAccessControlDelegationSubscriptionDataSourceFilterRepresentation}}
	DelegateAccessControlDelegationSubscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_delegate_access_control_delegation_subscription.test_delegation_subscription.id}`}},
	}

	DelegateAccessControlDelegationSubscriptionRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"service_provider_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.svcProviderId}`},
		"subscribed_service_type": acctest.Representation{RepType: acctest.Required, Create: `TROUBLESHOOTING`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DelegateAccessControlDelegationSubscriptionResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_service_providers", "test_service_providers", acctest.Required, acctest.Create, DelegateAccessControlServiceProviderDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: delegate_access_control/default
func TestDelegateAccessControlDelegationSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDelegateAccessControlDelegationSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_delegate_access_control_delegation_subscription.test_delegation_subscription"
	datasourceName := "data.oci_delegate_access_control_delegation_subscriptions.test_delegation_subscriptions"
	singularDatasourceName := "data.oci_delegate_access_control_delegation_subscription.test_delegation_subscription"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DelegateAccessControlDelegationSubscriptionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_subscription", "test_delegation_subscription", acctest.Optional, acctest.Create, DelegateAccessControlDelegationSubscriptionRepresentation), "delegateaccesscontrol", "delegationSubscription", t)

	acctest.ResourceTest(t, testAccCheckDelegateAccessControlDelegationSubscriptionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DelegateAccessControlDelegationSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_subscription", "test_delegation_subscription", acctest.Required, acctest.Create, DelegateAccessControlDelegationSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "service_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "subscribed_service_type", "TROUBLESHOOTING"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DelegateAccessControlDelegationSubscriptionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DelegateAccessControlDelegationSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_subscription", "test_delegation_subscription", acctest.Optional, acctest.Create, DelegateAccessControlDelegationSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "subscribed_service_type", "TROUBLESHOOTING"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DelegateAccessControlDelegationSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_subscription", "test_delegation_subscription", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DelegateAccessControlDelegationSubscriptionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "subscribed_service_type", "TROUBLESHOOTING"),

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
			Config: config + compartmentIdVariableStr + DelegateAccessControlDelegationSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_subscription", "test_delegation_subscription", acctest.Optional, acctest.Update, DelegateAccessControlDelegationSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "subscribed_service_type", "TROUBLESHOOTING"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegation_subscriptions", "test_delegation_subscriptions", acctest.Optional, acctest.Update, DelegateAccessControlDelegationSubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlDelegationSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_subscription", "test_delegation_subscription", acctest.Optional, acctest.Update, DelegateAccessControlDelegationSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "delegation_subscription_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "delegation_subscription_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegation_subscription", "test_delegation_subscription", acctest.Required, acctest.Create, DelegateAccessControlDelegationSubscriptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlDelegationSubscriptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "delegation_subscription_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subscribed_service_type", "TROUBLESHOOTING"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DelegateAccessControlDelegationSubscriptionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDelegateAccessControlDelegationSubscriptionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DelegateAccessControlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_delegate_access_control_delegation_subscription" {
			noResourceFound = false
			request := oci_delegate_access_control.GetDelegationSubscriptionRequest{}

			tmp := rs.Primary.ID
			request.DelegationSubscriptionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "delegate_access_control")

			response, err := client.GetDelegationSubscription(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_delegate_access_control.DelegationSubscriptionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DelegateAccessControlDelegationSubscription") {
		resource.AddTestSweepers("DelegateAccessControlDelegationSubscription", &resource.Sweeper{
			Name:         "DelegateAccessControlDelegationSubscription",
			Dependencies: acctest.DependencyGraph["delegationSubscription"],
			F:            sweepDelegateAccessControlDelegationSubscriptionResource,
		})
	}
}

func sweepDelegateAccessControlDelegationSubscriptionResource(compartment string) error {
	delegateAccessControlClient := acctest.GetTestClients(&schema.ResourceData{}).DelegateAccessControlClient()
	delegationSubscriptionIds, err := getDelegateAccessControlDelegationSubscriptionIds(compartment)
	if err != nil {
		return err
	}
	for _, delegationSubscriptionId := range delegationSubscriptionIds {
		if ok := acctest.SweeperDefaultResourceId[delegationSubscriptionId]; !ok {
			deleteDelegationSubscriptionRequest := oci_delegate_access_control.DeleteDelegationSubscriptionRequest{}

			deleteDelegationSubscriptionRequest.DelegationSubscriptionId = &delegationSubscriptionId

			deleteDelegationSubscriptionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "delegate_access_control")
			_, error := delegateAccessControlClient.DeleteDelegationSubscription(context.Background(), deleteDelegationSubscriptionRequest)
			if error != nil {
				fmt.Printf("Error deleting DelegationSubscription %s %s, It is possible that the resource is already deleted. Please verify manually \n", delegationSubscriptionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &delegationSubscriptionId, DelegateAccessControlDelegationSubscriptionSweepWaitCondition, time.Duration(3*time.Minute),
				DelegateAccessControlDelegationSubscriptionSweepResponseFetchOperation, "delegate_access_control", true)
		}
	}
	return nil
}

func getDelegateAccessControlDelegationSubscriptionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DelegationSubscriptionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	delegateAccessControlClient := acctest.GetTestClients(&schema.ResourceData{}).DelegateAccessControlClient()

	listDelegationSubscriptionsRequest := oci_delegate_access_control.ListDelegationSubscriptionsRequest{}
	listDelegationSubscriptionsRequest.CompartmentId = &compartmentId
	listDelegationSubscriptionsRequest.LifecycleState = oci_delegate_access_control.DelegationSubscriptionLifecycleStateActive
	listDelegationSubscriptionsResponse, err := delegateAccessControlClient.ListDelegationSubscriptions(context.Background(), listDelegationSubscriptionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DelegationSubscription list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, delegationSubscription := range listDelegationSubscriptionsResponse.Items {
		id := *delegationSubscription.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DelegationSubscriptionId", id)
	}
	return resourceIds, nil
}

func DelegateAccessControlDelegationSubscriptionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if delegationSubscriptionResponse, ok := response.Response.(oci_delegate_access_control.GetDelegationSubscriptionResponse); ok {
		return delegationSubscriptionResponse.LifecycleState != oci_delegate_access_control.DelegationSubscriptionLifecycleStateDeleted
	}
	return false
}

func DelegateAccessControlDelegationSubscriptionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DelegateAccessControlClient().GetDelegationSubscription(context.Background(), oci_delegate_access_control.GetDelegationSubscriptionRequest{
		DelegationSubscriptionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
