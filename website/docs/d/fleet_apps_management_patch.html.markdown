---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_patch"
sidebar_current: "docs-oci-datasource-fleet_apps_management-patch"
description: |-
  Provides details about a specific Patch in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_patch
This data source provides details about a specific Patch resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a Patch by identifier

## Example Usage

```hcl
data "oci_fleet_apps_management_patch" "test_patch" {
	#Required
	patch_id = oci_fleet_apps_management_patch.test_patch.id
}
```

## Argument Reference

The following arguments are supported:

* `patch_id` - (Required) unique Patch identifier


## Attributes Reference

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

