# oci\_core\_dhcp\_options

[DhcpOptions Reference][60fa58e0]

  [60fa58e0]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DhcpOptions/ "DhcpOptionsReference"

List Dhcp Options.

## Example Usage

```
data "oci_core_dhcp_options" "t" {
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
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - Oracle ID (OCID) for the set of DHCP options.
* `state` - The current state of DHCP options: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED].
* `options` - The collection of individual DHCP options.
* `time_created` - The date and time the set of DHCP options was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
* `vcn_id` - (Required) The OCID of the VCN that the set of DHCP options belongs to.
