---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_all_applications"
sidebar_current: "docs-oci-datasource-service_catalog-all_applications"
description: |-
  Provides the list of All Applications in Oracle Cloud Infrastructure Service Catalog service
---

# Data Source: oci_service_catalog_all_applications
This data source provides the list of All Applications in Oracle Cloud Infrastructure Service Catalog service.

Lists all the available listings and private applications in a compartment.
A new API for catalog manager use when creating/updating a service catalog.


## Example Usage

```hcl
data "oci_service_catalog_all_applications" "test_all_applications" {
	#Required
  	compartment_id = var.compartment_id

	#Optional
	display_name = var.all_application_display_name
	entity_id = oci_service_catalog_entity.test_entity.id
	entity_type = var.all_application_entity_type
	is_featured = var.all_application_is_featured
	package_type = var.all_application_package_type
	pricing = var.all_application_pricing
	publisher_id = oci_marketplace_publisher.test_publisher.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The unique identifier for the compartment.
* `display_name` - (Optional) Exact match name filter.
* `entity_id` - (Optional) The unique identifier of the entity associated with service catalog.
* `entity_type` - (Optional) The type of the application in the service catalog.
* `is_featured` - (Optional) Indicates whether to show only featured resources. If this is set to `false` or is omitted, then all resources will be returned. 
* `package_type` - (Optional) Name of the package type. If multiple package types are provided, then any resource with one or more matching package types will be returned. 
* `pricing` - (Optional) Name of the pricing type. If multiple pricing types are provided, then any resource with one or more matching pricing models will be returned. 
* `publisher_id` - (Optional) Limit results to just this publisher.


## Attributes Reference

The following attributes are exported:

* `application_collection` - The list of application_collection.

### AllApplication Reference

The following attributes are exported:

* `items` - Collection of service catalog applications.
	* `categories` - Product categories that the application belongs to.
	* `display_name` - The name that service catalog should use to display this application.
	* `entity_id` - Identifier of the application from a service catalog.
	* `entity_type` - The type of an application in the service catalog.
	* `is_featured` - Indicates whether the application is featured.
	* `logo` - The model for uploaded binary data, like logos and images.
		* `content_url` - The content URL of the uploaded data.
		* `display_name` - The name used to refer to the uploaded data.
		* `mime_type` - The MIME type of the uploaded data.
	* `package_type` - The type of the packages withing the application.
	* `pricing_type` - Summary of the pricing types available across all packages in the application.
	* `publisher` - Summary details about the publisher of the resource.
		* `display_name` - The name of the publisher.
		* `id` - The unique identifier for the publisher.
	* `short_description` - A short description of the application.
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 

