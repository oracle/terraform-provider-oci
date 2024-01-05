// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreVolumeRequiredOnlyResource = CoreVolumeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Required, acctest.Create, CoreVolumeRepresentation)

	CoreVolumeResourceConfig = CoreVolumeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Update, CoreVolumeRepresentation)

	CoreCoreVolumeSingularDataSourceRepresentation = map[string]interface{}{
		"volume_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_volume.test_volume.id}`},
	}

	CoreCoreVolumeDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVolumeDataSourceFilterRepresentation}}
	CoreVolumeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_volume.test_volume.id}`}},
	}

	CoreVolumeRepresentation = map[string]interface{}{
		"availability_domain":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"backup_policy_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":           acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"size_in_gbs":          acctest.Representation{RepType: acctest.Optional, Create: `51`, Update: `52`},
		"source_details":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreVolumeSourceDetailsRepresentation},
		"vpus_per_gb":          acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `10`},
		"autotune_policies":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: volumeAutotunePoliciesRepresentation},
		"is_auto_tune_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
	}
	volumeAutotunePoliciesRepresentation = map[string]interface{}{
		"autotune_type":   acctest.Representation{RepType: acctest.Required, Create: `PERFORMANCE_BASED`, Update: `PERFORMANCE_BASED`},
		"max_vpus_per_gb": acctest.Representation{RepType: acctest.Optional, Create: `20`, Update: `30`},
	}
	CoreVolumeBlockVolumeReplicasRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `KvuH:US-ASHBURN-AD-1`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}
	CoreVolumeSourceDetailsRepresentation = map[string]interface{}{
		"id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_volume.source_volume.id}`},
		"type": acctest.Representation{RepType: acctest.Required, Create: `volume`},
	}

	CoreVolumeResourceDependencies = utils.VolumeBackupPolicyDependency +
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "source_volume", acctest.Required, acctest.Create, CoreVolumeRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig
)

// issue-routing-tag: core/blockStorage
func TestCoreVolumeResource_basic(t *testing.T) {
	if httpreplay.ShouldRetryImmediately() {
		t.Skip("TestCoreVolumeResource_basic is running flaky in http replay mode, will skip this for checkin test.")
	}

	httpreplay.SetScenario("TestCoreVolumeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_volume.test_volume"
	datasourceName := "data.oci_core_volumes.test_volumes"
	singularDatasourceName := "data.oci_core_volume.test_volume"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreVolumeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Create, CoreVolumeRepresentation), "core", "volume", t)

	acctest.ResourceTest(t, testAccCheckCoreVolumeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreVolumeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Required, acctest.Create, CoreVolumeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckNoResourceAttr(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				// Check on default values used
				resource.TestCheckResourceAttr(resourceName, "size_in_mbs", "51200"),
				resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "50"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreVolumeResourceDependencies,
		},

		{
			Config: config + compartmentIdVariableStr + CoreVolumeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Create, CoreVolumeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.autotune_type", "PERFORMANCE_BASED"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.max_vpus_per_gb", "20"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "51"),
				resource.TestCheckResourceAttr(resourceName, "size_in_mbs", "52224"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volume"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckNoResourceAttr(resourceName, "volume_backup_id"),
				resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),
				resource.TestCheckResourceAttr(resourceName, "vpus_per_gb", "10"),

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

		//verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreVolumeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreVolumeRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.autotune_type", "PERFORMANCE_BASED"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.max_vpus_per_gb", "20"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "51"),
				resource.TestCheckResourceAttr(resourceName, "size_in_mbs", "52224"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volume"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckNoResourceAttr(resourceName, "volume_backup_id"),
				resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),
				resource.TestCheckResourceAttr(resourceName, "vpus_per_gb", "10"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreVolumeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Update, CoreVolumeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.autotune_type", "PERFORMANCE_BASED"),
				resource.TestCheckResourceAttr(resourceName, "autotune_policies.0.max_vpus_per_gb", "30"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "52"),
				resource.TestCheckResourceAttr(resourceName, "size_in_mbs", "53248"),
				resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volume"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckNoResourceAttr(resourceName, "volume_backup_id"),
				resource.TestCheckNoResourceAttr(resourceName, "volume_group_id"),
				resource.TestCheckResourceAttr(resourceName, "vpus_per_gb", "10"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_volumes", "test_volumes", acctest.Optional, acctest.Update, CoreCoreVolumeDataSourceRepresentation) +
				compartmentIdVariableStr + CoreVolumeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Update, CoreVolumeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckNoResourceAttr(datasourceName, "backup_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "state"),
				resource.TestCheckNoResourceAttr(datasourceName, "volume_backup_id"),
				resource.TestCheckNoResourceAttr(datasourceName, "volume_group_id"),

				resource.TestCheckResourceAttr(datasourceName, "volumes.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "volumes.0.autotune_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "volumes.0.autotune_policies.0.autotune_type", "PERFORMANCE_BASED"),
				resource.TestCheckResourceAttr(datasourceName, "volumes.0.autotune_policies.0.max_vpus_per_gb", "30"),
				resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "volumes.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "volumes.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "volumes.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.is_hydrated"),
				resource.TestCheckResourceAttr(datasourceName, "volumes.0.size_in_gbs", "52"),
				resource.TestCheckResourceAttr(datasourceName, "volumes.0.size_in_mbs", "53248"),
				resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "volumes.0.vpus_per_gb", "10"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Required, acctest.Create, CoreCoreVolumeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreVolumeResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Update, CoreVolumeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "volume_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "autotune_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "autotune_policies.0.autotune_type", "PERFORMANCE_BASED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "autotune_policies.0.max_vpus_per_gb", "30"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_hydrated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "size_in_gbs", "52"),
				resource.TestCheckResourceAttr(singularDatasourceName, "size_in_mbs", "53248"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_details.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_details.0.type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vpus_per_gb", "10"),
			),
		},
		// verify resource import
		{
			Config:            config + CoreVolumeRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"volume_backup_id",
			},
			ResourceName: resourceName,
		},
	})
}

