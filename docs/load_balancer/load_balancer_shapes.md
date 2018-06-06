
# oci_load_balancer_shapes

## LoadBalancerShape DataSource

Gets a list of load_balancer_shapes.

### List Operation
Lists the valid load balancer shapes.
The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer shapes to list.


The following attributes are exported:

* `shapes` - The list of shapes.

### Example Usage

```hcl
data "oci_load_balancer_shapes" "test_load_balancer_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```
### LoadBalancerShape Reference

The following attributes are exported:

* `name` - The name of the shape.  Example: `100Mbps` 
