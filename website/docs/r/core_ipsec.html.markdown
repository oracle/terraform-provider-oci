---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec"
sidebar_current: "docs-oci-resource-core-ipsec"
description: |-
  Provides the Ip Sec Connection resource in Oracle Cloud Infrastructure Core service
---

# oci_core_ipsec
This resource provides the Ip Sec Connection resource in Oracle Cloud Infrastructure Core service.

Creates a new IPSec connection between the specified DRG and CPE with two default static tunnels. For more information, see
[Site-to-Site VPN Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/overviewIPsec.htm).

For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want the
IPSec connection to reside. Notice that the IPSec connection doesn't have to be in the same compartment
as the DRG, CPE, or other Networking Service components. If you're not sure which compartment to
use, put the IPSec connection in the same compartment as the DRG. For more information about
compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

You may optionally specify a *display name* for the IPSec connection, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.

After creating the IPSec connection, you need to configure your on-premises router
with tunnel-specific information. For tunnel status and the required configuration information, see:

  * [IPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnel/)
  * [IPSecConnectionTunnelSharedSecret](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnelSharedSecret/)

To configure tunnel-specific information, use `oci_core_ipsec_connection_tunnel_management` to update the tunnels. If 
you configure at least one tunnel to use static routing, then in the oci_core_ipsec request you must provide
at least one valid static route (you're allowed a maximum of 10). For example: 10.0.0.0/16.
If you configure both tunnels to use BGP dynamic routing, the static routes will be ignored. However, you must provide a
static route in `oci_core_ipsec` even if you plan to use BGP routing because it defaults to two static tunnels.  For more 
information, see the important note in [IPSecConnection](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnection/).

For each tunnel, you need the IP address of Oracle's VPN headend and the shared secret
(that is, the pre-shared key). For more information, see
[CPE Configuration](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/configuringCPE.htm).

To configure tunnel-specific information for private ipsec connection over fastconnect, use attribute `tunnel_configuration`.
You can provide configuration for maximum of 2 tunnels. You can configure each tunnel with `oracle_tunnel_ip`,
`associated_virtual_circuits` and `drg_route_table_id` at time of creation. These attributes cannot be updated using IPSec
connection APIs. To update drg route table id, use `oci_core_drg_attachment_management` resource to update.


## Example Usage

```hcl
resource "oci_core_ipsec" "test_ip_sec_connection" {
	#Required
	compartment_id = var.compartment_id
	cpe_id = oci_core_cpe.test_cpe.id
	drg_id = oci_core_drg.test_drg.id
	static_routes = var.ip_sec_connection_static_routes

	#Optional
	cpe_local_identifier = var.ip_sec_connection_cpe_local_identifier
	cpe_local_identifier_type = var.ip_sec_connection_cpe_local_identifier_type
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.ip_sec_connection_display_name
	freeform_tags = {"Department"= "Finance"}
}

resource "oci_core_ipsec" "test_ip_sec_connection_over_fc" {
	#Required
	compartment_id = var.compartment_id
	cpe_id = oci_core_cpe.test_cpe.id
	drg_id = oci_core_drg.test_drg.id
	static_routes = var.ip_sec_connection_static_routes

	#Optional
	cpe_local_identifier = var.ip_sec_connection_cpe_local_identifier
	cpe_local_identifier_type = var.ip_sec_connection_cpe_local_identifier_type
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.ip_sec_connection_display_name
	freeform_tags = {"Department"= "Finance"}
	tunnel_configuration {
		oracle_tunnel_ip = "10.1.5.5"
		associated_virtual_circuits = [oci_core_virtual_circuit.test_ipsec_over_fc_virtual_circuit.id]
		drg_route_table_id = oci_core_drg_route_table.test_drg_ipsec_over_fc_route_table.id
	}
	tunnel_configuration {
		oracle_tunnel_ip = "10.1.7.7"
		associated_virtual_circuits = [oci_core_virtual_circuit.test_ipsec_over_fc_virtual_circuit.id]
		drg_route_table_id = oci_core_drg_route_table.test_drg_ipsec_over_fc_route_table.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the IPSec connection.
* `cpe_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [Cpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Cpe/) object.
* `cpe_local_identifier` - (Optional) (Updatable) Your identifier for your CPE device. Can be either an IP address or a hostname (specifically, the fully qualified domain name (FQDN)). The type of identifier you provide here must correspond to the value for `cpeLocalIdentifierType`.

    If you don't provide a value, the `ipAddress` attribute for the [Cpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Cpe/) object specified by `cpeId` is used as the `cpeLocalIdentifier`.

    For information about why you'd provide this value, see [If Your CPE Is Behind a NAT Device](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/overviewIPsec.htm#nat).

    Example IP address: `10.0.3.3`

    Example hostname: `cpe.example.com` 
* `cpe_local_identifier_type` - (Optional) (Updatable) The type of identifier for your CPE device. The value you provide here must correspond to the value for `cpeLocalIdentifier`. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `static_routes` - (Required) (Updatable) Static routes to the CPE. A static route's CIDR must not be a multicast address or class E address.

    Used for routing a given IPSec tunnel's traffic only if the tunnel is using static routing. If you configure at least one tunnel to use static routing, then you must provide at least one valid static route. If you configure both tunnels to use BGP dynamic routing, you can provide an empty list for the static routes on update. For more information, see the important note in [IPSecConnection](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnection/).

     Example: `10.0.1.0/24` 
* `tunnel_configuration` - (Optional) (Non-updatable) Tunnel configuration for private ipsec connection over fastconnect.

  Example: `
  tunnel_configuration {
    oracle_tunnel_ip = "10.1.5.5"
    associated_virtual_circuits = [oci_core_virtual_circuit.test_ipsec_over_fc_virtual_circuit.id]
    drg_route_table_id = oci_core_drg_route_table.test_drg_ipsec_over_fc_route_table.id
  }`


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the IPSec connection.
* `cpe_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [Cpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Cpe/) object.
* `cpe_local_identifier` - Your identifier for your CPE device. Can be either an IP address or a hostname (specifically, the fully qualified domain name (FQDN)). The type of identifier here must correspond to the value for `cpeLocalIdentifierType`.

    If you don't provide a value when creating the IPSec connection, the `ipAddress` attribute for the [Cpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Cpe/) object specified by `cpeId` is used as the `cpeLocalIdentifier`.

    For information about why you'd provide this value, see [If Your CPE Is Behind a NAT Device](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/overviewIPsec.htm#nat).

    Example IP address: `10.0.3.3`

    Example hostname: `cpe.example.com` 
* `cpe_local_identifier_type` - The type of identifier for your CPE device. The value here must correspond to the value for `cpeLocalIdentifier`. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The IPSec connection's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `state` - The IPSec connection's current state.
* `static_routes` - Static routes to the CPE. The CIDR must not be a multicast address or class E address.

    Used for routing a given IPSec tunnel's traffic only if the tunnel is using static routing. If you configure at least one tunnel to use static routing, then you must provide at least one valid static route. If you configure both tunnels to use BGP dynamic routing, the static routes will be ignored.

     Example: `10.0.1.0/24`
* `tunnel_configuration` - Tunnel configuration for private ipsec connection over fastconnect.

	Example: `
	tunnel_configuration {
	  oracle_tunnel_ip = "10.1.5.5"
	  associated_virtual_circuits = [oci_core_virtual_circuit.test_ipsec_over_fc_virtual_circuit.id]
	  drg_route_table_id = oci_core_drg_route_table.test_drg_ipsec_over_fc_route_table.id
	}`
* `time_created` - The date and time the IPSec connection was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `transport_type` - The transport type used for the IPSec connection. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ip Sec Connection
	* `update` - (Defaults to 20 minutes), when updating the Ip Sec Connection
	* `delete` - (Defaults to 20 minutes), when destroying the Ip Sec Connection


## Import

IpSecConnections can be imported using the `id`, e.g.

```
$ terraform import oci_core_ipsec.test_ip_sec_connection "id"
```

