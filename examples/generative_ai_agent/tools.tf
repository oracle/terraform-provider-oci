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

resource "oci_generative_ai_agent_tool" "http_test_tool" {
  agent_id       = var.test_agent_id
  compartment_id = var.compartment_ocid
  display_name   = var.test_agent_tool_display_name
  description    = var.test_agent_tool_description
  freeform_tags = {
    Department = "Finance"
  }
  tool_config {
    tool_config_type = "HTTP_ENDPOINT_TOOL_CONFIG"
    subnet_id        = var.test_subnet_id
    api_schema {
      api_schema_input_location_type = "INLINE"
      content = <<EOF
  {
    "openapi": "3.0.0",
    "info": {
      "title": "Minimal API",
      "version": "1.0"
    },
    "servers": [
      { "url": "https://example.com/api" }
    ],
    "paths": {
      "/ping": {
        "get": {
          "summary": "Ping for health check",
          "responses": {
            "200": {
              "description": "OK"
            }
          }
        }
      }
    }
  }
  EOF
    }

    http_endpoint_auth_config {
      http_endpoint_auth_sources {
        http_endpoint_auth_scope = "AGENT"
        http_endpoint_auth_scope_config {
          http_endpoint_auth_scope_config_type = "HTTP_ENDPOINT_NO_AUTH_SCOPE_CONFIG"
        }
      }
    }
  }
}

resource "oci_generative_ai_agent_tool" "agent_test_tool" {
  agent_id       = var.test_agent_id
  compartment_id = var.compartment_ocid
  display_name   = var.test_agent_tool_display_name
  description    = var.test_agent_tool_description
  freeform_tags = {
    Department = "Finance"
  }
  tool_config {
    tool_config_type = "AGENT_TOOL_CONFIG"
    agent_endpoint_id = var.test_agent_endpoint_id
  }
}