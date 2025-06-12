// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LogAnalyticsLogAnalyticsEntityTypeRequiredOnlyResource = LogAnalyticsLogAnalyticsEntityTypeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_type", "test_log_analytics_entity_type", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsEntityTypeRepresentation)

	LogAnalyticsLogAnalyticsEntityTypeDataSourceRepresentation = map[string]interface{}{
		"namespace":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"entity_type_name": acctest.Representation{RepType: acctest.Required, Create: `OCI TF Test Entity Type`}}

	LogAnalyticsLogAnalyticsEntityTypeDataSourceFilterRepresentation = map[string]interface{}{
		"entity_type_name": acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values":           acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_log_analytics_log_analytics_entity_type.test_log_analytics_entity_type.name}`}},
	}

	LogAnalyticsLogAnalyticsEntityTypeRepresentation = map[string]interface{}{
		"name":       acctest.Representation{RepType: acctest.Required, Create: `OCI TF Test Entity Type`},
		"namespace":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"category":   acctest.Representation{RepType: acctest.Optional, Create: `CUSTOM`},
		"properties": acctest.RepresentationGroup{RepType: acctest.Optional, Group: LogAnalyticsLogAnalyticsEntityTypePropertiesRepresentation},
	}

	LogAnalyticsLogAnalyticsEntityTypePropertiesRepresentation = map[string]interface{}{
		"name":        acctest.Representation{RepType: acctest.Required, Create: `propertyName`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `propertyDescription`},
	}

	LogAnalyticsLogAnalyticsEntityTypeResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsEntityTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsEntityTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_entity_type.test_log_analytics_entity_type"
	datasourceName := "data.oci_log_analytics_log_analytics_entity_type.test_log_analytics_entity_type"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsLogAnalyticsEntityTypeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_type", "test_log_analytics_entity_type", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsEntityTypeRepresentation), "loganalytics", "logAnalyticsEntityType", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsEntityTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_type", "test_log_analytics_entity_type", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsEntityTypeRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "OCI TF Test Entity Type"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				// resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsEntityTypeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsEntityTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_type", "test_log_analytics_entity_type", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsEntityTypeRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category", "CUSTOM"),
				// resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "OCI TF Test Entity Type"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.description", "propertyDescription"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.name", "propertyName"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_type", "test_log_analytics_entity_type", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsEntityTypeDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsLogAnalyticsEntityTypeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_type", "test_log_analytics_entity_type", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsEntityTypeRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(datasourceName, "cloud_type", "CLOUD"),
				resource.TestCheckResourceAttr(datasourceName, "entity_type_name", "OCI TF Test Entity Type"),
				// resource.TestCheckResourceAttr(datasourceName, "name_contains", "OCI TF Test"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace")),
		},
		// verify resource import
		{
			Config: config + LogAnalyticsLogAnalyticsEntityTypeRequiredOnlyResource,

			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"category",
				"name",
				"properties",
			},
			ResourceName: resourceName,
		},
	})
}
