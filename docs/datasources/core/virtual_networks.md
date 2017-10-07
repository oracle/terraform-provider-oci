# oci\_core\_virtual\_networks

**API:** [Vcn Reference][0d11fda6]

  [0d11fda6]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vcn/ "VcnReference"

Gets a list of Virtual Cloud Networks (VCNs).

## Example Usage

```
data "oci_core_virtual_networks" "t" {
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
* `display_name` - A user-friendly name. Does not have to be unique. Avoid entering confidential information.
* `id` - The OCID of the VNIC.
* `state` - The current state of the VCN. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the VCN was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
