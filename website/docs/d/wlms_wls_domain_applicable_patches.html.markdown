---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_wls_domain_applicable_patches"
sidebar_current: "docs-oci-datasource-wlms-wls_domain_applicable_patches"
description: |-
  Provides the list of Wls Domain Applicable Patches in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_wls_domain_applicable_patches
This data source provides the list of Wls Domain Applicable Patches in Oracle Cloud Infrastructure Wlms service.

Gets the latest patches that can be installed to the WebLogic domains.


## Example Usage

```hcl
data "oci_wlms_wls_domain_applicable_patches" "test_wls_domain_applicable_patches" {
	#Required
	wls_domain_id = oci_wlms_wls_domain.test_wls_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `wls_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.


## Attributes Reference

The following attributes are exported:

* `applicable_patch_collection` - The list of applicable_patch_collection.

### WlsDomainApplicablePatch Reference

The following attributes are exported:

* `items` - List of patches per WebLogic version and middleware type.
	* `description` - The description of the WebLogic patch.
	* `display_name` - The name of the WebLogic patch.
	* `id` - The ID of the WebLogic patch.
	* `middleware_type` - The type of middleware for which this patch is applicable. A patch can be applicable to more than one type of middleware.
	* `os_arch` - The operating system architecture for which the patch can be applied.
	* `weblogic_version` - The WebLogic version for this patch. The patch can be installed to domains with this version.

