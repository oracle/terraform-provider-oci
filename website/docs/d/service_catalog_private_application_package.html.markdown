---
subcategory: "Service Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_service_catalog_private_application_package"
sidebar_current: "docs-oci-datasource-service_catalog-private_application_package"
description: |-
  Provides details about a specific Private Application Package in Oracle Cloud Infrastructure Service Catalog service
---

# Data Source: oci_service_catalog_private_application_package
This data source provides details about a specific Private Application Package resource in Oracle Cloud Infrastructure Service Catalog service.

Gets the details of a specific package within a given private application.

## Example Usage

```hcl
data "oci_service_catalog_private_application_package" "test_private_application_package" {
	#Required
	private_application_package_id = oci_service_catalog_private_application_package.test_private_application_package.id
}
```

## Argument Reference

The following arguments are supported:

* `private_application_package_id` - (Required) The unique identifier for the private application package.


## Attributes Reference

The following attributes are exported:

* `content_url` - The content URL of the terraform configuration.
* `display_name` - The display name of the package.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private application package. 
* `mime_type` - The MIME type of the terraform configuration.
* `package_type` - The specified package's type.
* `private_application_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private application where the package is hosted. 
* `time_created` - The date and time the private application package was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2021-05-27T21:10:29.600Z` 
* `version` - The package version.

