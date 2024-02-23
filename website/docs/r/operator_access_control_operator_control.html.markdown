---
subcategory: "Operator Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_operator_access_control_operator_control"
sidebar_current: "docs-oci-resource-operator_access_control-operator_control"
description: |-
  Provides the Operator Control resource in Oracle Cloud Infrastructure Operator Access Control service
---

# oci_operator_access_control_operator_control
This resource provides the Operator Control resource in Oracle Cloud Infrastructure Operator Access Control service.

Creates an Operator Control.


## Example Usage

```hcl
resource "oci_operator_access_control_operator_control" "test_operator_control" {
	#Required
	approver_groups_list = var.operator_control_approver_groups_list
	compartment_id = var.compartment_id
	is_fully_pre_approved = var.operator_control_is_fully_pre_approved
	operator_control_name = oci_operator_access_control_operator_control.test_operator_control.name
	resource_type = var.operator_control_resource_type

	#Optional
	approvers_list = var.operator_control_approvers_list
	defined_tags = var.operator_control_defined_tags
	description = var.operator_control_description
	email_id_list = var.operator_control_email_id_list
	freeform_tags = var.operator_control_freeform_tags
	number_of_approvers = var.operator_control_number_of_approvers
	pre_approved_op_action_list = var.operator_control_pre_approved_op_action_list
	system_message = var.operator_control_system_message
}
```

## Argument Reference

The following arguments are supported:

* `approver_groups_list` - (Required) (Updatable) List of user groups who can approve an access request associated with a resource governed by this operator control.
* `approvers_list` - (Optional) (Updatable) List of users who can approve an access request associated with a resource governed by this operator control.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains this operator control.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. 
* `description` - (Optional) (Updatable) Description of the operator control.
* `email_id_list` - (Optional) (Updatable) List of emailId. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. 
* `is_fully_pre_approved` - (Required) (Updatable) Whether all the operator actions have been pre-approved. If yes, all access requests associated with a resource governed by this operator control  will be auto-approved.         
* `number_of_approvers` - (Optional) (Updatable) Number of approvers required to approve an access request.
* `operator_control_name` - (Required) (Updatable) Name of the operator control.
* `pre_approved_op_action_list` - (Optional) (Updatable) List of pre-approved operator actions. Access requests associated with a resource governed by this operator control will be auto-approved if the access request only contain operator actions in the pre-approved list. 
* `resource_type` - (Required) resourceType for which the OperatorControl is applicable
* `system_message` - (Optional) (Updatable) This is the message that will be displayed to the operator users while accessing the system.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `approval_required_op_action_list` - List of operator actions that need explicit approval. Any operator action not in the pre-approved list will require explicit approval. Access requests associated with a resource governed by this operator control will be require explicit approval if the access request contains any operator action in this list.  
* `approver_groups_list` - List of user groups who can approve an access request associated with a target resource under the governance of this operator control.
* `approvers_list` - List of users who can approve an access request associated with a target resource under the governance of this operator control.
* `compartment_id` - The OCID of the compartment that contains the operator control.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. 
* `description` - Description of operator control.
* `email_id_list` - List of emailId. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. 
* `id` - The OCID of the operator control.
* `is_default_operator_control` - Whether the operator control is a default Operator Control. 
* `is_fully_pre_approved` - Whether all the operator actions have been pre-approved. If yes, all access requests associated with a resource governed by this operator control  will be auto-approved. 
* `last_modified_info` - Description associated with the latest modification of the operator control.
* `number_of_approvers` - Number of approvers required to approve an access request.
* `operator_control_name` - Name of the operator control. The name must be unique.
* `pre_approved_op_action_list` - List of pre-approved operator actions. Access requests associated with a resource governed by this operator control will be automatically approved if the access request only contain operator actions in the pre-approved list.        
* `resource_type` - resourceType for which the OperatorControl is applicable
* `state` - The current lifecycle state of the operator control.
* `system_message` - System message that would be displayed to the operator users on accessing the target resource under the governance of this operator control.
* `time_of_creation` - Time when the operator control was created expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `time_of_deletion` - Time when deleted expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z'. Note a deleted operator control still stays in the system, so that you can still audit operator actions associated with access requests raised on target resources governed by the deleted operator control. 
* `time_of_modification` - Time when the operator control was last modified expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: '2020-05-22T21:10:29.600Z' 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Operator Control
	* `update` - (Defaults to 20 minutes), when updating the Operator Control
	* `delete` - (Defaults to 20 minutes), when destroying the Operator Control


## Import

OperatorControls can be imported using the `id`, e.g.

```
$ terraform import oci_operator_access_control_operator_control.test_operator_control "id"
```

