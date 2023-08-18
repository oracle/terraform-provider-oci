---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_get_os_patch"
sidebar_current: "docs-oci-datasource-bds-bds_instance_get_os_patch"
description: |-
  Provides the list of Bds Instance Get Os Patch in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_get_os_patch
This data source provides the list of Bds Instance Get Os Patch in Oracle Cloud Infrastructure Big Data Service service.

Get the details of an os patch

## Example Usage

```hcl
data "oci_bds_bds_instance_get_os_patch" "test_bds_instance_get_os_patch" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	os_patch_version = var.bds_instance_get_os_patch_os_patch_version
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `os_patch_version` - (Required) The version of the OS patch.


## Attributes Reference

The following attributes are exported:

* `os_patch_details` - The list of os_patch_details.

### BdsInstanceGetOsPatch Reference

The following attributes are exported:

* `min_bds_version` - Minimum BDS version required to install current OS patch.
* `min_compatible_odh_version_map` - Map of major ODH version to minimum ODH version required to install current OS patch. e.g. {ODH0.9: 0.9.1} 
* `os_patch_version` - Version of the os patch.
* `patch_type` - Type of a specific os patch. REGULAR means standard released os patches. CUSTOM means os patches with some customizations. EMERGENT means os patches with some emergency fixes that should be prioritized. 
* `release_date` - Released date of the OS patch.
* `target_packages` - List of summaries of individual target packages.
	* `package_name` - The package's name.
	* `related_cv_es` - Related CVEs of the package update.
	* `target_version` - The target version of the package.
	* `update_type` - The action that current package will be executed on the cluster.

