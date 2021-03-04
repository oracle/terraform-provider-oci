---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_publication"
sidebar_current: "docs-oci-resource-marketplace-publication"
description: |-
  Provides the Publication resource in Oracle Cloud Infrastructure Marketplace service
---

# oci_marketplace_publication
This resource provides the Publication resource in Oracle Cloud Infrastructure Marketplace service.

Creates a publication of the given type with an optional default package

## Example Usage

```hcl
resource "oci_marketplace_publication" "test_publication" {
	#Required
	compartment_id = var.compartment_id
	is_agreement_acknowledged = var.publication_is_agreement_acknowledged
	listing_type = var.publication_listing_type
	name = var.publication_name
	package_details {
		#Required
		eula {
			#Required
			eula_type = var.publication_package_details_eula_eula_type

			#Optional
			license_text = var.publication_package_details_eula_license_text
		}
		operating_system {

			#Optional
			name = var.publication_package_details_operating_system_name
		}
		package_type = var.publication_package_details_package_type
		package_version = var.publication_package_details_package_version

		#Optional
		image_id = oci_core_image.test_image.id
	}
	short_description = var.publication_short_description
	support_contacts {

		#Optional
		email = var.publication_support_contacts_email
		name = var.publication_support_contacts_name
		phone = var.publication_support_contacts_phone
		subject = var.publication_support_contacts_subject
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	long_description = var.publication_long_description
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment to create the resource within.
* `defined_tags` - (Optional) (Updatable) The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_agreement_acknowledged` - (Required) Acknowledgement that invoker has the right and authority to share this Community Image in accordance with their agreement with Oracle applicable to the Services and the related Service Specifications
* `listing_type` - (Required) In which catalog the listing should exist.
* `long_description` - (Optional) (Updatable) short description of the catalog listing
* `name` - (Required) (Updatable) The name of the listing.
* `package_details` - (Required) A base object for the properties of the package
	* `eula` - (Required) End User License Agreeement that a consumer of this listing has to accept
		* `eula_type` - (Required) the specified eula's type
		* `license_text` - (Optional) text of the eula
	* `image_id` - (Optional) base image id of the listing
	* `operating_system` - (Required) OS used by the listing.
		* `name` - (Optional) name of the operating system
	* `package_type` - (Required) Type of the artifact of the listing
	* `package_version` - (Required) The version of the package
* `short_description` - (Required) (Updatable) short description of the catalog listing
* `support_contacts` - (Required) (Updatable) Contact information to use to get support from the publisher for the listing.
	* `email` - (Optional) (Updatable) The email of the contact.
	* `name` - (Optional) (Updatable) The name of the contact.
	* `phone` - (Optional) (Updatable) The phone number of the contact.
	* `subject` - (Optional) (Updatable) The email subject line to use when contacting support.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The Compartment id where the listings exists
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `icon` - The model for upload data for images and icons.
	* `content_url` - The content URL of the upload data.
	* `file_extension` - The file extension of the upload data.
	* `mime_type` - The MIME type of the upload data.
	* `name` - The name used to refer to the upload data.
* `id` - The unique identifier for the listing in Marketplace.
* `listing_type` - In which catalog the listing should exist.
* `long_description` - A long description of the listing.
* `name` - The name of the listing.
* `package_type` - The listing's package type.
* `short_description` - A short description of the listing.
* `state` - The state of the listing in its lifecycle
* `support_contacts` - Contact information to use to get support from the publisher for the listing.
	* `email` - The email of the contact.
	* `name` - The name of the contact.
	* `phone` - The phone number of the contact.
	* `subject` - The email subject line to use when contacting support.
* `supported_operating_systems` - List of operating systems supprted.
	* `name` - name of the operating system
* `time_created` - The date and time this publication was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Import

Publications can be imported using the `id`, e.g.

```
$ terraform import oci_marketplace_publication.test_publication "id"
```

