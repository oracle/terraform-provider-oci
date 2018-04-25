# oci_core_service_gateway

## ServiceGateway Resource

### ServiceGateway Reference

The following attributes are exported:

* `block_traffic` - Whether the service gateway blocks all traffic through it. The default is `false`. When this is `true`, traffic is not routed to any services, regardless of route rules.  Example: `true` 
* `compartment_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the service gateway. 
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "foo-value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the service gateway.
* `services` - List of the services enabled on this service gateway. The list can be empty. You can enable a particular service by using [AttachServiceId](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ServiceGateway/AttachServiceId). 
	* `service_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the service. 
	* `service_name` - The name of the service.
* `state` - The service gateway's current state.
* `time_created` - The date and time the service gateway was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VCN the service gateway belongs to. 



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
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `services` - (Required) List of the Service OCIDs. These are the Services which will be enabled on the Service Gateway. This list can be empty.
	* `service_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the service. 
* `vcn_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VCN.


### Update Operation
Updates the specified service gateway. The information you provide overwrites the existing
attributes of the gateway.


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `services` - List of the Service OCIDs. These are the Services which will be enabled on the Service Gateway. This list can be empty.
	* `service_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the service. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_service_gateway" "test_service_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	services {
		#Required
		service_id = "${oci_core_service.test_service.id}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "foo-value"}
	display_name = "${var.service_gateway_display_name}"
	freeform_tags = {"bar-key"= "value"}
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

```hcl
data "oci_core_service_gateways" "test_service_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	state = "${var.service_gateway_state}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
```