// This test is separated from the basic test due to weird behavior from Terraform test framework.
// An test step that results in an error will result in the state being voided. Isolate such test steps to
// avoid interfering with regular tests that Create/Update resources.
// issue-routing-tag: core/blockStorage
func TestCoreVolumeResource_expectError(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeResource_expectError")
	defer httpreplay.SaveScenario()
	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume.test_volume"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeDestroy,
		Steps: []resource.TestStep{
			// verify baseline Create
			{
				Config: config + `
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = 50 }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + CoreVolumeResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						_, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// ensure that giving non-numeric characters in size_in_gbs will yield an error
			{
				Config: config + `
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = "abc" }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + CoreVolumeResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("size_in_gbs", acctest.Representation{RepType: acctest.Required, Create: "abc"}, CoreVolumeRepresentation)),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("must be a 64-bit integer"),
			},
			// specify size in MBs and GBs, expect error
			{
				Config: config + `
resource "oci_core_volume" "test_volume" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"

	#Optional
    defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.volume_defined_tags_value}")}"
    display_name = "${var.volume_display_name}"
    freeform_tags = "${var.volume_freeform_tags}"
	size_in_gbs = "${var.volume_size_in_gbs}"
	size_in_mbs = "${var.volume_size_in_mbs}"
	source_details {
		#Required
		id = "${oci_core_volume.source_volume.id}"
		type = "${var.volume_source_details_type}"
	}
}
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = "50" }
variable "volume_size_in_mbs" { default = "51200" }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }
				` + compartmentIdVariableStr + CoreVolumeResourceDependencies,
				ExpectError: regexp.MustCompile("Megabytes and Gigabytes"),
			},
		},
	})
}

