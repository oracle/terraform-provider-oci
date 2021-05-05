---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_service_catalog"
sidebar_current: "docs-oci-datasource-service_catalog-service_catalog"
description: |-
  Provides details about a specific Service Catalog in Oracle Cloud Infrastructure Service Catalog service
---

# Data Source: oci_service_catalog_service_catalog
This data source provides details about a specific Service Catalog resource in Oracle Cloud Infrastructure Service Catalog service.

Gets detailed information about the service catalog including name, compartmentId

## Example Usage

```hcl
data "oci_service_catalog_service_catalog" "test_service_catalog" {
	#Required
	service_catalog_id = oci_service_catalog_service_catalog.test_service_catalog.id
}
```

## Argument Reference

The following arguments are supported:

* `service_catalog_id` - (Required) The unique identifier for the service catalog.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The Compartment id where the service catalog exists
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The name of the service catalog.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The unique identifier for the Service catalog.
* `state` - The lifecycle state of the service catalog.
* `time_created` - The date and time the service catalog was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2021-05-26T21:10:29.600Z` 
* `time_updated` - The date and time the service catalog was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2021-12-10T05:10:29.721Z` 

