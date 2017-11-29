# oci\_core\_virtual_networks

[Vcn Reference][db318935]

  [db318935]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vcn/ "VcnReference"

Provides a Virtual Cloud Network (VCN) resource.

VCN resources have a default set of DHCP options, security list, and route table.
To learn more about managing these resources, see [Managing Default Virtual Cloud Network Resources](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Managing%20Default%20Resources.md).

## Example Usage

```
resource "oci_core_virtual_network" "t" {
    cidr_block = "cidr_block"
    compartment_id = "compartment_id"
    display_name = "display_name"
}
```

## Argument Reference

The following arguments are supported:

* `cidr_block` - (Required) The CIDR IP address block of the VCN.
* `compartment_id` - (Required) The OCID of the compartment to contain the VCN.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `dns_label` - (Optional) A DNS label for the VCN.

## Attributes Reference
* `compartment_id` - The OCID of the compartment.
* `cidr_block` - The CIDR IP address block of the VCN.
* `default_dhcp_options_id` - The OCID for the VCN's default set of DHCP options.
* `default_route_table_id` - The OCID for the VCN's default route table.
* `default_security_list_id` - The OCID for the VCN's default security list.
* `display_name` - A user-friendly name. Does not have to be unique.  Avoid entering confidential information.
* `id` - The OCID of the VCN.
* `state` - The current state of the VCN. Allowed values are: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the VCN was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
