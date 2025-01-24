---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_containers"
sidebar_current: "docs-oci-datasource-datascience-containers"
description: |-
  Provides the list of Containers in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_containers
This data source provides the list of Containers in Oracle Cloud Infrastructure Data Science service.

List containers.

## Example Usage

```hcl
data "oci_datascience_containers" "test_containers" {

	#Optional
	container_name = oci_datascience_container.test_container.name
	display_name = var.container_display_name
	is_latest = var.container_is_latest
	state = var.container_state
	tag_query_param = var.container_tag_query_param
	target_workload = var.container_target_workload
	usage_query_param = var.container_usage_query_param
}
```

## Argument Reference

The following arguments are supported:

* `container_name` - (Optional) <b>Filter</b> results by the container name.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `is_latest` - (Optional) if true, this returns latest version of container.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 
* `tag_query_param` - (Optional) <b>Filter</b> results by the container version tag.
* `target_workload` - (Optional) <b>Filter</b> results by the target workload.
* `usage_query_param` - (Optional) <b>Filter</b> results by the usage.


## Attributes Reference

The following attributes are exported:

* `containers` - The list of containers.

### Container Reference

The following attributes are exported:

* `container_name` - The name of the container. This can be same for different tags
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the container.
* `display_name` - The display name of the container.
* `family_name` - The family name of the container.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_latest` - The latest tag of the container.
* `state` - Container Version LifecycleState.
* `tag` - Container Tag.
* `tag_configuration_list` - An array of defined metadata details for the model.
	* `key` - Key of the container tag Metadata
	* `value` - Value of the container tag Metadata
* `target_workloads` - The list of target workload. This Container can be used with given data science resources.
* `usages` - The list of usages of this container. This Container can be used for given use-cases.
* `workload_configuration_details_list` - workload configuration of the container.
	* `additional_configurations` - The additional configurations
	* `cmd` - The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
	* `health_check_port` - The port on which the container [HEALTHCHECK](https://docs.docker.com/engine/reference/builder/#healthcheck) would listen. The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`. 
	* `server_port` - The port on which the web server serving the inference is running. The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`. 
	* `use_case_configuration` - The use-case configuration details 
		* `additional_configurations` - The additional configurations
		* `use_case_type` - The job-run use-case.
	* `workload_type` - The workload use case.

