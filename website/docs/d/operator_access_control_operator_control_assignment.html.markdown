---
subcategory: "Operator Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_operator_access_control_operator_control_assignment"
sidebar_current: "docs-oci-datasource-operator_access_control-operator_control_assignment"
description: |-
  Provides details about a specific Operator Control Assignment in Oracle Cloud Infrastructure Operator Access Control service
---

# Data Source: oci_operator_access_control_operator_control_assignment
This data source provides details about a specific Operator Control Assignment resource in Oracle Cloud Infrastructure Operator Access Control service.

Gets the details of an Operator Control Assignment of the specified ID.

## Example Usage

```hcl
data "oci_operator_access_control_operator_control_assignment" "test_operator_control_assignment" {
	#Required
	operator_control_assignment_id = oci_operator_access_control_operator_control_assignment.test_operator_control_assignment.id
}
```

## Argument Reference

The following arguments are supported:

* `operator_control_assignment_id` - (Required) unique OperatorControl identifier


## Attributes Reference

The following attributes are exported:

* `assigner_id` - The OCID of the user who created this operator control assignment.
* `comment` - Comment about the assignment of the operator control to this target resource.
* `compartment_id` - The OCID of the comparment that contains the operator control assignment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. 
* `detachment_description` - description containing reason for releasing of OperatorControl.
* `error_code` - The code identifying the error occurred during Assignment operation.
* `error_message` - The message describing the error occurred during Assignment operation.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. 
* `id` - The OCID of the operator control assignment.
* `is_auto_approve_during_maintenance` - The boolean if true would autoApprove during maintenance.
* `is_default_assignment` - Whether the assignment is a default assignment.    
* `is_enforced_always` - If set, then the target resource is always governed by the operator control.
* `is_hypervisor_log_forwarded` - If set, then the hypervisor audit logs will be forwarded to the relevant remote syslog server
* `is_log_forwarded` - If set indicates that the audit logs are being forwarded to the relevant remote logging server
* `lifecycle_details` - More in detail about the lifeCycleState.
* `op_control_name` - Name of the operator control name associated.
* `operator_control_id` - The OCID of the operator control.
* `remote_syslog_server_address` - The address of the remote syslog server where the audit logs are being forwarded to. Address in host or IP format.
* `remote_syslog_server_ca_cert` - The CA certificate of the remote syslog server.
* `remote_syslog_server_port` - The listening port of the remote syslog server. The port range is 0 - 65535. Only TCP supported.
* `resource_compartment_id` - The OCID of the compartment that contains the target resource.
* `resource_id` - The OCID of the target resource.
* `resource_name` - Name of the target resource.
* `resource_type` - resourceType for which the OperatorControlAssignment is applicable
* `state` - The current lifcycle state of the OperatorControl.
* `time_assignment_from` - The time at which the target resource will be brought under the governance of the operator control expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: '2020-05-22T21:10:29.600Z' 
* `time_assignment_to` - The time at which the target resource will leave the governance of the operator control expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `time_of_assignment` - Time when the operator control assignment is created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `time_of_deletion` - Time on which the operator control assignment was deleted in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format.Example: '2020-05-22T21:10:29.600Z' 
* `unassigner_id` - User id who released the operatorControl.

