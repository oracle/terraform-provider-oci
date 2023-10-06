---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_namespace_profile_version"
sidebar_current: "docs-oci-resource-containerengine-cluster_namespace_profile_version"
description: |-
  Provides the Cluster Namespace Profile Version resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_cluster_namespace_profile_version
This resource provides the Cluster Namespace Profile Version resource in Oracle Cloud Infrastructure Container Engine service.

Creates a new ClusterNamespaceProfileVersion.


## Example Usage

```hcl
resource "oci_containerengine_cluster_namespace_profile_version" "test_cluster_namespace_profile_version" {
	#Required
	admin_cluster_role_name = var.cluster_namespace_profile_version_admin_cluster_role_name
	cluster_namespace_profile_id = oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile.id
	compartment_id = var.compartment_id
	name = var.cluster_namespace_profile_version_name

	#Optional
	allowed_namespace_annotations {
		#Required
		key = var.cluster_namespace_profile_version_allowed_namespace_annotations_key

		#Optional
		value = var.cluster_namespace_profile_version_allowed_namespace_annotations_value
	}
	allowed_namespace_labels {
		#Required
		key = var.cluster_namespace_profile_version_allowed_namespace_labels_key

		#Optional
		value = var.cluster_namespace_profile_version_allowed_namespace_labels_value
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.cluster_namespace_profile_version_description
	fixed_namespace_annotations {
		#Required
		key = var.cluster_namespace_profile_version_fixed_namespace_annotations_key
		value = var.cluster_namespace_profile_version_fixed_namespace_annotations_value
	}
	fixed_namespace_labels {
		#Required
		key = var.cluster_namespace_profile_version_fixed_namespace_labels_key
		value = var.cluster_namespace_profile_version_fixed_namespace_labels_value
	}
	freeform_tags = {"Department"= "Finance"}
	is_deprecated = var.cluster_namespace_profile_version_is_deprecated
	required_namespace_annotations {
		#Required
		key = var.cluster_namespace_profile_version_required_namespace_annotations_key

		#Optional
		value = var.cluster_namespace_profile_version_required_namespace_annotations_value
	}
	required_namespace_labels {
		#Required
		key = var.cluster_namespace_profile_version_required_namespace_labels_key

		#Optional
		value = var.cluster_namespace_profile_version_required_namespace_labels_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `admin_cluster_role_name` - (Required) Name of the ClusterRole to bind to the admin account in the resulting namespace.
* `allowed_namespace_annotations` - (Optional) List of Kubernetes annotations that may be specified via Cluster Namespaces.
	* `key` - (Required) Unique annotation key
	* `value` - (Optional) Allowed set of values associated with annotation key, empty array means any value is allowed
* `allowed_namespace_labels` - (Optional) List of Kubernetes labels that may be specified via Cluster Namespaces.
	* `key` - (Required) Allowed unique label key
	* `value` - (Optional) Allowed set of values associated with label key, empty array means any value is allowed
* `cluster_namespace_profile_id` - (Required) The OCID of the ClusterNamespaceProfile
* `compartment_id` - (Required) (Updatable) OCID of compartment owning the Cluster Namespace Profile Version.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the resource. It can be changed after creation.
* `fixed_namespace_annotations` - (Optional) List of Kubernetes annotations to apply to the resulting namespace.
	* `key` - (Required) Unique annotation key
	* `value` - (Required) Value associated with annotation key
* `fixed_namespace_labels` - (Optional) List of Kubernetes labels to apply to the resulting namespace.
	* `key` - (Required) Unique label key
	* `value` - (Required) Value associated with label key
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_deprecated` - (Optional) (Updatable) If set to true, the Cluster Namespace Profile Version is not consumable by new Cluster Namespace configurations.
* `name` - (Required) A name for the Cluster Namespace Profile Version. Names are unique across versions in a Cluster Namespace Profile Profiles.
* `required_namespace_annotations` - (Optional) List of Kubernetes annotations that must be specified via Cluster Namespaces.
	* `key` - (Required) Unique annotation key
	* `value` - (Optional) Allowed set of values associated with annotation key, empty array means any value is allowed
* `required_namespace_labels` - (Optional) List of Kubernetes labels that must be specified via Cluster Namespaces.
	* `key` - (Required) Required unique label key
	* `value` - (Optional) Allowed set of values associated with label key, empty array means any value is allowed


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `admin_cluster_role_name` - Name of the ClusterRole to bind to the admin account in the resulting namespace.
* `allowed_namespace_annotations` - List of Kubernetes annotations that may be specified via Cluster Namespaces.
	* `key` - Unique annotation key
	* `value` - Allowed set of values associated with annotation key, empty array means any value is allowed
* `allowed_namespace_labels` - List of Kubernetes labels that may be specified via Cluster Namespaces.
	* `key` - Allowed unique label key
	* `value` - Allowed set of values associated with label key, empty array means any value is allowed
* `cluster_namespace_profile_id` - The OCID of the ClusterNamespaceProfile
* `compartment_id` - OCID of compartment owning the Cluster Namespace Profile Version.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the resource. It can be changed after creation.
* `fixed_namespace_annotations` - List of Kubernetes annotations to apply to the resulting namespace.
	* `key` - Unique annotation key
	* `value` - Value associated with annotation key
* `fixed_namespace_labels` - List of Kubernetes labels to apply to the resulting namespace.
	* `key` - Unique label key
	* `value` - Value associated with label key
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier that is immutable on creation.
* `is_deprecated` - If set to true, the Cluster Namespace Profile Version is not consumable by new Cluster Namespace configurations.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `name` - A name for the Cluster Namespace Profile Version. Names are unique across versions in a Cluster Namespace Profile Profiles.
* `required_namespace_annotations` - List of Kubernetes annotations that must be specified via Cluster Namespaces.
	* `key` - Unique annotation key
	* `value` - Allowed set of values associated with annotation key, empty array means any value is allowed
* `required_namespace_labels` - List of Kubernetes labels that must be specified via Cluster Namespaces.
	* `key` - Required unique label key
	* `value` - Allowed set of values associated with label key, empty array means any value is allowed
* `state` - The current state of the resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster Namespace Profile Version
	* `update` - (Defaults to 20 minutes), when updating the Cluster Namespace Profile Version
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster Namespace Profile Version


## Import

ClusterNamespaceProfileVersions can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_cluster_namespace_profile_version.test_cluster_namespace_profile_version "id"
```

