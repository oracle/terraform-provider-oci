---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_connection_assignment"
sidebar_current: "docs-oci-datasource-golden_gate-connection_assignment"
description: |-
  Provides details about a specific Connection Assignment in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_connection_assignment
This data source provides details about a specific Connection Assignment resource in Oracle Cloud Infrastructure Golden Gate service.

Retrieves a Connection Assignment.


## Example Usage

```hcl
data "oci_golden_gate_connection_assignment" "test_connection_assignment" {
	#Required
	connection_assignment_id = oci_golden_gate_connection_assignment.test_connection_assignment.id
}
```

## Argument Reference

The following arguments are supported:

* `connection_assignment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Connection Assignment. 


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

