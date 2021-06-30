---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deploy_environments"
sidebar_current: "docs-oci-datasource-devops-deploy_environments"
description: |-
  Provides the list of Deploy Environments in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_deploy_environments
This data source provides the list of Deploy Environments in Oracle Cloud Infrastructure Devops service.

Returns a list of deployment environments.

## Example Usage

```hcl
data "oci_devops_deploy_environments" "test_deploy_environments" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.deploy_environment_display_name
	id = var.deploy_environment_id
	project_id = oci_devops_project.test_project.id
	state = var.deploy_environment_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single resource by ID.
* `project_id` - (Optional) unique project identifier
* `state` - (Optional) A filter to return only DeployEnvironments that matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `deploy_environment_collection` - The list of deploy_environment_collection.

### DeployEnvironment Reference

The following attributes are exported:

* `cluster_id` - The OCID of the Kubernetes cluster.
* `compartment_id` - The OCID of a compartment.
* `compute_instance_group_selectors` - A collection of selectors. The combination of instances matching the selectors are included in the instance group.
	* `items` - A list of selectors for the instance group. UNION operator is used for combining the instances selected by each selector.
		* `compute_instance_ids` - Compute instance OCID identifiers that are members of this group.
		* `query` - Query expression confirming to the Oracle Cloud Infrastructure Search Language syntax to select compute instances for the group. The language is documented at https://docs.oracle.com/en-us/iaas/Content/Search/Concepts/querysyntax.htm
		* `region` - Region identifier referred by the deployment environment. Region identifiers are listed at https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
		* `selector_type` - Defines the type of the instance selector for the group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_environment_type` - Deployment environment type.
* `description` - Optional description about the deployment environment.
* `display_name` - Deployment environment display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `function_id` - The OCID of the Function.
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `project_id` - The OCID of a project.
* `state` - The current state of the deployment environment.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - Time the deployment environment was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - Time the deployment environment was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).

