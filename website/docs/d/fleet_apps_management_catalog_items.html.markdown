---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_catalog_items"
sidebar_current: "docs-oci-datasource-fleet_apps_management-catalog_items"
description: |-
  Provides the list of Catalog Items in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_catalog_items
This data source provides the list of Catalog Items in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a list of Catalog Items in a compartment.


## Example Usage

```hcl
data "oci_fleet_apps_management_catalog_items" "test_catalog_items" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	catalog_listing_id = oci_marketplace_listing.test_listing.id
	catalog_listing_version_criteria = var.catalog_item_catalog_listing_version_criteria
	config_source_type = var.catalog_item_config_source_type
	display_name = var.catalog_item_display_name
	package_type = var.catalog_item_package_type
	should_list_public_items = var.catalog_item_should_list_public_items
	state = var.catalog_item_state
}
```

## Argument Reference

The following arguments are supported:

* `catalog_listing_id` - (Optional) catalogListingId of the package. This is an integer whose min and max length are specified.
* `catalog_listing_version_criteria` - (Optional) Parameter to list all catalog items only with latest version or list all catalog items with all versions.
* `compartment_id` - (Required) (Updatable) The ID of the compartment in which to list resources.
* `config_source_type` - (Optional) The [ConfigSourceType](/definitions/CatalogItem/configSourceType) Eg: STACK_TEMPLATE_CATALOG_SOURCE, PAR_CATALOG_SOURCE, URL_CATALOG_SOURCE, MARKETPLACE_CATALOG_SOURCE. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `package_type` - (Optional) A filter to return only resources that match the given package type. The state value is case-insensitive.       
* `should_list_public_items` - (Optional) The indicator to append Public Items from the root compartment to any query, when set to TRUE.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `catalog_item_collection` - The list of catalog_item_collection.

### CatalogItem Reference

The following attributes are exported:

* `catalog_result_payload` - Catalog result payload. 
	* `branch_name` - branch Name 
	* `config_result_type` - config result type. 
	* `configuration_source_provider_id` - configuration Source Provider [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) 
	* `package_url` - package url 
	* `repository_url` - repository Url 
	* `template_id` - template id 
	* `time_expires` - The date and time expires, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
	* `working_directory` - working directory 
* `catalog_source_payload` - Catalog source payload.
	* `access_uri` - access uri 
	* `branch_name` - branch Name 
	* `bucket` - bucket name 
	* `config_source_type` - config source type. 
	* `configuration_source_provider_id` - configuration Source Provider [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) 
	* `description` - Template Description 
	* `listing_id` - This listing Id parameter of Payload.
	* `long_description` - Template Long Description 
	* `namespace` - nameSpace 
	* `object` - object name 
	* `repository_url` - repository Url 
	* `template_display_name` - Template Display Name 
	* `time_expires` - The date and time expires, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
	* `version` - This version parameter of Payload.
	* `working_directory` - File path to the directory to use for running Terraform. If not specified, the root directory is used. 
	* `zip_file_base64encoded` - The Base64 encoded template. This payload will trigger CreateTemplate API, where the parameter will be passed. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `config_source_type` - Config source type Eg: STACK_TEMPLATE_CATALOG_SOURCE, PAR_CATALOG_SOURCE, GIT_CATALOG_SOURCE, MARKETPLACE_CATALOG_SOURCE. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description about the catalog item.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the catalog.
* `is_item_locked` - Indicates if the CatalogItem is immutable or not.
* `lifecycle_details` - The details of lifecycle state CatalogItem.
* `listing_id` - The catalog listing Id. 
* `listing_version` - The catalog package version. 
* `package_type` - Config package type Eg: TF_PACKAGE, NON_TF_PACKAGE, CONFIG_FILE. 
* `short_description` - Short description about the catalog item.
* `should_list_public_items` - The indicator to append Public Items from the root compartment to any query, when set to TRUE.
* `state` - The current state of the CatalogItem.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_backfill_last_checked` - The date and time the CatalogItem was last checked by backfill job, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_created` - The date and time the CatalogItem was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_last_checked` - The date and time the CatalogItem was last checked, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_released` - The date and time the CatalogItem was released, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the CatalogItem was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `version_description` - Version description about the catalog item.

