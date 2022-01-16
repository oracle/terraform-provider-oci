---
subcategory: "Data Flow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_application"
sidebar_current: "docs-oci-resource-dataflow-application"
description: |-
  Provides the Application resource in Oracle Cloud Infrastructure Data Flow service
---

# oci_dataflow_application
This resource provides the Application resource in Oracle Cloud Infrastructure Data Flow service.

Creates an application.


## Example Usage

```hcl
resource "oci_dataflow_application" "test_application" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.application_display_name
	driver_shape = var.application_driver_shape
	executor_shape = var.application_executor_shape
	file_uri = var.application_file_uri
	language = var.application_language
	num_executors = var.application_num_executors
	spark_version = var.application_spark_version

	#Optional
	archive_uri = var.application_archive_uri
	arguments = var.application_arguments
	class_name = var.application_class_name
	configuration = var.application_configuration
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.application_description
	execute = var.application_execute
	freeform_tags = {"Department"= "Finance"}
	logs_bucket_uri = var.application_logs_bucket_uri
	metastore_id = var.metastore_id
	parameters {
		#Required
		name = var.application_parameters_name
		value = var.application_parameters_value
	}
	private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
	type = var.application_type
	warehouse_bucket_uri = var.application_warehouse_bucket_uri
}
```

## Argument Reference

The following arguments are supported:

* `archive_uri` - (Optional) (Updatable) An Oracle Cloud Infrastructure URI of an archive.zip file containing custom dependencies that may be used to support the execution a Python, Java, or Scala application. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `arguments` - (Optional) (Updatable) The arguments passed to the running application as command line arguments.  An argument is either a plain text or a placeholder. Placeholders are replaced using values from the parameters map.  Each placeholder specified must be represented in the parameters map else the request (POST or PUT) will fail with a HTTP 400 status code.  Placeholders are specified as `Service Api Spec`, where `name` is the name of the parameter. Example:  `[ "--input", "${input_file}", "--name", "John Doe" ]` If "input_file" has a value of "mydata.xml", then the value above will be translated to `--input mydata.xml --name "John Doe"` 
* `class_name` - (Optional) (Updatable) The class for the application. 
* `compartment_id` - (Required) (Updatable) The OCID of a compartment. 
* `configuration` - (Optional) (Updatable) The Spark configuration passed to the running process. See https://spark.apache.org/docs/latest/configuration.html#available-properties. Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" } Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is not allowed to be overwritten will cause a 400 status to be returned. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A user-friendly description. Avoid entering confidential information. 
* `display_name` - (Required) (Updatable) A user-friendly name. It does not have to be unique. Avoid entering confidential information. 
* `driver_shape` - (Required) (Updatable) The VM shape for the driver. Sets the driver cores and memory. 
* `execute` - (Optional) (Updatable) The input used for spark-submit command. For more details see https://spark.apache.org/docs/latest/submitting-applications.html#launching-applications-with-spark-submit. Supported options include ``--class``, ``--file``, ``--jars``, ``--conf``, ``--py-files``, and main application file with arguments. Example: ``--jars oci://path/to/a.jar,oci://path/to/b.jar --files oci://path/to/a.json,oci://path/to/b.csv --py-files oci://path/to/a.py,oci://path/to/b.py --conf spark.sql.crossJoin.enabled=true --class org.apache.spark.examples.SparkPi oci://path/to/main.jar 10`` Note: If execute is specified together with applicationId, className, configuration, fileUri, language, arguments, parameters during application create/update, or run create/submit, Data Flow service will use derived information from execute input only. 
* `executor_shape` - (Required) (Updatable) The VM shape for the executors. Sets the executor cores and memory. 
* `file_uri` - (Required) (Updatable) An Oracle Cloud Infrastructure URI of the file containing the application to execute. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `language` - (Required) (Updatable) The Spark language. 
* `logs_bucket_uri` - (Optional) (Updatable) An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `metastore_id` - (Optional) (Updatable) The OCID of Oracle Cloud Infrastructure Hive Metastore. 
* `num_executors` - (Required) (Updatable) The number of executor VMs requested. 
* `parameters` - (Optional) (Updatable) An array of name/value pairs used to fill placeholders found in properties like `Application.arguments`.  The name must be a string of one or more word characters (a-z, A-Z, 0-9, _).  The value can be a string of 0 or more characters of any kind. Example:  [ { name: "iterations", value: "10"}, { name: "input_file", value: "mydata.xml" }, { name: "variable_x", value: "${x}"} ] 
	* `name` - (Required) (Updatable) The name of the parameter.  It must be a string of one or more word characters (a-z, A-Z, 0-9, _). Examples: "iterations", "input_file" 
	* `value` - (Required) (Updatable) The value of the parameter. It must be a string of 0 or more characters of any kind. Examples: "" (empty string), "10", "mydata.xml", "${x}" 
