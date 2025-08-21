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
	* `health_check_type` - An enum type entry for each health check in the masking policy. Each enum describes a type of health check. INVALID_OBJECT_CHECK checks if there exist any invalid objects in the masking tables. PRIVILEGE_CHECK checks if the masking user has sufficient privilege to run masking. TABLESPACE_CHECK checks if the user has sufficient default and TEMP tablespace. DATABASE_OR_SYSTEM_TRIGGERS_CHECK checks if there exist any database/system triggers available. UNDO_TABLESPACE_CHECK checks if the AUTOEXTEND feature is enabled for the undo tablespace. If it's not enabled, it further checks if the undo tablespace has any space remaining STATE_STATS_CHECK checks if all the statistics of the masking table is upto date or not. OLS_POLICY_CHECK , VPD_POLICY_CHECK and REDACTION_POLICY_CHECK checks if the masking tables has Oracle Label Security (OLS) or Virtual Private Database (VPD) or Redaction policies enabled. DV_ENABLE_CHECK checks if database has Database Vault(DV) enabled DE_COL_SIZE_CHECK checks if any masking column with DETERMINISTIC ENCRYPTION as masking format has average column size greater than 27 or not. ACTIVE_MASK_JOB_CHECK checks if there is any active masking job running on the target database. DETERMINISTIC_ENCRYPTION_FORMAT_CHECK checks if any masking column has deterministic encryption masking format. COLUMN_EXIST_CHECK checks if the masking columns are available in the target database. TIME_TRAVEL_CHECK checks if the masking tables have Time Travel enabled. 
	* `message` - A human-readable log entry.
	* `message_type` - The log entry type.
	* `remediation` - A human-readable log entry to remedy any error or warnings in the masking policy.
	* `timestamp` - The date and time the log entry was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

