// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	instanceWithPreemptibleInstanceConfigRepresentation = map[string]interface{}{
		"availability_domain":         Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":              Representation{repType: Required, create: `${var.compartment_id}`},
		"shape":                       Representation{repType: Required, create: `VM.Standard2.1`},
		"image":                       Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
		"preemptible_instance_config": RepresentationGroup{Required, instancePreemptibleInstanceConfigRepresentation},
		"subnet_id":                   Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
	}

	InstanceWithPreemptibleInstanceConfigResourceConfig = InstanceResourceDependenciesWithoutDHV +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceWithPreemptibleInstanceConfigRepresentation)

	instanceWithPreemtibleInstanceConfigDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"state":               Representation{repType: Required, create: `RUNNING`},
		"filter":              RepresentationGroup{Required, instanceWithPreemtibleInstanceConfigDataSourceFilterRepresentation}}
	instanceWithPreemtibleInstanceConfigDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_instance.test_instance.id}`}},
	}

	instancePreemptibleInstanceConfigRepresentation = map[string]interface{}{
		"preemption_action": RepresentationGroup{Required, instancePreemptibleInstanceConfigPreemptionActionRepresentation},
	}
	instancePreemptibleInstanceConfigPreemptionActionRepresentation = map[string]interface{}{
		"type":                 Representation{repType: Required, create: `TERMINATE`},
		"preserve_boot_volume": Representation{repType: Required, create: `false`},
	}
)

func TestResourceCoreInstancePreemptibleInstanceConfig_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePreemptibleInstanceConfigResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + commonTestVariables()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: testProviderConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceWithPreemptibleInstanceConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "time_maintenance_reboot_due", ""),
					resource.TestCheckResourceAttr(resourceName, "preemptible_instance_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "preemptible_instance_config.0.preemption_action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "preemptible_instance_config.0.preemption_action.0.preserve_boot_volume", "false"),
					resource.TestCheckResourceAttr(resourceName, "preemptible_instance_config.0.preemption_action.0.type", "TERMINATE"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", Required, Create, instanceWithPreemtibleInstanceConfigDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceWithPreemptibleInstanceConfigRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),

					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.image"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.preemptible_instance_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.preemptible_instance_config.0.preemption_action.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.preemptible_instance_config.0.preemption_action.0.preserve_boot_volume", "false"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.preemptible_instance_config.0.preemption_action.0.type", "TERMINATE"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.ocpus", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.processor_description"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.source_details.0.source_id"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + InstanceWithPreemptibleInstanceConfigResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
