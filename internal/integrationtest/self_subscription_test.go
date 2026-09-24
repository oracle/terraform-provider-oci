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

const selfSubscriptionAcceptanceResource = `
resource "oci_self_subscription" "test_subscription" {
  compartment_id = var.compartment_id
  tenant_id      = var.tenant_id
  seller_id      = var.seller_id
  product_id     = var.product_id
  display_name   = "S-Open-AI-Test-Listing-${formatdate("YYYYMMDDhhmmss", timestamp())}"
  source_type    = "THIRD_PARTY"
  realm          = "OC1"
  region         = "us-ashburn-1"

  subscription_details {
    amount        = 89
    currency      = "USD"
    is_auto_renew = true
    partner_registration_url = "https://www.google.com/"

    billing_details {
      pricing_plan_key = "16ff7edf-a17f-4f77-88b8-1a09bef6df8f"
      billing_model = "FLAT_RATE"
      sku = "MP10257"
      metric_type = "INSTANCE_HOURS"
      rate_allocation = 1
      has_gov_sku = false
      meters {
        name = "MP_INSP_JX"
        rate_allocation = 1
      }
    }

    billing_details {
      pricing_plan_key = "16ff7edf-a17f-4f77-88b8-1a09bef6df8f"
      billing_model = "USAGE_BASED"
      sku = "MP10258"
      metric_type = "INSTANCE_HOURS"
      rate_allocation = 1
      has_gov_sku = false
      meters {
        name = "MP_INSP_JY"
        rate_allocation = 1
      }
    }

    pricing_plan {
      plan_type = "HYBRID"
      plan_name = "Pro-Enterprice"
      plan_description = "Enterprise marketplace listing with licensed-user and output-token usage dimensions."
      billing_frequency = "ANNUAL"
      plan_duration = "ANNUAL"
      rates {
        currency = "USD"
        rate = 89
      }

      dimensions {
        dimension_key = "f5325c5f-a0fd-4a10-b026-abe93e3e5963"
        dimension_name = "Starter-Monthly-Licensed-User-Subscription"
        dimension_description = "Starter-Monthly-Licensed-User-Subscription"
        metric_type = "EACH"
        dimension_billing_frequency = "ANNUAL"
        included_quantity = 98
        rates {
          currency = "USD"
          rate = 87
        }
      }

      dimensions {
        dimension_key = "373936ee-23cc-4f06-92bc-4e8f926714e8"
        dimension_name = "Starter-Output-Token-Generation-Volume"
        dimension_description = "Starter-Output-Token-Generation-Volume"
        metric_type = "EACH"
        dimension_billing_frequency = "ANNUAL"
        included_quantity = 909
        rates {
          currency = "USD"
          rate = 768
        }
      }
    }
  }
}
`

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
		"partner_registration_url": acctest.Representation{RepType: acctest.Required, Create: `https://www.google.com/`},
		"pricing_plan":             acctest.RepresentationGroup{RepType: acctest.Required, Group: SelfSubscriptionSubscriptionDetailsPricingPlanRepresentation},
		"amount":                   acctest.Representation{RepType: acctest.Required, Create: `89`},
		"currency":                 acctest.Representation{RepType: acctest.Required, Create: `USD`},
		"is_auto_renew":            acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	SelfSubscriptionAdditionalDetailsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}
	SelfSubscriptionSubscriptionDetailsBillingDetailsRepresentation = map[string]interface{}{
		"billing_model":    acctest.Representation{RepType: acctest.Required, Create: `FLAT_RATE`},
		"meters":           acctest.RepresentationGroup{RepType: acctest.Required, Group: SelfSubscriptionSubscriptionDetailsBillingDetailsMetersRepresentation},
		"metric_type":      acctest.Representation{RepType: acctest.Required, Create: `INSTANCE_HOURS`},
		"pricing_plan_key": acctest.Representation{RepType: acctest.Required, Create: `16ff7edf-a17f-4f77-88b8-1a09bef6df8f`},
		"rate_allocation":  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"sku":              acctest.Representation{RepType: acctest.Required, Create: `MP10257`},
		"has_gov_sku":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	SelfSubscriptionSubscriptionDetailsPricingPlanRepresentation = map[string]interface{}{
		"billing_frequency": acctest.Representation{RepType: acctest.Required, Create: `YEARLY`},
		"plan_name":         acctest.Representation{RepType: acctest.Required, Create: `Pro-Enterprice`},
		"plan_type":         acctest.Representation{RepType: acctest.Required, Create: `HYBRID`},
		"rates":             acctest.RepresentationGroup{RepType: acctest.Required, Group: SelfSubscriptionSubscriptionDetailsPricingPlanRatesRepresentation},
		"dimensions":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: SelfSubscriptionSubscriptionDetailsPricingPlanDimensionsRepresentation},
		"plan_description":  acctest.Representation{RepType: acctest.Required, Create: `These longer names are more descriptive and suitable for enterprise marketplace listings, billing dimensions, and SaaS metering systems.`},
		"plan_duration":     acctest.Representation{RepType: acctest.Required, Create: `ANNUAL`},
	}
	SelfSubscriptionSubscriptionDetailsBillingDetailsMetersRepresentation = map[string]interface{}{
		"name":              acctest.Representation{RepType: acctest.Required, Create: `MP_INSP_JX`},
		"rate_allocation":   acctest.Representation{RepType: acctest.Required, Create: `1`},
		"extended_metadata": acctest.RepresentationGroup{RepType: acctest.Optional, Group: SelfSubscriptionSubscriptionDetailsBillingDetailsMetersExtendedMetadataRepresentation},
	}
	SelfSubscriptionSubscriptionDetailsPricingPlanRatesRepresentation = map[string]interface{}{
		"currency": acctest.Representation{RepType: acctest.Required, Create: `USD`},
		"rate":     acctest.Representation{RepType: acctest.Required, Create: `89`},
	}
	SelfSubscriptionSubscriptionDetailsPricingPlanDimensionsRepresentation = map[string]interface{}{
		"dimension_billing_frequency": acctest.Representation{RepType: acctest.Required, Create: `MONTHLY`},
		"dimension_description":       acctest.Representation{RepType: acctest.Required, Create: `dimensionDescription`},
		"dimension_key":               acctest.Representation{RepType: acctest.Required, Create: `dimensionKey`},
		"dimension_name":              acctest.Representation{RepType: acctest.Required, Create: `dimensionName`},
		"metric_type":                 acctest.Representation{RepType: acctest.Required, Create: `OCPU_HOURS`},
		"rates":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: SelfSubscriptionSubscriptionDetailsPricingPlanDimensionsRatesRepresentation},
		"included_quantity":           acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
	}
	SelfSubscriptionSubscriptionDetailsBillingDetailsMetersExtendedMetadataRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}
	SelfSubscriptionSubscriptionDetailsPricingPlanDimensionsRatesRepresentation = map[string]interface{}{
		"currency": acctest.Representation{RepType: acctest.Required, Create: `currency`},
		"rate":     acctest.Representation{RepType: acctest.Required, Create: `1.0`},
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

	var resId string
	tenantId := utils.GetEnvSettingWithBlankDefault("tenant_id")
	tenantIdVariableStr := fmt.Sprintf("variable \"tenant_id\" { default = \"%s\" }\n", tenantId)

	resourceName := "oci_self_subscription.test_subscription"
	//datasourceName := "data.oci_self_subscriptions.test_subscriptions"
	//singularDatasourceName := "data.oci_self_subscription.test_subscription"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+productIdVariableStr+sellerIdVariableStr+tenantIdVariableStr+SelfSubscriptionResourceDependencies+
		selfSubscriptionAcceptanceResource, "self", "subscription", t)

	acctest.ResourceTest(t, testAccCheckSelfSubscriptionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SelfSubscriptionResourceDependencies + productIdVariableStr + sellerIdVariableStr + tenantIdVariableStr +
				selfSubscriptionAcceptanceResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "product_id"),
				resource.TestCheckResourceAttrSet(resourceName, "seller_id"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.billing_model", "FLAT_RATE"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.name", "MP_INSP_JX"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.rate_allocation", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.metric_type", "INSTANCE_HOURS"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.pricing_plan_key", "16ff7edf-a17f-4f77-88b8-1a09bef6df8f"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.rate_allocation", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.sku", "MP10257"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.partner_registration_url", "https://www.google.com/"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.billing_frequency", "ANNUAL"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_name", "Pro-Enterprice"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_type", "HYBRID"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.0.currency", "USD"),
				resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.0.rate", "89"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),

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

		////delete before next Create
		//{
		//	Config: config + compartmentIdVariableStr + SelfSubscriptionResourceDependencies + productIdVariableStr + sellerIdVariableStr + tenantIdVariableStr,
		//},
		/*
			},

			// verify Update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SelfSubscriptionResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_self_subscription", "test_subscription", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(SelfSubscriptionRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "additional_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "additional_details.0.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "additional_details.0.value", "value"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "product_id"),
					resource.TestCheckResourceAttr(resourceName, "realm", "realm"),
					resource.TestCheckResourceAttr(resourceName, "region", "region"),
					resource.TestCheckResourceAttrSet(resourceName, "seller_id"),
					resource.TestCheckResourceAttr(resourceName, "source_type", "OCI_NATIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.amount", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.billing_model", "FLAT_RATE"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.has_gov_sku", "false"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.0.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.0.value", "value"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.name", "name"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.rate_allocation", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.metric_type", "OCPU_HOURS"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.pricing_plan_key", "pricingPlanKey"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.rate_allocation", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.sku", "sku"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.currency", "currency"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.is_auto_renew", "false"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.partner_registration_url", "partnerRegistrationUrl"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.billing_frequency", "MONTHLY"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_billing_frequency", "MONTHLY"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_description", "dimensionDescription"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_key", "dimensionKey"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_name", "dimensionName"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.included_quantity", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.metric_type", "OCPU_HOURS"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.rates.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.rates.0.currency", "currency"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.rates.0.rate", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_description", "planDescription"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_duration", "MONTHLY"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_name", "planName"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_type", "FIXED"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.0.currency", "currency"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.0.rate", "1.0"),
					resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
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
				Config: config + compartmentIdVariableStr + SelfSubscriptionResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_self_subscription", "test_subscription", acctest.Optional, acctest.Update, SelfSubscriptionRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "additional_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "additional_details.0.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "additional_details.0.value", "value"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "product_id"),
					resource.TestCheckResourceAttr(resourceName, "realm", "realm"),
					resource.TestCheckResourceAttr(resourceName, "region", "region"),
					resource.TestCheckResourceAttrSet(resourceName, "seller_id"),
					resource.TestCheckResourceAttr(resourceName, "source_type", "OCI_NATIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.amount", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.billing_model", "FLAT_RATE"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.has_gov_sku", "false"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.0.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.0.value", "value"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.name", "name"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.meters.0.rate_allocation", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.metric_type", "OCPU_HOURS"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.pricing_plan_key", "pricingPlanKey"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.rate_allocation", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.billing_details.0.sku", "sku"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.currency", "currency"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.is_auto_renew", "false"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.partner_registration_url", "partnerRegistrationUrl"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.billing_frequency", "MONTHLY"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_billing_frequency", "MONTHLY"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_description", "dimensionDescription"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_key", "dimensionKey"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_name", "dimensionName"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.included_quantity", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.metric_type", "OCPU_HOURS"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.rates.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.rates.0.currency", "currency"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.dimensions.0.rates.0.rate", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_description", "planDescription"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_duration", "MONTHLY"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_name", "planName"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.plan_type", "FIXED"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.0.currency", "currency"),
					resource.TestCheckResourceAttr(resourceName, "subscription_details.0.pricing_plan.0.rates.0.rate", "1.0"),
					resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_self_subscriptions", "test_subscriptions", acctest.Optional, acctest.Update, SelfSubscriptionDataSourceRepresentation) +
					compartmentIdVariableStr + SelfSubscriptionResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_self_subscription", "test_subscription", acctest.Optional, acctest.Update, SelfSubscriptionRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "id", "id"),

					resource.TestCheckResourceAttr(datasourceName, "subscription_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "subscription_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_self_subscription", "test_subscription", acctest.Required, acctest.Create, SelfSubscriptionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + SelfSubscriptionResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "additional_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "additional_details.0.key", "key"),
					resource.TestCheckResourceAttr(singularDatasourceName, "additional_details.0.value", "value"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "realm", "realm"),
					resource.TestCheckResourceAttr(singularDatasourceName, "region", "region"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_type", "OCI_NATIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.amount", "1.0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.billing_model", "FLAT_RATE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.has_gov_sku", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.meters.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.0.key", "key"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.meters.0.extended_metadata.0.value", "value"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.meters.0.name", "name"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.meters.0.rate_allocation", "1.0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.metric_type", "OCPU_HOURS"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.pricing_plan_key", "pricingPlanKey"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.rate_allocation", "1.0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.billing_details.0.sku", "sku"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.currency", "currency"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.is_auto_renew", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.partner_registration_url", "partnerRegistrationUrl"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.billing_frequency", "MONTHLY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_billing_frequency", "MONTHLY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_description", "dimensionDescription"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_key", "dimensionKey"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.dimensions.0.dimension_name", "dimensionName"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.dimensions.0.included_quantity", "1.0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.dimensions.0.metric_type", "OCPU_HOURS"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.dimensions.0.rates.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.dimensions.0.rates.0.currency", "currency"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.dimensions.0.rates.0.rate", "1.0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.plan_description", "planDescription"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.plan_duration", "MONTHLY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.plan_name", "planName"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.plan_type", "FIXED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.rates.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.rates.0.currency", "currency"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subscription_details.0.pricing_plan.0.rates.0.rate", "1.0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// verify resource import
			{
				Config:                  config + SelfSubscriptionRequiredOnlyResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
		*/
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
