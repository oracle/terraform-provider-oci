---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_letter_of_authority"
sidebar_current: "docs-oci-datasource-core-letter_of_authority"
description: |-
  Provides details about a specific Letter Of Authority in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_letter_of_authority
This data source provides details about a specific Letter Of Authority resource in Oracle Cloud Infrastructure Core service.

Gets the Letter of Authority for the specified cross-connect.

## Example Usage

```hcl
data "oci_core_letter_of_authority" "test_letter_of_authority" {
	#Required
	cross_connect_id = oci_core_cross_connect.test_cross_connect.id
}
```

## Argument Reference

The following arguments are supported:

* `cross_connect_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect.


## Attributes Reference

The following attributes are exported:

* `authorized_agent` - Name of a customer authorized agent which will be appended to the LOA as 'Authorized Agent'.
* `authorized_entity_name` - The name of the entity authorized by this Letter of Authority.
* `circuit_type` - The type of cross-connect fiber, termination, and optical specification.
* `cross_connect_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect.
* `extension_details` - Data related to the extension of the Expiry date of the LOA. It gives you number of remaining extensions along with a history of past extensions made on the LOA. 
	* `history` - Chronologically sorted list of date and time when the Letter of Authority's expiration was last updated,  most recent first, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). List is empty if the LOA's expiration date has never been extended. 
	* `remaining_extensions` - The number of self-service LOA expiry extensions still available.
* `facility_location` - The address of the FastConnect location.
* `port_name` - The meet-me room port for this cross-connect.
* `time_expires` - The date and time when the Letter of Authority expires, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_issued` - The date and time the Letter of Authority was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
