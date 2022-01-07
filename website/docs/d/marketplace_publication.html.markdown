---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_publication"
sidebar_current: "docs-oci-datasource-marketplace-publication"
description: |-
  Provides details about a specific Publication in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_publication
This data source provides details about a specific Publication resource in Oracle Cloud Infrastructure Marketplace service.

Gets the details of the specified publication.

## Example Usage

```hcl
data "oci_marketplace_publication" "test_publication" {
	#Required
	publication_id = oci_marketplace_publication.test_publication.id
}
```

## Argument Reference

The following arguments are supported:

* `publication_id` - (Required) The unique identifier for the publication.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the publication exists.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `icon` - The model for upload data for images and icons.
	* `content_url` - The content URL of the upload data.
	* `file_extension` - The file extension of the upload data.
	* `mime_type` - The MIME type of the upload data.
	* `name` - The name used to refer to the upload data.
* `id` - The unique identifier for the publication in Marketplace.
* `listing_type` - The publisher category to which the publication belongs. The publisher category informs where the listing appears for use.
* `long_description` - A long description of the publication to use in the listing.
* `name` - The name of the publication, which is also used in the listing.
* `package_type` - The listing's package type.
* `short_description` - A short description of the publication to use in the listing.
* `state` - The lifecycle state of the publication.
* `support_contacts` - Contact information for getting support from the publisher for the listing.
	* `email` - The email of the contact.
	* `name` - The name of the contact.
	* `phone` - The phone number of the contact.
	* `subject` - The email subject line to use when contacting support.
* `supported_operating_systems` - The list of operating systems supported by the listing.
	* `name` - The name of the operating system.
* `time_created` - The date and time the publication was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

