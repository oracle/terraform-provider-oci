---
subcategory: "Batch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_batch_batch_job_pool"
sidebar_current: "docs-oci-resource-batch-batch_job_pool"
description: |-
  Provides the Batch Job Pool resource in Oracle Cloud Infrastructure Batch service
---

# oci_batch_batch_job_pool
This resource provides the Batch Job Pool resource in Oracle Cloud Infrastructure Batch service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/batch

Creates a batch job pool.

## Example Usage

```hcl
resource "oci_batch_batch_job_pool" "test_batch_job_pool" {
	#Required
	batch_context_id = oci_batch_batch_context.test_batch_context.id
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.batch_job_pool_description
	display_name = var.batch_job_pool_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `batch_context_id` - (Required) The OCID of batch context.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Summarized information about the batch job pool.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. If not specified or provided as null or empty string, it be generated as "<resourceType><timeCreated>", where timeCreated corresponds with the resource creation time in ISO 8601 basic format, i.e. omitting separating punctuation, at second-level precision and no UTC offset. Example: batchjobpool20250914115623. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `state` - (Optional) (Updatable) The target state for the Batch Job Pool. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `batch_context_id` - The OCID of batch context.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Summarized information about the batch job pool.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job pool.
* `state` - The current state of the batch job pool. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the batch job pool was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the batch job pool was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Batch Job Pool
	* `update` - (Defaults to 20 minutes), when updating the Batch Job Pool
	* `delete` - (Defaults to 20 minutes), when destroying the Batch Job Pool


## Import

BatchJobPools can be imported using the `id`, e.g.

```
$ terraform import oci_batch_batch_job_pool.test_batch_job_pool "id"
```

