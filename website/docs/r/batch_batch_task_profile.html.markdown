---
subcategory: "Batch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_batch_batch_task_profile"
sidebar_current: "docs-oci-resource-batch-batch_task_profile"
description: |-
  Provides the Batch Task Profile resource in Oracle Cloud Infrastructure Batch service
---

# oci_batch_batch_task_profile
This resource provides the Batch Task Profile resource in Oracle Cloud Infrastructure Batch service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/batch

Creates a batch task profile.

## Example Usage

```hcl
resource "oci_batch_batch_task_profile" "test_batch_task_profile" {
	#Required
	compartment_id = var.compartment_id
	min_memory_in_gbs = var.batch_task_profile_min_memory_in_gbs
	min_ocpus = var.batch_task_profile_min_ocpus

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.batch_task_profile_description
	display_name = var.batch_task_profile_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The batch task profile description.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. If not specified or provided as null or empty string, it be generated as "<resourceType><timeCreated>", where timeCreated corresponds with the resource creation time in ISO 8601 basic format, i.e. omitting separating punctuation, at second-level precision and no UTC offset. Example: batchtaskprofile20250914115623. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `min_memory_in_gbs` - (Required) The minimum required memory.
* `min_ocpus` - (Required) The minimum required OCPUs.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The batch task profile description.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task profile.
* `min_memory_in_gbs` - The minimum required memory.
* `min_ocpus` - The minimum required OCPUs.
* `state` - The current state of the batch task profile. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the batch task profile was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the batch task profile was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Batch Task Profile
	* `update` - (Defaults to 20 minutes), when updating the Batch Task Profile
	* `delete` - (Defaults to 20 minutes), when destroying the Batch Task Profile


## Import

BatchTaskProfiles can be imported using the `id`, e.g.

```
$ terraform import oci_batch_batch_task_profile.test_batch_task_profile "id"
```

