---
subcategory: "Datacc"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacc_vm_cluster_network"
sidebar_current: "docs-oci-resource-datacc-vm_cluster_network"
description: |-
  Provides the Vm Cluster Network resource in Oracle Cloud Infrastructure Datacc service
---

# oci_datacc_vm_cluster_network
This resource provides the Vm Cluster Network resource in Oracle Cloud Infrastructure Datacc service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datacc

Create an VM cluster on Database Infrastructure network using the specified details.


## Example Usage

```hcl
resource "oci_datacc_vm_cluster_network" "test_vm_cluster_network" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.vm_cluster_network_display_name
	infrastructure_id = oci_datacc_infrastructure.test_infrastructure.id
	vm_networks {
		#Required
		domain_name = oci_identity_domain.test_domain.name
		gateway = var.vm_cluster_network_vm_networks_gateway
		netmask = var.vm_cluster_network_vm_networks_netmask
		network_type = var.vm_cluster_network_vm_networks_network_type
		nodes {
			#Required
			hostname = var.vm_cluster_network_vm_networks_nodes_hostname
			ip = var.vm_cluster_network_vm_networks_nodes_ip

			#Optional
			vip = var.vm_cluster_network_vm_networks_nodes_vip
			vip_hostname = var.vm_cluster_network_vm_networks_nodes_vip_hostname
		}

		#Optional
		prefix = var.vm_cluster_network_vm_networks_prefix
		vlan_id = oci_core_vlan.test_vlan.id
	}

	#Optional
	consumer_type = var.vm_cluster_network_consumer_type
	defined_tags = {"foo-namespace.bar-key"= "value"}
	dns_servers = var.vm_cluster_network_dns_servers
	freeform_tags = {"bar-key"= "value"}
	listener_port = var.vm_cluster_network_listener_port
	listener_port_ssl = var.vm_cluster_network_listener_port_ssl
	node_count = var.vm_cluster_network_node_count
	ntp_servers = var.vm_cluster_network_ntp_servers
	scans {
		#Required
		hostname = var.vm_cluster_network_scans_hostname
		ips = var.vm_cluster_network_scans_ips
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VM cluster network. 
* `consumer_type` - (Optional) Consumer type for the VM cluster network.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) The user-friendly name for the VM cluster network. The name does not need to be unique.
* `dns_servers` - (Optional) (Updatable) The list of DNS server IP addresses. Maximum of 3 allowed.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
* `listener_port` - (Optional) (Updatable) The listener TCP/IP port.
* `listener_port_ssl` - (Optional) (Updatable) The listener TCP/IP SSL port. Default is 2484.
* `node_count` - (Optional) Count of virtual machines in this VM cluster.
* `ntp_servers` - (Optional) (Updatable) The list of NTP server IP addresses. Maximum of 3 allowed.
* `scans` - (Optional) (Updatable) The SCAN details.
	* `hostname` - (Required) (Updatable) The SCAN hostname.
	* `ips` - (Required) (Updatable) The list of SCAN IP addresses. Three addresses should be provided.
* `vm_networks` - (Required) (Updatable) Details of the client and backup networks.
	* `domain_name` - (Required) (Updatable) The network domain name.
	* `gateway` - (Required) (Updatable) The network gateway.
	* `netmask` - (Required) (Updatable) The network netmask.
	* `network_type` - (Required) (Updatable) The network type.
	* `nodes` - (Required) (Updatable) The list of node details.
		* `hostname` - (Required) (Updatable) The node host name.
		* `ip` - (Required) (Updatable) The node IP address.
		* `vip` - (Optional) (Updatable) The node virtual IP (VIP) address.
		* `vip_hostname` - (Optional) (Updatable) The node virtual IP (VIP) host name.
	* `prefix` - (Optional) (Updatable) The network domain name prefix.
	* `vlan_id` - (Optional) (Updatable) The network VLAN ID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Vm Cluster Network
	* `update` - (Defaults to 20 minutes), when updating the Vm Cluster Network
	* `delete` - (Defaults to 20 minutes), when destroying the Vm Cluster Network


## Import

VmClusterNetworks can be imported using the `id`, e.g.

```
$ terraform import oci_datacc_vm_cluster_network.test_vm_cluster_network "id"
```

