---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_workload_mappings"
sidebar_current: "docs-oci-datasource-containerengine-cluster_workload_mappings"
description: |-
  Provides the list of Cluster Workload Mappings in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster_workload_mappings
This data source provides the list of Cluster Workload Mappings in Oracle Cloud Infrastructure Container Engine service.

List workloadMappings for a provisioned cluster.

## Example Usage

```hcl
data "oci_containerengine_cluster_workload_mappings" "test_cluster_workload_mappings" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster.


## Attributes Reference

The following attributes are exported:

* `workload_mappings` - The list of workload_mappings.

### ClusterWorkloadMapping Reference

The following attributes are exported:

* `cluster_id` - The OCID of the cluster.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The ocid of the workloadMapping.
* `mapped_compartment_id` - The OCID of the mapped customer compartment.
* `mapped_tenancy_id` - The OCID of the mapped customer tenancy.
* `namespace` - The namespace of the workloadMapping.
* `state` - The state of the workloadMapping.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the cluster was created.

