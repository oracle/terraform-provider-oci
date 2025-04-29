---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_datastore"
sidebar_current: "docs-oci-resource-ocvp-datastore"
description: |-
  Provides the Datastore resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# oci_ocvp_datastore
This resource provides the Datastore resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Creates a Oracle Cloud VMware Solution Datastore.

Use the [WorkRequest](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/WorkRequest/) operations to track the
creation of the Datastore.


## Example Usage

```hcl
resource "oci_ocvp_datastore" "test_datastore" {
	#Required
	availability_domain = var.datastore_availability_domain
	block_volume_ids = var.datastore_block_volume_ids
	compartment_id = var.compartment_id
	display_name = var.datastore_display_name

	#Optional
	datastore_cluster_id = oci_ocvp_datastore_cluster.test_datastore_cluster.id
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain to create the Datastore in. 
* `block_volume_ids` - (Required) The List of Block volume [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s that belong to the Datastore. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the Datastore. 
* `datastore_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the datastore cluster that Datastore belongs to. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A descriptive name for the Datastore. It must be unique within a SDDC, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the Datastore. 
* `block_volume_details` - The list of Block Volume details that belong to the datastore. 
	* `attachments` - List of BlockVolumeAttachment objects containing information about attachment details
		* `esxi_host_id` - The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host that block volume is attached to. 
		* `ip_address` - The IP address of block volume attachment.
		* `port` - The port of block volume attachment.
	* `id` - An [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the Block Storage Volume.
	* `iqn` - An IQN of the Block Storage Volume.
* `block_volume_ids` - The List of Block volume [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s that belong to the Datastore. 
* `capacity_in_gbs` - Total size of the datastore in GB.
* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VMware Cluster that Datastore is attached to. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Datastore. 
* `datastore_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the datastore cluster that Datastore belongs to. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A descriptive name for the Datastore. It must be unique within a SDDC, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Datastore. 
* `sddc_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC that Datastore is associated with. 
* `state` - The current state of the Datastore.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The date and time the Datastore was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the Datastore was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Datastore
	* `update` - (Defaults to 20 minutes), when updating the Datastore
	* `delete` - (Defaults to 20 minutes), when destroying the Datastore


## Import

Datastores can be imported using the `id`, e.g.

```
$ terraform import oci_ocvp_datastore.test_datastore "id"
```

