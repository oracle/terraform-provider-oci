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

Lists the tunnel Security Associations information for the specified IPSec Tunnel ID.


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

* `cpe_subnet` - IP and mask of the Partner Subnet for Policy Based VPNs or Static Routes
* `oracle_subnet` - IP and mask of the Local Subnet for Policy Based VPNs or Static Routes
* `time` - Seconds in current state
* `tunnel_sa_error_info` - Current state if status is not up, including phase1/phase2 and possible reason for tunnel not up
* `tunnel_sa_status` - Phase 1 Status of the Tunnel

