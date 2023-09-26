---
subcategory: "Compute Cloud At Customer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_compute_cloud_at_customer_ccc_infrastructures"
sidebar_current: "docs-oci-datasource-compute_cloud_at_customer-ccc_infrastructures"
description: |-
  Provides the list of Ccc Infrastructures in Oracle Cloud Infrastructure Compute Cloud At Customer service
---

# Data Source: oci_compute_cloud_at_customer_ccc_infrastructures
This data source provides the list of Ccc Infrastructures in Oracle Cloud Infrastructure Compute Cloud At Customer service.

Returns a list of Compute Cloud@Customer infrastructures.


## Example Usage

```hcl
data "oci_compute_cloud_at_customer_ccc_infrastructures" "test_ccc_infrastructures" {

	#Optional
	access_level = var.ccc_infrastructure_access_level
	ccc_infrastructure_id = oci_compute_cloud_at_customer_ccc_infrastructure.test_ccc_infrastructure.id
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.ccc_infrastructure_compartment_id_in_subtree
	display_name = var.ccc_infrastructure_display_name
	display_name_contains = var.ccc_infrastructure_display_name_contains
	state = var.ccc_infrastructure_state
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `ccc_infrastructure_id` - (Optional) An [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for a  Compute Cloud@Customer Infrastructure. 
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and sub-compartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `display_name_contains` - (Optional) A filter to return only resources whose display name contains the substring. 
* `state` - (Optional) A filter used to return only resources that match the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `ccc_infrastructure_collection` - The list of ccc_infrastructure_collection.

### CccInfrastructure Reference

The following attributes are exported:

* `ccc_upgrade_schedule_id` - Schedule used for upgrades. If no schedule is associated with the infrastructure, it can be updated at any time. 
* `compartment_id` - The infrastructure compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `connection_details` - A message describing the current connection state in more detail. 
* `connection_state` - The current connection state of the infrastructure. A user can only update it from REQUEST to READY or from any state back to REJECT. The system automatically handles the REJECT to REQUEST, READY to CONNECTED, or CONNECTED to DISCONNECTED transitions. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A mutable client-meaningful text description of the Compute Cloud@Customer infrastructure. Avoid entering confidential information. 
* `display_name` - The name that will be used to display the Compute Cloud@Customer infrastructure in the Oracle Cloud Infrastructure console. Does not have to be unique and can be changed. Avoid entering confidential information. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The Compute Cloud@Customer infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). This cannot be changed once created. 
* `infrastructure_inventory` - Inventory for a Compute Cloud@Customer infrastructure. This information cannot be updated and is from the infrastructure. The information will only be available after the connectionState is transitioned to CONNECTED. 
	* `capacity_storage_tray_count` - The number of storage trays in the Compute Cloud@Customer infrastructure rack that are designated for capacity storage.
	* `compute_node_count` - The number of compute nodes that are available and usable on the Compute Cloud@Customer infrastructure rack. There is no distinction of compute node type in this information. 
	* `management_node_count` - The number of management nodes that are available and in active use on the Compute Cloud@Customer infrastructure rack. 
	* `performance_storage_tray_count` - The number of storage trays in the Compute Cloud@Customer infrastructure rack that are designated for performance storage.
	* `serial_number` - The serial number of the Compute Cloud@Customer infrastructure rack. 
* `infrastructure_network_configuration` - Configuration information for the Compute Cloud@Customer infrastructure. This  network configuration information cannot be updated and is retrieved from the data center. The information will only be available after the connectionState is transitioned to CONNECTED. 
	* `dns_ips` - The domain name system (DNS) addresses that the Compute Cloud@Customer infrastructure uses for the data center network. 
	* `infrastructure_routing_dynamic` - Dynamic routing information for the Compute Cloud@Customer infrastructure. 
		* `bgp_topology` - The topology in use for the Border Gateway Protocol (BGP) configuration. 
		* `oracle_asn` - The Oracle Autonomous System Number (ASN) to control routing and exchange information within the dynamic routing configuration. 
		* `peer_information` - The list of peer devices in the dynamic routing configuration.
			* `asn` - The Autonomous System Number (ASN) of the peer network.
			* `ip` - Neighbor Border Gateway Protocal (BGP) IP address. The IP address usually refers to the customer data center router. 
	* `infrastructure_routing_static` - Static routing information for a rack.
		* `uplink_hsrp_group` - The uplink Hot Standby Router Protocol (HSRP) group value for the switch in the Compute Cloud@Customer infrastructure. 
		* `uplink_vlan` - The virtual local area network (VLAN) identifier used to connect to the uplink (only access mode is supported). 
	* `management_nodes` - Information about the management nodes that are provisioned in the Compute Cloud@Customer infrastructure. 
		* `hostname` - Hostname for interface to the management node.
		* `ip` - Address of the management node.
	* `mgmt_vip_hostname` - The hostname corresponding to the virtual IP (VIP) address of the management nodes. 
	* `mgmt_vip_ip` - The IP address used as the virtual IP (VIP) address of the management nodes.
	* `spine_ips` - Addresses of the network spine switches.
	* `spine_vip` - The spine switch public virtual IP (VIP). Traffic routed to the Compute Cloud@Customer infrastructure and  and virtual cloud networks (VCNs) should have this address as next hop. 
	* `uplink_domain` - Domain name to be used as the base domain for the internal network and by  public facing services. 
	* `uplink_gateway_ip` - Uplink gateway in the datacenter network that the Compute Cloud@Customer connects to. 
	* `uplink_netmask` - Netmask of the subnet that the Compute Cloud@Customer infrastructure is connected to. 
	* `uplink_port_count` - Number of uplink ports per spine switch. Connectivity is identical on both spine switches. For example, if input is two 100 gigabyte ports; then port-1 and port-2 on both spines will be configured. 
	* `uplink_port_forward_error_correction` - The port forward error correction (FEC) setting for the uplink port on the Compute Cloud@Customer infrastructure. 
	* `uplink_port_speed_in_gbps` - Uplink port speed defined in gigabytes per second. All uplink ports must have identical speed. 
	* `uplink_vlan_mtu` - The virtual local area network (VLAN) maximum transmission unit (MTU) size for the uplink ports. 
* `lifecycle_details` - A message describing the current lifecycle state in more detail. For example, this can be used to provide actionable information for a resource that is in a Failed state. 
* `provisioning_fingerprint` - Fingerprint of a Compute Cloud@Customer infrastructure in a data center generated during the initial connection to this resource. The fingerprint should be verified by the administrator when changing the connectionState from REQUEST to READY. 
* `provisioning_pin` - Code that is required for service personnel to connect a Compute Cloud@Customer infrastructure in a data center to this resource. This code will only be available when the connectionState is REJECT (usually at create time of the Compute Cloud@Customer infrastructure). 
* `short_name` - The Compute Cloud@Customer infrastructure short name. This cannot be changed once created. The short name is used to refer to the infrastructure in several contexts and is unique. 
* `state` - The current state of the Compute Cloud@Customer infrastructure.
* `subnet_id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network subnet that is used to communicate with Compute Cloud@Customer infrastructure. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Compute Cloud@Customer infrastructure creation date and time, using an RFC3339 formatted datetime string. 
* `time_updated` - Compute Cloud@Customer infrastructure updated date and time, using an RFC3339 formatted datetime string. 
* `upgrade_information` - Upgrade information that relates to a Compute Cloud@Customer infrastructure. This information cannot be updated. 
	* `current_version` - The current version of software installed on the Compute Cloud@Customer infrastructure. 
	* `is_active` - Indication that the Compute Cloud@Customer infrastructure is in the process of an upgrade or an upgrade activity (such as preloading upgrade images). 
	* `scheduled_upgrade_duration` - Expected duration of Compute Cloud@Customer infrastructure scheduled upgrade. The actual upgrade time might be longer or shorter than this duration depending on rack activity, this is only an estimate. 
	* `time_of_scheduled_upgrade` - Compute Cloud@Customer infrastructure next upgrade time. The rack might have performance impacts during this time. 

