---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_deployment"
sidebar_current: "docs-oci-datasource-datascience-model_deployment"
description: |-
  Provides details about a specific Model Deployment in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_deployment
This data source provides details about a specific Model Deployment resource in Oracle Cloud Infrastructure Datascience service.

Retrieves the model deployment for the specified `modelDeploymentId`.

## Example Usage

```hcl
data "oci_datascience_model_deployment" "test_model_deployment" {
	#Required
	model_deployment_id = oci_datascience_model_deployment.test_model_deployment.id
}
```

## Argument Reference

The following arguments are supported:

* `model_deployment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment.


## Attributes Reference

The following attributes are exported:

* `category_log_details` - The log details for each category.
	* `access` - The log details.
		* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log group to work with.
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log to work with.
	* `predict` - The log details.
		* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log group to work with.
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log to work with.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment's compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model deployment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `description` - A short description of the model deployment.
* `display_name` - A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information. Example: `My ModelDeployment`
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment.
* `lifecycle_details` - Details about the state of the model deployment.
* `model_deployment_configuration_details` - The model deployment configuration details.
	* `deployment_type` - The type of the model deployment.
	* `environment_configuration_details` - The configuration to carry the environment details thats used in Model Deployment creation
		* `cmd` - The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
		* `entrypoint` - The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
		* `environment_configuration_type` - The environment configuration type
		* `environment_variables` - Environment variables to set for the web server container. The size of envVars must be less than 2048 bytes. Key should be under 32 characters. Key should contain only letters, digits and underscore (_) Key should start with a letter. Key should have at least 2 characters. Key should not end with underscore eg. `TEST_` Key if added cannot be empty. Value can be empty. No specific size limits on individual Values. But overall environment variables is limited to 2048 bytes. Key can't be reserved Model Deployment environment variables. 
		* `health_check_port` - The port on which the container [HEALTHCHECK](https://docs.docker.com/engine/reference/builder/#healthcheck) would listen. The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`. 
		* `image` - The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
		* `image_digest` - The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
		* `server_port` - The port on which the web server serving the inference is running. The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`. 
	* `model_configuration_details` - The model configuration details.
        * `bandwidth_mbps` - The minimum network bandwidth for the model deployment.
        * `instance_configuration` - The model deployment instance configuration
            * `instance_shape_name` - The shape used to launch the model deployment instances.
            * `model_deployment_instance_shape_config_details` - Details for the model-deployment instance shape configuration.
                * `cpu_baseline` - The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
                * `memory_in_gbs` - A model-deployment instance of type VM.Standard.E3.Flex or VM.Standard.E4.Flex allows the memory to be specified with in the range of 6 to 1024 GB. VM.Standard3.Flex memory range is between 6 to 512 GB and VM.Optimized3.Flex memory range is between 6 to 256 GB. 
                * `ocpus` - A model-deployment instance of type VM.Standard.E3.Flex or VM.Standard.E4.Flex allows the ocpu count to be specified with in the range of 1 to 64 ocpu. VM.Standard3.Flex OCPU range is between 1 to 32 ocpu and for VM.Optimized3.Flex OCPU range is 1 to 18 ocpu. 
            * `subnet_id` - A model deployment instance is provided with a VNIC for network access.  This specifies the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT/SGW gateway for egress.
        * `maximum_bandwidth_mbps` - The maximum network bandwidth for the model deployment.
        * `model_id` - The OCID of the model you want to deploy.
        * `scaling_policy` - The scaling policy to apply to each model of the deployment.
            * `auto_scaling_policies` - The list of autoscaling policy details. 
                * `auto_scaling_policy_type` - The type of autoscaling policy.
                * `initial_instance_count` - For a threshold-based autoscaling policy, this value is the initial number of instances to launch in the model deployment immediately after autoscaling is enabled. Note that anytime this value is updated, the number of instances will be reset to this value. After autoscaling retrieves performance metrics, the number of instances is automatically adjusted from this initial number to a number that is based on the limits that you set. 
                * `maximum_instance_count` - For a threshold-based autoscaling policy, this value is the maximum number of instances the model deployment is allowed to increase to (scale out). 
                * `minimum_instance_count` - For a threshold-based autoscaling policy, this value is the minimum number of instances the model deployment is allowed to decrease to (scale in). 
                * `rules` - The list of autoscaling policy rules. 
                    * `metric_expression_rule_type` - The metric expression for creating the alarm used to trigger autoscaling actions on the model deployment.

                        The following values are supported:
                        * `PREDEFINED_EXPRESSION`: An expression built using CPU or Memory metrics emitted by the Model Deployment Monitoring.
                        * `CUSTOM_EXPRESSION`: A custom Monitoring Query Language (MQL) expression. 
                    * `metric_type` - Metric type 
                    * `scale_in_configuration` - The scaling configuration for the predefined metric expression rule. 
                        * `instance_count_adjustment` - The value is used for adjusting the count of instances by. 
                        * `pending_duration` - The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five minutes before the alarm updates its state to "OK."

                            The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H` for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M. 
                        * `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval: `1m`-`60m` (also `1h`). You can optionally specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`.

                            Example of threshold alarm:

                            -----

                            CPUUtilization[1m]{resourceId = "MODEL_DEPLOYMENT_OCID"}.grouping().mean() < 25 CPUUtilization[1m]{resourceId = "MODEL_DEPLOYMENT_OCID"}.grouping().mean() > 75

                            ----- 
                        * `scaling_configuration_type` - The type of scaling configuration. 
                        * `threshold` - A metric value at which the scaling operation will be triggered. 
                    * `scale_out_configuration` - The scaling configuration for the predefined metric expression rule. 
                        * `instance_count_adjustment` - The value is used for adjusting the count of instances by. 
                        * `pending_duration` - The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five minutes before the alarm updates its state to "OK."

                            The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H` for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M. 
                        * `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval: `1m`-`60m` (also `1h`). You can optionally specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`.

                            Example of threshold alarm:

                            -----

                            CPUUtilization[1m]{resourceId = "MODEL_DEPLOYMENT_OCID"}.grouping().mean() < 25 CPUUtilization[1m]{resourceId = "MODEL_DEPLOYMENT_OCID"}.grouping().mean() > 75

                            ----- 
                        * `scaling_configuration_type` - The type of scaling configuration. 
                        * `threshold` - A metric value at which the scaling operation will be triggered. 
            * `cool_down_in_seconds` - For threshold-based autoscaling policies, this value is the minimum period of time to wait between scaling actions. The cooldown period gives the system time to stabilize before rescaling. The minimum value is 600 seconds, which is also the default. The cooldown period starts when the model deployment becomes ACTIVE after the scaling operation. 
            * `instance_count` - The number of instances for the model deployment.
            * `is_enabled` - Whether the autoscaling policy is enabled.
            * `policy_type` - The type of scaling policy.
* `model_deployment_system_data` - Model deployment system data.
	* `current_instance_count` - This value is the current count of the model deployment instances.
	* `system_infra_type` - The infrastructure type of the model deployment.
* `model_deployment_url` - The URL to interact with the model deployment.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model deployment.
* `state` - The state of the model deployment.
* `time_created` - The date and time the resource was created, in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 
