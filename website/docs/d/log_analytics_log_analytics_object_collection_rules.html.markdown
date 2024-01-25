---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_object_collection_rules"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_object_collection_rules"
description: |-
  Provides the list of Log Analytics Object Collection Rules in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_object_collection_rules
This data source provides the list of Log Analytics Object Collection Rules in Oracle Cloud Infrastructure Log Analytics service.

Gets list of configuration details of Object Storage based collection rules.

## Example Usage

```hcl
data "oci_log_analytics_log_analytics_object_collection_rules" "test_log_analytics_object_collection_rules" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.log_analytics_object_collection_rule_namespace

	#Optional
	name = var.log_analytics_object_collection_rule_name
	state = var.log_analytics_object_collection_rule_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `name` - (Optional) A filter to return rules only matching with this name.
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `state` - (Optional) Lifecycle state filter. 


## Attributes Reference

The following attributes are exported:

* `log_analytics_object_collection_rule_collection` - The list of log_analytics_object_collection_rule_collection.

### LogAnalyticsObjectCollectionRule Reference

The following attributes are exported:

* `char_encoding` - An optional character encoding to aid in detecting the character encoding of the contents of the objects while processing. It is recommended to set this value as ISO_8589_1 when configuring content of the objects having more numeric characters, and very few alphabets. For e.g. this applies when configuring VCN Flow Logs. 
* `collection_type` - The type of collection. Supported collection types: LIVE, HISTORIC, HISTORIC_LIVE 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A string that describes the details of the rule. It does not have to be unique, and can be changed. Avoid entering confidential information. 
* `entity_id` - Logging Analytics entity OCID to associate the processed logs with.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this rule.
* `is_enabled` - Whether or not this rule is currently enabled. 
* `is_force_historic_collection` - Flag to allow historic collection if poll period overlaps with existing ACTIVE collection rule 
* `lifecycle_details` - A detailed status of the life cycle state.
* `log_group_id` - Logging Analytics Log group OCID to associate the processed logs with.
* `log_set` - The logSet to be associated with the processed logs. The logSet feature can be used by customers with high volume of data  and this feature has to be enabled for a given tenancy prior to its usage. When logSetExtRegex value is provided, it will take precedence over this logSet value and logSet will be computed dynamically  using logSetKey and logSetExtRegex. 
* `log_set_ext_regex` - The regex to be applied against given logSetKey. Regex has to be in string escaped format. 
* `log_set_key` - An optional parameter to indicate from where the logSet to be extracted using logSetExtRegex. Default value is OBJECT_PATH (e.g. /n/<namespace>/b/<bucketname>/o/<objectname>). 
* `log_source_name` - Name of the Logging Analytics Source to use for the processing.
* `log_type` - Type of files/objects in this object collection rule. 
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
* `timezone` - Timezone to be used when processing log entries whose timestamps do not include an explicit timezone.  When this property is not specified, the timezone of the entity specified is used.  If the entity is also not specified or do not have a valid timezone then UTC is used. 

