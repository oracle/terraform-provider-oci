---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_crypto_analysis_results"
sidebar_current: "docs-oci-datasource-jms-fleet_crypto_analysis_results"
description: |-
  Provides the list of Fleet Crypto Analysis Results in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_crypto_analysis_results
This data source provides the list of Fleet Crypto Analysis Results in Oracle Cloud Infrastructure Jms service.

Lists the results of a Crypto event analysis.

## Example Usage

```hcl
data "oci_jms_fleet_crypto_analysis_results" "test_fleet_crypto_analysis_results" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id

	#Optional
	aggregation_mode = var.fleet_crypto_analysis_result_aggregation_mode
	managed_instance_id = oci_osmanagement_managed_instance.test_managed_instance.id
	time_end = var.fleet_crypto_analysis_result_time_end
	time_start = var.fleet_crypto_analysis_result_time_start
}
```

## Argument Reference

The following arguments are supported:

* `aggregation_mode` - (Optional) The aggregation mode of the crypto event analysis result.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the related managed instance.
* `time_end` - (Optional) The end of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `crypto_analysis_result_collection` - The list of crypto_analysis_result_collection.

### FleetCryptoAnalysisResult Reference

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
* `time_first_event` - Time of the first event in the analysis.
* `time_last_event` - Time of the last event in the analysis.
* `total_event_count` - Total number of events in the analysis.
* `work_request_id` - The OCID of the work request to start the analysis.

