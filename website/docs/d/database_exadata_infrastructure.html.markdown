---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadata_infrastructure"
sidebar_current: "docs-oci-datasource-database-exadata_infrastructure"
description: |-
  Provides details about a specific Exadata Infrastructure in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadata_infrastructure
This data source provides details about a specific Exadata Infrastructure resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Exadata infrastructure.

## Example Usage

```hcl
data "oci_database_exadata_infrastructure" "test_exadata_infrastructure" {
	#Required
	exadata_infrastructure_id = "${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}"
}
```

## Argument Reference

The following arguments are supported:

* `exadata_infrastructure_id` - (Required) The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `admin_network_cidr` - The CIDR block for the Exadata administration network.
* `cloud_control_plane_server1` - The IP address for the first control plane server.
* `cloud_control_plane_server2` - The IP address for the second control plane server.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `corporate_proxy` - The corporate network proxy for access to the control plane network.
* `cpus_enabled` - The number of enabled CPU cores.
* `data_storage_size_in_tbs` - Size, in terabytes, of the DATA disk group. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Exadata infrastructure. The name does not need to be unique.
* `dns_server` - The list of DNS server IP addresses. Maximum of 3 allowed.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway` - The gateway for the control plane network.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `infini_band_network_cidr` - The CIDR block for the Exadata InfiniBand interconnect.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `netmask` - The netmask for the control plane network.
* `ntp_server` - The list of NTP server IP addresses. Maximum of 3 allowed.
* `shape` - The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance. 
* `state` - The current lifecycle state of the Exadata infrastructure.
* `time_created` - The date and time the Exadata infrastructure was created.
* `time_zone` - The time zone of the Exadata infrastructure. For details, see [Exadata Infrastructure Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).

