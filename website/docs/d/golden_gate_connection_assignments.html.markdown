---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_connection_assignments"
sidebar_current: "docs-oci-datasource-golden_gate-connection_assignments"
description: |-
  Provides the list of Connection Assignments in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_connection_assignments
This data source provides the list of Connection Assignments in Oracle Cloud Infrastructure Golden Gate service.

Lists the Connection Assignments in the compartment.

## Example Usage

```hcl
data "oci_golden_gate_connection_assignments" "test_connection_assignments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	connection_id = oci_golden_gate_connection.test_connection.id
	deployment_id = oci_golden_gate_deployment.test_deployment.id
	name = var.connection_assignment_name
	state = var.connection_assignment_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `connection_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection. 
* `deployment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment in which to list resources. 
* `name` - (Optional) The name of the connection in the assignment (aliasName).
* `state` - (Optional) A filter to return only connection assignments having the 'lifecycleState' given.


## Attributes Reference

The following attributes are exported:

* `connection_assignment_collection` - The list of connection_assignment_collection.

### ConnectionAssignment Reference

The following attributes are exported:

* `alias_name` - Credential store alias. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection being referenced. 
* `deployment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connection assignment being referenced. 
* `state` - Possible lifecycle states for connection assignments.
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 

