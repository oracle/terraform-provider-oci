# baremetal\_identity\_availability\_domain

Lists availability domains.

## Example Usage

```
data "baremetal_identity_availability_domains" "t" {
  compartment_id = "compartmentID"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy.

## AvailabilityDomain Reference
* `compartment_id` - The OCID of the tenancy.
* `name` - The name of the Availability Domain.
