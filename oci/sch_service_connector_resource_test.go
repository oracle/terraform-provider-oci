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
	// streaming as a source definition
	serviceConnectorStreamingSourceCursorRepresentation = map[string]interface{}{
		"kind": Representation{repType: Optional, create: `LATEST`, update: `TRIM_HORIZON`},
	}

	serviceConnectorStreamingSourceRepresentation = map[string]interface{}{
		"kind":      Representation{repType: Required, create: `streaming`},
		"cursor":    RepresentationGroup{Optional, serviceConnectorStreamingSourceCursorRepresentation},
		"stream_id": Representation{repType: Required, create: `${oci_streaming_stream.test_stream.id}`},
	}

	// function as a task
	serviceConnectorFunctionTasksRepresentation = map[string]interface{}{
		"kind":              Representation{repType: Required, create: `function`},
		"batch_size_in_kbs": Representation{repType: Required, create: `60`},
		"batch_time_in_sec": Representation{repType: Required, create: `60`},
		"function_id":       Representation{repType: Required, create: `${oci_functions_function.test_function.id}`},
	}

	// create serviceConnector definitions
	serviceConnectorRepresentationNoTargetStreamingSource = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `My_Service_Connector`, update: `displayName2`},
		"source":         RepresentationGroup{Required, serviceConnectorStreamingSourceRepresentation},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `My service connector description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"tasks":          RepresentationGroup{Optional, serviceConnectorTasksRepresentation},
	}

	serviceConnectorRepresentationNoTargetStreamingSourceFunctionTask = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `My_Service_Connector`, update: `displayName2`},
		"source":         RepresentationGroup{Required, serviceConnectorStreamingSourceRepresentation},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `My service connector description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"tasks":          RepresentationGroup{Required, serviceConnectorFunctionTasksRepresentation},
	}

	// targets for streaming as a source
	serviceConnectorFunctionTargetStreamingSourceRepresentation             = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTargetStreamingSource, functionTargetRepresentation)
	serviceConnectorFunctionTargetStreamingSourceFunctionTaskRepresentation = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTargetStreamingSourceFunctionTask, functionTargetRepresentation)

	updatedServiceConnectorFunctionTasksRepresentation = map[string]interface{}{
		"kind":              Representation{repType: Optional, update: `function`},
		"batch_size_in_kbs": Representation{repType: Optional, update: `60`},
		"batch_time_in_sec": Representation{repType: Optional, update: `60`},
		"function_id":       Representation{repType: Optional, update: `${oci_functions_function.test_function.id}`},
	}

	updatedServiceConnectorStreamingSourceRepresentation = map[string]interface{}{
		"kind":      Representation{repType: Optional, update: `streaming`},
		"cursor":    RepresentationGroup{Optional, serviceConnectorStreamingSourceCursorRepresentation},
		"stream_id": Representation{repType: Optional, create: `${oci_streaming_stream.test_stream.id}`},
	}
)

func TestSchServiceConnectorResource_streamingAnalytics(t *testing.T) {
	httpreplay.SetScenario("TestSchServiceConnectorResource_streamingAnalytics")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	image := getEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	resourceName := "oci_sch_service_connector.test_service_connector"
	singularDatasourceName := "data.oci_sch_service_connector.test_service_connector"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},

		CheckDestroy: testAccCheckSchServiceConnectorDestroy,
		Steps: []resource.TestStep{
			// verify streaming as a source with functions target
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Required, Create, serviceConnectorFunctionTargetStreamingSourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "streaming"),
					resource.TestCheckResourceAttr(resourceName, "source.0.cursor.0.kind", "LATEST"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.stream_id"),
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

			// verify streaming as a source with functions task and functions target
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Required, Create, serviceConnectorFunctionTargetStreamingSourceFunctionTaskRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
					resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "streaming"),
					resource.TestCheckResourceAttr(resourceName, "source.0.cursor.0.kind", "LATEST"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.stream_id"),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "functions"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.function_id"),
					resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "tasks.0.function_id"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.batch_size_in_kbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.batch_time_in_sec", "60"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "function"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Update,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(serviceConnectorFunctionTargetStreamingSourceFunctionTaskRepresentation, []string{"target"}), map[string]interface{}{
							"source": RepresentationGroup{Optional, serviceConnectorStreamingSourceRepresentation},
							"target": RepresentationGroup{Required, updatedServiceConnectorTargetRepresentation},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source.0.kind", "streaming"),
					resource.TestCheckResourceAttr(resourceName, "source.0.cursor.0.kind", "TRIM_HORIZON"),
					resource.TestCheckResourceAttrSet(resourceName, "source.0.stream_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target.0.kind", "streaming"),
					resource.TestCheckResourceAttrSet(resourceName, "target.0.stream_id"),
					resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "tasks.0.function_id"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.batch_size_in_kbs", "60"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.batch_time_in_sec", "60"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "function"),
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

			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Required, Create, serviceConnectorSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ServiceConnectorResourceDependencies + imageVariableStr +
					generateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", Optional, Update,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"source", "task", "target"}), map[string]interface{}{
							"source": RepresentationGroup{Optional, updatedServiceConnectorStreamingSourceRepresentation},
							"tasks":  RepresentationGroup{Optional, updatedServiceConnectorFunctionTasksRepresentation},
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
					resource.TestCheckResourceAttr(singularDatasourceName, "source.0.kind", "streaming"),
					resource.TestCheckResourceAttr(singularDatasourceName, "source.0.cursor.0.kind", "TRIM_HORIZON"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "source.0.stream_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "target.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "target.0.kind", "streaming"),
					resource.TestCheckResourceAttr(singularDatasourceName, "tasks.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tasks.0.function_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.batch_size_in_kbs", "60"),
					resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.batch_time_in_sec", "60"),
					resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.kind", "function"),
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
