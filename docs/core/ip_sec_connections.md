# oci_core_ipsec

## IpSecConnection Resource

### IpSecConnection Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the IPSec connection.
* `cpe_id` - The OCID of the CPE.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The OCID of the DRG.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The IPSec connection's Oracle ID (OCID).
* `state` - The IPSec connection's current state.
* `static_routes` - Static routes to the CPE. At least one route must be included. The CIDR must not be a multicast address or class E address.  Example: `10.0.1.0/24` 
* `time_created` - The date and time the IPSec connection was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new IPSec connection between the specified DRG and CPE. For more information, see
[IPSec VPNs](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIPsec.htm).

In the request, you must include at least one static route to the CPE object (you're allowed a maximum
of 10). For example: 10.0.8.0/16.

For the purposes of access control, you must provide the OCID of the compartment where you want the
IPSec connection to reside. Notice that the IPSec connection doesn't have to be in the same compartment
as the DRG, CPE, or other Networking Service components. If you're not sure which compartment to
use, put the IPSec connection in the same compartment as the DRG. For more information about
compartments and access control, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the IPSec connection, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.

After creating the IPSec connection, you need to configure your on-premises router
with tunnel-specific information returned by
[GetIPSecConnectionDeviceConfig](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnectionDeviceConfig/GetIPSecConnectionDeviceConfig).
For each tunnel, that operation gives you the IP address of Oracle's VPN headend and the shared secret
(that is, the pre-shared key). For more information, see
[Configuring Your On-Premises Router for an IPSec VPN](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/configuringCPE.htm).

To get the status of the tunnels (whether they're up or down), use
[GetIPSecConnectionDeviceStatus](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnectionDeviceStatus/GetIPSecConnectionDeviceStatus).


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the IPSec connection.
* `cpe_id` - (Required) The OCID of the CPE.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `drg_id` - (Required) The OCID of the DRG.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `static_routes` - (Required) Static routes to the CPE. At least one route must be included. The CIDR must not be a multicast address or class E address.  Example: `10.0.1.0/24` 


### Update Operation
Updates the display name or tags for the specified IPSec connection.
Avoid entering confidential information.


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_ipsec" "test_ip_sec_connection" {
	#Required
	compartment_id = "${var.compartment_id}"
	cpe_id = "${oci_core_cpe.test_cpe.id}"
	drg_id = "${oci_core_drg.test_drg.id}"
	static_routes = "${var.ip_sec_connection_static_routes}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.ip_sec_connection_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

# oci_core_ipsec_connections

## IpSecConnection DataSource

Gets a list of ip_sec_connections.

### List Operation
Lists the IPSec connections for the specified compartment. You can filter the
results by DRG or CPE.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `cpe_id` - (Optional) The OCID of the CPE.
* `drg_id` - (Optional) The OCID of the DRG.


The following attributes are exported:

* `connections` - The list of IPSec connections.

### Example Usage

```hcl
data "oci_core_ipsec_connections" "test_ip_sec_connections" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	cpe_id = "${oci_core_cpe.test_cpe.id}"
	drg_id = "${oci_core_drg.test_drg.id}"
}
```