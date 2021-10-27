// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v50/common"
	oci_core "github.com/oracle/oci-go-sdk/v50/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ComputeImageCapabilitySchemaRequiredOnlyResource = ComputeImageCapabilitySchemaResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_compute_image_capability_schema", "test_compute_image_capability_schema", Required, Create, computeImageCapabilitySchemaRepresentation)

	ComputeImageCapabilitySchemaResourceConfig = ComputeImageCapabilitySchemaResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_compute_image_capability_schema", "test_compute_image_capability_schema", Optional, Update, computeImageCapabilitySchemaRepresentation)

	computeImageCapabilitySchemaSingularDataSourceRepresentation = map[string]interface{}{
		"compute_image_capability_schema_id": Representation{RepType: Required, Create: `${oci_core_compute_image_capability_schema.test_compute_image_capability_schema.id}`},
		"is_merge_enabled":                   Representation{RepType: Required, Create: `false`},
	}

	computeImageCapabilitySchemaDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"image_id":       Representation{RepType: Optional, Create: `${oci_core_compute_image_capability_schema.test_compute_image_capability_schema.image_id}`},
		"filter":         RepresentationGroup{Required, computeImageCapabilitySchemaDataSourceFilterRepresentation}}
	computeImageCapabilitySchemaDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_compute_image_capability_schema.test_compute_image_capability_schema.id}`}},
	}

	computeImageCapabilitySchemaRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compute_global_image_capability_schema_version_name": Representation{RepType: Required, Create: `02b83c1f-a4db-4867-80df-d50d50f3b759`},
		"image_id": Representation{RepType: Required, Create: `${oci_core_image.custom_image.id}`},
		"schema_data": Representation{RepType: Required, Create: map[string]string{
			"Network.AttachmentType": "{\\\"descriptorType\\\": \\\"enumstring\\\",\\\"source\\\": \\\"IMAGE\\\", \\\"defaultValue\\\": \\\"PARAVIRTUALIZED\\\", \\\"values\\\": [\\\"PARAVIRTUALIZED\\\"]}",
		}, Update: map[string]string{
			"Network.AttachmentType": "{\\\"descriptorType\\\": \\\"enumstring\\\", \\\"source\\\": \\\"IMAGE\\\", \\\"defaultValue\\\": \\\"PARAVIRTUALIZED\\\", \\\"values\\\": [\\\"PARAVIRTUALIZED\\\", \\\"E1000\\\"]}",
		}},
		"defined_tags":  Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	computeImageCapabilitySchemaSchemaDataRepresentation = map[string]interface{}{
		"descriptor_type": Representation{RepType: Required, Create: `descriptorType`, Update: `descriptorType2`},
		"source":          Representation{RepType: Required, Create: `source`, Update: `source2`},
		"default_value":   Representation{RepType: Optional, Create: `false`, Update: `true`},
		"values":          Representation{RepType: Optional, Create: []string{`values`}, Update: []string{`values2`}},
	}

	ComputeImageCapabilitySchemaResourceDependencies = AvailabilityDomainConfig + `data "oci_core_images" "image_capability_images" {
		compartment_id   = "${var.tenancy_ocid}"
		display_name = "Windows-Server-2019-Standard-Edition-VM-E3-2020.06.10-0"
	}

	resource "oci_core_instance" "t" {
					availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.test_subnet.id}"
					image = "${data.oci_core_images.image_capability_images.images.0.id}"
					shape = "VM.Standard.E3.Flex"
					shape_config {
						ocpus = "1"
					}	
				}

