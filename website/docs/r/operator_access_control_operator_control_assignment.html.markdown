---
subcategory: "Operator Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_operator_access_control_operator_control_assignment"
sidebar_current: "docs-oci-resource-operator_access_control-operator_control_assignment"
description: |-
  Provides the Operator Control Assignment resource in Oracle Cloud Infrastructure Operator Access Control service
---

# oci_operator_access_control_operator_control_assignment
This resource provides the Operator Control Assignment resource in Oracle Cloud Infrastructure Operator Access Control service.

Creates an Operator Control Assignment resource. In effect, this brings the target resource under the governance of the Operator Control for specified time duration.

## Example Usage

```hcl
resource "oci_operator_access_control_operator_control_assignment" "test_operator_control_assignment" {
	#Required
	compartment_id = var.compartment_id
	is_enforced_always = var.operator_control_assignment_is_enforced_always
	operator_control_id = oci_operator_access_control_operator_control.test_operator_control.id
	resource_compartment_id = oci_identity_compartment.test_compartment.id
	resource_id = oci_operator_access_control_resource.test_resource.id
	resource_name = var.operator_control_assignment_resource_name
	resource_type = var.operator_control_assignment_resource_type

	#Optional
	comment = var.operator_control_assignment_comment
	defined_tags = var.operator_control_assignment_defined_tags
	freeform_tags = var.operator_control_assignment_freeform_tags
	is_auto_approve_during_maintenance = var.operator_control_assignment_is_auto_approve_during_maintenance
	is_hypervisor_log_forwarded = var.operator_control_assignment_is_hypervisor_log_forwarded
	is_log_forwarded = var.operator_control_assignment_is_log_forwarded
	remote_syslog_server_address = var.operator_control_assignment_remote_syslog_server_address
	remote_syslog_server_ca_cert = var.operator_control_assignment_remote_syslog_server_ca_cert
	remote_syslog_server_port = var.operator_control_assignment_remote_syslog_server_port
	time_assignment_from = var.operator_control_assignment_time_assignment_from
	time_assignment_to = var.operator_control_assignment_time_assignment_to
}
```

## Argument Reference

The following arguments are supported:

* `comment` - (Optional) (Updatable) Comment about the assignment of the operator control to this target resource.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the operator control assignment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. 
* `is_auto_approve_during_maintenance` - (Optional) (Updatable) The boolean if true would autoApprove during maintenance.
* `is_enforced_always` - (Required) (Updatable) If set, then the target resource is always governed by the operator control.
* `is_hypervisor_log_forwarded` - (Optional) (Updatable) If set, then the hypervisor audit logs will be forwarded to the relevant remote syslog server
* `is_log_forwarded` - (Optional) (Updatable) If set, then the audit logs will be forwarded to the relevant remote logging server
* `operator_control_id` - (Required) The OCID of the operator control that is being assigned to a target resource.
* `remote_syslog_server_address` - (Optional) (Updatable) The address of the remote syslog server where the audit logs will be forwarded to. Address in host or IP format.
* `remote_syslog_server_ca_cert` - (Optional) (Updatable) The CA certificate of the remote syslog server. Identity of the remote syslog server will be asserted based on this certificate.
* `remote_syslog_server_port` - (Optional) (Updatable) The listening port of the remote syslog server. The port range is 0 - 65535. Only TCP supported.
* `resource_compartment_id` - (Required) The OCID of the compartment that contains the target resource.
* `resource_id` - (Required) The OCID of the target resource being brought under the governance of the operator control.
* `resource_name` - (Required) Name of the target resource.
* `resource_type` - (Required) Type of the target resource.
* `time_assignment_from` - (Optional) (Updatable) The time at which the target resource will be brought under the governance of the operator control in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `time_assignment_to` - (Optional) (Updatable) The time at which the target resource will leave the governance of the operator control in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format.Example: '2020-05-22T21:10:29.600Z' 
* `validate_assignment_trigger` - (Optional) (Updatable) An optional property when incremented triggers Validate Assignment. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Operator Control Assignment
	* `update` - (Defaults to 20 minutes), when updating the Operator Control Assignment
	* `delete` - (Defaults to 20 minutes), when destroying the Operator Control Assignment


## Import

OperatorControlAssignments can be imported using the `id`, e.g.

```
$ terraform import oci_operator_access_control_operator_control_assignment.test_operator_control_assignment "id"
```

