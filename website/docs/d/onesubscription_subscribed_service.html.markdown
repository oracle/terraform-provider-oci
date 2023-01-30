---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_subscribed_service"
sidebar_current: "docs-oci-datasource-onesubscription-subscribed_service"
description: |-
  Provides details about a specific Subscribed Service in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_subscribed_service
This data source provides details about a specific Subscribed Service resource in Oracle Cloud Infrastructure Onesubscription service.

This API returns the subscribed service details corresponding to the id provided


## Example Usage

```hcl
data "oci_onesubscription_subscribed_service" "test_subscribed_service" {
	#Required
	subscribed_service_id = oci_onesubscription_subscribed_service.test_subscribed_service.id

	#Optional
	fields = var.subscribed_service_fields
}
```

## Argument Reference

The following arguments are supported:

* `fields` - (Optional) Partial response refers to an optimization technique offered by the RESTful web APIs to return only the information  (fields) required by the client. In this mechanism, the client sends the required field names as the query parameters for an API to the server, and the server trims down the default response content by removing the fields that are not required by the client. The parameter used to control what fields to return should be a query string parameter called "fields" of type array, and usecollectionFormat 
* `subscribed_service_id` - (Required) The Subscribed Service Id


## Attributes Reference

The following attributes are exported:

* `admin_email` - Subscribed service admin email id 
* `agreement_id` - Subscribed service agreement ID 
* `agreement_name` - Subscribed service agrrement name 
* `agreement_type` - Subscribed service agrrement type 
* `available_amount` - Subscribed sercice available or remaining amount 
* `bill_to_address` - Address. 
	* `bill_site_use_id` - Bill to site use Id. 
	* `is_bill_to` - Identify as the customer shipping address. 
	* `is_ship_to` - Identify as the customer invoicing address. 
	* `location` - Address location. 
		* `address1` - Address first line. 
		* `address2` - Address second line. 
		* `city` - City. 
		* `country` - Country. 
		* `postal_code` - Postal code. 
		* `region` - Region. 
		* `tca_location_id` - Region. 
	* `name` - Address name identifier. 
	* `phone` - Phone. 
	* `service2site_use_id` - Service to site use Id. 
	* `tca_cust_acct_site_id` - TCA customer account site Id. 
	* `tca_party_site_number` - Party site number. 
* `bill_to_contact` - User. 
	* `email` - Email. 
	* `first_name` - First name. 
	* `last_name` - Last name. 
	* `name` - Name. 
	* `tca_contact_id` - TCA contact ID. 
	* `tca_cust_accnt_site_id` - TCA customer account site ID. 
	* `tca_party_id` - TCA party ID. 
	* `username` - Username. 
* `bill_to_customer` - Business partner. 
	* `customer_chain_type` - Customer chain type. 
	* `is_chain_customer` - The business partner is chain customer or not. 
	* `is_public_sector` - The business partner is part of the public sector or not. 
	* `name` - Commercial name also called customer name. 
	* `name_phonetic` - Phonetic name. 
	* `tca_cust_account_number` - TCA customer account number. 
	* `tca_customer_account_id` - TCA customer account ID. 
	* `tca_party_id` - TCA party ID. 
	* `tca_party_number` - TCA party number. 
* `billing_frequency` - Subscribed service invoice frequency 
* `booking_opty_number` - Booking Opportunity Number of Subscribed Service 
* `buyer_email` - Subscribed service buyer email id 
* `commitment_schedule_id` - Subscribed service commitment schedule Id 
* `commitment_services` - List of Commitment services of a line 
	* `available_amount` - Commitment available amount 
	* `funded_allocation_value` - Funded Allocation line value 
	* `line_net_amount` - Commitment line net amount 
	* `quantity` - Commitment quantity 
	* `time_end` - Commitment end date 
	* `time_start` - Commitment start date 
