// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ExaccDatabaseAutonomousDatabaseSoftwareImageRequiredOnlyResource = ExaccDatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Required, acctest.Create, ExaccDatabaseAutonomousDatabaseSoftwareImageRepresentation)

	DatabaseAutonomousDatabaseSoftwareImageRequiredOnlyResource = DatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseSoftwareImageRepresentation)

	DatabaseAutonomousDatabaseSoftwareImageResourceConfig = DatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseSoftwareImageRepresentation)

	ExaccDatabaseAutonomousDatabaseSoftwareImageResourceConfig = ExaccDatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Update, ExaccDatabaseAutonomousDatabaseSoftwareImageRepresentation)

	DatabaseAutonomousDatabaseSoftwareImageSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_software_image_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database_software_image.test_autonomous_database_software_image.id}`},
	}

	ExaccDatabaseAutonomousDatabaseSoftwareImageSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_software_image_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database_software_image.test_autonomous_database_software_image.id}`},
	}

	DatabaseAutonomousDatabaseSoftwareImageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"image_shape_family": acctest.Representation{RepType: acctest.Required, Create: `EXADATA_SHAPE`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `image1`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseSoftwareImageDataSourceFilterRepresentation}}

	DatabaseAutonomousDatabaseSoftwareImageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_autonomous_database_software_image.test_autonomous_database_software_image.id}`}},
	}

	ExaccDatabaseAutonomousDatabaseSoftwareImageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"image_shape_family": acctest.Representation{RepType: acctest.Required, Create: `EXACC_SHAPE`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `image1`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ExaccDatabaseAutonomousDatabaseSoftwareImageDataSourceFilterRepresentation}}

	ExaccDatabaseAutonomousDatabaseSoftwareImageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_autonomous_database_software_image.test_autonomous_database_software_image.id}`}},
	}

	DatabaseAutonomousDatabaseSoftwareImageRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `image1`},
		"image_shape_family": acctest.Representation{RepType: acctest.Required, Create: `EXADATA_SHAPE`},
		"source_cdb_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ExaccDatabaseAutonomousDatabaseSoftwareImageRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `image1`},
		"image_shape_family": acctest.Representation{RepType: acctest.Required, Create: `EXACC_SHAPE`},
		"source_cdb_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseAutonomousDatabaseSoftwareImageResourceDependencies = ATPDAutonomousContainerDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousContainerDatabaseRepresentation, []string{"backup_config"}))

	//acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousContainerDatabaseResourceConfig, []string{"backup_config"}))

	ExaccDatabaseAutonomousDatabaseSoftwareImageResourceDependencies = ExaccACDResourceConfig
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseExaccAutonomousDatabaseSoftwareImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccAutonomousDatabaseSoftwareImageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_database_software_image.test_autonomous_database_software_image"
	datasourceName := "data.oci_database_autonomous_database_software_images.test_autonomous_database_software_images"
	singularDatasourceName := "data.oci_database_autonomous_database_software_image.test_autonomous_database_software_image"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ExaccDatabaseAutonomousDatabaseSoftwareImageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Create, ExaccDatabaseAutonomousDatabaseSoftwareImageRepresentation), "database", "autonomousDatabaseSoftwareImage", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseSoftwareImageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Required, acctest.Create, ExaccDatabaseAutonomousDatabaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cdb_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousDatabaseSoftwareImageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Create, ExaccDatabaseAutonomousDatabaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttrSet(resourceName, "release_update"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cdb_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ExaccDatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ExaccDatabaseAutonomousDatabaseSoftwareImageRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "database_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttrSet(resourceName, "release_update"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cdb_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Update, ExaccDatabaseAutonomousDatabaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttrSet(resourceName, "release_update"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cdb_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_software_images", "test_autonomous_database_software_images", acctest.Optional, acctest.Update, ExaccDatabaseAutonomousDatabaseSoftwareImageDataSourceRepresentation) +
				compartmentIdVariableStr + ExaccDatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Update, ExaccDatabaseAutonomousDatabaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(datasourceName, "image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_database_software_image_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_database_software_image_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Required, acctest.Create, ExaccDatabaseAutonomousDatabaseSoftwareImageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ExaccDatabaseAutonomousDatabaseSoftwareImageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_software_image_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "autonomous_dsi_one_off_patches.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_shape_family", "EXACC_SHAPE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_update"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + ExaccDatabaseAutonomousDatabaseSoftwareImageRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"source_cdb_id",
			},
			ResourceName: resourceName,
		},
	})
}

