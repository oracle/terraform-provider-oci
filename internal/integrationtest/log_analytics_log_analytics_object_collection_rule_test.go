// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	LogAnalyticsLogAnalyticsObjectCollectionRuleRequiredOnlyResource = LogAnalyticsLogAnalyticsObjectCollectionRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsObjectCollectionRuleRepresentation)

	LogAnalyticsLogAnalyticsObjectCollectionRuleResourceConfig = LogAnalyticsLogAnalyticsObjectCollectionRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsObjectCollectionRuleRepresentation)

	LogAnalyticsLogAnalyticsLogAnalyticsObjectCollectionRuleSingularDataSourceRepresentation = map[string]interface{}{
		"log_analytics_object_collection_rule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_log_analytics_log_analytics_object_collection_rule.test_log_analytics_object_collection_rule.id}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsLogAnalyticsLogAnalyticsObjectCollectionRuleDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `test_terraform_rule`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsLogAnalyticsObjectCollectionRuleDataSourceFilterRepresentation}}
	LogAnalyticsLogAnalyticsObjectCollectionRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_log_analytics_log_analytics_object_collection_rule.test_log_analytics_object_collection_rule.id}`}},
	}

	logAnalyticsObjectCollectionRulePropertyOverridesRepresentation = map[string]interface{}{
		"match_type":     acctest.Representation{RepType: acctest.Optional, Create: `contains`, Update: `contains`},
		"match_value":    acctest.Representation{RepType: acctest.Optional, Create: `db`, Update: `db`},
		"property_name":  acctest.Representation{RepType: acctest.Optional, Create: `charEncoding`, Update: `charEncoding`},
		"property_value": acctest.Representation{RepType: acctest.Optional, Create: `utf-8`, Update: `utf-16`},
	}

	LogAnalyticsLogAnalyticsObjectCollectionRuleRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"log_group_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.log_analytics_log_group_id}`},
		"log_source_name":              acctest.Representation{RepType: acctest.Required, Create: `LinuxSyslogSource`, Update: `LinuxSyslogSource`},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: `test_terraform_rule`},
		"namespace":                    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"os_bucket_name":               acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"os_namespace":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"char_encoding":                acctest.Representation{RepType: acctest.Optional, Create: `utf-8`, Update: `utf-16`},
		"collection_type":              acctest.Representation{RepType: acctest.Optional, Create: `LIVE`},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `test terraform rule create`, Update: `test terraform rule update`},
		"entity_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.log_analytics_entity_id}`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: nil},
		"is_force_historic_collection": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"log_set":                      acctest.Representation{RepType: acctest.Optional, Create: `logSet`, Update: `logSet2`},
		"log_type":                     acctest.Representation{RepType: acctest.Optional, Create: `LOG`},
		"object_name_filters":          acctest.Representation{RepType: acctest.Optional, Create: []string{`objectNameFilters`}, Update: []string{`objectNameFilters2`}},
		"timezone":                     acctest.Representation{RepType: acctest.Optional, Create: `Asia/Dhaka`, Update: `America/New_York`},
		"overrides":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: logAnalyticsObjectCollectionRulePropertyOverridesRepresentation},
	}

	// Log Analytics Log Group and Log Analytics Entity dependencies are removed and values are provided as environment variables.
	LogAnalyticsLogAnalyticsObjectCollectionRuleResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation)
)

