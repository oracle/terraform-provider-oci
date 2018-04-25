# oci_core_service_gateway

## ServiceGateway Resource

### ServiceGateway Reference

The following attributes are exported:

* `block_traffic` - Boolean to allow/disallow traffic through Service Gateway. This will be False by default
* `compartment_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the Service Gateway.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The Service Gateway's Oracle ID ([OCID])(/Content/General/Concepts/identifiers.htm).
* `services` - List of objects of Service OCID and name. These are the Services which have been enabled on the Service Gateway.
	* `service_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Service. 
	* `service_name` - Name of the Service
* `state` - The Service Gateway's current state.
* `time_created` - The date and time the Service Gateway was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VCN.



### Create Operation
Creates a new service gateway in the specified compartment.

For the purposes of access control, you must provide the OCID of the compartment where you want
the service gateway to reside. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the service gateway, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.


The following arguments are supported:

* `compartment_id` - (Required) The [OCID] (/Content/General/Concepts/identifiers.htm)  of the compartment to contain the Service Gateway. 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `services` - (Required) List of the Service OCIDs. These are the Services which will be enabled on the Service Gateway. This list can be empty.
	* `service_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Service. 
* `vcn_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VCN.


### Update Operation
Updates the specified service gateway. The information you provide overwrites the existing
attributes of the gateway.


The following arguments support updates:
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `services` - List of the Service OCIDs. These are the Services which will be enabled on the Service Gateway. This list can be empty.
	* `service_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Service. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_core_service_gateway" "test_service_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	services {
		#Required
		service_id = "${oci_core_service.test_service.id}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.service_gateway_display_name}"
}
```

# oci_core_service_gateways

## ServiceGateway DataSource

Gets a list of service_gateways.

### List Operation
Lists the service gateways in the specified compartment. You may optionally specify a VCN OCID
to filter the results by VCN.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state.  The state value is case-insensitive. 
* `vcn_id` - (Optional) The OCID of the VCN.


The following attributes are exported:

* `service_gateways` - The list of service_gateways.

### Example Usage

```
data "oci_core_service_gateways" "test_service_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	state = "${var.service_gateway_state}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
```