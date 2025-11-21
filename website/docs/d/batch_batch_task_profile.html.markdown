---
subcategory: "Batch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_batch_batch_task_profile"
sidebar_current: "docs-oci-datasource-batch-batch_task_profile"
description: |-
  Provides details about a specific Batch Task Profile in Oracle Cloud Infrastructure Batch service
---

# Data Source: oci_batch_batch_task_profile
This data source provides details about a specific Batch Task Profile resource in Oracle Cloud Infrastructure Batch service.

Gets information about a batch task profile.

## Example Usage

```hcl
data "oci_batch_batch_task_profile" "test_batch_task_profile" {
	#Required
	batch_task_profile_id = oci_batch_batch_task_profile.test_batch_task_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `batch_task_profile_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task profile.


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

