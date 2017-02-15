# baremetal\_core\_dhcp\_options

List Dhcp Options.

## Example Usage

```
data "baremetal_core_dhcp_options" "t" {
  compartment_id = "compartment_id"
  limit = 1
  page = "page"
  vcn_id = "vcn_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vcn_id` - (Required) The OCID of the VCN.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The page to fetch.

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the set of DHCP options.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - Oracle ID (OCID) for the set of DHCP options.
* `state` - The DRG's current state: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED].
* `options` - The collection of individual DHCP options.
* `time_created` - The date and time the image was created.
* `vcn_id` - (Required) The OCID of the VCN the set of DHCP options belongs to.
