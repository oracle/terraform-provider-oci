// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dif "github.com/oracle/oci-go-sdk/v65/dif"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	testCompartmentId                  = utils.GetEnvSettingWithBlankDefault("compartment_id")
	compartmentIdVariableStr           = fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", testCompartmentId)
	testCompartmentIdForMove           = utils.GetEnvSettingWithDefault("compartment_id_for_move", testCompartmentId)
	compartmentIdMoveVariableStr       = fmt.Sprintf("variable \"compartment_id_for_move\" { default = \"%s\" }\n", testCompartmentIdForMove)
	testPasswordSecretId               = utils.GetEnvSettingWithBlankDefault("password_secret_id")
	testGGCSPasswordSecretId           = utils.GetEnvSettingWithBlankDefault("ggcs_password_secret_id")
	testGGAdminPasswordSecretId        = utils.GetEnvSettingWithBlankDefault("gg_admin_password_secret_id")
	dataflow_file_uri                  = utils.GetEnvSettingWithBlankDefault("dataflow_file_uri")
	dataflow_archive_uri               = utils.GetEnvSettingWithBlankDefault("dataflow_archive_uri")
	dif_ggcs_connection_name           = "ggcsConn"
	passwordSecretIdVariableStr        = fmt.Sprintf("variable \"password_secret_id\" { default = \"%s\" }\n", testPasswordSecretId)
	ggcsPasswordSecretIdVariableStr    = fmt.Sprintf("variable \"ggcs_password_secret_id\" { default = \"%s\" }\n", testGGCSPasswordSecretId)
	ggAdminPasswordSecretIdVariableStr = fmt.Sprintf("variable \"gg_admin_password_secret_id\" { default = \"%s\" }\n", testGGAdminPasswordSecretId)
	dataflowFileUriVariableStr         = fmt.Sprintf("variable \"dataflow_file_uri\" { default = \"%s\" }\n", dataflow_file_uri)
	dataflowArchiveUriVariableStr      = fmt.Sprintf("variable \"dataflow_archive_uri\" { default = \"%s\" }\n", dataflow_archive_uri)

	DefinedTagsDifDependencies = `
		variable defined_tag_namespace_name { default = "" }
		resource "oci_identity_tag_namespace" "tag-namespace-terraform" {
				#Required
				compartment_id = "${var.compartment_id}"
				description = "example tag namespace"
				name = "${var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"}"

				is_retired = false
		}

		resource "oci_identity_tag" "tag1" {
				#Required
				description = "example tag"
				name = "example-tag"
				tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace-terraform.id}"

				is_retired = false
		}
		`
	PrivateSubnetDifDependency = `
		resource "oci_core_subnet" "test_private_subnet" {
				cidr_block = "10.0.1.0/24"
				compartment_id = "${var.compartment_id}"
				lifecycle {
					ignore_changes = ["defined_tags"]
				}
				vcn_id = "${oci_core_vcn.test_vcn.id}"
	            dns_label = "testsubnet"
				prohibit_public_ip_on_vnic = true
		}
		`

	ignoredTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{
			RepType: acctest.Required,
			Create: []string{
				"defined_tags",
				"system_tags",
				"freeform_tags",
				"time_created",
				"time_updated",
				"service_details",
			}},
	}

	DifStackRequiredOnlyResource = DifStackResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Create, DifStackRepresentation)

	DifStackResourceConfig = DifStackResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Optional, acctest.Update, DifStackRepresentation)

	DifStackSingularDataSourceRepresentation = map[string]interface{}{
		"stack_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dif_stack.test_stack.id}`},
	}

	DifStackDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: testCompartmentId},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `testStack`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_dif_stack.test_stack.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: string(oci_dif.StackLifecycleStateActive)},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataSourceFilterRepresentation}}
	DifStackDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dif_stack.test_stack.id}`}},
	}

	DifStackRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `testStack`},
		"services":        acctest.Representation{RepType: acctest.Required, Create: []string{`OBJECTSTORAGE`, `ADB`, `DATAFLOW`, `GGCS`}, Update: []string{`OBJECTSTORAGE`, `ADB`, `DATAFLOW`, `GGCS`, `GENAI`}},
		"stack_templates": acctest.Representation{RepType: acctest.Required, Create: []string{`DATALAKE`, `DATATRANSFORMATION`, `DATAPIPELINE`}, Update: []string{`DATALAKE`, `DATATRANSFORMATION`, `DATAPIPELINE`, `AISERVICES`}},
		"adb":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackAdbRepresentation},
		"dataflow":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackDataflowRepresentation},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("example-tag-namespace-all.example-tag", "value")}`, Update: `${map("example-tag-namespace-all.example-tag", "updatedValue")}`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		//"genai":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackGenaiRepresentation},
		"ggcs":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackGgcsRepresentation},
		"notification_email": acctest.Representation{RepType: acctest.Optional, Create: `test@example.com`},
		"objectstorage":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackObjectstorageRepresentation},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoredTagsRepresentation},
	}
	DifStackAdbRepresentation = map[string]interface{}{
		"admin_password_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.password_secret_id}`},
		"data_storage_size_in_tbs":    acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_version":                  acctest.Representation{RepType: acctest.Required, Create: `19c`},
		"db_workload":                 acctest.Representation{RepType: acctest.Required, Create: `DW`},
		"ecpu":                        acctest.Representation{RepType: acctest.Required, Create: `2`},
		"instance_id":                 acctest.Representation{RepType: acctest.Required, Create: `testAdb`},
		"is_mtls_connection_required": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_public":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subnet_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_private_subnet.id}`},
		"tools_public_access":         acctest.Representation{RepType: acctest.Optional, Create: `10.0.1.10`},
	}
	DifStackDataflowRepresentation = map[string]interface{}{
		"driver_shape":                 acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`, Update: `VM.Standard.E5.Flex`},
		"executor_shape":               acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`, Update: `VM.Standard.E5.Flex`},
		"instance_id":                  acctest.Representation{RepType: acctest.Required, Create: `testApp`},
		"log_bucket_instance_id":       acctest.Representation{RepType: acctest.Required, Create: `testBucket`},
		"num_executors":                acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"spark_version":                acctest.Representation{RepType: acctest.Required, Create: `3.5.0`, Update: `3.5.0`},
		"driver_shape_config":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackDataflowDriverShapeConfigRepresentation},
		"executor_shape_config":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackDataflowExecutorShapeConfigRepresentation},
		"warehouse_bucket_instance_id": acctest.Representation{RepType: acctest.Optional, Create: `testBucket`},
	}
	DifStackGenaiRepresentation = map[string]interface{}{
		"base_model":   acctest.Representation{RepType: acctest.Required, Create: `baseModel`},
		"cluster_type": acctest.Representation{RepType: acctest.Required, Create: `HOSTING`},
		"instance_id":  acctest.Representation{RepType: acctest.Required, Create: `genAiCluster1`},
		"oci_region":   acctest.Representation{RepType: acctest.Required, Create: `ociRegion`},
		"unit_count":   acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"endpoints":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackGenaiEndpointsRepresentation},
	}
	DifStackGgcsRepresentation = map[string]interface{}{
		"instance_id":        acctest.Representation{RepType: acctest.Required, Create: `testGgcsInstance`},
		"ocpu":               acctest.Representation{RepType: acctest.Required, Create: `1`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.ggcs_password_secret_id}`},
		"subnet_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_private_subnet.id}`},
		"ogg_version":        acctest.Representation{RepType: acctest.Optional, Create: `oggoracle:23.9.0.0.0_250903.1222_1314`},
	}
	DifStackObjectstorageRepresentation = map[string]interface{}{
		"instance_id":       acctest.Representation{RepType: acctest.Required, Create: `testBucket`},
		"object_versioning": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`},
		"storage_tier":      acctest.Representation{RepType: acctest.Required, Create: `STANDARD`},
		"auto_tiering":      acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},
	}
	DifStackDataflowConnectionsRepresentation = map[string]interface{}{
		"connection_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowConnectionDetailsRepresentation},
		"subnet_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_private_subnet.id}`},
	}
	DifStackDataflowDriverShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `16`, Update: `32`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
	}
	DifStackDataflowExecutorShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `16`, Update: `32`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
	}
	DifStackGenaiEndpointsRepresentation = map[string]interface{}{
		"endpoint_name":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_language_endpoint.test_endpoint.name}`},
		"is_content_moderation_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}
	DifStackGgcsConnectionsRepresentation = map[string]interface{}{
		"dif_dependencies":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackGgcsConnectionsDifDependenciesRepresentation},
		"gg_admin_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.gg_admin_password_secret_id}`},
	}
	DataflowConnectionDetailsRepresentation = map[string]interface{}{
		"dif_dependencies": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowDifDependenciesRepresentation},
		"domain_names":     acctest.Representation{RepType: acctest.Optional, Create: []string{`custpvtsubnet.oraclevcn.com`}, Update: []string{`custpvtsubnet.oraclevcn.com`, `db.custpvtsubnet.oraclevcn.com`}},
	}
	DifStackGgcsConnectionsDifDependenciesRepresentation = map[string]interface{}{
		"service_instance_id": acctest.Representation{RepType: acctest.Required, Create: `testAdb`},
		"service_type":        acctest.Representation{RepType: acctest.Required, Create: `ADB`},
	}
	DataflowDifDependenciesRepresentation = map[string]interface{}{
		"service_instance_id": acctest.Representation{RepType: acctest.Required, Create: `testAdb`},
		"service_type":        acctest.Representation{RepType: acctest.Required, Create: `ADB`},
	}

	MinimumValidDifStackAdbRepresentation = acctest.RepresentationCopyWithNewProperties(
		DifStackAdbRepresentation,
		map[string]interface{}{
			"is_public": acctest.Representation{RepType: acctest.Required, Create: `false`},
			"subnet_id": acctest.Representation{
				RepType: acctest.Required,
				Create:  `${oci_core_subnet.test_private_subnet.id}`,
			},
		},
	)

	DifStackResourceDependencies = compartmentIdVariableStr + compartmentIdMoveVariableStr + DefinedTagsDifDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{"dns_label": acctest.Representation{RepType: acctest.Required, Create: `testvcn`}})) +
		PrivateSubnetDifDependency
)

