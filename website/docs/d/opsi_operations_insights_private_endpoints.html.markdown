---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_operations_insights_private_endpoints"
sidebar_current: "docs-oci-datasource-opsi-operations_insights_private_endpoints"
description: |-
  Provides the list of Operations Insights Private Endpoints in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_operations_insights_private_endpoints
This data source provides the list of Operations Insights Private Endpoints in Oracle Cloud Infrastructure Opsi service.

Gets a list of Operation Insights private endpoints.

## Example Usage

```hcl
data "oci_opsi_operations_insights_private_endpoints" "test_operations_insights_private_endpoints" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.operations_insights_private_endpoint_compartment_id_in_subtree
	display_name = var.operations_insights_private_endpoint_display_name
	is_used_for_rac_dbs = var.operations_insights_private_endpoint_is_used_for_rac_dbs
	opsi_private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
	state = var.operations_insights_private_endpoint_state
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) A flag to search all resources within a given compartment and all sub-compartments. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name.
* `is_used_for_rac_dbs` - (Optional) The option to filter OPSI private endpoints that can used for RAC. Should be used along with vcnId query parameter.
* `opsi_private_endpoint_id` - (Optional) Unique Operations Insights PrivateEndpoint identifier
* `state` - (Optional) Lifecycle states
* `vcn_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


## Attributes Reference

The following attributes are exported:

* `operations_insights_private_endpoint_collection` - The list of operations_insights_private_endpoint_collection.

### OperationsInsightsPrivateEndpoint Reference

The following attributes are exported:

* `compartment_id` - The compartment OCID of the Private service accessed database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The description of the private endpoint.
* `display_name` - The display name of the private endpoint.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the Private service accessed database.
* `is_used_for_rac_dbs` - The flag is to identify if private endpoint is used for rac database or not. This flag is deprecated and no longer is used.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `nsg_ids` - The OCIDs of the network security groups that the private endpoint belongs to. 
* `private_endpoint_status_details` - A message describing the status of the private endpoint connection of this resource. For example, it can be used to provide actionable information about the validity of the private endpoint connection.
* `private_ip` - The private IP addresses assigned to the private endpoint. All IP addresses will be concatenated if it is RAC DBs. 
* `state` - The current state of the private endpoint.
* `subnet_id` - The OCID of the subnet.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcn_id` - The OCID of the VCN.

