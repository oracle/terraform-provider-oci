// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v46/common"
	oci_core "github.com/oracle/oci-go-sdk/v46/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VolumeRequiredOnlyResource = VolumeResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Required, Create, volumeRepresentation)

	VolumeResourceConfig = VolumeResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Update, volumeRepresentation)

	volumeSingularDataSourceRepresentation = map[string]interface{}{
		"volume_id": Representation{repType: Required, create: `${oci_core_volume.test_volume.id}`},
	}

	volumeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":               Representation{repType: Optional, create: `AVAILABLE`},
		"filter":              RepresentationGroup{Required, volumeDataSourceFilterRepresentation}}
	volumeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume.test_volume.id}`}},
	}

	volumeRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"backup_policy_id":    Representation{repType: Optional, create: `${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}`},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":          Representation{repType: Optional, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"size_in_gbs":         Representation{repType: Optional, create: `51`, update: `52`},
		"source_details":      RepresentationGroup{Optional, volumeSourceDetailsRepresentation},
		"vpus_per_gb":         Representation{repType: Optional, create: `10`, update: `20`},
	}
	volumeBlockVolumeReplicasRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `NyKp:US-ASHBURN-AD-1`},
		"display_name":        Representation{repType: Optional, create: `displayName`},
	}
	volumeSourceDetailsRepresentation = map[string]interface{}{
		"id":   Representation{repType: Required, create: `${oci_core_volume.source_volume.id}`},
		"type": Representation{repType: Required, create: `volume`},
	}

	VolumeResourceDependencies = VolumeBackupPolicyDependency +
		generateResourceFromRepresentationMap("oci_core_volume", "source_volume", Required, Create, volumeRepresentation) +
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

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_volume.test_volume"
	datasourceName := "data.oci_core_volumes.test_volumes"
	singularDatasourceName := "data.oci_core_volume.test_volume"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+VolumeResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Create, volumeRepresentation), "core", "volume", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Required, Create, volumeRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckNoResourceAttr(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					// Check on default values used
					resource.TestCheckResourceAttr(resourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(resourceName, "size_in_gbs", "50"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeResourceDependencies,
			},

			{
				Config: config + compartmentIdVariableStr + VolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Create, volumeRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			//verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Create,
						representationCopyWithNewProperties(volumeRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + VolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Update, volumeRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resource.TestCheckResourceAttr(resourceName, "vpus_per_gb", "20"),

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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_volumes", "test_volumes", Optional, Update, volumeDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Update, volumeRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckNoResourceAttr(datasourceName, "backup_policy_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "state"),
					resource.TestCheckNoResourceAttr(datasourceName, "volume_backup_id"),
					resource.TestCheckNoResourceAttr(datasourceName, "volume_group_id"),

					resource.TestCheckResourceAttr(datasourceName, "volumes.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.is_hydrated"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.size_in_gbs", "52"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.size_in_mbs", "53248"),
					resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volumes.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "volumes.0.vpus_per_gb", "20"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_volume", "test_volume", Required, Create, volumeSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeResourceDependencies + generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Update, volumeRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "volume_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
					resource.TestCheckResourceAttr(singularDatasourceName, "vpus_per_gb", "20"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + VolumeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Update, volumeRepresentation),
			},
			//verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"volume_backup_id",
				},
				ResourceName: resourceName,
			},
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
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume.test_volume"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeDestroy,
		Steps: []resource.TestStep{
			// verify baseline create
			{
				Config: config + `
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = 50 }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + VolumeResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
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

                ` + compartmentIdVariableStr + VolumeResourceDependencies + generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Optional, Create,
					getUpdatedRepresentationCopy("size_in_gbs", Representation{repType: Required, create: "abc"}, volumeRepresentation)),
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
				` + compartmentIdVariableStr + VolumeResourceDependencies,
				ExpectError: regexp.MustCompile("Megabytes and Gigabytes"),
			},
		},
	})
}

// This is a test to validate that interpolation syntax can be passed into int64
// fields that are being represented as strings in the schema. This is a regression
// test for issue found in https://github.com/terraform-providers/terraform-provider-oci/issues/607
// issue-routing-tag: core/blockStorage
func TestCoreVolumeResource_int64_interpolation(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeResource_int64_interpolation")
	defer httpreplay.SaveScenario()
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume.test_volume"
	resourceName2 := "oci_core_volume.test_volume2"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VolumeResourceConfig + `
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
				Check: ComposeAggregateTestCheckFuncWrapper(
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
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume.test_volume"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeDestroy,
		Steps: []resource.TestStep{
			// verify baseline create
			{
				Config: config + `
variable "volume_defined_tags_value" { default = "updatedValue" }
variable "volume_display_name" { default = "displayName2" }
variable "volume_freeform_tags" { default = {"Department"= "Accounting"} }
variable "volume_size_in_gbs" { default = 50 }
variable "volume_source_details_type" { default = "volume" }
variable "volume_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + VolumeResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
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

                ` + compartmentIdVariableStr + VolumeResourceConfig,
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

                ` + compartmentIdVariableStr + VolumeResourceConfig,
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
				` + compartmentIdVariableStr + VolumeResourceDependencies,
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
				` + compartmentIdVariableStr + VolumeResourceDependencies,
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
				` + compartmentIdVariableStr + VolumeResourceDependencies,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func testAccCheckCoreVolumeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume" {
			noResourceFound = false
			request := oci_core.GetVolumeRequest{}

			tmp := rs.Primary.ID
			request.VolumeId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreVolume") {
		resource.AddTestSweepers("CoreVolume", &resource.Sweeper{
			Name:         "CoreVolume",
			Dependencies: DependencyGraph["volume"],
			F:            sweepCoreVolumeResource,
		})
	}
}

func sweepCoreVolumeResource(compartment string) error {
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient()
	volumeIds, err := getVolumeIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeId := range volumeIds {
		if ok := SweeperDefaultResourceId[volumeId]; !ok {
			deleteVolumeRequest := oci_core.DeleteVolumeRequest{}

			deleteVolumeRequest.VolumeId = &volumeId

			deleteVolumeRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteVolume(context.Background(), deleteVolumeRequest)
			if error != nil {
				fmt.Printf("Error deleting Volume %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeId, error)
				continue
			}
			waitTillCondition(testAccProvider, &volumeId, volumeSweepWaitCondition, time.Duration(3*time.Minute),
				volumeSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVolumeIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VolumeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "VolumeId", id)
	}
	return resourceIds, nil
}

func volumeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if volumeResponse, ok := response.Response.(oci_core.GetVolumeResponse); ok {
		return volumeResponse.LifecycleState != oci_core.VolumeLifecycleStateTerminated
	}
	return false
}

func volumeSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.blockstorageClient().GetVolume(context.Background(), oci_core.GetVolumeRequest{
		VolumeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
