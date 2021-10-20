---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_object_collection_rule"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_object_collection_rule"
description: |-
  Provides details about a specific Log Analytics Object Collection Rule in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_object_collection_rule
This data source provides details about a specific Log Analytics Object Collection Rule resource in Oracle Cloud Infrastructure Log Analytics service.

Gets a configured object storage based collection rule by given id

## Example Usage

```hcl
data "oci_log_analytics_log_analytics_object_collection_rule" "test_log_analytics_object_collection_rule" {
	#Required
	log_analytics_object_collection_rule_id = oci_log_analytics_log_analytics_object_collection_rule.test_log_analytics_object_collection_rule.id
	namespace = var.log_analytics_object_collection_rule_namespace
}
```

## Argument Reference

The following arguments are supported:

* `log_analytics_object_collection_rule_id` - (Required) The Logging Analytics Object Collection Rule [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `char_encoding` - An optional character encoding to aid in detecting the character encoding of the contents of the objects while processing. It is recommended to set this value as ISO_8589_1 when configuring content of the objects having more numeric characters, and very few alphabets. For e.g. this applies when configuring VCN Flow Logs. 
* `collection_type` - The type of collection. Supported collection types: LIVE, HISTORIC, HISTORIC_LIVE 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A string that describes the details of the rule. It does not have to be unique, and can be changed. Avoid entering confidential information. 
* `entity_id` - Logging Analytics entity OCID to associate the processed logs with.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this rule.
* `lifecycle_details` - A detailed status of the life cycle state.
* `log_group_id` - Logging Analytics Log group OCID to associate the processed logs with.
* `log_source_name` - Name of the Logging Analytics Source to use for the processing.
* `name` - A unique name to the rule. The name must be unique, within the tenancy, and cannot be changed.
* `object_name_filters` - When the filters are provided, only the objects matching the filters are picked up for processing. The matchType supported is exact match and accommodates wildcard "*". For more information on filters, see [Event Filters](https://docs.oracle.com/en-us/iaas/Content/Events/Concepts/filterevents.htm). 
* `os_bucket_name` - Name of the Object Storage bucket.
* `os_namespace` - Object Storage namespace.
* `overrides` - Use this to override some property values which are defined at bucket level to the scope of object. Supported propeties for override are, logSourceName, charEncoding. Supported matchType for override are "contains". 
* `poll_since` - The oldest time of the file in the bucket to consider for collection. Accepted values are: BEGINNING or CURRENT_TIME or RFC3339 formatted datetime string. When collectionType is LIVE, specifying pollSince value other than CURRENT_TIME will result in error. 
* `poll_till` - The oldest time of the file in the bucket to consider for collection. Accepted values are: CURRENT_TIME or RFC3339 formatted datetime string. When collectionType is LIVE, specifying pollTill will result in error. 
* `state` - The current state of the rule. 
* `time_created` - The time when this rule was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when this rule was last updated. An RFC3339 formatted datetime string.

