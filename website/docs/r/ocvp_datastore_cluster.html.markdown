---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_datastore_cluster"
sidebar_current: "docs-oci-resource-ocvp-datastore_cluster"
description: |-
  Provides the Datastore Cluster resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# oci_ocvp_datastore_cluster
This resource provides the Datastore Cluster resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/vmware/latest/DatastoreCluster

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/ocvp

Creates a Oracle Cloud VMware Solution Datastore Cluster.


## Example Usage

```hcl
resource "oci_ocvp_datastore_cluster" "test_datastore_cluster" {
	#Required
	availability_domain = var.datastore_cluster_availability_domain
	compartment_id = var.compartment_id
	datastore_cluster_type = var.datastore_cluster_datastore_cluster_type
	display_name = var.datastore_cluster_display_name

	#Optional
	datastore_ids = var.datastore_cluster_datastore_ids
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain to create the Datastore Cluster in. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the Datastore Cluster. 
* `datastore_cluster_type` - (Required) Type of the datastore.
* `datastore_ids` - (Optional) The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Datastores that belong to the Datastore Cluster. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A descriptive name for the Datastore Cluster. It must be unique within a SDDC, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Datastore Cluster
	* `update` - (Defaults to 20 minutes), when updating the Datastore Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Datastore Cluster


## Import

DatastoreClusters can be imported using the `id`, e.g.

```
$ terraform import oci_ocvp_datastore_cluster.test_datastore_cluster "id"
```

