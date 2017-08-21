# baremetal\_core\_vnic

Gets information about a specific vnic.

## Example Usage

```
data "baremetal_core_vnic" "t" {
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
* `display_name` - A user-friendly name. Does not have to be unique.
* `id` - The OCID of the VNIC.
* `hostname_label` - The hostname for the VNIC that is created during instance launch. Used for DNS.
* `is_primary` - Whether the VNIC is the primary VNIC (the VNIC that is automatically created and attached during instance launch).
* `state` - The current state of the VNIC. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `mac_address` - The MAC address of the VNIC.
* `private_ip_address` - The private IP addresses of the VNIC, which is within the VNIC subnet and is accessible within the VCN.
* `public_ip_address` - The public IP address of the VNIC, which Oracle performs NAT for at the gateway.
* `subnet_id` - The OCID of the subnet the VNIC is in.
* `time_created` - The date and time the VNIC was created.
