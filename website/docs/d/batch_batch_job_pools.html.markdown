---
subcategory: "Batch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_batch_batch_job_pools"
sidebar_current: "docs-oci-datasource-batch-batch_job_pools"
description: |-
  Provides the list of Batch Job Pools in Oracle Cloud Infrastructure Batch service
---

# Data Source: oci_batch_batch_job_pools
This data source provides the list of Batch Job Pools in Oracle Cloud Infrastructure Batch service.

Lists the batch job pools by compartment or job pool [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). You can filter and sort them by various properties like lifecycle state, display name and also ocid. All properties require an exact match. List operation only provides a summary information, use GetBatchJobPool to get the full details on a specific context

## Example Usage

```hcl
data "oci_batch_batch_job_pools" "test_batch_job_pools" {

	#Optional
	batch_context_id = oci_batch_batch_context.test_batch_context.id
	compartment_id = var.compartment_id
	display_name = var.batch_job_pool_display_name
	id = var.batch_job_pool_id
	state = var.batch_job_pool_state
}
```

## Argument Reference

The following arguments are supported:

* `batch_context_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch context.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job pool.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `batch_job_pool_collection` - The list of batch_job_pool_collection.

### BatchJobPool Reference

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

