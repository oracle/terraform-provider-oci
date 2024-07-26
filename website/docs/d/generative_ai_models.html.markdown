---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_models"
sidebar_current: "docs-oci-datasource-generative_ai-models"
description: |-
  Provides the list of Models in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_models
This data source provides the list of Models in Oracle Cloud Infrastructure Generative AI service.

Lists the models in a specific compartment. Includes pretrained base models and fine-tuned custom models.

## Example Usage

```hcl
data "oci_generative_ai_models" "test_models" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	capability = var.model_capability
	display_name = var.model_display_name
	id = var.model_id
	state = var.model_state
	vendor = var.model_vendor
}
```

## Argument Reference

The following arguments are supported:

* `capability` - (Optional) A filter to return only resources their capability matches the given capability.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The ID of the model.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.
* `vendor` - (Optional) A filter to return only resources that match the entire vendor given.


## Attributes Reference

The following attributes are exported:

* `model_collection` - The list of model_collection.

### Model Reference

The following attributes are exported:

* `base_model_id` - The OCID of the base model that's used for fine-tuning. For pretrained models, the value is null.
* `capabilities` - Describes what this model can be used for.
* `compartment_id` - The compartment OCID for fine-tuned models. For pretrained models, this value is null.
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
* `state` - The lifecycle state of the model.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the model was created in the format of an RFC3339 datetime string.
* `time_deprecated` - Corresponds to the time when the custom model and its associated foundation model will be deprecated.
* `time_updated` - The date and time that the model was updated in the format of an RFC3339 datetime string.
* `vendor` - The provider of the base model.
* `version` - The version of the model.

