---
subcategory: "Blockchain"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_blockchain_osns"
sidebar_current: "docs-oci-datasource-blockchain-osns"
description: |-
  Provides the list of Osns in Oracle Cloud Infrastructure Blockchain service
---

# Data Source: oci_blockchain_osns
This data source provides the list of Osns in Oracle Cloud Infrastructure Blockchain service.

List Blockchain Platform OSNs

## Example Usage

```hcl
data "oci_blockchain_osns" "test_osns" {
	#Required
	blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id

	#Optional
	display_name = var.osn_display_name
}
```

## Argument Reference

The following arguments are supported:

* `blockchain_platform_id` - (Required) Unique service identifier.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Example: `My new resource` 


## Attributes Reference

The following attributes are exported:

* `osn_collection` - The list of osn_collection.

### Osn Reference

The following attributes are exported:

* `ad` - Availability Domain of OSN
* `ocpu_allocation_param` - OCPU allocation parameter
	* `ocpu_allocation_number` - Number of OCPU allocation
* `osn_key` - OSN identifier
* `state` - The current state of the OSN.

