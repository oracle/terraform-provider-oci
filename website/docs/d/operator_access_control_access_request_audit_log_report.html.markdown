---
subcategory: "Operator Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_operator_access_control_access_request_audit_log_report"
sidebar_current: "docs-oci-datasource-operator_access_control-access_request_audit_log_report"
description: |-
  Provides details about a specific Access Request Audit Log Report in Oracle Cloud Infrastructure Operator Access Control service
---

# Data Source: oci_operator_access_control_access_request_audit_log_report
This data source provides details about a specific Access Request Audit Log Report resource in Oracle Cloud Infrastructure Operator Access Control service.

Gets the Audit Log Report for the given access requestId.

## Example Usage

```hcl
data "oci_operator_access_control_access_request_audit_log_report" "test_access_request_audit_log_report" {
	#Required
	access_request_id = oci_operator_access_control_access_request.test_access_request.id

	#Optional
	enable_process_tree = var.access_request_audit_log_report_enable_process_tree
}
```

## Argument Reference

The following arguments are supported:

* `access_request_id` - (Required) unique AccessRequest identifier
* `enable_process_tree` - (Optional) To enable process tree computation in audit report


## Attributes Reference

The following attributes are exported:

* `audit_report_status` - auditReportStatus for the accessRequestId
* `process_tree` - Contains the process tree data
* `report` - Contains the report data.
* `time_of_report_generation` - Time when the audit report was generated [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z' 

