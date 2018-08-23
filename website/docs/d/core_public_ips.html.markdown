---
layout: "oci"
page_title: "OCI: oci_core_public_ips"
sidebar_current: "docs-oci-datasource-core-public_ips"
description: |-
  Provides a list of PublicIps
---

# Data Source: oci_core_public_ips
The `oci_core_public_ips` data source allows access to the list of OCI public_ips

Lists either the ephemeral or reserved [PublicIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PublicIp/) objects
in the specified compartment.

To list your reserved public IPs, set `scope` = `REGION`, and leave the
`availabilityDomain` parameter empty.

To list your ephemeral public IPs, set `scope` = `AVAILABILITY_DOMAIN`, and set the
`availabilityDomain` parameter to the desired Availability Domain. An ephemeral public IP
is always in the same Availability Domain and compartment as the private IP it's assigned to.


## Example Usage

```hcl
data "oci_core_public_ips" "test_public_ips" {
	#Required
	compartment_id = "${var.compartment_id}"
	scope = "${var.public_ip_scope}"

	#Optional
	availability_domain = "${var.public_ip_availability_domain}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `scope` - (Required) Whether the public IP is regional or specific to a particular Availability Domain.
	* `REGION`: The public IP exists within a region and can be assigned to a private IP in any Availability Domain in the region. Reserved public IPs have `scope` = `REGION`.
	* `AVAILABILITY_DOMAIN`: The public IP exists within the Availability Domain of the private IP it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs have `scope` = `AVAILABILITY_DOMAIN`. 


## Attributes Reference

The following attributes are exported:

* `public_ips` - The list of public_ips.

### PublicIp Reference

The following attributes are exported:

* `availability_domain` - The public IP's Availability Domain. This property is set only for ephemeral public IPs (that is, when the `scope` of the public IP is set to AVAILABILITY_DOMAIN). The value is the Availability Domain of the assigned private IP.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment containing the public IP. For an ephemeral public IP, this is the same compartment as the private IP's. For a reserved public IP that is currently assigned, this can be a different compartment than the assigned private IP's. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The public IP's Oracle ID (OCID).
* `ip_address` - The public IP address of the `publicIp` object.  Example: `129.146.2.1` 
* `lifetime` - Defines when the public IP is deleted and released back to Oracle's public IP pool.
	* `EPHEMERAL`: The lifetime is tied to the lifetime of its assigned private IP. The ephemeral public IP is automatically deleted when its private IP is deleted, when the VNIC is terminated, or when the instance is terminated. An ephemeral public IP must always be assigned to a private IP.
	* `RESERVED`: You control the public IP's lifetime. You can delete a reserved public IP whenever you like. It does not need to be assigned to a private IP at all times.

	For more information and comparison of the two types, see [Public IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingpublicIPs.htm). 
* `private_ip_id` - The OCID of the private IP that the public IP is currently assigned to, or in the process of being assigned to. 
* `scope` - Whether the public IP is regional or specific to a particular Availability Domain.
	* `REGION`: The public IP exists within a region and can be assigned to a private IP in any Availability Domain in the region. Reserved public IPs have `scope` = `REGION`.
	* `AVAILABILITY_DOMAIN`: The public IP exists within the Availability Domain of the private IP it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs have `scope` = `AVAILABILITY_DOMAIN`. 
* `state` - The public IP's current state.
* `time_created` - The date and time the public IP was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