* `private_endpoint_id` - (Optional) (Updatable) The OCID of a private endpoint. 
* `spark_version` - (Required) (Updatable) The Spark version utilized to run the application. 
* `type` - (Optional) The Spark application processing type. 
* `warehouse_bucket_uri` - (Optional) (Updatable) An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory for BATCH SQL runs. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `archive_uri` - An Oracle Cloud Infrastructure URI of an archive.zip file containing custom dependencies that may be used to support the execution a Python, Java, or Scala application. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `arguments` - The arguments passed to the running application as command line arguments.  An argument is either a plain text or a placeholder. Placeholders are replaced using values from the parameters map.  Each placeholder specified must be represented in the parameters map else the request (POST or PUT) will fail with a HTTP 400 status code.  Placeholders are specified as `Service Api Spec`, where `name` is the name of the parameter. Example:  `[ "--input", "${input_file}", "--name", "John Doe" ]` If "input_file" has a value of "mydata.xml", then the value above will be translated to `--input mydata.xml --name "John Doe"` 
* `class_name` - The class for the application. 
* `compartment_id` - The OCID of a compartment. 
* `configuration` - The Spark configuration passed to the running process. See https://spark.apache.org/docs/latest/configuration.html#available-properties. Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" } Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is not allowed to be overwritten will cause a 400 status to be returned. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A user-friendly description. 
* `display_name` - A user-friendly name. This name is not necessarily unique. 
* `driver_shape` - The VM shape for the driver. Sets the driver cores and memory. 
* `execute` - The input used for spark-submit command. For more details see https://spark.apache.org/docs/latest/submitting-applications.html#launching-applications-with-spark-submit. Supported options include ``--class``, ``--file``, ``--jars``, ``--conf``, ``--py-files``, and main application file with arguments. Example: ``--jars oci://path/to/a.jar,oci://path/to/b.jar --files oci://path/to/a.json,oci://path/to/b.csv --py-files oci://path/to/a.py,oci://path/to/b.py --conf spark.sql.crossJoin.enabled=true --class org.apache.spark.examples.SparkPi oci://path/to/main.jar 10`` Note: If execute is specified together with applicationId, className, configuration, fileUri, language, arguments, parameters during application create/update, or run create/submit, Data Flow service will use derived information from execute input only. 
* `executor_shape` - The VM shape for the executors. Sets the executor cores and memory. 
* `file_uri` - An Oracle Cloud Infrastructure URI of the file containing the application to execute. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The application ID. 
* `language` - The Spark language. 
* `logs_bucket_uri` - An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `metastore_id` - The OCID of Oracle Cloud Infrastructure Hive Metastore. 
* `num_executors` - The number of executor VMs requested. 
* `owner_principal_id` - The OCID of the user who created the resource. 
* `owner_user_name` - The username of the user who created the resource.  If the username of the owner does not exist, `null` will be returned and the caller should refer to the ownerPrincipalId value instead. 
* `parameters` - An array of name/value pairs used to fill placeholders found in properties like `Application.arguments`.  The name must be a string of one or more word characters (a-z, A-Z, 0-9, _).  The value can be a string of 0 or more characters of any kind. Example:  [ { name: "iterations", value: "10"}, { name: "input_file", value: "mydata.xml" }, { name: "variable_x", value: "${x}"} ] 
	* `name` - The name of the parameter.  It must be a string of one or more word characters (a-z, A-Z, 0-9, _). Examples: "iterations", "input_file" 
	* `value` - The value of the parameter. It must be a string of 0 or more characters of any kind. Examples: "" (empty string), "10", "mydata.xml", "${x}" 
* `private_endpoint_id` - The OCID of a private endpoint. 
* `spark_version` - The Spark version utilized to run the application. 
* `state` - The current state of this application. 
* `time_created` - The date and time a application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time a application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `type` - The Spark application processing type. 
* `warehouse_bucket_uri` - An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory for BATCH SQL runs. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Application
	* `update` - (Defaults to 20 minutes), when updating the Application
	* `delete` - (Defaults to 20 minutes), when destroying the Application


## Import

Applications can be imported using the `id`, e.g.

```
$ terraform import oci_dataflow_application.test_application "id"
```

