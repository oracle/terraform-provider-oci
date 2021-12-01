---
subcategory: "Blockchain"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_blockchain_blockchain_platform_patch"
sidebar_current: "docs-oci-datasource-blockchain-blockchain_platform_patch"
description: |-
  Provides details about a specific Blockchain Platform Patch in Oracle Cloud Infrastructure Blockchain service
---

# Data Source: oci_blockchain_blockchain_platform_patch
This data source provides details about a specific Blockchain Platform Patch resource in Oracle Cloud Infrastructure Blockchain service.

List Blockchain Platform Patches

## Example Usage

```hcl
data "oci_blockchain_blockchain_platform_patch" "test_blockchain_platform_patch" {
	#Required
	blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id
}
```

## Argument Reference

The following arguments are supported:

* `blockchain_platform_id` - (Required) Unique service identifier.


## Attributes Reference

The following attributes are exported:

* `items` - Collection of PatchSummary
	* `id` - patch id
	* `patch_info_url` - A URL for the patch specific documentation
	* `service_version` - patch service version
	* `time_patch_due` - patch due date for customer initiated patching

