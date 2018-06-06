
# oci_core_peer_region_for_remote_peerings

## PeerRegionForRemotePeering DataSource

Gets a list of peer_region_for_remote_peerings.

### List Operation
Lists the regions that support remote VCN peering (which is peering across regions).
For more information, see [VCN Peering](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/VCNpeering.htm).

The following arguments are supported:



The following attributes are exported:

* `peer_region_for_remote_peerings` - The list of peer_region_for_remote_peerings.

### Example Usage

```hcl
data "oci_core_peer_region_for_remote_peerings" "test_peer_region_for_remote_peerings" {
}
```
### PeerRegionForRemotePeering Reference

The following attributes are exported:

* `name` - The region's name.  Example: `us-phoenix-1` 
