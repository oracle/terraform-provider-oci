---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_namespace"
sidebar_current: "docs-oci-resource-containerengine-cluster_namespace"
description: |-
  Provides the Cluster Namespace resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_cluster_namespace
This resource provides the Cluster Namespace resource in Oracle Cloud Infrastructure Container Engine service.

Creates a new ClusterNamespace.


## Example Usage

```hcl
resource "oci_containerengine_cluster_namespace" "test_cluster_namespace" {
	#Required
	cluster_namespace_profile_version_id = oci_containerengine_cluster_namespace_profile_version.test_cluster_namespace_profile_version.id
	compartment_id = var.compartment_id
	name = var.cluster_namespace_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.cluster_namespace_description
	freeform_tags = {"Department"= "Finance"}
	namespace_annotations {
		#Required
		key = var.cluster_namespace_namespace_annotations_key
		value = var.cluster_namespace_namespace_annotations_value
	}
	namespace_labels {
		#Required
		key = var.cluster_namespace_namespace_labels_key
		value = var.cluster_namespace_namespace_labels_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `cluster_namespace_profile_version_id` - (Required) (Updatable) OCID of Cluster Namespace Profile Version to use.
* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the resource. It can be changed after creation.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) ClusterNamespace Identifier
* `namespace_annotations` - (Optional) (Updatable) List of Kubernetes annotations to apply to the resulting namespace. 
	* `key` - (Required) (Updatable) Unique annotation key
	* `value` - (Required) (Updatable) Value associated with annotation key
* `namespace_labels` - (Optional) (Updatable) List of Kubernetes labels to apply to the resulting namespace. 
	* `key` - (Required) (Updatable) Unique label key
	* `value` - (Required) (Updatable) Value associated with label key


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cluster_ids` - List of OKE Cluster OCIDs the Cluster Namespace is provisioned upon
* `cluster_namespace_profile_version_id` - OCID of Cluster Namespace Profile Version to use.
* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the resource. It can be changed after creation.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier that is immutable on creation
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `name` - Name of the cluster namespace.
* `namespace` - Name of the resulting Kubernetes namespace
* `namespace_annotations` - List of Kubernetes annotations to apply to the resulting namespace. 
	* `key` - Unique annotation key
	* `value` - Value associated with annotation key
* `namespace_labels` - List of Kubernetes labels to apply to the resulting namespace. 
	* `key` - Unique label key
	* `value` - Value associated with label key
* `state` - The current state of the ClusterNamespace.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster Namespace
	* `update` - (Defaults to 20 minutes), when updating the Cluster Namespace
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster Namespace


## Import

ClusterNamespaces can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_cluster_namespace.test_cluster_namespace "id"
```

