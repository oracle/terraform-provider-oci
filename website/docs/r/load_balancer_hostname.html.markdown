---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_hostname"
sidebar_current: "docs-oci-resource-load_balancer-hostname"
description: |-
  Provides the Hostname resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_hostname
This resource provides the Hostname resource in Oracle Cloud Infrastructure Load Balancer service.

Adds a hostname resource to the specified load balancer. For more information, see
[Managing Request Routing](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrequest.htm).


## Example Usage

```hcl
resource "oci_load_balancer_hostname" "test_hostname" {
	#Required
	hostname = "${var.hostname_hostname}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.hostname_name}"
}
```

## Argument Reference

The following arguments are supported:

* `hostname` - (Required) (Updatable) A virtual hostname. For more information about virtual hostname string construction, see [Managing Request Routing](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrequest.htm#routing).  Example: `app.example.com` 
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer to add the hostname to.
* `name` - (Required) A friendly name for the hostname resource. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_hostname_001` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `hostname` - A virtual hostname. For more information about virtual hostname string construction, see [Managing Request Routing](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/managingrequest.htm#routing).  Example: `app.example.com` 
* `name` - A friendly name for the hostname resource. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_hostname_001` 

## Import

Hostnames can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_hostname.test_hostname "loadBalancers/{loadBalancerId}/hostnames/{name}" 
```

