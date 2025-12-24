// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LogAnalyticsNamespaceAssociationRequiredOnlyResource = LogAnalyticsNamespaceAssociationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_association", "test_namespace_association", acctest.Required, acctest.Create, LogAnalyticsNamespaceAssociationRepresentation)

	LogAnalyticsNamespaceAssociationRepresentation = map[string]interface{}{
		"entity_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.entity_id}`},
		"log_group_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.log_group_id}`, Update: `${var.log_group_id_updated}`},
		"namespace":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"source_name":            acctest.Representation{RepType: acctest.Required, Create: `${var.source_name}`},
		"association_properties": acctest.RepresentationGroup{RepType: acctest.Optional, Group: LogAnalyticsNamespaceAssociationAssociationPropertiesRepresentation},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_create}`},
		"is_from_republish":      acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
	}
	LogAnalyticsNamespaceAssociationAssociationPropertiesRepresentation = map[string]interface{}{
		"name":     acctest.Representation{RepType: acctest.Required, Create: `management_agent.os_file.timezone`},
		"patterns": acctest.RepresentationGroup{RepType: acctest.Optional, Group: LogAnalyticsNamespaceAssociationItemsAssociationPropertiesPatternsRepresentation},
		"value":    acctest.Representation{RepType: acctest.Optional, Create: `IST`, Update: `GMT`},
	}

	LogAnalyticsNamespaceAssociationItemsAssociationPropertiesPatternsRepresentation = map[string]interface{}{
		"id":    acctest.Representation{RepType: acctest.Required, Create: `${var.pattern_id}`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `PST`},
	}

	LogAnalyticsNamespaceAssociationResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id_for_create\" { default = \"%s\" }\n", compartmentId)

	entityId := utils.GetEnvSettingWithBlankDefault("entity_id")
	entityIdVariableStr := fmt.Sprintf("variable \"entity_id\" { default = \"%s\" }\n", entityId)

	logGroupId := utils.GetEnvSettingWithBlankDefault("log_group_id")
	logGroupIdVariableStr := fmt.Sprintf("variable \"log_group_id\" { default = \"%s\" }\n", logGroupId)

	logGroupIdU := utils.GetEnvSettingWithBlankDefault("log_group_id_updated")
	logGroupIdVariableStrU := fmt.Sprintf("variable \"log_group_id_updated\" { default = \"%s\" }\n", logGroupIdU)

	sourceName := utils.GetEnvSettingWithBlankDefault("source_name")
	sourceNameVariableStr := fmt.Sprintf("variable \"source_name\" { default = \"%s\" }\n", sourceName)

	patternId := utils.GetEnvSettingWithBlankDefault("pattern_id")
	patternIdVariableStr := fmt.Sprintf("variable \"pattern_id\" { default = \"%s\" }\n", patternId)

	resourceName := "oci_log_analytics_namespace_association.test_namespace_association"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+entityIdVariableStr+logGroupIdVariableStr+patternIdVariableStr+LogAnalyticsNamespaceAssociationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_association", "test_namespace_association", acctest.Optional, acctest.Create, LogAnalyticsNamespaceAssociationRepresentation), "loganalytics", "NamespaceAssociation", t)

	acctest.ResourceTest(t, testAccCheckLogAnalyticsNamespaceAssociationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + entityIdVariableStr + logGroupIdVariableStr + sourceNameVariableStr + LogAnalyticsNamespaceAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_association", "test_namespace_association", acctest.Required, acctest.Create, LogAnalyticsNamespaceAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "source_name", sourceName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "entity_id", entityId),
				resource.TestCheckResourceAttr(resourceName, "log_group_id", logGroupId),
			),
		},
		// delete before next Create
		{
			Config: config + LogAnalyticsNamespaceAssociationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + entityIdVariableStr + logGroupIdVariableStr + sourceNameVariableStr + patternIdVariableStr + LogAnalyticsNamespaceAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_association", "test_namespace_association", acctest.Optional, acctest.Create, LogAnalyticsNamespaceAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "association_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "association_properties.0.name", "management_agent.os_file.timezone"),
				resource.TestCheckResourceAttr(resourceName, "association_properties.0.value", "IST"),
				resource.TestCheckResourceAttr(resourceName, "association_properties.0.patterns.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "association_properties.0.patterns.0.id"),
				resource.TestCheckResourceAttr(resourceName, "association_properties.0.patterns.0.value", "PST"),
				resource.TestCheckResourceAttr(resourceName, "entity_id", entityId),
				resource.TestCheckResourceAttr(resourceName, "log_group_id", logGroupId),
				resource.TestCheckResourceAttr(resourceName, "source_name", sourceName),
				resource.TestCheckResourceAttr(resourceName, "is_from_republish", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

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
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + entityIdVariableStr + logGroupIdVariableStr + logGroupIdVariableStrU + sourceNameVariableStr + patternIdVariableStr + LogAnalyticsNamespaceAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_association", "test_namespace_association", acctest.Optional, acctest.Update, LogAnalyticsNamespaceAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "log_group_id", logGroupIdU),
				resource.TestCheckResourceAttr(resourceName, "association_properties.0.value", "GMT"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "is_from_republish", "false"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config + LogAnalyticsNamespaceAssociationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getNamespaceAssociationEndpointImportId(resourceName),
			ImportStateVerifyIgnore: []string{"is_from_republish"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLogAnalyticsNamespaceAssociationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LogAnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_log_analytics_namespace_association" {
			noResourceFound = false
			request := oci_log_analytics.ListSourceAssociationsRequest{}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			if value, ok := rs.Primary.Attributes["source_name"]; ok {
				request.SourceName = &value
			}

			if value, ok := rs.Primary.Attributes["entity_id"]; ok {
				request.EntityId = &value
			}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

			response, err := client.ListSourceAssociations(context.Background(), request)

			if len(response.LogAnalyticsAssociationCollection.Items) == 0 {
				log.Println("Items is empty in Response. Resource has been deleted")
				return nil
			}

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

func getNamespaceAssociationEndpointImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf(
			"namespaces/%s/associations/%s/%s/%s",
			rs.Primary.Attributes["namespace"],
			rs.Primary.Attributes["compartment_id"],
			rs.Primary.Attributes["entity_id"],
			rs.Primary.Attributes["source_name"],
		), nil
	}
}
