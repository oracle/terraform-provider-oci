# baremetal\_load\_balancer_protocols 

Provide a list of supported load balancer protocols.

## Example Usage

```
data "baremetal_load_balancer_pprotocols" "t" {
  compartment_id = "ocid1.compartment.stub_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attribute Reference
* `protocols` - The list of shapes

## Protocol Reference
* `name` - The name of the protocol.
