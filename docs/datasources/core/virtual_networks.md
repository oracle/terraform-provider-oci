# baremetal\_core\_virtual_networks

Gets a list of virtual networks.

## Example Usage

```
data "baremetal_core_virtual_networks" "t" {
      compartment_id = "compartment_id"
      limit = 1
      page = "page"
    }
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `offset` - (Optional) Offset of the snapshot data to retrieve.
* `length` - (Optional) Length of the snapshot data to retrieve.

## Attributes Reference

The following attributes are exported:

* `virtual_networks` - The list of virtual networks.

## Virtual Network reference
* `cidr_block` - The CIDR IP address block of the VCN.
* `compartment_id` - The OCID of the compartment containing the VCN.
* `default_dhcp_options_id` - The OCID for the VCN's default set of DHCP options.
* `default_route_table_id` - The OCID for the VCN's default route table.
* `default_security_list_id` - The OCID for the VCN's default security list.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The VCN's Oracle ID (OCID).
* `state` - The VCN's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the VCN was created.