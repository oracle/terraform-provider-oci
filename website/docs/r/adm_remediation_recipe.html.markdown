---
subcategory: "Adm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_adm_remediation_recipe"
sidebar_current: "docs-oci-resource-adm-remediation_recipe"
description: |-
  Provides the Remediation Recipe resource in Oracle Cloud Infrastructure Adm service
---

# oci_adm_remediation_recipe
This resource provides the Remediation Recipe resource in Oracle Cloud Infrastructure Adm service.

Creates a new Remediation Recipe.

## Example Usage

```hcl
resource "oci_adm_remediation_recipe" "test_remediation_recipe" {
	#Required
	compartment_id = var.compartment_id
	detect_configuration {

		#Optional
		exclusions = var.remediation_recipe_detect_configuration_exclusions
		max_permissible_cvss_v2score = var.remediation_recipe_detect_configuration_max_permissible_cvss_v2score
		max_permissible_cvss_v3score = var.remediation_recipe_detect_configuration_max_permissible_cvss_v3score
		max_permissible_severity = var.remediation_recipe_detect_configuration_max_permissible_severity
		upgrade_policy = var.remediation_recipe_detect_configuration_upgrade_policy
	}
	is_run_triggered_on_kb_change = var.remediation_recipe_is_run_triggered_on_kb_change
	knowledge_base_id = oci_adm_knowledge_base.test_knowledge_base.id
	network_configuration {
		#Required
		subnet_id = oci_core_subnet.test_subnet.id

		#Optional
		nsg_ids = var.remediation_recipe_network_configuration_nsg_ids
	}
	scm_configuration {
		#Required
		branch = var.remediation_recipe_scm_configuration_branch
		is_automerge_enabled = var.remediation_recipe_scm_configuration_is_automerge_enabled
		scm_type = var.remediation_recipe_scm_configuration_scm_type

		#Optional
		build_file_location = var.remediation_recipe_scm_configuration_build_file_location
		external_scm_type = var.remediation_recipe_scm_configuration_external_scm_type
		oci_code_repository_id = oci_artifacts_repository.test_repository.id
		pat_secret_id = oci_vault_secret.test_secret.id
		repository_url = var.remediation_recipe_scm_configuration_repository_url
		username = var.remediation_recipe_scm_configuration_username
	}
	verify_configuration {
		#Required
		build_service_type = var.remediation_recipe_verify_configuration_build_service_type

		#Optional
		additional_parameters = var.remediation_recipe_verify_configuration_additional_parameters
		jenkins_url = var.remediation_recipe_verify_configuration_jenkins_url
		job_name = oci_database_migration_job.test_job.name
		pat_secret_id = oci_vault_secret.test_secret.id
		pipeline_id = oci_datascience_pipeline.test_pipeline.id
		repository_url = var.remediation_recipe_verify_configuration_repository_url
		trigger_secret_id = oci_vault_secret.test_secret.id
		username = var.remediation_recipe_verify_configuration_username
		workflow_name = var.remediation_recipe_verify_configuration_workflow_name
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.remediation_recipe_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the remediation recipe.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `detect_configuration` - (Required) (Updatable) A configuration to define the constraints when detecting vulnerable dependencies. 
	* `exclusions` - (Optional) (Updatable) The list of dependencies to be ignored by the recommendation algorithm. The dependency pattern is matched against the 'group:artifact:version' or the purl of a dependency. An asterisk (*) at the end in the dependency pattern acts as a wildcard and matches zero or more characters. 
	* `max_permissible_cvss_v2score` - (Optional) (Updatable) The maximum Common Vulnerability Scoring System Version 2 (CVSS V2) score. An artifact with a CVSS V2 score below this value is not considered for patching.
	* `max_permissible_cvss_v3score` - (Optional) (Updatable) The maximum Common Vulnerability Scoring System Version 3 (CVSS V3) score. An artifact with a CVSS V3 score below this value is not considered for patching.
	* `max_permissible_severity` - (Optional) (Updatable) The maximum ADM Severity. An artifact with an ADM Severity below this value is not considered for patching.
	* `upgrade_policy` - (Optional) (Updatable) The upgrade policy for recommendations. The `Nearest` upgrade policy upgrades a dependency to the oldest version that meets both of the following criteria: it is newer than the current version and it is not affected by a vulnerability. 
* `display_name` - (Optional) (Updatable) The name of the remediation recipe.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_run_triggered_on_kb_change` - (Required) (Updatable) Boolean indicating if a run should be automatically triggered once the knowledge base is updated.
* `knowledge_base_id` - (Required) (Updatable) The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the knowledge base.
* `network_configuration` - (Required) (Updatable) A network configuration defines the required network characteristics for an ADM remediation recipe. A network configuration is required if the build service is one of: GitHub Actions, GitLab Pipeline, or Jenkins Pipeline. 
	* `nsg_ids` - (Optional) (Updatable) The list of Oracle Cloud Identifiers ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) corresponding to Network Security Groups.
	* `subnet_id` - (Required) (Updatable) The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the subnet.
