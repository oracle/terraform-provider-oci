---
subcategory: "Blockchain"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_blockchain_osn"
sidebar_current: "docs-oci-resource-blockchain-osn"
description: |-
  Provides the Osn resource in Oracle Cloud Infrastructure Blockchain service
---

# oci_blockchain_osn
This resource provides the Osn resource in Oracle Cloud Infrastructure Blockchain service.

Create Blockchain Platform Osn

## Example Usage

```hcl
resource "oci_blockchain_osn" "test_osn" {
	#Required
	ad = var.osn_ad
	blockchain_platform_id = oci_blockchain_blockchain_platform.test_blockchain_platform.id

	#Optional
	ocpu_allocation_param {
		#Required
		ocpu_allocation_number = var.osn_ocpu_allocation_param_ocpu_allocation_number
	}
}
```

## Argument Reference

The following arguments are supported:

* `ad` - (Required) Availability Domain to place new OSN
* `blockchain_platform_id` - (Required) Unique service identifier.
* `ocpu_allocation_param` - (Optional) (Updatable) OCPU allocation parameter
	* `ocpu_allocation_number` - (Required) (Updatable) Number of OCPU allocation


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `ad` - Availability Domain of OSN
* `ocpu_allocation_param` - OCPU allocation parameter
	* `ocpu_allocation_number` - Number of OCPU allocation
* `osn_key` - OSN identifier
* `state` - The current state of the OSN.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 30 minutes), when creating the Osn
	* `update` - (Defaults to 30 minutes), when updating the Osn
	* `delete` - (Defaults to 30 minutes), when destroying the Osn


## Import

Osns can be imported using the `id`, e.g.

```
$ terraform import oci_blockchain_osn.test_osn "blockchainPlatforms/{blockchainPlatformId}/osns/{osnId}" 
```

