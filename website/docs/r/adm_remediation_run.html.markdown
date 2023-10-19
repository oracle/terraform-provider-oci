---
subcategory: "Adm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_adm_remediation_run"
sidebar_current: "docs-oci-resource-adm-remediation_run"
description: |-
  Provides the Remediation Run resource in Oracle Cloud Infrastructure Adm service
---

# oci_adm_remediation_run
This resource provides the Remediation Run resource in Oracle Cloud Infrastructure Adm service.

Creates a new remediation run.

## Example Usage

```hcl
resource "oci_adm_remediation_run" "test_remediation_run" {
	#Required
	remediation_recipe_id = oci_adm_remediation_recipe.test_remediation_recipe.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.remediation_run_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The compartment Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the remediation run.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) The name of the remediation run.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `remediation_recipe_id` - (Required) The Oracle Cloud identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Remediation Recipe.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the remediation run.
* `current_stage_type` - The type of the current stage of the remediation run.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The name of the remediation run.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the remediation run.
* `remediation_recipe_id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Remediation Recipe.
* `remediation_run_source` - The source that triggered the Remediation Recipe.
* `stages` - The list of remediation run stage summaries.
	* `summary` - Information about the current step within the given stage.
	* `time_created` - The creation date and time of the stage (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
	* `time_finished` - The date and time of the finish of the stage (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
	* `time_started` - The date and time of the start of the stage (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
	* `type` - The type of stage.
* `state` - The current lifecycle state of the remediation run.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The creation date and time of the remediation run (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_finished` - The date and time of the finish of the remediation run (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_started` - The date and time of the start of the remediation run (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_updated` - The date and time the remediation run was last updated (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Remediation Run
	* `update` - (Defaults to 20 minutes), when updating the Remediation Run
	* `delete` - (Defaults to 20 minutes), when destroying the Remediation Run


## Import

RemediationRuns can be imported using the `id`, e.g.

```
$ terraform import oci_adm_remediation_run.test_remediation_run "id"
```

