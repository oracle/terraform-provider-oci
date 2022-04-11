/* The below .tf script must be run only once per realm. Cloud Guard is a realm-level service and once a region is selected as the reporting region,
*  all other existing and newly subscribed regions in the realm are automatically considered as monitoring regions.
*  The reporting region must be same as tenant's home region.
*  Replace <your_boat_group> in the IAM policy below with your BOAT group.
*/

//Common variables required
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

//****** Capability Check to verify Cloudguard is available in the given realm ******
data "capability" "cloudguard_available" {
  name = "cloudguard_available"
}

locals {
  cloud_guard_available = data.capability.cloudguard_available.is_available
  //Only enable Cloud Guard in reporting region as Cloud Guard is a realm-level service
  cloud_guard_enable    = var.cloud_guard_configuration_reporting_region == var.region
}

//****** Add Cloud Guard IAM Policies ******

resource "oci_identity_policy" "cloud_guard_policy" {
  count          = data.capability.cloudguard_available.is_available ? 1 : 0
  compartment_id = var.tenancy_ocid
  description    = "This policy allows cloud guard service to detect security vulnerabilities in tenancy"
  name           = "CloudGuardPolicies"
  statements     = [
    "allow group <your_boat_group> to read threat-intel-family in tenancy",
    "allow service cloudguard to read keys in tenancy",
    "allow service cloudguard to read compartments in tenancy",
    "allow service cloudguard to read compute-management-family in tenancy",
    "allow service cloudguard to read instance-family in tenancy",
    "allow service cloudguard to read virtual-network-family in tenancy",
    "allow service cloudguard to read volume-family in tenancy",
    "allow service cloudguard to read tenancies in tenancy",
    "allow service cloudguard to read audit-events in tenancy",
    "allow service cloudguard to read vaults in tenancy",
    "allow service cloudguard to read object-family in tenancy",
    "allow service cloudguard to read load-balancers in tenancy",
    "allow service cloudguard to read groups in tenancy",
    "allow service cloudguard to read dynamic-groups in tenancy",
    "allow service cloudguard to read users in tenancy",
    "allow service cloudguard to read database-family in tenancy",
    "allow service cloudguard to read authentication-policies in tenancy",
    "allow service cloudguard to read policies in tenancy",
    "allow service cloudguard to use network-security-groups in tenancy",
    "allow service cloudguard to read data-safe-family in tenancy",
    "allow service cloudguard to read autonomous-database-family in tenancy",
    "allow service cloudguard to manage cloudevents-rules in tenancy where target.rule.type='managed'"
  ]
}

//****** Enable Cloud Guard ******

//The reporting region needs to be a valid reporting region where the tenancy is subscribed to.
//In most cases the home-region of the tenancy is its reporting region.
//In a single region tenancy, the home region, reporting region and the monitoring region are all same.
variable "cloud_guard_configuration_reporting_region" {
  default = "us-ashburn-1"
}

//The acceptable values for status are `ENABLED` and `DISABLED`.
//DISABLING the tenancy is equivalent to off-boarding resulting in deletion of all the Control Plane entities, also disallowing most of the CloudGuard Operations.
//Once ENABLED, the reporting region can't be switched unless it is DISABLED and then ENABLED again for another region.
//However, The reporting region needs to be a valid reporting region where the tenancy is subscribed to.
variable "cloud_guard_configuration_status" {
  default = "ENABLED"
}

// Refer to the cloudguard_advancedMode_example to see its usage
variable "cloud_guard_configuration_self_manage_resources" {
  default = false
}

//Cloud Guard enabling and disabling is a tenant-level operation so the compartment-id needs to be a tenant-ocid.
resource "oci_cloud_guard_cloud_guard_configuration" "enable_cloud_guard" {
  count            = local.cloud_guard_available && local.cloud_guard_enable ? 1 : 0
  compartment_id   = var.tenancy_ocid
  reporting_region = var.cloud_guard_configuration_reporting_region
  status           = var.cloud_guard_configuration_status
  depends_on       = [oci_identity_policy.cloud_guard_policy]
}

//****** List Cloud Guard Responder Recipes ******

variable "responder_recipe_state" {
  default = "ACTIVE"
}

data "oci_cloud_guard_responder_recipes" "compartment_responder_recipes" {
  compartment_id = var.tenancy_ocid
  state          = var.responder_recipe_state
  depends_on     = [oci_cloud_guard_cloud_guard_configuration.enable_cloud_guard]
}

//****** List Cloud Guard Detector Recipes ******

variable "detector_recipe_state" {
  default = "ACTIVE"
}

data "oci_cloud_guard_detector_recipes" "compartment_detector_recipes" {
  compartment_id = var.tenancy_ocid
  state          = var.detector_recipe_state
  depends_on     = [oci_cloud_guard_cloud_guard_configuration.enable_cloud_guard]
}

//****** Create a Cloud Guard Target with all the default recipes attached ******
/* PS: refer to cloudguard_advancedMode_example and cloudguard_basic_example on
 *     how to further tweaks input settings for creation of customer control plane resources
 *     eg: targets, recipes, managedLists etc.
*/

variable "target_display_name" {
  default = "root"
}
variable "target_state" {
  default = "ACTIVE"
}
variable "target_resource_type" {
  default = "COMPARTMENT"
}

resource "oci_cloud_guard_target" "root_target" {
  compartment_id       = var.compartment_ocid
  display_name         = var.target_display_name
  target_resource_id   = var.compartment_ocid
  target_resource_type = var.target_resource_type
  state                = var.target_state

  dynamic "target_detector_recipes" {
    for_each = length(data.oci_cloud_guard_detector_recipes.compartment_detector_recipes.detector_recipe_collection) > 0 ?
    data.oci_cloud_guard_detector_recipes.compartment_detector_recipes.detector_recipe_collection[0].items : []
    iterator = detector_recipe
    content {
      detector_recipe_id = detector_recipe.value["id"]
    }
  }

  dynamic "target_responder_recipes" {
    for_each = length(data.oci_cloud_guard_responder_recipes.compartment_responder_recipes.responder_recipe_collection) > 0 ?
    data.oci_cloud_guard_responder_recipes.compartment_responder_recipes.responder_recipe_collection[0].items : []
    iterator = responder_recipe
    content {
      responder_recipe_id = responder_recipe.value["id"]
    }
  }

  depends_on = [oci_cloud_guard_cloud_guard_configuration.enable_cloud_guard]
}