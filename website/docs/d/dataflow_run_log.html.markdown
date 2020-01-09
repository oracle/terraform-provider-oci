---
subcategory: "Dataflow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_run_log"
sidebar_current: "docs-oci-datasource-dataflow-run_log"
description: |-
  Provides details about a specific Run Log in Oracle Cloud Infrastructure Dataflow service
---

# Data Source: oci_dataflow_run_log
This data source provides details about a specific Run Log resource in Oracle Cloud Infrastructure Dataflow service.

Retrieves the content of an run log.


## Example Usage

```hcl
data "oci_dataflow_run_log" "test_run_log" {
	#Required
	name = "${var.run_log_name}"
	run_id = "${oci_dataflow_run.test_run.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the log. Avoid entering confidential information. 
* `run_id` - (Required) The unique ID for the run 
* `base64_encode_content` - (Optional) Encodes the downloaded content in base64. It is recommended to set this to `true` for binary content to avoid corrupting the zip file in Terraform state. The default value is `false` to preserve backwards compatibility with Terraform v0.11 configurations.
If passing the base64 encoded content to a `local_file` resource, please use the `content_base64` attribute of the `local_file` resource.

## Attributes Reference

The following attributes are exported:

* `content` - The content of the run log.
* `content_type` - The content type of the run log.

