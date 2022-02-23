---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_tunnel_security_associations"
sidebar_current: "docs-oci-datasource-core-tunnel_security_associations"
description: |-
  Provides the list of Tunnel Security Associations in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_tunnel_security_associations
This data source provides the list of Tunnel Security Associations in Oracle Cloud Infrastructure Core service.

Lists the tunnel security associations information for the specified IPSec tunnel ID.


## Example Usage

```hcl
data "oci_core_tunnel_security_associations" "test_tunnel_security_associations" {
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

* `tunnel_security_associations` - The list of tunnel_security_associations.

### TunnelSecurityAssociation Reference

The following attributes are exported:

* `cpe_subnet` - The IP address and mask of the partner subnet used in policy based VPNs or static routes.
* `oracle_subnet` - The IP address and mask of the local subnet used in policy based VPNs or static routes.
* `time` - Time in the current state, in seconds.
* `tunnel_sa_error_info` - Current state if the IPSec tunnel status is not `UP`, including phase one and phase two details and a possible reason the tunnel is not `UP`. 
* `tunnel_sa_status` - The IPSec tunnel's phase one status.