func checkServiceDetailFields(resourceName, serviceType, expectedInstanceID string, requireId, requireStatus bool, extra func(map[string]string, string) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		attrs := rs.Primary.Attributes
		countStr, ok := attrs["service_details.#"]
		if !ok {
			return fmt.Errorf("service_details not found in state")
		}
		n, _ := strconv.Atoi(countStr)
		for i := 0; i < n; i++ {
			prefix := fmt.Sprintf("service_details.%d.", i)
			if t, ok := attrs[prefix+"service_type"]; ok && t == serviceType {
				if expectedInstanceID != "" {
					if v, ok := attrs[prefix+"instance_id"]; !ok || v != expectedInstanceID {
						return fmt.Errorf("expected instance_id=%s for %s, got %s", expectedInstanceID, serviceType, v)
					}
				}
				if requireStatus {
					if v, ok := attrs[prefix+"status"]; !ok || v == "" {
						return fmt.Errorf("expected non-empty status for %s service_details", serviceType)
					}
				}
				if requireId {
					if v, ok := attrs[prefix+"service_id"]; !ok || v == "" {
						return fmt.Errorf("expected non-empty service_id for %s service_details", serviceType)
					}
				}
				if extra != nil {
					if err := extra(attrs, prefix); err != nil {
						return err
					}
				}
				return nil
			}
		}
		return fmt.Errorf("no service_details entry with service_type %s", serviceType)
	}
}

