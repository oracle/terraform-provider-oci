---
subcategory: "Data Connectivity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_connectivity_registry"
sidebar_current: "docs-oci-datasource-data_connectivity-registry"
description: |-
  Provides details about a specific Registry in Oracle Cloud Infrastructure Data Connectivity service
---

# Data Source: oci_data_connectivity_registry
This data source provides details about a specific Registry resource in Oracle Cloud Infrastructure Data Connectivity service.

Gets a Data Connectivity Management Registry by identifier

## Example Usage

```hcl
data "oci_data_connectivity_registry" "test_registry" {
	#Required
	registry_id = oci_data_connectivity_registry.test_registry.id
}
```

## Argument Reference

The following arguments are supported:

* `registry_id` - (Required) The registry Ocid.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Registry description
* `display_name` - Data Connectivity Management Registry display name, registries can be renamed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `state` - Lifecycle states for registries in Data Connectivity Management Service CREATING - The resource is being created and may not be usable until the entire metadata is defined UPDATING - The resource is being updated and may not be usable until all changes are commited DELETING - The resource is being deleted and might require deep cleanup of children. ACTIVE   - The resource is valid and available for access INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for administrative reasons DELETED  - The resource has been deleted and isn't available FAILED   - The resource is in a failed state due to validation or other errors 
* `state_message` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `time_created` - The time the Data Connectivity Management Registry was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Data Connectivity Management Registry was updated. An RFC3339 formatted datetime string
* `updated_by` - Name of the user who updated the DCMS Registry.

