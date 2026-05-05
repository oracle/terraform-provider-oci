variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  # version          = "~> 8.6"
}

# Toggle to simulate delete (resource removed from config)
variable "create_enabled" {
  type    = bool
  default = true
}

# Toggle to simulate update (switch values)
variable "do_update" {
  type    = bool
  default = false
}

locals {
  data_config = var.do_update ? (
    <<-JSON
    [{"key3":"key4"}]
    JSON
    ) : (
    <<-JSON
    [{"key1":"key2"}]
    JSON
  )

  drilldown_config = var.do_update ? (
    <<-JSON
    [{"key3":"key4"}]
    JSON
    ) : (
    <<-JSON
    [{"key1":"key2"}]
    JSON
  )

  parameters_config = var.do_update ? (
    <<-JSON
    [{"key3":"key4"}]
    JSON
    ) : (
    <<-JSON
    [{"key1":"key2"}]
    JSON
  )

  display_name     = var.do_update ? "displayName2" : "displayName"
  description      = var.do_update ? "description2" : "description"
  provider_name    = var.do_update ? "providerName2" : "providerName"
  provider_version = var.do_update ? "providerVersion2" : "providerVersion"
  screen_image     = var.do_update ? "screenImage2" : "screenImage"
  search_type      = var.do_update ? "SEARCH_DONT_SHOW_IN_DASHBOARD" : "SEARCH_SHOW_IN_DASHBOARD"
  widget_template  = var.do_update ? "widgetTemplate2" : "widgetTemplate"
  widget_vm        = var.do_update ? "widgetVM2" : "widgetVM"

  nls = <<-JSON
  {"key1":"key2","key3":"key4"}
  JSON

  ui_config = <<-JSON
  {"key1":"key2","key3":"key4"}
  JSON
}

# CREATE/UPDATE resource (count=0 simulates DELETE by removing it from config)
resource "oci_management_dashboard_management_saved_search" "test_management_saved_search" {
  count = var.create_enabled ? 1 : 0

  compartment_id      = var.compartment_ocid
  data_config         = local.data_config
  description         = local.description
  display_name        = local.display_name
  is_oob_saved_search = var.do_update ? true : false
  metadata_version    = "2.0"
  nls                 = local.nls
  provider_id         = "management-dashboard"
  provider_name       = local.provider_name
  provider_version    = local.provider_version
  screen_image        = local.screen_image
  type                = local.search_type
  ui_config           = local.ui_config
  widget_template     = local.widget_template
  widget_vm           = local.widget_vm
  drilldown_config    = local.drilldown_config
  parameters_config   = local.parameters_config
}

# GET test (singular datasource) - only valid when resource exists
data "oci_management_dashboard_management_saved_search" "by_id" {
  count = var.create_enabled ? 1 : 0

  management_saved_search_id = oci_management_dashboard_management_saved_search.test_management_saved_search[0].id
}

# GET test (plural datasource) - list in compartment and optionally filter by display name
data "oci_management_dashboard_management_saved_searches" "list" {
  compartment_id = var.compartment_ocid

  # Optional: narrow results by display name. Remove if you want full list.
  display_name = local.display_name

  # If your provider supports "filter" blocks (as in your Go test), include it too.
  # This makes the list deterministic (it should return exactly the resource you created).
  filter {
    name   = "id"
    values = var.create_enabled ? [oci_management_dashboard_management_saved_search.test_management_saved_search[0].id] : []
  }
}

output "saved_search_id" {
  value = var.create_enabled ? oci_management_dashboard_management_saved_search.test_management_saved_search[0].id : null
}

output "saved_search_display_name_from_ds" {
  value = var.create_enabled ? data.oci_management_dashboard_management_saved_search.by_id[0].display_name : null
}

output "saved_searches_list_count" {
  # collection size depends on filtering; with the filter block above it should be 1 when created.
  value = length(try(data.oci_management_dashboard_management_saved_searches.list.management_saved_search_collection[0].items, []))
}

