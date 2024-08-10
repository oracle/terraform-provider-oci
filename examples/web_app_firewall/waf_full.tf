variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "waf_policy_display_name" {
  default = "tf_example_waf_policy_test"
}

variable "waf_web_app_firewall" {
  default = "tf_example_waf_web_app_firewall"
}

variable "waf_network_address_list_vcn" {
  default = "tf_example_waf_network_address_list_vcn"
}

variable "waf_network_address_list" {
  default = "tf_example_waf_network_address_list"
}

resource "oci_waf_web_app_firewall_policy" "test_waf_web_app_firewall_policy" {
  #Required
  compartment_id = var.compartment_ocid

  # Optional
  display_name = var.waf_policy_display_name

  actions {
    #Required
    name = "defaultAction"
    type = "ALLOW"
  }

  actions {
    #Required
    name = "return401Response"
    type = "RETURN_HTTP_RESPONSE"
    code = 401
    body {
      #Required
      type = "STATIC_TEXT"
      text = "{\n\"code\": 401,\n\"message\":\"Unauthorised\"\n}"
    }

    #Optional
    headers {
      #Required
      name = "Header1"
      value = "Value1"
    }

    headers {
      #Required
      name = "Header2"
      value = "Value2"
    }
  }

  actions {
    #Required
    name = "dynamicReturn401Response"
    type = "RETURN_HTTP_RESPONSE"
    code = 401
    body {
      #Required
      type = "DYNAMIC"
      # need $${ so that terraform doesn't try to replace ${http.request.id}
      template = "{\n\"code\": 401,\n\"message\":\"Unauthorised: requestId: $${http.request.id}}"
    }

    #Optional
    headers {
      #Required
      name = "Header1"
      value = "Value1"
    }

    headers {
      #Required
      name = "Header2"
      value = "Value2"
    }
  }

  request_access_control {
    #Required
    default_action_name = "defaultAction"
    #Optional
    rules {
      #Required
      type = "ACCESS_CONTROL"
      name = "requestAccessRule"
      action_name = "return401Response"
      condition = "i_contains(keys(http.request.headers), 'Header1')"
      #Optional
      condition_language = "JMESPATH"
    }
  }

  request_protection {
    #Optional
    body_inspection_size_limit_exceeded_action_name = "return401Response"
    body_inspection_size_limit_in_bytes = 8192
    rules {
      #Required
      type = "PROTECTION"
      name = "requestProtectionRule"
      action_name = "return401Response"
      is_body_inspection_enabled = true
      protection_capabilities {
        #Required
        key = "9300000"
        version = 1
        #Optional
        collaborative_action_threshold = 4
        collaborative_weights {
          key = "9301000"
          weight = 2
        }
        collaborative_weights {
          key = "9301100"
          weight = 2
        }
        collaborative_weights {
          key = "9301200"
          weight = 2
        }
        collaborative_weights {
          key = "9301300"
          weight = 2
        }
        exclusions {
          args = ["argName1", "argName2"]
          request_cookies = ["cookieName1", "cookieName2"]
        }
      }
    }
  }

  request_rate_limiting {
    #Optional
    rules {
      #Required
      type = "REQUEST_RATE_LIMITING"
      name = "requestRateLimitingRule"
      action_name = "return401Response"
      configurations {
        #Required
        period_in_seconds = 100
        requests_limit = 10
        #Optional
        action_duration_in_seconds = 10
      }
      #Optional
      condition = "i_contains(keys(http.request.headers), 'Header1')"
      condition_language = "JMESPATH"
    }
  }

  response_access_control {
    #Optional
    rules {
      #Required
      type = "ACCESS_CONTROL"
      name = "responseAccessRule"
      action_name = "return401Response"
      condition = "i_contains(keys(http.response.headers), 'Header1')"
      #Optional
      condition_language = "JMESPATH"
    }
  }

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_waf_network_address_list" "test_waf_network_address_list_vcn" {
  #Required
  compartment_id = var.compartment_ocid
  type = "VCN_ADDRESSES"

  vcn_addresses {
    addresses = "10.1.1.0/24"
    vcn_id = oci_core_vcn.vcn.id
  }

  vcn_addresses {
    addresses = "10.1.2.0/24"
    vcn_id = oci_core_vcn.vcn.id
  }

  #Optional
  display_name = var.waf_network_address_list_vcn

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_waf_network_address_list" "test_waf_network_address_list" {
  #Required
  compartment_id = var.compartment_ocid
  type = "ADDRESSES"
  addresses = ["10.1.1.0/24", "10.2.1.3"]

  #Optional
  display_name = var.waf_network_address_list

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_waf_web_app_firewall" "test_waf_web_app_firewall" {
  #Required
  compartment_id = var.compartment_ocid
  backend_type = "LOAD_BALANCER"
  load_balancer_id = oci_load_balancer.lb.id
  web_app_firewall_policy_id = oci_waf_web_app_firewall_policy.test_waf_web_app_firewall_policy.id

  #Optional
  display_name = var.waf_web_app_firewall

  freeform_tags = {
    "Department" = "Finance"
  }
}

/* VCN for Network Address List and Load Balancers */
resource "oci_core_vcn" "vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "vcn1"
  dns_label      = "vcn1"
}

/* Data needed to create Load Balancer, which is needed for Web App Firewall */
resource "oci_core_subnet" "subnet1" {
  cidr_block                 = "10.1.20.0/24"
  display_name               = "subnet1"
  dns_label                  = "subnet1"
  compartment_id             = var.compartment_ocid
  vcn_id                     = oci_core_vcn.vcn.id
  security_list_ids          = [oci_core_vcn.vcn.default_security_list_id]
  route_table_id             = oci_core_vcn.vcn.default_route_table_id
  dhcp_options_id            = oci_core_vcn.vcn.default_dhcp_options_id
  prohibit_public_ip_on_vnic = true

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.vcn.id
}

resource "oci_load_balancer" "lb" {
  shape          = "100Mbps"
  compartment_id = var.compartment_ocid

  subnet_ids = [
    oci_core_subnet.subnet1.id,
  ]

  display_name               = "lb1"
  is_private                 = true
  network_security_group_ids = [oci_core_network_security_group.test_network_security_group.id]
}
