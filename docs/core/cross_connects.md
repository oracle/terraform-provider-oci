# oci_core_cross_connect

## CrossConnect Resource

### CrossConnect Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the cross-connect group.
* `cross_connect_group_id` - The OCID of the cross-connect group this cross-connect belongs to (if any).
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The cross-connect's Oracle ID (OCID).
* `location_name` - The name of the FastConnect location where this cross-connect is installed.
* `port_name` - A string identifying the meet-me room port for this cross-connect.
* `port_speed_shape_name` - The port speed for this cross-connect.  Example: `10 Gbps` 
* `state` - The cross-connect's current state.
* `time_created` - The date and time the cross-connect was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new cross-connect. Oracle recommends you create each cross-connect in a
[CrossConnectGroup](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/CrossConnectGroup) so you can use link aggregation
with the connection.

After creating the `CrossConnect` object, you need to go the FastConnect location
and request to have the physical cable installed. For more information, see
[FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).

For the purposes of access control, you must provide the OCID of the
compartment where you want the cross-connect to reside. If you're
not sure which compartment to use, put the cross-connect in the
same compartment with your VCN. For more information about
compartments and access control, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the cross-connect.
It does not have to be unique, and you can change it. Avoid entering confidential information.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the cross-connect.
* `cross_connect_group_id` - (Optional) The OCID of the cross-connect group to put this cross-connect in.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `far_cross_connect_or_cross_connect_group_id` - (Optional) If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on a different router (for the purposes of redundancy), provide the OCID of that existing cross-connect or cross-connect group. 
* `location_name` - (Required) The name of the FastConnect location where this cross-connect will be installed. To get a list of the available locations, see [ListCrossConnectLocations](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/CrossConnectLocation/ListCrossConnectLocations).  Example: `CyrusOne, Chandler, AZ` 
* `near_cross_connect_or_cross_connect_group_id` - (Optional) If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on the same router, provide the OCID of that existing cross-connect or cross-connect group. 
* `port_speed_shape_name` - (Required) The port speed for this cross-connect. To get a list of the available port speeds, see [ListCrossConnectPortSpeedShapes](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/CrossConnectPortSpeedShape/ListCrossconnectPortSpeedShapes).  Example: `10 Gbps` 
* `is_active` - (Optional) Set to true to activate the cross-connect. You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up. Activation indicates to Oracle that the physical connection is ready. 


### Update Operation
Updates the specified cross-connect.

The following arguments support updates:
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `is_active` - (Optional) Set to true to activate the cross-connect. You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up. Activation indicates to Oracle that the physical connection is ready.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_cross_connect" "test_cross_connect" {
	#Required
	compartment_id = "${var.compartment_id}"
	location_name = "${var.cross_connect_location_name}"
	port_speed_shape_name = "${var.cross_connect_port_speed_shape_name}"

	#Optional
	cross_connect_group_id = "${oci_core_cross_connect_group.test_cross_connect_group.id}"
	display_name = "${var.cross_connect_display_name}"
	far_cross_connect_or_cross_connect_group_id = "${oci_core_far_cross_connect_or_cross_connect_group.test_far_cross_connect_or_cross_connect_group.id}"
	near_cross_connect_or_cross_connect_group_id = "${oci_core_near_cross_connect_or_cross_connect_group.test_near_cross_connect_or_cross_connect_group.id}"
}
```


## CrossConnect Singular DataSource


### Get Operation
Gets the specified cross-connect's information.

The following arguments are supported:

* `cross_connect_id` - (Required) The OCID of the cross-connect.


### Example Usage

```hcl
data "oci_core_cross_connect" "test_cross_connect" {
	#Required
	cross_connect_id = "${var.cross_connect_cross_connect_id}"
}
```
# oci_core_cross_connects

## CrossConnect DataSource

Gets a list of cross_connects.

### List Operation
Lists the cross-connects in the specified compartment. You can filter the list
by specifying the OCID of a cross-connect group.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `cross_connect_group_id` - (Optional) The OCID of the cross-connect group.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive. 


The following attributes are exported:

* `cross_connects` - The list of cross_connects.

### Example Usage

```hcl
data "oci_core_cross_connects" "test_cross_connects" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	cross_connect_group_id = "${oci_core_cross_connect_group.test_cross_connect_group.id}"
	display_name = "${var.cross_connect_display_name}"
	state = "${var.cross_connect_state}"
}
```