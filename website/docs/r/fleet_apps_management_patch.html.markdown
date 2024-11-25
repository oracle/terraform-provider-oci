---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_patch"
sidebar_current: "docs-oci-resource-fleet_apps_management-patch"
description: |-
  Provides the Patch resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_patch
This resource provides the Patch resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Creates a new Patch.


## Example Usage

```hcl
resource "oci_fleet_apps_management_patch" "test_patch" {
	#Required
	artifact_details {
		#Required
		category = var.patch_artifact_details_category

		#Optional
		artifact {

			#Optional
			content {
				#Required
				bucket = var.patch_artifact_details_artifact_content_bucket
				checksum = var.patch_artifact_details_artifact_content_checksum
				namespace = var.patch_artifact_details_artifact_content_namespace
				object = var.patch_artifact_details_artifact_content_object
				source_type = var.patch_artifact_details_artifact_content_source_type
			}
		}
		artifacts {

			#Optional
			architecture = var.patch_artifact_details_artifacts_architecture
			content {
				#Required
				bucket = var.patch_artifact_details_artifacts_content_bucket
				checksum = var.patch_artifact_details_artifacts_content_checksum
				namespace = var.patch_artifact_details_artifacts_content_namespace
				object = var.patch_artifact_details_artifacts_content_object
				source_type = var.patch_artifact_details_artifacts_content_source_type
			}
			os_type = var.patch_artifact_details_artifacts_os_type
		}
	}
	compartment_id = var.compartment_id
	name = var.patch_name
	patch_type {
		#Required
		platform_configuration_id = oci_fleet_apps_management_platform_configuration.test_platform_configuration.id
	}
	product {
		#Required
		platform_configuration_id = oci_fleet_apps_management_platform_configuration.test_platform_configuration.id

		#Optional
		version = var.patch_product_version
	}
	severity = var.patch_severity
	time_released = var.patch_time_released

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	dependent_patches {
		#Required
		id = var.patch_dependent_patches_id
	}
	description = var.patch_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `artifact_details` - (Required) (Updatable) Patch artifact description and content details.
	* `artifact` - (Required when category=GENERIC) (Updatable) Patch artifact metadata Details which is common for all platforms. 
		* `content` - (Required when category=GENERIC) (Updatable) Content Source details.
			* `bucket` - (Required) (Updatable) Bucket Name.
			* `checksum` - (Required) (Updatable) md5 checksum of the artifact.
			* `namespace` - (Required) (Updatable) Namespace.
			* `object` - (Required) (Updatable) Object Name.
			* `source_type` - (Required) (Updatable) Content Source type details. 
	* `artifacts` - (Required when category=PLATFORM_SPECIFIC) (Updatable) Artifacts. 
		* `architecture` - (Required when category=PLATFORM_SPECIFIC) (Updatable) System architecture.
		* `content` - (Required when category=PLATFORM_SPECIFIC) (Updatable) Content Source details.
			* `bucket` - (Required) (Updatable) Bucket Name.
			* `checksum` - (Required) (Updatable) md5 checksum of the artifact.
			* `namespace` - (Required) (Updatable) Namespace.
			* `object` - (Required) (Updatable) Object Name.
			* `source_type` - (Required) (Updatable) Content Source type details. 
		* `os_type` - (Required when category=PLATFORM_SPECIFIC) (Updatable) The OS type the patch is applicable for.
	* `category` - (Required) (Updatable) Artifact category details.
* `compartment_id` - (Required) (Updatable) 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `dependent_patches` - (Optional) (Updatable) Dependent Patches for this patch. 
	* `id` - (Required) (Updatable) The OCID of the resource.
* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `name` - (Required) A user-friendly name. Should be unique within the tenancy, and cannot be changed after creation.  Avoid entering confidential information. 
* `patch_type` - (Required) (Updatable) Patch Type
	* `platform_configuration_id` - (Required) (Updatable) PlatformConfiguration Id corresponding to the Patch Type
* `product` - (Required) (Updatable) Product
	* `platform_configuration_id` - (Required) (Updatable) PlatformConfiguration Id corresponding to the Product
	* `version` - (Optional) (Updatable) product version.
* `severity` - (Required) (Updatable) Patch Severity.
* `time_released` - (Required) (Updatable) Date when the patch was released.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Patch
	* `update` - (Defaults to 20 minutes), when updating the Patch
	* `delete` - (Defaults to 20 minutes), when destroying the Patch


## Import

Patches can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_patch.test_patch "id"
```

