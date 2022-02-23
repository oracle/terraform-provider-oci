---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec_connection_tunnel_error"
sidebar_current: "docs-oci-datasource-core-ipsec_connection_tunnel_error"
description: |-
  Provides details about a specific Ipsec Connection Tunnel Error in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ipsec_connection_tunnel_error
This data source provides details about a specific Ipsec Connection Tunnel Error resource in Oracle Cloud Infrastructure Core service.

Gets the identified error for the specified IPSec tunnel ID.


## Example Usage

```hcl
data "oci_core_ipsec_connection_tunnel_error" "test_ipsec_connection_tunnel_error" {
	#Required
	ipsec_id = oci_core_ipsec.test_ipsec.id
	tunnel_id = oci_core_tunnel.test_tunnel.id
}
```

## Argument Reference

The following arguments are supported:

* `ipsec_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPSec connection.
* `tunnel_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tunnel.


## Attributes Reference

The following attributes are exported:

* `error_code` - Unique code describes the error type.
* `error_description` - A detailed description of the error.
* `id` - Unique ID generated for each error report.
* `oci_resources_link` - Link to more Oracle resources or relevant documentation.
* `solution` - Resolution for the error.
* `timestamp` - Timestamp when the error occurred.

