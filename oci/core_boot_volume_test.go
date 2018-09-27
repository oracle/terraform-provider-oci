// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	BootVolumeWaitConditionDuration = time.Duration(20 * time.Minute)
	BootVolumeRequiredOnlyResource  = BootVolumeResourceDependencies + `
resource "oci_core_boot_volume" "test_boot_volume" {
	#Required
	availability_domain = "${oci_core_instance.test_instance.availability_domain}"
	compartment_id = "${var.compartment_id}"
	source_details {
		#Required
		id = "${oci_core_instance.test_instance.boot_volume_id}"
		type = "${var.boot_volume_source_details_type}"
	}
}
`

	BootVolumeResourceConfig = BootVolumeResourceDependencies + `
resource "oci_core_boot_volume" "test_boot_volume" {
	#Required
	availability_domain = "${oci_core_instance.test_instance.availability_domain}"
	compartment_id = "${var.compartment_id}"
	source_details {
		#Required
		id = "${oci_core_instance.test_instance.boot_volume_id}"
		type = "${var.boot_volume_source_details_type}"
	}

	#Optional
	backup_policy_id = "${data.oci_core_volume_backup_policies.test_boot_volume_backup_policies.volume_backup_policies.0.id}"
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.boot_volume_defined_tags_value}")}"
	display_name = "${var.boot_volume_display_name}"
	freeform_tags = "${var.boot_volume_freeform_tags}"
	size_in_gbs = "${var.boot_volume_size_in_gbs}"
}
`
	BootVolumePropertyVariables = `
variable "boot_volume_availability_domain" { default = "availabilityDomain" }
variable "boot_volume_defined_tags_value" { default = "value" }
variable "boot_volume_display_name" { default = "displayName" }
variable "boot_volume_freeform_tags" { default = {"Department"= "Finance"} }
variable "boot_volume_size_in_gbs" { default = 50 }
variable "boot_volume_source_details_id" { default = "id" }
variable "boot_volume_source_details_type" { default = "bootVolume" }

`
	BootVolumeResourceDependencies = DefinedTagsDependencies + InstancePropertyVariables + InstanceResourceAsDependencyConfig + `
data "oci_core_volume_backup_policies" "test_boot_volume_backup_policies" {
	filter {
		name = "display_name"
		values = [ "silver" ]
	}
}
`
)

func TestCoreBootVolumeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_boot_volume.test_boot_volume"
	datasourceName := "data.oci_core_boot_volumes.test_boot_volumes"
	singularDatasourceName := "data.oci_core_boot_volume.test_boot_volume"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreBootVolumeDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + BootVolumePropertyVariables + compartmentIdVariableStr + BootVolumeRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckNoResourceAttr(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "bootVolume"),
					resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BootVolumeResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + BootVolumePropertyVariables + compartmentIdVariableStr + BootVolumeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "bootVolume"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				PreConfig: waitTillCondition(testAccProvider, &resId, bootVolumeWaitCondition, BootVolumeWaitConditionDuration,
					bootVolumeResponseFetchOperation, "core", false),
				Config: config + `
variable "boot_volume_availability_domain" { default = "availabilityDomain" }
variable "boot_volume_defined_tags_value" { default = "updatedValue" }
variable "boot_volume_display_name" { default = "displayName2" }
variable "boot_volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "boot_volume_size_in_gbs" { default = 51 }
variable "boot_volume_source_details_id" { default = "id" }
variable "boot_volume_source_details_type" { default = "bootVolume" }

                ` + compartmentIdVariableStr + BootVolumeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "51"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "source_details.0.id"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "bootVolume"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				PreConfig: waitTillCondition(testAccProvider, &resId, bootVolumeWaitCondition, BootVolumeWaitConditionDuration,
					bootVolumeResponseFetchOperation, "core", false),
				Config: config + `
variable "boot_volume_availability_domain" { default = "availabilityDomain" }
variable "boot_volume_defined_tags_value" { default = "updatedValue" }
variable "boot_volume_display_name" { default = "displayName2" }
variable "boot_volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "boot_volume_size_in_gbs" { default = 51 }
variable "boot_volume_source_details_id" { default = "id" }
variable "boot_volume_source_details_type" { default = "bootVolume" }

data "oci_core_boot_volumes" "test_boot_volumes" {
	#Required
	availability_domain = "${oci_core_instance.test_instance.availability_domain}"
	compartment_id = "${var.compartment_id}"

    filter {
    	name = "id"
    	values = ["${oci_core_boot_volume.test_boot_volume.id}"]
    }
}
                ` + compartmentIdVariableStr + BootVolumeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckNoResourceAttr(datasourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckNoResourceAttr(datasourceName, "volume_group_id"),

					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.size_in_gbs", "51"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.size_in_mbs"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.source_details.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.source_details.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "boot_volumes.0.source_details.0.type", "bootVolume"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volumes.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
variable "boot_volume_availability_domain" { default = "availabilityDomain" }
variable "boot_volume_defined_tags_value" { default = "updatedValue" }
variable "boot_volume_display_name" { default = "displayName2" }
variable "boot_volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "boot_volume_size_in_gbs" { default = 51 }
variable "boot_volume_source_details_id" { default = "id" }
variable "boot_volume_source_details_type" { default = "bootVolume" }

data "oci_core_boot_volume" "test_boot_volume" {
	#Required
	boot_volume_id = "${oci_core_boot_volume.test_boot_volume.id}"
}
                ` + compartmentIdVariableStr + BootVolumeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr(singularDatasourceName, "backup_policy_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_id"),
					resource.TestCheckNoResourceAttr(singularDatasourceName, "volume_group_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_hydrated", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "size_in_gbs", "51"),
					resource.TestCheckResourceAttr(singularDatasourceName, "size_in_mbs", "52224"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "source_details.0.id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source_details.0.type", "bootVolume"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}

func testAccCheckCoreBootVolumeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_boot_volume" {
			noResourceFound = false
			request := oci_core.GetBootVolumeRequest{}

			tmp := rs.Primary.ID
			request.BootVolumeId = &tmp

			response, err := client.GetBootVolume(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.BootVolumeLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func bootVolumeResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *oci_common.RetryPolicy) error {
	_, err := client.blockstorageClient.GetBootVolume(context.Background(), oci_core.GetBootVolumeRequest{
		BootVolumeId: resourceId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func bootVolumeWaitCondition(response oci_common.OCIOperationResponse) bool {
	// Only stop if the volume is hydrated
	if bootVolumeResponse, ok := response.Response.(oci_core.GetBootVolumeResponse); ok {
		return *bootVolumeResponse.IsHydrated == false
	}
	return false
}
