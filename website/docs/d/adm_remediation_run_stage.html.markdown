---
subcategory: "Adm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_adm_remediation_run_stage"
sidebar_current: "docs-oci-datasource-adm-remediation_run_stage"
description: |-
  Provides details about a specific Remediation Run Stage in Oracle Cloud Infrastructure Adm service
---

# Data Source: oci_adm_remediation_run_stage
This data source provides details about a specific Remediation Run Stage resource in Oracle Cloud Infrastructure Adm service.

Returns the details of the specified Remediation Run Stage.

## Example Usage

```hcl
data "oci_adm_remediation_run_stage" "test_remediation_run_stage" {
	#Required
	remediation_run_id = oci_adm_remediation_run.test_remediation_run.id
	stage_type = var.remediation_run_stage_stage_type
}
```

## Argument Reference

The following arguments are supported:

* `remediation_run_id` - (Required) Unique Remediation Run identifier path parameter.
* `stage_type` - (Required) The type of Remediation Run Stage, as a URL path parameter.


## Attributes Reference

The following attributes are exported:

* `audit_id` - The Oracle Cloud identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the vulnerability audit.
* `next_stage_type` - The next type of stage in the remediation run.
* `pipeline_properties` - Pipeline properties which result from the run of the verify stage.
	* `pipeline_identifier` - Unique identifier for the pipeline or action created in the Verify stage.
	* `pipeline_url` - The web link to the pipeline from the Verify stage.
* `previous_stage_type` - The previous type of stage in the remediation run.
* `pull_request_properties` - Pull request properties from recommend stage of the remediation run.
	* `pull_request_identifier` - Unique identifier for the pull or merge request created in the recommend stage.
	* `pull_request_url` - The web link to the pull or merge request created in the recommend stage.
* `recommended_updates_count` - Count of recommended application dependencies to update.
* `remediation_run_id` - The Oracle Cloud identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the remediation run.
* `status` - The current status of a remediation run stage.
* `summary` - Information about the current step within the stage.
* `time_created` - The creation date and time of the remediation run stage (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_finished` - The date and time of the finish of the remediation run stage (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_started` - The date and time of the start of the remediation run stage (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `type` - The type of the remediation run stage.

