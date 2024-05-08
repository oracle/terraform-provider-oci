```

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present
---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_adhoc_query_results"
sidebar_current: "docs-oci-datasource-cloud_guard-adhoc_query_results"
description: |-
  Provides the list of Adhoc Query Results in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_adhoc_query_results
This data source provides the list of Adhoc Query Results in Oracle Cloud Infrastructure Cloud Guard service.

Lists the results for a given adhoc ID (from includes results from all monitoring regions).

## Example Usage

```hcl
data "oci_cloud_guard_adhoc_query_results" "test_adhoc_query_results" {
	#Required
	adhoc_query_id = oci_cloud_guard_adhoc_query.test_adhoc_query.id
	compartment_id = var.compartment_id
}


## Argument Reference

The following arguments are supported:

* `adhoc_query_id` - (Required) Adhoc query OCID.
* `compartment_id` - (Required) The OCID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `adhoc_query_result_collection` - The list of adhoc_query_result_collection.

### AdhocQueryResult Reference

The following attributes are exported:

* `items` - List of adhoc query results
    * `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
    * `error_message` - Optional error message
    * `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

      Avoid entering confidential information.
    * `host_id` - Resource this result belongs to
    * `region` - The region this adhoc work request is running in, needed for tracking when work request is synced to reporting region
    * `result` - Result of the adhoc query this result resource is associated with
    * `result_count` - Number of records returned for the query results on this host
    * `state` - Status of the query
    * `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
    * `time_submitted` - The time the adhoc result was submitted. An RFC3339 formatted datetime string

require('bitbucket/internal/layout/base/base').onReady({id : 936064, active: true, name : "kmujumda", slug : "kmujumda", displayName : "Kartik Mujumdar", avatarUrl : "https:\/\/bitbucket.oci.oraclecorp.com\/s\/910479833\/c68db0f\/j23pqb\/1.0\/_\/download\/resources\/com.atlassian.bitbucket.server.bitbucket-webpack-INTERNAL:avatar\/avatar\/default-avatar-48.png", emailAddress : "kartik.mujumdar@oracle.com", type : "NORMAL"}, "OCI Bitbucket" ); require('bitbucket/internal/widget/keyboard-shortcuts/keyboard-shortcuts').onReady();require('bitbucket/internal/layout/repository/repository').onReady({"slug":"terraform-provider","id":2913,"name":"terraform-provider","hierarchyId":"c5483866f10cfed4aba8","scmId":"git","state":"AVAILABLE","statusMessage":"Available","forkable":true,"project":{"key":"ORC","id":856,"name":"Resource Manager","description":"Resource Manager Service","public":false,"type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC"}]},"avatarUrl":"/projects/ORC/avatar.png?s=64&v=1498072335242"},"public":false,"links":{"clone":[{"href":"ssh://git@bitbucket.oci.oraclecorp.com:7999/orc/terraform-provider.git","name":"ssh"}],"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC/repos/terraform-provider/browse"}]}}, '#clone-repo-button');define('@bitbucket/apps/pull-requests/initial-data', {currentUser: {"name":"kmujumda","emailAddress":"kartik.mujumdar@oracle.com","id":936064,"displayName":"Kartik Mujumdar","active":true,"slug":"kmujumda","type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/users/kmujumda"}]},"avatarUrl":"https://bitbucket.oci.oraclecorp.com/s/910479833/c68db0f/j23pqb/1.0/_/download/resources/com.atlassian.bitbucket.server.bitbucket-webpack-INTERNAL:avatar/avatar/default-avatar-64.png"}, repository: {"slug":"terraform-provider","id":2913,"name":"terraform-provider","hierarchyId":"c5483866f10cfed4aba8","scmId":"git","state":"AVAILABLE","statusMessage":"Available","forkable":true,"project":{"key":"ORC","id":856,"name":"Resource Manager","description":"Resource Manager Service","public":false,"type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC"}]},"avatarUrl":"/projects/ORC/avatar.png?s=64&v=1498072335242"},"public":false,"links":{"clone":[{"href":"ssh://git@bitbucket.oci.oraclecorp.com:7999/orc/terraform-provider.git","name":"ssh"}],"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC/repos/terraform-provider/browse"}]}},pullRequest: {"id":5947,"version":77,"title":"TERSI-2299:- Added - Support for CloudGuard - Container Workload Integration","state":"OPEN","open":true,"closed":false,"createdDate":1712045009657,"updatedDate":1715067117444,"fromRef":{"id":"refs/heads/dev_provider_dexreq_5821","displayId":"dev_provider_dexreq_5821","latestCommit":"928efbf96137f3d71016502a5c316ad7442b9a46","type":"BRANCH","repository":{"slug":"terraform-provider","id":2913,"name":"terraform-provider","hierarchyId":"c5483866f10cfed4aba8","scmId":"git","state":"AVAILABLE","statusMessage":"Available","forkable":true,"project":{"key":"ORC","id":856,"name":"Resource Manager","description":"Resource Manager Service","public":false,"type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC"}]},"avatarUrl":"/projects/ORC/avatar.png?s=48&v=1498072335242"},"public":false,"links":{"clone":[{"href":"ssh://git@bitbucket.oci.oraclecorp.com:7999/orc/terraform-provider.git","name":"ssh"}],"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC/repos/terraform-provider/browse"}]}}},"toRef":{"id":"refs/heads/provider_preview","displayId":"provider_preview","latestCommit":"7a85e0f934b4d1e45e2f3db1fac52c2efc8f40c3","type":"BRANCH","repository":{"slug":"terraform-provider","id":2913,"name":"terraform-provider","hierarchyId":"c5483866f10cfed4aba8","scmId":"git","state":"AVAILABLE","statusMessage":"Available","forkable":true,"project":{"key":"ORC","id":856,"name":"Resource Manager","description":"Resource Manager Service","public":false,"type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC"}]},"avatarUrl":"/projects/ORC/avatar.png?s=48&v=1498072335242"},"public":false,"links":{"clone":[{"href":"ssh://git@bitbucket.oci.oraclecorp.com:7999/orc/terraform-provider.git","name":"ssh"}],"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC/repos/terraform-provider/browse"}]}}},"locked":false,"author":{"user":{"name":"gear-TER","emailAddress":"no-reply@oracle.com","id":949936,"displayName":"Terraform Team Automation","active":true,"slug":"gear-ter","type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/users/gear-ter"}]},"avatarUrl":"https://bitbucket.oci.oraclecorp.com/s/910479833/c68db0f/j23pqb/1.0/_/download/resources/com.atlassian.bitbucket.server.bitbucket-webpack-INTERNAL:avatar/avatar/default-avatar-48.png"},"role":"AUTHOR","approved":false,"status":"UNAPPROVED"},"reviewers":[],"participants":[{"user":{"name":"3adc96b3b93443e4","id":930202,"displayName":"Access Key User - Terraform_TC_Key","active":true,"slug":"3adc96b3b93443e4","type":"SERVICE","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/bots/3adc96b3b93443e4"}]},"avatarUrl":"https://bitbucket.oci.oraclecorp.com/s/910479833/c68db0f/j23pqb/1.0/_/download/resources/com.atlassian.bitbucket.server.bitbucket-webpack-INTERNAL:avatar/avatar/default-avatar-48.png"},"role":"PARTICIPANT","approved":false,"status":"UNAPPROVED"},{"user":{"name":"kmujumda","emailAddress":"kartik.mujumdar@oracle.com","id":936064,"displayName":"Kartik Mujumdar","active":true,"slug":"kmujumda","type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/users/kmujumda"}]},"avatarUrl":"https://bitbucket.oci.oraclecorp.com/s/910479833/c68db0f/j23pqb/1.0/_/download/resources/com.atlassian.bitbucket.server.bitbucket-webpack-INTERNAL:avatar/avatar/default-avatar-48.png"},"role":"PARTICIPANT","approved":false,"status":"UNAPPROVED"}],"links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/projects/ORC/repos/terraform-provider/pull-requests/5947"}]},"descriptionAsHtml":null},relevantContextLines: 10, userAttributes: {"canDelete":false,"canWrite":true,"canReadSourceRepo":true,"isWatching":true,"canWriteSourceRepo":true},});require('bitbucket/internal/layout/base/menu/repositories/recent').initMenu('repositories-menu-trigger'); require('bitbucket-plugin-awesome-graphs/recent-people').initMenu('#people-menu-trigger'); Jira Issuesjira-issues-listClose
```