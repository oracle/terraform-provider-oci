---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_crypto_analysis_result"
sidebar_current: "docs-oci-datasource-jms-fleet_crypto_analysis_result"
description: |-
  Provides details about a specific Fleet Crypto Analysis Result in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_crypto_analysis_result
This data source provides details about a specific Fleet Crypto Analysis Result resource in Oracle Cloud Infrastructure Jms service.

Retrieve the metadata for the result of a Crypto event analysis.

## Example Usage

```hcl
data "oci_jms_fleet_crypto_analysis_result" "test_fleet_crypto_analysis_result" {
	#Required
	crypto_analysis_result_id = var.fleet_crypto_analysis_result_id
	fleet_id = oci_jms_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `crypto_analysis_result_id` - (Required) The OCID of the analysis result.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.


## Attributes Reference

The following attributes are exported:

* `aggregation_mode` - The result aggregation mode
* `bucket` - The Object Storage bucket name of this analysis result.
* `crypto_roadmap_version` - The Crypto Roadmap version used to perform the analysis.
* `finding_count` - Total number of findings with the analysis.
* `fleet_id` - The fleet OCID.
* `host_name` - The hostname of the managed instance.
* `id` - The OCID to identify this analysis results.
* `managed_instance_id` - The managed instance OCID.
* `namespace` - The Object Storage namespace of this analysis result.
* `non_compliant_finding_count` - Total number of non-compliant findings with the analysis. A non-compliant finding means the application won't work properly with the changes introduced by the Crypto Roadmap version used by the analysis. 
* `object` - The Object Storage object name of this analysis result.
* `summarized_event_count` - Total number of summarized events. Summarized events are deduplicated events of interest.
* `time_created` - The time the result is compiled.
* `time_finished` - The time the JFR recording has finished.
* `time_first_event` - Time of the first event in the analysis.
* `time_last_event` - Time of the last event in the analysis.
* `time_started` - The time the JFR recording has started.
* `total_event_count` - Total number of events in the analysis.
* `work_request_id` - The OCID of the work request to start the analysis.

