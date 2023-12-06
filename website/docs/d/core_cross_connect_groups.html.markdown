---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect_groups"
sidebar_current: "docs-oci-datasource-core-cross_connect_groups"
description: |-
  Provides the list of Cross Connect Groups in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cross_connect_groups
This data source provides the list of Cross Connect Groups in Oracle Cloud Infrastructure Core service.

Lists the cross-connect groups in the specified compartment.


## Example Usage

```hcl
data "oci_core_cross_connect_groups" "test_cross_connect_groups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.cross_connect_group_display_name
	state = var.cross_connect_group_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive. 


## Attributes Reference

The following attributes are exported:

* `cross_connect_groups` - The list of cross_connect_groups.

### CrossConnectGroup Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the cross-connect group.
* `customer_reference_name` - A reference name or identifier for the physical fiber connection that this cross-connect group uses. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The cross-connect group's Oracle ID (OCID).
* `macsec_properties` - Properties used for MACsec (if capable).
	* `encryption_cipher` - Type of encryption cipher suite to use for the MACsec connection.
	* `is_unprotected_traffic_allowed` - Indicates whether unencrypted traffic is allowed if MACsec Key Agreement protocol (MKA) fails.
	* `primary_key` - An object defining the Secrets-in-Vault OCIDs representing the MACsec key.
		* `connectivity_association_key_secret_id` - Secret [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) containing the Connectivity Association Key (CAK) of this MACsec key.
		* `connectivity_association_key_secret_version` - The secret version of the `connectivityAssociationKey` secret in Vault.
		* `connectivity_association_name_secret_id` - Secret [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) containing the Connectivity association Key Name (CKN) of this MACsec key.
		* `connectivity_association_name_secret_version` - The secret version of the connectivity association name secret in Vault.
	* `state` - Indicates whether or not MACsec is enabled.
* `oci_logical_device_name` - The FastConnect device that terminates the logical connection. This device might be different than the device that terminates the physical connection. 
* `oci_physical_device_name` - The FastConnect device that terminates the physical connection. 
* `state` - The cross-connect group's current state.
* `time_created` - The date and time the cross-connect group was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

