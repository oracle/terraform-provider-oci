---
subcategory: "Resource Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resourcemanager_private_endpoint"
sidebar_current: "docs-oci-datasource-resourcemanager-private_endpoint"
description: |-
  Provides details about a specific Private Endpoint in Oracle Cloud Infrastructure Resource Manager service
---

# Data Source: oci_resourcemanager_private_endpoint
This data source provides details about a specific Private Endpoint resource in Oracle Cloud Infrastructure Resource Manager service.

Gets the specified private endpoint.
For more information, see
[Getting a Private Endpoint's Details](https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/get-private-endpoints.htm).


## Example Usage

```hcl
data "oci_resourcemanager_private_endpoint" "test_private_endpoint" {
	#Required
	private_endpoint_id = oci_resourcemanager_private_endpoint.test_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `private_endpoint_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this private endpoint.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the private endpoint. Avoid entering confidential information.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dns_zones` - DNS zones to use for accessing private Git servers. For private Git server instructions, see [Private Git Server](https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/private-endpoints.htm#private-git). Specify DNS fully qualified domain names (FQDNs); DNS Proxy forwards related DNS FQDN queries to the consumer DNS resolver. For DNS FQDNs not specified, queries go to service provider VCN resolver. Example: `abc.oraclevcn.com` 
* `freeform_tags` - Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
* `is_used_with_configuration_source_provider` - When `true`, allows the private endpoint to be used with a configuration source provider.
* `nsg_id_list` - The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of [network security groups (NSGs)](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm) for the private endpoint. Order does not matter. 
* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}` 
* `source_ips` - The source IP addresses that Resource Manager uses to connect to your network. Automatically assigned by Resource Manager.
* `state` - The current lifecycle state of the private endpoint. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet within the VCN for the private endpoint.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The date and time at which the private endpoint was created. Format is defined by RFC3339. Example: `2020-11-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN for the private endpoint.

