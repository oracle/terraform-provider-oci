---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ip_sec_connections"
sidebar_current: "docs-oci-datasource-core-ipsecs"
description: |-
  Provides the list of Ip Sec Connections in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ip_sec_connections
This data source provides the list of Ip Sec Connections in Oracle Cloud Infrastructure Core service.

Lists the IPSec connections for the specified compartment. You can filter the
results by DRG or CPE.


## Example Usage

```hcl
data "oci_core_ipsec_connections" "test_ip_sec_connections" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	cpe_id = "${oci_core_cpe.test_cpe.id}"
	drg_id = "${oci_core_drg.test_drg.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `cpe_id` - (Optional) The OCID of the CPE.
* `drg_id` - (Optional) The OCID of the DRG.


## Attributes Reference

The following attributes are exported:

* `connections` - The list of IPSec connections.

### IpSecConnection Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the IPSec connection.
* `cpe_id` - The OCID of the CPE.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The OCID of the DRG.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The IPSec connection's Oracle ID (OCID).
* `state` - The IPSec connection's current state.
* `static_routes` - Static routes to the CPE. At least one route must be included. The CIDR must not be a multicast address or class E address.  Example: `10.0.1.0/24` 
* `time_created` - The date and time the IPSec connection was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

