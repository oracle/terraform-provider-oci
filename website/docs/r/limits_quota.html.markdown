---
subcategory: "Limits"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_limits_quota"
sidebar_current: "docs-oci-resource-limits-quota"
description: |-
  Provides the Quota resource in Oracle Cloud Infrastructure Limits service
---

# oci_limits_quota
This resource provides the Quota resource in Oracle Cloud Infrastructure Limits service.

Creates a new quota with the details supplied.

## Example Usage

```hcl
resource "oci_limits_quota" "test_quota" {
	#Required
	compartment_id = var.tenancy_ocid
	description = var.quota_description
	name = var.quota_name
	statements = var.quota_statements

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	locks {
		#Required
		type = var.quota_locks_type

		#Optional
		message = var.quota_locks_message
		related_resource_id = oci_limits_related_resource.test_related_resource.id
	}	
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the resource this quota applies to.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the quota.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the lock creator. The message typically gives an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The resource ID that is locking this resource. Indicates that deleting this resource removes the lock. 
	* `type` - (Required) Lock type.
* `name` - (Required) The name you assign to the quota during creation. The name must be unique across all quotas in the tenancy and cannot be changed. 
* `statements` - (Required) (Updatable) An array of quota statements written in the declarative quota statement language. 

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the resource this quota applies to. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the quota.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the quota.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the lock creator. The message typically gives an indication of why the resource is locked. 
	* `related_resource_id` - The resource ID that is locking this resource. Indicates that deleting this resource removes the lock. 
	* `time_created` - Indicates when the lock was created, in the format defined by RFC 3339.
	* `type` - Lock type.
* `name` - The name you assign to the quota during creation. The name must be unique across all quotas in the tenancy and cannot be changed. 
* `state` - The quota's current state.
* `statements` - An array of one or more quota statements written in the declarative quota statement language.
* `time_created` - Date and time the quota was created, in the format defined by RFC 3339. Example: `2016-08-25T21:10:29.600Z` 
* `is_lock_override` - this is a computed field which is set to true if any lock is present` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Quota
	* `update` - (Defaults to 20 minutes), when updating the Quota
	* `delete` - (Defaults to 20 minutes), when destroying the Quota


## Import

Quotas can be imported using the `id`, e.g.

```
$ terraform import oci_limits_quota.test_quota "id"
```

