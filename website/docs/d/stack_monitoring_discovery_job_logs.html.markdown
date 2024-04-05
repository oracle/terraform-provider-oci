---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_discovery_job_logs"
sidebar_current: "docs-oci-datasource-stack_monitoring-discovery_job_logs"
description: |-
  Provides the list of Discovery Job Logs in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_discovery_job_logs
This data source provides the list of Discovery Job Logs in Oracle Cloud Infrastructure Stack Monitoring service.

API to get all the logs of a Discovery Job.


## Example Usage

```hcl
data "oci_stack_monitoring_discovery_job_logs" "test_discovery_job_logs" {
	#Required
	discovery_job_id = oci_stack_monitoring_discovery_job.test_discovery_job.id

	#Optional
	log_type = var.discovery_job_log_log_type
}
```

## Argument Reference

The following arguments are supported:

* `discovery_job_id` - (Required) The Discovery Job ID
* `log_type` - (Optional) The log type like INFO, WARNING, ERROR, SUCCESS


## Attributes Reference

The following attributes are exported:

* `discovery_job_log_collection` - The list of discovery_job_log_collection.

### DiscoveryJobLog Reference

The following attributes are exported:

* `items` - List of logs
	* `id` - The OCID of Discovery job
	* `log_message` - Log message
	* `log_type` - Type of log (INFO, WARNING, ERROR or SUCCESS)
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - Time the Job log was created

