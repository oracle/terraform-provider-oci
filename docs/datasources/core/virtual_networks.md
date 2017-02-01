
# baremetal\_core\_virtual\_networks

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
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The pagination token to continue listing from.

## Attributes Reference

The following attributes are exported:

* `virtual_networks` - The list of virtual networks.

## Virtual Networks Reference
* `compartment_id` - The OCID of the compartment.
* `cidr_block` - The CIDR IP address block of the VCN.
* `default_dhcp_options_id` - The OCID for the VCN's default set of DHCP options.
* `default_route_table_id` - The OCID for the VCN's default route table.
* `default_security_list_id` - The OCID for the VCN's default security list.
* `display_name` - A user-friendly name. Does not have to be unique.
* `id` - The OCID of the VNIC.
* `state` - The current state of the VNIC. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the VNIC was created.
