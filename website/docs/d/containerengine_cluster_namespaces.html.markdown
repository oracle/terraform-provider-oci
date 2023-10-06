---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_namespaces"
sidebar_current: "docs-oci-datasource-containerengine-cluster_namespaces"
description: |-
  Provides the list of Cluster Namespaces in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster_namespaces
This data source provides the list of Cluster Namespaces in Oracle Cloud Infrastructure Container Engine service.

Returns a list of ClusterNamespaces.


## Example Usage

```hcl
data "oci_containerengine_cluster_namespaces" "test_cluster_namespaces" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	id = var.cluster_namespace_id
	name = var.cluster_namespace_name
	state = var.cluster_namespace_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `id` - (Optional) unique ClusterNamespace identifier
* `name` - (Optional) The name to filter on.
* `state` - (Optional) A filter to return only resources with lifecycleState matching the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `cluster_namespace_collection` - The list of cluster_namespace_collection.

### ClusterNamespace Reference

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

