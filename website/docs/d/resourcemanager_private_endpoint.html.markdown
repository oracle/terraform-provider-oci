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

