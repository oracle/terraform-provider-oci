---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_service_catalog_association"
sidebar_current: "docs-oci-datasource-service_catalog-service_catalog_association"
description: |-
  Provides details about a specific Service Catalog Association in Oracle Cloud Infrastructure Service Catalog service
---

# Data Source: oci_service_catalog_service_catalog_association
This data source provides details about a specific Service Catalog Association resource in Oracle Cloud Infrastructure Service Catalog service.

Gets detailed information about specific service catalog association.

## Example Usage

```hcl
data "oci_service_catalog_service_catalog_association" "test_service_catalog_association" {
	#Required
	service_catalog_association_id = oci_service_catalog_service_catalog_association.test_service_catalog_association.id
}
```

## Argument Reference

The following arguments are supported:

* `service_catalog_association_id` - (Required) The unique identifier of the service catalog association.


## Attributes Reference

The following attributes are exported:

* `entity_id` - Identifier of the entity being associated with service catalog.
* `entity_type` - The type of the entity that is associated with the service catalog.
* `id` - Identifier of the association.
* `service_catalog_id` - Identifier of the service catalog.
* `time_created` - Timestamp of when the resource was associated with service catalog.

