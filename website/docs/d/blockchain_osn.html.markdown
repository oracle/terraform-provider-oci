---
subcategory: "Blockchain"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_blockchain_osn"
sidebar_current: "docs-oci-datasource-blockchain-osn"
description: |-
  Provides details about a specific Osn in Oracle Cloud Infrastructure Blockchain service
---

# Data Source: oci_blockchain_osn
This data source provides details about a specific Osn resource in Oracle Cloud Infrastructure Blockchain service.

Gets information about an OSN identified by the specific id

## Example Usage

```hcl
data "oci_blockchain_osn" "test_osn" {
	#Required
	blockchain_platform_id = "${oci_blockchain_blockchain_platform.test_blockchain_platform.id}"
	osn_id = "${oci_blockchain_osn.test_osn.id}"
}
```

## Argument Reference

The following arguments are supported:

* `blockchain_platform_id` - (Required) Unique service identifier.
* `osn_id` - (Required) OSN identifier.


## Attributes Reference

The following attributes are exported:

* `ad` - Availability Domain of OSN
* `ocpu_allocation_param` - 
	* `ocpu_allocation_number` - Number of OCPU allocation
* `osn_key` - OSN identifier
* `state` - The current state of the OSN.

