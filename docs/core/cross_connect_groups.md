# oci_core_cross_connect_group

## CrossConnectGroup Resource

### CrossConnectGroup Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the cross-connect group.
* `display_name` - The display name of A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The cross-connect group's Oracle ID (OCID).
* `state` - The cross-connect group's current state.
* `time_created` - The date and time the cross-connect group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new cross-connect group to use with Oracle Cloud Infrastructure
FastConnect. For more information, see
[FastConnect Overview](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/fastconnect.htm).

For the purposes of access control, you must provide the OCID of the
compartment where you want the cross-connect group to reside. If you're
not sure which compartment to use, put the cross-connect group in the
same compartment with your VCN. For more information about
compartments and access control, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the cross-connect group.
It does not have to be unique, and you can change it. Avoid entering confidential information.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the cross-connect group.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 


### Update Operation
Updates the specified cross-connect group's display name.
Avoid entering confidential information.


The following arguments support updates:
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_cross_connect_group" "test_cross_connect_group" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.cross_connect_group_display_name}"
}
```


## CrossConnectGroup Singular DataSource


### Get Operation
Gets the specified cross-connect group's information.

The following arguments are supported:

* `cross_connect_group_id` - (Required) The OCID of the cross-connect group.


### Example Usage

```hcl
data "oci_core_cross_connect_group" "test_cross_connect_group" {
	#Required
	cross_connect_group_id = "${oci_core_cross_connect_group.test_cross_connect_group.id}"
}
```
# oci_core_cross_connect_groups

## CrossConnectGroup DataSource

Gets a list of cross_connect_groups.

### List Operation
Lists the cross-connect groups in the specified compartment.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive. 


The following attributes are exported:

* `cross_connect_groups` - The list of cross_connect_groups.

### Example Usage

```hcl
data "oci_core_cross_connect_groups" "test_cross_connect_groups" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.cross_connect_group_display_name}"
	state = "${var.cross_connect_group_state}"
}
```