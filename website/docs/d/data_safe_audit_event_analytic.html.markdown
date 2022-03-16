---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_event_analytic"
sidebar_current: "docs-oci-datasource-data_safe-audit_event_analytic"
description: |-
  Provides details about a specific Audit Event Analytic in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_event_analytic
This data source provides details about a specific Audit Event Analytic resource in Oracle Cloud Infrastructure Data Safe service.

By default ListAuditEventAnalytics operation will return all of the summary columns. To filter desired summary columns, specify
it in the `summaryOf` query parameter.

**Example:** /ListAuditEventAnalytics?summaryField=targetName&summaryField=userName&summaryField=clientHostName&summaryField
             &summaryField=dmls&summaryField=privilege_changes&summaryField=ddls&summaryField=login_failure&summaryField=login_success
             &summaryField=eventcount&q=(operationTime ge '2021-06-13T23:49:14')&groupBy=targetName


## Example Usage

```hcl
data "oci_data_safe_audit_event_analytic" "test_audit_event_analytic" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.audit_event_analytic_access_level
	compartment_id_in_subtree = var.audit_event_analytic_compartment_id_in_subtree
	group_by = var.audit_event_analytic_group_by
	query_time_zone = var.audit_event_analytic_query_time_zone
	scim_query = var.audit_event_analytic_scim_query
	summary_field = var.audit_event_analytic_summary_field
	time_ended = var.audit_event_analytic_time_ended
	time_started = var.audit_event_analytic_time_started
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `group_by` - (Optional) A groupBy can only be used in combination with summaryField parameter. A groupBy value has to be a subset of the values mentioned in summaryField parameter. 
* `query_time_zone` - (Optional) Default time zone is UTC if no time zone provided. The date-time considerations of the resource will be in accordance with the specified time zone. 
* `scim_query` - (Optional) The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2 of the System for Cross-Domain Identity Management (SCIM) specification, which is available at [RFC3339](https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions, text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format. (Numeric and boolean values should not be quoted.)

	**Example:** query=(operationTime ge '2021-06-04T01-00-26') and (eventName eq 'LOGON') 
* `summary_field` - (Optional) Specifies a subset of summarized fields to be returned in the response.
* `time_ended` - (Optional) An optional filter to return audit events whose creation time in the database is less than and equal to the date-time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_started` - (Optional) An optional filter to return audit events whose creation time in the database is greater than and equal to the date-time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 


## Attributes Reference

The following attributes are exported:

* `items` - The aggregated data point items.
	* `count` - Total count of aggregated value.
	* `dimensions` - Details of aggregation dimensions used for summarizing audit events.
		* `audit_event_time` - Time of audit event occurrence in the target database.
		* `audit_type` - Type of auditing.
		* `client_hostname` - Name of the host machine from which the session was spawned.
		* `client_id` - The client identifier in each Oracle session.
		* `client_program` - The application from which the audit event was generated. Examples SQL Plus or SQL Developer.
		* `db_user_name` - Name of the database user whose actions were audited.
		* `event_name` - Name of the detail action executed by the user on the target database. i.e ALTER SEQUENCE, CREATE TRIGGER, CREATE INDEX.
		* `object_type` - Type of object in the source database affected by the action. i.e PL/SQL, SYNONYM, PACKAGE BODY.
		* `target_class` - Class of the target that was audited.
		* `target_id` - The OCID of the target database that was audited.
		* `target_name` - The name of the target database that was audited.
	* `display_name` - Display Name of aggregation field.
	* `metric_name` - Name of the aggregation.
	* `time_ended` - The time at which the aggregation ended.
	* `time_started` - The time at which the aggregation started.

