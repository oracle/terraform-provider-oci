---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_windows_update"
sidebar_current: "docs-oci-datasource-os_management_hub-windows_update"
description: |-
  Provides details about a specific Windows Update in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_windows_update
This data source provides details about a specific Windows Update resource in Oracle Cloud Infrastructure Os Management Hub service.

Returns a Windows Update object.


## Example Usage

```hcl
data "oci_os_management_hub_windows_update" "test_windows_update" {
	#Required
	windows_update_id = oci_os_management_hub_windows_update.test_windows_update.id
}
```

## Argument Reference

The following arguments are supported:

* `windows_update_id` - (Required) The unique identifier for the Windows update. Note that this is not an OCID, but is a unique identifier assigned by Microsoft.  Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed' 


## Attributes Reference

The following attributes are exported:

* `description` - Description of the update.
* `installable` - Indicates whether the update can be installed using the service.
* `installation_requirements` - List of requirements for installing the update on the managed instance.
* `is_reboot_required_for_installation` - Indicates whether a reboot is required to complete the installation of this update.
* `kb_article_ids` - List of the Microsoft Knowledge Base Article Ids related to this Windows Update.
* `name` - Name of the Windows update.
* `size_in_bytes` - size of the package in bytes
* `update_id` - Unique identifier for the Windows update. Note that this is not an OCID, but is a unique identifier assigned by Microsoft.  Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed' 
* `update_type` - The type of Windows update.

