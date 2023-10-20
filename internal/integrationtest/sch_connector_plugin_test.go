// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	serviceConnectorRepresentationNoTargetPluginSource = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `My_Service_Connector`, Update: `displayName`},
		"source":         acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorQueueSourceRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `My service connector description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"tasks":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: SchServiceConnectorTasksRepresentation},
	}

	serviceConnectorQueueSourceRepresentation = map[string]interface{}{
		"kind":        acctest.Representation{RepType: acctest.Required, Create: `plugin`},
		"plugin_name": acctest.Representation{RepType: acctest.Required, Create: `QueueSource`},
		"config_map":  acctest.Representation{RepType: acctest.Required, Create: `{\"queueId\":\"${var.queue_id}\"}`},
	}

	serviceConnectorFunctionBatchTargetQueueRepresentation = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTargetPluginSource, functionTargetBatchRepresentation)
)

// issue-routing-tag: sch/default
func TestSchConnectorPluginResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSchConnectorPluginResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	testQueueId := utils.GetEnvSettingWithBlankDefault("queue_ocid")
	queueIdVariableStr := fmt.Sprintf("variable \"queue_id\" { default = \"%s\" }\n", testQueueId)

	image := utils.GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	var resId string

	//resourceName := "oci_sch_connector_plugins.test_connector_plugins"
	resourceName := "oci_sch_service_connector.test_service_connector"

	acctest.ResourceTest(t, testAccCheckSchServiceConnectorDestroy, []resource.TestStep{

		// verify Create with queue source and function as target with batching
		{
			Config: config + compartmentIdVariableStr + queueIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorFunctionBatchTargetQueueRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "plugin"),
				resource.TestCheckResourceAttr(resourceName, "source.0.plugin_name", "QueueSource"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.config_map"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "functions"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.function_id"),
				resource.TestCheckResourceAttr(resourceName, "target.0.batch_size_in_kbs", "6144"),
				resource.TestCheckResourceAttr(resourceName, "target.0.batch_size_in_num", "6291456"),
				resource.TestCheckResourceAttr(resourceName, "target.0.batch_time_in_sec", "600"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					_ = resId
					return err
				},
			),
		},

		// verify resource import
		{
			Config:                  config + SchServiceConnectorRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
