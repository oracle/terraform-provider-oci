---
subcategory: "Zpr"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_zpr_configuration"
sidebar_current: "docs-oci-resource-zpr-configuration"
description: |-
  Provides the Configuration resource in Oracle Cloud Infrastructure Zpr service
---

# oci_zpr_configuration
This resource provides the Configuration resource in Oracle Cloud Infrastructure Zpr service.

Initiates the process to onboard ZPR
in a root compartment (the root compartment is the tenancy). It creates an object of ZPR configuration as part of onboarding.


## Example Usage

```hcl
resource "oci_zpr_configuration" "test_configuration" {
	#Required
	compartment_id = var.tenancy_ocid

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	zpr_status = var.configuration_zpr_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy into which ZPR resources will be bootstrapped. 
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `zpr_status` - (Optional) The enabled or disabled status of ZPR in the tenancy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy into which ZPR will be onboarded. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ZprConfiguration.
* `lifecycle_details` - A message that describes the current state of ZPR in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `state` - The current state of ZPR in the tenancy.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that ZPR was onboarded to the tenancy, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time that ZPR was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `zpr_status` - The enabled or disabled status of ZPR in tenancy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Configuration
	* `update` - (Defaults to 20 minutes), when updating the Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Configuration


## Import

Configuration can be imported using the 'compartment_id' and `id`, e.g.

```
$ terraform import oci_zpr_configuration.test_configuration "{compartment_id}/{id}"
```

