# oci\_core\_route\_tables

Provide a route table resource.

## Example Usage

```
resource "oci_core_route_table" "t" {
    compartment_id = "compartment_id"
    display_name = "display_name"
    route_rules {
        cidr_block = "cidr_block"
        network_entity_id = "network_entity_id"
    }
    route_rules {
        cidr_block = "cidr_block"
        network_entity_id = "network_entity_id"
    }
    vcn_id = "vcn_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.
* `route_rules` - (Required) The collection of rules for routing destination IPs to network devices.
* `vcn_id` - (Required) The OCID of the VCN.

## Attributes reference

* `compartment_id` - The OCID of the compartment containing the route table.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The route table's Oracle Cloud ID (OCID).
* `state` - The route table's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `route_rules` - The collection of rules for routing destination IPs to network devices.
* `time_created` - The date and time the security list was created.
* `vcn_id` - The OCID of the VCN the security list belongs to.
