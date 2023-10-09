---
subcategory: "Jms Java Downloads"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_java_downloads_java_download_tokens"
sidebar_current: "docs-oci-datasource-jms_java_downloads-java_download_tokens"
description: |-
  Provides the list of Java Download Tokens in Oracle Cloud Infrastructure Jms Java Downloads service
---

# Data Source: oci_jms_java_downloads_java_download_tokens
This data source provides the list of Java Download Tokens in Oracle Cloud Infrastructure Jms Java Downloads service.

Returns a list of JavaDownloadTokens.


## Example Usage

```hcl
data "oci_jms_java_downloads_java_download_tokens" "test_java_download_tokens" {
	#Required
	compartment_id = var.tenancy_ocid

	#Optional
	display_name = var.java_download_token_display_name
	family_version = var.java_download_token_family_version
	id = var.java_download_token_id
	search_by_user = var.java_download_token_search_by_user
	state = var.java_download_token_state
	value = var.java_download_token_value
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `family_version` - (Optional) Unique Java family version identifier.
* `id` - (Optional) Unique JavaDownloadToken identifier.
* `search_by_user` - (Optional) A filter to return only resources that match the user principal detail.  The search string can be any of the property values from the [Principal](https://docs.cloud.oracle.com/iaas/api/#/en/jms/latest/datatypes/Principal) object. This object is used as response datatype for the `createdBy` and `lastUpdatedBy` fields in applicable resource. 
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.
* `value` - (Optional) Unique JavaDownloadToken value.


## Attributes Reference

The following attributes are exported:

* `java_download_token_collection` - The list of java_download_token_collection.

### JavaDownloadToken Reference

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

