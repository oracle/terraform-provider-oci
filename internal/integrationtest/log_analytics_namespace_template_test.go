// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LogAnalyticsNamespaceTemplateSingularDataSourceRepresentation = map[string]interface{}{
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"template_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_log_analytics_namespace_templates.test_namespace_templates.log_analytics_template_collection[0].items[0].id}`},
	}

	LogAnalyticsNamespaceTemplateDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"namespace":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"name":                  acctest.Representation{RepType: acctest.Optional, Create: `Linux ROOT Logins`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"template_display_text": acctest.Representation{RepType: acctest.Optional, Create: `root`},
		"type":                  acctest.Representation{RepType: acctest.Optional, Create: `Scheduled Search`},
	}

	LogAnalyticsNamespaceTemplateResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceTemplateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceTemplateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_log_analytics_namespace_templates.test_namespace_templates"
	singularDatasourceName := "data.oci_log_analytics_namespace_template.test_namespace_template"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource with required attributes
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_templates", "test_namespace_templates", acctest.Required, acctest.Create, LogAnalyticsNamespaceTemplateDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceTemplateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

				resource.TestCheckResourceAttrSet(datasourceName, "log_analytics_template_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "log_analytics_template_collection.0.items.0.name"),
			),
		},
		// verify datasource with optional attributes
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_templates", "test_namespace_templates", acctest.Optional, acctest.Create, LogAnalyticsNamespaceTemplateDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceTemplateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

				resource.TestCheckResourceAttr(datasourceName, "name", "Linux ROOT Logins"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "template_display_text", "root"),
				resource.TestCheckResourceAttr(datasourceName, "type", "Scheduled Search"),

				resource.TestCheckResourceAttr(datasourceName, "log_analytics_template_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "log_analytics_template_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "log_analytics_template_collection.0.items.0.name", "Linux ROOT Logins"),
				resource.TestCheckResourceAttr(datasourceName, "log_analytics_template_collection.0.items.0.is_system", "true"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_templates", "test_namespace_templates", acctest.Optional, acctest.Create, LogAnalyticsNamespaceTemplateDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_template", "test_namespace_template", acctest.Required, acctest.Create, LogAnalyticsNamespaceTemplateSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceTemplateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "template_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content_format"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "facets.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_system", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "Linux ROOT Logins"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parameters"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parameters_format"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parameters_metadata"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "Scheduled Search"),
			),
		},
	})
}
