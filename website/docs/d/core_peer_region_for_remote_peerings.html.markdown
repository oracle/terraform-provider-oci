---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_peer_region_for_remote_peerings"
sidebar_current: "docs-oci-datasource-core-peer_region_for_remote_peerings"
description: |-
  Provides the list of Peer Region For Remote Peerings in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_peer_region_for_remote_peerings
This data source provides the list of Peer Region For Remote Peerings in Oracle Cloud Infrastructure Core service.

Lists the regions that support remote VCN peering (which is peering across regions).
For more information, see [VCN Peering](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/VCNpeering.htm).


## Example Usage

```hcl
data "oci_core_peer_region_for_remote_peerings" "test_peer_region_for_remote_peerings" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `peer_region_for_remote_peerings` - The list of peer_region_for_remote_peerings.

### PeerRegionForRemotePeering Reference

The following attributes are exported:

* `name` - The region's name.  Example: `us-phoenix-1` 

