# baremetal\_load\_balancer_shapes

Provide a list of supported load balancer shapes.

## Example Usage

```
data "baremetal_load_balancer_shapes" "t" {
  compartment_id = "ocid1.compartment.stub_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attribute Reference
* `shapes` - The list of shapes

## Shape Reference
* `name` - The name of the shape.
