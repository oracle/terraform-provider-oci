---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_deployment_disaster_recovery_precheck_report"
sidebar_current: "docs-oci-datasource-golden_gate-deployment_disaster_recovery_precheck_report"
description: |-
  Provides details about a specific Deployment Disaster Recovery Precheck Report in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_deployment_disaster_recovery_precheck_report
This data source provides details about a specific Deployment Disaster Recovery Precheck Report resource in Oracle Cloud Infrastructure Golden Gate service.

Returns DR precheck report for a standby peer with the specified placement (availabilityDomain and faultDomain).


## Example Usage

```hcl
data "oci_golden_gate_deployment_disaster_recovery_precheck_report" "test_deployment_disaster_recovery_precheck_report" {
	#Required
	availability_domain = var.deployment_disaster_recovery_precheck_report_availability_domain
	deployment_id = oci_golden_gate_deployment.test_deployment.id
	fault_domain = var.deployment_disaster_recovery_precheck_report_fault_domain
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain of the placement to evaluate DR pre-checks for. 
* `deployment_id` - (Required) A unique Deployment identifier. 
* `fault_domain` - (Required) The fault domain of the placement to evaluate DR pre-checks for. 


## Attributes Reference

The following attributes are exported:

* `checks` - A list of precheck results. 
	* `code` - The code returned when GoldenGate reports an error while running a step during pipeline initialization. https://docs.oracle.com/en/middleware/goldengate/core/23/error-messages/ogg-00001-ogg-40000.html#GUID-97FF7AA7-7A5C-4AA7-B29F-3CC8D26761F2 
	* `corrective_action` - The corrective action for non-passing checks. Null for passed checks. 
	* `description` - Metadata about this specific object. 
	* `display_name` - An object's Display Name. 
	* `key` - UUID to uniquely identify the each check result. 
	* `related_resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource related to the corresponding check. 
	* `related_resource_type` - Type of resource related to corresponding check. 
	* `status` - Status of the DR precheck result.
* `precheck_status` - Status of the DR precheck result.
* `time_precheck_finished` - The timestamp when pre-check operation finished. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2024-10-26T20:19:29.600Z`. 
* `time_precheck_started` - The timestamp when pre-check started. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2024-10-26T20:19:29.600Z`. 

