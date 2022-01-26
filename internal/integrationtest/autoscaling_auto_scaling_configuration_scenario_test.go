// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/oracle/oci-go-sdk/v56/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/suite"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

var (
	AutoScalingConfigurationResourceConfigForScheduledExecution = AutoScalingConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, autoScalingConfigurationRepresentation2)

	AutoScalingConfigurationResourceConfigForScheduledExecutionResourceAction = AutoScalingConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, autoScalingConfigurationRepresentation3)

	autoScalingConfigurationRepresentation2 = map[string]interface{}{
		"auto_scaling_resources": acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationAutoScalingResourcesRepresentation},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"policies":               acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesRepresentationForScheduledExecution},
		"cool_down_in_seconds":   acctest.Representation{RepType: acctest.Optional, Create: `300`, Update: `400`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `example_autoscaling_configuration`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	autoScalingConfigurationRepresentation3 = map[string]interface{}{
		"auto_scaling_resources": acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationAutoScalingResourcesRepresentation},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"policies":               acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesRepresentationForScheduledExecution3},
		"cool_down_in_seconds":   acctest.Representation{RepType: acctest.Optional, Create: `300`, Update: `400`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `example_autoscaling_configuration`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	autoScalingConfigurationPoliciesRepresentationForScheduledExecution = map[string]interface{}{
		"capacity":           acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesCapacityRepresentation},
		"policy_type":        acctest.Representation{RepType: acctest.Required, Create: `scheduled`, Update: `scheduled`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `example_autoscaling_configuration`, Update: `displayName2`},
		"execution_schedule": acctest.RepresentationGroup{RepType: acctest.Optional, Group: autoScalingConfigurationPoliciesExecutionScheduleRepresentation},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	autoScalingConfigurationPoliciesRepresentationForScheduledExecution3 = map[string]interface{}{
		"policy_type":        acctest.Representation{RepType: acctest.Required, Create: `scheduled`, Update: `scheduled`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `example_autoscaling_configuration`, Update: `displayName2`},
		"execution_schedule": acctest.RepresentationGroup{RepType: acctest.Optional, Group: autoScalingConfigurationPoliciesExecutionScheduleRepresentation},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"resource_action":    acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesResourceActionRepresentation},
	}

	autoScalingConfigurationPoliciesResourceActionRepresentation = map[string]interface{}{
		"action":      acctest.Representation{RepType: acctest.Required, Create: `STOP`, Update: `START`},
		"action_type": acctest.Representation{RepType: acctest.Optional, Create: `power`},
	}

	autoScalingConfigurationPoliciesExecutionScheduleRepresentation = map[string]interface{}{
		"expression": acctest.Representation{RepType: acctest.Required, Create: `0 15 10 ? * *`},
		"timezone":   acctest.Representation{RepType: acctest.Required, Create: `UTC`},
		"type":       acctest.Representation{RepType: acctest.Required, Create: `cron`},
	}
)

type ResourceAutoScalingConfigurationTestSuite struct {
	suite.Suite
	Providers              map[string]terraform.ResourceProvider
	Config                 string
	OperatingSystem        string
	OperatingSystemVersion string
}

func (s *ResourceAutoScalingConfigurationTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	//testAccPreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + utils.OciImageIdsVariable + `
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

	_, tokenFn := acctest.TokenizeWithHttpReplay("instance_pool")

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config:             s.Config + tokenFn(TFInstancePool, values),
				ExpectNonEmptyPlan: true,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet("oci_core_instance_pool.TFInstancePool", "id"),
					resource.TestCheckResourceAttr("oci_core_instance_pool.TFInstancePool", "display_name", "TFInstancePool"),
					resource.TestCheckResourceAttr("oci_core_instance_pool.TFInstancePool", "size", "2"),
					resource.TestCheckResourceAttr("oci_core_instance_pool.TFInstancePool", "actual_size", "2"),

					resource.TestCheckResourceAttrSet("oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfiguration", "id"),
					resource.TestCheckResourceAttr("oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfiguration", "display_name", "TFAutoScalingConfiguration"),
					resource.TestCheckResourceAttr("oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfiguration", "is_enabled", "true"),
					resource.TestCheckResourceAttr("oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfiguration", "cool_down_in_seconds", "300"),

					func(s *terraform.State) (err error) {
						instancePoolId, err = acctest.FromInstanceState(s, "oci_core_instance_pool.TFInstancePool", "id")
						return err
					},
				),
			},
			{
				PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &instancePoolId, instancePoolRunningWaitCondition, time.Duration(10*time.Minute),
					instancePoolSweepResponseFetchOperation, "auto_scaling", true),
				Config: s.Config + tokenFn(TFInstancePool, values),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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

