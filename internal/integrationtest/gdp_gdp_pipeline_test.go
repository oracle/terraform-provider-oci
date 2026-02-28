// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_gdp "github.com/oracle/oci-go-sdk/v65/gdp"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	compartmentIdBucket            = utils.GetEnvSettingWithDefault("compartment_id_for_bucket", compartmentId)
	compartmentIdBucketVariableStr = fmt.Sprintf("variable \"compartment_id_for_bucket\" { default = \"%s\" }\n", compartmentIdBucket)

	compartmentIdPipeline            = utils.GetEnvSettingWithDefault("compartment_id_for_create", compartmentId)
	compartmentIdPipelineVariableStr = fmt.Sprintf("variable \"compartment_id_for_create\" { default = \"%s\" }\n", compartmentIdPipeline)

	GdpGdpPipelineRequiredOnlyResource = GdpGdpPipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_gdp_gdp_pipeline", "test_gdp_pipeline", acctest.Required, acctest.Create, GdpSenderPipelineRepresentation)

	GdpGdpPipelineResourceConfig = GdpGdpPipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_gdp_gdp_pipeline", "test_gdp_pipeline", acctest.Optional, acctest.Update, GdpSenderPipelineRepresentation)

	GdpGdpPipelineSingularDataSourceRepresentation = map[string]interface{}{
		"gdp_pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_gdp_gdp_pipeline.test_gdp_pipeline.id}`},
	}

	GdpGdpPipelineDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id_for_create}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: GdpGdpPipelineDataSourceFilterRepresentation}}
	GdpGdpPipelineDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Optional, Create: `state`},
		"values": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_gdp_gdp_pipeline.test_gdp_pipeline.state}`}},
	}

	lifecycleRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{"defined_tags"}},
	}

	GdpSenderPipelineRepresentation = map[string]interface{}{
		"bucket_details":                          []acctest.RepresentationGroup{{RepType: acctest.Required, Group: GdpPipelineSourceBucketDetailsRepresentation}, {RepType: acctest.Required, Group: GdpPipelineTransferBucketDetailsRepresentation}, {RepType: acctest.Required, Group: GdpPipelineRejectBucketDetailsRepresentation}},
		"compartment_id":                          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_create}`},
		"display_name":                            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"peering_region":                          acctest.Representation{RepType: acctest.Required, Create: `us-dcc-phoenix-1`},
		"pipeline_type":                           acctest.Representation{RepType: acctest.Required, Create: `SENDER`},
		"approval_key_vault_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_kms_vault.test_vault.id}`},
		"authorization_details":                   acctest.Representation{RepType: acctest.Optional, Create: `authorizationDetails`, Update: `authorizationDetails2`},
		"defined_tags":                            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"file_types":                              acctest.Representation{RepType: acctest.Optional, Create: []string{`fileTypes`}, Update: []string{`fileTypes2`}},
		"freeform_tags":                           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_approval_needed":                      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_chunking_enabled":                     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_file_override_in_destination_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_scanning_enabled":                     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"service_log_group_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_logging_log_group.test_log_group.id}`},
		"lifecycle":                               acctest.RepresentationGroup{RepType: acctest.Optional, Group: lifecycleRepresentation},
		"env":                                     acctest.Representation{RepType: acctest.Optional, Create: `COMMERCIAL`},
	}

	GdpPipelineSourceBucketRepresentation = acctest.RepresentationCopyWithNewProperties(ObjectStorageBucketRepresentation, map[string]interface{}{
		"name":                  acctest.Representation{RepType: acctest.Required, Create: testBucketName + "-source"},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_bucket}`},
		"object_events_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	})
	GdpPipelineSourceBucketDetailsRepresentation = map[string]interface{}{
		"bucket_type": acctest.Representation{RepType: acctest.Required, Create: `SOURCE`},
		"id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_source_bucket.bucket_id}`},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_source_bucket.name}`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_source_bucket.namespace}`},
	}

	GdpPipelineTransferBucketRepresentation = acctest.RepresentationCopyWithNewProperties(ObjectStorageBucketRepresentation, map[string]interface{}{
		"name":           acctest.Representation{RepType: acctest.Required, Create: testBucketName + "-transfer"},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_bucket}`},
	})
	GdpPipelineTransferBucketDetailsRepresentation = map[string]interface{}{
		"bucket_type": acctest.Representation{RepType: acctest.Required, Create: `TRANSFER`},
		"id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_transfer_bucket.bucket_id}`},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_transfer_bucket.name}`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_transfer_bucket.namespace}`},
	}

	GdpPipelineRejectBucketRepresentation = acctest.RepresentationCopyWithNewProperties(ObjectStorageBucketRepresentation, map[string]interface{}{
		"name":           acctest.Representation{RepType: acctest.Required, Create: testBucketName + "-reject"},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_bucket}`},
	})
	GdpPipelineRejectBucketDetailsRepresentation = map[string]interface{}{
		"bucket_type": acctest.Representation{RepType: acctest.Required, Create: `REJECT`},
		"id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_reject_bucket.bucket_id}`},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_reject_bucket.name}`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_reject_bucket.namespace}`},
	}

	GdpGdpPipelineResourceDependencies = DefinedTagsDependencies +
		compartmentIdBucketVariableStr + compartmentIdPipelineVariableStr +
		KmsKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, LoggingLogGroupRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_source_bucket", acctest.Required, acctest.Create, GdpPipelineSourceBucketRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_transfer_bucket", acctest.Required, acctest.Create, GdpPipelineTransferBucketRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_reject_bucket", acctest.Required, acctest.Create, GdpPipelineRejectBucketRepresentation)
)

