---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_package_group"
sidebar_current: "docs-oci-datasource-os_management_hub-software_source_package_group"
description: |-
  Provides details about a specific Software Source Package Group in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_software_source_package_group
This data source provides details about a specific Software Source Package Group resource in Oracle Cloud Infrastructure Os Management Hub service.

Returns information about the specified package group from a software source.


## Example Usage

```hcl
data "oci_os_management_hub_software_source_package_group" "test_software_source_package_group" {
	#Required
	package_group_id = oci_identity_group.test_group.id
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
}
```

## Argument Reference

The following arguments are supported:

* `package_group_id` - (Required) The unique package group identifier.
* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


## Attributes Reference

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

