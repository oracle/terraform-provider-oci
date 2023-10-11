---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_workload_mapping"
sidebar_current: "docs-oci-resource-containerengine-cluster_workload_mapping"
description: |-
  Provides the Cluster Workload Mapping resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_cluster_workload_mapping
This resource provides the Cluster Workload Mapping resource in Oracle Cloud Infrastructure Container Engine service.

Create the specified workloadMapping for a cluster.

## Example Usage

```hcl
resource "oci_containerengine_cluster_workload_mapping" "test_cluster_workload_mapping" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
	mapped_compartment_id = oci_identity_compartment.test_compartment.id
	namespace = var.cluster_workload_mapping_namespace

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `mapped_compartment_id` - (Required) (Updatable) The OCID of the mapped customer compartment.
* `namespace` - (Required) The namespace of the workloadMapping.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cluster_id` - The OCID of the cluster.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The ocid of the workloadMapping.
* `mapped_compartment_id` - The OCID of the mapped customer compartment.
* `mapped_tenancy_id` - The OCID of the mapped customer tenancy.
* `namespace` - The namespace of the workloadMapping.
* `state` - The state of the workloadMapping.
* `time_created` - The time the cluster was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster Workload Mapping
	* `update` - (Defaults to 20 minutes), when updating the Cluster Workload Mapping
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster Workload Mapping


## Import

ClusterWorkloadMappings can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_cluster_workload_mapping.test_cluster_workload_mapping "clusters/{clusterId}/workloadMappings/{workloadMappingId}" 
```

