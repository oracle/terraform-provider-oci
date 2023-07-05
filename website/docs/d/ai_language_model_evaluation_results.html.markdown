---
subcategory: "Ai Language"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_language_model_evaluation_results"
sidebar_current: "docs-oci-datasource-ai_language-model_evaluation_results"
description: |-
  Provides the list of Model Evaluation Results in Oracle Cloud Infrastructure Ai Language service
---

# Data Source: oci_ai_language_model_evaluation_results
This data source provides the list of Model Evaluation Results in Oracle Cloud Infrastructure Ai Language service.

Get a (paginated) list of evaluation results for a given model.

## Example Usage

```hcl
data "oci_ai_language_model_evaluation_results" "test_model_evaluation_results" {
	#Required
	model_id = oci_ai_language_model.test_model.id
}
```

## Argument Reference

The following arguments are supported:

* `model_id` - (Required) unique model OCID.


## Attributes Reference

The following attributes are exported:

* `evaluation_result_collection` - The list of evaluation_result_collection.

### ModelEvaluationResult Reference

The following attributes are exported:

* `items` - List of model evaluation analysis
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `location` - For CSV format location is rowId(1 is header) and for JSONL location is jsonL line sequence(1 is metadata)
	* `model_type` - Model type
	* `predicted_entities` - List of true(actual) entities in test data for NER model
		* `length` - Length of text
		* `offset` - Starting index on text.
		* `type` - Type of entity text like PER, LOC, GPE, NOPE etc.
	* `predicted_labels` - List of predicted labels by custom multi class or multi label TextClassification model
	* `record` - For CSV format location is rowId(1 is header) and for JSONL location is jsonL line sequence(1 is metadata)
	* `true_entities` - List of true(actual) entities in test data for NER model
		* `length` - Length of text
		* `offset` - Starting index on text.
		* `type` - Type of entity text like PER, LOC, GPE, NOPE etc.
	* `true_labels` - List of true(actual) labels in test data for multi class or multi label TextClassification

