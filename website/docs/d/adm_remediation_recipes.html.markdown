---
subcategory: "Adm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_adm_remediation_recipes"
sidebar_current: "docs-oci-datasource-adm-remediation_recipes"
description: |-
  Provides the list of Remediation Recipes in Oracle Cloud Infrastructure Adm service
---

# Data Source: oci_adm_remediation_recipes
This data source provides the list of Remediation Recipes in Oracle Cloud Infrastructure Adm service.

Returns a list of Remediation Recipes based on the specified query parameters.
The query parameters `compartmentId` or `id` must be provided.


## Example Usage

```hcl
data "oci_adm_remediation_recipes" "test_remediation_recipes" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.remediation_recipe_display_name
	id = var.remediation_recipe_id
	state = var.remediation_recipe_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) A filter to return only resources that belong to the specified compartment identifier. Required only if the id query param is not specified. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) A filter to return only resources that match the specified identifier. Required only if the compartmentId query parameter is not specified. 
* `state` - (Optional) A filter to return only Remediation Recipes that match the specified lifecycleState.


## Attributes Reference

The following attributes are exported:

* `remediation_recipe_collection` - The list of remediation_recipe_collection.

### RemediationRecipe Reference

The following attributes are exported:

* `compartment_id` - The compartment Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the remediation recipe.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `detect_configuration` - A configuration to define the constraints when detecting vulnerable dependencies. 
	* `exclusions` - The list of dependencies to be ignored by the recommendation algorithm. The dependency pattern is matched against the 'group:artifact:version' or the purl of a dependency. An asterisk (*) at the end in the dependency pattern acts as a wildcard and matches zero or more characters. 
	* `max_permissible_cvss_v2score` - The maximum Common Vulnerability Scoring System Version 2 (CVSS V2) score. An artifact with a CVSS V2 score below this value is not considered for patching.
	* `max_permissible_cvss_v3score` - The maximum Common Vulnerability Scoring System Version 3 (CVSS V3) score. An artifact with a CVSS V3 score below this value is not considered for patching.
	* `max_permissible_severity` - The maximum ADM Severity. An artifact with an ADM Severity below this value is not considered for patching.
	* `upgrade_policy` - The upgrade policy for recommendations. The `Nearest` upgrade policy upgrades a dependency to the oldest version that meets both of the following criteria: it is newer than the current version and it is not affected by a vulnerability. 
* `display_name` - The name of the Remediation Recipe.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the remediation recipe.
* `is_run_triggered_on_kb_change` - Boolean indicating if a run should be automatically triggered once the Knowledge Base contents are updated.
* `knowledge_base_id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the knowledge base.
* `network_configuration` - A network configuration defines the required network characteristics for an ADM remediation recipe. A network configuration is required if the build service is one of: GitHub Actions, GitLab Pipeline, or Jenkins Pipeline. 
	* `nsg_ids` - The list of Oracle Cloud Identifiers ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) corresponding to Network Security Groups.
	* `subnet_id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the subnet.
* `scm_configuration` - A configuration for the Source Code Management tool/platform used by a remediation recipe.
	* `branch` - The branch used by ADM to patch vulnerabilities.
	* `build_file_location` - The location of the build file relative to the root of the repository. Only Maven build files (POM) are currently supported. If this property is not specified, ADM will use the build file located at the root of the repository. 
	* `external_scm_type` - The type of External Source Code Management.
	* `is_automerge_enabled` - If true, the Pull Request (PR) will be merged after the verify stage completes successfully     If false, the PR with the proposed changes must be reviewed and manually merged. 
	* `oci_code_repository_id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Oracle Cloud Infrastructure DevOps repository.
	* `pat_secret_id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Private Access Token (PAT) Secret. The secret provides the credentials necessary to authenticate against the SCM. 
	* `repository_url` - The repository URL for the SCM. For Non-Enterprise GitHub the expected format is https://github.com/[owner]/[repoName] For Enterprise GitHub the expected format is http(s)://[hostname]/api/v3/repos/[owner]/[repoName] For GitLab the expected format is https://gitlab.com/[groupName]/[repoName] 
	* `scm_type` - The type of Source Code Management.
	* `username` - The username for the SCM (to perform operations such as cloning or pushing via HTTP).
* `state` - The current lifecycle state of the Remediation Recipe.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The creation date and time of the Remediation Recipe (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_updated` - The date and time the Remediation Recipe was last updated (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `verify_configuration` - The Verify stage configuration specifies a build service to run a pipeline for the recommended code changes. The build pipeline will be initiated to ensure that there is no breaking change after the dependency versions have been updated in source to avoid vulnerabilities. 
	* `additional_parameters` - Additional key-value pairs passed as parameters to the build service when running an experiment.
	* `build_service_type` - The type of Build Service.
	* `jenkins_url` - The URL that locates the Jenkins pipeline.
	* `job_name` - The name of the Jenkins pipeline job that identifies the build pipeline.
	* `pat_secret_id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Private Access Token (PAT) Secret. The PAT provides the credentials to access the Jenkins Pipeline. 
	* `pipeline_id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the user's DevOps Build Pipeline.
	* `repository_url` - The location of the repository where the GitHub Actions is defined. For Non-Enterprise GitHub the expected format is https://github.com/[owner]/[repoName] For Enterprise GitHub the expected format is http(s)://[hostname]/api/v3/repos/[owner]/[repoName] 
	* `trigger_secret_id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the trigger Secret. The Secret provides access to the trigger for a GitLab pipeline. 
	* `username` - The username that will be used to authenticate with Jenkins.
	* `workflow_name` - The name of the GitHub Actions workflow that defines the build pipeline.

