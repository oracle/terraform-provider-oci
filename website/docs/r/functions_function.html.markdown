---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_function"
sidebar_current: "docs-oci-resource-functions-function"
description: |-
  Provides the Function resource in Oracle Cloud Infrastructure Functions service
---

# oci_functions_function
This resource provides the Function resource in Oracle Cloud Infrastructure Functions service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/functions/latest/Function

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/functions

Creates a new function.

## Example Usage

```hcl
resource "oci_functions_function" "test_function" {
	#Required
	application_id = oci_functions_application.test_application.id
	display_name = var.function_display_name
	memory_in_mbs = var.function_memory_in_mbs
	source_details {
		#Required
		source_type = var.function_source_details_source_type

		#Optional
		archive_source_details {
			#Required
			archive_source_type = var.function_source_details_archive_source_details_archive_source_type

			#Optional
			archive_file = var.function_source_details_archive_source_details_archive_file
			bucket = var.function_source_details_archive_source_details_bucket
			namespace = var.function_source_details_archive_source_details_namespace
			object = var.function_source_details_archive_source_details_object
			object_version_id = oci_objectstorage_object_version.test_object_version.id
		}
		handler = var.function_source_details_handler
		image = var.function_source_details_image
		image_digest = var.function_source_details_image_digest
		pbf_listing_id = oci_functions_pbf_listing.test_pbf_listing.id
		runtime_config {
			#Required
				functions_runtime_name = var.functions_runtime_name
				functions_runtime_version_id = var.functions_runtime_version_id
			runtime_config_type = var.function_source_details_runtime_config_runtime_config_type
		}
	}

	#Optional
	config = var.function_config
	defined_tags = {"Operations.CostCenter"= "42"}
	detached_mode_timeout_in_seconds = var.function_detached_mode_timeout_in_seconds
	failure_destination {
		#Required
		kind = var.function_failure_destination_kind

		#Optional
		channel_id = oci_mysql_channel.test_channel.id
		queue_id = oci_queue_queue.test_queue.id
		stream_id = oci_streaming_stream.test_stream.id
		topic_id = oci_ons_notification_topic.test_notification_topic.id
	}
	freeform_tags = {"Department"= "Finance"}
	provisioned_concurrency_config {
		#Required
		strategy = var.function_provisioned_concurrency_config_strategy

		#Optional
		count = var.function_provisioned_concurrency_config_count
	}
	success_destination {
		#Required
		kind = var.function_success_destination_kind

		#Optional
		channel_id = oci_mysql_channel.test_channel.id
		queue_id = oci_queue_queue.test_queue.id
		stream_id = oci_streaming_stream.test_stream.id
		topic_id = oci_ons_notification_topic.test_notification_topic.id
	}
	timeout_in_seconds = var.function_timeout_in_seconds
	trace_config {

		#Optional
		is_enabled = var.function_trace_config_is_enabled
	}
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Required) The OCID of the application this function belongs to.
* `config` - (Optional) (Updatable) Function configuration. These values are passed on to the function as environment variables, this overrides application configuration values. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`

	The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `detached_mode_timeout_in_seconds` - (Optional) (Updatable) Timeout for detached function invocations. Value in seconds.
