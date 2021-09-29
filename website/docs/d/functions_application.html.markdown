---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_application"
sidebar_current: "docs-oci-datasource-functions-application"
description: |-
  Provides details about a specific Application in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_application
This data source provides details about a specific Application resource in Oracle Cloud Infrastructure Functions service.

Retrieves an application.

## Example Usage

```hcl
data "oci_functions_application" "test_application" {
	#Required
	application_id = oci_functions_application.test_application.id
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this application. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the application. 
* `config` - Application configuration for functions in this application (passed as environment variables). Can be overridden by function configuration. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`

	The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of the application. The display name is unique within the compartment containing the application. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application. 
* `image_policy_config` - Define the image signature verification policy for an application. 
	* `is_policy_enabled` - Define if image signature verification policy is enabled for the application. 
	* `key_details` - A list of KMS key details.
		* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the KMS key that will be used to verify the image signature. 
* `network_security_group_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the Network Security Groups to add the application to. 
* `state` - The current state of the application. 
* `subnet_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the subnets in which to run functions in the application. 
* `syslog_url` - A syslog URL to which to send all function logs. Supports tcp, udp, and tcp+tls. The syslog URL must be reachable from all of the subnets configured for the application. Note: If you enable the Oracle Cloud Infrastructure Logging service for this application, the syslogUrl value is ignored. Function logs are sent to the Oracle Cloud Infrastructure Logging service, and not to the syslog URL.  Example: `tcp://logserver.myserver:1234` 
* `time_created` - The time the application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-09-12T22:47:12.613Z` 
* `time_updated` - The time the application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-09-12T22:47:12.613Z` 
* `trace_config` - Define the tracing configuration for an application. 
	* `domain_id` - The OCID of the collector (e.g. an APM Domain) trace events will be sent to.  
	* `is_enabled` - Define if tracing is enabled for the resource. 

