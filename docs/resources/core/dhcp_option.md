# oci\_core\_dhcp\_option

Provide a Dhcp Options resource.

For more information, see 
[DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm)

## Example Usage

#### VCN Local with Internet 
```
resource "oci_core_dhcp_options" "dhcp-options1" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${var.vcn_ocid}"
  display_name = "dhcp-options1"
  
  // required
  options {
    type = "DomainNameServer"
    server_type = "VcnLocalPlusInternet"
  }
  
  // optional
  options {
    type = "SearchDomain"
    search_domain_names = [ "test.com" ]
  }
}
```

#### Custom DNS Server

```
resource "oci_core_dhcp_options" "dhcp-options2" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${var.vcn_ocid}"
  display_name = "dhcp-options3"
  
  // required
  options {
    type = "DomainNameServer"
    server_type = "CustomDnsServer"
    custom_dns_servers = [ "192.168.0.2", "192.168.0.11", "192.168.0.19" ]
  }
  
  // optional
  options {
    type = "SearchDomain"
    search_domain_names = [ "test.com" ]
  }
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vcn_id` - (Required) The OCID of the VCN.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.
* `options` - (Required) A set of [DHCP Options](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DhcpDnsOption/).

## Attributes Reference
* `compartment_id` - The OCID of the compartment containing the set of DHCP options.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - Oracle ID (OCID) for the set of DHCP options.
* `state` - The DRG's current state: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED].
* `options` - The collection of individual DHCP options.
* `time_created` - The date and time the image was created.
* `vcn_id` - (Required) The OCID of the VCN the set of DHCP options belongs to.