resource "oci_core_image" "custom_image" {
    compartment_id   = "${var.tenancy_ocid}"
    instance_id = "${oci_core_instance.t.id}"
}` +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/computeImaging
func TestCoreComputeImageCapabilitySchemaResource_basic(t *testing.T) {
	t.Skip("Skip test for Windows image")
	httpreplay.SetScenario("TestCoreComputeImageCapabilitySchemaResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_compute_image_capability_schema.test_compute_image_capability_schema"
	datasourceName := "data.oci_core_compute_image_capability_schemas.test_compute_image_capability_schemas"
	singularDatasourceName := "data.oci_core_compute_image_capability_schema.test_compute_image_capability_schema"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ComputeImageCapabilitySchemaResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_compute_image_capability_schema", "test_compute_image_capability_schema", Optional, Create, computeImageCapabilitySchemaRepresentation), "core", "computeImageCapabilitySchema", t)

	ResourceTest(t, testAccCheckCoreComputeImageCapabilitySchemaDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ComputeImageCapabilitySchemaResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_compute_image_capability_schema", "test_compute_image_capability_schema", Required, Create, computeImageCapabilitySchemaRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compute_global_image_capability_schema_version_name"),
				resource.TestCheckResourceAttrSet(resourceName, "image_id"),
				resource.TestCheckResourceAttr(resourceName, "schema_data.%", "1"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ComputeImageCapabilitySchemaResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ComputeImageCapabilitySchemaResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_compute_image_capability_schema", "test_compute_image_capability_schema", Optional, Create, computeImageCapabilitySchemaRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compute_global_image_capability_schema_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_global_image_capability_schema_version_name"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "image_id"),
				resource.TestCheckResourceAttr(resourceName, "schema_data.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ComputeImageCapabilitySchemaResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_compute_image_capability_schema", "test_compute_image_capability_schema", Optional, Create,
					RepresentationCopyWithNewProperties(computeImageCapabilitySchemaRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "compute_global_image_capability_schema_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_global_image_capability_schema_version_name"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "image_id"),
				resource.TestCheckResourceAttr(resourceName, "schema_data.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ComputeImageCapabilitySchemaResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_compute_image_capability_schema", "test_compute_image_capability_schema", Optional, Update, computeImageCapabilitySchemaRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compute_global_image_capability_schema_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_global_image_capability_schema_version_name"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "image_id"),
				resource.TestCheckResourceAttr(resourceName, "schema_data.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_core_compute_image_capability_schemas", "test_compute_image_capability_schemas", Optional, Update, computeImageCapabilitySchemaDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeImageCapabilitySchemaResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_compute_image_capability_schema", "test_compute_image_capability_schema", Optional, Update, computeImageCapabilitySchemaRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "image_id"),

				resource.TestCheckResourceAttr(datasourceName, "compute_image_capability_schemas.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compute_image_capability_schemas.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_image_capability_schemas.0.compute_global_image_capability_schema_version_name"),
				resource.TestCheckResourceAttr(datasourceName, "compute_image_capability_schemas.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compute_image_capability_schemas.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "compute_image_capability_schemas.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_image_capability_schemas.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_image_capability_schemas.0.image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_image_capability_schemas.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_compute_image_capability_schema", "test_compute_image_capability_schema", Required, Create, computeImageCapabilitySchemaSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeImageCapabilitySchemaResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_image_capability_schema_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_merge_enabled", "false"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_global_image_capability_schema_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_global_image_capability_schema_version_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schema_data.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ComputeImageCapabilitySchemaResourceConfig,
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

func testAccCheckCoreComputeImageCapabilitySchemaDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_compute_image_capability_schema" {
			noResourceFound = false
			request := oci_core.GetComputeImageCapabilitySchemaRequest{}

			tmp := rs.Primary.ID
			request.ComputeImageCapabilitySchemaId = &tmp

			if value, ok := rs.Primary.Attributes["is_merge_enabled"]; ok {
				boolVal, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				request.IsMergeEnabled = &boolVal
			}

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

			_, err := client.GetComputeImageCapabilitySchema(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !InSweeperExcludeList("CoreComputeImageCapabilitySchema") {
		resource.AddTestSweepers("CoreComputeImageCapabilitySchema", &resource.Sweeper{
			Name:         "CoreComputeImageCapabilitySchema",
			Dependencies: DependencyGraph["computeImageCapabilitySchema"],
			F:            sweepCoreComputeImageCapabilitySchemaResource,
		})
	}
}

func sweepCoreComputeImageCapabilitySchemaResource(compartment string) error {
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()
	computeImageCapabilitySchemaIds, err := getComputeImageCapabilitySchemaIds(compartment)
	if err != nil {
		return err
	}
	for _, computeImageCapabilitySchemaId := range computeImageCapabilitySchemaIds {
		if ok := SweeperDefaultResourceId[computeImageCapabilitySchemaId]; !ok {
			deleteComputeImageCapabilitySchemaRequest := oci_core.DeleteComputeImageCapabilitySchemaRequest{}

			deleteComputeImageCapabilitySchemaRequest.ComputeImageCapabilitySchemaId = &computeImageCapabilitySchemaId

			deleteComputeImageCapabilitySchemaRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteComputeImageCapabilitySchema(context.Background(), deleteComputeImageCapabilitySchemaRequest)
			if error != nil {
				fmt.Printf("Error deleting ComputeImageCapabilitySchema %s %s, It is possible that the resource is already deleted. Please verify manually \n", computeImageCapabilitySchemaId, error)
				continue
			}
		}
	}
	return nil
}

func getComputeImageCapabilitySchemaIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ComputeImageCapabilitySchemaId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient()

	listComputeImageCapabilitySchemasRequest := oci_core.ListComputeImageCapabilitySchemasRequest{}
	listComputeImageCapabilitySchemasRequest.CompartmentId = &compartmentId
	listComputeImageCapabilitySchemasResponse, err := computeClient.ListComputeImageCapabilitySchemas(context.Background(), listComputeImageCapabilitySchemasRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ComputeImageCapabilitySchema list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, computeImageCapabilitySchema := range listComputeImageCapabilitySchemasResponse.Items {
		id := *computeImageCapabilitySchema.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ComputeImageCapabilitySchemaId", id)
	}
	return resourceIds, nil
}
