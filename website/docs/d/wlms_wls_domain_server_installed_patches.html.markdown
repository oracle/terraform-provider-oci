---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_wls_domain_server_installed_patches"
sidebar_current: "docs-oci-datasource-wlms-wls_domain_server_installed_patches"
description: |-
  Provides the list of Wls Domain Server Installed Patches in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_wls_domain_server_installed_patches
This data source provides the list of Wls Domain Server Installed Patches in Oracle Cloud Infrastructure Wlms service.

Gets a list of installed patches on a server for a domain.


## Example Usage

```hcl
data "oci_wlms_wls_domain_server_installed_patches" "test_wls_domain_server_installed_patches" {
	#Required
	server_id = oci_wlms_server.test_server.id
	wls_domain_id = oci_wlms_wls_domain.test_wls_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `server_id` - (Required) The unique identifier of a server.

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `wls_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.


## Attributes Reference

The following attributes are exported:

* `installed_patch_collection` - The list of installed_patch_collection.

### WlsDomainServerInstalledPatch Reference

The following attributes are exported:

* `items` - List of installed patches per server
	* `description` - The description of the WebLogic patch.
	* `display_name` - The name of the WebLogic patch.
	* `id` - The ID of the WebLogic patch.

		**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 

