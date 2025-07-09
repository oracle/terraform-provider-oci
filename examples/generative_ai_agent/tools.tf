resource "oci_generative_ai_agent_tool" "rag_test_tool" {
  agent_id       = var.test_agent_id
  compartment_id = var.compartment_ocid
  display_name   = var.test_agent_tool_display_name
  description    = var.test_agent_tool_description
  freeform_tags = {
    Department = "Finance"
  }
  tool_config {
    knowledge_base_configs {
      knowledge_base_id = var.test_knowledge_base_id
    }
    tool_config_type = "RAG_TOOL_CONFIG"
    generation_llm_customization {
      instruction = "instruction"
    }
  }
}


resource "oci_generative_ai_agent_tool" "sql_test_tool" {
  agent_id       = var.test_agent_id
  compartment_id = var.compartment_ocid
  display_name   = var.test_agent_tool_display_name
  description    = var.test_agent_tool_description
  tool_config {
    database_schema {
      content             = "CREATE TABLE example ();"
      input_location_type = "INLINE"
    }
    dialect                       = "SQL_LITE"
    model_size                    = "SMALL"
    should_enable_self_correction = false
    should_enable_sql_execution   = false
    tool_config_type              = "SQL_TOOL_CONFIG"
  }
}

resource "oci_generative_ai_agent_tool" "fc_test_tool" {
  agent_id       = var.test_agent_id
  compartment_id = var.compartment_ocid
  display_name   = var.test_agent_tool_display_name
  description    = var.test_agent_tool_description
  freeform_tags = {
    Department = "Finance"
  }
  tool_config {
    function {
      description = "description"
      name        = "name"
      parameters = {
        "parameters" = "parameters"
      }
    }
    tool_config_type = "FUNCTION_CALLING_TOOL_CONFIG"
  }
}
