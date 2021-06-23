---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_service_catalog"
sidebar_current: "docs-oci-resource-service_catalog-service_catalog"
description: |-
  Provides the Service Catalog resource in Oracle Cloud Infrastructure Service Catalog service
---

# oci_service_catalog_service_catalog
This resource provides the Service Catalog resource in Oracle Cloud Infrastructure Service Catalog service.

Creates a brand new service catalog in a given compartment.

## Example Usage

```hcl
resource "oci_service_catalog_service_catalog" "test_service_catalog" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.service_catalog_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The unique identifier for the compartment where the service catalog will be created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) The display name of the service catalog.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Service Catalog
	* `update` - (Defaults to 20 minutes), when updating the Service Catalog
	* `delete` - (Defaults to 20 minutes), when destroying the Service Catalog


## Import

ServiceCatalogs can be imported using the `id`, e.g.

```
$ terraform import oci_service_catalog_service_catalog.test_service_catalog "id"
```