// This is a test to validate that interpolation syntax can be passed into int64
// fields that are being represented as strings in the schema. This is a regression
// test for issue found in https://github.com/oracle/terraform-provider-oci/issues/607
// issue-routing-tag: core/blockStorage
func TestCoreVolumeResource_int64_interpolation(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeResource_int64_interpolation")
	defer httpreplay.SaveScenario()
	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume.test_volume"
	resourceName2 := "oci_core_volume.test_volume2"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + CoreVolumeResourceConfig + `
data "oci_core_volumes" "test_volumes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"

	filter {
		name = "id"
		values = ["${oci_core_volume.test_volume.id}"]
	}
}

resource "oci_core_volume" "test_volume2" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"

	size_in_gbs = "${data.oci_core_volumes.test_volumes.volumes.0.size_in_gbs}"
}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					// Check on default values used
					resource.TestCheckResourceAttr(resourceName, "size_in_mbs", "53248"),
					resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "52"),
					resource.TestCheckResourceAttr(resourceName2, "size_in_mbs", "53248"),
					resource.TestCheckResourceAttr(resourceName2, "size_in_gbs", "52"),
				),
			},
		},
	})
}

// This test is separated from the basic test due to weird behavior from Terraform test framework.
// An test step that results in an error will result in the state being voided. Isolate such test steps to
// avoid interfering with regular tests that Create/Update resources.
// issue-routing-tag: core/blockStorage
func TestCoreVolumeResource_validations(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeResource_validations")
	defer httpreplay.SaveScenario()
	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume.test_volume"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeDestroy,
		Steps: []resource.TestStep{
			// verify baseline Create
			{
				Config: config + `
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = 50 }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + CoreVolumeResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						_, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// ensure that changing datatype of size_in_gbs is a no-op
			{
				Config: config + `
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = "50" }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + CoreVolumeResourceConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// ensure that adding leading zeroes to size_in_gbs is a no-op
			{
				Config: config + `
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = "0050" }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + CoreVolumeResourceConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// explicit volume size in MBs, noop
			{
				Config: config + `
resource "oci_core_volume" "test_volume" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"

	#Optional
 	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.volume_defined_tags_value}")}"
	display_name = "${var.volume_display_name}"
 	freeform_tags = "${var.volume_freeform_tags}"
	size_in_mbs = "${var.volume_size_in_mbs}"
	source_details {
		#Required
		id = "${oci_core_volume.source_volume.id}"
		type = "${var.volume_source_details_type}"
	}
}
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_mbs" { default = "51200" }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }
				` + compartmentIdVariableStr + CoreVolumeResourceDependencies,
				ExpectNonEmptyPlan: false,
			},
			// migrate size_in_mbs to size_in_gbs
			{
				Config: config + `
resource "oci_core_volume" "test_volume" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"

	#Optional
 	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.volume_defined_tags_value}")}"
	display_name = "${var.volume_display_name}"
	freeform_tags = "${var.volume_freeform_tags}"
	size_in_gbs = "${var.volume_size_in_gbs}"
	source_details {
		#Required
		id = "${oci_core_volume.source_volume.id}"
		type = "${var.volume_source_details_type}"
	}
}
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = "50" }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }
				` + compartmentIdVariableStr + CoreVolumeResourceDependencies,
				ExpectNonEmptyPlan: false,
			},
			// ensure that changing the case for source_details.?.type (polymorphic discriminator) is a no-op.
			{
				Config: config + `
resource "oci_core_volume" "test_volume" {
	#Required
	availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"

	#Optional
 	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.volume_defined_tags_value}")}"
	display_name = "${var.volume_display_name}"
	freeform_tags = "${var.volume_freeform_tags}"
	size_in_gbs = "${var.volume_size_in_gbs}"
	source_details {
		#Required
		id = "${oci_core_volume.source_volume.id}"
		type = "${var.volume_source_details_type}"
	}
}
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = "50" }
variable "volume_source_details_type" { default = "VoLume" } # case-insensitive
variable "volume_state" { default = "AVAILABLE" }
				` + compartmentIdVariableStr + CoreVolumeResourceDependencies,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func testAccCheckCoreVolumeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlockstorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume" {
			noResourceFound = false
			request := oci_core.GetVolumeRequest{}

			tmp := rs.Primary.ID
			request.VolumeId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetVolume(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VolumeLifecycleStateTerminated): true,
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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreVolume") {
		resource.AddTestSweepers("CoreVolume", &resource.Sweeper{
			Name:         "CoreVolume",
			Dependencies: acctest.DependencyGraph["volume"],
			F:            sweepCoreVolumeResource,
		})
	}
}

func sweepCoreVolumeResource(compartment string) error {
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()
	volumeIds, err := getCoreVolumeIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeId := range volumeIds {
		if ok := acctest.SweeperDefaultResourceId[volumeId]; !ok {
			deleteVolumeRequest := oci_core.DeleteVolumeRequest{}

			deleteVolumeRequest.VolumeId = &volumeId

			deleteVolumeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteVolume(context.Background(), deleteVolumeRequest)
			if error != nil {
				fmt.Printf("Error deleting Volume %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &volumeId, CoreVolumeSweepWaitCondition, time.Duration(3*time.Minute),
				CoreVolumeSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreVolumeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VolumeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()

	listVolumesRequest := oci_core.ListVolumesRequest{}
	listVolumesRequest.CompartmentId = &compartmentId
	listVolumesRequest.LifecycleState = oci_core.VolumeLifecycleStateAvailable
	listVolumesResponse, err := blockstorageClient.ListVolumes(context.Background(), listVolumesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Volume list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, volume := range listVolumesResponse.Items {
		id := *volume.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VolumeId", id)
	}
	return resourceIds, nil
}

func CoreVolumeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if volumeResponse, ok := response.Response.(oci_core.GetVolumeResponse); ok {
		return volumeResponse.LifecycleState != oci_core.VolumeLifecycleStateTerminated
	}
	return false
}

func CoreVolumeSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BlockstorageClient().GetVolume(context.Background(), oci_core.GetVolumeRequest{
		VolumeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
