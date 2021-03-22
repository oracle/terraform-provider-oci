---
subcategory: "Load Balancer"
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

Set the terraform flag `lifecycle { create_before_destroy = true }` in your hostname to facilitate rotating hostnames. 
A hostname cannot be deleted if it is attached to another resource (a listener for example).
Because hostname_names in the listener is an updatable parameter, terraform will attempt to recreate the hostname first and then update the listener but the hostname cannot be deleted while it is attached to a listener so it will fail.
Setting the flag makes it so that when a hostname is recreated, the new hostname will be created first before the old one gets deleted.

## Example Usage

```hcl
resource "oci_load_balancer_hostname" "test_hostname" {
	#Required
	hostname = var.hostname_hostname
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	name = var.hostname_name

    #Optional
    lifecycle {
	    create_before_destroy = true
	}
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Hostname
	* `update` - (Defaults to 20 minutes), when updating the Hostname
	* `delete` - (Defaults to 20 minutes), when destroying the Hostname


## Import

Hostnames can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_hostname.test_hostname "loadBalancers/{loadBalancerId}/hostnames/{name}" 
```

