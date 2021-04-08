---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_service_catalog_associations"
sidebar_current: "docs-oci-datasource-service_catalog-service_catalog_associations"
description: |-
  Provides the list of Service Catalog Associations in Oracle Cloud Infrastructure Service Catalog service
---

# Data Source: oci_service_catalog_service_catalog_associations
This data source provides the list of Service Catalog Associations in Oracle Cloud Infrastructure Service Catalog service.

Lists all the resource associations for a specific service catalog.

## Example Usage

```hcl
data "oci_service_catalog_service_catalog_associations" "test_service_catalog_associations" {

	#Optional
	entity_id = oci_service_catalog_entity.test_entity.id
	entity_type = var.service_catalog_association_entity_type
	service_catalog_association_id = oci_service_catalog_service_catalog_association.test_service_catalog_association.id
	service_catalog_id = oci_service_catalog_service_catalog.test_service_catalog.id
}
```

## Argument Reference

The following arguments are supported:

* `entity_id` - (Optional) The unique identifier of the entity associated with service catalog.
* `entity_type` - (Optional) The type of the application in the service catalog.
* `service_catalog_association_id` - (Optional) The unique identifier for the service catalog association.
* `service_catalog_id` - (Optional) The unique identifier for the service catalog.


## Attributes Reference

The following attributes are exported:

* `service_catalog_association_collection` - The list of service_catalog_association_collection.

### ServiceCatalogAssociation Reference

The following attributes are exported:

* `entity_id` - Identifier of the entity being associated with service catalog.
* `entity_type` - The type of the entity that is associated with the service catalog.
* `id` - Identifier of the association.
* `service_catalog_id` - Identifier of the service catalog.
* `time_created` - Timestamp of when the resource was associated with service catalog.

