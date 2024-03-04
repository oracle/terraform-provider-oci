---
subcategory: "Cluster Placement Groups"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cluster_placement_groups_cluster_placement_group"
sidebar_current: "docs-oci-datasource-cluster_placement_groups-cluster_placement_group"
description: |-
  Provides details about a specific Cluster Placement Group in Oracle Cloud Infrastructure Cluster Placement Groups service
---

# Data Source: oci_cluster_placement_groups_cluster_placement_group
This data source provides details about a specific Cluster Placement Group resource in Oracle Cloud Infrastructure Cluster Placement Groups service.

Gets the specified cluster placement group.

## Example Usage

```hcl
data "oci_cluster_placement_groups_cluster_placement_group" "test_cluster_placement_group" {
	#Required
	cluster_placement_group_id = oci_cluster_placement_groups_cluster_placement_group.test_cluster_placement_group.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_placement_group_id` - (Required) A unique cluster placement group identifier.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the cluster placement group.
* `capabilities` - A list of resources that you can create in a cluster placement group. 
	* `items` - The supported resources.
		* `name` - The type of resource.
		* `service` - The service that the resource is part of.
* `cluster_placement_group_type` - The type of cluster placement group.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the cluster placement group. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description of the cluster placement group.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group. 
* `lifecycle_details` - A message describing the current state in more detail. For example, lifecycle details for a resource in a Failed state might include information to act on. 
* `name` - The user-friendly name of the cluster placement group. The display name for a cluster placement must be unique and you cannot change it. Avoid entering confidential information. 
* `placement_instruction` - Details that inform cluster placement group provisioning.
	* `type` - The type of placement instruction.
	* `value` - The value of the token designated for placement of the cluster placement group upon creation.
* `state` - The current state of the ClusterPlacementGroup.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the cluster placement group was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.
* `time_updated` - The time the cluster placement group was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.

