---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_network"
sidebar_current: "docs-oci-datasource-database-vm_cluster_network"
description: |-
  Provides details about a specific Vm Cluster Network in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_network
This data source provides details about a specific Vm Cluster Network resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified VM cluster network. Applies to Exadata Cloud@Customer instances only.
To get information about a cloud VM cluster in an Exadata Cloud Service instance, use the [GetCloudVmCluster ](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudVmCluster/GetCloudVmCluster) operation.


## Example Usage

```hcl
data "oci_database_vm_cluster_network" "test_vm_cluster_network" {
	#Required
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
	vm_cluster_network_id = oci_database_vm_cluster_network.test_vm_cluster_network.id
}
```

## Argument Reference

The following arguments are supported:

* `exadata_infrastructure_id` - (Required) The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `vm_cluster_network_id` - (Required) The VM cluster network [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the VM cluster network. The name does not need to be unique.
* `dns` - The list of DNS server IP addresses. Maximum of 3 allowed.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `ntp` - The list of NTP server IP addresses. Maximum of 3 allowed.
* `scans` - The SCAN details.
	* `hostname` - The SCAN hostname.
	* `ips` - The list of SCAN IP addresses. Three addresses should be provided.
	* `port` - The SCAN TCPIP port. Default is 1521.
	* `scan_listener_port_tcp` - The SCAN TCPIP port. Default is 1521.
	* `scan_listener_port_tcp_ssl` - The SCAN TCPIP SSL port. Default is 2484.
* `state` - The current state of the VM cluster network.
* `time_created` - The date and time when the VM cluster network was created.
* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated VM Cluster.
* `vm_networks` - Details of the client and backup networks.
	* `domain_name` - The network domain name.
	* `gateway` - The network gateway.
	* `netmask` - The network netmask.
	* `network_type` - The network type.
	* `nodes` - The list of node details.
		* `hostname` - The node host name.
		* `ip` - The node IP address.
		* `vip` - The node virtual IP (VIP) address.
		* `vip_hostname` - The node virtual IP (VIP) host name.
	* `vlan_id` - The network VLAN ID.

