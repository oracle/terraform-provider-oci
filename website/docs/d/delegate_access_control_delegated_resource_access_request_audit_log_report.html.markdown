---
subcategory: "Delegate Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_delegate_access_control_delegated_resource_access_request_audit_log_report"
sidebar_current: "docs-oci-datasource-delegate_access_control-delegated_resource_access_request_audit_log_report"
description: |-
  Provides details about a specific Delegated Resource Access Request Audit Log Report in Oracle Cloud Infrastructure Delegate Access Control service
---

# Data Source: oci_delegate_access_control_delegated_resource_access_request_audit_log_report
This data source provides details about a specific Delegated Resource Access Request Audit Log Report resource in Oracle Cloud Infrastructure Delegate Access Control service.

Gets the audit log report for the given Delegated Resource Access Request.

## Example Usage

```hcl
data "oci_delegate_access_control_delegated_resource_access_request_audit_log_report" "test_delegated_resource_access_request_audit_log_report" {
	#Required
	delegated_resource_access_request_id = oci_delegate_access_control_delegated_resource_access_request.test_delegated_resource_access_request.id

	#Optional
	is_process_tree_enabled = var.delegated_resource_access_request_audit_log_report_is_process_tree_enabled
}
```

## Argument Reference

The following arguments are supported:

* `delegated_resource_access_request_id` - (Required) Unique Delegated Resource Access Request identifier
* `is_process_tree_enabled` - (Optional) Set to true to enable process tree computation in audit report


## Attributes Reference

The following attributes are exported:

* `audit_report_status` - Status of the audit report
* `process_tree` - The process tree data
* `report` - Audit log report.
* `time_report_generated` - Time when the audit report was generated [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

