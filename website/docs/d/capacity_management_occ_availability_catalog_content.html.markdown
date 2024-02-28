---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_availability_catalog_content"
sidebar_current: "docs-oci-datasource-capacity_management-occ_availability_catalog_content"
description: |-
  Provides details about a specific Occ Availability Catalog Content in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_occ_availability_catalog_content
This data source provides details about a specific Occ Availability Catalog Content resource in Oracle Cloud Infrastructure Capacity Management service.

Returns the binary contents of the availability catalog. Can be saved as a csv file.


## Example Usage

```hcl
data "oci_capacity_management_occ_availability_catalog_content" "test_occ_availability_catalog_content" {
	#Required
	occ_availability_catalog_id = oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id
}
```

## Argument Reference

The following arguments are supported:

* `occ_availability_catalog_id` - (Required) The OCID of the availability catalog.


## Attributes Reference

The following attributes are exported:


