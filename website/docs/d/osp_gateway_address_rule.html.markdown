---
subcategory: "Osp Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osp_gateway_address_rule"
sidebar_current: "docs-oci-datasource-osp_gateway-address_rule"
description: |-
  Provides details about a specific Address Rule in Oracle Cloud Infrastructure Osp Gateway service
---

# Data Source: oci_osp_gateway_address_rule
This data source provides details about a specific Address Rule resource in Oracle Cloud Infrastructure Osp Gateway service.

Get the address rule for the compartment based on the country code

## Example Usage

```hcl
data "oci_osp_gateway_address_rule" "test_address_rule" {
	#Required
	compartment_id = var.compartment_id
	country_code = var.address_rule_country_code
	osp_home_region = var.address_rule_osp_home_region
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `country_code` - (Required) Country code for the address rule in ISO-3166-1 2-letter format. 
* `osp_home_region` - (Required) The home region's public name of the logged in user. 


## Attributes Reference

The following attributes are exported:

* `address` - Address type rule information
	* `fields` - Address type rule fields
		* `format` - Format information
			* `example` - Example of the desired format that matches the regex
			* `value` - Regex format specification
		* `is_required` - The given field is requeired or not
		* `label` - Label information
			* `example` - English translation of the label (for reference only - translation is not provided)
			* `value` - Language token of the required label
		* `language` - Locale code (rfc4646 format) of a forced language (e.g.: jp addresses require jp always)
		* `name` - The field name
	* `third_party_validation` - Third party validation.
* `contact` - Contact type rule information
	* `fields` - Contact type rule fields
		* `format` - Format information
			* `example` - Example of the desired format that matches the regex
			* `value` - Regex format specification
		* `is_required` - The given field is requeired or not
		* `label` - Label information
			* `example` - English translation of the label (for reference only - translation is not provided)
			* `value` - Language token of the required label
		* `language` - Locale code (rfc4646 format) of a forced language (e.g.: jp addresses require jp always)
		* `name` - The field name
* `country_code` - Country code for the address rule in ISO-3166-1 2-letter format
* `tax` - Tax type rule information
	* `fields` - Tax type rule fields
		* `format` - Format information
			* `example` - Example of the desired format that matches the regex
			* `value` - Regex format specification
		* `is_required` - The given field is requeired or not
		* `label` - Label information
			* `example` - English translation of the label (for reference only - translation is not provided)
			* `value` - Language token of the required label
		* `language` - Locale code (rfc4646 format) of a forced language (e.g.: jp addresses require jp always)
		* `name` - The field name
	* `value_set` - Label value pair for allowed values. Used for GIRO
		* `name` - User friendly name
		* `value` - Value

