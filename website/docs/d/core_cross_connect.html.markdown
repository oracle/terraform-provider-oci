---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect"
sidebar_current: "docs-oci-datasource-core-cross_connect"
description: |-
  Provides details about a specific Cross Connect in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cross_connect
This data source provides details about a specific Cross Connect resource in Oracle Cloud Infrastructure Core service.

Gets the specified cross-connect's information.

## Example Usage

```hcl
data "oci_core_cross_connect" "test_cross_connect" {
	#Required
	cross_connect_id = oci_core_cross_connect.test_cross_connect.id
}
```

## Argument Reference

The following arguments are supported:

* `cross_connect_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the cross-connect group.
* `cross_connect_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect group this cross-connect belongs to (if any). 
* `customer_reference_name` - A reference name or identifier for the physical fiber connection that this cross-connect uses. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The cross-connect's Oracle ID (OCID).
* `location_name` - The name of the FastConnect location where this cross-connect is installed. 
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
* `port_name` - A string identifying the meet-me room port for this cross-connect.
* `port_speed_shape_name` - The port speed for this cross-connect.  Example: `10 Gbps` 
* `state` - The cross-connect's current state.
* `time_created` - The date and time the cross-connect was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

