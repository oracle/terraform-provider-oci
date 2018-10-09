---
layout: "oci"
page_title: "Provider: Oracle Cloud Infrastructure"
sidebar_current: "docs-oci-guide-managing_default_resources"
description: |-
  The Oracle Cloud Infrastructure provider. Managing Default Resources
---
## Managing Default Virtual Cloud Network Resources

When you create an [oci_core_vcn](../r/core_vcn.html)
resource, it will also create the following associated resources by default.

- [oci_core_security_list](../r/core_security_list.html)
- [oci_core_dhcp_options](../r/core_dhcp_options.html)
- [oci_core_route_table](../r/core_route_table.html)

These default resources will be implicitly created even if they are not specified in the Terraform configuration.
Their OCIDs are returned by the following attributes under the `oci_core_vcn` resource:

- `default_security_list_id`
- `default_dhcp_options_id`
- `default_route_table_id`

Default resources must be configured in Terraform using a separate resource type. Here are
the mappings between the resource and the new resource type to use for configuring default
resources:
- `oci_core_security_list` => `oci_core_default_security_list`
- `oci_core_dhcp_options` => `oci_core_default_dhcp_options`
- `oci_core_route_table` => `oci_core_default_route_table`

Default resources types are configured in the same way as their non-default counterparts. 
The only difference is specifying the ID of the default resource using the
`manage_default_resource_id` argument.

Consequently, the `compartment_id` and `vcn_id` are no longer necessary for default resources.


### Example Usage
#### Modifying a VCN's default DHCP options

```
resource "oci_core_vcn" "vcn1" {
  cidr_block = "10.0.0.0/16"
  dns_label = "vcn1"
  compartment_id = "${var.compartment_ocid}"
  display_name = "vcn1"
}

resource "oci_core_default_dhcp_options" "default-dhcp-options" {
  manage_default_resource_id = "${oci_core_vcn.vcn1.default_dhcp_options_id}"

  // required
  options {
    type = "DomainNameServer"
    server_type = "VcnLocalPlusInternet"
  }

  // optional
  options {
    type = "SearchDomain"
    search_domain_names = [ "abc.com" ]
  }
}
```

### Limitations

Default resources can only be removed when the associated `oci_core_vcn resource` is removed. When attempting
a targeted removal of a default resource, the resource will be removed from the Terraform state file but the resource may
still exist in OCI with empty settings.
 
Examples of targeted removal include:

- Removing a default resource from a Terraform configuration that was previously applied
- Running a `terraform destroy -target=<default resource>` command
- Changing the `manage_default_resource_id` for a default resource that was previously applied