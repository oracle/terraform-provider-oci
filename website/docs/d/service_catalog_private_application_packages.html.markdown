---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_private_application_packages"
sidebar_current: "docs-oci-datasource-service_catalog-private_application_packages"
description: |-
  Provides the list of Private Application Packages in Oracle Cloud Infrastructure Service Catalog service
---

# Data Source: oci_service_catalog_private_application_packages
This data source provides the list of Private Application Packages in Oracle Cloud Infrastructure Service Catalog service.

Lists the packages in the specified private application.

## Example Usage

```hcl
data "oci_service_catalog_private_application_packages" "test_private_application_packages" {
	#Required
	private_application_id = oci_service_catalog_private_application.test_private_application.id

	#Optional
	display_name = var.private_application_package_display_name
	package_type = var.private_application_package_package_type
	private_application_package_id = oci_service_catalog_private_application_package.test_private_application_package.id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) Exact match name filter.
* `package_type` - (Optional) Name of the package type. If multiple package types are provided, then any resource with one or more matching package types will be returned. 
* `private_application_id` - (Required) The unique identifier for the private application.
* `private_application_package_id` - (Optional) The unique identifier for the private application package.


## Attributes Reference

The following attributes are exported:

* `private_application_package_collection` - The list of private_application_package_collection.

### PrivateApplicationPackage Reference

The following attributes are exported:

* `content_url` - The content URL of the terraform configuration.
* `display_name` - The display name of the package.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private application package. 
* `mime_type` - The MIME type of the terraform configuration.
* `package_type` - The specified package's type.
* `private_application_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private application where the package is hosted. 
* `time_created` - The date and time the private application package was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2021-05-27T21:10:29.600Z` 
* `version` - The package version.

