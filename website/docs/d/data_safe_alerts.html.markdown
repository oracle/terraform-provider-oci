---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_alerts"
sidebar_current: "docs-oci-datasource-data_safe-alerts"
description: |-
  Provides the list of Alerts in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_alerts
This data source provides the list of Alerts in Oracle Cloud Infrastructure Data Safe service.

Gets a list of all alerts.


## Example Usage

```hcl
data "oci_data_safe_alerts" "test_alerts" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.alert_access_level
	compartment_id_in_subtree = var.alert_compartment_id_in_subtree
	field = var.alert_field
	id = var.alert_id
	scim_query = var.alert_scim_query
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
* `field` - (Optional) Specifies a subset of fields to be returned in the response.
* `id` - (Optional) A filter to return alert by it's OCID.
* `scim_query` - (Optional) The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2 of the System for Cross-Domain Identity Management (SCIM) specification, which is available at [RFC3339](https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions, text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format. (Numeric and boolean values should not be quoted.)

  **Example:** | query=(timeCreated ge '2021-06-04T01-00-26') and (targetNames eq 'target_1') query=(featureDetails.userName eq "user") and (targetNames eq "target_1") Supported fields: severity status alertType targetIds targetNames operationTime lifecycleState displayName timeCreated timeUpdated featureDetails.* (* can be any field in nestedStrMap in Feature Attributes in Alert Summary. For example -  userName,object,clientHostname,osUserName,clientIPs,clientId,commandText,commandParam,clientProgram,objectType,targetOwner)


## Attributes Reference

The following attributes are exported:

* `alert_collection` - The list of alert_collection.

### Alert Reference

The following attributes are exported:

* `alert_policy_rule_key` - The key of the rule of alert policy that triggered alert.
* `alert_policy_rule_name` - The display name of the rule of alert policy that triggered alert.
* `alert_type` - Type of the alert. Indicates the Data Safe feature triggering the alert.
* `comment` - A comment for the alert. Entered by the user.
* `compartment_id` - The OCID of the compartment that contains the alert.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
* `description` - The description of the alert.
* `display_name` - The display name of the alert.
* `feature_details` - Map that contains maps of values. Example: `{"Operations": {"CostCenter": "42"}}`
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
* `id` - The OCID of the alert.
* `operation` - The operation (event) that triggered alert.
* `operation_status` - The result of the operation (event) that triggered alert.
* `operation_time` - Creation date and time of the operation that triggered alert, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `policy_id` - The OCID of the policy that triggered alert.
* `resource_name` - The resource endpoint that triggered the alert.
* `severity` - Severity level of the alert.
* `state` - The current state of the alert.
* `status` - The status of the alert.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `target_ids` - Array of OCIDs of the target database which are associated with the alert.
* `target_names` - Array of names of the target database.
* `time_created` - Creation date and time of the alert, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - Last date and time the alert was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
