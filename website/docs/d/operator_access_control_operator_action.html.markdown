---
subcategory: "Operator Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_operator_access_control_operator_action"
sidebar_current: "docs-oci-datasource-operator_access_control-operator_action"
description: |-
  Provides details about a specific Operator Action in Oracle Cloud Infrastructure Operator Access Control service
---

# Data Source: oci_operator_access_control_operator_action
This data source provides details about a specific Operator Action resource in Oracle Cloud Infrastructure Operator Access Control service.

Gets the operator action associated with the specified operator action ID.

## Example Usage

```hcl
data "oci_operator_access_control_operator_action" "test_operator_action" {
	#Required
	operator_action_id = oci_operator_access_control_operator_action.test_operator_action.id
}
```

## Argument Reference

The following arguments are supported:

* `operator_action_id` - (Required) Unique Oracle supplied identifier associated with the operator action.


## Attributes Reference

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

