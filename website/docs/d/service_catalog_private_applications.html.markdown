---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_private_applications"
sidebar_current: "docs-oci-datasource-service_catalog-private_applications"
description: |-
  Provides the list of Private Applications in Oracle Cloud Infrastructure Service Catalog service
---

# Data Source: oci_service_catalog_private_applications
This data source provides the list of Private Applications in Oracle Cloud Infrastructure Service Catalog service.

Lists all the private applications in a given compartment.

## Example Usage

```hcl
data "oci_service_catalog_private_applications" "test_private_applications" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.private_application_display_name
	private_application_id = oci_service_catalog_private_application.test_private_application.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The unique identifier for the compartment.
* `display_name` - (Optional) Exact match name filter.
* `private_application_id` - (Optional) The unique identifier for the private application.


## Attributes Reference

The following attributes are exported:

* `private_application_collection` - The list of private_application_collection.

### PrivateApplication Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the private application resides. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The name of the private application.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The unique identifier for the private application in Marketplace.
* `logo` - The model for uploaded binary data, like logos and images.
	* `content_url` - The content URL of the uploaded data.
	* `display_name` - The name used to refer to the uploaded data.
	* `mime_type` - The MIME type of the uploaded data.
* `long_description` - A long description of the private application.
* `package_type` - Type of packages within this private application.
* `short_description` - A short description of the private application.
* `state` - The lifecycle state of the private application.
* `time_created` - The date and time the private application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2021-05-26T21:10:29.600Z` 
* `time_updated` - The date and time the private application was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2021-12-10T05:10:29.721Z` 

