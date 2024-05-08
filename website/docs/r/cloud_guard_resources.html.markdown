```

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present
---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_resources"
sidebar_current: "docs-oci-datasource-cloud_guard-resources"
description: |-
  Provides the list of Resources in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_resources
This data source provides the list of Resources in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of all resources in a compartment

The ListResources operation returns only the resources in `compartmentId` passed.
The list does not include any subcompartments of the compartmentId passed.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListResources on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_cloud_guard_resources" "test_resources" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.resource_access_level
	compartment_id_in_subtree = var.resource_compartment_id_in_subtree
	cve_id = oci_cloud_guard_cve.test_cve.id
	cvss_score = var.resource_cvss_score
	cvss_score_greater_than = var.resource_cvss_score_greater_than
	cvss_score_less_than = var.resource_cvss_score_less_than
	detector_rule_id_list = var.resource_detector_rule_id_list
	detector_type = var.resource_detector_type
	region = var.resource_region
	risk_level = var.resource_risk_level
	risk_level_greater_than = var.resource_risk_level_greater_than
	risk_level_less_than = var.resource_risk_level_less_than
	target_id = oci_cloud_guard_target.test_target.id
}


## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to `RESTRICTED` permissions are checked and no partial results are displayed.
* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the setting of `accessLevel`.
* `cve_id` - (Optional) CVE ID associated with the resource.
* `cvss_score` - (Optional) Cvss score associated with the resource.
* `cvss_score_greater_than` - (Optional) Cvss score greater than associated with the resource.
* `cvss_score_less_than` - (Optional) Cvss score less than associated with the resource.
* `detector_rule_id_list` - (Optional) Comma seperated list of detector rule IDs to be passed in to match against Problems.
* `detector_type` - (Optional) The field to list the problems by detector type.
* `region` - (Optional) Oracle Cloud Infrastructure monitoring region.
* `risk_level` - (Optional) Risk level of the problem.
* `risk_level_greater_than` - (Optional) To filter risk level greater than the one mentioned in query param
* `risk_level_less_than` - (Optional) To filter risk level less than the one mentioned in query param
* `target_id` - (Optional) The ID of the target in which to list resources.


## Attributes Reference

The following attributes are exported:

* `resource_collection` - The list of resource_collection.

### Resource Reference

The following attributes are exported:

* `additional_details` - Optional details of a resource
    * `os_info` - Type of OS present in the resource
* `compartment_id` - CompartmentId of CG Resource
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

  Avoid entering confidential information.
* `id` - Ocid for CG resource
* `open_ports_count` - Number of open ports in a resource
* `problem_count` - Count of existing problems for a resource
* `region` - region of CG Resource
* `resource_name` - Name for the CG resource
* `resource_type` - resource type of the CG resource
* `risk_level` - The Risk Level
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `target_id` - TargetId of CG Resource
* `target_name` - Target name for the CG Resource
* `time_first_monitored` - First monitored time
* `time_last_monitored` - Last monitored time
* `vulnerability_count` - Count of existing number of vulnerabilities in the resource

require('bitbucket/internal/layout/base/base').onReady({id : 936064, active: true, name : "kmujumda", slug : "kmujumda", displayName : "Kartik Mujumdar", avatarUrl : "https:\/\/bitbucket.oci.oraclecorp.com\/s\/910479833\/c68db0f\/j23pqb\/1.0\/_\/download\/resources\/com.atlassian.bitbucket.server.bitbucket-webpack-INTERNAL:avatar\/avatar\/default-avatar-48.png", emailAddress : "kartik.mujumdar@oracle.com", type : "NORMAL"}, "OCI Bitbucket" ); require('bitbucket/internal/widget/keyboard-shortcuts/keyboard-shortcuts').onReady();require('bitbucket/internal/layout/repository/repository').onReady({"slug":"terraform-provider","id":2913,"name":"terraform-provider","hierarchyId":"c5483866f10cfed4aba8","scmId":"git","state":"AVAILABLE","statusMessage":"Available","forkable":true,"project":{"key":"ORC","id":856,"name":"Resource Manager","description":"Resource Manager Service","public":false,"type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC"}]},"avatarUrl":"/projects/ORC/avatar.png?s=64&v=1498072335242"},"public":false,"links":{"clone":[{"href":"ssh://git@bitbucket.oci.oraclecorp.com:7999/orc/terraform-provider.git","name":"ssh"}],"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC/repos/terraform-provider/browse"}]}}, '#clone-repo-button');define('@bitbucket/apps/pull-requests/initial-data', {currentUser: {"name":"kmujumda","emailAddress":"kartik.mujumdar@oracle.com","id":936064,"displayName":"Kartik Mujumdar","active":true,"slug":"kmujumda","type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/users/kmujumda"}]},"avatarUrl":"https://bitbucket.oci.oraclecorp.com/s/910479833/c68db0f/j23pqb/1.0/_/download/resources/com.atlassian.bitbucket.server.bitbucket-webpack-INTERNAL:avatar/avatar/default-avatar-64.png"}, repository: {"slug":"terraform-provider","id":2913,"name":"terraform-provider","hierarchyId":"c5483866f10cfed4aba8","scmId":"git","state":"AVAILABLE","statusMessage":"Available","forkable":true,"project":{"key":"ORC","id":856,"name":"Resource Manager","description":"Resource Manager Service","public":false,"type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC"}]},"avatarUrl":"/projects/ORC/avatar.png?s=64&v=1498072335242"},"public":false,"links":{"clone":[{"href":"ssh://git@bitbucket.oci.oraclecorp.com:7999/orc/terraform-provider.git","name":"ssh"}],"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC/repos/terraform-provider/browse"}]}},pullRequest: {"id":5947,"version":77,"title":"TERSI-2299:- Added - Support for CloudGuard - Container Workload Integration","state":"OPEN","open":true,"closed":false,"createdDate":1712045009657,"updatedDate":1715067117444,"fromRef":{"id":"refs/heads/dev_provider_dexreq_5821","displayId":"dev_provider_dexreq_5821","latestCommit":"928efbf96137f3d71016502a5c316ad7442b9a46","type":"BRANCH","repository":{"slug":"terraform-provider","id":2913,"name":"terraform-provider","hierarchyId":"c5483866f10cfed4aba8","scmId":"git","state":"AVAILABLE","statusMessage":"Available","forkable":true,"project":{"key":"ORC","id":856,"name":"Resource Manager","description":"Resource Manager Service","public":false,"type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC"}]},"avatarUrl":"/projects/ORC/avatar.png?s=48&v=1498072335242"},"public":false,"links":{"clone":[{"href":"ssh://git@bitbucket.oci.oraclecorp.com:7999/orc/terraform-provider.git","name":"ssh"}],"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC/repos/terraform-provider/browse"}]}}},"toRef":{"id":"refs/heads/provider_preview","displayId":"provider_preview","latestCommit":"7a85e0f934b4d1e45e2f3db1fac52c2efc8f40c3","type":"BRANCH","repository":{"slug":"terraform-provider","id":2913,"name":"terraform-provider","hierarchyId":"c5483866f10cfed4aba8","scmId":"git","state":"AVAILABLE","statusMessage":"Available","forkable":true,"project":{"key":"ORC","id":856,"name":"Resource Manager","description":"Resource Manager Service","public":false,"type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC"}]},"avatarUrl":"/projects/ORC/avatar.png?s=48&v=1498072335242"},"public":false,"links":{"clone":[{"href":"ssh://git@bitbucket.oci.oraclecorp.com:7999/orc/terraform-provider.git","name":"ssh"}],"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC/repos/terraform-provider/browse"}]}}},"locked":false,"author":{"user":{"name":"gear-TER","emailAddress":"no-reply@oracle.com","id":949936,"displayName":"Terraform Team Automation","active":true,"slug":"gear-ter","type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/users/gear-ter"}]},"avatarUrl":"https://bitbucket.oci.oraclecorp.com/s/910479833/c68db0f/j23pqb/1.0/_/download/resources/com.atlassian.bitbucket.server.bitbucket-webpack-INTERNAL:avatar/avatar/default-avatar-48.png"},"role":"AUTHOR","approved":false,"status":"UNAPPROVED"},"reviewers":[],"participants":[{"user":{"name":"3adc96b3b93443e4","id":930202,"displayName":"Access Key User - Terraform_TC_Key","active":true,"slug":"3adc96b3b93443e4","type":"SERVICE","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/bots/3adc96b3b93443e4"}]},"avatarUrl":"https://bitbucket.oci.oraclecorp.com/s/910479833/c68db0f/j23pqb/1.0/_/download/resources/com.atlassian.bitbucket.server.bitbucket-webpack-INTERNAL:avatar/avatar/default-avatar-48.png"},"role":"PARTICIPANT","approved":false,"status":"UNAPPROVED"},{"user":{"name":"kmujumda","emailAddress":"kartik.mujumdar@oracle.com","id":936064,"displayName":"Kartik Mujumdar","active":true,"slug":"kmujumda","type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/users/kmujumda"}]},"avatarUrl":"https://bitbucket.oci.oraclecorp.com/s/910479833/c68db0f/j23pqb/1.0/_/download/resources/com.atlassian.bitbucket.server.bitbucket-webpack-INTERNAL:avatar/avatar/default-avatar-48.png"},"role":"PARTICIPANT","approved":false,"status":"UNAPPROVED"}],"links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC/repos/terraform-provider/pull-requests/5947"}]},"descriptionAsHtml":null},relevantContextLines: 10, userAttributes: {"canDelete":false,"canWrite":true,"canReadSourceRepo":true,"isWatching":true,"canWriteSourceRepo":true},});require('bitbucket/internal/layout/base/menu/repositories/recent').initMenu('repositories-menu-trigger'); require('bitbucket-plugin-awesome-graphs/recent-people').initMenu('#people-menu-trigger'); Jira Issuesjira-issues-listClose
```