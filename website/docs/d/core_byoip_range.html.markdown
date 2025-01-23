---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_byoip_range"
sidebar_current: "docs-oci-datasource-core-byoip_range"
description: |-
  Provides details about a specific Byoip Range in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_byoip_range
This data source provides details about a specific Byoip Range resource in Oracle Cloud Infrastructure Core service.

Gets the `ByoipRange` resource. You must specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_core_byoip_range" "test_byoip_range" {
	#Required
	byoip_range_id = oci_core_byoip_range.test_byoip_range.id
}
```

## Argument Reference

The following arguments are supported:

* `byoip_range_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `ByoipRange` resource containing the BYOIP CIDR block.


## Attributes Reference

The following attributes are exported:

* `byoip_range_vcn_ipv6allocations` - A list of `ByoipRangeVcnIpv6AllocationSummary` objects. 
	* `byoip_range_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `ByoipRange` resource to which the CIDR block belongs.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the `ByoipRange`. 
	* `ipv6cidr_block` - The BYOIPv6 prefix range or subrange allocated to a VCN. This could be all or part of a BYOIPv6 prefix. Each VCN allocation must be /64 or larger. 
	* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `Vcn` resource to which the ByoipRange belongs. 
* `cidr_block` - The public IPv4 CIDR block being imported from on-premises to the Oracle cloud.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the BYOIP CIDR block. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `ByoipRange` resource.
* `ipv6cidr_block` - The IPv6 prefix being imported to the Oracle cloud. This prefix must be /48 or larger, and can be subdivided into sub-ranges used across multiple VCNs. A BYOIPv6 prefix can be also assigned across multiple VCNs, and each VCN must be /64 or larger. You may specify a ULA or private IPv6 prefix of /64 or larger to use in the VCN. IPv6-enabled subnets will remain a fixed /64 in size. 
* `lifecycle_details` - The `ByoipRange` resource's current status.
* `origin_asn` - Information about the origin asn.
	* `as_path_prepend_length` - The as path prepend length.
	* `asn` - The Autonomous System Number (ASN) you are importing to the Oracle cloud.
	* `byoasn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `Byoasn` resource.
* `state` - The `ByoipRange` resource's current state.
* `time_advertised` - The date and time the `ByoipRange` resource was advertised to the internet by BGP, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_created` - The date and time the `ByoipRange` resource was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_validated` - The date and time the `ByoipRange` resource was validated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_withdrawn` - The date and time the `ByoipRange` resource was withdrawn from advertisement by BGP to the internet, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `validation_token` - The validation token is an internally-generated ASCII string used in the validation process. See [Importing a CIDR block](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/BYOIP.htm#import_cidr) for details.

