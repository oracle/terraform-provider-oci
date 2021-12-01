---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec_connections"
sidebar_current: "docs-oci-datasource-core-ipsec_connections"
description: |-
  Provides the list of Ip Sec Connections in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ipsec_connections
This data source provides the list of Ip Sec Connections in Oracle Cloud Infrastructure Core service.

Lists the IPSec connections for the specified compartment. You can filter the
results by DRG or CPE.


## Example Usage

```hcl
data "oci_core_ipsec_connections" "test_ip_sec_connections" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	cpe_id = oci_core_cpe.test_cpe.id
	drg_id = oci_core_drg.test_drg.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpe_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CPE.
* `drg_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.


## Attributes Reference

The following attributes are exported:

* `connections` - The list of connections.

### IpSecConnection Reference

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

	Used for routing a given IPSec tunnel's traffic only if the tunnel is using static routing. If you configure at least one tunnel to use static routing, then you must provide at least one valid static route. If you configure both tunnels to use BGP dynamic routing, you can provide an empty list for the static routes.

	 Example: `10.0.1.0/24` 
* `time_created` - The date and time the IPSec connection was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

