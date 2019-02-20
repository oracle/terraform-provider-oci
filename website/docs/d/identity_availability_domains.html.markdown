---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_availability_domains"
sidebar_current: "docs-oci-datasource-identity-availability_domains"
description: |-
  Provides the list of Availability Domains in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_availability_domains
This data source provides the list of Availability Domains in Oracle Cloud Infrastructure Identity service.

Lists the availability domains in your tenancy. Specify the OCID of either the tenancy or another
of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm#five).
Note that the order of the results returned can change if availability domains are added or removed; therefore, do not
create a dependency on the list order.


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
* `id` - The OCID of the Availability Domain.
* `name` - The name of the Availability Domain.