* `scm_configuration` - (Required) (Updatable) A configuration for the Source Code Management tool/platform used by a remediation recipe.
	* `branch` - (Required) (Updatable) The branch used by ADM to patch vulnerabilities.
	* `build_file_location` - (Optional) (Updatable) The location of the build file relative to the root of the repository. Only Maven build files (POM) are currently supported. If this property is not specified, ADM will use the build file located at the root of the repository. 
	* `external_scm_type` - (Required when scm_type=EXTERNAL_SCM) (Updatable) The type of External Source Code Management.
	* `is_automerge_enabled` - (Required) (Updatable) If true, the Pull Request (PR) will be merged after the verify stage completes successfully     If false, the PR with the proposed changes must be reviewed and manually merged. 
	* `oci_code_repository_id` - (Required when scm_type=OCI_CODE_REPOSITORY) (Updatable) The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Oracle Cloud Infrastructure DevOps repository.
	* `pat_secret_id` - (Required when scm_type=EXTERNAL_SCM) (Updatable) The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Private Access Token (PAT) Secret. The secret provides the credentials necessary to authenticate against the SCM. 
	* `repository_url` - (Required when scm_type=EXTERNAL_SCM) (Updatable) The repository URL for the SCM. For Non-Enterprise GitHub the expected format is https://github.com/[owner]/[repoName] For Enterprise GitHub the expected format is http(s)://[hostname]/api/v3/repos/[owner]/[repoName] For GitLab the expected format is https://gitlab.com/[groupName]/[repoName] 
	* `scm_type` - (Required) (Updatable) The type of Source Code Management.
	* `username` - (Applicable when scm_type=EXTERNAL_SCM) (Updatable) The username for the SCM (to perform operations such as cloning or pushing via HTTP).
* `verify_configuration` - (Required) (Updatable) The Verify stage configuration specifies a build service to run a pipeline for the recommended code changes. The build pipeline will be initiated to ensure that there is no breaking change after the dependency versions have been updated in source to avoid vulnerabilities. 
	* `additional_parameters` - (Applicable when build_service_type=GITHUB_ACTIONS | GITLAB_PIPELINE | JENKINS_PIPELINE | OCI_DEVOPS_BUILD) (Updatable) Additional key-value pairs passed as parameters to the build service when running an experiment.
	* `build_service_type` - (Required) (Updatable) The type of Build Service.
	* `jenkins_url` - (Required when build_service_type=JENKINS_PIPELINE) (Updatable) The URL that locates the Jenkins pipeline.
	* `job_name` - (Required when build_service_type=JENKINS_PIPELINE) (Updatable) The name of the Jenkins pipeline job that identifies the build pipeline.
	* `pat_secret_id` - (Required when build_service_type=GITHUB_ACTIONS | GITLAB_PIPELINE | JENKINS_PIPELINE) (Updatable) The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Private Access Token (PAT) Secret. The PAT provides the credentials to access the Jenkins Pipeline. 
	* `pipeline_id` - (Required when build_service_type=OCI_DEVOPS_BUILD) (Updatable) The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the user's DevOps Build Pipeline.
	* `repository_url` - (Required when build_service_type=GITHUB_ACTIONS | GITLAB_PIPELINE) (Updatable) The location of the repository where the GitHub Actions is defined. For Non-Enterprise GitHub the expected format is https://github.com/[owner]/[repoName] For Enterprise GitHub the expected format is http(s)://[hostname]/api/v3/repos/[owner]/[repoName] 
	* `trigger_secret_id` - (Required when build_service_type=GITLAB_PIPELINE) (Updatable) The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the trigger Secret. The Secret provides access to the trigger for a GitLab pipeline. 
	* `username` - (Required when build_service_type=GITHUB_ACTIONS | GITLAB_PIPELINE | JENKINS_PIPELINE) (Updatable) The username that will be used to authenticate with Jenkins.
	* `workflow_name` - (Required when build_service_type=GITHUB_ACTIONS) (Updatable) The name of the GitHub Actions workflow that defines the build pipeline.
* `state` - (Optional) (Updatable) The target state for the Remediation Recipe. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Remediation Recipe
	* `update` - (Defaults to 20 minutes), when updating the Remediation Recipe
	* `delete` - (Defaults to 20 minutes), when destroying the Remediation Recipe


## Import

RemediationRecipes can be imported using the `id`, e.g.

```
$ terraform import oci_adm_remediation_recipe.test_remediation_recipe "id"
```

