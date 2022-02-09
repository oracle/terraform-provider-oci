---
subcategory: "Osp Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osp_gateway_invoices"
sidebar_current: "docs-oci-datasource-osp_gateway-invoices"
description: |-
  Provides the list of Invoices in Oracle Cloud Infrastructure Osp Gateway service
---

# Data Source: oci_osp_gateway_invoices
This data source provides the list of Invoices in Oracle Cloud Infrastructure Osp Gateway service.

Returns a list of invoices

## Example Usage

```hcl
data "oci_osp_gateway_invoices" "test_invoices" {
	#Required
	compartment_id = var.compartment_id
	osp_home_region = var.invoice_osp_home_region

	#Optional
	invoice_id = oci_osp_gateway_invoice.test_invoice.id
	search_text = var.invoice_search_text
	status = var.invoice_status
	time_invoice_end = var.invoice_time_invoice_end
	time_invoice_start = var.invoice_time_invoice_start
	time_payment_end = var.invoice_time_payment_end
	time_payment_start = var.invoice_time_payment_start
	type = var.invoice_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `invoice_id` - (Optional) The invoice query param (not unique).
* `osp_home_region` - (Required) The home region's public name of the logged in user. 
* `search_text` - (Optional) A filter to only return resources that match the given value. Looking for partial matches in the following fileds: Invoice No., Reference No. (plan number), Payment Ref, Total Amount(plan number), Balance Due(plan number) and Party/Customer Name 
* `status` - (Optional) A filter to only return resources that match one of the status elements. 
* `time_invoice_end` - (Optional) description: End time (UTC) of the target invoice date range for which to fetch invoice data (exclusive). 
* `time_invoice_start` - (Optional) description: Start time (UTC) of the target invoice date range for which to fetch invoice data (inclusive). 
* `time_payment_end` - (Optional) description: End time (UTC) of the target payment date range for which to fetch invoice data (exclusive). 
* `time_payment_start` - (Optional) description: Start time (UTC) of the target payment date range for which to fetch invoice data (inclusive). 
* `type` - (Optional) A filter to only return resources that match the given type exactly. 


## Attributes Reference

The following attributes are exported:

* `invoice_collection` - The list of invoice_collection.

### Invoice Reference

The following attributes are exported:

* `bill_to_address` - Address details model
	* `address_line1` - Address line 1
	* `address_line2` - Address line 2
	* `address_line3` - Address line 3
	* `address_line4` - Address line 4
	* `city` - Name of the city
	* `company_name` - Name of the customer company
	* `contact_name` - Name of the contact person
	* `country` - Country details model
		* `ascii3country_code` - Country code in ISO-3166-1 3-letter format
		* `country_code` - Country code in ISO-3166-1 2-letter format
		* `country_id` - Indentifier of the country. This is a DB side unique id which was generated when the entity was created in the table
		* `country_name` - Name of the country
		* `language_id` - Language identifier
	* `county` - County name
	* `postal_code` - ZIP no
	* `province` - Name of the province
	* `state` - Name of the state
	* `street_name` - Street name
	* `street_number` - House no
* `currency` - Currency details model
	* `currency_code` - Currency code
	* `currency_symbol` - Currency symbol
	* `name` - Name of the currency
	* `round_decimal_point` - Round decimal point
	* `usd_conversion` - USD conversion rate of the currency
* `internal_invoice_id` - Transaction identifier
* `invoice_amount` - Total amount of invoice
* `invoice_amount_adjusted` - Invoice amount adjust
* `invoice_amount_applied` - Invoice amount applied
* `invoice_amount_credited` - Invoice amount credit
* `invoice_amount_due` - Balance of invoice
* `invoice_id` - Invoice identifier which is generated on the on-premise sie. Pls note this is not an OCID
* `invoice_number` - Invoice external reference
* `invoice_po_number` - Invoice PO number
* `invoice_ref_number` - Invoice reference number
* `invoice_status` - Invoice status
* `invoice_type` - Type of invoice
* `is_credit_card_payable` - Is credit card payment eligible
* `is_display_download_pdf` - Is pdf download access allowed
* `is_payable` - Whether invoice can be payed
* `is_pdf_email_available` - Is emailing pdf allowed
* `last_payment_detail` - Payment related details
	* `amount_paid` - Amount that paid
	* `paid_by` - example
	* `payment_method` - Payment method
	* `time_paid_on` - Paid the invoice on this day
* `payment_terms` - Payment terms
* `preferred_email` - Preferred Email on the invoice
* `subscription_ids` - List of subscription identifiers
* `tax` - Tax of invoice amount
* `time_invoice` - Date of invoice
* `time_invoice_due` - Due date of invoice

