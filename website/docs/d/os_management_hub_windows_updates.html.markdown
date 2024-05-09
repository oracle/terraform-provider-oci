---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_windows_updates"
sidebar_current: "docs-oci-datasource-os_management_hub-windows_updates"
description: |-
  Provides the list of Windows Updates in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_windows_updates
This data source provides the list of Windows Updates in Oracle Cloud Infrastructure Os Management Hub service.

Lists Windows updates that have been reported to the service.


## Example Usage

```hcl
data "oci_os_management_hub_windows_updates" "test_windows_updates" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	classification_type = var.windows_update_classification_type
	display_name_contains = var.windows_update_display_name_contains
	name = var.windows_update_name
}
```

## Argument Reference

The following arguments are supported:

* `classification_type` - (Optional) A filter to return only packages that match the given update classification type.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This parameter is required and returns only resources contained within the specified compartment.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `name` - (Optional) A filter based on the unique identifier for the Windows update. Note that this is not an OCID, but is a unique identifier assigned by Microsoft.  Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed' 


## Attributes Reference

The following attributes are exported:

* `windows_update_collection` - The list of windows_update_collection.

### WindowsUpdate Reference

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