func checkServiceDetail(resourceName, serviceType, expectedInstanceID string) resource.TestCheckFunc {
	return checkServiceDetailFields(resourceName, serviceType, expectedInstanceID, true, true, nil)
}

// issue-routing-tag: dif/default
func TestDifStackResource_basic(t *testing.T) {
	t.Parallel()
	httpreplay.SetScenario("TestDifStackResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() +
		passwordSecretIdVariableStr +
		ggcsPasswordSecretIdVariableStr +
		ggAdminPasswordSecretIdVariableStr

	resourceName := "oci_dif_stack.test_stack"
	datasourceName := "data.oci_dif_stacks.test_stacks"
	singularDatasourceName := "data.oci_dif_stack.test_stack"

	var resId string

	// Ensure unique stack display name for parallel test runs
	stackDisplayName := "tfdif" + utils.RandomString(8, utils.CharsetLowerCaseWithoutDigits)

	MinimumValidDifStackRepresentation := acctest.RepresentationCopyWithNewProperties(
		DifStackRepresentation,
		map[string]interface{}{
			"display_name":  acctest.Representation{RepType: acctest.Required, Create: stackDisplayName},
			"objectstorage": acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackObjectstorageRepresentation},
			"adb":           acctest.RepresentationGroup{RepType: acctest.Required, Group: MinimumValidDifStackAdbRepresentation},
			"dataflow":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowRepresentation},
			"ggcs":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackGgcsRepresentation},
		},
	)

	// Ensure Flex shapes include required shape configs
	DifStackDataflowFlexRequired := acctest.RepresentationCopyWithNewProperties(
		DifStackDataflowRepresentation,
		map[string]interface{}{
			"driver_shape_config":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowDriverShapeConfigRepresentation},
			"executor_shape_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowExecutorShapeConfigRepresentation},
		},
	)
	MinimumValidDifStackRepresentationWithFlexShape := acctest.RepresentationCopyWithNewProperties(
		MinimumValidDifStackRepresentation,
		map[string]interface{}{
			"dataflow": acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowFlexRequired},
		},
	)
	DifStackRepresentationWithFlexShape := acctest.RepresentationCopyWithNewProperties(
		DifStackRepresentation,
		map[string]interface{}{
			"display_name":    acctest.Representation{RepType: acctest.Required, Create: stackDisplayName},
			"services":        acctest.Representation{RepType: acctest.Required, Create: []string{`OBJECTSTORAGE`, `ADB`, `DATAFLOW`, `GGCS`}, Update: []string{`OBJECTSTORAGE`, `ADB`, `DATAFLOW`, `GGCS`}},
			"stack_templates": acctest.Representation{RepType: acctest.Required, Create: []string{`DATALAKE`, `DATATRANSFORMATION`, `DATAPIPELINE`}, Update: []string{`DATALAKE`, `DATATRANSFORMATION`, `DATAPIPELINE`}},
			"dataflow":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowFlexRequired},
			"objectstorage":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackObjectstorageRepresentation},
		},
	)

	// Use the actual created stack's display name for the list data source to avoid filtering out the item.
	DifStackDataSourceRepresentationWithName := acctest.RepresentationCopyWithNewProperties(
		DifStackDataSourceRepresentation,
		map[string]interface{}{
			"display_name": acctest.Representation{RepType: acctest.Optional, Create: stackDisplayName},
		},
	)

	DifStackRepresentationWithFlexShapeNoEmail := acctest.
		RepresentationCopyWithRemovedProperties(
			DifStackRepresentationWithFlexShape,
			[]string{"notification_email"},
		)

	// Save TF content to Create resource with optional properties.
	acctest.SaveConfigContent(config+DifStackResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Optional, acctest.Create, MinimumValidDifStackRepresentationWithFlexShape), "dif", "stack", t)
	acctest.ResourceTest(t, testAccCheckDifStackDestroy, []resource.TestStep{
		{
			Config: config + DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Create, MinimumValidDifStackRepresentationWithFlexShape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", testCompartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", stackDisplayName),
				resource.TestCheckResourceAttr(resourceName, "services.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "stack_templates.#", "3"),
				func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] Step 1: verify Create - Starting")
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					log.Printf("[DEBUG] Step 1: verify Create - Completed, resId: %s", resId)
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Optional, acctest.Update, DifStackRepresentationWithFlexShapeNoEmail) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dif_stacks", "test_stacks", acctest.Optional, acctest.Update, DifStackDataSourceRepresentationWithName),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", testCompartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", stackDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "state", string(oci_dif.StackLifecycleStateActive)),

				resource.TestCheckResourceAttr(datasourceName, "stack_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "stack_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Create, DifStackSingularDataSourceRepresentation) +
				(DifStackResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Optional, acctest.Update, DifStackRepresentationWithFlexShapeNoEmail)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", testCompartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", stackDisplayName),

				resource.TestCheckResourceAttr(singularDatasourceName, "adb.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dataflow.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "genai.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ggcs.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "objectstorage.#", "0"),

				checkServiceDetail(singularDatasourceName, "OBJECTSTORAGE", "testBucket"),
				checkServiceDetail(singularDatasourceName, "ADB", "testAdb"),
				checkServiceDetail(singularDatasourceName, "GGCS", "testGgcsInstance"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DifStackRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"adb",
				"dataflow",
				"genai",
				"ggcs",
				"objectstorage",
				"add_service_trigger",
				"deploy_artifacts_trigger",
				"notification_email",
				"defined_tags",
				"freeform_tags",
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: dif/default
func TestDifStackResource_dataflow_with_connections(t *testing.T) {
	t.Parallel()
	httpreplay.SetScenario("TestDifStackResource_dataflow_with_connections")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() +
		passwordSecretIdVariableStr

	resourceName := "oci_dif_stack.test_stack"

	//unique stack display name for parallel test runs
	stackDisplayName := "tfdif" + utils.RandomString(8, utils.CharsetLowerCaseWithoutDigits)

	DifStackDataflowConnectionDetailsRepresentation := map[string]interface{}{
		"dif_dependencies": acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowDifDependenciesRepresentation},
		"domain_names":     acctest.Representation{RepType: acctest.Required, Create: []string{`custpvtsubnet.oraclevcn.com`}, Update: []string{`custpvtsubnet.oraclevcn.com`, `db.custpvtsubnet.oraclevcn.com`}},
	}

	DifStackDataflowConnectionsRepresentation := map[string]interface{}{
		"connection_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowConnectionDetailsRepresentation},
		"subnet_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_private_subnet.id}`},
	}

	DifStackDataflowRepresentationWithConnections := acctest.RepresentationCopyWithNewProperties(
		DifStackDataflowRepresentation,
		map[string]interface{}{
			"connections":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowConnectionsRepresentation},
			"driver_shape_config":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowDriverShapeConfigRepresentation},
			"executor_shape_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowExecutorShapeConfigRepresentation},
		},
	)

	MinimumValidDifStackRepresentationWithConnections := map[string]interface{}{
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: stackDisplayName},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"services":        acctest.Representation{RepType: acctest.Required, Create: []string{`OBJECTSTORAGE`, `ADB`, `DATAFLOW`}},
		"stack_templates": acctest.Representation{RepType: acctest.Required, Create: []string{`DATALAKE`, `DATATRANSFORMATION`}},
		"objectstorage":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackObjectstorageRepresentation},
		"adb":             acctest.RepresentationGroup{RepType: acctest.Required, Group: MinimumValidDifStackAdbRepresentation},
		"dataflow":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowRepresentationWithConnections},
	}

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+DifStackResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Optional, acctest.Create, MinimumValidDifStackRepresentationWithConnections), "dif", "stack", t)
	acctest.ResourceTest(t, testAccCheckDifStackDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Create, MinimumValidDifStackRepresentationWithConnections),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", testCompartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", stackDisplayName),
				resource.TestCheckResourceAttr(resourceName, "services.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "stack_templates.#", "2"),
				checkServiceDetail(resourceName, "OBJECTSTORAGE", "testBucket"),
				checkServiceDetail(resourceName, "ADB", "testAdb"),
				checkServiceDetailFields(resourceName, "DATAFLOW", "testApp", false, false, nil),
				checkServiceDetailFields(resourceName, "DATAFLOW", "testApp", false, false, func(attrs map[string]string, prefix string) error {
					if v, ok := attrs[prefix+"additional_details.0.private_endpoint_id"]; !ok || v == "" {
						return fmt.Errorf("expected private_endpoint_id for DATAFLOW service_details")
					}
					return nil
				}),
			),
		},

		// delete before next Create
		{
			Config: config + DifStackResourceDependencies,
		},
		// verify Create with Optionals
		{
			Config: config + DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Optional, acctest.Create, MinimumValidDifStackRepresentationWithConnections),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", testCompartmentId),
				resource.TestCheckResourceAttr(resourceName, "services.#", "3"),
				checkServiceDetail(resourceName, "OBJECTSTORAGE", "testBucket"),
				checkServiceDetail(resourceName, "ADB", "testAdb"),
				checkServiceDetailFields(resourceName, "DATAFLOW", "testApp", false, false, nil),

				func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] Test completed: DataFlow with connections works")
					return err
				},
			),
		},
	})
}

