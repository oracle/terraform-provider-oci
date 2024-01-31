---
subcategory: "Osp Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osp_gateway_subscriptions"
sidebar_current: "docs-oci-datasource-osp_gateway-subscriptions"
description: |-
  Provides the list of Subscriptions in Oracle Cloud Infrastructure Osp Gateway service
---

# Data Source: oci_osp_gateway_subscriptions
This data source provides the list of Subscriptions in Oracle Cloud Infrastructure Osp Gateway service.

Get the subscription data for the compartment

## Example Usage

```hcl
data "oci_osp_gateway_subscriptions" "test_subscriptions" {
	#Required
	compartment_id = var.compartment_id
	osp_home_region = var.subscription_osp_home_region
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `osp_home_region` - (Required) The home region's public name of the logged in user. 


## Attributes Reference

The following attributes are exported:

* `subscription_collection` - The list of subscription_collection.

### Subscription Reference

The following attributes are exported:

* `account_type` - Account type.
* `bill_to_cust_account_id` - Bill to customer Account id.
* `billing_address` - Address details model.
	* `address_key` - Address identifier.
	* `city` - Name of the city.
	* `company_name` - Name of the customer company.
	* `contributor_class` - Contributor class of the customer company.
	* `country` - Country of the address.
	* `county` - County of the address.
	* `department_name` - Department name of the customer company.
	* `email_address` - Contact person email address.
	* `first_name` - First name of the contact person.
	* `internal_number` - Internal number of the customer company.
	* `job_title` - Job title of the contact person.
	* `last_name` - Last name of the contact person.
	* `line1` - Address line 1.
	* `line2` - Address line 2.
	* `line3` - Address line 3.
	* `line4` - Address line 4.
	* `middle_name` - Middle name of the contact person.
	* `municipal_inscription` - Municipal Inscription.
	* `phone_country_code` - Phone country code of the contact person.
	* `phone_number` - Phone number of the contact person.
	* `postal_code` - Post code of the address.
	* `province` - Province of the address.
	* `state` - State of the address.
	* `state_inscription` - State Inscription.
	* `street_name` - Street name of the address.
	* `street_number` - Street number of the address.
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
	* `credit_card_type` - Credit card type.
	* `email_address` - The email address of the paypal user.
	* `ext_billing_agreement_id` - Agreement id for the paypal account.
	* `first_name` - First name of the paypal user.
	* `last_digits` - Last four digits of the card.
	* `last_name` - Last name of the paypal user.
	* `name_on_card` - Name on the credit card.
	* `payment_method` - Payment method
	* `time_expiration` - Expired date of the credit card.
	* `wallet_instrument_id` - Wallet instrument internal id.
	* `wallet_transaction_id` - Wallet transaction id.
* `plan_type` - Subscription plan type.
* `ship_to_cust_acct_role_id` - Ship to customer account role.
* `ship_to_cust_acct_site_id` - Ship to customer account site address id.
* `subscription_plan_number` - Subscription plan number.
* `tax_info` - Tax details.
	* `giro` - Companies' GIRO code
	* `no_tax_reason_code` - Tax exemption reason code.
	* `no_tax_reason_code_details` - Tax exemption reason description.
	* `tax_cnpj` - Brazilian companies' CNPJ number.
	* `tax_payer_id` - Tay payer identifier.
	* `tax_reg_number` - Tax registration number.
* `time_personal_to_corporate_conv` - Date of upgrade/conversion when account type changed from PERSONAL to CORPORATE
* `time_plan_upgrade` - Date of upgrade/conversion when planType changed from FREE_TIER to PAYG
* `time_start` - Start date of the subscription.
* `upgrade_state` - Status of the upgrade.
* `upgrade_state_details` - This field is used to describe the Upgrade State in case of error (E.g. Upgrade failure caused by interfacing Tax details- TaxError)

