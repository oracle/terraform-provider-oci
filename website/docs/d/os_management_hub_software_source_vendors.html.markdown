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

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This parameter is required and returns only resources contained within the specified compartment.
* `name` - (Optional) The name of the entity to be queried.


## Attributes Reference

The following attributes are exported:

* `software_source_vendor_collection` - The list of software_source_vendor_collection.

### SoftwareSourceVendor Reference

The following attributes are exported:

* `items` - List of software source vendors.
	* `arch_types` - List of corresponding architecture types.
	* `name` - Name of the vendor providing the software source.
	* `os_families` - List of corresponding operating system families.

