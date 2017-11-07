# oci\_load\_balancer_protocols

[LoadBalancerProtocol Reference][b632188a]

  [b632188a]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerProtocol/ "LoadBalancerProtocolReference"

Provide a list of supported load balancer protocols. The protocol that defines the type of traffic accepted by a listener.

## Example Usage

```
data "oci_load_balancer_pprotocols" "t" {
  compartment_id = "ocid1.compartment.stub_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attribute Reference
* `protocols` - The list of supported traffic protocols.

## Protocol Reference
* `name` - The name of the protocol.
