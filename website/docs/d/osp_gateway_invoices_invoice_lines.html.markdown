---
subcategory: "Osp Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osp_gateway_invoices_invoice_lines"
sidebar_current: "docs-oci-datasource-osp_gateway-invoices_invoice_lines"
description: |-
  Provides the list of Invoices Invoice Lines in Oracle Cloud Infrastructure Osp Gateway service
---

# Data Source: oci_osp_gateway_invoices_invoice_lines
This data source provides the list of Invoices Invoice Lines in Oracle Cloud Infrastructure Osp Gateway service.

Returns the invoice product list by invoice id

## Example Usage

```hcl
data "oci_osp_gateway_invoices_invoice_lines" "test_invoices_invoice_lines" {
	#Required
	compartment_id = var.compartment_id
	internal_invoice_id = oci_osp_gateway_invoice.test_invoice.id
	osp_home_region = var.invoices_invoice_line_osp_home_region
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `internal_invoice_id` - (Required) The identifier of the invoice.
* `osp_home_region` - (Required) The home region's public name of the logged in user. 


## Attributes Reference

The following attributes are exported:

* `invoice_line_collection` - The list of invoice_line_collection.

### InvoicesInvoiceLine Reference

The following attributes are exported:

* `items` - Invoice line list elements
	* `currency` - Currency details model
		* `currency_code` - Currency code
		* `currency_symbol` - Currency symbol
		* `name` - Name of the currency
		* `round_decimal_point` - Round decimal point
		* `usd_conversion` - USD conversion rate of the currency
	* `net_unit_price` - Unit price of the ordered product
	* `order_no` - Product of the item
	* `part_number` - Part number
	* `product` - Product of the item
	* `quantity` - Quantity of the ordered product
	* `time_end` - End date
	* `time_start` - Start date
	* `total_price` - Total price of the ordered product (Net unit price x quantity)

