
# oci_identity_availability_domains

## AvailabilityDomain DataSource

Gets a list of availability_domains.

### List Operation
Lists the Availability Domains in your tenancy. Specify the OCID of either the tenancy or another
of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


The following attributes are exported:

* `availability_domains` - The list of availability_domains.

### Example Usage

```hcl
data "oci_identity_availability_domains" "test_availability_domains" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```
### AvailabilityDomain Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy.
* `name` - The name of the Availability Domain.
