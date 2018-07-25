---
layout: "oci"
page_title: "OCI: oci_core_peer_region_for_remote_peerings"
sidebar_current: "docs-oci-datasource-core-peer_region_for_remote_peerings"
description: |-
  Provides a list of PeerRegionForRemotePeerings
---

# Data Source: oci_core_peer_region_for_remote_peerings
The PeerRegionForRemotePeerings data source allows access to the list of OCI peer_region_for_remote_peerings

Lists the regions that support remote VCN peering (which is peering across regions).
For more information, see [VCN Peering](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/VCNpeering.htm).


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

