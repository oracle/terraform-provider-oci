# baremetal\_load\_balancer

Provide a load balancer resource.

## Example Usage

```
resource "baremetal_load_balancer" "t" {
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
* `display_name` - (optional) A user-friendly name. Does not have to be unique, and it's changeable.

## Attributes Reference
* `id` - The OCID of the load balancer.
* `ip_addresses` - An array of IP Addresses
* `time_created` - The date and time the image was created.
