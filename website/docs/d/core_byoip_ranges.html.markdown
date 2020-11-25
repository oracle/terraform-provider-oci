---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_byoip_ranges"
sidebar_current: "docs-oci-datasource-core-byoip_ranges"
description: |-
  Provides the list of Byoip Ranges in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_byoip_ranges
This data source provides the list of Byoip Ranges in Oracle Cloud Infrastructure Core service.

Lists the `ByoipRange` resources in the specified compartment.
You can filter the list using query parameters.


## Example Usage

```hcl
data "oci_core_byoip_ranges" "test_byoip_ranges" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.byoip_range_display_name
	state = var.byoip_range_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state name exactly. 


## Attributes Reference

The following attributes are exported:

* `byoip_range_collection` - The list of byoip_range_collection.

### ByoipRange Reference

The following attributes are exported:

* `cidr_block` - The public IPv4 CIDR block being imported from on-premises to the Oracle cloud.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the BYOIP CIDR block. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `ByoipRange` resource.
* `lifecycle_details` - The `ByoipRange` resource's current status.
* `state` - The `ByoipRange` resource's current state.
* `time_advertised` - The date and time the `ByoipRange` resource was advertised to the internet by BGP, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_created` - The date and time the `ByoipRange` resource was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_validated` - The date and time the `ByoipRange` resource was validated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_withdrawn` - The date and time the `ByoipRange` resource was withdrawn from advertisement by BGP to the internet, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `validation_token` - The validation token is an internally-generated ASCII string used in the validation process. See [Importing a CIDR block](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/BYOIP.htm#import_cidr) for details.

