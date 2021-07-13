---
subcategory: "Apm Synthetics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_synthetics_result"
sidebar_current: "docs-oci-datasource-apm_synthetics-result"
description: |-
  Provides details about a specific Result in Oracle Cloud Infrastructure Apm Synthetics service
---

# Data Source: oci_apm_synthetics_result
This data source provides details about a specific Result resource in Oracle Cloud Infrastructure Apm Synthetics service.

Gets the results for a specific execution of a monitor identified by OCID. The results are in a HAR file, Screenshot, Console Log or Network details.


## Example Usage

```hcl
data "oci_apm_synthetics_result" "test_result" {
	#Required
	apm_domain_id = oci_apm_synthetics_apm_domain.test_apm_domain.id
	execution_time = var.result_execution_time
	monitor_id = oci_apm_synthetics_monitor.test_monitor.id
	result_content_type = var.result_result_content_type
	result_type = var.result_result_type
	vantage_point = var.result_vantage_point
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM domain ID the request is intended for. 
* `execution_time` - (Required) The time the object was posted. 
* `monitor_id` - (Required) The OCID of the monitor.
* `result_content_type` - (Required) The result content type zip or raw. 
* `result_type` - (Required) The result type har, screenshot, log or network. 
* `vantage_point` - (Required) The vantagePoint name. 


## Attributes Reference

The following attributes are exported:

* `execution_time` - The specific point of time when the result of an execution is collected.
* `monitor_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitor.
* `result_content_type` - Type of result content. Example: Zip or Raw file. 
* `result_data_set` - Monitor result data set.
	* `byte_content` - Data content in byte format. Example: Zip or Screenshot. 
	* `name` - Name of the data.
	* `string_content` - Data content in string format. Example: HAR. 
	* `timestamp` - The time when the data was generated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 
* `result_type` - Type of result. Example: HAR, Screenshot, Log or Network. 
* `vantage_point` - The name of the vantage point.

