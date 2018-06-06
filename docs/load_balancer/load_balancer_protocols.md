
# oci_load_balancer_protocols

## LoadBalancerProtocol DataSource

Gets a list of load_balancer_protocols.

### List Operation
Lists all supported traffic protocols.
The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer protocols to list.


The following attributes are exported:

* `protocols` - The list of protocols.

### Example Usage

```hcl
data "oci_load_balancer_protocols" "test_load_balancer_protocols" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```
### LoadBalancerProtocol Reference

The following attributes are exported:

* `name` - The name of a protocol.  Example: 'HTTP' 
