---
subcategory: "Cluster Placement Groups"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cluster_placement_groups_cluster_placement_groups"
sidebar_current: "docs-oci-datasource-cluster_placement_groups-cluster_placement_groups"
description: |-
  Provides the list of Cluster Placement Groups in Oracle Cloud Infrastructure Cluster Placement Groups service
---

# Data Source: oci_cluster_placement_groups_cluster_placement_groups
This data source provides the list of Cluster Placement Groups in Oracle Cloud Infrastructure Cluster Placement Groups service.

Gets a list of all cluster placement groups in the specified compartment.


## Example Usage

```hcl
data "oci_cluster_placement_groups_cluster_placement_groups" "test_cluster_placement_groups" {

	#Optional
	ad = var.cluster_placement_group_ad
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.cluster_placement_group_compartment_id_in_subtree
	id = var.cluster_placement_group_id
	name = var.cluster_placement_group_name
	state = var.cluster_placement_group_state
}
```

## Argument Reference

The following arguments are supported:

* `ad` - (Optional) A filter to return only the resources that match the specified availability domain.
* `compartment_id` - (Optional) A filter to return only the resources that match the specified compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `compartment_id_in_subtree` - (Optional) When set to `true`, cluster placement groups in all compartments under the specified compartment are returned. The default is set to `false`. 
* `id` - (Optional) A filter to return only the resources that match the specified unique cluster placement group identifier.
* `name` - (Optional) A filter to return only the resources that match the entire display name specified.
* `state` - (Optional) A filter to return only the resources that match the specified lifecycle state.


## Attributes Reference

The following attributes are exported:

* `cluster_placement_group_collection` - The list of cluster_placement_group_collection.

### ClusterPlacementGroup Reference

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

