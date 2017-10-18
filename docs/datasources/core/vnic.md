# oci\_core\_vnic

**API:** [Vnic Reference][38037a66]

  [38037a66]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vnic/ "VnicReference"

Gets information about a specific VNIC.

## Example Usage

```
data "oci_core_vnic" "t" {
  vnic_id = "vnicid"
}
```

## Argument Reference

The following arguments are supported:

* `vnic_id` - (Required) The OCID of the VNIC.

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `availability_domain` - The VNIC's Availability Domain.
* `display_name` - A user-friendly name. Does not have to be unique. Avoid entering confidential information.
* `id` - The OCID of the VNIC.
* `hostname_label` - The hostname for the VNIC that is created during instance launch. Used for DNS.
* `is_primary` - Whether the VNIC is the primary VNIC (the VNIC that is automatically created and attached during instance launch).
* `state` - The current state of the VNIC. Allowed values are: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `mac_address` - The MAC address of the VNIC.
* `private_ip_address` - The private IP addresses of the VNIC, which is within the VNIC subnet and is accessible within the VCN.
* `public_ip_address` - The public IP address of the VNIC, which Oracle performs NAT for at the gateway.
* `subnet_id` - The OCID of the subnet the VNIC is in.
* `time_created` - The date and time the VNIC was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. Defaults to `false`.
