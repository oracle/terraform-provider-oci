---
subcategory: "Resource Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resourcemanager_private_endpoints"
sidebar_current: "docs-oci-datasource-resourcemanager-private_endpoints"
description: |-
  Provides the list of Private Endpoints in Oracle Cloud Infrastructure Resource Manager service
---

# Data Source: oci_resourcemanager_private_endpoints
This data source provides the list of Private Endpoints in Oracle Cloud Infrastructure Resource Manager service.

Lists private endpoints according to the specified filter.
- For `compartmentId`, lists all private endpoint in the matching compartment.
- For `privateEndpointId`, lists the matching private endpoint.


## Example Usage

```hcl
data "oci_resourcemanager_private_endpoints" "test_private_endpoints" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.private_endpoint_display_name
	private_endpoint_id = oci_resourcemanager_private_endpoint.test_private_endpoint.id
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) A filter to return only resources that exist in the compartment, identified by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. Use this filter to list a resource by name. Requires `sortBy` set to `DISPLAYNAME`. Alternatively, when you know the resource OCID, use the related Get operation. 
* `private_endpoint_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint. 
* `vcn_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


## Attributes Reference

The following attributes are exported:

* `private_endpoint_collection` - The list of private_endpoint_collection.

### PrivateEndpoint Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this private endpoint details.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the private endpoint. Avoid entering confidential information.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `dns_zones` - DNS Proxy forwards any DNS FQDN queries over into the consumer DNS resolver if the DNS FQDN is included in the dns zones list otherwise it goes to service provider VCN resolver. 
* `freeform_tags` - Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the private endpoint details.
* `is_used_with_configuration_source_provider` - When `true`, allows the private endpoint to be used with a configuration source provider.
* `nsg_id_list` - An array of network security groups (NSG) that the customer can optionally provide.
* `source_ips` - The source IPs which resource manager service will use to connect to customer's network. Automatically assigned by Resource Manager Service.
* `state` - The current lifecycle state of the private endpoint. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet within the VCN for the private endpoint.
* `time_created` - The date and time at which the private endpoint was created. Format is defined by RFC3339. Example: `2020-11-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN for the private endpoint.

