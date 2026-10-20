---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_health_report_logs"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_policy_health_report_logs"
description: |-
  Provides the list of Subsetting Policy Health Report Logs in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_policy_health_report_logs
This data source provides the list of Subsetting Policy Health Report Logs in Oracle Cloud Infrastructure Data Safe service.

Gets a list of errors and warnings from a subsetting policy health check.


## Example Usage

```hcl
data "oci_data_safe_subsetting_policy_health_report_logs" "test_subsetting_policy_health_report_logs" {
	#Required
	subsetting_policy_health_report_id = oci_data_safe_subsetting_policy_health_report.test_subsetting_policy_health_report.id

	#Optional
	message_type = var.subsetting_policy_health_report_log_message_type
}
```

## Argument Reference

The following arguments are supported:

* `message_type` - (Optional) A filter to return only the resources that match the specified log message type.
* `subsetting_policy_health_report_id` - (Required) The OCID of the subsetting health report.


## Attributes Reference

The following attributes are exported:

* `subsetting_policy_health_report_log_collection` - The list of subsetting_policy_health_report_log_collection.

### SubsettingPolicyHealthReportLog Reference

The following attributes are exported:

* `items` - An array of subsetting policy health report objects.
	* `description` - A human-readable description for the log entry.
	* `health_check_type` - An enum type entry for each health check in the subsetting policy. Each enum describes a type of health check. INVALID_OBJECT_CHECK checks if there exist any invalid objects in the subsetting tables. PRIVILEGE_CHECK checks if the subsetting user has sufficient privilege to run subsetting. TABLESPACE_CHECK checks if the user has sufficient default and TEMP tablespace. Also verifies that the specified tablespace by the user is valid, if user has provided one DATABASE_OR_SYSTEM_TRIGGERS_CHECK checks if there exist any database/system triggers available. UNDO_TABLESPACE_CHECK checks if for all the instances of undo tablespace the AUTOEXTEND feature is enabled.  If it's not enabled, it further checks if the undo tablespace has any space remaining. STATE_STATS_CHECK checks if all the statistics of the subsetting table is upto date or not. OLS_POLICY_CHECK , VPD_POLICY_CHECK and REDACTION_POLICY_CHECK checks if the subsetting tables has Oracle Label Security (OLS) or Virtual Private Database (VPD) or Redaction policies enabled. DV_ENABLE_CHECK checks if database has Database Vault(DV) enabled ACTIVE_JOB_CHECK checks if there is any active subsetting job running on the target database. TABLE_EXIST_CHECK checks if the subsetting tables are available in the target database. TIME_TRAVEL_CHECK checks if the subsetting tables have Time Travel enabled. SYSTEM_OBJECTS_CHECK checks if the subsetting tables have dependent objects present in SYS schema. INVALID_PACKAGE_CHECK checks if any of the required packages are in invalid state. AUDIT_POLICY_CHECK checks if the subsetting tables have Audit policies enabled. VALID_RULES_CHECK if the subsetting rules on the tables are valid. 
	* `message` - A human-readable log entry.
	* `message_type` - The log entry type.
	* `remediation` - A human-readable log entry to remedy any error or warnings in the subsetting policy.
	* `timestamp` - The date and time the log entry was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