// issue-routing-tag: auto_scaling/default
func TestAutoScalingAutoScalingConfigurationResource_scheduledExecution(t *testing.T) {
	httpreplay.SetScenario("TestAutoScalingAutoScalingConfigurationResource_scheduledExecution")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration"
	datasourceName := "data.oci_autoscaling_auto_scaling_configurations.test_auto_scaling_configurations"
	singularDatasourceName := "data.oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckAutoScalingAutoScalingConfigurationDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create, autoScalingConfigurationRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.initial", "2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.max", "3"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.min", "2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "example_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.expression", "0 15 10 ? * *"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.timezone", "UTC"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.type", "cron"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "scheduled"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autoScalingConfigurationRepresentation2, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.initial", "2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.max", "3"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.min", "2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "example_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.expression", "0 15 10 ? * *"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.timezone", "UTC"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.type", "cron"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "scheduled"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, autoScalingConfigurationRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "400"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.initial", "4"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.max", "5"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.min", "3"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.expression", "0 15 10 ? * *"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.timezone", "UTC"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.type", "cron"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "scheduled"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("Resource updated when it was supposed to be recreated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_autoscaling_auto_scaling_configurations", "test_auto_scaling_configurations", acctest.Optional, acctest.Update, autoScalingConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, autoScalingConfigurationRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.cool_down_in_seconds", "400"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Required, acctest.Create, autoScalingConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutoScalingConfigurationResourceConfigForScheduledExecution,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cool_down_in_seconds", "400"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				// max_resource_count and min_resource_count are set as per the recent policy executed
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "max_resource_count"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "min_resource_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.0.initial", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.0.max", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.0.min", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.execution_schedule.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.execution_schedule.0.expression", "0 15 10 ? * *"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.execution_schedule.0.timezone", "UTC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.execution_schedule.0.type", "cron"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.policy_type", "scheduled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceConfigForScheduledExecution,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: auto_scaling/default
func TestAutoScalingAutoScalingConfigurationResource_scheduledExecution_ResourceAction(t *testing.T) {
	httpreplay.SetScenario("TestAutoScalingAutoScalingConfigurationResource_scheduledExecution_ResourceAction")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration"
	datasourceName := "data.oci_autoscaling_auto_scaling_configurations.test_auto_scaling_configurations"
	singularDatasourceName := "data.oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckAutoScalingAutoScalingConfigurationDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create, autoScalingConfigurationRepresentation3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "example_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.expression", "0 15 10 ? * *"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.timezone", "UTC"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.type", "cron"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.0.action", "STOP"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.0.action_type", "power"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "scheduled"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autoScalingConfigurationRepresentation3, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.0.action", "STOP"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.0.action_type", "power"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "example_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.expression", "0 15 10 ? * *"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.timezone", "UTC"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.type", "cron"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "scheduled"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, autoScalingConfigurationRepresentation3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "400"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.0.action", "START"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.0.action_type", "power"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.expression", "0 15 10 ? * *"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.timezone", "UTC"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.execution_schedule.0.type", "cron"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "scheduled"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("Resource updated when it was supposed to be recreated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_autoscaling_auto_scaling_configurations", "test_auto_scaling_configurations", acctest.Optional, acctest.Update, autoScalingConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, autoScalingConfigurationRepresentation3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.cool_down_in_seconds", "400"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Required, acctest.Create, autoScalingConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutoScalingConfigurationResourceConfigForScheduledExecutionResourceAction,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cool_down_in_seconds", "400"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				// max_resource_count and min_resource_count are set as per the recent policy executed
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "max_resource_count"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "min_resource_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.0.action", "START"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.resource_action.0.action_type", "power"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.execution_schedule.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.execution_schedule.0.expression", "0 15 10 ? * *"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.execution_schedule.0.timezone", "UTC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.execution_schedule.0.type", "cron"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.policy_type", "scheduled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceConfigForScheduledExecutionResourceAction,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func instancePoolRunningWaitCondition(response common.OCIOperationResponse) bool {
	if instancePoolResponse, ok := response.Response.(oci_core.GetInstancePoolResponse); ok {
		return instancePoolResponse.LifecycleState != oci_core.InstancePoolLifecycleStateRunning
	}
	return false
}

// issue-routing-tag: auto_scaling/default
func TestResourceAutoScalingConfigurationTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceAutoScalingConfigurationTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceAutoScalingConfigurationTestSuite))
}
