---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadata_infrastructures"
sidebar_current: "docs-oci-datasource-database-exadata_infrastructures"
description: |-
  Provides the list of Exadata Infrastructures in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadata_infrastructures
This data source provides the list of Exadata Infrastructures in Oracle Cloud Infrastructure Database service.

Gets a list of the Exadata infrastructure in the specified compartment.


## Example Usage

```hcl
data "oci_database_exadata_infrastructures" "test_exadata_infrastructures" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.exadata_infrastructure_display_name
	state = var.exadata_infrastructure_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `exadata_infrastructures` - The list of exadata_infrastructures.

### ExadataInfrastructure Reference

The following attributes are exported:

* `admin_network_cidr` - The CIDR block for the Exadata administration network.
* `cloud_control_plane_server1` - The IP address for the first control plane server.
* `cloud_control_plane_server2` - The IP address for the second control plane server.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `contacts` - The list of contacts for the Exadata Infrastructure.
	* `email` - The email for the Exadata Infrastructure contact.
	* `is_primary` - True, if this Exadata Infrastructure contact is a primary contact. False, if this Exadata Infrastructure is a secondary contact.
	* `name` - The name of the Exadata Infrastructure contact.
	* `phone_number` - The phone number for the Exadata Infrastructure contact.
* `corporate_proxy` - The corporate network proxy for access to the control plane network.
* `cpus_enabled` - The number of enabled CPU cores.
* `csi_number` - The CSI Number of the Exadata Infrastructure.
* `data_storage_size_in_tbs` - Size, in terabytes, of the DATA disk group. 
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Exadata infrastructure. The name does not need to be unique.
* `dns_server` - The list of DNS server IP addresses. Maximum of 3 allowed.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway` - The gateway for the control plane network.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `infini_band_network_cidr` - The CIDR block for the Exadata InfiniBand interconnect.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `max_cpu_count` - The total number of CPU cores available.
* `max_data_storage_in_tbs` - The total available DATA disk group size.
* `max_db_node_storage_in_gbs` - The total local node storage available in GBs.
* `max_memory_in_gbs` - The total memory available in GBs.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `netmask` - The netmask for the control plane network.
* `ntp_server` - The list of NTP server IP addresses. Maximum of 3 allowed.
* `shape` - The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance. 
* `state` - The current lifecycle state of the Exadata infrastructure.
* `time_created` - The date and time the Exadata infrastructure was created.
* `time_zone` - The time zone of the Exadata infrastructure. For details, see [Exadata Infrastructure Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).

