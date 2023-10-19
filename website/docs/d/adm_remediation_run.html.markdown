---
subcategory: "Adm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_adm_remediation_run"
sidebar_current: "docs-oci-datasource-adm-remediation_run"
description: |-
  Provides details about a specific Remediation Run in Oracle Cloud Infrastructure Adm service
---

# Data Source: oci_adm_remediation_run
This data source provides details about a specific Remediation Run resource in Oracle Cloud Infrastructure Adm service.

Returns the details of the specified remediation run.

## Example Usage

```hcl
data "oci_adm_remediation_run" "test_remediation_run" {
	#Required
	remediation_run_id = oci_adm_remediation_run.test_remediation_run.id
}
```

## Argument Reference

The following arguments are supported:

* `remediation_run_id` - (Required) Unique Remediation Run identifier path parameter.


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

