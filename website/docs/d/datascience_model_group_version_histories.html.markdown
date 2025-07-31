---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_group_version_histories"
sidebar_current: "docs-oci-datasource-datascience-model_group_version_histories"
description: |-
  Provides the list of Model Group Version Histories in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_group_version_histories
This data source provides the list of Model Group Version Histories in Oracle Cloud Infrastructure Data Science service.

List all modelGroupVersionHistories in the specified compartment. The query must include compartmentId.

## Example Usage

```hcl
data "oci_datascience_model_group_version_histories" "test_model_group_version_histories" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	created_by = var.model_group_version_history_created_by
	display_name = var.model_group_version_history_display_name
	id = var.model_group_version_history_id
	project_id = oci_datascience_project.test_project.id
	state = var.model_group_version_history_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `id` - (Optional) <b>Filter</b> results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type. 
* `project_id` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
* `state` - (Optional) A filter to return resources matching the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `model_group_version_histories` - The list of model_group_version_histories.

### ModelGroupVersionHistory Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the modelGroupVersionHistory's compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the modelGroupVersionHistory.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the modelGroupVersionHistory.
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the modelGroupVersionHistory.
* `latest_model_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the latest version of the model group associated.
* `lifecycle_details` - Details about the lifecycle state of the model group version history.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the modelGroupVersionHistory.
* `state` - The state of the modelGroupVersionHistory.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 
* `time_updated` - The date and time the resource was last updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

