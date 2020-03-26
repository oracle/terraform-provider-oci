---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadata_infrastructure"
sidebar_current: "docs-oci-resource-database-exadata_infrastructure"
description: |-
  Provides the Exadata Infrastructure resource in Oracle Cloud Infrastructure Database service
---

# oci_database_exadata_infrastructure
This resource provides the Exadata Infrastructure resource in Oracle Cloud Infrastructure Database service.

Create Exadata infrastructure.

## Example Usage

```hcl
resource "oci_database_exadata_infrastructure" "test_exadata_infrastructure" {
	#Required
	admin_network_cidr = "${var.exadata_infrastructure_admin_network_cidr}"
	cloud_control_plane_server1 = "${var.exadata_infrastructure_cloud_control_plane_server1}"
	cloud_control_plane_server2 = "${var.exadata_infrastructure_cloud_control_plane_server2}"
	compartment_id = "${var.compartment_id}"
	display_name = "${var.exadata_infrastructure_display_name}"
	dns_server = "${var.exadata_infrastructure_dns_server}"
	gateway = "${var.exadata_infrastructure_gateway}"
	infini_band_network_cidr = "${var.exadata_infrastructure_infini_band_network_cidr}"
	netmask = "${var.exadata_infrastructure_netmask}"
	ntp_server = "${var.exadata_infrastructure_ntp_server}"
	shape = "${var.exadata_infrastructure_shape}"
	time_zone = "${var.exadata_infrastructure_time_zone}"

	#Optional
	activation_file = "${var.exadata_infrastructure_activation_file}"
	corporate_proxy = "${var.exadata_infrastructure_corporate_proxy}"
	defined_tags = "${var.exadata_infrastructure_defined_tags}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `activation_file` - (Optional) (Updatable) The activation zip file. If provided in config, exadata infrastructure will be activated after creation. Updates are not allowed on activated exadata infrastructure.
* `admin_network_cidr` - (Required) (Updatable) The CIDR block for the Exadata administration network.
* `cloud_control_plane_server1` - (Required) (Updatable) The IP address for the first control plane server.
* `cloud_control_plane_server2` - (Required) (Updatable) The IP address for the second control plane server.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `corporate_proxy` - (Optional) (Updatable) The corporate network proxy for access to the control plane network. Oracle recommends using an HTTPS proxy when possible for enhanced security. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) The user-friendly name for the Exadata infrastructure. The name does not need to be unique. 
* `dns_server` - (Required) (Updatable) The list of DNS server IP addresses. Maximum of 3 allowed.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway` - (Required) (Updatable) The gateway for the control plane network.
* `infini_band_network_cidr` - (Required) (Updatable) The CIDR block for the Exadata InfiniBand interconnect.
* `netmask` - (Required) (Updatable) The netmask for the control plane network.
* `ntp_server` - (Required) (Updatable) The list of NTP server IP addresses. Maximum of 3 allowed.
* `shape` - (Required) The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance. 
* `time_zone` - (Required) (Updatable) The time zone of the Exadata infrastructure. For details, see [Exadata Infrastructure Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Import

ExadataInfrastructures can be imported using the `id`, e.g.

```
$ terraform import oci_database_exadata_infrastructure.test_exadata_infrastructure "id"
```

