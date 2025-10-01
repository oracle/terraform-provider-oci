---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_datastore_clusters"
sidebar_current: "docs-oci-datasource-ocvp-datastore_clusters"
description: |-
  Provides the list of Datastore Clusters in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_datastore_clusters
This data source provides the list of Datastore Clusters in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

List the Datastore Clusters in the specified compartment. The list can be filtered
by compartment, Datastore Cluster, Display name and Lifecycle state


## Example Usage

```hcl
data "oci_ocvp_datastore_clusters" "test_datastore_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	cluster_id = oci_ocvp_cluster.test_cluster.id
	datastore_cluster_id = oci_ocvp_datastore_cluster.test_datastore_cluster.id
	display_name = var.datastore_cluster_display_name
	state = var.datastore_cluster_state
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC Cluster. 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `datastore_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Datastore Cluster. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `state` - (Optional) The lifecycle state of the resource.


## Attributes Reference

The following attributes are exported:

* `datastore_cluster_collection` - The list of datastore_cluster_collection.

### DatastoreCluster Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the Datastore Cluster. 
* `capacity_in_gbs` - Total size of all datastores associated with the datastore cluster in GB.
* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VMware Cluster that Datastore cluster is attached to. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Datastore. 
* `datastore_cluster_type` - Type of the datastore cluster.
* `datastore_ids` - The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Datastores that belong to the Datastore Cluster 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A descriptive name for the Datastore Cluster. It must be unique within a SDDC, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information. 
* `esxi_host_ids` - The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi hosts to attach the datastore to. All ESXi hosts must belong to the same VMware cluster. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Datastore cluster. 
* `sddc_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC that Datastore cluster is associated with. 
* `state` - The current state of the Datastore Cluster.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The date and time the Datastore Cluster was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the Datastore Cluster was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

