package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func TestCoreInstancePoolMaintenancePolicyWithoutLoadBalancers(test *testing.T) {
	for _, scenario := range []struct {
		name string
		ipv6 bool
	}{
		{name: "ipv4"},
		{name: "ipv6", ipv6: true},
	} {
		test.Run(scenario.name, func(test *testing.T) {
			httpreplay.SetScenario("TestCoreInstancePoolMaintenancePolicyWithoutLoadBalancers_" + scenario.name)
			defer httpreplay.SaveScenario()

			availabilityDomain := utils.GetEnvSettingWithBlankDefault("availability_domain")
			const resourceName = "oci_core_instance_pool.test_instance_pool"
			config := func(fullRefresh bool, maxHealthy int) string {
				return coreInstancePoolMaintenancePolicyLiveConfig(scenario.ipv6, fullRefresh, maxHealthy, availabilityDomain)
			}
			check := func(fullRefresh bool, maxHealthy int) resource.TestCheckFunc {
				checks := []resource.TestCheckFunc{
					resource.TestCheckResourceAttr(resourceName, "size", "2"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "0"),
					resource.TestCheckResourceAttr("data.oci_core_instance_pools.test_instance_pools", "instance_pools.#", "1"),
				}
				for _, location := range []struct{ resourceName, prefix string }{
					{resourceName, ""},
					{"data.oci_core_instance_pool.test_instance_pool", ""},
					{"data.oci_core_instance_pools.test_instance_pools", "instance_pools.0."},
				} {
					prefix := location.prefix + "instance_maintenance_policy.0.launch_before_terminating.0."
					checks = append(checks,
						resource.TestCheckResourceAttr(location.resourceName, prefix+"is_full_refresh", fmt.Sprint(fullRefresh)),
						resource.TestCheckResourceAttr(location.resourceName, prefix+"max_healthy_percent", fmt.Sprint(maxHealthy)))
				}
				if availabilityDomain != "" {
					checks = append(checks, resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.availability_domain", availabilityDomain))
				}
				if scenario.ipv6 {
					checks = append(checks,
						resource.TestCheckResourceAttr(resourceName, "placement_configurations.0.primary_vnic_subnets.0.is_assign_ipv6ip", "true"),
						resource.TestCheckResourceAttrSet(resourceName, "placement_configurations.0.primary_vnic_subnets.0.ipv6address_ipv6subnet_cidr_pair_details.0.ipv6subnet_cidr"))
				}
				return resource.ComposeTestCheckFunc(checks...)
			}
			acctest.ResourceTest(test, testAccCheckCoreInstancePoolDestroy, []resource.TestStep{
				{Config: config(false, 110), Check: check(false, 110)},
				{Config: config(false, 110), PlanOnly: true, ExpectNonEmptyPlan: false},
				{Config: config(true, 120), Check: check(true, 120)},
				{ResourceName: resourceName, ImportState: true, ImportStateVerify: true},
				{Config: config(false, 110), Check: check(false, 110)},
				{Config: config(false, 110), PlanOnly: true, ExpectNonEmptyPlan: false},
			})
		})
	}
}

func coreInstancePoolMaintenancePolicyLiveConfig(ipv6, fullRefresh bool, maxHealthy int, availabilityDomain string) string {
	placement := `primary_subnet_id = oci_core_subnet.test_subnet.id`
	if ipv6 {
		placement = `primary_vnic_subnets {
      subnet_id = oci_core_subnet.test_subnet.id
      is_assign_ipv6ip = true
      ipv6address_ipv6subnet_cidr_pair_details {
        ipv6subnet_cidr = oci_core_subnet.test_subnet.ipv6cidr_blocks[0]
      }
    }`
	}
	return acctest.ProviderTestConfig() + utils.OciImageIdsVariable + fmt.Sprintf(`
variable "compartment_id" {
  default = %q
}
variable "availability_domain" {
  default = %q
}
data "oci_identity_availability_domains" "maintenance_policy" {
  compartment_id = var.tenancy_ocid
}
locals {
  availability_domain = var.availability_domain != "" ? var.availability_domain : data.oci_identity_availability_domains.maintenance_policy.availability_domains[0].name
}
resource "oci_core_vcn" "test_vcn" {
  compartment_id = var.compartment_id
  cidr_block = "10.0.0.0/16"
  display_name = "maintenance-policy-no-lb-test"
  is_ipv6enabled = %t
}
resource "oci_core_subnet" "test_subnet" {
  compartment_id = var.compartment_id
  vcn_id = oci_core_vcn.test_vcn.id
  cidr_block = "10.0.0.0/24"
  availability_domain = local.availability_domain
  prohibit_public_ip_on_vnic = true
  ipv6cidr_blocks = %t ? [cidrsubnet(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 8, 0)] : null
}
resource "oci_core_instance_configuration" "test_instance_configuration" {
  compartment_id = var.compartment_id
  display_name = "maintenance-policy-no-lb-test"
  instance_details {
    instance_type = "compute"
    launch_details {
      compartment_id = var.compartment_id
      availability_domain = local.availability_domain
      shape = "VM.Standard2.1"
      create_vnic_details {
        subnet_id = oci_core_subnet.test_subnet.id
        assign_public_ip = false
      }
      instance_options {
        are_legacy_imds_endpoints_disabled = true
      }
      source_details {
        source_type = "image"
        image_id = var.InstanceImageOCID[var.region]
      }
    }
  }
}
resource "oci_core_instance_pool" "test_instance_pool" {
  compartment_id = var.compartment_id
  instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id
  display_name = "maintenance-policy-no-lb-test"
  size = 2
  placement_configurations {
    availability_domain = local.availability_domain
    %s
  }
  instance_maintenance_policy {
    launch_before_terminating {
      is_full_refresh = %t
      max_healthy_percent = %d
    }
  }
  timeouts {
    create = "30m"
    update = "30m"
    delete = "30m"
  }
}
data "oci_core_instance_pool" "test_instance_pool" {
  instance_pool_id = oci_core_instance_pool.test_instance_pool.id
}
data "oci_core_instance_pools" "test_instance_pools" {
  compartment_id = var.compartment_id
  filter {
    name = "id"
    values = [oci_core_instance_pool.test_instance_pool.id]
  }
  depends_on = [oci_core_instance_pool.test_instance_pool]
}
`, utils.GetEnvSettingWithBlankDefault("compartment_ocid"), availabilityDomain, ipv6, ipv6, placement, fullRefresh, maxHealthy)
}
