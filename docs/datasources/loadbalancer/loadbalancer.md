# baremetal\_load\_balancers

Provide a list of load balancer resources.

## Example Usage

```
data "baremetal_load_balancers" "t" {
  compartment_id = "ocid1.compartment.stub_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attributes Reference
* `load_balancers` - The list of load balancers

## Load Balancer Reference
* `id` - The OCID of the load balancer.
* `ip_addresses` - An array of IP Addresses
* `is_private` - Whether the load balancer has a VCN-local (private) IP address
* `time_created` - The date and time the image was created.
* `shape` - A template that determines the total pre-provisioned bandwidth (ingress plus egress).
* `subnet_ids` - An array of subnet OCIDs
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.