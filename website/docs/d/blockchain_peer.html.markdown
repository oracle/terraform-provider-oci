---
subcategory: "Blockchain"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_blockchain_peer"
sidebar_current: "docs-oci-datasource-blockchain-peer"
description: |-
  Provides details about a specific Peer in Oracle Cloud Infrastructure Blockchain service
---

# Data Source: oci_blockchain_peer
This data source provides details about a specific Peer resource in Oracle Cloud Infrastructure Blockchain service.

Gets information about a peer identified by the specific id

## Example Usage

```hcl
data "oci_blockchain_peer" "test_peer" {
	#Required
	blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id
	peer_id = oci_blockchain_peer.test_peer.id
}
```

## Argument Reference

The following arguments are supported:

* `blockchain_platform_id` - (Required) Unique service identifier.
* `peer_id` - (Required) Peer identifier.


## Attributes Reference

The following attributes are exported:

* `ad` - Availability Domain of peer
* `alias` - peer alias
* `host` - Host on which the Peer exists
* `ocpu_allocation_param` - 
	* `ocpu_allocation_number` - Number of OCPU allocation
* `peer_key` - peer identifier
* `role` - Peer role
* `state` - The current state of the peer.

