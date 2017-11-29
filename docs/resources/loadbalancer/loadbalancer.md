# oci\_load\_balancer

[LoadBalancer Reference][9603545e]

  [9603545e]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancer/ "LoadBalancerReference"


## Example Usage

```
resource "oci_load_balancer" "t" {
  shape          = "stub_shape_id"
  compartment_id = "ocid1.compartment.stub_id"
  subnet_ids     = ["ocid1.subnet.stub_id"]
  display_name   = "stub_display_name"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `shape` - (Required) A template that determines the total pre-provisioned bandwidth (ingress plus egress).
* `subnet_ids` - (Required) An array of subnet OCIDs
* `display_name` - (optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `is_private` - (optional) Whether the load balancer has a VCN-local (private) IP address. Example: `true`

## Attributes Reference
* `id` - The OCID of the load balancer.
* `ip_addresses` - An array of IP Addresses.
* `time_created` - The date and time the load balancer was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
