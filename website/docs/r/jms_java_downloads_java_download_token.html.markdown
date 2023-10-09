---
subcategory: "Jms Java Downloads"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_downloads_java_download_token"
sidebar_current: "docs-oci-resource-jms_java_downloads-java_download_token"
description: |-
  Provides the Java Download Token resource in Oracle Cloud Infrastructure Jms Java Downloads service
---

# oci_jms_java_downloads_java_download_token
This resource provides the Java Download Token resource in Oracle Cloud Infrastructure Jms Java Downloads service.

Creates a new JavaDownloadToken in the tenancy with specified attributes.


## Example Usage

```hcl
resource "oci_jms_java_downloads_java_download_token" "test_java_download_token" {
	#Required
	compartment_id = var.tenancy_ocid
	description = var.java_download_token_description
	display_name = var.java_download_token_display_name
	java_version = var.java_download_token_java_version
	license_type = var.java_download_token_license_type
	time_expires = var.java_download_token_time_expires

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	is_default = var.java_download_token_is_default
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy scoped to the JavaDownloadToken.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)). 
* `description` - (Required) (Updatable) User provided description of the JavaDownloadToken.
* `display_name` - (Required) (Updatable) User provided display name of the JavaDownloadToken.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).) 
* `is_default` - (Optional) (Updatable) The token default attribute.
* `java_version` - (Required) The Java version associated with the token.
* `license_type` - (Required) (Updatable) The license type(s) associated with the JavaDownloadToken.
* `time_expires` - (Required) (Updatable) Expiry time of the token.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy scoped to the JavaDownloadToken. 
* `created_by` - An authorized principal.
	* `display_name` - The name of the principal.
	* `email` - The email of the principal.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)). 
* `description` - User provided description of the JavaDownloadToken.
* `display_name` - User provided display name of the JavaDownloadToken.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).) 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the JavaDownloadToken. 
* `is_default` - A flag to indicate if the token is default.
* `java_version` - The associated Java version of the JavaDownloadToken.
* `last_updated_by` - An authorized principal.
	* `display_name` - The name of the principal.
	* `email` - The email of the principal.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal.
* `license_type` - The license type(s) associated with the JavaDownloadToken.
* `lifecycle_details` - Possible lifecycle substates.
* `state` - The current state of the JavaDownloadToken.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the JavaDownloadToken was created. An RFC3339 formatted datetime string.
* `time_expires` - The expiry time of the JavaDownloadToken. An RFC3339 formatted datetime string.
* `time_last_used` - The time the JavaDownloadToken was last used for download. An RFC3339 formatted datetime string.
* `time_updated` - The time the JavaDownloadToken was updated. An RFC3339 formatted datetime string.
* `value` - Uniquely generated value for the JavaDownloadToken.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Java Download Token
	* `update` - (Defaults to 20 minutes), when updating the Java Download Token
	* `delete` - (Defaults to 20 minutes), when destroying the Java Download Token


## Import

Import is not supported for this resource.

