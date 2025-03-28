---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect_group"
sidebar_current: "docs-oci-resource-core-cross_connect_group"
description: |-
  Provides the Cross Connect Group resource in Oracle Cloud Infrastructure Core service
---

# oci_core_cross_connect_group
This resource provides the Cross Connect Group resource in Oracle Cloud Infrastructure Core service.

Creates a new cross-connect group to use with Oracle Cloud Infrastructure
FastConnect. For more information, see
[FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).

For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the
compartment where you want the cross-connect group to reside. If you're
not sure which compartment to use, put the cross-connect group in the
same compartment with your VCN. For more information about
compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the cross-connect group.
It does not have to be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_cross_connect_group" "test_cross_connect_group" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	customer_reference_name = var.cross_connect_group_customer_reference_name
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.cross_connect_group_display_name
	freeform_tags = {"Department"= "Finance"}
	macsec_properties {
		#Required
		state = var.cross_connect_group_macsec_properties_state

		#Optional
		encryption_cipher = var.cross_connect_group_macsec_properties_encryption_cipher
		is_unprotected_traffic_allowed = var.cross_connect_group_macsec_properties_is_unprotected_traffic_allowed
		primary_key {
			#Required
			connectivity_association_key_secret_id = oci_vault_secret.test_secret.id
			connectivity_association_name_secret_id = oci_vault_secret.test_secret.id
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the cross-connect group.
* `customer_reference_name` - (Optional) (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect group uses. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `macsec_properties` - (Optional) (Updatable) Properties used to configure MACsec (if capable).
	* `encryption_cipher` - (Optional) (Updatable) Type of encryption cipher suite to use for the MACsec connection.
	* `is_unprotected_traffic_allowed` - (Optional) (Updatable) Indicates whether unencrypted traffic is allowed if MACsec Key Agreement protocol (MKA) fails.
	* `primary_key` - (Optional) (Updatable) Defines the secret [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s held in Vault that represent the MACsec key.
		* `connectivity_association_key_secret_id` - (Required) (Updatable) Secret [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) containing the Connectivity Association Key (CAK) of this MACsec key.
		* `connectivity_association_key_secret_version` - (Optional) (Updatable) The secret version of the `connectivity_association_key_secret_id` secret in Vault.
		
			NOTE: Only the latest secret version will be used. 
		* `connectivity_association_name_secret_id` - (Required) (Updatable) Secret [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) containing the Connectivity association Key Name (CKN) of this MACsec key.
		* `connectivity_association_name_secret_version` - (Optional) (Updatable) The secret version of the `connectivity_association_name_secret_id` secret in Vault.

			NOTE: Only the latest secret version will be used. 
	* `state` - (Required) (Updatable) Indicates whether or not MACsec is enabled.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cross Connect Group
	* `update` - (Defaults to 20 minutes), when updating the Cross Connect Group
	* `delete` - (Defaults to 20 minutes), when destroying the Cross Connect Group


## Import

CrossConnectGroups can be imported using the `id`, e.g.

```
$ terraform import oci_core_cross_connect_group.test_cross_connect_group "id"
```

