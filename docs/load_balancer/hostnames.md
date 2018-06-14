# oci_load_balancer_hostname

## Hostname Resource

### Hostname Reference

The following attributes are exported:

* `hostname` - A virtual hostname. For more information about virtual hostname string construction, see [Managing Request Routing](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm#routing).  Example: `app.example.com` 
* `name` - A friendly name for the hostname resource. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_hostname_001` 



### Create Operation
Adds a hostname resource to the specified load balancer. For more information, see
[Managing Request Routing](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm).


The following arguments are supported:

* `hostname` - (Required) A virtual hostname. For more information about virtual hostname string construction, see [Managing Request Routing](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm#routing).  Example: `app.example.com` 
* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer to add the hostname to.
* `name` - (Required) A friendly name for the hostname resource. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_hostname_001` 


### Update Operation
Overwrites an existing hostname resource on the specified load balancer. Use this operation to change a
virtual hostname.


The following arguments support updates:
* `hostname` - A virtual hostname. For more information about virtual hostname string construction, see [Managing Request Routing](https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Tasks/managingrequest.htm#routing).  Example: `app.example.com` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_load_balancer_hostname" "test_hostname" {
	#Required
	hostname = "${var.hostname_hostname}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.hostname_name}"
}
```

# oci_load_balancer_hostnames

## Hostname DataSource

Gets a list of hostnames.

### List Operation
Lists all hostname resources associated with the specified load balancer.
The following arguments are supported:

* `load_balancer_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the load balancer associated with the hostnames to retrieve. 


The following attributes are exported:

* `hostnames` - The list of hostnames.

### Example Usage

```hcl
data "oci_load_balancer_hostnames" "test_hostnames" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```