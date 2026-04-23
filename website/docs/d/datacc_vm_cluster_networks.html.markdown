---
subcategory: "Datacc"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacc_vm_cluster_networks"
sidebar_current: "docs-oci-datasource-datacc-vm_cluster_networks"
description: |-
  Provides the list of Vm Cluster Networks in Oracle Cloud Infrastructure Datacc service
---

# Data Source: oci_datacc_vm_cluster_networks
This data source provides the list of Vm Cluster Networks in Oracle Cloud Infrastructure Datacc service.

Obtain a list of VM cluster networks on Database Infrastructure.


## Example Usage

```hcl
data "oci_datacc_vm_cluster_networks" "test_vm_cluster_networks" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.vm_cluster_network_display_name
	infrastructure_id = oci_datacc_infrastructure.test_infrastructure.id
	is_scan_enabled = var.vm_cluster_network_is_scan_enabled
	node_count = var.vm_cluster_network_node_count
	state = var.vm_cluster_network_state
	vm_network_consumer_type = var.vm_cluster_network_vm_network_consumer_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. For list operations, you may provide the tenant [OCID] in this field. When a tenant OCID is provided, it will be validated against the caller's tenant and then treated as tenant scope (compartmentId filtering is not applied). 
* `display_name` - (Optional) A filter to return resources that match the entire display name given. The match is case sensitive.
* `infrastructure_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
* `is_scan_enabled` - (Optional) A filter to return VM cluster network resources that matches the specified value.
* `node_count` - (Optional) Count of virtual machines in this VM cluster.
* `state` - (Optional) A filter to return resources that match the specified lifecycle state.
* `vm_network_consumer_type` - (Optional) VM network consumer type.


## Attributes Reference

The following attributes are exported:

* `vm_cluster_network_collection` - The list of vm_cluster_network_collection.

### VmClusterNetwork Reference

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

