// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v40/common"
	oci_sch "github.com/oracle/oci-go-sdk/v40/sch"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	// Dependency definition
	ServiceConnectorResourceDependencies = generateResourceFromRepresentationMap("oci_logging_log", "test_log", Required, Create, logRepresentation) +
		generateResourceFromRepresentationMap("oci_logging_log", "test_update_log", Required, Update, getUpdatedRepresentationCopy("configuration.source.category", Representation{repType: Required, create: `read`}, logRepresentation)) +
		LogResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
		generateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation) +
		generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Required, Create, streamRepresentation) +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)

	// source definitions
	serviceConnectorSourceLogSourcesRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"log_group_id":   Representation{repType: Optional, create: `${oci_logging_log_group.test_log_group.id}`, update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_id":         Representation{repType: Optional, create: `${oci_logging_log.test_log.id}`, update: `${oci_logging_log.test_update_log.id}`},
	}

	serviceConnectorSourceRepresentation = map[string]interface{}{
		"kind":        Representation{repType: Required, create: `logging`},
		"log_sources": RepresentationGroup{Required, serviceConnectorSourceLogSourcesRepresentation},
	}

	serviceConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `My_Service_Connector`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, serviceConnectorDataSourceFilterRepresentation},
	}

	// task definitions
	serviceConnectorTasksRepresentation = map[string]interface{}{
		"condition": Representation{repType: Required, create: `data.action='REJECT'`, update: `logContent='20'`},
		"kind":      Representation{repType: Required, create: `logRule`},
	}

	// target definitions
	functionTargetRepresentation = map[string]interface{}{
		"kind":        Representation{repType: Required, create: `functions`},
		"function_id": Representation{repType: Required, create: `${oci_functions_function.test_function.id}`},
	}

	objectStorageTargetRepresentation = map[string]interface{}{
		"kind":                       Representation{repType: Required, create: `objectStorage`},
		"bucket":                     Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":                  Representation{repType: Optional, create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object_name_prefix":         Representation{repType: Optional, create: `test_prefix`},
		"batch_rollover_size_in_mbs": Representation{repType: Optional, create: `10`},
		"batch_rollover_time_in_ms":  Representation{repType: Optional, create: `80000`},
	}

	logAnTargetRepresentation = map[string]interface{}{
		"kind":         Representation{repType: Required, create: `loggingAnalytics`},
		"log_group_id": Representation{repType: Required, create: `${var.logAn_log_group_ocid}`},
	}

	onsTargetRepresentation = map[string]interface{}{
		"kind":                       Representation{repType: Required, create: `notifications`},
		"topic_id":                   Representation{repType: Required, create: `${oci_ons_notification_topic.test_notification_topic.id}`},
		"enable_formatted_messaging": Representation{repType: Optional, create: `true`},
	}

	// create serviceConnector definitions
	serviceConnectorRepresentationNoTarget = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `My_Service_Connector`, update: `displayName2`},
		"source":         RepresentationGroup{Required, serviceConnectorSourceRepresentation},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `My service connector description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"tasks":          RepresentationGroup{Optional, serviceConnectorTasksRepresentation},
	}

	// targets for logging as a source
	serviceConnectorFunctionTargetRepresentation      = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, functionTargetRepresentation)
	serviceConnectorObjectStorageTargetRepresentation = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, objectStorageTargetRepresentation)
	serviceConnectorLogAnTargetRepresentation         = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, logAnTargetRepresentation)
	serviceConnectorOnsTargetRepresentation           = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, onsTargetRepresentation)

	serviceConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"service_connector_id": Representation{repType: Required, create: `${oci_sch_service_connector.test_service_connector.id}`},
	}

	serviceConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_sch_service_connector.test_service_connector.id}`}},
	}

	// update serviceConnector definitions
	updatedServiceConnectorTargetRepresentation = map[string]interface{}{
		"kind":      Representation{repType: Required, create: `streaming`},
		"stream_id": Representation{repType: Optional, create: `${oci_streaming_stream.test_stream.id}`},
	}

	ServiceConnectorResourceConfig = ServiceConnectorResourceDependencies +
		generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Update, serviceConnectorFunctionTargetRepresentation)
)

func TestSchServiceConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSchServiceConnectorResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	image := getEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	logAnLogGroupId := getEnvSettingWithBlankDefault("logAn_log_group_ocid")
	logAnLogGroupIdVariableStr := fmt.Sprintf("variable \"logAn_log_group_ocid\" { default = \"%s\" }\n", logAnLogGroupId)

	resourceName := "oci_sch_service_connector.test_service_connector"
	datasourceName := "data.oci_sch_service_connectors.test_service_connectors"
	singularDatasourceName := "data.oci_sch_service_connector.test_service_connector"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ServiceConnectorResourceDependencies+imageVariableStr+
		generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Create, serviceConnectorObjectStorageTargetRepresentation), "sch", "serviceConnector", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},

		CheckDestroy: testAccCheckSchServiceConnectorDestroy,
		Steps: []resource.TestStep{
			// verify create with functions
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Required, Create, serviceConnectorFunctionTargetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "functions"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.function_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr,
			},

			// verify create with objectstorage
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Create, serviceConnectorObjectStorageTargetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "objectStorage"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.bucket"),
					resource.TestCheckResourceAttr(resourceName, "target.0.batch_rollover_size_in_mbs", "10"),
					resource.TestCheckResourceAttr(resourceName, "target.0.batch_rollover_time_in_ms", "80000"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr,
			},

			// verify create with log analytics
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr + logAnLogGroupIdVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Required, Create, serviceConnectorLogAnTargetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "loggingAnalytics"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.log_group_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr,
			},

			// verify create with ons
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Create, serviceConnectorOnsTargetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "My service connector description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "notifications"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.topic_id"),
					resource.TestCheckResourceAttr(resourceName, "target.0.enable_formatted_messaging", "true"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr,
			},

			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Create, serviceConnectorFunctionTargetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "My service connector description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "functions"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.function_id"),
					resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.condition", "data.action='REJECT'"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "logRule"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Create,
						representationCopyWithNewProperties(serviceConnectorFunctionTargetRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "My service connector description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "functions"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.function_id"),
					resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.condition", "data.action='REJECT'"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "logRule"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Update,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"target"}), map[string]interface{}{
							"target": RepresentationGroup{Required, updatedServiceConnectorTargetRepresentation},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "streaming"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.stream_id"),
					resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.condition", "logContent='20'"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "logRule"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// verify stop service connector
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Update,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"target"}), map[string]interface{}{
							"target": RepresentationGroup{Required, updatedServiceConnectorTargetRepresentation},
							"state":  Representation{repType: Optional, create: `INACTIVE`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "streaming"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.stream_id"),
					resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.condition", "logContent='20'"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "logRule"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// verify start service connector
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Update,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"target"}), map[string]interface{}{
							"target": RepresentationGroup{Required, updatedServiceConnectorTargetRepresentation},
							"state":  Representation{repType: Optional, create: `ACTIVE`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "streaming"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.stream_id"),
					resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.condition", "logContent='20'"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "logRule"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
					generateDataSourceFromRepresentationMap("oci_sch_service_connectors", "test_service_connectors", Optional, Update, serviceConnectorDataSourceRepresentation) +
					compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Update,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"target"}), map[string]interface{}{
							"target": RepresentationGroup{Required, updatedServiceConnectorTargetRepresentation},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "service_connector_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "service_connector_collection.0.items.#", "1"),
				),
			},

			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Required, Create, serviceConnectorSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Update,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"target"}), map[string]interface{}{
							"target": RepresentationGroup{Required, updatedServiceConnectorTargetRepresentation},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "service_connector_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source.0.kind", "logging"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source.0.log_sources.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source.0.log_sources.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "target.0.kind", "streaming"),
					resource.TestCheckResourceAttr(singularDatasourceName, "tasks.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.condition", "logContent='20'"),
					resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.kind", "logRule"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},

			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceConfig + imageVariableStr,
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

func testAccCheckSchServiceConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).serviceConnectorClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_sch_service_connector" {
			noResourceFound = false
			request := oci_sch.GetServiceConnectorRequest{}

			tmp := rs.Primary.ID
			request.ServiceConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "sch")

			response, err := client.GetServiceConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_sch.LifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("SchServiceConnector") {
		resource.AddTestSweepers("SchServiceConnector", &resource.Sweeper{
			Name:         "SchServiceConnector",
			Dependencies: DependencyGraph["serviceConnector"],
			F:            sweepSchServiceConnectorResource,
		})
	}
}

func sweepSchServiceConnectorResource(compartment string) error {
	serviceConnectorClient := GetTestClients(&schema.ResourceData{}).serviceConnectorClient()
	serviceConnectorIds, err := getServiceConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, serviceConnectorId := range serviceConnectorIds {
		if ok := SweeperDefaultResourceId[serviceConnectorId]; !ok {
			deleteServiceConnectorRequest := oci_sch.DeleteServiceConnectorRequest{}

			deleteServiceConnectorRequest.ServiceConnectorId = &serviceConnectorId

			deleteServiceConnectorRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "sch")
			_, error := serviceConnectorClient.DeleteServiceConnector(context.Background(), deleteServiceConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting ServiceConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", serviceConnectorId, error)
				continue
			}
			waitTillCondition(testAccProvider, &serviceConnectorId, serviceConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				serviceConnectorSweepResponseFetchOperation, "sch", true)
		}
	}
	return nil
}

func getServiceConnectorIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ServiceConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceConnectorClient := GetTestClients(&schema.ResourceData{}).serviceConnectorClient()

	listServiceConnectorsRequest := oci_sch.ListServiceConnectorsRequest{}
	listServiceConnectorsRequest.CompartmentId = &compartmentId
	listServiceConnectorsRequest.LifecycleState = oci_sch.ListServiceConnectorsLifecycleStateActive
	listServiceConnectorsResponse, err := serviceConnectorClient.ListServiceConnectors(context.Background(), listServiceConnectorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ServiceConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, serviceConnector := range listServiceConnectorsResponse.Items {
		id := *serviceConnector.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ServiceConnectorId", id)
	}
	return resourceIds, nil
}

func serviceConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if serviceConnectorResponse, ok := response.Response.(oci_sch.GetServiceConnectorResponse); ok {
		return serviceConnectorResponse.LifecycleState != oci_sch.LifecycleStateDeleted
	}
	return false
}

func serviceConnectorSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.serviceConnectorClient().GetServiceConnector(context.Background(), oci_sch.GetServiceConnectorRequest{
		ServiceConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func createServiceConnectorRepresentation(sc map[string]interface{}, target map[string]interface{}) map[string]interface{} {
	serviceConnector := make(map[string]interface{})

	// Copy map and populate target
	for key, value := range sc {
		serviceConnector[key] = value
	}
	serviceConnector["target"] = RepresentationGroup{Required, target}

	return serviceConnector
}
