# oci\_load\_balancers

[LoadBalancer Reference][4c71d901]

  [4c71d901]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancer/ "LoadBalancerReference"

Provides a list of the properties that define a load balancer.

## Example Usage

```
data "oci_load_balancers" "t" {
  compartment_id = "ocid1.compartment.stub_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attributes Reference
* `load_balancers` - The list of load balancers.

## Load Balancer Reference
* `id` - The OCID of the load balancer.
* `ip_addresses` - An array of IP Addresses
* `is_private` - Whether the load balancer has a VCN-local (private) IP address. Example: `true`
* `time_created` - The date and time the load balancer was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `shape` - A template that determines the total pre-provisioned bandwidth (ingress plus egress).
* `subnet_ids` - An array of subnet OCIDs
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
