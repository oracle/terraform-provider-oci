
# oci_core_services

## Service DataSource

Gets a list of services.

### List Operation
Lists the available services that you can access through a service gateway in this region.

The following arguments are supported:



The following attributes are exported:

* `services` - The list of services.

### Example Usage

```
data "oci_core_services" "test_services" {
}
```
### Service Reference

The following attributes are exported:

* `cidr_block` - This value will be used as Destination CidrBlock while creating a route rule with service gateway as target.
* `description` - Description of this particular Service, provided by the Service owner. 
* `id` - The Service's Oracle ID ([OCID])(/Content/General/Concepts/identifiers.htm).
* `name` - Name of the Service.
