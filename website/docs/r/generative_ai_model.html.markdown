---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_model"
sidebar_current: "docs-oci-resource-generative_ai-model"
description: |-
  Provides the Model resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_model
This resource provides the Model resource in Oracle Cloud Infrastructure Generative AI service.

Creates a custom model by fine-tuning a base model with your own dataset. You can create a new custom models or create a new version of existing custom model..

The header contains an opc-work-request-id, which is the id for the WorkRequest that tracks the model creation progress.


## Example Usage

```hcl
resource "oci_generative_ai_model" "test_model" {
	#Required
	base_model_id = oci_generative_ai_model.test_model.id
	compartment_id = var.compartment_id
	fine_tune_details {
		#Required
		dedicated_ai_cluster_id = oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster.id
		training_dataset {
			#Required
			bucket = var.model_fine_tune_details_training_dataset_bucket
			dataset_type = var.model_fine_tune_details_training_dataset_dataset_type
			namespace = var.model_fine_tune_details_training_dataset_namespace
			object = var.model_fine_tune_details_training_dataset_object
		}

		#Optional
		training_config {
			#Required
			training_config_type = var.model_fine_tune_details_training_config_training_config_type

			#Optional
			early_stopping_patience = var.model_fine_tune_details_training_config_early_stopping_patience
			early_stopping_threshold = var.model_fine_tune_details_training_config_early_stopping_threshold
			learning_rate = var.model_fine_tune_details_training_config_learning_rate
			log_model_metrics_interval_in_steps = var.model_fine_tune_details_training_config_log_model_metrics_interval_in_steps
			lora_alpha = var.model_fine_tune_details_training_config_lora_alpha
			lora_dropout = var.model_fine_tune_details_training_config_lora_dropout
			lora_r = var.model_fine_tune_details_training_config_lora_r
			num_of_last_layers = var.model_fine_tune_details_training_config_num_of_last_layers
			total_training_epochs = var.model_fine_tune_details_training_config_total_training_epochs
			training_batch_size = var.model_fine_tune_details_training_config_training_batch_size
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.model_description
	display_name = var.model_display_name
	freeform_tags = {"Department"= "Finance"}
	vendor = var.model_vendor
	version = var.model_version
}
```

## Argument Reference

The following arguments are supported:

* `base_model_id` - (Required) The OCID of the base model that's used for fine-tuning.
* `compartment_id` - (Required) (Updatable) The compartment OCID for fine-tuned models. For pretrained models, this value is null.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) An optional description of the model.
* `display_name` - (Optional) (Updatable) A user-friendly name.
* `fine_tune_details` - (Required) Details about fine-tuning a custom model. 
	* `dedicated_ai_cluster_id` - (Required) The OCID of the dedicated AI cluster this fine-tuning runs on.
	* `training_config` - (Optional) The fine-tuning method and hyperparameters used for fine-tuning a custom model.
		* `early_stopping_patience` - (Optional) Stop training if the loss metric does not improve beyond 'early_stopping_threshold' for this many times of evaluation. 
		* `early_stopping_threshold` - (Optional) How much the loss must improve to prevent early stopping.
		* `learning_rate` - (Optional) The initial learning rate to be used during training
		* `log_model_metrics_interval_in_steps` - (Optional) Determines how frequently to log model metrics. 

			Every step is logged for the first 20 steps and then follows this parameter for log frequency. Set to 0 to disable logging the model metrics. 
		* `lora_alpha` - (Applicable when training_config_type=LORA_TRAINING_CONFIG) This parameter represents the scaling factor for the weight matrices in LoRA.
		* `lora_dropout` - (Applicable when training_config_type=LORA_TRAINING_CONFIG) This parameter indicates the dropout probability for LoRA layers.
		* `lora_r` - (Applicable when training_config_type=LORA_TRAINING_CONFIG) This parameter represents the LoRA rank of the update matrices.
		* `num_of_last_layers` - (Applicable when training_config_type=VANILLA_TRAINING_CONFIG) The number of last layers to be fine-tuned.
		* `total_training_epochs` - (Optional) The maximum number of training epochs to run for.
		* `training_batch_size` - (Optional) The batch size used during training.
		* `training_config_type` - (Required) The fine-tuning method for training a custom model.
	* `training_dataset` - (Required) The dataset used to fine-tune the model. 

		Only one dataset is allowed per custom model, which is split 80-20 for training and validating. You must provide the dataset in a JSON Lines (JSONL) file. Each line in the JSONL file must have the format: `{"prompt": "<first prompt>", "completion": "<expected completion given first prompt>"}` 
		* `bucket` - (Required) The Object Storage bucket name.
		* `dataset_type` - (Required) The type of the data asset.
		* `namespace` - (Required) The Object Storage namespace.
		* `object` - (Required) The Object Storage object name.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `vendor` - (Optional) (Updatable) The provider of the model.
