---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tenancy"
sidebar_current: "docs-oci-datasource-identity-tenancy"
description: |-
  Provides details about a specific Tenancy in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_tenancy
This data source provides details about a specific Tenancy resource in Oracle Cloud Infrastructure Identity service.

Get the specified tenancy's information.

## Example Usage

```hcl
data "oci_identity_tenancy" "test_tenancy" {
	#Required
	tenancy_id = "${var.tenancy_ocid}"
}
```

## Argument Reference

The following arguments are supported:

* `tenancy_id` - (Required) The OCID of the tenancy.


## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the tenancy.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `home_region_key` - The region key for the tenancy's home region. For more information about regions, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).

	Allowed values are:
	* `IAD`
	* `PHX`
	* `FRA`
	* `LHR` 
* `id` - The OCID of the tenancy.
* `name` - The name of the tenancy.

