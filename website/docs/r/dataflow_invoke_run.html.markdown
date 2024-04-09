---
subcategory: "Data Flow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_invoke_run"
sidebar_current: "docs-oci-resource-dataflow-invoke_run"
description: |-
  Provides the Invoke Run resource in Oracle Cloud Infrastructure Data Flow service
---

# oci_dataflow_invoke_run
This resource provides the Invoke Run resource in Oracle Cloud Infrastructure Data Flow service.

Creates a run for an application.


## Example Usage

```hcl
resource "oci_dataflow_invoke_run" "test_invoke_run" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	application_id = oci_dataflow_application.test_application.id
	application_log_config {
		#Required
		log_group_id = oci_logging_log_group.test_log_group.id
		log_id = oci_logging_log.test_log.id
	}
	archive_uri = var.invoke_run_archive_uri
	arguments = var.invoke_run_arguments
	configuration = var.invoke_run_configuration
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.invoke_run_display_name
	driver_shape = var.invoke_run_driver_shape
	driver_shape_config {

		#Optional
		memory_in_gbs = var.invoke_run_driver_shape_config_memory_in_gbs
		ocpus = var.invoke_run_driver_shape_config_ocpus
	}
	execute = var.invoke_run_execute
	executor_shape = var.invoke_run_executor_shape
	executor_shape_config {

		#Optional
		memory_in_gbs = var.invoke_run_executor_shape_config_memory_in_gbs
		ocpus = var.invoke_run_executor_shape_config_ocpus
	}
	freeform_tags = {"Department"= "Finance"}
	idle_timeout_in_minutes = var.invoke_run_idle_timeout_in_minutes
	logs_bucket_uri = var.invoke_run_logs_bucket_uri
	max_duration_in_minutes = var.invoke_run_max_duration_in_minutes
	metastore_id = var.metastore_id
	num_executors = var.invoke_run_num_executors
	opc_parent_rpt_url = var.invoke_run_opc_parent_rpt_url
	parameters {
		#Required
		name = var.invoke_run_parameters_name
		value = var.invoke_run_parameters_value
	}
	pool_id = oci_dataflow_pool.test_pool.id
	spark_version = var.invoke_run_spark_version
	type = var.invoke_run_type
	warehouse_bucket_uri = var.invoke_run_warehouse_bucket_uri
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The OCID of the associated application. If this value is set, then no value for the execute parameter is required. If this value is not set, then a value for the execute parameter is required, and a new application is created and associated with the new run. 
* `application_log_config` - (Optional) Logging details of Application logs for Data Flow Run. 
	* `log_group_id` - (Required) The log group id for where log objects will be for Data Flow Runs. 
	* `log_id` - (Required) The log id of the log object the Application Logs of Data Flow Run will be shipped to. 
* `archive_uri` - (Optional) A comma separated list of one or more archive files as Oracle Cloud Infrastructure URIs. For example, ``oci://path/to/a.zip,oci://path/to/b.zip``. An Oracle Cloud Infrastructure URI of an archive.zip file containing custom dependencies that may be used to support the execution of a Python, Java, or Scala application. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `arguments` - (Optional) The arguments passed to the running application as command line arguments.  An argument is either a plain text or a placeholder. Placeholders are replaced using values from the parameters map.  Each placeholder specified must be represented in the parameters map else the request (POST or PUT) will fail with a HTTP 400 status code.  Placeholders are specified as `Service Api Spec`, where `name` is the name of the parameter. Example:  `[ "--input", "${input_file}", "--name", "John Doe" ]` If "input_file" has a value of "mydata.xml", then the value above will be translated to `--input mydata.xml --name "John Doe"` 
* `asynchronous` -  (Optional) Flag to invoke run asynchronously. The default is true and Terraform provider will not wait for run resource to reach target state of `SUCCEEDED`, `FAILED` or `CANCELLED` before exiting. User must wait to perform operations that need resource to be in target states. Set this to false to override this behavior. 
* `compartment_id` - (Required) (Updatable) The OCID of a compartment. 
* `configuration` - (Optional) The Spark configuration passed to the running process. See https://spark.apache.org/docs/latest/configuration.html#available-properties Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" } Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is not allowed to be overwritten will cause a 400 status to be returned.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name that does not have to be unique. Avoid entering confidential information. If this value is not specified, it will be derived from the associated application's displayName or set by API using fileUri's application file name. 
* `driver_shape` - (Optional) The VM shape for the driver. Sets the driver cores and memory. 
* `driver_shape_config` - (Optional) This is used to configure the shape of the driver or executor if a flexible shape is used. 
	* `memory_in_gbs` - (Optional) The amount of memory used for the driver or executors. 
	* `ocpus` - (Optional) The total number of OCPUs used for the driver or executors. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `execute` - (Optional) The input used for spark-submit command. For more details see https://spark.apache.org/docs/latest/submitting-applications.html#launching-applications-with-spark-submit. Supported options include ``--class``, ``--file``, ``--jars``, ``--conf``, ``--py-files``, and main application file with arguments. Example: ``--jars oci://path/to/a.jar,oci://path/to/b.jar --files oci://path/to/a.json,oci://path/to/b.csv --py-files oci://path/to/a.py,oci://path/to/b.py --conf spark.sql.crossJoin.enabled=true --class org.apache.spark.examples.SparkPi oci://path/to/main.jar 10`` Note: If execute is specified together with applicationId, className, configuration, fileUri, language, arguments, parameters during application create/update, or run create/submit, Data Flow service will use derived information from execute input only. 
* `executor_shape` - (Optional) The VM shape for the executors. Sets the executor cores and memory. 
* `executor_shape_config` - (Optional) This is used to configure the shape of the driver or executor if a flexible shape is used. 
	* `memory_in_gbs` - (Optional) The amount of memory used for the driver or executors. 
	* `ocpus` - (Optional) The total number of OCPUs used for the driver or executors. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `idle_timeout_in_minutes` - (Optional) (Updatable) The timeout value in minutes used to manage Runs. A Run would be stopped after inactivity for this amount of time period. Note: This parameter is currently only applicable for Runs of type `SESSION`. Default value is 2880 minutes (2 days) 
* `logs_bucket_uri` - (Optional) An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `max_duration_in_minutes` - (Optional) (Updatable) The maximum duration in minutes for which an Application should run. Data Flow Run would be terminated once it reaches this duration from the time it transitions to `IN_PROGRESS` state. 
* `metastore_id` - (Optional) The OCID of Oracle Cloud Infrastructure Hive Metastore. 
* `num_executors` - (Optional) The number of executor VMs requested. 
* `opc_parent_rpt_url` - (Optional) (Optional header param, required for Resource Principal version 3.0+) Parent resource control plane endpoint used to exchange for upper level resource principal token. 
* `parameters` - (Optional) An array of name/value pairs used to fill placeholders found in properties like `Application.arguments`.  The name must be a string of one or more word characters (a-z, A-Z, 0-9, _).  The value can be a string of 0 or more characters of any kind. Example:  [ { name: "iterations", value: "10"}, { name: "input_file", value: "mydata.xml" }, { name: "variable_x", value: "${x}"} ] 
	* `name` - (Required) The name of the parameter.  It must be a string of one or more word characters (a-z, A-Z, 0-9, _). Examples: "iterations", "input_file" 
	* `value` - (Required) The value of the parameter. It must be a string of 0 or more characters of any kind. Examples: "" (empty string), "10", "mydata.xml", "${x}" 
* `pool_id` - (Optional) The OCID of a pool. Unique Id to indentify a dataflow pool resource. 
* `spark_version` - (Optional) The Spark version utilized to run the application. This value may be set if applicationId is not since the Spark version will be taken from the associated application. 
* `type` - (Optional) The Spark application processing type. 
* `warehouse_bucket_uri` - (Optional) An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory for BATCH SQL runs. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `application_id` - The application ID. 
* `application_log_config` - Logging details of Application logs for Data Flow Run. 
	* `log_group_id` - The log group id for where log objects will be for Data Flow Runs. 
	* `log_id` - The log id of the log object the Application Logs of Data Flow Run will be shipped to. 
* `archive_uri` - A comma separated list of one or more archive files as Oracle Cloud Infrastructure URIs. For example, ``oci://path/to/a.zip,oci://path/to/b.zip``. An Oracle Cloud Infrastructure URI of an archive.zip file containing custom dependencies that may be used to support the execution of a Python, Java, or Scala application. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `arguments` - The arguments passed to the running application as command line arguments.  An argument is either a plain text or a placeholder. Placeholders are replaced using values from the parameters map.  Each placeholder specified must be represented in the parameters map else the request (POST or PUT) will fail with a HTTP 400 status code.  Placeholders are specified as `Service Api Spec`, where `name` is the name of the parameter. Example:  `[ "--input", "${input_file}", "--name", "John Doe" ]` If "input_file" has a value of "mydata.xml", then the value above will be translated to `--input mydata.xml --name "John Doe"` 
* `class_name` - The class for the application. 
* `compartment_id` - The OCID of a compartment. 
* `configuration` - The Spark configuration passed to the running process. See https://spark.apache.org/docs/latest/configuration.html#available-properties. Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" } Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is not allowed to be overwritten will cause a 400 status to be returned. 
* `data_read_in_bytes` - The data read by the run in bytes. 
* `data_written_in_bytes` - The data written by the run in bytes. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. This name is not necessarily unique. 
* `driver_shape` - The VM shape for the driver. Sets the driver cores and memory. 
* `driver_shape_config` - This is used to configure the shape of the driver or executor if a flexible shape is used. 
	* `memory_in_gbs` - The amount of memory used for the driver or executors. 
	* `ocpus` - The total number of OCPUs used for the driver or executors. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `execute` - The input used for spark-submit command. For more details see https://spark.apache.org/docs/latest/submitting-applications.html#launching-applications-with-spark-submit. Supported options include ``--class``, ``--file``, ``--jars``, ``--conf``, ``--py-files``, and main application file with arguments. Example: ``--jars oci://path/to/a.jar,oci://path/to/b.jar --files oci://path/to/a.json,oci://path/to/b.csv --py-files oci://path/to/a.py,oci://path/to/b.py --conf spark.sql.crossJoin.enabled=true --class org.apache.spark.examples.SparkPi oci://path/to/main.jar 10`` Note: If execute is specified together with applicationId, className, configuration, fileUri, language, arguments, parameters during application create/update, or run create/submit, Data Flow service will use derived information from execute input only. 
* `executor_shape` - The VM shape for the executors. Sets the executor cores and memory. 
* `executor_shape_config` - This is used to configure the shape of the driver or executor if a flexible shape is used. 
	* `memory_in_gbs` - The amount of memory used for the driver or executors. 
	* `ocpus` - The total number of OCPUs used for the driver or executors. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `file_uri` - An Oracle Cloud Infrastructure URI of the file containing the application to execute. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The ID of a run. 
* `idle_timeout_in_minutes` - The timeout value in minutes used to manage Runs. A Run would be stopped after inactivity for this amount of time period. Note: This parameter is currently only applicable for Runs of type `SESSION`. Default value is 2880 minutes (2 days) 
* `language` - The Spark language. 
* `lifecycle_details` - The detailed messages about the lifecycle state. 
* `logs_bucket_uri` - An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `max_duration_in_minutes` - The maximum duration in minutes for which an Application should run. Data Flow Run would be terminated once it reaches this duration from the time it transitions to `IN_PROGRESS` state. 
* `metastore_id` - The OCID of Oracle Cloud Infrastructure Hive Metastore. 
* `num_executors` - The number of executor VMs requested. 
* `opc_request_id` - Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID. 
* `owner_principal_id` - The OCID of the user who created the resource. 
* `owner_user_name` - The username of the user who created the resource.  If the username of the owner does not exist, `null` will be returned and the caller should refer to the ownerPrincipalId value instead. 
* `parameters` - An array of name/value pairs used to fill placeholders found in properties like `Application.arguments`.  The name must be a string of one or more word characters (a-z, A-Z, 0-9, _).  The value can be a string of 0 or more characters of any kind. Example:  [ { name: "iterations", value: "10"}, { name: "input_file", value: "mydata.xml" }, { name: "variable_x", value: "${x}"} ] 
	* `name` - The name of the parameter.  It must be a string of one or more word characters (a-z, A-Z, 0-9, _). Examples: "iterations", "input_file" 
	* `value` - The value of the parameter. It must be a string of 0 or more characters of any kind. Examples: "" (empty string), "10", "mydata.xml", "${x}" 
* `pool_id` - The OCID of a pool. Unique Id to indentify a dataflow pool resource. 
* `private_endpoint_dns_zones` - An array of DNS zone names. Example: `[ "app.examplecorp.com", "app.examplecorp2.com" ]` 
* `private_endpoint_id` - The OCID of a private endpoint. 
* `private_endpoint_max_host_count` - The maximum number of hosts to be accessed through the private endpoint. This value is used to calculate the relevant CIDR block and should be a multiple of 256.  If the value is not a multiple of 256, it is rounded up to the next multiple of 256. For example, 300 is rounded up to 512. 
* `private_endpoint_nsg_ids` - An array of network security group OCIDs. 
* `private_endpoint_subnet_id` - The OCID of a subnet. 
* `run_duration_in_milliseconds` - The duration of the run in milliseconds. 
* `spark_version` - The Spark version utilized to run the application. 
* `state` - The current state of this run. 
* `time_created` - The date and time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `total_ocpu` - The total number of oCPU requested by the run. 
* `type` - The Spark application processing type. 
* `warehouse_bucket_uri` - An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory for BATCH SQL runs. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Invoke Run
	* `update` - (Defaults to 20 minutes), when updating the Invoke Run
	* `delete` - (Defaults to 20 minutes), when destroying the Invoke Run


## Import

InvokeRuns can be imported using the `id`, e.g.

```
$ terraform import oci_dataflow_invoke_run.test_invoke_run "id"
```

## Note

At a time service allows only one run to succeed if user is trying to invoke runs on multiple applications which have Private Endpoints and service will proceed invoking only one run and put the rest of them in failed state.
