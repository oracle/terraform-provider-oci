---
subcategory: "Psa"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psa_private_service_access"
sidebar_current: "docs-oci-datasource-psa-private_service_access"
description: |-
  Provides details about a specific Private Service Access in Oracle Cloud Infrastructure Psa service
---

# Data Source: oci_psa_private_service_access
This data source provides details about a specific Private Service Access resource in Oracle Cloud Infrastructure Psa service.

Gets information about a PrivateServiceAccess.

## Example Usage

```hcl
data "oci_psa_private_service_access" "test_private_service_access" {
	#Required
	private_service_access_id = oci_psa_private_service_access.test_private_service_access.id
}
```

## Argument Reference

The following arguments are supported:

* `private_service_access_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PrivateServiceAccess.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the private service access. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of this private service access. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `fqdns` - The private service access FQDNs, which are going to be used to access the service.  Example: `xyz.oraclecloud.com` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private service access. 
* `ipv4ip` - The private IPv4 address (in the consumer's VCN) that represents the access point for the associated service. 
* `nsg_ids` - A list of the OCIDs of the network security groups that the private service access's VNIC belongs to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
* `service_id` - A unique service identifier for which the private service access was created. 
* `state` - The private service access's current lifecycle state.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet that the private service access belongs to. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the private service access was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the PrivateServiceAccess was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN that the private service access belongs to. 
* `vnic_id` - An [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private service access's VNIC, which resides in the private service access's VCN . 

