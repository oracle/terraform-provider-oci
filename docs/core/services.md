
# oci_core_services

## Service DataSource

Gets a list of services.

### List Operation
Lists the available services that you can access through a service gateway in this region.

The following arguments are supported:



The following attributes are exported:

* `services` - The list of services.

### Example Usage

```hcl
data "oci_core_services" "test_services" {
}
```
### Service Reference

The following attributes are exported:

* `cidr_block` - A string that represents the public endpoints for the service. When you set up a route rule to route traffic to the service gateway, use this value as the destination CIDR block for the rule. See [Route Table](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/RouteTable/). 
* `description` - Description of the service. 
* `id` - The service's [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `name` - Name of the service.
