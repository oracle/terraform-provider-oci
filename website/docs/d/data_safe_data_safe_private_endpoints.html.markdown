---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_data_safe_private_endpoints"
sidebar_current: "docs-oci-datasource-data_safe-data_safe_private_endpoints"
description: |-
  Provides the list of Data Safe Private Endpoints in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_data_safe_private_endpoints
This data source provides the list of Data Safe Private Endpoints in Oracle Cloud Infrastructure Data Safe service.

Gets a list of Data Safe private endpoints.


## Example Usage

```hcl
data "oci_data_safe_data_safe_private_endpoints" "test_data_safe_private_endpoints" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.data_safe_private_endpoint_access_level
	compartment_id_in_subtree = var.data_safe_private_endpoint_compartment_id_in_subtree
	display_name = var.data_safe_private_endpoint_display_name
	state = var.data_safe_private_endpoint_state
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state.
* `vcn_id` - (Optional) A filter to return only resources that match the specified VCN OCID.


## Attributes Reference

The following attributes are exported:

* `data_safe_private_endpoints` - The list of data_safe_private_endpoints.

### DataSafePrivateEndpoint Reference

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