// issue-routing-tag: gdp/default
func TestGdpGdpPipelineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGdpGdpPipelineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_gdp_gdp_pipeline.test_gdp_pipeline"
	datasourceName := "data.oci_gdp_gdp_pipelines.test_gdp_pipelines"
	singularDatasourceName := "data.oci_gdp_gdp_pipeline.test_gdp_pipeline"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GdpGdpPipelineResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_gdp_gdp_pipeline", "test_gdp_pipeline", acctest.Optional, acctest.Create, GdpSenderPipelineRepresentation), "gdp", "gdpPipeline", t)

	acctest.ResourceTest(t, testAccCheckGdpGdpPipelineDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GdpGdpPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_gdp_gdp_pipeline", "test_gdp_pipeline", acctest.Required, acctest.Create, GdpSenderPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket_details.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.0.bucket_type", "SOURCE"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.0.name", testBucketName+"-source"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.1.bucket_type", "TRANSFER"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.1.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.1.name", testBucketName+"-transfer"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.2.bucket_type", "REJECT"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.2.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.2.name", testBucketName+"-reject"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdPipeline),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "peering_region", "us-dcc-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "pipeline_type", "SENDER"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GdpGdpPipelineResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GdpGdpPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_gdp_gdp_pipeline", "test_gdp_pipeline", acctest.Optional, acctest.Create, GdpSenderPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "approval_key_vault_id"),
				resource.TestCheckResourceAttr(resourceName, "authorization_details", "authorizationDetails"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.0.bucket_type", "SOURCE"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.0.name", testBucketName+"-source"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.1.bucket_type", "TRANSFER"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.1.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.1.name", testBucketName+"-transfer"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.2.bucket_type", "REJECT"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.2.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.2.name", testBucketName+"-reject"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdPipeline),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "file_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_approval_needed", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_chunking_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_file_override_in_destination_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_scanning_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "peering_region", "us-dcc-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "pipeline_type", "SENDER"),
				resource.TestCheckResourceAttrSet(resourceName, "service_log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GdpGdpPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_gdp_gdp_pipeline", "test_gdp_pipeline", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GdpSenderPipelineRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "approval_key_vault_id"),
				resource.TestCheckResourceAttr(resourceName, "authorization_details", "authorizationDetails"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.0.bucket_type", "SOURCE"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.0.name", testBucketName+"-source"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.1.bucket_type", "TRANSFER"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.1.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.1.name", testBucketName+"-transfer"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.2.bucket_type", "REJECT"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.2.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.2.name", testBucketName+"-reject"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "file_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_approval_needed", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_chunking_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_file_override_in_destination_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_scanning_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "peering_region", "us-dcc-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "pipeline_type", "SENDER"),
				resource.TestCheckResourceAttrSet(resourceName, "service_log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + GdpGdpPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_gdp_gdp_pipeline", "test_gdp_pipeline", acctest.Optional, acctest.Update, GdpSenderPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "approval_key_vault_id"),
				resource.TestCheckResourceAttr(resourceName, "authorization_details", "authorizationDetails2"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.0.bucket_type", "SOURCE"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.0.name", testBucketName+"-source"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.1.bucket_type", "TRANSFER"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.1.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.1.name", testBucketName+"-transfer"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.2.bucket_type", "REJECT"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.2.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.2.name", testBucketName+"-reject"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdPipeline),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "file_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_approval_needed", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_chunking_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_file_override_in_destination_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_scanning_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "peering_region", "us-dcc-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "pipeline_type", "SENDER"),
				resource.TestCheckResourceAttrSet(resourceName, "service_log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_gdp_gdp_pipelines", "test_gdp_pipelines", acctest.Optional, acctest.Update, GdpGdpPipelineDataSourceRepresentation) +
				compartmentIdVariableStr + GdpGdpPipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_gdp_gdp_pipeline", "test_gdp_pipeline", acctest.Optional, acctest.Update, GdpSenderPipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentIdPipeline),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "gdp_pipeline_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "gdp_pipeline_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_gdp_gdp_pipeline", "test_gdp_pipeline", acctest.Required, acctest.Create, GdpGdpPipelineSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GdpGdpPipelineResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gdp_pipeline_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "authorization_details", "authorizationDetails2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket_details.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.0.bucket_type", "SOURCE"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.0.name", testBucketName+"-source"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.1.bucket_type", "TRANSFER"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.1.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.1.name", testBucketName+"-transfer"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.2.bucket_type", "REJECT"),
				resource.TestCheckResourceAttrSet(resourceName, "bucket_details.2.id"),
				resource.TestCheckResourceAttr(resourceName, "bucket_details.2.name", testBucketName+"-reject"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentIdPipeline),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_types.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_approval_needed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_chunking_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_file_override_in_destination_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_scanning_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "peering_region", "us-dcc-phoenix-1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pipeline_type", "SENDER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GdpGdpPipelineRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"env"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGdpGdpPipelineDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GuardedDataPipelineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_gdp_gdp_pipeline" {
			noResourceFound = false
			request := oci_gdp.GetGdpPipelineRequest{}

			tmp := rs.Primary.ID
			request.GdpPipelineId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "gdp")

			response, err := client.GetGdpPipeline(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_gdp.GdpPipelineLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GdpGdpPipeline") {
		resource.AddTestSweepers("GdpGdpPipeline", &resource.Sweeper{
			Name:         "GdpGdpPipeline",
			Dependencies: acctest.DependencyGraph["gdpPipeline"],
			F:            sweepGdpGdpPipelineResource,
		})
	}
}

