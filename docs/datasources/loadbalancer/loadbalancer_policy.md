# oci\_load\_balancer_policies

[LoadBalancerPolicy Reference][ff23fb78]

  [ff23fb78]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ "LoadBalancerPolicyReference"

Provides a list of supported load balancer policies. A policy determines how traffic is distributed among backend servers.

## Example Usage

```
data "oci_load_balancer_policies" "t" {
  compartment_id = "ocid1.compartment.stub_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attribute Reference
* `policies` - The list of available load balancer policies.

## Policy Reference
* `name` - The name of the policy.
