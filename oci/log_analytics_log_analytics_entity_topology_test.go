// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	logAnalyticsEntityTopologySingularDataSourceRepresentation = map[string]interface{}{
		"log_analytics_entity_id": Representation{RepType: Required, Create: `${oci_log_analytics_log_analytics_entity.test_entity.id}`},
		"namespace":               Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"state":                   Representation{RepType: Optional, Create: `ACTIVE`},
	}

	logAnalyticsEntityForTopologyRepresentation = map[string]interface{}{
		"compartment_id":    Representation{RepType: Required, Create: `${var.compartment_id}`},
		"entity_type_name":  Representation{RepType: Required, Create: `Host (Linux)`},
		"name":              Representation{RepType: Required, Create: `TF_LA_ENTITY`},
		"namespace":         Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"cloud_resource_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"hostname":          Representation{RepType: Optional, Create: `hostname`},
		"source_id":         Representation{RepType: Optional, Create: `source1`},
		"timezone_region":   Representation{RepType: Optional, Create: `PST8PDT`},
	}

	LoganObjectStoreDependency = GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace",
		Required, Create, namespaceSingularDataSourceRepresentation)

	LogAnalyticsEntityTopologyResourceConfig = GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_entity",
		Required, Create, logAnalyticsEntityForTopologyRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsEntityTopologyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsEntityTopologyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	requiredDatasourceName := "data.oci_log_analytics_log_analytics_entity_topology.test_entity_topology_required"
	optionalDatasourceName := "data.oci_log_analytics_log_analytics_entity_topology.test_entity_topology_optional"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify required inputs
		{
			Config: config + LoganObjectStoreDependency +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_topology",
					"test_entity_topology_required",
					Required, Create, logAnalyticsEntityTopologySingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsEntityTopologyResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_topology",
					"test_entity_topology_optional",
					Optional, Create, logAnalyticsEntityTopologySingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsEntityTopologyResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
