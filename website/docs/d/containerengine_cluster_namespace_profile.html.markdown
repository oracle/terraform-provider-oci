---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_namespace_profile"
sidebar_current: "docs-oci-datasource-containerengine-cluster_namespace_profile"
description: |-
  Provides details about a specific Cluster Namespace Profile in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster_namespace_profile
This data source provides details about a specific Cluster Namespace Profile resource in Oracle Cloud Infrastructure Container Engine service.

Gets a ClusterNamespaceProfile by identifier

## Example Usage

```hcl
data "oci_containerengine_cluster_namespace_profile" "test_cluster_namespace_profile" {
	#Required
	cluster_namespace_profile_id = oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_namespace_profile_id` - (Required) unique ClusterNamespaceProfile identifier


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

