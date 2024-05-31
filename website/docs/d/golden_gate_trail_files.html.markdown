---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_trail_files"
sidebar_current: "docs-oci-datasource-golden_gate-trail_files"
description: |-
  Provides the list of Trail Files in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_trail_files
This data source provides the list of Trail Files in Oracle Cloud Infrastructure Golden Gate service.

Lists the TrailFiles for a deployment. Deprecated: Please access trail file management functions directly on OGG console which are available since version Oracle GoldenGate 23c.


## Example Usage

```hcl
data "oci_golden_gate_trail_files" "test_trail_files" {
	#Required
	deployment_id = oci_golden_gate_deployment.test_deployment.id
	trail_file_id = oci_golden_gate_trail_file.test_trail_file.id

	#Optional
	display_name = var.trail_file_display_name
}
```

## Argument Reference

The following arguments are supported:

* `deployment_id` - (Required) A unique Deployment identifier. 
* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `trail_file_id` - (Required) A Trail File identifier 


## Attributes Reference

The following attributes are exported:

* `trail_file_collection` - The list of trail_file_collection.

### TrailFile Reference

The following attributes are exported:

* `items` - An array of TrailFiles. 
	* `consumers` - array of consumer process names 
	* `display_name` - An object's Display Name. 
	* `max_sequence_number` - Maximum sequence number 
	* `min_sequence_number` - Minimum sequence number 
	* `number_of_sequences` - Number of sequences for a specific trail file 
	* `producer` - Producer Process Name if any. 
	* `size_in_bytes` - The size of the backup stored in object storage (in bytes) 
	* `time_last_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
	* `trail_file_id` - The TrailFile Id. 
* `time_last_fetched` - The time the data was last fetched from the deployment. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

