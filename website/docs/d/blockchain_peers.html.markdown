---
subcategory: "Blockchain"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_blockchain_peers"
sidebar_current: "docs-oci-datasource-blockchain-peers"
description: |-
  Provides the list of Peers in Oracle Cloud Infrastructure Blockchain service
---

# Data Source: oci_blockchain_peers
This data source provides the list of Peers in Oracle Cloud Infrastructure Blockchain service.

List Blockchain Platform Peers

## Example Usage

```hcl
data "oci_blockchain_peers" "test_peers" {
	#Required
	blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id

	#Optional
	display_name = var.peer_display_name
}
```

## Argument Reference

The following arguments are supported:

* `blockchain_platform_id` - (Required) Unique service identifier.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Example: `My new resource` 


## Attributes Reference

The following attributes are exported:

* `peer_collection` - The list of peer_collection.

### Peer Reference

The following attributes are exported:

* `ad` - Availability Domain of peer
* `alias` - peer alias
* `host` - Host on which the Peer exists
* `ocpu_allocation_param` - OCPU allocation parameter
	* `ocpu_allocation_number` - Number of OCPU allocation
* `peer_key` - peer identifier
* `role` - Peer role
* `state` - The current state of the peer.

