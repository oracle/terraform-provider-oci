---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_security_feature_analytics"
sidebar_current: "docs-oci-datasource-data_safe-security_assessment_security_feature_analytics"
description: |-
  Provides the list of Security Assessment Security Feature Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessment_security_feature_analytics
This data source provides the list of Security Assessment Security Feature Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets a list of Database security feature usage aggregated details in the specified compartment. This provides information about the
overall security controls, by returning the counting number of the target databases using the security features.

When you perform the ListSecurityFeatureAnalytics operation, if the parameter compartmentIdInSubtree is set to "true," and if the
parameter accessLevel is set to ACCESSIBLE, then the operation returns statistics from the compartments in which the requestor has INSPECT
permissions on at least one resource, directly or indirectly (in subcompartments). If the operation is performed at the
root compartment and the requestor does not have access to at least one subcompartment of the compartment specified by
compartmentId, then "Not Authorized" is returned.


## Example Usage

```hcl
data "oci_data_safe_security_assessment_security_feature_analytics" "test_security_assessment_security_feature_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.security_assessment_security_feature_analytic_access_level
	compartment_id_in_subtree = var.security_assessment_security_feature_analytic_compartment_id_in_subtree
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `security_feature_analytics_collection` - The list of security_feature_analytics_collection.

### SecurityAssessmentSecurityFeatureAnalytic Reference

The following attributes are exported:

* `items` - The array of database security feature analytics summary.
	* `dimensions` - The scope of analytics data.
		* `security_feature` - The name of the security feature.
	* `metric_name` - The name of the aggregation metric.
	* `security_assessment_security_feature_analytic_count` - The total count for the aggregation metric.

