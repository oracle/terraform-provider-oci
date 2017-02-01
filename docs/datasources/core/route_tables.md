# baremetal\_core\_route\_tables

Gets a list of route tables.

## Example Usage

```
    data "baremetal_core_route_tables" "t" {
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
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The route table's Oracle Cloud ID (OCID).
* `state` - The route table's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `route_rules` - The collection of rules for routing destination IPs to network devices.
* `time_created` - The date and time the security list was created.
* `vcn_id` - The OCID of the VCN the security list belongs to.
