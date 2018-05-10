
# oci_identity_tenancies

## Tenancy DataSource

Gets a single tenancy

### Get Operation
Get the specified tenancy's information.
The following arguments are supported:

* `tenancy_id` - (Required) The OCID of the tenancy.


The following attributes are exported:

* `description` - The description of the tenancy.
* `home_region_key` - The region key for the tenancy's home region. For more information about regions, see [Regions and Availability Domains](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm).  Allowed values are: - `IAD` - `PHX` - `FRA` - `LHR` 
* `tenancy_id` - The OCID of the tenancy.
* `name` - The name of the tenancy.


### Example Usage

```
data "oci_identity_tenancies" "test_tenancies" {
	#Required
	tenancy_id = "${var.tenancy_tenancy_id}"
}
```
### Tenancy Reference

The following attributes are exported:

* `description` - The description of the tenancy.
* `home_region_key` - The region key for the tenancy's home region. For more information about regions, see [Regions and Availability Domains](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm).  Allowed values are: - `IAD` - `PHX` - `FRA` - `LHR` 
* `tenancy_id` - The OCID of the tenancy.
* `name` - The name of the tenancy.