* `version` - (Optional) (Updatable) The version of the model.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `base_model_id` - The OCID of the base model that's used for fine-tuning. For pretrained models, the value is null.
* `capabilities` - Describes what this model can be used for.
* `compartment_id` - The compartment OCID for fine-tuned models. For pretrained models, this value is null.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - An optional description of the model.
* `display_name` - A user-friendly name.
* `fine_tune_details` - Details about fine-tuning a custom model. 
	* `dedicated_ai_cluster_id` - The OCID of the dedicated AI cluster this fine-tuning runs on.
	* `training_config` - The fine-tuning method and hyperparameters used for fine-tuning a custom model.
		* `early_stopping_patience` - Stop training if the loss metric does not improve beyond 'early_stopping_threshold' for this many times of evaluation. 
		* `early_stopping_threshold` - How much the loss must improve to prevent early stopping.
		* `learning_rate` - The initial learning rate to be used during training
		* `log_model_metrics_interval_in_steps` - Determines how frequently to log model metrics. 

			Every step is logged for the first 20 steps and then follows this parameter for log frequency. Set to 0 to disable logging the model metrics. 
		* `lora_alpha` - This parameter represents the scaling factor for the weight matrices in LoRA.
		* `lora_dropout` - This parameter indicates the dropout probability for LoRA layers.
		* `lora_r` - This parameter represents the LoRA rank of the update matrices.
		* `num_of_last_layers` - The number of last layers to be fine-tuned.
		* `total_training_epochs` - The maximum number of training epochs to run for.
		* `training_batch_size` - The batch size used during training.
		* `training_config_type` - The fine-tuning method for training a custom model.
	* `training_dataset` - The dataset used to fine-tune the model. 

		Only one dataset is allowed per custom model, which is split 80-20 for training and validating. You must provide the dataset in a JSON Lines (JSONL) file. Each line in the JSONL file must have the format: `{"prompt": "<first prompt>", "completion": "<expected completion given first prompt>"}` 
		* `bucket` - The Object Storage bucket name.
		* `dataset_type` - The type of the data asset.
		* `namespace` - The Object Storage namespace.
		* `object` - The Object Storage object name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - An ID that uniquely identifies a pretrained or fine-tuned model.
* `is_long_term_supported` - Whether a model is supported long-term. Only applicable to base models.
* `lifecycle_details` - A message describing the current state of the model in more detail that can provide actionable information.
* `model_metrics` - Model metrics during the creation of a new model.
	* `final_accuracy` - Fine-tuned model accuracy.
	* `final_loss` - Fine-tuned model loss.
	* `model_metrics_type` - The type of the model metrics. Each type of model can expect a different set of model metrics.
* `previous_state` - You can create a custom model by using your dataset to fine-tune an out-of-the-box text generation base model. Have your dataset ready before you create a custom model. See [Training Data Requirements](https://docs.cloud.oracle.com/iaas/Content/generative-ai/training-data-requirements.htm).

	To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives Oracle Cloud Infrastructure resource access to users. See [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and [Getting Access to Generative AI Resouces](https://docs.cloud.oracle.com/iaas/Content/generative-ai/iam-policies.htm). 
* `state` - The lifecycle state of the model.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the model was created in the format of an RFC3339 datetime string.
* `time_deprecated` - Corresponds to the time when the custom model and its associated foundation model will be deprecated.
* `time_updated` - The date and time that the model was updated in the format of an RFC3339 datetime string.
* `type` - The model type indicating whether this is a pretrained/base model or a custom/fine-tuned model.
* `vendor` - The provider of the base model.
* `version` - The version of the model.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model
	* `update` - (Defaults to 20 minutes), when updating the Model
	* `delete` - (Defaults to 20 minutes), when destroying the Model


## Import

Models can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_model.test_model "id"
```

