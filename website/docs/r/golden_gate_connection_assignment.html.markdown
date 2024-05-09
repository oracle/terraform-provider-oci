---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_connection_assignment"
sidebar_current: "docs-oci-resource-golden_gate-connection_assignment"
description: |-
  Provides the Connection Assignment resource in Oracle Cloud Infrastructure Golden Gate service
---

# oci_golden_gate_connection_assignment
This resource provides the Connection Assignment resource in Oracle Cloud Infrastructure Golden Gate service.

Creates a new Connection Assignment.

## Example Usage

```hcl
resource "oci_golden_gate_connection_assignment" "test_connection_assignment" {
	#Required
	connection_id = oci_golden_gate_connection.test_connection.id
	deployment_id = oci_golden_gate_deployment.test_deployment.id

	#Optional
	is_lock_override = var.connection_assignment_is_lock_override
}
```

## Argument Reference

The following arguments are supported:

* `connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being referenced. 
* `deployment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `is_lock_override` - (Optional) Whether to override locks (if any exist).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `alias_name` - Credential store alias. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being referenced. 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection assignment being referenced. 
* `state` - Possible lifecycle states for connection assignments.
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Connection Assignment
	* `update` - (Defaults to 20 minutes), when updating the Connection Assignment
	* `delete` - (Defaults to 20 minutes), when destroying the Connection Assignment


## Import

ConnectionAssignments can be imported using the `id`, e.g.

```
$ terraform import oci_golden_gate_connection_assignment.test_connection_assignment "id"
```