* `created_by` - User that created the subscribed service 
* `credit_percentage` - Subscribed service credit percentage 
* `csi` - Subscribed service CSI number 
* `customer_transaction_reference` - Identifier for a customer's transactions for purchase of ay oracle services 
* `data_center` - Subscribed service data center 
* `data_center_region` - Subscribed service data center region 
* `eligible_to_renew` - Subscribed service eligible to renew field 
* `end_user_address` - Address. 
	* `bill_site_use_id` - Bill to site use Id. 
	* `is_bill_to` - Identify as the customer shipping address. 
	* `is_ship_to` - Identify as the customer invoicing address. 
	* `location` - Address location. 
		* `address1` - Address first line. 
		* `address2` - Address second line. 
		* `city` - City. 
		* `country` - Country. 
		* `postal_code` - Postal code. 
		* `region` - Region. 
		* `tca_location_id` - Region. 
	* `name` - Address name identifier. 
	* `phone` - Phone. 
	* `service2site_use_id` - Service to site use Id. 
	* `tca_cust_acct_site_id` - TCA customer account site Id. 
	* `tca_party_site_number` - Party site number. 
* `end_user_contact` - User. 
	* `email` - Email. 
	* `first_name` - First name. 
	* `last_name` - Last name. 
	* `name` - Name. 
	* `tca_contact_id` - TCA contact ID. 
	* `tca_cust_accnt_site_id` - TCA customer account site ID. 
	* `tca_party_id` - TCA party ID. 
	* `username` - Username. 
* `end_user_customer` - Business partner. 
	* `customer_chain_type` - Customer chain type. 
	* `is_chain_customer` - The business partner is chain customer or not. 
	* `is_public_sector` - The business partner is part of the public sector or not. 
	* `name` - Commercial name also called customer name. 
	* `name_phonetic` - Phonetic name. 
	* `tca_cust_account_number` - TCA customer account number. 
	* `tca_customer_account_id` - TCA customer account ID. 
	* `tca_party_id` - TCA party ID. 
	* `tca_party_number` - TCA party number. 
* `fulfillment_set` - Subscribed service fulfillment set 
* `funded_allocation_value` - Funded Allocation line value example: 12000.00 
* `id` - SPM internal Subscribed Service ID 
* `is_allowance` - Indicates if a service can recieve usages and consequently have available amounts computed 
* `is_cap_to_price_list` - If true compares rate between ratecard and the active pricelist and minimum rate would be fetched 
* `is_credit_enabled` - Used in context of service credit lines 
* `is_having_usage` - Indicator on whether or not there has been usage for the subscribed service 
* `is_intent_to_pay` - Subscribed service intent to pay flag 
* `is_payg` - Subscribed service payg flag 
* `is_single_rate_card` - Indicates if the Subscribed service has a single ratecard 
* `is_variable_commitment` - Indicates if the commitment lines can have different quantities 
* `line_net_amount` - Subscribed service line net amount 
* `major_set` - Subscribed service Major Set 
* `net_unit_price` - Subscribed service net unit price 
* `operation_type` - Subscribed service operation type 
* `order_header_id` - Sales Order Header associated to the subscribed service 
* `order_line_id` - Sales Order Line Id associated to the subscribed service 
* `order_line_number` - Sales Order Line Number associated to the subscribed service 
* `order_number` - Sales Order Number associated to the subscribed service 
* `order_type` - Order Type of Subscribed Service 
* `original_promo_amount` - Subscribed service Promotion Amount 
* `overage_bill_to` - Overage Bill To of Subscribed Service 
* `overage_discount_percentage` - Subscribed service Overage Discount Percentage 
* `overage_policy` - Overage Policy of Subscribed Service 
* `partner_credit_amount` - Subscribed service partner credit amount 
* `partner_transaction_type` - This field contains the name of the partner to which the subscription belongs - depending on which the invoicing may differ 
* `payg_policy` - Pay As You Go policy of Subscribed Service (Can be null - indicating no payg policy) 
* `payment_method` - Payment Method of Subscribed Service 
* `payment_number` - Payment Number of Subscribed Service 
* `payment_term` - Payment Term details 
	* `created_by` - User that created the Payment term 
	* `description` - Payment term Description 
	* `is_active` - Payment term active flag 
	* `name` - Payment Term name 
	* `time_created` - Payment term last update date 
	* `time_updated` - Payment term last update date 
	* `updated_by` - User that updated the Payment term 
	* `value` - Payment Term value 
* `price_period` - Indicates the period for which the commitment amount can be utilised exceeding which the amount lapses. Also used in calculation of total contract line value 
* `pricing_model` - Subscribed service pricing model 
* `product` - Product description 
	* `billing_category` - Metered service billing category 
	* `name` - Product name 
	* `part_number` - Product part numner 
	* `product_category` - Product category 
	* `ucm_rate_card_part_type` - Rate card part type of Product 
	* `unit_of_measure` - Unit of measure 
