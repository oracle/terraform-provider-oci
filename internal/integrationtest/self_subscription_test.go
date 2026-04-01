// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_self "github.com/oracle/oci-go-sdk/v65/self"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

const dispName = "display_name21"
const updatedDispName = "updated_display_name5"

var (
	SelfSubscriptionRequiredOnlyResource = SelfSubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_self_subscription", "test_subscription", acctest.Required, acctest.Create, SelfSubscriptionRepresentation)

	SelfSubscriptionResourceConfig = SelfSubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_self_subscription", "test_subscription", acctest.Optional, acctest.Update, SelfSubscriptionRepresentation)

	SelfSubscriptionSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_self_subscription.test_subscription.id}`},
	}

	SelfSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: dispName, Update: updatedDispName},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_self_subscription.test_subscription.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: SelfSubscriptionDataSourceFilterRepresentation}}
	SelfSubscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_self_subscription.test_subscription.id}`}},
	}

	SelfSubscriptionRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"product_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.product_id}`},
		"seller_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.seller_id}`},
		"subscription_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: SelfSubscriptionSubscriptionDetailsRepresentation},
		"tenant_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.tenant_id}`},
		//"additional_details":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: SelfSubscriptionAdditionalDetailsRepresentation},
		//"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: dispName, Update: updatedDispName},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Accounting"}, Update: map[string]string{"Department": "Finance"}},
		//"realm":         acctest.Representation{RepType: acctest.Optional, Create: `OC1`},
		//"region":        acctest.Representation{RepType: acctest.Optional, Create: `us-ashburn-1`},
		//"source_type": acctest.Representation{RepType: acctest.Optional, Create: `OCI_NATIVE`},
	}
	SelfSubscriptionSubscriptionDetailsRepresentation = map[string]interface{}{
		"billing_details":          acctest.RepresentationGroup{RepType: acctest.Required, Group: SelfSubscriptionSubscriptionDetailsBillingDetailsRepresentation},
		"partner_registration_url": acctest.Representation{RepType: acctest.Required, Create: `https://oracle.com`},
		"pricing_plan":             acctest.RepresentationGroup{RepType: acctest.Required, Group: SelfSubscriptionSubscriptionDetailsPricingPlanRepresentation},
		"amount":                   acctest.Representation{RepType: acctest.Optional, Create: `5000`},
		"currency":                 acctest.Representation{RepType: acctest.Optional, Create: `USD`},
		"is_auto_renew":            acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	SelfSubscriptionAdditionalDetailsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}
	SelfSubscriptionSubscriptionDetailsBillingDetailsRepresentation = map[string]interface{}{
		"meters":          acctest.RepresentationGroup{RepType: acctest.Required, Group: SelfSubscriptionSubscriptionDetailsBillingDetailsMetersRepresentation},
		"metric_type":     acctest.Representation{RepType: acctest.Required, Create: `OCPU_HOURS`},
		"rate_allocation": acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"sku":             acctest.Representation{RepType: acctest.Required, Create: `MP00385`},
		"has_gov_sku":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	SelfSubscriptionSubscriptionDetailsPricingPlanRepresentation = map[string]interface{}{
		"billing_frequency": acctest.Representation{RepType: acctest.Required, Create: `YEARLY`},
		"plan_name":         acctest.Representation{RepType: acctest.Required, Create: `Base`},
		"plan_type":         acctest.Representation{RepType: acctest.Required, Create: `FIXED`},
		"rates":             acctest.RepresentationGroup{RepType: acctest.Required, Group: SelfSubscriptionSubscriptionDetailsPricingPlanRatesRepresentation},
		"plan_description":  acctest.Representation{RepType: acctest.Optional, Create: `planDescription`},
		"plan_duration":     acctest.Representation{RepType: acctest.Optional, Create: `ANNUAL`},
	}
	SelfSubscriptionSubscriptionDetailsBillingDetailsMetersRepresentation = map[string]interface{}{
		"name":              acctest.Representation{RepType: acctest.Required, Create: `MP_BOBO_OX`},
		"rate_allocation":   acctest.Representation{RepType: acctest.Required, Create: `1`},
		"extended_metadata": acctest.RepresentationGroup{RepType: acctest.Optional, Group: SelfSubscriptionSubscriptionDetailsBillingDetailsMetersExtendedMetadataRepresentation},
	}
	SelfSubscriptionSubscriptionDetailsPricingPlanRatesRepresentation = map[string]interface{}{
		"currency": acctest.Representation{RepType: acctest.Required, Create: `USD`},
		"rate":     acctest.Representation{RepType: acctest.Required, Create: `5000`},
	}
	SelfSubscriptionSubscriptionDetailsBillingDetailsMetersExtendedMetadataRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	//SelfSubscriptionResourceDependencies = DefinedTagsDependencies +
	//	acctest.GenerateDataSourceFromRepresentationMap("oci_self_products", "test_products", acctest.Required, acctest.Create, SelfProductSingularDataSourceRepresentation)

	SelfSubscriptionResourceDependencies = ""
)

