---
subcategory: "Datacc"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacc_vm_cluster_network"
sidebar_current: "docs-oci-datasource-datacc-vm_cluster_network"
description: |-
  Provides details about a specific Vm Cluster Network in Oracle Cloud Infrastructure Datacc service
---

# Data Source: oci_datacc_vm_cluster_network
This data source provides details about a specific Vm Cluster Network resource in Oracle Cloud Infrastructure Datacc service.

Obtain the VM cluster network on Database Infrastructure that has the specified
[OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_datacc_vm_cluster_network" "test_vm_cluster_network" {
	#Required
	vm_cluster_network_id = oci_datacc_vm_cluster_network.test_vm_cluster_network.id
}
```

## Argument Reference

The following arguments are supported:

* `vm_cluster_network_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.


## Attributes Reference

The following attributes are exported:

* `associated_resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated resource.
* `base_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated VM cluster.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VM cluster network. 
* `consumer_type` - Consumer type for the VM cluster network.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The user-friendly name for the VM cluster network. The name does not need to be unique.
* `dns_servers` - The list of DNS server IP addresses. Maximum of 3 allowed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.
* `infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
* `is_scan_enabled` - Indicates whether Single Client Access Name (SCAN) is enabled on the VM cluster. 
* `lifecycle_details` - Lifecycle state details of the VM cluster network.
* `listener_port` - The listener TCP/IP port.
* `listener_port_ssl` - The listener TCP/IP SSL port.
* `node_count` - Count of virtual machines in this VM cluster.
* `ntp_servers` - The list of NTP server IP addresses. Maximum of 3 allowed.
* `scans` - The SCAN details.
	* `hostname` - The SCAN hostname.
	* `ips` - The list of SCAN IP addresses. Three addresses should be provided.
* `state` - The current state of the virtual machine cluster network.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time that the VM cluster network was created. An RFC3339 formatted datetime string. 
* `time_updated` - The time that the VM cluster network was last updated. An RFC3339 formatted datetime string. 
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
	* `prefix` - The network domain name prefix.
	* `vlan_id` - The network VLAN ID.

