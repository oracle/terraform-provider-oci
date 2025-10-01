---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_datastores"
sidebar_current: "docs-oci-datasource-ocvp-datastores"
description: |-
  Provides the list of Datastores in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_datastores
This data source provides the list of Datastores in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

List the Datastores in the specified compartment. The list can be filtered
by compartment, datastore id, display name and lifecycle state.


## Example Usage

```hcl
data "oci_ocvp_datastores" "test_datastores" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	cluster_id = oci_ocvp_cluster.test_cluster.id
	datastore_id = oci_ocvp_datastore.test_datastore.id
	display_name = var.datastore_display_name
	state = var.datastore_state
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC Cluster. 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `datastore_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Datastore. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `state` - (Optional) The lifecycle state of the resource.


## Attributes Reference

The following attributes are exported:

* `datastore_collection` - The list of datastore_collection.

### Datastore Reference

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

