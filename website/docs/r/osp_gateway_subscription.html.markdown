---
subcategory: "Osp Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osp_gateway_subscription"
sidebar_current: "docs-oci-resource-osp_gateway-subscription"
description: |-
  Provides the Subscription resource in Oracle Cloud Infrastructure Osp Gateway service
---

# oci_osp_gateway_subscription
This resource provides the Subscription resource in Oracle Cloud Infrastructure Osp Gateway service.

Update plan of the subscription.

## Example Usage

```hcl
resource "oci_osp_gateway_subscription" "test_subscription" {
	#Required
	compartment_id = var.compartment_id
	email = var.subscription_email
	osp_home_region = var.subscription_osp_home_region
	subscription {
		#Required
		subscription_plan_number = var.subscription_subscription_subscription_plan_number

		#Optional
		bill_to_cust_account_id = oci_osp_gateway_bill_to_cust_account.test_bill_to_cust_account.id
		billing_address {

			#Optional
			address_key = var.subscription_subscription_billing_address_address_key
			city = var.subscription_subscription_billing_address_city
			company_name = var.subscription_subscription_billing_address_company_name
			country = var.subscription_subscription_billing_address_country
			email_address = var.subscription_subscription_billing_address_email_address
			first_name = var.subscription_subscription_billing_address_first_name
			last_name = var.subscription_subscription_billing_address_last_name
			line1 = var.subscription_subscription_billing_address_line1
			line2 = var.subscription_subscription_billing_address_line2
			postal_code = var.subscription_subscription_billing_address_postal_code
			state = var.subscription_subscription_billing_address_state
		}
		currency_code = var.subscription_subscription_currency_code
		gsi_org_code = var.subscription_subscription_gsi_org_code
		id = var.subscription_subscription_id
		is_intent_to_pay = var.subscription_subscription_is_intent_to_pay
		language_code = var.subscription_subscription_language_code
		organization_id = oci_osp_gateway_organization.test_organization.id
		payment_gateway {

			#Optional
			merchant_defined_data {

				#Optional
				cloud_account_name = var.subscription_subscription_payment_gateway_merchant_defined_data_cloud_account_name
				promo_type = var.subscription_subscription_payment_gateway_merchant_defined_data_promo_type
			}
		}
		payment_options {
			#Required
			payment_method = var.subscription_subscription_payment_options_payment_method

			#Optional
			wallet_instrument_id = oci_osp_gateway_wallet_instrument.test_wallet_instrument.id
			wallet_transaction_id = oci_osp_gateway_wallet_transaction.test_wallet_transaction.id
		}
		plan_type = var.subscription_subscription_plan_type
		ship_to_cust_acct_role_id = oci_osp_gateway_ship_to_cust_acct_role.test_ship_to_cust_acct_role.id
		ship_to_cust_acct_site_id = oci_osp_gateway_ship_to_cust_acct_site.test_ship_to_cust_acct_site.id
		tax_info {

			#Optional
			no_tax_reason_code = var.subscription_subscription_tax_info_no_tax_reason_code
			no_tax_reason_code_details = var.subscription_subscription_tax_info_no_tax_reason_code_details
			tax_cnpj = var.subscription_subscription_tax_info_tax_cnpj
			tax_payer_id = oci_osp_gateway_tax_payer.test_tax_payer.id
			tax_reg_number = var.subscription_subscription_tax_info_tax_reg_number
		}
		time_plan_upgrade = var.subscription_subscription_time_plan_upgrade
		time_start = var.subscription_subscription_time_start
		upgrade_state = var.subscription_subscription_upgrade_state
		upgrade_state_details = var.subscription_subscription_upgrade_state_details
	}
	subscription_id = oci_osp_gateway_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `email` - (Required) (Updatable) User email
* `osp_home_region` - (Required) (Updatable) The home region's public name of the logged in user. 
* `subscription` - (Required) (Updatable) Subscription details object which extends the SubscriptionSummary
	* `bill_to_cust_account_id` - (Optional) (Updatable) Bill to customer Account id.
	* `billing_address` - (Optional) (Updatable) Billing address details model.
		* `address_key` - (Optional) (Updatable) Address identifier.
		* `city` - (Optional) (Updatable) Name of the city.
		* `company_name` - (Optional) (Updatable) Name of the customer company.
		* `country` - (Optional) (Updatable) Country of the address.
		* `email_address` - (Optional) (Updatable) Contact person email address.
		* `first_name` - (Optional) (Updatable) First name of the contact person.
		* `last_name` - (Optional) (Updatable) Last name of the contact person.
		* `line1` - (Optional) (Updatable) Address line 1.
		* `line2` - (Optional) (Updatable) Address line 2.
		* `postal_code` - (Optional) (Updatable) Post code of the address.
		* `state` - (Optional) (Updatable) State of the address.
	* `currency_code` - (Optional) (Updatable) Currency code
	* `gsi_org_code` - (Optional) (Updatable) GSI Subscription external code.
	* `id` - (Optional) (Updatable) Subscription id identifier (OCID).
	* `is_intent_to_pay` - (Optional) (Updatable) Payment intension.
	* `language_code` - (Optional) (Updatable) Language short code (en, de, hu, etc)
	* `organization_id` - (Optional) (Updatable) GSI organization external identifier.
	* `payment_gateway` - (Optional) (Updatable) Payment gateway details.
		* `merchant_defined_data` - (Optional) (Updatable) Merchant details.
			* `cloud_account_name` - (Optional) (Updatable) Cloud account name.
			* `promo_type` - (Optional) (Updatable) Promotion type code.
	* `payment_options` - (Optional) (Updatable) Payment option list of a subscription.
		* `payment_method` - (Required) (Updatable) Payment method
		* `wallet_instrument_id` - (Optional) (Updatable) Wallet instrument internal id.
		* `wallet_transaction_id` - (Optional) (Updatable) Wallet transaction id.
	* `plan_type` - (Optional) (Updatable) Subscription plan type.
	* `ship_to_cust_acct_role_id` - (Optional) (Updatable) Ship to customer account role.
	* `ship_to_cust_acct_site_id` - (Optional) (Updatable) Ship to customer account site address id.
	* `subscription_plan_number` - (Required) (Updatable) Subscription plan number.
	* `tax_info` - (Optional) (Updatable) Tax details.
		* `no_tax_reason_code` - (Optional) (Updatable) Tax exemption reason code.
		* `no_tax_reason_code_details` - (Optional) (Updatable) Tax exemption reason description.
		* `tax_cnpj` - (Optional) (Updatable) Brazilian companies' CNPJ number.
		* `tax_payer_id` - (Optional) (Updatable) Tay payer identifier.
		* `tax_reg_number` - (Optional) (Updatable) Tax registration number.
	* `time_plan_upgrade` - (Optional) (Updatable) Date of upgrade/conversion when planType changed from FREE_TIER to PAYG
	* `time_start` - (Optional) (Updatable) Start date of the subscription.
	* `upgrade_state` - (Optional) (Updatable) Status of the upgrade.
	* `upgrade_state_details` - (Optional) (Updatable) This field is used to describe the Upgrade State in case of error (E.g. Upgrade failure caused by interfacing Tax details- TaxError)
* `subscription_id` - (Required) Subscription id(OCID).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bill_to_cust_account_id` - Bill to customer Account id.
* `billing_address` - Billing address details model.
	* `address_key` - Address identifier.
	* `city` - Name of the city.
	* `company_name` - Name of the customer company.
	* `country` - Country of the address.
	* `email_address` - Contact person email address.
	* `first_name` - First name of the contact person.
	* `last_name` - Last name of the contact person.
	* `line1` - Address line 1.
	* `line2` - Address line 2.
	* `postal_code` - Post code of the address.
	* `state` - State of the address.
