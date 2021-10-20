---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_object_collection_rule"
sidebar_current: "docs-oci-resource-log_analytics-log_analytics_object_collection_rule"
description: |-
  Provides the Log Analytics Object Collection Rule resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_log_analytics_object_collection_rule
This resource provides the Log Analytics Object Collection Rule resource in Oracle Cloud Infrastructure Log Analytics service.

Create a configuration to collect logs from object storage bucket.

## Example Usage

```hcl
resource "oci_log_analytics_log_analytics_object_collection_rule" "test_log_analytics_object_collection_rule" {
	#Required
	compartment_id = var.compartment_id
	log_group_id = oci_logging_log_group.test_log_group.id
	log_source_name = var.log_analytics_object_collection_rule_log_source_name
	name = var.log_analytics_object_collection_rule_name
	namespace = var.log_analytics_object_collection_rule_namespace
	os_bucket_name = oci_objectstorage_bucket.test_bucket.name
	os_namespace = var.log_analytics_object_collection_rule_os_namespace

	#Optional
	char_encoding = var.log_analytics_object_collection_rule_char_encoding
	collection_type = var.log_analytics_object_collection_rule_collection_type
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.log_analytics_object_collection_rule_description
	entity_id = oci_log_analytics_entity.test_entity.id
	freeform_tags = {"bar-key"= "value"}
	object_name_filters = var.log_analytics_object_collection_rule_object_name_filters
	overrides = var.log_analytics_object_collection_rule_overrides
	poll_since = var.log_analytics_object_collection_rule_poll_since
	poll_till = var.log_analytics_object_collection_rule_poll_till
}
```

## Argument Reference

The following arguments are supported:

* `char_encoding` - (Optional) (Updatable) An optional character encoding to aid in detecting the character encoding of the contents of the objects while processing. It is recommended to set this value as ISO_8589_1 when configuring content of the objects having more numeric characters, and very few alphabets. For e.g. this applies when configuring VCN Flow Logs. 
* `collection_type` - (Optional) The type of collection. Supported collection types: LIVE, HISTORIC, HISTORIC_LIVE 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A string that describes the details of the rule. It does not have to be unique, and can be changed. Avoid entering confidential information. 
* `entity_id` - (Optional) (Updatable) Logging Analytics entity OCID. Associates the processed logs with the given entity (optional).
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `log_group_id` - (Required) (Updatable) Logging Analytics Log group OCID to associate the processed logs with.
* `log_source_name` - (Required) (Updatable) Name of the Logging Analytics Source to use for the processing.
* `name` - (Required) A unique name given to the rule. The name must be unique within the tenancy, and cannot be modified.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `object_name_filters` - (Optional) (Updatable) When the filters are provided, only the objects matching the filters are picked up for processing. The matchType supported is exact match and accommodates wildcard "*". For more information on filters, see [Event Filters](https://docs.oracle.com/en-us/iaas/Content/Events/Concepts/filterevents.htm). 
* `os_bucket_name` - (Required) Name of the Object Storage bucket.
* `os_namespace` - (Required) Object Storage namespace.
* `overrides` - (Optional) (Updatable) The override is used to modify some important configuration properties for objects matching a specific pattern inside the bucket. Supported propeties for override are - logSourceName, charEncoding. Supported matchType for override are "contains". 
* `poll_since` - (Optional) The oldest time of the file in the bucket to consider for collection. Accepted values are: BEGINNING or CURRENT_TIME or RFC3339 formatted datetime string. When collectionType is LIVE, specifying pollSince value other than CURRENT_TIME will result in error. 
* `poll_till` - (Optional) The oldest time of the file in the bucket to consider for collection. Accepted values are: CURRENT_TIME or RFC3339 formatted datetime string. When collectionType is LIVE, specifying pollTill will result in error. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Import

LogAnalyticsObjectCollectionRules can be imported using the `id`, e.g.

```
$ terraform import oci_log_analytics_log_analytics_object_collection_rule.test_log_analytics_object_collection_rule "namespaces/{namespaceName}/logAnalyticsObjectCollectionRules/{logAnalyticsObjectCollectionRuleId}" 
```

