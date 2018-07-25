---
layout: "oci"
page_title: "OCI: oci_identity_availability_domains"
sidebar_current: "docs-oci-datasource-identity-availability_domains"
description: |-
  Provides a list of AvailabilityDomains
---

# Data Source: oci_identity_availability_domains
The AvailabilityDomains data source allows access to the list of OCI availability_domains

Lists the Availability Domains in your tenancy. Specify the OCID of either the tenancy or another
of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five).


## Example Usage

```hcl
data "oci_identity_availability_domains" "test_availability_domains" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


## Attributes Reference

The following attributes are exported:

* `availability_domains` - The list of availability_domains.

### AvailabilityDomain Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy.
* `name` - The name of the Availability Domain.