* `program_type` - Subscribed service program type 
* `promo_order_line_id` - Not null if this service has an associated promotion line in SPM. Contains the line identifier from Order Management of  the associated promo line. 
* `promo_type` - Subscribed service promotion type 
* `promotion_pricing_type` - Promotion Pricing Type of Subscribed Service (Can be null - indicating no promotion pricing) 
* `provisioning_source` - Subscribed service provisioning source 
* `quantity` - Subscribed service quantity 
* `rate_card_discount_percentage` - Subscribed service Rate Card Discount Percentage 
* `rate_cards` - List of Rate Cards of a Subscribed Service 
	* `currency` - Currency details 
		* `iso_code` - Currency Code 
		* `name` - Currency name 
		* `std_precision` - Standard Precision of the Currency 
	* `discretionary_discount_percentage` - Rate card discretionary discount percentage 
	* `is_tier` - Rate card price tier flag 
	* `net_unit_price` - Rate card net unit price 
	* `overage_price` - Rate card overage price 
	* `product` - Product description 
		* `billing_category` - Metered service billing category 
		* `name` - Product name 
		* `part_number` - Product part numner 
		* `product_category` - Product category 
		* `ucm_rate_card_part_type` - Rate card part type of Product 
		* `unit_of_measure` - Unit of measure 
	* `rate_card_tiers` - List of tiered rate card prices 
		* `net_unit_price` - Rate card tier net unit price 
		* `overage_price` - Rate card tier overage price 
		* `up_to_quantity` - Rate card tier quantity range 
	* `subscribed_service_id` - SPM internal Subscribed Service ID 
	* `time_end` - Rate card end date 
	* `time_start` - Rate card start date 
* `ratecard_type` - SPM Ratecard Type 
* `renewal_opty_id` - Subscribed service Opportunity Id 
* `renewal_opty_number` - Renewal Opportunity Number of Subscribed Service 
* `renewal_opty_type` - Renewal Opportunity Type of Subscribed Service 
* `renewed_subscribed_service_id` - SPM renewed Subscription ID 
* `reseller_address` - Address. 
	* `bill_site_use_id` - Bill to site use Id. 
	* `is_bill_to` - Identify as the customer shipping address. 
	* `is_ship_to` - Identify as the customer invoicing address. 
	* `location` - Address location. 
		* `address1` - Address first line. 
		* `address2` - Address second line. 
		* `city` - City. 
		* `country` - Country. 
		* `postal_code` - Postal code. 
		* `region` - Region. 
		* `tca_location_id` - Region. 
	* `name` - Address name identifier. 
	* `phone` - Phone. 
	* `service2site_use_id` - Service to site use Id. 
	* `tca_cust_acct_site_id` - TCA customer account site Id. 
	* `tca_party_site_number` - Party site number. 
* `reseller_contact` - User. 
	* `email` - Email. 
	* `first_name` - First name. 
	* `last_name` - Last name. 
	* `name` - Name. 
	* `tca_contact_id` - TCA contact ID. 
	* `tca_cust_accnt_site_id` - TCA customer account site ID. 
	* `tca_party_id` - TCA party ID. 
	* `username` - Username. 
* `reseller_customer` - Business partner. 
	* `customer_chain_type` - Customer chain type. 
	* `is_chain_customer` - The business partner is chain customer or not. 
	* `is_public_sector` - The business partner is part of the public sector or not. 
	* `name` - Commercial name also called customer name. 
	* `name_phonetic` - Phonetic name. 
	* `tca_cust_account_number` - TCA customer account number. 
	* `tca_customer_account_id` - TCA customer account ID. 
	* `tca_party_id` - TCA party ID. 
	* `tca_party_number` - TCA party number. 
