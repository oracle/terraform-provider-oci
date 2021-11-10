---
subcategory: "Operator Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_operator_access_control_operator_actions"
sidebar_current: "docs-oci-datasource-operator_access_control-operator_actions"
description: |-
  Provides the list of Operator Actions in Oracle Cloud Infrastructure Operator Access Control service
---

# Data Source: oci_operator_access_control_operator_actions
This data source provides the list of Operator Actions in Oracle Cloud Infrastructure Operator Access Control service.

Lists all the OperatorActions available in the system.


## Example Usage

```hcl
data "oci_operator_access_control_operator_actions" "test_operator_actions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.operator_action_name
	resource_type = var.operator_action_resource_type
	state = var.operator_action_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `name` - (Optional) A filter to return only resources that match the entire display name given.
* `resource_type` - (Optional) A filter to return only lists of resources that match the entire given service type.
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the given OperatorAction lifecycleState.


## Attributes Reference

The following attributes are exported:

* `operator_action_collection` - The list of operator_action_collection.

### OperatorAction Reference

The following attributes are exported:

* `component` - Name of the infrastructure layer associated with the operator action.
* `customer_display_name` - Display Name of the operator action.
* `description` - Description of the operator action in terms of associated risk profile, and characteristics of the operating system commands made available to the operator under this operator action. 
* `id` - Unique Oracle assigned identifier for the operator action.
* `name` - Unique name of the operator action.
* `properties` - Fine grained properties associated with the operator control.
	* `name` - Name of the property
	* `value` - value of the property
* `resource_type` - resourceType for which the OperatorAction is applicable