func sweepGdpGdpPipelineResource(compartment string) error {
	guardedDataPipelineClient := acctest.GetTestClients(&schema.ResourceData{}).GuardedDataPipelineClient()
	gdpPipelineIds, err := getGdpGdpPipelineIds(compartment)
	if err != nil {
		return err
	}
	for _, gdpPipelineId := range gdpPipelineIds {
		if ok := acctest.SweeperDefaultResourceId[gdpPipelineId]; !ok {
			deleteGdpPipelineRequest := oci_gdp.DeleteGdpPipelineRequest{}

			deleteGdpPipelineRequest.GdpPipelineId = &gdpPipelineId

			deleteGdpPipelineRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "gdp")
			_, error := guardedDataPipelineClient.DeleteGdpPipeline(context.Background(), deleteGdpPipelineRequest)
			if error != nil {
				fmt.Printf("Error deleting GdpPipeline %s %s, It is possible that the resource is already deleted. Please verify manually \n", gdpPipelineId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &gdpPipelineId, GdpGdpPipelineSweepWaitCondition, time.Duration(3*time.Minute),
				GdpGdpPipelineSweepResponseFetchOperation, "gdp", true)
		}
	}
	return nil
}

func getGdpGdpPipelineIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "GdpPipelineId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	guardedDataPipelineClient := acctest.GetTestClients(&schema.ResourceData{}).GuardedDataPipelineClient()

	listGdpPipelinesRequest := oci_gdp.ListGdpPipelinesRequest{}
	listGdpPipelinesRequest.CompartmentId = &compartmentId
	listGdpPipelinesRequest.LifecycleState = oci_gdp.GdpPipelineLifecycleStateNeedsAttention
	listGdpPipelinesResponse, err := guardedDataPipelineClient.ListGdpPipelines(context.Background(), listGdpPipelinesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting GdpPipeline list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, gdpPipeline := range listGdpPipelinesResponse.Items {
		id := *gdpPipeline.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "GdpPipelineId", id)
	}
	return resourceIds, nil
}

func GdpGdpPipelineSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if gdpPipelineResponse, ok := response.Response.(oci_gdp.GetGdpPipelineResponse); ok {
		return gdpPipelineResponse.LifecycleState != oci_gdp.GdpPipelineLifecycleStateDeleted
	}
	return false
}

func GdpGdpPipelineSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GuardedDataPipelineClient().GetGdpPipeline(context.Background(), oci_gdp.GetGdpPipelineRequest{
		GdpPipelineId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
