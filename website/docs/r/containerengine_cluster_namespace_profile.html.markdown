---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_namespace_profile"
sidebar_current: "docs-oci-resource-containerengine-cluster_namespace_profile"
description: |-
  Provides the Cluster Namespace Profile resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_cluster_namespace_profile
This resource provides the Cluster Namespace Profile resource in Oracle Cloud Infrastructure Container Engine service.

Creates a new ClusterNamespaceProfile.


## Example Usage

```hcl
resource "oci_containerengine_cluster_namespace_profile" "test_cluster_namespace_profile" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.cluster_namespace_profile_display_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.cluster_namespace_profile_description
	freeform_tags = {"Department"= "Finance"}
	namespace_suffix = var.cluster_namespace_profile_namespace_suffix
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) OCID of compartment owning the cluster namespace.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the resource. It can be changed after creation.
* `display_name` - (Required) (Updatable) Name of the cluster namespace.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `namespace_suffix` - (Optional) Suffix to append to the end of the namespaces generated from this Profile


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - OCID of compartment owning the cluster namespace.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the resource. It can be changed after creation.
* `display_name` - Name of the cluster namespace.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier that is immutable on creation.
* `kubernetes_version` - Minimum Kubernetes version supported by the Cluster Namespace Profile. Effectively the minimum version of Kubernetes clusters attached to the Profile. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `namespace_suffix` - Suffix to append to the end of the namespaces generated from this Profile
* `state` - The current state of the ClusterNamespaceProfile.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster Namespace Profile
	* `update` - (Defaults to 20 minutes), when updating the Cluster Namespace Profile
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster Namespace Profile


## Import

ClusterNamespaceProfiles can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile "id"
```

