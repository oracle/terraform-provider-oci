# oci\_core\_route\_tables

**API:** [RouteTable Reference][c78b8cc0]

  [c78b8cc0]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/RouteTable/ "RouteTableReference"

Gets a list of route tables. A route table is a collection of `RouteRule` objects, which are used to route packets based on destination IP to a particular network entity.

## Example Usage

```
    data "oci_core_route_tables" "t" {
      compartment_id = "compartment_id"
      vcn_id = "vcn_id"
    }
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vcn_id` - (Required) The OCID of the VCN.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) Length of the snapshot data to retrieve.

## Attributes Reference

The following attributes are exported:

* `route_tables` - The list of security lists.

## Route Table reference
* `compartment_id` - The OCID of the compartment containing the route table.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The route table's Oracle Cloud ID (OCID).
* `state` - The route table's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `route_rules` - The collection of rules for routing destination IPs to network devices.
* `time_created` - The date and time the route table was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `vcn_id` - The OCID of the VCN the security list belongs to.
