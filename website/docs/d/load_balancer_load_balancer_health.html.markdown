---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_health"
sidebar_current: "docs-oci-datasource-load_balancer-load_balancer_health"
description: |-
  Provides details about a specific Load Balancer Health in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_health
This data source provides details about a specific Load Balancer Health resource in Oracle Cloud Infrastructure Load Balancer service.

Gets the health status for the specified load balancer.

## Example Usage

```hcl
data "oci_load_balancer_health" "test_load_balancer_health" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer to return health status for.


## Attributes Reference

The following attributes are exported:

* `critical_state_backend_set_names` - A list of backend sets that are currently in the `CRITICAL` health state. The list identifies each backend set by the friendly name you assigned when you created it.  Example: `example_backend_set` 
* `status` - The overall health status of the load balancer.
	*  **OK:** All backend sets associated with the load balancer return a status of `OK`.
	*  **WARNING:** At least one of the backend sets associated with the load balancer returns a status of `WARNING`, no backend sets return a status of `CRITICAL`, and the load balancer life cycle state is `ACTIVE`.
	*  **CRITICAL:** One or more of the backend sets associated with the load balancer return a status of `CRITICAL`.
	*  **UNKNOWN:** If any one of the following conditions is true:
	*  The load balancer life cycle state is not `ACTIVE`.
	*  No backend sets are defined for the load balancer.
	*  More than half of the backend sets associated with the load balancer return a status of `UNKNOWN`, none of the backend sets return a status of `WARNING` or `CRITICAL`, and the load balancer life cycle state is `ACTIVE`.
	*  The system could not retrieve metrics for any reason. 
* `total_backend_set_count` - The total number of backend sets associated with this load balancer.  Example: `4` 
* `unknown_state_backend_set_names` - A list of backend sets that are currently in the `UNKNOWN` health state. The list identifies each backend set by the friendly name you assigned when you created it.  Example: `example_backend_set2` 
* `warning_state_backend_set_names` - A list of backend sets that are currently in the `WARNING` health state. The list identifies each backend set by the friendly name you assigned when you created it.  Example: `example_backend_set3` 