* `display_name` - (Required) The display name of the function. The display name must be unique within the application containing the function. Avoid entering confidential information.
* `failure_destination` - (Optional) (Updatable) An object that represents the destination to which Oracle Functions will send an invocation record with the details of the error of the failed detached function invocation. A notification is an example of a failure destination.  Example: `{"kind": "NOTIFICATION", "topicId": "topic_OCID"}`
	* `channel_id` - (Applicable when kind=QUEUE) (Updatable) The ID of the channel in the queue.
	* `kind` - (Required) (Updatable) The type of destination for the response to a failed detached function invocation.
	* `queue_id` - (Required when kind=QUEUE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the queue.
	* `stream_id` - (Required when kind=STREAM) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream.
	* `topic_id` - (Required when kind=NOTIFICATION) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `image` - (Optional) (Updatable) Deprecated. The 'image' field has been deprecated. Use `source_details.image` in a `CONTAINER_IMAGE` source_details block instead. If both fields are specified, then 'source_details.image' will be used. Example: `phx.ocir.io/ten/functions/function:0.0.1`
* `image_digest` - (Optional) (Updatable) Deprecated. The 'image_digest' field has been deprecated. Use `source_details.image_digest` in a `CONTAINER_IMAGE` source_details block instead. If both fields are specified, then 'source_details.image_digest' will be used. Example: `sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7`
* `memory_in_mbs` - (Required) (Updatable) Maximum usable memory for the function (MiB).
* `provisioned_concurrency_config` - (Optional) (Updatable) Define the strategy for provisioned concurrency for the function.
	* `count` - (Required when strategy=CONSTANT) (Updatable) Configuration specifying a constant amount of provisioned concurrency.
	* `strategy` - (Required) (Updatable) The strategy for provisioned concurrency to be used.
* `source_details` - (Required) (Updatable) The source details for creating the Function. The function can be created from various sources.
	* `archive_source_details` - (Required when source_type=ARCHIVE) (Updatable) The details required to create an Archive-based function source.  This mode is used when the function code is provided as an archive, either from Object Storage or directly uploaded by the API caller.  It is suitable for scenarios where the function code is packaged as a single archive file.
		* `archive_file` - (Required when archive_source_type=DIRECT_ARCHIVE) (Updatable) The base64-encoded archive file of the function code. The archive file must contain all the files for the function. Please refer to functions documentation for maximum allowed size and supported archive formats.
		* `archive_source_type` - (Required) (Updatable) Type of the Archive Source. Possible values: OBJECT_STORAGE_ARCHIVE and DIRECT_ARCHIVE.
		* `bucket` - (Required when archive_source_type=OBJECT_STORAGE_ARCHIVE) (Updatable) The name of the Object Storage bucket.
		* `namespace` - (Required when archive_source_type=OBJECT_STORAGE_ARCHIVE) (Updatable) The Object Storage namespace.
		* `object` - (Required when archive_source_type=OBJECT_STORAGE_ARCHIVE) (Updatable) The name of the Object Storage object.
		* `object_version_id` - (Applicable when archive_source_type=OBJECT_STORAGE_ARCHIVE) (Updatable) VersionId used to identify a particular version of the object. If not specified, the latest version of the object is used.
	* `handler` - (Applicable when source_type=ARCHIVE) (Updatable) The function handler that is executed when the function is invoked. The value of this field depends on the runtime used
	* `image` - (Required when source_type=CONTAINER_IMAGE) (Updatable) The qualified name of the Docker image to use in the function, including the image tag. The image should be in the Oracle Cloud Infrastructure Registry that is in the same region as the function itself. Example: `phx.ocir.io/ten/functions/function:0.0.1`
	* `image_digest` - (Applicable when source_type=CONTAINER_IMAGE) (Updatable) The image digest for the version of the image that will be pulled when invoking this function. If no value is specified, the digest currently associated with the image in the Oracle Cloud Infrastructure Registry will be used. Example: `sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7`
	* `pbf_listing_id` - (Required when source_type=PRE_BUILT_FUNCTIONS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PbfListing this function is sourced from.
	* `runtime_config` - (Required when source_type=ARCHIVE) (Updatable) FunctionsRuntime configuration used to create a function.
		* `functions_runtime_name` - (Required) (Updatable) The name of the FunctionsRuntime this function is to be associated with.
		* `functions_runtime_version_id` - (Required) (Updatable) The OCID of the FunctionsRuntimeVersion to use for the Function in manual mode.
		* `runtime_config_type` - (Required) (Updatable) Type of the FunctionsRuntime Config. Possible values: FUNCTION_UPDATE and MANUAL.
		* `source_type` - (Required) (Updatable) Type of the Function Source. Possible values: CONTAINER_IMAGE, PRE_BUILT_FUNCTIONS and ARCHIVE.
* `success_destination` - (Optional) (Updatable) An object that represents the destination to which Oracle Functions will send an invocation record with the details of the successful detached function invocation. A stream is an example of a success destination.  Example: `{"kind": "STREAM", "streamId": "stream_OCID"}`
	* `channel_id` - (Applicable when kind=QUEUE) (Updatable) The ID of the channel in the queue.
	* `kind` - (Required) (Updatable) The type of destination for the response to a successful detached function invocation.
	* `queue_id` - (Required when kind=QUEUE) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the queue.
	* `stream_id` - (Required when kind=STREAM) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream.
	* `topic_id` - (Required when kind=NOTIFICATION) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
* `timeout_in_seconds` - (Optional) (Updatable) Timeout for executions of the function. Value in seconds.
* `trace_config` - (Optional) (Updatable) Define the tracing configuration for a function.
	* `is_enabled` - (Optional) (Updatable) Define if tracing is enabled for the resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `application_id` - The OCID of the application the function belongs to.
* `compartment_id` - The OCID of the compartment that contains the function.
* `config` - Function configuration. Overrides application configuration. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`

	The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `detached_mode_timeout_in_seconds` - Timeout for detached function invocations. Value in seconds.  Example: `{"detachedModeTimeoutInSeconds": 900}`
* `display_name` - The display name of the function. The display name is unique within the application containing the function.
* `failure_destination` - An object that represents the destination to which Oracle Functions will send an invocation record with the details of the error of the failed detached function invocation. A notification is an example of a failure destination.  Example: `{"kind": "NOTIFICATION", "topicId": "topic_OCID"}`
	* `channel_id` - The ID of the channel in the queue.
	* `kind` - The type of destination for the response to a failed detached function invocation.
	* `queue_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the queue.
	* `stream_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream.
	* `topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function.
* `image` - Deprecated compatibility field. Use `source_details.image` in a `CONTAINER_IMAGE` source_details block instead. If both fields are specified, then 'source_details.image' will be used.
* `image_digest` - Deprecated compatibility field. Use `source_details.image_digest` in a `CONTAINER_IMAGE` source_details block instead. If both fields are specified, then 'source_details.image_digest' will be used.
* `invoke_endpoint` - The base https invoke URL to set on a client in order to invoke a function. This URL will never change over the lifetime of the function and can be cached.
* `memory_in_mbs` - Maximum usable memory for the function (MiB).
* `provisioned_concurrency_config` - Define the strategy for provisioned concurrency for the function.
	* `count` - Configuration specifying a constant amount of provisioned concurrency.
	* `strategy` - The strategy for provisioned concurrency to be used.
* `shape` - The processor shape (`GENERIC_X86`/`GENERIC_ARM`) on which to run functions in the application, extracted from the image manifest.
* `source_details` - The source details for the Function. The function can be created from various sources.
	* `archive_source_details` - The details for the Archive source of the Function.  This mode is used when the function code is provided as an archive, either from Object Storage or directly uploaded by the API caller.  It is suitable for scenarios where the function code is packaged as a single archive file.
		* `archive_source_type` - Type of the Archive Source. Possible values: OBJECT_STORAGE_ARCHIVE and DIRECT_ARCHIVE.
		* `bucket` - The name of the Object Storage bucket.
		* `namespace` - The Object Storage namespace.
		* `object` - The name of the Object Storage object.
		* `object_version_id` - VersionId used to identify a particular version of the object. If not specified, the latest version of the object is used.
	* `handler` - The function handler that is executed when the function is invoked. The value of this field depends on the runtime used
	* `image` - The qualified name of the Docker image to use in the function, including the image tag. The image should be in the Oracle Cloud Infrastructure Registry that is in the same region as the function itself. Example: `phx.ocir.io/ten/functions/function:0.0.1`
	* `image_digest` - The image digest for the version of the image that will be pulled when invoking this function. If no value is specified, the digest currently associated with the image in the Oracle Cloud Infrastructure Registry will be used. Example: `sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7`
	* `pbf_listing_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PbfListing this function is sourced from.
		* `runtime_config` - FunctionsRuntime configuration for a function.
			* `functions_runtime_name` - The name of the FunctionsRuntime this function is to be associated with.
			* `functions_runtime_version_id` - The OCID of the FunctionsRuntimeVersion that is currently in use for the function.
			* `runtime_config_type` - Type of the FunctionsRuntime Config. Possible values: FUNCTION_UPDATE and MANUAL.
		* `source_code_sha256` - The SHA256 hash of the function source code archive, base64-encoded.
		* `source_type` - Type of the Function Source. Possible values: CONTAINER_IMAGE, PRE_BUILT_FUNCTIONS and ARCHIVE.
* `state` - The current state of the function.
* `success_destination` - An object that represents the destination to which Oracle Functions will send an invocation record with the details of the successful detached function invocation. A stream is an example of a success destination.  Example: `{"kind": "STREAM", "streamId": "stream_OCID"}`
	* `channel_id` - The ID of the channel in the queue.
	* `kind` - The type of destination for the response to a successful detached function invocation.
	* `queue_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the queue.
	* `stream_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream.
	* `topic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
* `time_created` - The time the function was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-09-12T22:47:12.613Z`
* `time_updated` - The time the function was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-09-12T22:47:12.613Z`
* `timeout_in_seconds` - Timeout for executions of the function. Value in seconds.
* `trace_config` - Define the tracing configuration for a function.
	* `is_enabled` - Define if tracing is enabled for the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Function
	* `update` - (Defaults to 20 minutes), when updating the Function
	* `delete` - (Defaults to 20 minutes), when destroying the Function


## Import

Functions can be imported using the `id`, e.g.

```
$ terraform import oci_functions_function.test_function "id"
```