// issue-routing-tag: dif/default
func TestDifStackResource_adb_with_ggcs_connections(t *testing.T) {
	t.Parallel()
	httpreplay.SetScenario("TestDifStackResource_adb_with_ggcs_connections")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() +
		passwordSecretIdVariableStr +
		ggcsPasswordSecretIdVariableStr +
		ggAdminPasswordSecretIdVariableStr
	resourceName := "oci_dif_stack.test_stack"
	// Ensure unique stack display name for parallel test runs
	stackDisplayName := "tfdif" + utils.RandomString(8, utils.CharsetLowerCaseWithoutDigits)

	// Use GGCS "connections" (no connection_id) minimal
	// Ensure gg_admin_secret_id present when difDependencies are used
	DifStackGgcsConnectionsDifDependencies := acctest.RepresentationCopyWithNewProperties(
		DifStackGgcsConnectionsRepresentation,
		map[string]interface{}{
			"dif_dependencies":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackGgcsConnectionsDifDependenciesRepresentation},
			"connection_id":      acctest.Representation{RepType: acctest.Optional, Create: nil},
			"connection_name":    acctest.Representation{RepType: acctest.Required, Create: dif_ggcs_connection_name + stackDisplayName},
			"gg_admin_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.gg_admin_password_secret_id}`},
		},
	)
	DifStackGgcsWithConnections := acctest.RepresentationCopyWithNewProperties(
		DifStackGgcsRepresentation,
		map[string]interface{}{
			"connections": acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackGgcsConnectionsDifDependencies},
		},
	)

	MinimumValidDifStackRepresentationWithGgcsConnections := acctest.RepresentationCopyWithNewProperties(
		DifStackRepresentation,
		map[string]interface{}{
			"display_name":    acctest.Representation{RepType: acctest.Required, Create: stackDisplayName},
			"services":        acctest.Representation{RepType: acctest.Required, Create: []string{`ADB`, `GGCS`}},
			"stack_templates": acctest.Representation{RepType: acctest.Required, Create: []string{`DATALAKE`, `DATAPIPELINE`}},
			"objectstorage":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: nil},
			"adb":             acctest.RepresentationGroup{RepType: acctest.Required, Group: MinimumValidDifStackAdbRepresentation},
			"ggcs":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackGgcsWithConnections},
			"dataflow":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: nil},
		},
	)

	acctest.SaveConfigContent(config+DifStackResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Create, MinimumValidDifStackRepresentationWithGgcsConnections), "dif", "stack", t)

	acctest.ResourceTest(t, testAccCheckDifStackDestroy, []resource.TestStep{
		{
			Config: config + DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Create, MinimumValidDifStackRepresentationWithGgcsConnections),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", testCompartmentId),
				resource.TestCheckResourceAttr(resourceName, "services.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "stack_templates.#", "2"),
				checkServiceDetail(resourceName, "ADB", "testAdb"),
				checkServiceDetailFields(resourceName, "GGCS", "testGgcsInstance", true, true, func(attrs map[string]string, prefix string) error {
					// additionalDetails with assigned_connections should be present
					if v, ok := attrs[prefix+"additional_details.#"]; !ok || v == "" {
						return fmt.Errorf("expected additional_details for GGCS")
					}
					if v, ok := attrs[prefix+"additional_details.0.assigned_connections.#"]; ok {
						if cnt, _ := strconv.Atoi(v); cnt < 1 {
							return fmt.Errorf("expected at least one assigned_connection for GGCS")
						}
					} else {
						return fmt.Errorf("expected assigned_connections for GGCS")
					}
					return nil
				}),
			),
		},
	})
}

