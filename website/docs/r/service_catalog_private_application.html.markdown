---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_private_application"
sidebar_current: "docs-oci-resource-service_catalog-private_application"
description: |-
  Provides the Private Application resource in Oracle Cloud Infrastructure Service Catalog service
---

# oci_service_catalog_private_application
This resource provides the Private Application resource in Oracle Cloud Infrastructure Service Catalog service.

Creates a private application along with a single package to be hosted.

## Example Usage

```hcl
resource "oci_service_catalog_private_application" "test_private_application" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.private_application_display_name
	package_details {
		#Required
		package_type = var.private_application_package_details_package_type
		version = var.private_application_package_details_version

		#Optional
		zip_file_base64encoded = var.private_application_package_details_zip_file_base64encoded
	}
	short_description = var.private_application_short_description

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	logo_file_base64encoded = var.private_application_logo_file_base64encoded
	long_description = var.private_application_long_description
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the private application. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) The name of the private application.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `logo_file_base64encoded` - (Optional) (Updatable) Base64-encoded logo to use as the private application icon. Template icon file requirements: PNG format, 50 KB maximum, 130 x 130 pixels. 
* `long_description` - (Optional) (Updatable) A long description of the private application.
* `package_details` - (Required) A base object for creating a private application package.
	* `package_type` - (Required) The package's type.
	* `version` - (Required) The package version.
	* `zip_file_base64encoded` - (Optional) Base-64 payload of the Terraform zip package.
* `short_description` - (Required) (Updatable) A short description of the private application.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Private Application
	* `update` - (Defaults to 20 minutes), when updating the Private Application
	* `delete` - (Defaults to 20 minutes), when destroying the Private Application


## Import

PrivateApplications can be imported using the `id`, e.g.

```
$ terraform import oci_service_catalog_private_application.test_private_application "id"
```

