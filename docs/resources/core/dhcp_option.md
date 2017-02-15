# baremetal\_core\_dhcp\_option

Provide a Dhcp Options resource.

## Example Usage

```
resource "baremetal_core_dhcp_options" "t" {
    compartment_id = "compartment_id"
    display_name = "display_name"
    options {
        type = "type"
        custom_dns_servers = [ "custom_dns_servers" ]
        server_type = "server_type"
    }
    options {
        type = "type"
        custom_dns_servers = [ "custom_dns_servers" ]
        server_type = "server_type"
    }
    vcn_id = "vcn_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vcn_id` - (Required) The OCID of the VCN.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.
* `options` - (Required) A set of DHCP options.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the set of DHCP options.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - Oracle ID (OCID) for the set of DHCP options.
* `state` - The DRG's current state: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED].
* `options` - The collection of individual DHCP options.
* `time_created` - The date and time the image was created.
* `vcn_id` - (Required) The OCID of the VCN the set of DHCP options belongs to.
