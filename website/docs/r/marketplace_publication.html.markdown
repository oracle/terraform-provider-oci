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

Creates a publication of the specified listing type with an optional default package.

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

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the publication.
* `defined_tags` - (Optional) (Updatable) The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_agreement_acknowledged` - (Required) Whether the publisher acknowledged that they have the right and authority to share the contents of the publication and that they accepted the Oracle terms of use agreements required to create a publication. 
* `listing_type` - (Required) The publisher category to which the publication belongs. The publisher category informs where the listing appears for use. 
* `long_description` - (Optional) (Updatable) A long description of the publication to use in the listing.
* `name` - (Required) (Updatable) The name of the publication, which is also used in the listing.
* `package_details` - (Required) A base object for creating a publication package.
	* `eula` - (Required) The end user license agreeement (EULA) that consumers of this listing must accept.
		* `eula_type` - (Required) The end user license agreement's type.
		* `license_text` - (Optional) The text of the end user license agreement.
	* `image_id` - (Optional) The unique identifier for the base image of the publication.
	* `operating_system` - (Required) The operating system used by the listing.
		* `name` - (Optional) The name of the operating system.
	* `package_type` - (Required) The package's type.
	* `package_version` - (Required) The package version.
* `short_description` - (Required) (Updatable) A short description of the publication to use in the listing.
* `support_contacts` - (Required) (Updatable) Contact information for getting support from the publisher for the listing.
	* `email` - (Optional) (Updatable) The email of the contact.
	* `name` - (Optional) (Updatable) The name of the contact.
	* `phone` - (Optional) (Updatable) The phone number of the contact.
	* `subject` - (Optional) (Updatable) The email subject line to use when contacting support.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The date and time the publication was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Publication
	* `update` - (Defaults to 20 minutes), when updating the Publication
	* `delete` - (Defaults to 20 minutes), when destroying the Publication


## Import

Publications can be imported using the `id`, e.g.

```
$ terraform import oci_marketplace_publication.test_publication "id"
```

