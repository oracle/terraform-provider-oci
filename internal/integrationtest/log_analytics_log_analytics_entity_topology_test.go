// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	logAnalyticsEntityTopologySingularDataSourceRepresentation = map[string]interface{}{
		"log_analytics_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_log_analytics_log_analytics_entity.test_entity.id}`},
		"namespace":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	logAnalyticsEntityForTopologyRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"entity_type_name":  acctest.Representation{RepType: acctest.Required, Create: `Host (Linux)`},
		"name":              acctest.Representation{RepType: acctest.Required, Create: `TF_LA_ENTITY`},
		"namespace":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"cloud_resource_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"hostname":          acctest.Representation{RepType: acctest.Optional, Create: `hostname`},
		"source_id":         acctest.Representation{RepType: acctest.Optional, Create: `source1`},
		"timezone_region":   acctest.Representation{RepType: acctest.Optional, Create: `PST8PDT`},
	}

	LoganObjectStoreDependency = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace",
		acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation)

	LogAnalyticsEntityTopologyResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_entity",
		acctest.Required, acctest.Create, logAnalyticsEntityForTopologyRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsEntityTopologyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsEntityTopologyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	requiredDatasourceName := "data.oci_log_analytics_log_analytics_entity_topology.test_entity_topology_required"
	optionalDatasourceName := "data.oci_log_analytics_log_analytics_entity_topology.test_entity_topology_optional"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify required inputs
		{
			Config: config + LoganObjectStoreDependency +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_topology",
					"test_entity_topology_required",
					acctest.Required, acctest.Create, logAnalyticsEntityTopologySingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsEntityTopologyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "log_analytics_entity_id"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(requiredDatasourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(requiredDatasourceName, "items.0.nodes.0.items.#", "1"),
				resource.TestCheckResourceAttr(requiredDatasourceName, "items.0.nodes.0.items.0.name", "TF_LA_ENTITY"),
				resource.TestCheckResourceAttr(requiredDatasourceName, "items.0.links.#", "1"),
			),
		},
		// verify optionals
		{
			Config: config + LoganObjectStoreDependency +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_topology",
					"test_entity_topology_optional",
					acctest.Optional, acctest.Create, logAnalyticsEntityTopologySingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsEntityTopologyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "log_analytics_entity_id"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "items.0.nodes.0.items.#", "1"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "items.0.nodes.0.items.0.name", "TF_LA_ENTITY"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "items.0.links.#", "1"),
			),
		},
	})
}
