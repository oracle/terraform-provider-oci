---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_applications"
sidebar_current: "docs-oci-datasource-functions-applications"
description: |-
  Provides the list of Applications in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_applications
This data source provides the list of Applications in Oracle Cloud Infrastructure Functions service.

Lists applications for a compartment.

## Example Usage

```hcl
data "oci_functions_applications" "test_applications" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.application_display_name}"
	id = "${var.application_id}"
	state = "${var.application_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this resource belongs. 
* `display_name` - (Optional) A filter to return only applications with display names that match the display name string. Matching is exact. 
* `id` - (Optional) A filter to return only applications with the specfied OCID. 
* `state` - (Optional) A filter to return only applications that match the lifecycle state in this parameter. Example: `Creating` 


## Attributes Reference

The following attributes are exported:

* `applications` - The list of applications.

### Application Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the application. 
* `config` - Application configuration for functions in this application (passed as environment variables). Can be overridden by function configuration. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`

	The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of the application. The display name is unique within the compartment containing the application. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application. 
* `state` - The current state of the application. 
* `subnet_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the subnets in which to run functions in the application. 
* `time_created` - The time the application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-09-12T22:47:12.613Z` 
* `time_updated` - The time the application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-09-12T22:47:12.613Z` 

