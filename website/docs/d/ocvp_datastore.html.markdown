---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_datastore"
sidebar_current: "docs-oci-datasource-ocvp-datastore"
description: |-
  Provides details about a specific Datastore in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_datastore
This data source provides details about a specific Datastore resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Get the specified Datastore's information.

## Example Usage

```hcl
data "oci_ocvp_datastore" "test_datastore" {
	#Required
	datastore_id = oci_ocvp_datastore.test_datastore.id
}
```

## Argument Reference

The following arguments are supported:

* `datastore_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Datastore. 


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