// issue-routing-tag: self/default
func TestSelfSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSelfSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	//compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	productId := utils.GetEnvSettingWithBlankDefault("product_id")
	productIdVariableStr := fmt.Sprintf("variable \"product_id\" { default = \"%s\" }\n", productId)

	sellerId := utils.GetEnvSettingWithBlankDefault("seller_id")
	sellerIdVariableStr := fmt.Sprintf("variable \"seller_id\" { default = \"%s\" }\n", sellerId)

	//var resId, resId2 string
	tenantId := utils.GetEnvSettingWithBlankDefault("tenant_id")
	tenantIdVariableStr := fmt.Sprintf("variable \"tenant_id\" { default = \"%s\" }\n", tenantId)

	resourceName := "oci_self_subscription.test_subscription"
	//datasourceName := "data.oci_self_subscriptions.test_subscriptions"
	//singularDatasourceName := "data.oci_self_subscription.test_subscription"

	//var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SelfSubscriptionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_self_subscription", "test_subscription", acctest.Optional, acctest.Create, SelfSubscriptionRepresentation), "self", "subscription", t)

	acctest.ResourceTest(t, testAccCheckSelfSubscriptionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SelfSubscriptionResourceDependencies + productIdVariableStr + sellerIdVariableStr + tenantIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_self_subscription", "test_subscription", acctest.Required, acctest.Create, SelfSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "product_id"),
				resource.TestCheckResourceAttrSet(resourceName, "seller_id"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.name", "MP_BOBO_OX"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.rate_allocation", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.metric_type", "OCPU_HOURS"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.rate_allocation", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.sku", "MP00385"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.partner_registration_url", "https://oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.billing_frequency", "YEARLY"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_name", "Base"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_type", "FIXED"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.0.currency", "USD"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.0.rate", "5000"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
			ExpectNonEmptyPlan: true,
		},

		////delete before next Create
		//{
		//	Config: config + compartmentIdVariableStr + SelfSubscriptionResourceDependencies + productIdVariableStr + sellerIdVariableStr + tenantIdVariableStr,
		//},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SelfSubscriptionResourceDependencies + productIdVariableStr + sellerIdVariableStr + tenantIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_self_subscription", "test_subscription", acctest.Optional, acctest.Create, SelfSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "additional_details.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "additional_details.0.key", "key"),
				//resource.TestCheckResourceAttr(resourceName, "additional_details.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", dispName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "product_id"),
				//resource.TestCheckResourceAttr(resourceName, "realm", "realm"),
				//resource.TestCheckResourceAttr(resourceName, "region", "region"),
				resource.TestCheckResourceAttrSet(resourceName, "seller_id"),
				//resource.TestCheckResourceAttr(resourceName, "source_type", "OCI_NATIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.amount", "5000"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.has_gov_sku", "false"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.name", "MP_BOBO_OX"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.rate_allocation", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.metric_type", "OCPU_HOURS"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.rate_allocation", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.sku", "MP00385"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.currency", "USD"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.is_auto_renew", "false"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.partner_registration_url", "https://oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.billing_frequency", "YEARLY"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_description", "planDescription"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_duration", "ANNUAL"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_name", "Base"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_type", "FIXED"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.0.currency", "USD"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.0.rate", "5000"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
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
			ExpectNonEmptyPlan: true,
		},
	})
}

func testAccCheckSelfSubscriptionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).SelfSubscriptionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_self_subscription" {
			noResourceFound = false
			request := oci_self.GetSubscriptionRequest{}

			tmp := rs.Primary.ID
			request.SubscriptionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "self")

			response, err := client.GetSubscription(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_self.LifecycleStateEnumDeleted): true,
					string(oci_self.LifecycleStateEnumActive):  true,
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
	if !acctest.InSweeperExcludeList("SelfSubscription") {
		resource.AddTestSweepers("SelfSubscription", &resource.Sweeper{
			Name:         "SelfSubscription",
			Dependencies: acctest.DependencyGraph["subscription"],
			F:            sweepSelfSubscriptionResource,
		})
	}
}

func sweepSelfSubscriptionResource(compartment string) error {
	subscriptionClient := acctest.GetTestClients(&schema.ResourceData{}).SelfSubscriptionClient()
	subscriptionIds, err := getSelfSubscriptionIds(compartment)
	if err != nil {
		return err
	}
	for _, subscriptionId := range subscriptionIds {
		if ok := acctest.SweeperDefaultResourceId[subscriptionId]; !ok {
			deleteSubscriptionRequest := oci_self.DeleteSubscriptionRequest{}

			deleteSubscriptionRequest.SubscriptionId = &subscriptionId

			deleteSubscriptionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "self")
			_, error := subscriptionClient.DeleteSubscription(context.Background(), deleteSubscriptionRequest)
			if error != nil {
				fmt.Printf("Error deleting Subscription %s %s, It is possible that the resource is already deleted. Please verify manually \n", subscriptionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &subscriptionId, SelfSubscriptionSweepWaitCondition, time.Duration(3*time.Minute),
				SelfSubscriptionSweepResponseFetchOperation, "self", true)
		}
	}
	return nil
}

func getSelfSubscriptionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubscriptionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	subscriptionClient := acctest.GetTestClients(&schema.ResourceData{}).SelfSubscriptionClient()

	listSubscriptionsRequest := oci_self.ListSubscriptionsRequest{}
	listSubscriptionsRequest.CompartmentId = &compartmentId
	listSubscriptionsResponse, err := subscriptionClient.ListSubscriptions(context.Background(), listSubscriptionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Subscription list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subscription := range listSubscriptionsResponse.Items {
		id := *subscription.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubscriptionId", id)
	}
	return resourceIds, nil
}

func SelfSubscriptionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if subscriptionResponse, ok := response.Response.(oci_self.GetSubscriptionResponse); ok {
		return subscriptionResponse.LifecycleState != oci_self.LifecycleStateEnumDeleted
	}
	return false
}

func SelfSubscriptionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.SelfSubscriptionClient().GetSubscription(context.Background(), oci_self.GetSubscriptionRequest{
		SubscriptionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
