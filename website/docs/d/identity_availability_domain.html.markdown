---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_availability_domain"
sidebar_current: "docs-oci-datasource-identity-availability-domain"
description: |-
  Provides details about a specific Availability Domain in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_availability_domain
This data source provides the details of a single Availability Domain in Oracle Cloud Infrastructure Identity service.


## Example Usage

```hcl
data "oci_identity_availability_domain" "test_compartment" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	
	#Optional (one or the other is required)
	id = "${var.id}"
	ad_number = "${var.ad_number}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy.
* `id` - (Optional) The OCID of the Availability Domain. Required if `ad_number` is not specified.
* `ad_number` - (Optional) The number of the Availability Domain. Required if `id` is not specified. This number corresponds to the integer in the Availability Domain `name`.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy. 
* `id` - The OCID of the Availability Domain.
* `ad_number` - The number of the Availability Domain. For example, the `ad_number` for YXol:US-ASHBURN-AD-1 would be "1"
* `name` - The name of the Availability Domain.  

