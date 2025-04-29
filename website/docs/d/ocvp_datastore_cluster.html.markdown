---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_datastore_cluster"
sidebar_current: "docs-oci-datasource-ocvp-datastore_cluster"
description: |-
  Provides details about a specific Datastore Cluster in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_datastore_cluster
This data source provides details about a specific Datastore Cluster resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Get the specified Datastore Cluster information.

## Example Usage

```hcl
data "oci_ocvp_datastore_cluster" "test_datastore_cluster" {
	#Required
	datastore_cluster_id = oci_ocvp_datastore_cluster.test_datastore_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `datastore_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Datastore Cluster. 


## Attributes Reference

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

