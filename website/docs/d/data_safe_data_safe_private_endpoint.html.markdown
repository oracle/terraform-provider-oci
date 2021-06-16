---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_data_safe_private_endpoint"
sidebar_current: "docs-oci-datasource-data_safe-data_safe_private_endpoint"
description: |-
  Provides details about a specific Data Safe Private Endpoint in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_data_safe_private_endpoint
This data source provides details about a specific Data Safe Private Endpoint resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified Data Safe private endpoint.

## Example Usage

```hcl
data "oci_data_safe_data_safe_private_endpoint" "test_data_safe_private_endpoint" {
	#Required
	data_safe_private_endpoint_id = oci_data_safe_data_safe_private_endpoint.test_data_safe_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `data_safe_private_endpoint_id` - (Required) The OCID of the private endpoint.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the private endpoint.
* `display_name` - The display name of the private endpoint.
* `endpoint_fqdn` - The three-label fully qualified domain name (FQDN) of the private endpoint. The customer VCN's DNS records are updated with this FQDN.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Data Safe private endpoint.
* `nsg_ids` - The OCIDs of the network security groups that the private endpoint belongs to. 
* `private_endpoint_id` - The OCID of the underlying private endpoint.
* `private_endpoint_ip` - The private IP address of the private endpoint. 
* `state` - The current state of the private endpoint.
* `subnet_id` - The OCID of the subnet.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcn_id` - The OCID of the VCN.

