---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_byoasn"
sidebar_current: "docs-oci-resource-core-byoasn"
description: |-
  Provides the Byoasn resource in Oracle Cloud Infrastructure Core service
---

# oci_core_byoasn
This resource provides the Byoasn resource in Oracle Cloud Infrastructure Core service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iaas/latest/Byoasn

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates a BYOASN Resource


## Example Usage

```hcl
resource "oci_core_byoasn" "test_byoasn" {
	#Required
	asn = var.byoasn_asn
	compartment_id = var.compartment_id
	display_name = var.byoasn_display_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `asn` - (Required) The Autonomous System Number (ASN) you are importing to the Oracle cloud. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the BYOASN Resource. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Byoasn
	* `update` - (Defaults to 20 minutes), when updating the Byoasn
	* `delete` - (Defaults to 20 minutes), when destroying the Byoasn


## Import

Byoasns can be imported using the `id`, e.g.

```
$ terraform import oci_core_byoasn.test_byoasn "id"
```

