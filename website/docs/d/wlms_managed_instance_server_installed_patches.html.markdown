---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_managed_instance_server_installed_patches"
sidebar_current: "docs-oci-datasource-wlms-managed_instance_server_installed_patches"
description: |-
  Provides the list of Managed Instance Server Installed Patches in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_managed_instance_server_installed_patches
This data source provides the list of Managed Instance Server Installed Patches in Oracle Cloud Infrastructure Wlms service.

Gets a list of installed patches on a server in a managed instance.


## Example Usage

```hcl
data "oci_wlms_managed_instance_server_installed_patches" "test_managed_instance_server_installed_patches" {
	#Required
	managed_instance_id = oci_wlms_managed_instance.test_managed_instance.id
	server_id = oci_wlms_server.test_server.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
* `server_id` - (Required) The unique identifier of a server.

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 


## Attributes Reference

The following attributes are exported:

* `installed_patch_collection` - The list of installed_patch_collection.

### ManagedInstanceServerInstalledPatch Reference

The following attributes are exported:

* `items` - List of installed patches per server
	* `description` - The description of the WebLogic patch.
	* `display_name` - The name of the WebLogic patch.
	* `id` - The ID of the WebLogic patch.

		**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 

