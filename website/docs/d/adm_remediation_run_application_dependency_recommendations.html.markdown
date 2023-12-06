---
subcategory: "Adm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_adm_remediation_run_application_dependency_recommendations"
sidebar_current: "docs-oci-datasource-adm-remediation_run_application_dependency_recommendations"
description: |-
  Provides the list of Remediation Run Application Dependency Recommendations in Oracle Cloud Infrastructure Adm service
---

# Data Source: oci_adm_remediation_run_application_dependency_recommendations
This data source provides the list of Remediation Run Application Dependency Recommendations in Oracle Cloud Infrastructure Adm service.

Returns a list of application dependency with their associated recommendations.

## Example Usage

```hcl
data "oci_adm_remediation_run_application_dependency_recommendations" "test_remediation_run_application_dependency_recommendations" {
	#Required
	remediation_run_id = oci_adm_remediation_run.test_remediation_run.id

	#Optional
	gav = var.remediation_run_application_dependency_recommendation_gav
}
```

## Argument Reference

The following arguments are supported:

* `gav` - (Optional) A filter to return only resources that match the entire GAV (Group Artifact Version) identifier given.
* `remediation_run_id` - (Required) Unique Remediation Run identifier path parameter.


## Attributes Reference

The following attributes are exported:

* `application_dependency_recommendation_collection` - The list of application_dependency_recommendation_collection.

### RemediationRunApplicationDependencyRecommendation Reference

The following attributes are exported:

* `items` - List of application recommendation summaries.
	* `application_dependency_node_ids` - List of (application dependencies) node identifiers from which this node depends.
	* `gav` - Unique Group Artifact Version (GAV) identifier in the format _Group:Artifact:Version_, e.g. org.graalvm.nativeimage:svm:21.1.0.
	* `node_id` - Unique node identifier of an application dependency with an associated Recommendation, e.g. nodeId1.
	* `recommended_gav` - Recommended application dependency in "group:artifact:version" (GAV) format, e.g. org.graalvm.nativeimage:svm:21.2.0.

