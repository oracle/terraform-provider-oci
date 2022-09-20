---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_trail_sequence"
sidebar_current: "docs-oci-datasource-golden_gate-trail_sequence"
description: |-
  Provides details about a specific Trail Sequence in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_trail_sequence
This data source provides details about a specific Trail Sequence resource in Oracle Cloud Infrastructure Golden Gate service.

Lists the Trail Sequences for a TrailFile in a given deployment.


## Example Usage

```hcl
data "oci_golden_gate_trail_sequence" "test_trail_sequence" {
	#Required
	deployment_id = oci_golden_gate_deployment.test_deployment.id
	trail_file_id = oci_golden_gate_trail_file.test_trail_file.id
	display_name = var.trail_sequence_display_name
	trail_sequence_id = oci_golden_gate_trail_sequence.test_trail_sequence.id
	#Optional
}
```

## Argument Reference

The following arguments are supported:

* `deployment_id` - (Required) A unique Deployment identifier. 
* `display_name` - (Required) A filter to return only the resources that match the entire 'displayName' given. 
* `trail_file_id` - (Required) A Trail File identifier 
* `trail_sequence_id` - (Required) A Trail Sequence identifier 


## Attributes Reference

The following attributes are exported:

* `items` - An array of TrailSequences. 
	* `display_name` - An object's Display Name. 
	* `sequence_id` - Sequence Id 
	* `size_in_bytes` - The size of the backup stored in object storage (in bytes) 
	* `time_last_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_last_fetched` - The time the data was last fetched from the deployment. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