func TestDifStackResource_add_services_deploy_updates(t *testing.T) {
	t.Parallel()
	httpreplay.SetScenario("TestDifStackResource_add_services_deploy_updates")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() +
		passwordSecretIdVariableStr +
		ggcsPasswordSecretIdVariableStr +
		dataflowFileUriVariableStr +
		dataflowArchiveUriVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DataflowPrivateEndpointRepresentation, map[string]interface{}{"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_private_subnet.id}`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(GoldenGateConnectionRepresentation, map[string]interface{}{
			"host":         acctest.Representation{RepType: acctest.Required, Create: `10.0.0.127`},
			"port":         acctest.Representation{RepType: acctest.Required, Create: `12`},
			"display_name": acctest.Representation{RepType: acctest.Required, Create: dif_ggcs_connection_name},
		}))

	resourceName := "oci_dif_stack.test_stack"
	// Ensure unique stack display name for parallel test runs
	stackDisplayName := "tfdif" + utils.RandomString(8, utils.CharsetLowerCaseWithoutDigits)

	DifStackInitialStackRepresentation := acctest.RepresentationCopyWithNewProperties(
		DifStackRepresentation,
		map[string]interface{}{
			"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id_for_move}`},
			"display_name":        acctest.Representation{RepType: acctest.Required, Create: stackDisplayName},
			"notification_email":  acctest.Representation{RepType: acctest.Required, Create: `test@example.com`},
			"services":            acctest.Representation{RepType: acctest.Required, Create: []string{`OBJECTSTORAGE`}},
			"stack_templates":     acctest.Representation{RepType: acctest.Required, Create: []string{`DATALAKE`}},
			"objectstorage":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackObjectstorageRepresentation},
			"add_service_trigger": acctest.Representation{RepType: acctest.Required, Create: `1`},
			"adb":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: nil},
			"dataflow":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: nil},
			"ggcs":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: nil},
		},
	)

	DifStackGgcsConnectionsWithConnectionId := map[string]interface{}{
		"connection_name": acctest.Representation{RepType: acctest.Required, Create: dif_ggcs_connection_name},
		"connection_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_connection.test_connection.id}`},
	}

	DifStackGgcsWithConnectionId := acctest.RepresentationCopyWithNewProperties(
		DifStackGgcsRepresentation,
		map[string]interface{}{
			"connections": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackGgcsConnectionsWithConnectionId},
		},
	)

	DifStackDataflowWithFlexShapeConfigs := acctest.RepresentationCopyWithNewProperties(
		DifStackDataflowRepresentation,
		map[string]interface{}{
			"private_endpoint_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_dataflow_private_endpoint.test_private_endpoint.id}`},
			"driver_shape_config":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowDriverShapeConfigRepresentation},
			"executor_shape_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowExecutorShapeConfigRepresentation},
		},
	)
	DifStackAddServicesRepresentation := acctest.RepresentationCopyWithNewProperties(
		DifStackInitialStackRepresentation,
		map[string]interface{}{
			"add_service_trigger": acctest.Representation{RepType: acctest.Required, Update: `2`},
			"ggcs":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackGgcsWithConnectionId},
			"dataflow":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DifStackDataflowWithFlexShapeConfigs},
			"services":            acctest.Representation{RepType: acctest.Required, Create: []string{`OBJECTSTORAGE`, `DATAFLOW`, `GGCS`}},
			"stack_templates":     acctest.Representation{RepType: acctest.Required, Create: []string{`DATALAKE`, `DATATRANSFORMATION`, `DATAPIPELINE`}},
		},
	)
	DifStackDeployArtifactsRepresentation := acctest.RepresentationCopyWithNewProperties(
		DifStackAddServicesRepresentation,
		map[string]interface{}{
			"deploy_artifacts_trigger": acctest.Representation{RepType: acctest.Required, Update: `2`},
			"dataflow": acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(DifStackDataflowWithFlexShapeConfigs, map[string]interface{}{
				"execute":     acctest.Representation{RepType: acctest.Required, Create: `${var.dataflow_file_uri}`, Update: `${var.dataflow_file_uri}`},
				"archive_uri": acctest.Representation{RepType: acctest.Required, Create: `${var.dataflow_archive_uri}`, Update: `${var.dataflow_archive_uri}`},
			})},
		},
	)
	DifStackAddADBRepresentation := acctest.RepresentationCopyWithNewProperties(
		DifStackDeployArtifactsRepresentation,
		map[string]interface{}{
			"add_service_trigger": acctest.Representation{RepType: acctest.Required, Update: `3`},
			"services":            acctest.Representation{RepType: acctest.Required, Create: []string{`OBJECTSTORAGE`, `DATAFLOW`, `GGCS`}, Update: []string{`OBJECTSTORAGE`, `DATAFLOW`, `GGCS`, `ADB`}},
			"stack_templates":     acctest.Representation{RepType: acctest.Required, Create: []string{`DATALAKE`, `DATATRANSFORMATION`, `DATAPIPELINE`}, Update: []string{`DATALAKE`, `DATATRANSFORMATION`, `DATAPIPELINE`}},
			"adb":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: MinimumValidDifStackAdbRepresentation},
		},
	)

	acctest.SaveConfigContent(config+DifStackResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Create, DifStackInitialStackRepresentation), "dif", "stack", t)

	acctest.ResourceTest(t, testAccCheckDifStackDestroy, []resource.TestStep{
		// 1) Create ObjectStorage-only
		{
			Config: config + DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Create, DifStackInitialStackRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "stack_templates.#", "1"),
				checkServiceDetail(resourceName, "OBJECTSTORAGE", "testBucket"),
			),
		},
		// 2) Change compartment
		{
			Config: config + DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DifStackInitialStackRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_move}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", testCompartmentIdForMove),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "stack_templates.#", "1"),
				checkServiceDetail(resourceName, "OBJECTSTORAGE", "testBucket"),
			),
		},
		// 3) AddService: add GGCS and DataFlow
		{
			Config: config + DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Update, DifStackAddServicesRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "services.#", "3"),
				checkServiceDetail(resourceName, "OBJECTSTORAGE", "testBucket"),
				checkServiceDetailFields(resourceName, "DATAFLOW", "testApp", false, false, nil),
				checkServiceDetail(resourceName, "GGCS", "testGgcsInstance"),
			),
		},
		// 4) Deploy artifacts via dataflow.execute
		{
			Config: config + DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Update, DifStackDeployArtifactsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				checkServiceDetailFields(resourceName, "DATAFLOW", "testApp", true, true, func(attrs map[string]string, prefix string) error {
					if v, ok := attrs[prefix+"current_artifact_path"]; !ok || v != dataflow_file_uri {
						return fmt.Errorf("expected current_artifact_path to equal %s, got %s", dataflow_file_uri, v)
					}
					return nil
				}),
			),
		},
		// 5) AddService: add ADB
		{
			Config: config + DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Required, acctest.Update,
					DifStackAddADBRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "services.#", "4"),
				checkServiceDetail(resourceName, "ADB", "testAdb"),
			),
		},
		// 6) Update optionalslog
		{
			Config: config + DifStackResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dif_stack", "test_stack", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DifStackAddADBRepresentation, map[string]interface{}{
						// update some optionals (ecpu, data_storage_size_in_tbs, object_versioning, shape configs)
						"adb":           acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(MinimumValidDifStackAdbRepresentation, map[string]interface{}{"ecpu": acctest.Representation{RepType: acctest.Required, Update: `2`}, "data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Update: `2`}})},
						"dataflow":      acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(DifStackDataflowWithFlexShapeConfigs, map[string]interface{}{"driver_shape_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DifStackDataflowDriverShapeConfigRepresentation}})},
						"objectstorage": acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(DifStackObjectstorageRepresentation, map[string]interface{}{"object_versioning": acctest.Representation{RepType: acctest.Required, Update: `SUSPENDED`}})},
						"ggcs":          acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(DifStackGgcsWithConnectionId, map[string]interface{}{"ocpu": acctest.Representation{RepType: acctest.Required, Update: `2`}})},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// Validate updated optional fields only
				// ADB updates
				resource.TestCheckResourceAttr(resourceName, "adb.0.ecpu", "2"),
				resource.TestCheckResourceAttr(resourceName, "adb.0.data_storage_size_in_tbs", "2"),
				// Dataflow driver shape config updates
				resource.TestCheckResourceAttr(resourceName, "dataflow.0.driver_shape_config.0.memory_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "dataflow.0.driver_shape_config.0.ocpus", "2"),
				// ObjectStorage versioning update
				resource.TestCheckResourceAttr(resourceName, "objectstorage.0.object_versioning", "SUSPENDED"),
			),
		},
	})
}

func testAccCheckDifStackDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dif_stack" {
			noResourceFound = false
			request := oci_dif.GetStackRequest{}

			tmp := rs.Primary.ID
			request.StackId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dif")

			response, err := client.GetStack(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dif.StackLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DifStack") {
		resource.AddTestSweepers("DifStack", &resource.Sweeper{
			Name:         "DifStack",
			Dependencies: acctest.DependencyGraph["stack"],
			F:            sweepDifStackResource,
		})
	}
}

func sweepDifStackResource(compartment string) error {
	stackClient := acctest.GetTestClients(&schema.ResourceData{}).StackClient()
	stackIds, err := getDifStackIds(compartment)
	if err != nil {
		return err
	}
	for _, stackId := range stackIds {
		if ok := acctest.SweeperDefaultResourceId[stackId]; !ok {
			deleteStackRequest := oci_dif.DeleteStackRequest{}

			deleteStackRequest.StackId = &stackId

			deleteStackRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dif")
			_, error := stackClient.DeleteStack(context.Background(), deleteStackRequest)
			if error != nil {
				fmt.Printf("Error deleting Stack %s %s, It is possible that the resource is already deleted. Please verify manually \n", stackId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &stackId, DifStackSweepWaitCondition, time.Duration(10*time.Minute),
				DifStackSweepResponseFetchOperation, "dif", true)
		}
	}
	return nil
}

func getDifStackIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "StackId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackClient := acctest.GetTestClients(&schema.ResourceData{}).StackClient()

	listStacksRequest := oci_dif.ListStacksRequest{}
	listStacksRequest.CompartmentId = &compartmentId
	listStacksRequest.LifecycleState = oci_dif.StackLifecycleStateActive
	listStacksResponse, err := stackClient.ListStacks(context.Background(), listStacksRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Stack list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, stack := range listStacksResponse.Items {
		id := *stack.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "StackId", id)
	}
	return resourceIds, nil
}

func DifStackSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if stackResponse, ok := response.Response.(oci_dif.GetStackResponse); ok {
		return stackResponse.LifecycleState != oci_dif.StackLifecycleStateDeleted
	}
	return false
}

func DifStackSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackClient().GetStack(context.Background(), oci_dif.GetStackRequest{
		StackId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
