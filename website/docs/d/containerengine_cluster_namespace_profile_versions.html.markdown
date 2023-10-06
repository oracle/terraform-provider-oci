---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_namespace_profile_versions"
sidebar_current: "docs-oci-datasource-containerengine-cluster_namespace_profile_versions"
description: |-
  Provides the list of Cluster Namespace Profile Versions in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster_namespace_profile_versions
This data source provides the list of Cluster Namespace Profile Versions in Oracle Cloud Infrastructure Container Engine service.

Returns a list of ClusterNamespaceProfileVersions.


## Example Usage

```hcl
data "oci_containerengine_cluster_namespace_profile_versions" "test_cluster_namespace_profile_versions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.cluster_namespace_profile_version_display_name
	id = var.cluster_namespace_profile_version_id
	state = var.cluster_namespace_profile_version_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) unique ClusterNamespaceProfileVersion identifier
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `cluster_namespace_profile_version_collection` - The list of cluster_namespace_profile_version_collection.

### ClusterNamespaceProfileVersion Reference

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

