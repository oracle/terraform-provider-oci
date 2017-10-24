# oci\_load\_balancer_shapes

[LoadBalancerShape Reference][843ad06b]

  [843ad06b]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerShape/ "LoadBalancerShapeReference"

Provides a list of supported load balancer shapes. A shape is a template that determines the total pre-provisioned bandwidth (ingress plus egress) for the load balancer.

## Example Usage

```
data "oci_load_balancer_shapes" "t" {
  compartment_id = "ocid1.compartment.stub_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attribute Reference
* `shapes` - The list of valid load balancer shapes.

## Shape Reference
* `name` - The name of the shape.
