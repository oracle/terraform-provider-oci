---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_patches"
sidebar_current: "docs-oci-datasource-bds-bds_instance_patches"
description: |-
  Provides the list of Bds Instance Patches in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_patches
This data source provides the list of Bds Instance Patches in Oracle Cloud Infrastructure Big Data Service service.

List all the available patches for this cluster.


## Example Usage

```hcl
data "oci_bds_bds_instance_patches" "test_bds_instance_patches" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `patches` - The list of patches.

### BdsInstancePatch Reference

The following attributes are exported:

* `time_released` - The time when the patch was released.
* `version` - The version of the patch.

