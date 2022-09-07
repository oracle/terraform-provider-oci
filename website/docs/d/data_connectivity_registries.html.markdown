---
subcategory: "Data Connectivity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_connectivity_registries"
sidebar_current: "docs-oci-datasource-data_connectivity-registries"
description: |-
  Provides the list of Registries in Oracle Cloud Infrastructure Data Connectivity service
---

# Data Source: oci_data_connectivity_registries
This data source provides the list of Registries in Oracle Cloud Infrastructure Data Connectivity service.

Retrieves a list of Data Connectivity Management registries.


## Example Usage

```hcl
data "oci_data_connectivity_registries" "test_registries" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	is_deep_lookup = var.registry_is_deep_lookup
	name = var.registry_name
	state = var.registry_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the resources you want to list.
* `is_deep_lookup` - (Optional) This parameter allows list registries to deep look at the whole tenancy.
* `name` - (Optional) Used to filter by the name of the object.
* `state` - (Optional) Lifecycle state of the resource.


## Attributes Reference

The following attributes are exported:

* `registry_summary_collection` - The list of registry_summary_collection.

### Registry Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Registry description
* `display_name` - Data Connectivity Management registry display name; registries can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
* `id` - A unique identifier that is immutable on creation.
* `state` - Lifecycle states for registries in the Data Connectivity Management Service CREATING - The resource is being created and may not be usable until the entire metadata is defined. UPDATING - The resource is being updated and may not be usable until all changes are commited. DELETING - The resource is being deleted and might require deep cleanup of children. ACTIVE   - The resource is valid and available for access. INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for administrative reasons. DELETED  - The resource has been deleted and isn't available. FAILED   - The resource is in a failed state due to validation or other errors. 
* `state_message` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `time_created` - Time when the Data Connectivity Management registry was created. An RFC3339 formatted datetime string.
* `time_updated` - Time when the Data Connectivity Management registry was updated. An RFC3339 formatted datetime string.
* `updated_by` - Name of the user who updated the DCMS registry.

