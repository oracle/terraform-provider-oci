---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_fault_domains"
sidebar_current: "docs-oci-datasource-identity-fault_domains"
description: |-
  Provides the list of Fault Domains in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_fault_domains
This data source provides the list of Fault Domains in Oracle Cloud Infrastructure Identity service.

Lists the Fault Domains in your tenancy. Specify the OCID of either the tenancy or another
of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm#five).


## Example Usage

```hcl
data "oci_identity_fault_domains" "test_fault_domains" {
	#Required
	availability_domain = var.fault_domain_availability_domain
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the availibilityDomain. 
* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


## Attributes Reference

The following attributes are exported:

* `fault_domains` - The list of fault_domains.

### FaultDomain Reference

The following attributes are exported:

* `availability_domain` - The name of the availabilityDomain where the Fault Domain belongs.
* `compartment_id` - The OCID of the compartment. Currently only tenancy (root) compartment can be provided.
* `id` - The OCID of the Fault Domain.
* `name` - The name of the Fault Domain.

