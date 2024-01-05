// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OnesubscriptionOnesubscriptionSubscribedServiceSingularDataSourceRepresentation = map[string]interface{}{
		"subscribed_service_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscribed_service_id}`},
		"fields":                acctest.Representation{RepType: acctest.Optional, Create: `${var.fields}`},
	}

	OnesubscriptionOnesubscriptionSubscribedServiceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"order_line_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.order_line_id}`},
		"status":          acctest.Representation{RepType: acctest.Optional, Create: `status`},
	}
)

// issue-routing-tag: onesubscription/default
func TestOnesubscriptionSubscribedServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOnesubscriptionSubscribedServiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subsServId := utils.GetEnvSettingWithBlankDefault("subscribed_service_id")
	subsServIdVariableStr := fmt.Sprintf("variable \"subscribed_service_id\" { default = \"%s\" }\n", subsServId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_onesubscription_subscribed_services.test_subscribed_services"
	singularDatasourceName := "data.oci_onesubscription_subscribed_service.test_subscribed_service"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_subscribed_services", "test_subscribed_services", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionSubscribedServiceDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttrSet(datasourceName, "order_line_id"),
				//resource.TestCheckResourceAttr(datasourceName, "status", "status"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.admin_email"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.agreement_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.agreement_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.agreement_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.available_amount"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.bill_to_address.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.bill_to_contact.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.bill_to_customer.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.billing_frequency"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.booking_opty_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.buyer_email"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.commitment_schedule_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.created_by"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.credit_percentage"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.csi"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.customer_transaction_reference"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.data_center"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.data_center_region"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.eligible_to_renew"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.end_user_address.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.end_user_contact.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.end_user_customer.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.fulfillment_set"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.funded_allocation_value"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.is_allowance"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.is_cap_to_price_list"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.is_credit_enabled"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.is_having_usage"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.is_intent_to_pay"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.is_payg"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.is_single_rate_card"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.is_variable_commitment"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.line_net_amount"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.major_set"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.net_unit_price"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.operation_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.order_header_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.order_line_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.order_line_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.order_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.order_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.original_promo_amount"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.overage_bill_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.overage_discount_percentage"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.overage_policy"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.partner_credit_amount"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.partner_transaction_type"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.payg_policy"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.payment_method"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.payment_number"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.payment_term.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.price_period"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.pricing_model"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.product.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.promo_order_line_id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.promo_type"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.promotion_pricing_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.provisioning_source"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.quantity"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.rate_card_discount_percentage"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.ratecard_type"),
				/*resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.renewal_opty_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.renewal_opty_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.renewal_opty_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.renewed_subscribed_service_id"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.reseller_address.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.reseller_contact.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.reseller_customer.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.revenue_line_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.revenue_line_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.revised_arr_in_lc"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.revised_arr_in_sc"),*/
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.sales_account_party_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.sales_channel"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.serial_number"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.service_to_address.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.service_to_contact.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.service_to_customer.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.sold_to_contact.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscribed_services.0.sold_to_customer.#", "1"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.start_date_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.subscription_source"),
				/*resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.system_arr_in_lc"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.system_arr_in_sc"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.system_atr_arr_in_lc"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.system_atr_arr_in_sc"),*/
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.term_value"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.term_value_uom"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_agreement_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_created"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_customer_config"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_majorset_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_majorset_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_payment_expiry"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_provisioned"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_service_configuration_email_sent"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_updated"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.time_welcome_email_sent"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.total_value"),
				//resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.transaction_extension_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.updated_by"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscribed_services.0.used_amount"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_subscribed_service", "test_subscribed_service", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionSubscribedServiceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subsServIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(singularDatasourceName, "fields.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscribed_service_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "admin_email"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "agreement_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "agreement_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "agreement_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_amount"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bill_to_address.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bill_to_contact.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bill_to_customer.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "billing_frequency"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "booking_opty_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "buyer_email"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "commitment_schedule_id"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "commitment_services.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "credit_percentage"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "csi"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "customer_transaction_reference"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_center"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_center_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "eligible_to_renew"),
				resource.TestCheckResourceAttr(singularDatasourceName, "end_user_address.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "end_user_contact.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "end_user_customer.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fulfillment_set"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "funded_allocation_value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_allowance"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cap_to_price_list"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_credit_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_having_usage"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_intent_to_pay"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "is_payg"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_single_rate_card"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_variable_commitment"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "line_net_amount"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "major_set"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "net_unit_price"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operation_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "order_header_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "order_line_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "order_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "order_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "original_promo_amount"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "overage_bill_to"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "overage_discount_percentage"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "overage_policy"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "partner_credit_amount"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "partner_transaction_type"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "payg_policy"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "payment_method"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "payment_number"),
				resource.TestCheckResourceAttr(singularDatasourceName, "payment_term.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "price_period"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pricing_model"),
				resource.TestCheckResourceAttr(singularDatasourceName, "product.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "promo_order_line_id"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "promo_type"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "promotion_pricing_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "provisioning_source"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "quantity"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "rate_card_discount_percentage"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rate_cards.#", "0"),
				/*resource.TestCheckResourceAttrSet(singularDatasourceName, "ratecard_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "renewal_opty_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "renewal_opty_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "renewal_opty_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "renewed_subscribed_service_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "reseller_address.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "reseller_contact.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "reseller_customer.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "revenue_line_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "revenue_line_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "revised_arr_in_lc"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "revised_arr_in_sc"),*/
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sales_account_party_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sales_channel"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "serial_number"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_to_address.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_to_contact.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_to_customer.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sold_to_contact.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sold_to_customer.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "start_date_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_source"),
				/*resource.TestCheckResourceAttrSet(singularDatasourceName, "system_arr_in_lc"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_arr_in_sc"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_atr_arr_in_lc"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_atr_arr_in_sc"),*/
				resource.TestCheckResourceAttrSet(singularDatasourceName, "term_value"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "term_value_uom"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_agreement_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_customer_config"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_majorset_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_majorset_start"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_payment_expiry"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_provisioned"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_service_configuration_email_sent"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_start"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_welcome_email_sent"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_value"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "transaction_extension_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "updated_by"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "used_amount"),
			),
		},
	})
}
