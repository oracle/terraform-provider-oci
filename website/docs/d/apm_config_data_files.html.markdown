---
subcategory: "Apm Config"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_config_data_files"
sidebar_current: "docs-oci-datasource-apm_config-data_files"
description: |-
  Provides the list of Data Files in Oracle Cloud Infrastructure Apm Config service
---

# Data Source: oci_apm_config_data_files
This data source provides the list of Data Files in Oracle Cloud Infrastructure Apm Config service.

Fetches a list of Data files using some parameters.


## Example Usage

```hcl
data "oci_apm_config_data_files" "test_data_files" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

	#Optional
	apm_type = var.data_file_apm_type
	metadata = var.data_file_metadata
	name = var.data_file_name
	time_last_modified_after = var.data_file_time_last_modified_after
	time_last_modified_before = var.data_file_time_last_modified_before
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID the request is intended for. 
* `apm_type` - (Optional) The type of the data file.
* `metadata` - (Optional) Optional user-defined metadata key and value to search by. 
* `name` - (Optional) A filter to return resources that match the specified name. Supports regular expressions to filter data files.
* `time_last_modified_after` - (Optional) Return data files with the 'timeLastModified' after the specified time, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-19T22:47:12.613Z` 
* `time_last_modified_before` - (Optional) Return data files with time 'timeLastModified' before the specified time, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-19T22:47:12.613Z` 


## Attributes Reference

The following attributes are exported:

* `data_file_summary_collection` - The list of data_file_summary_collection.

### DataFile Reference

The following attributes are exported:


