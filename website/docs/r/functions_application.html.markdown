---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_application"
sidebar_current: "docs-oci-resource-functions-application"
description: |-
  Provides the Application resource in Oracle Cloud Infrastructure Functions service
---

# oci_functions_application
This resource provides the Application resource in Oracle Cloud Infrastructure Functions service.

Creates a new application.

**NOTE:** Terraform will take 5 minutes after destroying an application due to a known service issue. Refer [here](https://docs.cloud.oracle.com/iaas/Content/Functions/Tasks/functionsdeleting.htm)

## Example Usage

```hcl
resource "oci_functions_application" "test_application" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.application_display_name
	subnet_ids = var.application_subnet_ids

	#Optional
	config = var.application_config
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	network_security_group_ids = var.application_network_security_group_ids
	image_policy_config {
		#Required
		is_policy_enabled = var.application_image_policy_config_is_policy_enabled

		#Optional
		key_details {
			#Required
			kms_key_id = oci_kms_key.test_key.id
		}
	}
	syslog_url = var.application_syslog_url
	trace_config {

		#Optional
		domain_id = oci_functions_domain.test_domain.id
		is_enabled = var.application_trace_config_is_enabled
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment to create the application within. 
* `config` - (Optional) (Updatable) Application configuration. These values are passed on to the function as environment variables, functions may override application configuration. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`

    The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) The display name of the application. The display name must be unique within the compartment containing the application. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `network_security_group_ids` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the Network Security Groups to add the application to.
* `image_policy_config` - (Optional) (Updatable) Define the image signature verification policy for an application. 
    * `is_policy_enabled` - (Required) (Updatable) Define if image signature verification policy is enabled for the application. 
    * `key_details` - (Optional) (Updatable) A list of KMS key details.
        * `kms_key_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the KMS key that will be used to verify the image signature. 
* `subnet_ids` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the subnets in which to run functions in the application. 
* `syslog_url` - (Optional) (Updatable) A syslog URL to which to send all function logs. Supports tcp, udp, and tcp+tls. The syslog URL must be reachable from all of the subnets configured for the application. Note: If you enable the Oracle Cloud Infrastructure Logging service for this application, the syslogUrl value is ignored. Function logs are sent to the Oracle Cloud Infrastructure Logging service, and not to the syslog URL.  Example: `tcp://logserver.myserver:1234` 
* `trace_config` - (Optional) (Updatable) Define the tracing configuration for an application. 
    * `domain_id` - (Optional) (Updatable) The OCID of the collector (e.g. an APM Domain) trace events will be sent to.  
    * `is_enabled` - (Optional) (Updatable) Define if tracing is enabled for the resource. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the application. 
* `config` - Application configuration for functions in this application (passed as environment variables). Can be overridden by function configuration. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`

    The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of the application. The display name is unique within the compartment containing the application. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application. 
* `network_security_group_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the Network Security Groups to add the application to.
* `image_policy_config` - Define the image signature verification policy for an application. 
    * `is_policy_enabled` - Define if image signature verification policy is enabled for the application. 
    * `key_details` - A list of KMS key details.
        * `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the KMS key that will be used to verify the image signature. 
* `state` - The current state of the application. 
* `subnet_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the subnets in which to run functions in the application. 
* `syslog_url` - A syslog URL to which to send all function logs. Supports tcp, udp, and tcp+tls. The syslog URL must be reachable from all of the subnets configured for the application. Note: If you enable the Oracle Cloud Infrastructure Logging service for this application, the syslogUrl value is ignored. Function logs are sent to the Oracle Cloud Infrastructure Logging service, and not to the syslog URL.  Example: `tcp://logserver.myserver:1234` 
* `time_created` - The time the application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-09-12T22:47:12.613Z` 
* `time_updated` - The time the application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-09-12T22:47:12.613Z` 
* `trace_config` - Define the tracing configuration for an application. 
    * `domain_id` - The OCID of the collector (e.g. an APM Domain) trace events will be sent to.  
    * `is_enabled` - Define if tracing is enabled for the resource. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Application
	* `update` - (Defaults to 20 minutes), when updating the Application
	* `delete` - (Defaults to 20 minutes), when destroying the Application


## Import

Applications can be imported using the `id`, e.g.

```
$ terraform import oci_functions_application.test_application "id"
```

