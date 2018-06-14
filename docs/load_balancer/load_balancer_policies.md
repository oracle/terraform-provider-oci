
# oci_load_balancer_policies

## LoadBalancerPolicy DataSource

Gets a list of load_balancer_policies.

### List Operation
Lists the available load balancer policies.
The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer policies to list.


The following attributes are exported:

* `policies` - The list of policies.

### Example Usage

```hcl
data "oci_load_balancer_policies" "test_load_balancer_policies" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```
### LoadBalancerPolicy Reference

The following attributes are exported:

* `name` - The name of a load balancing policy.  Example: 'LEAST_CONNECTIONS' 
