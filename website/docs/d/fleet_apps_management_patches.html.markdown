---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_patches"
sidebar_current: "docs-oci-datasource-fleet_apps_management-patches"
description: |-
  Provides the list of Patches in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_patches
This data source provides the list of Patches in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of Patches.


## Example Usage

```hcl
data "oci_fleet_apps_management_patches" "test_patches" {

	#Optional
	compartment_id = var.compartment_id
	id = var.patch_id
	name = var.patch_name
	patch_type_id = oci_fleet_apps_management_patch_type.test_patch_type.id
	product_id = oci_fleet_apps_management_product.test_product.id
	should_compliance_policy_rules_be_applied = var.patch_should_compliance_policy_rules_be_applied
	state = var.patch_state
	time_released_greater_than_or_equal_to = var.patch_time_released_greater_than_or_equal_to
	time_released_less_than = var.patch_time_released_less_than
	type = var.patch_type
	version = var.patch_version
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `id` - (Optional) unique Patch identifier
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `patch_type_id` - (Optional) Patch Type platformConfigurationId associated with the Patch.
* `product_id` - (Optional) Product platformConfigurationId associated with the Patch.
* `should_compliance_policy_rules_be_applied` - (Optional) Filter patch based on compliance policy rules for the Product
* `state` - (Optional) The current state of the Patch.
* `time_released_greater_than_or_equal_to` - (Optional) Patch Released Date
* `time_released_less_than` - (Optional) Patch Released Date
* `type` - (Optional) DefinedBy type.
* `version` - (Optional) Product version


## Attributes Reference

The following attributes are exported:

* `patch_collection` - The list of patch_collection.

### Patch Reference

The following attributes are exported:

* `artifact_details` - Patch artifact description and content details.
	* `artifact` - Patch artifact metadata Details which is common for all platforms. 
		* `content` - Content Source details.
			* `bucket` - Bucket Name.
			* `checksum` - md5 checksum of the artifact.
			* `namespace` - Namespace.
			* `object` - Object Name.
			* `source_type` - Content Source type details. 
	* `artifacts` - Artifacts. 
		* `architecture` - System architecture.
		* `content` - Content Source details.
			* `bucket` - Bucket Name.
			* `checksum` - md5 checksum of the artifact.
			* `namespace` - Namespace.
			* `object` - Object Name.
			* `source_type` - Content Source type details. 
		* `os_type` - The OS type the patch is applicable for.
	* `category` - Artifact category details.
* `compartment_id` - 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `dependent_patches` - Dependent Patches for this patch. 
	* `id` - The OCID of the resource.
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `name` - A user-friendly name. Should be unique within the tenancy, and cannot be changed after creation.  Avoid entering confidential information. 
* `patch_type` - Patch Type
	* `platform_configuration_id` - PlatformConfiguration Id corresponding to the Patch Type
* `product` - Product
	* `platform_configuration_id` - PlatformConfiguration Id corresponding to the Product
	* `version` - product version.
* `resource_region` - Associated region
* `severity` - Patch Severity.
* `state` - The current state of the Patch.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_released` - Date when the patch was released.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `type` - Provide information on who defined the patch. Example: For Custom Patches the value will be USER_DEFINED For Oracle Defined Patches the value will be ORACLE_DEFINED 

