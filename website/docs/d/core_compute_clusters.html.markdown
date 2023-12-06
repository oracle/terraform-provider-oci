---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_clusters"
sidebar_current: "docs-oci-datasource-core-compute_clusters"
description: |-
  Provides the list of Compute Clusters in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_clusters
This data source provides the list of Compute Clusters in Oracle Cloud Infrastructure Core service.

Lists the compute clusters in the specified compartment.
A [compute cluster](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/compute-clusters.htm) is a remote direct memory access (RDMA) network group.


## Example Usage

```hcl
data "oci_core_compute_clusters" "test_compute_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.compute_cluster_availability_domain
	display_name = var.compute_cluster_display_name
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 


## Attributes Reference

The following attributes are exported:

* `compute_cluster_collection` - The list of compute_cluster_collection.

### ComputeCluster Reference

The following attributes are exported:

* `availability_domain` - The availability domain the compute cluster is running in.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the compute cluster.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute cluster.
* `state` - The current state of the compute cluster.
* `time_created` - The date and time the compute cluster was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

