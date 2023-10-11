---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_workload_mapping"
sidebar_current: "docs-oci-datasource-containerengine-cluster_workload_mapping"
description: |-
  Provides details about a specific Cluster Workload Mapping in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster_workload_mapping
This data source provides details about a specific Cluster Workload Mapping resource in Oracle Cloud Infrastructure Container Engine service.

Get the specified workloadMapping for a cluster.

## Example Usage

```hcl
data "oci_containerengine_cluster_workload_mapping" "test_cluster_workload_mapping" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
	workload_mapping_id = oci_containerengine_workload_mapping.test_workload_mapping.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster.
* `workload_mapping_id` - (Required) The OCID of the workloadMapping.


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

