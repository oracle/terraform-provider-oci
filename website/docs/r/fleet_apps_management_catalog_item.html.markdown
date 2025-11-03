---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_catalog_item"
sidebar_current: "docs-oci-resource-fleet_apps_management-catalog_item"
description: |-
  Provides the Catalog Item resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_catalog_item
This resource provides the Catalog Item resource in Oracle Cloud Infrastructure Fleet Apps Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/fleet-management/latest/CatalogItem

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/fleet_apps_management

Creates a CatalogItem.


## Example Usage

```hcl
resource "oci_fleet_apps_management_catalog_item" "test_catalog_item" {
	#Required
	compartment_id = var.compartment_id
	config_source_type = var.catalog_item_config_source_type
	description = var.catalog_item_description
	display_name = var.catalog_item_display_name
	package_type = var.catalog_item_package_type

	#Optional
	catalog_source_payload {
		#Required
		config_source_type = var.catalog_item_catalog_source_payload_config_source_type

		#Optional
		access_uri = var.catalog_item_catalog_source_payload_access_uri
		branch_name = var.catalog_item_catalog_source_payload_branch_name
		bucket = var.catalog_item_catalog_source_payload_bucket
		configuration_source_provider_id = oci_fleet_apps_management_configuration_source_provider.test_configuration_source_provider.id
		description = var.catalog_item_catalog_source_payload_description
		listing_id = oci_marketplace_listing.test_listing.id
		long_description = var.catalog_item_catalog_source_payload_long_description
		namespace = var.catalog_item_catalog_source_payload_namespace
		object = var.catalog_item_catalog_source_payload_object
		repository_url = var.catalog_item_catalog_source_payload_repository_url
		template_display_name = var.catalog_item_catalog_source_payload_template_display_name
		time_expires = var.catalog_item_catalog_source_payload_time_expires
		version = var.catalog_item_catalog_source_payload_version
		working_directory = var.catalog_item_catalog_source_payload_working_directory
		zip_file_base64encoded = var.catalog_item_catalog_source_payload_zip_file_base64encoded
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	is_item_locked = var.catalog_item_is_item_locked
	listing_id = oci_marketplace_listing.test_listing.id
	listing_version = var.catalog_item_listing_version
	short_description = var.catalog_item_short_description
	time_released = var.catalog_item_time_released
	version_description = var.catalog_item_version_description
}
```

## Argument Reference

The following arguments are supported:

* `catalog_source_payload` - (Optional) Catalog source payload.
	* `access_uri` - (Applicable when config_source_type=PAR_CATALOG_SOURCE) access uri 
	* `branch_name` - (Applicable when config_source_type=GIT_CATALOG_SOURCE) branch Name 
	* `bucket` - (Applicable when config_source_type=PAR_CATALOG_SOURCE) bucket name 
	* `config_source_type` - (Required) config source type. 
	* `configuration_source_provider_id` - (Applicable when config_source_type=GIT_CATALOG_SOURCE) configuration Source Provider [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) 
	* `description` - (Applicable when config_source_type=STACK_TEMPLATE_CATALOG_SOURCE) Template Description 
	* `listing_id` - (Applicable when config_source_type=MARKETPLACE_CATALOG_SOURCE) This listing Id parameter of Payload.
	* `long_description` - (Applicable when config_source_type=STACK_TEMPLATE_CATALOG_SOURCE) Template Long Description 
	* `namespace` - (Applicable when config_source_type=PAR_CATALOG_SOURCE) nameSpace 
	* `object` - (Applicable when config_source_type=PAR_CATALOG_SOURCE) object name 
	* `repository_url` - (Applicable when config_source_type=GIT_CATALOG_SOURCE) repository Url 
	* `template_display_name` - (Applicable when config_source_type=STACK_TEMPLATE_CATALOG_SOURCE) Template Display Name 
	* `time_expires` - (Applicable when config_source_type=PAR_CATALOG_SOURCE) The date and time expires, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
	* `version` - (Applicable when config_source_type=MARKETPLACE_CATALOG_SOURCE) This version parameter of Payload.
	* `working_directory` - (Optional) File path to the directory to use for running Terraform. If not specified, the root directory is used. 
	* `zip_file_base64encoded` - (Applicable when config_source_type=STACK_TEMPLATE_CATALOG_SOURCE) The Base64 encoded template. This payload will trigger CreateTemplate API, where the parameter will be passed. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `config_source_type` - (Required) Config source type Eg: STACK_TEMPLATE_CATALOG_SOURCE, PAR_CATALOG_SOURCE, GIT_CATALOG_SOURCE, MARKETPLACE_CATALOG_SOURCE. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Required) (Updatable) The description of the CatalogItem. 
* `display_name` - (Required) (Updatable) The CatalogItem name.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_item_locked` - (Optional) (Updatable) Indicates if the CatalogItem is immutable or not.
* `listing_id` - (Optional) The catalog listing Id. 
* `listing_version` - (Optional) The catalog package version. 
* `package_type` - (Required) Config package type Eg: TF_PACKAGE, NON_TF_PACKAGE, CONFIG_FILE. 
* `short_description` - (Optional) (Updatable) Short description about the catalog item.
* `time_released` - (Optional) The date and time the CatalogItem was released, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `version_description` - (Optional) (Updatable) Version description about the catalog item.
* `clone_catalog_item_trigger` - (Optional) (Updatable) An optional property when incremented triggers Clone Catalog Item. Could be set to any integer value.
* `configure_trigger` - (Optional) (Updatable) An optional property when incremented triggers Configure. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Catalog Item
	* `update` - (Defaults to 20 minutes), when updating the Catalog Item
	* `delete` - (Defaults to 20 minutes), when destroying the Catalog Item


## Import

CatalogItems can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_catalog_item.test_catalog_item "id"
```

