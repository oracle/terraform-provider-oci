---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_marketplace_external_attested_metadata"
sidebar_current: "docs-oci-resource-marketplace-marketplace_external_attested_metadata"
description: |-
  Provides the Marketplace External Attested Metadata resource in Oracle Cloud Infrastructure Marketplace service
---

# oci_marketplace_marketplace_external_attested_metadata
This resource provides the Marketplace External Attested Metadata resource in Oracle Cloud Infrastructure Marketplace service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/marketplace/latest/MarketplaceExternalAttestedMetadatum

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/marketplace

Generates attested marketplace metadata

## Example Usage

```hcl
resource "oci_marketplace_marketplace_external_attested_metadata" "test_marketplace_external_attested_metadata" {
	#Required
	compartment_id = var.compartment_id
	instance_id = oci_core_instance.test_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) compartment that associated instance is in
* `instance_id` - (Required) unique id that identifies the associated instance


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `serialized_jwt` - The serialized JWT token, containing header, payload, signature

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Marketplace External Attested Metadata
	* `update` - (Defaults to 20 minutes), when updating the Marketplace External Attested Metadata
	* `delete` - (Defaults to 20 minutes), when destroying the Marketplace External Attested Metadata


## Import

MarketplaceExternalAttestedMetadata can be imported using the `id`, e.g.

```
$ terraform import oci_marketplace_marketplace_external_attested_metadata.test_marketplace_external_attested_metadata "id"
```

