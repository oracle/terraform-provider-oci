# oci_core_vnic

## Vnic Singular DataSource

### Vnic Reference

The following attributes are exported:

* `availability_domain` - The VNIC's Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment containing the VNIC.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).  For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).  Example: `bminstance-1` 
* `id` - The OCID of the VNIC.
* `is_primary` - Whether the VNIC is the primary VNIC (the VNIC that is automatically created and attached during instance launch). 
* `mac_address` - The MAC address of the VNIC.  Example: `00:00:17:B6:4D:DD` 
* `private_ip_address` - The private IP address of the primary `privateIp` object on the VNIC. The address is within the CIDR of the VNIC's subnet.  Example: `10.0.3.3` 
* `public_ip_address` - The public IP address of the VNIC, if one is assigned. 
* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you would skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm#privateip).  Example: `true` 
* `state` - The current state of the VNIC.
* `subnet_id` - The OCID of the subnet the VNIC is in.
* `time_created` - The date and time the VNIC was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Get Operation
Gets the information for the specified virtual network interface card (VNIC).
You can get the VNIC OCID from the
[ListVnicAttachments](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VnicAttachment/ListVnicAttachments)
operation.


The following arguments are supported:

* `vnic_id` - (Required) The OCID of the VNIC.


### Example Usage

```hcl
data "oci_core_vnic" "test_vnic" {
	#Required
	vnic_id = "${oci_core_vnic.test_vnic.id}"
}
```
