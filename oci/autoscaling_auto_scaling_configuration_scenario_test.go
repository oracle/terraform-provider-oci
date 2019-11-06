// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/oracle/oci-go-sdk/common"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

type ResourceAutoScalingConfigurationTestSuite struct {
	suite.Suite
	Providers              map[string]terraform.ResourceProvider
	Config                 string
	OperatingSystem        string
	OperatingSystemVersion string
}

func (s *ResourceAutoScalingConfigurationTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + OciImageIdsVariable + `
		data "oci_identity_availability_domains" "ADs" {
			compartment_id = "${var.tenancy_ocid}"
		}

		resource "oci_core_virtual_network" "ExampleVCN" {
			cidr_block     = "10.1.0.0/16"
			compartment_id = "${var.compartment_id}"
			display_name   = "TFExampleVCN"
			dns_label      = "tfexamplevcn"
		}

		resource "oci_core_subnet" "ExampleSubnet" {
			availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
			cidr_block          = "10.1.20.0/24"
			display_name        = "TFExampleSubnet"
			dns_label           = "tfexamplesubnet"
			security_list_ids   = ["${oci_core_virtual_network.ExampleVCN.default_security_list_id}"]
			compartment_id      = "${var.compartment_id}"
			vcn_id              = "${oci_core_virtual_network.ExampleVCN.id}"
			route_table_id      = "${oci_core_route_table.ExampleRT.id}"
			dhcp_options_id     = "${oci_core_virtual_network.ExampleVCN.default_dhcp_options_id}"
		}

		resource "oci_core_internet_gateway" "ExampleIG" {
			compartment_id = "${var.compartment_id}"
			display_name   = "TFExampleIG"
			vcn_id         = "${oci_core_virtual_network.ExampleVCN.id}"
		}

		resource "oci_core_route_table" "ExampleRT" {
			compartment_id = "${var.compartment_id}"
			vcn_id         = "${oci_core_virtual_network.ExampleVCN.id}"
			display_name   = "TFExampleRouteTable"

			route_rules {
				destination       = "0.0.0.0/0"
				destination_type  = "CIDR_BLOCK"
				network_entity_id = "${oci_core_internet_gateway.ExampleIG.id}"
			}
		}

		resource "oci_core_instance_configuration" "TFInstanceConfiguration" {
			compartment_id = "${var.compartment_id}"
			display_name   = "TFExampleInstanceConfiguration"

			instance_details {
				instance_type = "compute"

				launch_details {
					compartment_id = "${var.compartment_id}"
					ipxe_script    = "ipxeScript"
					shape          = "VM.Standard2.1"
					display_name   = "TFExampleInstanceConfigurationLaunchDetails"
	
					create_vnic_details {
						assign_public_ip       = true
						display_name           = "TFExampleInstanceConfigurationVNIC"
						skip_source_dest_check = false
					}
	
					extended_metadata = {
						some_string   = "stringA"
						nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
					}
	
					source_details {
						source_type = "image"
						image_id    = "${var.InstanceImageOCID[var.region]}"
					}
				}
			}
		}

		resource "oci_autoscaling_auto_scaling_configuration" "TFAutoScalingConfiguration" {
			compartment_id       = "${var.compartment_id}"
			cool_down_in_seconds = "300"
			display_name         = "TFAutoScalingConfiguration"
			is_enabled           = "true"
	
			policies {
				capacity {
					initial = "4"
					max     = "4"
					min     = "2"
				}
	
				display_name = "TFPolicy"
				policy_type  = "threshold"
	
				rules {
					action {
						type  = "CHANGE_COUNT_BY"
						value = "1"
					}
	
					display_name = "TFScaleOutRule"
	
					metric {
						metric_type = "CPU_UTILIZATION"
	
						threshold {
							operator = "GT"
							value    = "1"
						}
					}
				}
	
				rules {
					action {
						type  = "CHANGE_COUNT_BY"
						value = "-1"
					}
	
					display_name = "TFScaleInRule"
	
					metric {
						metric_type = "CPU_UTILIZATION"
	
						threshold {
							operator = "LT"
							value    = "1"
						}
					}
				}
			}
	
			auto_scaling_resources {
				id   = "${oci_core_instance_pool.TFInstancePool.id}"
				type = "instancePool"
			}
		}
	`
	s.OperatingSystem = "Oracle Linux"
}

func (s *ResourceAutoScalingConfigurationTestSuite) TestAccResourceAutoScalingConfiguration_InstancePoolSources() {
	var instancePoolId string

	var values = make(map[string]string)
	values["instance_pool_size"] = "2"

	var updatedValues = make(map[string]string)
	updatedValues["instance_pool_size"] = "3"

	var TFInstancePool = `
		resource "oci_core_instance_pool" "TFInstancePool" {
			compartment_id            = "${var.compartment_id}"
			instance_configuration_id = "${oci_core_instance_configuration.TFInstanceConfiguration.id}"
			size                      = {{.instance_pool_size}}
			state                     = "RUNNING"
			display_name              = "TFInstancePool"

			placement_configurations {
				availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
				primary_subnet_id   = "${oci_core_subnet.ExampleSubnet.id}"
			}
		}
	`

	_, tokenFn := tokenizeWithHttpReplay("instance_pool")

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config:             s.Config + tokenFn(TFInstancePool, values),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("oci_core_instance_pool.TFInstancePool", "id"),
					resource.TestCheckResourceAttr("oci_core_instance_pool.TFInstancePool", "display_name", "TFInstancePool"),
					resource.TestCheckResourceAttr("oci_core_instance_pool.TFInstancePool", "size", "2"),
					resource.TestCheckResourceAttr("oci_core_instance_pool.TFInstancePool", "actual_size", "2"),

					resource.TestCheckResourceAttrSet("oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfiguration", "id"),
					resource.TestCheckResourceAttr("oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfiguration", "display_name", "TFAutoScalingConfiguration"),
					resource.TestCheckResourceAttr("oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfiguration", "is_enabled", "true"),
					resource.TestCheckResourceAttr("oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfiguration", "cool_down_in_seconds", "300"),

					func(s *terraform.State) (err error) {
						instancePoolId, err = fromInstanceState(s, "oci_core_instance_pool.TFInstancePool", "id")
						return err
					},
				),
			},
			{
				PreConfig: waitTillCondition(testAccProvider, &instancePoolId, instancePoolRunningWaitCondition, time.Duration(10*time.Minute),
					instancePoolSweepResponseFetchOperation, "auto_scaling", true),
				Config: s.Config + tokenFn(TFInstancePool, values),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("oci_core_instance_pool.TFInstancePool", "display_name", "TFInstancePool"),
					resource.TestCheckResourceAttr("oci_core_instance_pool.TFInstancePool", "size", "2"),
					resource.TestCheckResourceAttr("oci_core_instance_pool.TFInstancePool", "actual_size", "4"),
				),
			},
			{
				Config:             s.Config + tokenFn(TFInstancePool, values),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config:             s.Config + tokenFn(TFInstancePool, updatedValues),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func instancePoolRunningWaitCondition(response common.OCIOperationResponse) bool {
	if instancePoolResponse, ok := response.Response.(oci_core.GetInstancePoolResponse); ok {
		return instancePoolResponse.LifecycleState != oci_core.InstancePoolLifecycleStateRunning
	}
	return false
}

func TestResourceAutoScalingConfigurationTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceAutoScalingConfigurationTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceAutoScalingConfigurationTestSuite))
}