* `revenue_line_id` - Subscribed service Revenue Line Id 
* `revenue_line_number` - Revenue Line NUmber of Subscribed Service 
* `revised_arr_in_lc` - Subscribed service Revised ARR 
* `revised_arr_in_sc` - Subscribed service Revised ARR in Standard Currency 
* `sales_account_party_id` - Subscribed service sales account party id 
* `sales_channel` - Sales Channel of Subscribed Service 
* `serial_number` - Subscribed service line number 
* `service_to_address` - Address. 
	* `bill_site_use_id` - Bill to site use Id. 
	* `is_bill_to` - Identify as the customer shipping address. 
	* `is_ship_to` - Identify as the customer invoicing address. 
	* `location` - Address location. 
		* `address1` - Address first line. 
		* `address2` - Address second line. 
		* `city` - City. 
		* `country` - Country. 
		* `postal_code` - Postal code. 
		* `region` - Region. 
		* `tca_location_id` - Region. 
	* `name` - Address name identifier. 
	* `phone` - Phone. 
	* `service2site_use_id` - Service to site use Id. 
	* `tca_cust_acct_site_id` - TCA customer account site Id. 
	* `tca_party_site_number` - Party site number. 
* `service_to_contact` - User. 
	* `email` - Email. 
	* `first_name` - First name. 
	* `last_name` - Last name. 
	* `name` - Name. 
	* `tca_contact_id` - TCA contact ID. 
	* `tca_cust_accnt_site_id` - TCA customer account site ID. 
	* `tca_party_id` - TCA party ID. 
	* `username` - Username. 
* `service_to_customer` - Business partner. 
	* `customer_chain_type` - Customer chain type. 
	* `is_chain_customer` - The business partner is chain customer or not. 
	* `is_public_sector` - The business partner is part of the public sector or not. 
	* `name` - Commercial name also called customer name. 
	* `name_phonetic` - Phonetic name. 
	* `tca_cust_account_number` - TCA customer account number. 
	* `tca_customer_account_id` - TCA customer account ID. 
	* `tca_party_id` - TCA party ID. 
	* `tca_party_number` - TCA party number. 
* `sold_to_contact` - User. 
	* `email` - Email. 
	* `first_name` - First name. 
	* `last_name` - Last name. 
	* `name` - Name. 
	* `tca_contact_id` - TCA contact ID. 
	* `tca_cust_accnt_site_id` - TCA customer account site ID. 
	* `tca_party_id` - TCA party ID. 
	* `username` - Username. 
* `sold_to_customer` - Business partner. 
	* `customer_chain_type` - Customer chain type. 
	* `is_chain_customer` - The business partner is chain customer or not. 
	* `is_public_sector` - The business partner is part of the public sector or not. 
	* `name` - Commercial name also called customer name. 
	* `name_phonetic` - Phonetic name. 
	* `tca_cust_account_number` - TCA customer account number. 
	* `tca_customer_account_id` - TCA customer account ID. 
	* `tca_party_id` - TCA party ID. 
	* `tca_party_number` - TCA party number. 
* `start_date_type` - Subscribed service start date type 
* `status` - Subscribed service status 
* `subscription_id` - Subscription ID associated to the subscribed service 
* `subscription_source` - Subscribed service source 
* `system_arr_in_lc` - Subscribed service System ARR 
* `system_arr_in_sc` - Subscribed service System ARR in Standard Currency 
* `system_atr_arr_in_lc` - Subscribed service System ATR-ARR 
* `system_atr_arr_in_sc` - Subscribed service System ATR-ARR in Standard Currency 
* `term_value` - Term value in Months 
* `term_value_uom` - Term value UOM 
* `time_agreement_end` - Subscribed service agrrement end date 
* `time_created` - Subscribed service creation date 
* `time_customer_config` - Subscribed service customer config date 
* `time_end` - Subscribed service end date 
* `time_majorset_end` - Subscribed service Major Set End date 
* `time_majorset_start` - Subscribed service Major Set Start date 
* `time_payment_expiry` - Subscribed service payment expiry date 
* `time_provisioned` - Subscribed service provisioning date 
* `time_service_configuration_email_sent` - Subscribed service service configuration email sent date 
* `time_start` - Subscribed service start date 
* `time_updated` - Subscribed service last update date 
* `time_welcome_email_sent` - Subscribed service welcome email sent date 
* `total_value` - Subscribed service total value 
* `transaction_extension_id` - Subscribed service Transaction Extension Id 
* `type` - Subscribed Service line type 
* `updated_by` - User that updated the subscribed service 
* `used_amount` - Subscribed service used amount 

