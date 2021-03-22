---
subcategory: "Blockchain"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_blockchain_peer"
sidebar_current: "docs-oci-resource-blockchain-peer"
description: |-
  Provides the Peer resource in Oracle Cloud Infrastructure Blockchain service
---

# oci_blockchain_peer
This resource provides the Peer resource in Oracle Cloud Infrastructure Blockchain service.

Create Blockchain Platform Peer

## Example Usage

```hcl
resource "oci_blockchain_peer" "test_peer" {
	#Required
	ad = var.peer_ad
	blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id
	ocpu_allocation_param {
		#Required
		ocpu_allocation_number = var.peer_ocpu_allocation_param_ocpu_allocation_number
	}
	role = var.peer_role

	#Optional
	alias = var.peer_alias
}
```

## Argument Reference

The following arguments are supported:

* `ad` - (Required) Availability Domain to place new peer
* `alias` - (Optional) peer alias
* `blockchain_platform_id` - (Required) Unique service identifier.
* `ocpu_allocation_param` - (Required) (Updatable) OCPU allocation parameter
	* `ocpu_allocation_number` - (Required) (Updatable) Number of OCPU allocation
* `role` - (Required) Peer role


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `ad` - Availability Domain of peer
* `alias` - peer alias
* `host` - Host on which the Peer exists
* `ocpu_allocation_param` - OCPU allocation parameter
	* `ocpu_allocation_number` - Number of OCPU allocation
* `peer_key` - peer identifier
* `role` - Peer role
* `state` - The current state of the peer.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 30 minutes), when creating the Peer
	* `update` - (Defaults to 30 minutes), when updating the Peer
	* `delete` - (Defaults to 30 minutes), when destroying the Peer


## Import

Peers can be imported using the `id`, e.g.

```
$ terraform import oci_blockchain_peer.test_peer "blockchainPlatforms/{blockchainPlatformId}/peers/{peerId}" 
```