func TestLogAnalyticsLogAnalyticsObjectCollectionRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsObjectCollectionRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	logGroupId := utils.GetEnvSettingWithBlankDefault("log_analytics_log_group_id")
	if logGroupId == "" {
		t.Skip("Set log_analytics_log_group_id to run this test")
	}
	logGroupIdVariableStr := fmt.Sprintf("variable \"log_analytics_log_group_id\" { default = \"%s\" }\n", logGroupId)

	entityId := utils.GetEnvSettingWithBlankDefault("log_analytics_entity_id")
	if logGroupId == "" {
		t.Skip("Set log_analytics_entity_id to run this test")
	}
	entityIdVariableStr := fmt.Sprintf("variable \"log_analytics_entity_id\" { default = \"%s\" }\n", entityId)

	resourceName := "oci_log_analytics_log_analytics_object_collection_rule.test_log_analytics_object_collection_rule"
	datasourceName := "data.oci_log_analytics_log_analytics_object_collection_rules.test_log_analytics_object_collection_rules"
	singularDatasourceName := "data.oci_log_analytics_log_analytics_object_collection_rule.test_log_analytics_object_collection_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+logGroupIdVariableStr+entityIdVariableStr+LogAnalyticsLogAnalyticsObjectCollectionRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsObjectCollectionRuleRepresentation), "loganalytics", "logAnalyticsObjectCollectionRule", t)

	acctest.ResourceTest(t, testAccCheckLogAnalyticsLogAnalyticsObjectCollectionRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + logGroupIdVariableStr + entityIdVariableStr + LogAnalyticsLogAnalyticsObjectCollectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsObjectCollectionRuleRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_source_name", "LinuxSyslogSource"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "os_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "os_namespace"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + logGroupIdVariableStr + entityIdVariableStr + LogAnalyticsLogAnalyticsObjectCollectionRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + logGroupIdVariableStr + entityIdVariableStr + LogAnalyticsLogAnalyticsObjectCollectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsObjectCollectionRuleRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "char_encoding", "utf-8"),
				resource.TestCheckResourceAttr(resourceName, "collection_type", "LIVE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "test terraform rule create"),
				resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_force_historic_collection", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_set", "logSet"),
				resource.TestCheckResourceAttr(resourceName, "log_source_name", "LinuxSyslogSource"),
				resource.TestCheckResourceAttr(resourceName, "log_type", "LOG"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object_name_filters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "os_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "os_namespace"),
				resource.TestCheckResourceAttr(resourceName, "overrides.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_name", "charEncoding"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_value", "utf-8"),
				resource.TestCheckResourceAttrSet(resourceName, "poll_since"),
				resource.TestCheckNoResourceAttr(resourceName, "poll_till"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "timezone", "Asia/Dhaka"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + logGroupIdVariableStr + entityIdVariableStr + LogAnalyticsLogAnalyticsObjectCollectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsLogAnalyticsObjectCollectionRuleRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "char_encoding", "utf-8"),
				resource.TestCheckResourceAttr(resourceName, "collection_type", "LIVE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "test terraform rule create"),
				resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_force_historic_collection", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_set", "logSet"),
				resource.TestCheckResourceAttr(resourceName, "log_source_name", "LinuxSyslogSource"),
				resource.TestCheckResourceAttr(resourceName, "log_type", "LOG"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object_name_filters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "os_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "os_namespace"),
				resource.TestCheckResourceAttr(resourceName, "overrides.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_name", "charEncoding"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_value", "utf-8"),
				resource.TestCheckResourceAttrSet(resourceName, "poll_since"),
				resource.TestCheckNoResourceAttr(resourceName, "poll_till"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "timezone", "Asia/Dhaka"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						utils.Logf("TestLogAnalyticsLogAnalyticsObjectCollectionRuleResource:: resource Ids not matching \n%s\n%s", fmt.Sprintf(resId), fmt.Sprintf(resId2))
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + logGroupIdVariableStr + entityIdVariableStr + LogAnalyticsLogAnalyticsObjectCollectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsObjectCollectionRuleRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "char_encoding", "utf-16"),
				resource.TestCheckResourceAttr(resourceName, "collection_type", "LIVE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "test terraform rule update"),
				resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_force_historic_collection", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_set", "logSet2"),
				resource.TestCheckResourceAttr(resourceName, "log_source_name", "LinuxSyslogSource"),
				resource.TestCheckResourceAttr(resourceName, "log_type", "LOG"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object_name_filters.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "os_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "os_namespace"),
				resource.TestCheckResourceAttr(resourceName, "overrides.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_name", "charEncoding"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_value", "utf-16"),
				resource.TestCheckResourceAttrSet(resourceName, "poll_since"),
				resource.TestCheckNoResourceAttr(resourceName, "poll_till"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "timezone", "America/New_York"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						utils.Logf("TestLogAnalyticsLogAnalyticsObjectCollectionRuleResource:: resource Ids not matching \n%s\n%s", fmt.Sprintf(resId), fmt.Sprintf(resId2))
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rules", "test_log_analytics_object_collection_rules", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsLogAnalyticsObjectCollectionRuleDataSourceRepresentation) +
				compartmentIdVariableStr + logGroupIdVariableStr + entityIdVariableStr + LogAnalyticsLogAnalyticsObjectCollectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsObjectCollectionRuleRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "log_analytics_object_collection_rule_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "log_analytics_object_collection_rule_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "log_analytics_object_collection_rule_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "log_analytics_object_collection_rule_collection.0.items.0.is_enabled", "true"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsLogAnalyticsObjectCollectionRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + logGroupIdVariableStr + entityIdVariableStr + LogAnalyticsLogAnalyticsObjectCollectionRuleResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_analytics_object_collection_rule_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttr(singularDatasourceName, "char_encoding", "utf-16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "collection_type", "LIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "test terraform rule update"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_force_historic_collection", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_set", "logSet2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_source_name", "LinuxSyslogSource"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_type", "LOG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_name_filters.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "overrides.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "overrides.0.property_name", "charEncoding"),
				resource.TestCheckResourceAttr(singularDatasourceName, "overrides.0.property_value", "utf-16"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "poll_since"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "poll_till"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timezone", "America/New_York"),
			),
		},
		// verify resource import
		{
			Config:                  config + LogAnalyticsLogAnalyticsObjectCollectionRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getLogAnalyticsObjectCollectionRulesEndpointImportId(resourceName),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLogAnalyticsLogAnalyticsObjectCollectionRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LogAnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_log_analytics_log_analytics_object_collection_rule" {
			noResourceFound = false
			request := oci_log_analytics.GetLogAnalyticsObjectCollectionRuleRequest{}

			tmp := rs.Primary.ID
			request.LogAnalyticsObjectCollectionRuleId = &tmp

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

			response, err := client.GetLogAnalyticsObjectCollectionRule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_log_analytics.ObjectCollectionRuleLifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("LogAnalyticsLogAnalyticsObjectCollectionRule") {
		resource.AddTestSweepers("LogAnalyticsLogAnalyticsObjectCollectionRule", &resource.Sweeper{
			Name:         "LogAnalyticsLogAnalyticsObjectCollectionRule",
			Dependencies: acctest.DependencyGraph["logAnalyticsObjectCollectionRule"],
			F:            sweepLogAnalyticsLogAnalyticsObjectCollectionRuleResource,
		})
	}
}

func sweepLogAnalyticsLogAnalyticsObjectCollectionRuleResource(compartment string) error {
	logAnalyticsClient := acctest.GetTestClients(&schema.ResourceData{}).LogAnalyticsClient()
	logAnalyticsObjectCollectionRuleIds, err := getLogAnalyticsLogAnalyticsObjectCollectionRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, logAnalyticsObjectCollectionRuleId := range logAnalyticsObjectCollectionRuleIds {
		if ok := acctest.SweeperDefaultResourceId[logAnalyticsObjectCollectionRuleId]; !ok {
			deleteLogAnalyticsObjectCollectionRuleRequest := oci_log_analytics.DeleteLogAnalyticsObjectCollectionRuleRequest{}

			deleteLogAnalyticsObjectCollectionRuleRequest.LogAnalyticsObjectCollectionRuleId = &logAnalyticsObjectCollectionRuleId

			deleteLogAnalyticsObjectCollectionRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")
			_, error := logAnalyticsClient.DeleteLogAnalyticsObjectCollectionRule(context.Background(), deleteLogAnalyticsObjectCollectionRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting LogAnalyticsObjectCollectionRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", logAnalyticsObjectCollectionRuleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &logAnalyticsObjectCollectionRuleId, LogAnalyticsLogAnalyticsObjectCollectionRuleSweepWaitCondition, time.Duration(3*time.Minute),
				LogAnalyticsLogAnalyticsObjectCollectionRuleSweepResponseFetchOperation, "log_analytics", true)
		}
	}
	return nil
}

func getLogAnalyticsLogAnalyticsObjectCollectionRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LogAnalyticsObjectCollectionRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	logAnalyticsClient := acctest.GetTestClients(&schema.ResourceData{}).LogAnalyticsClient()

	listLogAnalyticsObjectCollectionRulesRequest := oci_log_analytics.ListLogAnalyticsObjectCollectionRulesRequest{}
	listLogAnalyticsObjectCollectionRulesRequest.CompartmentId = &compartmentId

	namespaces, error := getNamespaces(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting namespace required for LogAnalyticsObjectCollectionRule resource requests \n")
	}
	for _, namespace := range namespaces {
		listLogAnalyticsObjectCollectionRulesRequest.NamespaceName = &namespace

		listLogAnalyticsObjectCollectionRulesRequest.LifecycleState = oci_log_analytics.ListLogAnalyticsObjectCollectionRulesLifecycleStateActive
		listLogAnalyticsObjectCollectionRulesResponse, err := logAnalyticsClient.ListLogAnalyticsObjectCollectionRules(context.Background(), listLogAnalyticsObjectCollectionRulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting LogAnalyticsObjectCollectionRule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, logAnalyticsObjectCollectionRule := range listLogAnalyticsObjectCollectionRulesResponse.Items {
			id := *logAnalyticsObjectCollectionRule.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LogAnalyticsObjectCollectionRuleId", id)
		}

	}
	return resourceIds, nil
}

func LogAnalyticsLogAnalyticsObjectCollectionRuleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if logAnalyticsObjectCollectionRuleResponse, ok := response.Response.(oci_log_analytics.GetLogAnalyticsObjectCollectionRuleResponse); ok {
		return logAnalyticsObjectCollectionRuleResponse.LifecycleState != oci_log_analytics.ObjectCollectionRuleLifecycleStatesDeleted
	}
	return false
}

func LogAnalyticsLogAnalyticsObjectCollectionRuleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.LogAnalyticsClient().GetLogAnalyticsObjectCollectionRule(context.Background(), oci_log_analytics.GetLogAnalyticsObjectCollectionRuleRequest{
		LogAnalyticsObjectCollectionRuleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func getLogAnalyticsObjectCollectionRulesEndpointImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("namespaces/" + rs.Primary.Attributes["namespace"] + "/logAnalyticsObjectCollectionRules/" + rs.Primary.Attributes["id"]), nil
	}
}
