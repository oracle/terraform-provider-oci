resource "oci_generative_ai_model" "llama3_test_model" {
  #Required
  compartment_id                 = var.compartment_ocid
  base_model_id                  = local.llama_base_model_id
  fine_tune_details {
    dedicated_ai_cluster_id      = data.oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster_large_generic.id
    training_dataset {
      bucket                       = oci_objectstorage_bucket.fine_tune_bucket.name
      dataset_type                 = "OBJECT_STORAGE"
      namespace                    = data.oci_objectstorage_namespace.ns.namespace
      object                       = oci_objectstorage_object.fine_tune_object.object
    }
    training_config {
      training_config_type                = "LORA_TRAINING_CONFIG"
      total_training_epochs                = "3"
      learning_rate                        = "0.0002"
      training_batch_size                  = "8"
      early_stopping_patience              = "15"
      early_stopping_threshold             = "0.0001"
      log_model_metrics_interval_in_steps  = "10"
      lora_r                               = "8"
      lora_alpha                           = "8"
      lora_dropout                         = "0.1"
    }
  }

  #Optional
  display_name                  = var.llama3_test_model_display_name
  description                   = var.test_model_description
  vendor                        = var.test_model_vendor
  version                       = var.test_model_version
  #defined_tags not tested - cannot test in home region
  freeform_tags                 = var.test_freeform_tags
}

resource "oci_generative_ai_model" "test_model" {
  #Required
  compartment_id                 = var.compartment_ocid
  base_model_id                  = local.cohere_base_model_id
  fine_tune_details {
    dedicated_ai_cluster_id      = data.oci_generative_ai_dedicated_ai_cluster.test_fine_tuning_cluster.id
    training_dataset {
    bucket                       = oci_objectstorage_bucket.fine_tune_bucket.name
    dataset_type                 = "OBJECT_STORAGE"
    namespace                    = data.oci_objectstorage_namespace.ns.namespace
    object                       = oci_objectstorage_object.fine_tune_object.object
  }
    training_config {
      training_config_type         = "TFEW_TRAINING_CONFIG"
    }
  }

  #Optional
  display_name                  = var.test_model_display_name
  description                   = var.test_model_description
  vendor                        = var.test_model_vendor
  version                       = var.test_model_version
  #defined_tags not tested - cannot test in home region
  freeform_tags                 = var.test_freeform_tags
}

data "oci_generative_ai_model" "test_model" {
  #Required
  model_id                   = oci_generative_ai_model.test_model.id
}

data "oci_generative_ai_models" "test_models" {
  #Required
  compartment_id                = var.compartment_ocid
}

locals {

  filtered_base_models = [
	for item in data.oci_generative_ai_models.base_models.model_collection[0].items : item
	  if (
		(item.version == "14.2")
		&& contains(item.capabilities, "FINE_TUNE")
		&& (item.display_name == "cohere.command-light")
	  )
	]

  cohere_base_model_id = local.filtered_base_models[0].id

  llama_filtered_models = [
    for item in data.oci_generative_ai_models.llama_base_models.model_collection[0].items : item
    if (
    (item.version == "1.0.0")
    && contains(item.capabilities, "FINE_TUNE")
    && (item.display_name == "meta.llama-3-70b-instruct")
    )
  ]

  llama_base_model_id = local.llama_filtered_models[0].id
}

data "oci_generative_ai_models" "llama_base_models" {
  compartment_id = var.compartment_ocid
}

data "oci_generative_ai_models" "base_models" {
  compartment_id = var.compartment_ocid
  display_name = "cohere.command-light"
}

data "oci_objectstorage_namespace" "ns" {
  #Optional
  compartment_id = var.compartment_ocid
}

resource "oci_objectstorage_bucket" "fine_tune_bucket" {
    compartment_id               = var.compartment_ocid
    name                         = "fineTuneData"
    namespace                    = data.oci_objectstorage_namespace.ns.namespace
}

resource "oci_objectstorage_object" "fine_tune_object" {
    bucket                       = oci_objectstorage_bucket.fine_tune_bucket.name
    object                       = "uhc_data.jsonl"
    namespace                    = data.oci_objectstorage_namespace.ns.namespace
    content                      = <<EOF
{"prompt": "1", "completion": "one"}
{"prompt": "2", "completion": "two"}
{"prompt": "3", "completion": "three"}
{"prompt": "4", "completion": "four"}
{"prompt": "5", "completion": "five"}
{"prompt": "6", "completion": "six"}
{"prompt": "7", "completion": "seven"}
{"prompt": "8", "completion": "eight"}
{"prompt": "9", "completion": "nine"}
{"prompt": "10", "completion": "ten"}
{"prompt": "11", "completion": "eleven"}
{"prompt": "12", "completion": "twelve"}
{"prompt": "13", "completion": "thirteen"}
{"prompt": "14", "completion": "fourteen"}
{"prompt": "15", "completion": "fifteen"}
{"prompt": "16", "completion": "sixteen"}
{"prompt": "17", "completion": "seventeen"}
{"prompt": "18", "completion": "eighteen"}
{"prompt": "19", "completion": "nineteen"}
{"prompt": "20", "completion": "twenty"}
{"prompt": "21", "completion": "twenty-one"}
{"prompt": "22", "completion": "twenty-two"}
{"prompt": "23", "completion": "twenty-three"}
{"prompt": "24", "completion": "twenty-four"}
{"prompt": "25", "completion": "twenty-five"}
{"prompt": "26", "completion": "twenty-six"}
{"prompt": "27", "completion": "twenty-seven"}
{"prompt": "28", "completion": "twenty-eight"}
{"prompt": "29", "completion": "twenty-nine"}
{"prompt": "30", "completion": "thirty"}
{"prompt": "31", "completion": "thirty-one"}
{"prompt": "32", "completion": "thirty-two"}
EOF
}
