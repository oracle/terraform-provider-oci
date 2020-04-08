---
subcategory: "Dataflow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_invoke_run"
sidebar_current: "docs-oci-datasource-dataflow-invoke_run"
description: |-
  Provides details about a specific Invoke Run in Oracle Cloud Infrastructure Dataflow service
---

# Data Source: oci_dataflow_invoke_run
This data source provides details about a specific Invoke Run resource in Oracle Cloud Infrastructure Dataflow service.

Retrieves the run for the specified `runId`.


## Example Usage

```hcl
data "oci_dataflow_invoke_run" "test_invoke_run" {
	#Required
	run_id = "${oci_dataflow_run.test_run.id}"
}
```

## Argument Reference

The following arguments are supported:

* `run_id` - (Required) The unique ID for the run 


## Attributes Reference

The following attributes are exported:

* `application_id` - The application ID. 
* `arguments` - The arguments passed to the running application as command line arguments.  An argument is either a plain text or a placeholder. Placeholders are replaced using values from the parameters map.  Each placeholder specified must be represented in the parameters map else the request (POST or PUT) will fail with a HTTP 400 status code.  Placeholders are specified as `Service Api Spec`, where `name` is the name of the parameter. Example:  `[ "--input", "${input_file}", "--name", "John Doe" ]` If "input_file" has a value of "mydata.xml", then the value above will be translated to `--input mydata.xml --name "John Doe"` 
* `class_name` - The class for the application. 
* `compartment_id` - The OCID of a compartment. 
* `configuration` - The Spark configuration passed to the running process. See https://spark.apache.org/docs/latest/configuration.html#available-properties Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" } Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is not allowed to be overwritten will cause a 400 status to be returned. 
* `data_read_in_bytes` - The data read by the run in bytes. 
* `data_written_in_bytes` - The data written by the run in bytes. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. This name is not necessarily unique. 
* `driver_shape` - The VM shape for the driver. Sets the driver cores and memory. 
* `executor_shape` - The VM shape for the executors. Sets the executor cores and memory. 
* `file_uri` - An Oracle Cloud Infrastructure URI of the file containing the application to execute. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The ID of a run. 
* `language` - The Spark language. 
* `lifecycle_details` - The detailed messages about the lifecycle state. 
* `logs_bucket_uri` - An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat 
* `num_executors` - The number of executor VMs requested. 
* `opc_request_id` - Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID. 
* `owner_principal_id` - The OCID of the user who created the resource. 
* `owner_user_name` - The username of the user who created the resource.  If the username of the owner does not exist, `null` will be returned and the caller should refer to the ownerPrincipalId value instead. 
* `parameters` - An array of name/value pairs used to fill placeholders found in properties like `Application.arguments`.  The name must be a string of one or more word characters (a-z, A-Z, 0-9, _).  The value can be a string of 0 or more characters of any kind. Example:  [ { name: "iterations", value: "10"}, { name: "input_file", value: "mydata.xml" }, { name: "variable_x", value: "${x}"} ] 
	* `name` - The name of the parameter.  It must be a string of one or more word characters (a-z, A-Z, 0-9, _). Examples: "iterations", "input_file" 
	* `value` - The value of the parameter. It must be a string of 0 or more characters of any kind. Examples: "" (empty string), "10", "mydata.xml", "${x}" 
* `run_duration_in_milliseconds` - The duration of the run in milliseconds. 
* `spark_version` - The Spark version utilized to run the application. 
* `state` - The current state of this run. 
* `time_created` - The date and time a application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time a application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `total_ocpu` - The total number of oCPU requested by the run. 
* `warehouse_bucket_uri` - An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory for BATCH SQL runs. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat 

