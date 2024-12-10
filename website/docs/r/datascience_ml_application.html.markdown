---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_ml_application"
sidebar_current: "docs-oci-resource-datascience-ml_application"
description: |-
  Provides the Ml Application resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_ml_application
This resource provides the Ml Application resource in Oracle Cloud Infrastructure Data Science service.

Creates a new MlApplication.


## Example Usage

```hcl
resource "oci_datascience_ml_application" "test_ml_application" {
	#Required
	compartment_id = var.compartment_id
	name = var.ml_application_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.ml_application_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the MlApplication is created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Optional description of the ML Application
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) The name of MlApplication. It is unique in a given tenancy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ml Application
	* `update` - (Defaults to 20 minutes), when updating the Ml Application
	* `delete` - (Defaults to 20 minutes), when destroying the Ml Application


## Import

MlApplications can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_ml_application.test_ml_application "id"
```

