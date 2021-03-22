---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_service_catalog_association"
sidebar_current: "docs-oci-resource-service_catalog-service_catalog_association"
description: |-
  Provides the Service Catalog Association resource in Oracle Cloud Infrastructure Service Catalog service
---

# oci_service_catalog_service_catalog_association
This resource provides the Service Catalog Association resource in Oracle Cloud Infrastructure Service Catalog service.

Creates an association between service catalog and a resource.

## Example Usage

```hcl
resource "oci_service_catalog_service_catalog_association" "test_service_catalog_association" {
	#Required
	entity_id = oci_service_catalog_entity.test_entity.id
	service_catalog_id = oci_service_catalog_service_catalog.test_service_catalog.id

	#Optional
	entity_type = var.service_catalog_association_entity_type
}
```

## Argument Reference

The following arguments are supported:

* `entity_id` - (Required) Identifier of the entity being associated with service catalog.
* `entity_type` - (Optional) The type of the entity that is associated with the service catalog.
* `service_catalog_id` - (Required) Identifier of the service catalog.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `entity_id` - Identifier of the entity being associated with service catalog.
* `entity_type` - The type of the entity that is associated with the service catalog.
* `id` - Identifier of the association.
* `service_catalog_id` - Identifier of the service catalog.
* `time_created` - Timestamp of when the resource was associated with service catalog.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Service Catalog Association
	* `update` - (Defaults to 20 minutes), when updating the Service Catalog Association
	* `delete` - (Defaults to 20 minutes), when destroying the Service Catalog Association


## Import

ServiceCatalogAssociations can be imported using the `id`, e.g.

```
$ terraform import oci_service_catalog_service_catalog_association.test_service_catalog_association "id"
```

