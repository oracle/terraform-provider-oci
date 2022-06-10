---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_invoices"
sidebar_current: "docs-oci-datasource-onesubscription-invoices"
description: |-
  Provides the list of Invoices in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_invoices
This data source provides the list of Invoices in Oracle Cloud Infrastructure Onesubscription service.

This is a collection API which returns a list of Invoices for given filters.


## Example Usage

```hcl
data "oci_onesubscription_invoices" "test_invoices" {
	#Required
	ar_customer_transaction_id = oci_onesubscription_ar_customer_transaction.test_ar_customer_transaction.id
	compartment_id = var.compartment_id

	#Optional
	fields = var.invoice_fields
	time_from = var.invoice_time_from
	time_to = var.invoice_time_to
}
```

## Argument Reference

The following arguments are supported:

* `ar_customer_transaction_id` - (Required) AR Unique identifier for an invoice .
* `compartment_id` - (Required) The OCID of the root compartment.
* `fields` - (Optional) Partial response refers to an optimization technique offered by the RESTful web APIs to return only the information  (fields) required by the client. This parameter is used to control what fields to return. 
* `time_from` - (Optional) Initial date to filter Invoice data in SPM. 
* `time_to` - (Optional) Final date to filter Invoice data in SPM. 


## Attributes Reference

The following attributes are exported:

* `invoices` - The list of invoices.

### Invoice Reference

The following attributes are exported:

* `ar_invoices` - AR Invoice Numbers comma separated under one invoice 
* `bill_to_address` - Address. 
	* `bill_site_use_id` - Bill to site use Id. 
	* `is_bill_to` - Identify as the customer's billing address. 
	* `is_ship_to` - Identify as the customer's shipping address. 
	* `location` - Address location. 
		* `address1` - Address first line. 
		* `address2` - Address second line. 
		* `city` - City. 
		* `country` - Country. 
		* `postal_code` - Postal code. 
		* `region` - Region. 
		* `tca_location_id` - TCA Location identifier. 
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
	* `user_name` - userName. 
* `bill_to_customer` - Business partner. 
	* `customer_chain_type` - Customer chain type. 
	* `is_chain_customer` - The business partner is chain customer or not. 
	* `is_public_sector` - The business partner is part of the public sector or not. 
	* `name` - Commercial name also called customer name. 
	* `name_phonetic` - Phonetic name. 
	* `tca_customer_account_id` - TCA customer account ID. 
	* `tca_customer_account_number` - TCA customer account number. 
	* `tca_party_id` - TCA party ID. 
	* `tca_party_number` - TCA party number. 
* `created_by` - User that executed SPM Invoice process 
* `currency` - Currency details 
	* `iso_code` - Currency Code 
	* `name` - Currency name 
	* `std_precision` - Standard Precision of the Currency 
* `invoice_lines` - Invoice Lines under particular invoice. 
	* `ar_invoice_number` - AR Invoice Number for Invoice Line 
	* `data_center` - Data Center Attribute. 
	* `id` - SPM Invoice Line internal identifier 
	* `product` - Product description 
		* `billing_category` - Metered service billing category 
		* `name` - Product name 
		* `part_number` - Product part number 
		* `product_category` - Product category 
		* `ucm_rate_card_part_type` - Rate card part type of Product 
		* `unit_of_measure` - Unit of Measure 
	* `time_end` - Usage end time 
	* `time_start` - Usage start time 
* `organization` - Organization details 
	* `name` - Organization name 
	* `number` - Organization ID 
* `payment_method` - Payment Method 
* `payment_term` - Payment Term details 
	* `created_by` - User that created the Payment term 
	* `description` - Payment term Description 
	* `is_active` - Payment term active flag 
	* `name` - Payment Term name 
	* `time_created` - Payment term last update date 
	* `time_updated` - Payment term last update date 
	* `updated_by` - User that updated the Payment term 
	* `value` - Payment Term value 
* `receipt_method` - Receipt Method of Payment Mode 
* `spm_invoice_number` - SPM Document Number is an functional identifier for invoice in SPM 
* `status` - Document Status in SPM which depicts current state of invoice 
* `subscription_number` - Invoice associated subscription plan number. 
* `time_created` - SPM Invocie creation date 
* `time_invoice_date` - Invoice Date 
* `time_updated` - SPM Invoice updated date 
* `type` - Document Type in SPM like SPM Invoice,SPM Credit Memo etc., 
* `updated_by` - User that updated SPM Invoice 

