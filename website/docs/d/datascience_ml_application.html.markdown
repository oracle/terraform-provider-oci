---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_ml_application"
sidebar_current: "docs-oci-datasource-datascience-ml_application"
description: |-
  Provides details about a specific Ml Application in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_ml_application
This data source provides details about a specific Ml Application resource in Oracle Cloud Infrastructure Data Science service.

Gets a MlApplication by identifier

## Example Usage

```hcl
data "oci_datascience_ml_application" "test_ml_application" {
	#Required
	ml_application_id = oci_datascience_ml_application.test_ml_application.id
}
```

## Argument Reference

The following arguments are supported:

* `ml_application_id` - (Required) unique MlApplication identifier


## Attributes Reference

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

