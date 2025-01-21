---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_byoasns"
sidebar_current: "docs-oci-datasource-core-byoasns"
description: |-
  Provides the list of Byoasns in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_byoasns
This data source provides the list of Byoasns in Oracle Cloud Infrastructure Core service.

Lists the `Byoasn` resources in the specified compartment.
You can filter the list using query parameters.


## Example Usage

```hcl
data "oci_core_byoasns" "test_byoasns" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.byoasn_display_name
	state = var.byoasn_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state name exactly. 


## Attributes Reference

The following attributes are exported:

* `byoasn_collection` - The list of byoasn_collection.

### Byoasn Reference

The following attributes are exported:

* `asn` - The Autonomous System Number (ASN) you are importing to the Oracle cloud.
* `byoip_ranges` - The BYOIP Ranges that has the `Byoasn` as origin.
	* `as_path_prepend_length` - The as path prepend length.
	* `byoip_range_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `ByoipRange` resource to which the CIDR block belongs.
	* `cidr_block` - The BYOIP CIDR block range or subrange allocated to an IP pool. This could be all or part of a BYOIP CIDR block.
	* `ipv6cidr_block` - The IPv6 prefix being imported to the Oracle cloud. This prefix must be /48 or larger, and can  be subdivided into sub-ranges used across multiple VCNs. A BYOIPv6 prefix can be assigned across multiple VCNs, and each VCN must be /64 or larger. You may specify a ULA or private IPv6 prefix of /64 or larger to use in the VCN. IPv6-enabled subnets will remain a fixed /64 in size. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the `Byoasn` resource. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `Byoasn` resource.
* `state` - The `Byoasn` resource's current state.
* `time_created` - The date and time the `Byoasn` resource was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the `Byoasn` resource was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_validated` - The date and time the `Byoasn` resource was validated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `validation_token` - The validation token is an internally-generated ASCII string used in the validation process. See [Importing a Byoasn](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/BYOASN.htm) for details.

