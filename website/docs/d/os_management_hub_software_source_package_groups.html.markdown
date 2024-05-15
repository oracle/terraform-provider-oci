---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_package_groups"
sidebar_current: "docs-oci-datasource-os_management_hub-software_source_package_groups"
description: |-
  Provides the list of Software Source Package Groups in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_source_package_groups
This data source provides the list of Software Source Package Groups in Oracle Cloud Infrastructure Os Management Hub service.

Lists package groups that are associated with the specified software source [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Filter the list against a
variety of criteria including but not limited to its name, and package group type.


## Example Usage

```hcl
data "oci_os_management_hub_software_source_package_groups" "test_software_source_package_groups" {
	#Required
	software_source_id = oci_os_management_hub_software_source.test_software_source.id

	#Optional
	compartment_id = var.compartment_id
	group_type = var.software_source_package_group_group_type
	name = var.software_source_package_group_name
	name_contains = var.software_source_package_group_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `group_type` - (Optional) A filter to return only package groups of the specified type.
* `name` - (Optional) The name of the entity to be queried.
* `name_contains` - (Optional) A filter to return resources that may partially match the name given.
* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


## Attributes Reference

The following attributes are exported:

* `package_group_collection` - The list of package_group_collection.

### SoftwareSourcePackageGroup Reference

The following attributes are exported:

* `description` - Description of the package group.
* `display_order` - Indicates the order to display category or environment.
* `group_type` - Indicates if this is a group, category, or environment.
* `id` - Package group identifier.
* `is_default` - Indicates if this package group is the default.
* `is_user_visible` - Indicates if this package group is visible to users.
* `name` - Package group name.
* `packages` - The list of packages in the package group.
* `repositories` - The repository IDs of the package group's repositories.

