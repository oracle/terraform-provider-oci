# baremetal\_core\_virtual_networks

Provides a virtual network resource.

## Example Usage

```
resource "baremetal_core_virtual_network" "t" {
			cidr_block = "cidr_block"
			compartment_id = "compartment_id"
			display_name = "display_name"
		}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `cidr_block` - (Required) a single, contiguous IPv4 CIDR block in the private IP address ranges specified in RFC 1918 (10.0.0.0/8, 172.16/12, and 192.168/16). Example: 172.16.0.0/16. The CIDR block can range from /16 to /30, and it must not overlap with your on-premise network. You can't change the size of the VCN after creation.
* `display_name` - (Optional) A human readable name. It need not be unique.

## Attributes Reference

The following attributes are exported:

* `cidr_block` - The CIDR IP address block of the VCN.
* `compartment_id` - The OCID of the compartment containing the VCN.
* `default_dhcp_options_id` - The OCID for the VCN's default set of DHCP options.
* `default_route_table_id` - The OCID for the VCN's default route table.
* `default_security_list_id` - The OCID for the VCN's default security list.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The VCN's Oracle ID (OCID).
* `state` - The VCN's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the VCN was created.