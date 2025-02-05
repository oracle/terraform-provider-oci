---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduled_action_params"
sidebar_current: "docs-oci-datasource-database-scheduled_action_params"
description: |-
  Provides the list of Scheduled Action Params in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_scheduled_action_params
This data source provides the list of Scheduled Action Params in Oracle Cloud Infrastructure Database service.

List all the action params and their possible values for a given action type


## Example Usage

```hcl
data "oci_database_scheduled_action_params" "test_scheduled_action_params" {
	#Required
	type = var.scheduled_action_param_type
}
```

## Argument Reference

The following arguments are supported:

* `type` - (Required) The type of the scheduled action


## Attributes Reference

The following attributes are exported:

* `action_param_values_collection` - The list of action_param_values_collection.

### ScheduledActionParam Reference

The following attributes are exported:

* `items` - List of Action Parameters and their possible values.
	* `default_value` - The default value for this parameter.
	* `is_required` - Whether this parameter is required or not for this action type.、
	* `parameter_name` - The name of this parameter.
	* `parameter_type` - The type of the parameter.
	* `parameter_values` - Possible values for this parameter. In case of integer it's min and max values.

