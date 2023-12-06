---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_vendors"
sidebar_current: "docs-oci-datasource-os_management_hub-software_source_vendors"
description: |-
  Provides the list of Software Source Vendors in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_source_vendors
This data source provides the list of Software Source Vendors in Oracle Cloud Infrastructure Os Management Hub service.

Lists available software source vendors. Filter the list against a variety of criteria including but not limited
to its name.


## Example Usage

```hcl
data "oci_os_management_hub_software_source_vendors" "test_software_source_vendors" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.software_source_vendor_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment that contains the resources to list. This parameter is required.
* `name` - (Optional) The name of the entity to be queried.


## Attributes Reference

The following attributes are exported:

* `software_source_vendor_collection` - The list of software_source_vendor_collection.

### SoftwareSourceVendor Reference

The following attributes are exported:

* `items` - List of SoftwareSourceVendor.
	* `arch_types` - List of corresponding archTypes.
	* `name` - Name of the vendor providing the software source.
	* `os_families` - List of corresponding osFamilies.

