# baremetal\_load\_balancer_policies

Provide a list of supported load balancer policies.

## Example Usage

```
data "baremetal_load_balancer_policies" "t" {
  compartment_id = "ocid1.compartment.stub_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attribute Reference
* `policies` - The list of shapes

## Policy Reference
* `name` - The name of the policy.
