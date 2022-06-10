---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_invoice_line_computed_usages"
sidebar_current: "docs-oci-datasource-onesubscription-invoice_line_computed_usages"
description: |-
  Provides the list of Invoice Line Computed Usages in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_invoice_line_computed_usages
This data source provides the list of Invoice Line Computed Usages in Oracle Cloud Infrastructure Onesubscription service.

This is a collection API which returns a list of Invoiced Computed Usages for given Invoiceline id.


## Example Usage

```hcl
data "oci_onesubscription_invoice_line_computed_usages" "test_invoice_line_computed_usages" {
	#Required
	compartment_id = var.compartment_id
	invoice_line_id = oci_onesubscription_invoice_line.test_invoice_line.id

	#Optional
	fields = var.invoice_line_computed_usage_fields
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.
* `fields` - (Optional) Partial response refers to an optimization technique offered by the RESTful web APIs to return only the information  (fields) required by the client. This parameter is used to control what fields to return. 
* `invoice_line_id` - (Required) Invoice Line Identifier - Primary Key SPM 


## Attributes Reference

The following attributes are exported:

* `invoiceline_computed_usages` - The list of invoiceline_computed_usages.

### InvoiceLineComputedUsage Reference

The following attributes are exported:

* `cost` - Sum of Usage/Service Billing Line net Amount 
* `cost_rounded` - Computed Line Amount rounded. 
* `net_unit_price` - Net Unit Price for the product in consideration, price actual. 
* `parent_product` - Product description 
	* `billing_category` - Metered service billing category 
	* `name` - Product name 
	* `part_number` - Product part number 
	* `product_category` - Product category 
	* `ucm_rate_card_part_type` - Rate card part type of Product 
	* `unit_of_measure` - Unit of Measure 
* `product` - Product description 
	* `billing_category` - Metered service billing category 
	* `name` - Product name 
	* `part_number` - Product part number 
	* `product_category` - Product category 
	* `ucm_rate_card_part_type` - Rate card part type of Product 
	* `unit_of_measure` - Unit of Measure 
* `quantity` - Total Quantity that was used for computation 
* `time_metered_on` - Metered Service date. 
* `type` - Usage compute type in SPM. 

