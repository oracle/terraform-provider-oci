---
subcategory: "Blockchain"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_blockchain_blockchain_platform_patches"
sidebar_current: "docs-oci-datasource-blockchain-blockchain_platform_patches"
description: |-
  Provides the list of Blockchain Platform Patches in Oracle Cloud Infrastructure Blockchain service
---

# Data Source: oci_blockchain_blockchain_platform_patches
This data source provides the list of Blockchain Platform Patches in Oracle Cloud Infrastructure Blockchain service.

List Blockchain Platform Patches

## Example Usage

```hcl
data "oci_blockchain_blockchain_platform_patches" "test_blockchain_platform_patches" {
	#Required
	blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id
}
```

## Argument Reference

The following arguments are supported:

* `blockchain_platform_id` - (Required) Unique service identifier.


## Attributes Reference

The following attributes are exported:

* `blockchain_platform_patch_collection` - The list of blockchain_platform_patch_collection.

### BlockchainPlatformPatch Reference

The following attributes are exported:

* `items` - Collection of PatchSummary
	* `id` - patch id
	* `patch_info_url` - A URL for the patch specific documentation
	* `service_version` - patch service version
	* `time_patch_due` - patch due date for customer initiated patching

