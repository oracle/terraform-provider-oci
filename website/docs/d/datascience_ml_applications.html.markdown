---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_ml_applications"
sidebar_current: "docs-oci-datasource-datascience-ml_applications"
description: |-
  Provides the list of Ml Applications in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_ml_applications
This data source provides the list of Ml Applications in Oracle Cloud Infrastructure Data Science service.

Returns a list of MlApplications.


## Example Usage

```hcl
data "oci_datascience_ml_applications" "test_ml_applications" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.ml_application_compartment_id_in_subtree
	ml_application_id = oci_datascience_ml_application.test_ml_application.id
	name = var.ml_application_name
	state = var.ml_application_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) If it is true search must include all results from descendant compartments. Value true is allowed only if compartmentId refers to root compartment.
* `ml_application_id` - (Optional) unique MlApplication identifier
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `state` - (Optional) A filter to return only resources with lifecycleState matching the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `ml_application_collection` - The list of ml_application_collection.

### MlApplication Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment where the MlApplication is created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Optional description of the ML Application
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the MlApplication. Unique identifier that is immutable after creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `name` - The name of MlApplication. It is unique in a given tenancy.
* `state` - The current state of the MlApplication.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Creation time of MlApplication in the format defined by RFC 3339.
* `time_updated` - Time of last MlApplication update in the format defined by RFC 3339.

