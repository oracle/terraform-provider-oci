---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_list_os_patches"
sidebar_current: "docs-oci-datasource-bds-bds_instance_list_os_patches"
description: |-
  Provides the list of Bds Instance List Os Patches in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_list_os_patches
This data source provides the list of Bds Instance List Os Patches in Oracle Cloud Infrastructure Big Data Service service.

List all available os patches for a given cluster

## Example Usage

```hcl
data "oci_bds_bds_instance_list_os_patches" "test_bds_instance_list_os_patches" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `os_patches` - The list of os_patches.

### BdsInstanceListOsPatch Reference

The following attributes are exported:

* `os_patch_version` - Patch version of the os patch.
* `release_date` - The time when the OS patch was released.

