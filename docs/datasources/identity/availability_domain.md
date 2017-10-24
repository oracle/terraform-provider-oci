# oci\_identity\_availability\_domain

[AvailabilityDomain Reference][65774c17]

  [65774c17]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/AvailabilityDomain/ "AvailabilityDomainReference"

Lists Availability Domains (ADs).

## Example Usage

```
data "oci_identity_availability_domains" "t" {
  compartment_id = "compartmentID"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy.

## AvailabilityDomain Reference
* `compartment_id` - The OCID of the tenancy.
* `name` - The name of the Availability Domain.
