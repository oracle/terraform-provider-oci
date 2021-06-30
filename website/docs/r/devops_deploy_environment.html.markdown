---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deploy_environment"
sidebar_current: "docs-oci-resource-devops-deploy_environment"
description: |-
  Provides the Deploy Environment resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_deploy_environment
This resource provides the Deploy Environment resource in Oracle Cloud Infrastructure Devops service.

Creates a new deployment environment.

## Example Usage

```hcl
resource "oci_devops_deploy_environment" "test_deploy_environment" {
	#Required
	deploy_environment_type = var.deploy_environment_deploy_environment_type
	project_id = oci_devops_project.test_project.id

	#Optional
	cluster_id = oci_containerengine_cluster.test_cluster.id
	compute_instance_group_selectors {

		#Optional
		items {
			#Required
			selector_type = var.deploy_environment_compute_instance_group_selectors_items_selector_type

			#Optional
			compute_instance_ids = var.deploy_environment_compute_instance_group_selectors_items_compute_instance_ids
			query = var.deploy_environment_compute_instance_group_selectors_items_query
			region = var.deploy_environment_compute_instance_group_selectors_items_region
		}
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.deploy_environment_description
	display_name = var.deploy_environment_display_name
	freeform_tags = {"bar-key"= "value"}
	function_id = oci_functions_function.test_function.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required when deploy_environment_type=OKE_CLUSTER) (Updatable) The OCID of the Kubernetes cluster.
* `compute_instance_group_selectors` - (Required when deploy_environment_type=COMPUTE_INSTANCE_GROUP) (Updatable) A collection of selectors. The combination of instances matching the selectors are included in the instance group.
	* `items` - (Required when deploy_environment_type=COMPUTE_INSTANCE_GROUP) (Updatable) A list of selectors for the instance group. UNION operator is used for combining the instances selected by each selector.
		* `compute_instance_ids` - (Required when selector_type=INSTANCE_IDS) (Updatable) Compute instance OCID identifiers that are members of this group.
		* `query` - (Required when selector_type=INSTANCE_QUERY) (Updatable) Query expression confirming to the Oracle Cloud Infrastructure Search Language syntax to select compute instances for the group. The language is documented at https://docs.oracle.com/en-us/iaas/Content/Search/Concepts/querysyntax.htm
		* `region` - (Required when selector_type=INSTANCE_QUERY) (Updatable) Region identifier referred by the deployment environment. Region identifiers are listed at https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
		* `selector_type` - (Required) (Updatable) Defines the type of the instance selector for the group.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_environment_type` - (Required) (Updatable) Deployment environment type.
* `description` - (Optional) (Updatable) Optional description about the deployment environment.
* `display_name` - (Optional) (Updatable) Deployment environment display name. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `function_id` - (Required when deploy_environment_type=FUNCTION) (Updatable) The OCID of the Function.
* `project_id` - (Required) The OCID of a project.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Import

DeployEnvironments can be imported using the `id`, e.g.

```
$ terraform import oci_devops_deploy_environment.test_deploy_environment "id"
```

