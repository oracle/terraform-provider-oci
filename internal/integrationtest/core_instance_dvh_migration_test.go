// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreInstanceForMigrationRep = map[string]interface{}{
		"availability_domain":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"image":                acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"subnet_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"dedicated_vm_host_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
	}

	CoreInstanceTargetDVH = acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_target_dedicated_vm_host", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(CoreDedicatedVmHostRepresentation, map[string]interface{}{
			"display_name": acctest.Representation{RepType: acctest.Optional, Create: `TargetDVH`},
		}))
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceResource_DVHMigration(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceResource_DVHMigration")
	defer httpreplay.SaveScenario()

	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + acctest.CommonTestVariables()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_instance.test_dvh_migration_instance"

	managementEndpoint := utils.GetEnvSettingWithBlankDefault("management_endpoint")
	managementEndpointStr := fmt.Sprintf("variable \"management_endpoint\" { default = \"%s\" }\n", managementEndpoint)

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckCoreInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: acctest.ProviderTestConfig() + compartmentIdVariableStr + managementEndpointStr + CoreInstanceResourceDependencies + CoreInstanceTargetDVH +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_dvh_migration_instance", acctest.Optional, acctest.Create, CoreInstanceForMigrationRep),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "time_maintenance_reboot_due", ""),
				resource.TestCheckResourceAttr(resourceName, "launch_options.0.network_type", "VFIO"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update to the DVH Id, this should trigger a migration
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + managementEndpointStr + CoreInstanceResourceDependencies + CoreInstanceTargetDVH +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_dvh_migration_instance", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("dedicated_vm_host_id", acctest.Representation{RepType: acctest.Optional,
						Update: `${oci_core_dedicated_vm_host.test_target_dedicated_vm_host.id}`}, CoreInstanceForMigrationRep)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "dedicated_vm_host_id"),
				func(s *terraform.State) (err error) {
					targetDvhId, _ := acctest.FromInstanceState(s, "oci_core_dedicated_vm_host.test_target_dedicated_vm_host", "id")
					return resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_id", targetDvhId)(s)
				},
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Delete instance first before deleting all other dependencies
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + managementEndpointStr + CoreInstanceResourceDependencies + CoreInstanceTargetDVH,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					time.Sleep(10 * time.Minute) // Let instance terminate before deleting the DVH
					return err
				},
			),
		},
		// Finish cleanup
		{
			Config: config,
		},
	})
}
