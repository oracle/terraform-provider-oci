---
subcategory: "Logging"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_logging_log_saved_search"
sidebar_current: "docs-oci-datasource-logging-log_saved_search"
description: |-
  Provides details about a specific Log Saved Search in Oracle Cloud Infrastructure Logging service
---

# Data Source: oci_logging_log_saved_search
This data source provides details about a specific Log Saved Search resource in Oracle Cloud Infrastructure Logging service.

Retrieves a log saved search.

## Example Usage

```hcl
data "oci_logging_log_saved_search" "test_log_saved_search" {
	#Required
	log_saved_search_id = oci_logging_log_saved_search.test_log_saved_search.id
}
```

## Argument Reference

The following arguments are supported:

* `log_saved_search_id` - (Required) OCID of the logSavedSearch 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that the resource belongs to.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description for this resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the resource.
* `name` - The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information. 
* `query` - The search query that is saved. 
* `state` - The state of the LogSavedSearch 
* `time_created` - Time the resource was created.
* `time_last_modified` - Time the resource was last modified.

