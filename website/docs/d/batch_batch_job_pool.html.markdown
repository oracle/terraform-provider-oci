---
subcategory: "Batch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_batch_batch_job_pool"
sidebar_current: "docs-oci-datasource-batch-batch_job_pool"
description: |-
  Provides details about a specific Batch Job Pool in Oracle Cloud Infrastructure Batch service
---

# Data Source: oci_batch_batch_job_pool
This data source provides details about a specific Batch Job Pool resource in Oracle Cloud Infrastructure Batch service.

Gets information about a batch job pool.

## Example Usage

```hcl
data "oci_batch_batch_job_pool" "test_batch_job_pool" {
	#Required
	batch_job_pool_id = oci_batch_batch_job_pool.test_batch_job_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `batch_job_pool_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job pool.


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

