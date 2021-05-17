---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_on_prem_connectors"
sidebar_current: "docs-oci-datasource-data_safe-on_prem_connectors"
description: |-
  Provides the list of On Prem Connectors in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_on_prem_connectors
This data source provides the list of On Prem Connectors in Oracle Cloud Infrastructure Data Safe service.

Gets a list of on-premises connectors.


## Example Usage

```hcl
data "oci_data_safe_on_prem_connectors" "test_on_prem_connectors" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.on_prem_connector_access_level
	compartment_id_in_subtree = var.on_prem_connector_compartment_id_in_subtree
	display_name = var.on_prem_connector_display_name
	on_prem_connector_id = oci_data_safe_on_prem_connector.test_on_prem_connector.id
	on_prem_connector_lifecycle_state = var.on_prem_connector_on_prem_connector_lifecycle_state
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `on_prem_connector_id` - (Optional) A filter to return only the on-premises connector that matches the specified id.
* `on_prem_connector_lifecycle_state` - (Optional) A filter to return only on-premises connector resources that match the specified lifecycle state.


## Attributes Reference

The following attributes are exported:

* `on_prem_connectors` - The list of on_prem_connectors.

### OnPremConnector Reference

The following attributes are exported:

* `available_version` - Latest available version of the on-premises connector.
* `compartment_id` - The OCID of the compartment that contains the on-premises connector.
* `created_version` - Created version of the on-premises connector.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the on-premises connector.
* `display_name` - The display name of the on-premises connector.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the on-premises connector.
* `lifecycle_details` - Details about the current state of the on-premises connector.
* `state` - The current state of the on-premises connector.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the on-premises connector was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

