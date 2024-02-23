---
subcategory: "Operator Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_operator_access_control_operator_controls"
sidebar_current: "docs-oci-datasource-operator_access_control-operator_controls"
description: |-
  Provides the list of Operator Controls in Oracle Cloud Infrastructure Operator Access Control service
---

# Data Source: oci_operator_access_control_operator_controls
This data source provides the list of Operator Controls in Oracle Cloud Infrastructure Operator Access Control service.

Lists the operator controls in the compartment.


## Example Usage

```hcl
data "oci_operator_access_control_operator_controls" "test_operator_controls" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.operator_control_display_name
	resource_type = var.operator_control_resource_type
	state = var.operator_control_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return OperatorControl that match the entire display name given.
* `resource_type` - (Optional) A filter to return only lists of resources that match the entire given service type.
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the given OperatorControl lifecycleState.


## Attributes Reference

The following attributes are exported:

* `operator_control_collection` - The list of operator_control_collection.

### OperatorControl Reference

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