func TestDatabaseAutonomousDatabaseSoftwareImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseSoftwareImageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_database_software_image.test_autonomous_database_software_image"
	datasourceName := "data.oci_database_autonomous_database_software_images.test_autonomous_database_software_images"
	singularDatasourceName := "data.oci_database_autonomous_database_software_image.test_autonomous_database_software_image"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousDatabaseSoftwareImageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseSoftwareImageRepresentation), "database", "autonomousDatabaseSoftwareImage", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseSoftwareImageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cdb_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseSoftwareImageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttrSet(resourceName, "release_update"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cdb_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseSoftwareImageRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "database_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttrSet(resourceName, "release_update"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cdb_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttrSet(resourceName, "release_update"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cdb_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_software_images", "test_autonomous_database_software_images", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseSoftwareImageDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabaseSoftwareImageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseSoftwareImageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(datasourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_database_software_image_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_database_software_image_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_software_image", "test_autonomous_database_software_image", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseSoftwareImageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabaseSoftwareImageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_software_image_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "autonomous_dsi_one_off_patches.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "image1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_shape_family", "EXADATA_SHAPE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_update"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseAutonomousDatabaseSoftwareImageRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"source_cdb_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseAutonomousDatabaseSoftwareImageDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_database_software_image" {
			noResourceFound = false
			request := oci_database.GetAutonomousDatabaseSoftwareImageRequest{}

			tmp := rs.Primary.ID
			request.AutonomousDatabaseSoftwareImageId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetAutonomousDatabaseSoftwareImage(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AutonomousDatabaseSoftwareImageLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseAutonomousDatabaseSoftwareImage") {
		resource.AddTestSweepers("DatabaseAutonomousDatabaseSoftwareImage", &resource.Sweeper{
			Name:         "DatabaseAutonomousDatabaseSoftwareImage",
			Dependencies: acctest.DependencyGraph["autonomousDatabaseSoftwareImage"],
			F:            sweepDatabaseAutonomousDatabaseSoftwareImageResource,
		})
	}
}

func sweepDatabaseAutonomousDatabaseSoftwareImageResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	autonomousDatabaseSoftwareImageIds, err := getDatabaseAutonomousDatabaseSoftwareImageIds(compartment)
	if err != nil {
		return err
	}
	for _, autonomousDatabaseSoftwareImageId := range autonomousDatabaseSoftwareImageIds {
		if ok := acctest.SweeperDefaultResourceId[autonomousDatabaseSoftwareImageId]; !ok {
			deleteAutonomousDatabaseSoftwareImageRequest := oci_database.DeleteAutonomousDatabaseSoftwareImageRequest{}

			deleteAutonomousDatabaseSoftwareImageRequest.AutonomousDatabaseSoftwareImageId = &autonomousDatabaseSoftwareImageId

			deleteAutonomousDatabaseSoftwareImageRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteAutonomousDatabaseSoftwareImage(context.Background(), deleteAutonomousDatabaseSoftwareImageRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousDatabaseSoftwareImage %s %s, It is possible that the resource is already deleted. Please verify manually \n", autonomousDatabaseSoftwareImageId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &autonomousDatabaseSoftwareImageId, DatabaseAutonomousDatabaseSoftwareImageSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseAutonomousDatabaseSoftwareImageSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseAutonomousDatabaseSoftwareImageIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AutonomousDatabaseSoftwareImageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listAutonomousDatabaseSoftwareImagesRequest := oci_database.ListAutonomousDatabaseSoftwareImagesRequest{}
	listAutonomousDatabaseSoftwareImagesRequest.CompartmentId = &compartmentId

	listAutonomousDatabaseSoftwareImagesRequest.ImageShapeFamily = oci_database.AutonomousDatabaseSoftwareImageImageShapeFamilyExaccShape

	listAutonomousDatabaseSoftwareImagesRequest.LifecycleState = oci_database.AutonomousDatabaseSoftwareImageLifecycleStateAvailable
	listAutonomousDatabaseSoftwareImagesResponse, err := databaseClient.ListAutonomousDatabaseSoftwareImages(context.Background(), listAutonomousDatabaseSoftwareImagesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AutonomousDatabaseSoftwareImage list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, autonomousDatabaseSoftwareImage := range listAutonomousDatabaseSoftwareImagesResponse.Items {
		id := *autonomousDatabaseSoftwareImage.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AutonomousDatabaseSoftwareImageId", id)
	}

	return resourceIds, nil
}

func DatabaseAutonomousDatabaseSoftwareImageSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousDatabaseSoftwareImageResponse, ok := response.Response.(oci_database.GetAutonomousDatabaseSoftwareImageResponse); ok {
		return autonomousDatabaseSoftwareImageResponse.LifecycleState != oci_database.AutonomousDatabaseSoftwareImageLifecycleStateTerminated
	}
	return false
}

func DatabaseAutonomousDatabaseSoftwareImageSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetAutonomousDatabaseSoftwareImage(context.Background(), oci_database.GetAutonomousDatabaseSoftwareImageRequest{
		AutonomousDatabaseSoftwareImageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