* `currency_code` - Currency code
* `gsi_org_code` - GSI Subscription external code.
* `id` - Subscription id identifier (OCID).
* `is_intent_to_pay` - Payment intension.
* `language_code` - Language short code (en, de, hu, etc)
* `organization_id` - GSI organization external identifier.
* `payment_gateway` - Payment gateway details.
	* `merchant_defined_data` - Merchant details.
		* `cloud_account_name` - Cloud account name.
		* `promo_type` - Promotion type code.
* `payment_options` - Payment option list of a subscription.
	* `payment_method` - Payment method
	* `wallet_instrument_id` - Wallet instrument internal id.
	* `wallet_transaction_id` - Wallet transaction id.
* `plan_type` - Subscription plan type.
* `ship_to_cust_acct_role_id` - Ship to customer account role.
* `ship_to_cust_acct_site_id` - Ship to customer account site address id.
* `subscription_plan_number` - Subscription plan number.
* `tax_info` - Tax details.
	* `no_tax_reason_code` - Tax exemption reason code.
	* `no_tax_reason_code_details` - Tax exemption reason description.
	* `tax_cnpj` - Brazilian companies' CNPJ number.
	* `tax_payer_id` - Tay payer identifier.
	* `tax_reg_number` - Tax registration number.
* `time_plan_upgrade` - Date of upgrade/conversion when planType changed from FREE_TIER to PAYG
* `time_start` - Start date of the subscription.
* `upgrade_state` - Status of the upgrade.
* `upgrade_state_details` - This field is used to describe the Upgrade State in case of error (E.g. Upgrade failure caused by interfacing Tax details- TaxError)

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subscription
	* `update` - (Defaults to 20 minutes), when updating the Subscription
	* `delete` - (Defaults to 20 minutes), when destroying the Subscription


## Import

Subscriptions can be imported using the `id`, e.g.

```
$ terraform import oci_osp_gateway_subscription.test_subscription "subscriptions/{subscriptionId}/compartmentId/{compartmentId}/ospHomeRegion/{ospHomeRegion}" 
```

