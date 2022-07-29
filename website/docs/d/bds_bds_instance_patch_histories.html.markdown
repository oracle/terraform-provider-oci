---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_patch_histories"
sidebar_current: "docs-oci-datasource-bds-bds_instance_patch_histories"
description: |-
  Provides the list of Bds Instance Patch Histories in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_patch_histories
This data source provides the list of Bds Instance Patch Histories in Oracle Cloud Infrastructure Big Data Service service.

List the patch history of this cluster.


## Example Usage

```hcl
data "oci_bds_bds_instance_patch_histories" "test_bds_instance_patch_histories" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id

	#Optional
	patch_version = var.bds_instance_patch_history_patch_version
	state = var.bds_instance_patch_history_state
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `patch_version` - (Optional) The version of the patch
* `state` - (Optional) The status of the patch.


## Attributes Reference

The following attributes are exported:

* `patch_histories` - The list of patch_histories.

### BdsInstancePatchHistory Reference

The following attributes are exported:

* `state` - The status of this patch.
* `time_updated` - The time when the patch history was last updated.
* `version` - The version of the patch.

