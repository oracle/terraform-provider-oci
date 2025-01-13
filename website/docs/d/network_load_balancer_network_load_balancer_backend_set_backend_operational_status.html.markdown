---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_network_load_balancer_backend_set_backend_operational_status"
sidebar_current: "docs-oci-datasource-network_load_balancer-network_load_balancer_backend_set_backend_operational_status"
description: |-
  Provides details about a specific Network Load Balancer Backend Set Backend Operational Status in Oracle Cloud Infrastructure Network Load Balancer service
---

# Data Source: oci_network_load_balancer_network_load_balancer_backend_set_backend_operational_status
This data source provides details about a specific Network Load Balancer Backend Set Backend Operational Status resource in Oracle Cloud Infrastructure Network Load Balancer service.

Retrieves the current operational status of the specified backend server.

## Example Usage

```hcl
data "oci_network_load_balancer_network_load_balancer_backend_set_backend_operational_status" "test_network_load_balancer_backend_set_backend_operational_status" {
	#Required
	backend_name = oci_network_load_balancer_backend.test_backend.name
	backend_set_name = oci_network_load_balancer_backend_set.test_backend_set.name
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `backend_name` - (Required) The name of the backend server to retrieve health status for. If the backend was created with an explicitly specified name, that name should be used here. If the backend was created without explicitly specifying the name, but was created using ipAddress, this is specified as <ipAddress>:<port>. If the backend was created without explicitly specifying the name, but was created using targetId, this is specified as <targetId>:<port>.  Example: `10.0.0.3:8080` or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:8080` 
* `backend_set_name` - (Required) The name of the backend set associated with the backend server for which to retrieve the operational status.  Example: `example_backend_set` 
* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.


## Attributes Reference

The following attributes are exported:

* `status` - The operational status. 

