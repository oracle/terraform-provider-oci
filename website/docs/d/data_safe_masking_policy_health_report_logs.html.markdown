---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policy_health_report_logs"
sidebar_current: "docs-oci-datasource-data_safe-masking_policy_health_report_logs"
description: |-
  Provides the list of Masking Policy Health Report Logs in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_policy_health_report_logs
This data source provides the list of Masking Policy Health Report Logs in Oracle Cloud Infrastructure Data Safe service.

Gets a list of errors and warnings from a masking policy health check.


## Example Usage

```hcl
data "oci_data_safe_masking_policy_health_report_logs" "test_masking_policy_health_report_logs" {
	#Required
	masking_policy_health_report_id = oci_data_safe_masking_policy_health_report.test_masking_policy_health_report.id

	#Optional
	message_type = var.masking_policy_health_report_log_message_type
}
```

## Argument Reference

The following arguments are supported:

* `masking_policy_health_report_id` - (Required) The OCID of the masking health report.
* `message_type` - (Optional) A filter to return only the resources that match the specified log message type.


## Attributes Reference

The following attributes are exported:

* `masking_policy_health_report_log_collection` - The list of masking_policy_health_report_log_collection.

### MaskingPolicyHealthReportLog Reference

The following attributes are exported:

* `items` - An array of masking policy health report objects.
	* `description` - A human-readable description for the log entry.
	* `message` - A human-readable log entry.
	* `message_type` - The log entry type.
	* `remediation` - A human-readable log entry to remedy any error or warnings in the masking policy.
	* `timestamp` - The date and time the log entry was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

