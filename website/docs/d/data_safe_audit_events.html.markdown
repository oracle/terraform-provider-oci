---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_events"
sidebar_current: "docs-oci-datasource-data_safe-audit_events"
description: |-
  Provides the list of Audit Events in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_events
This data source provides the list of Audit Events in Oracle Cloud Infrastructure Data Safe service.

The ListAuditEvents operation returns specified `compartmentId` audit Events only.
The list does not include any audit Events associated with the `subcompartments` of the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListAuditEvents on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_audit_events" "test_audit_events" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.audit_event_access_level
	compartment_id_in_subtree = var.audit_event_compartment_id_in_subtree
	scim_query = var.audit_event_scim_query
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `scim_query` - (Optional) The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2 of the System for Cross-Domain Identity Management (SCIM) specification, which is available at [RFC3339](https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions, text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format. (Numeric and boolean values should not be quoted.)

	**Example:** (operationTime ge "2021-06-04T12:00:00.000Z") and (eventName eq "LOGON") 


## Attributes Reference

The following attributes are exported:

* `audit_event_collection` - The list of audit_event_collection.

### AuditEvent Reference

The following attributes are exported:

* `items` - Array of audit event summary.
	* `action_taken` - The action taken for this audit event.
	* `application_contexts` - Semicolon-seperated list of application context namespace, attribute, value information in (APPCTX_NSPACE,APPCTX_ATTRIBUTE=<value>) format.
	* `audit_event_time` - The time that the audit event occurs in the target database.
	* `audit_location` - The location of the audit. Currently the value is audit table.
	* `audit_policies` - Comma-seperated list of audit policies that caused the current audit event.
	* `audit_trail_id` - The OCID of the audit trail that generated this audit event. To be noted, this field has been deprecated.
	* `audit_type` - The type of the auditing.
	* `client_hostname` - The name of the host machine from which the session was spawned.
	* `client_id` - The client identifier in each Oracle session.
	* `client_ip` - The IP address of the host machine from which the session was spawned.
	* `client_program` - The application from which the audit event was generated. For example SQL Plus or SQL Developer.
	* `command_param` - List of bind variables associated with the command text.
	* `command_text` - The SQL associated with the audit event.
	* `compartment_id` - The OCID of the compartment containing the audit event. The compartment is the same as that of audit profile of the target database resource.
	* `database_type` - The type of the target database that was audited. Allowed values are
		* DATABASE_CLOUD_SERVICE - Represents Oracle Database Cloud Services.
		* AUTONOMOUS_DATABASE - Represents Oracle Autonomous Databases.
		* INSTALLED_DATABASE - Represents databases running on-premises or on compute instances. 
	* `database_unique_name` - Unique name of the database associated to the peer target database.
	* `db_user_name` - The name of the database user whose actions were audited.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
	* `error_code` - Oracle Error code generated by the action. Zero indicates the action was successful.
	* `error_message` - The detailed message on why the error occurred.
	* `event_name` - The name of the detail action executed by the user on the target database. For example ALTER SEQUENCE, CREATE TRIGGER or CREATE INDEX.
	* `extended_event_attributes` - List of all other attributes of the audit event seperated by a colon other than the one returned in audit record.
	* `fga_policy_name` - Fine-grained auditing (FGA) policy name that generated this audit record.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
	* `id` - The OCID of the audit event.
	* `is_alerted` - Indicates whether an alert was raised for this audit event.
	* `object` - The name of the object affected by the action.
	* `object_owner` - The schema name of the object affected by the action.
	* `object_type` - The type of the object in the source database affected by the action. For example PL/SQL, SYNONYM or PACKAGE BODY.
	* `operation` - The name of the action executed by the user on the target database. For example ALTER, CREATE or DROP.
	* `operation_status` - Indicates whether the operation was a success or a failure.
	* `os_terminal` - The operating system terminal of the user session.
	* `os_user_name` - The name of the operating system user for the database session.
	* `peer_target_database_key` - The secondary id assigned for the peer database registered with Data Safe.
	* `target_class` - The class of the target that was audited.
	* `target_id` - The OCID of the target database that was audited.
	* `target_name` - The name of the target database that was audited.
	* `time_collected` - The timestamp when this audit event was collected from the target database by Data Safe.
	* `trail_source` - The underlying source of unified audit trail.

