// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LogAnalyticsNamespaceLookupRequiredOnlyResource = LogAnalyticsNamespaceLookupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Required, acctest.Create, LogAnalyticsNamespaceLookupRequiredRepresentation)

	LogAnalyticsNamespaceLookupResourceConfig = LogAnalyticsNamespaceLookupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Optional, acctest.Update, LogAnalyticsNamespaceLookupRepresentation)

	LogAnalyticsNamespaceLookupSingularDataSourceRepresentation = map[string]interface{}{
		"lookup_name": acctest.Representation{RepType: acctest.Required, Create: `TFLookup`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsNamespaceLookupRequiredRepresentation = map[string]interface{}{
		"lookup_name":          acctest.Representation{RepType: acctest.Required, Create: `TFLookup`},
		"namespace":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                 acctest.Representation{RepType: acctest.Required, Create: `Lookup`},
		"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: ``},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsNamespaceLookupIgnoreChangesRepresentation},
	}

	LogAnalyticsNamespaceLookupRepresentation = map[string]interface{}{
		"lookup_name":          acctest.Representation{RepType: acctest.Required, Create: `TFLookup`},
		"namespace":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                 acctest.Representation{RepType: acctest.Required, Create: `Lookup`},
		"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: ``},
		"char_encoding":        acctest.Representation{RepType: acctest.Optional, Create: `UTF-8`},
		"categories":           []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: LogAnalyticsNamespaceLookupCategoriesRepresentation1}, {RepType: acctest.Optional, Group: LogAnalyticsNamespaceLookupCategoriesRepresentation2}},
		"default_match_value":  acctest.Representation{RepType: acctest.Optional, Create: `WILDCARD`, Update: `EXACT`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":          acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"fields":               []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: LogAnalyticsNamespaceLookupFieldsRepresentation1}, {RepType: acctest.Optional, Group: LogAnalyticsNamespaceLookupFieldsRepresentation2}},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_hidden":            acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"max_matches":          acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsNamespaceLookupIgnoreChangesRepresentation},
	}
	LogAnalyticsNamespaceLookupIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `fields`}},
	}
	LogAnalyticsNamespaceLookupCategoriesRepresentation1 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Optional, Create: `linux`, Update: `oracle`},
	}
	LogAnalyticsNamespaceLookupCategoriesRepresentation2 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Optional, Create: `cloud`, Update: `database`},
	}
	LogAnalyticsNamespaceLookupFieldsRepresentation1 = map[string]interface{}{
		"match_operator": acctest.Representation{RepType: acctest.Optional, Create: `WILDCARD`, Update: `WILDCARD`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name`},
	}
	LogAnalyticsNamespaceLookupFieldsRepresentation2 = map[string]interface{}{
		"match_operator": acctest.Representation{RepType: acctest.Optional, Create: `WILDCARD`, Update: `WILDCARD`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `profitmodel`, Update: `profitmodel`},
	}

	LogAnalyticsNamespaceLookupResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceLookupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceLookupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_log_analytics_namespace_lookup.test_namespace_lookup"
	singularDatasourceName := "data.oci_log_analytics_namespace_lookup.test_namespace_lookup"

	lookupFilePath, err := createLookupFile()
	if err != nil {
		t.Fatalf("Unable to create file to upload. Error: %q", err)
	}

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+LogAnalyticsNamespaceLookupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupRepresentation, map[string]interface{}{
				"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
				"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
			})), "loganalytics", "namespaceLookup", t)

	acctest.ResourceTest(t, testAccCheckLogAnalyticsNamespaceLookupDestroy, []resource.TestStep{
		// verify Create with required attributes
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceLookupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupRequiredRepresentation, map[string]interface{}{
						"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "lookup_name", "TFLookup"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "lookup_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "type", "Lookup"),

				resource.TestCheckResourceAttrSet(resourceName, "active_edit_version"),
				resource.TestCheckResourceAttrSet(resourceName, "edit_version"),
				resource.TestCheckResourceAttr(resourceName, "fields.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "is_built_in", "0"),
				resource.TestCheckResourceAttr(resourceName, "is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "lookup_display_name", "TFLookup"),
				resource.TestCheckResourceAttrSet(resourceName, "lookup_reference"),
				resource.TestCheckResourceAttrSet(resourceName, "lookup_reference_string"),
				resource.TestCheckResourceAttr(resourceName, "status_summary.0.status", "successful"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceLookupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LogAnalyticsNamespaceLookupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupRepresentation, map[string]interface{}{
						"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "lookup_name", "TFLookup"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "lookup_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "type", "Lookup"),

				resource.TestCheckResourceAttrSet(resourceName, "active_edit_version"),
				resource.TestCheckResourceAttr(resourceName, "categories.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "default_match_value", "WILDCARD"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "edit_version"),
				resource.TestCheckResourceAttr(resourceName, "fields.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_built_in", "0"),
				resource.TestCheckResourceAttr(resourceName, "is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "lookup_display_name", "TFLookup"),
				resource.TestCheckResourceAttrSet(resourceName, "lookup_reference"),
				resource.TestCheckResourceAttrSet(resourceName, "lookup_reference_string"),
				resource.TestCheckResourceAttr(resourceName, "status_summary.0.status", "successful"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "max_matches", "10"),

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
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceLookupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupRepresentation, map[string]interface{}{
						"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
						"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "lookup_name", "TFLookup"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "lookup_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "type", "Lookup"),

				resource.TestCheckResourceAttrSet(resourceName, "active_edit_version"),
				resource.TestCheckResourceAttr(resourceName, "categories.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "default_match_value", "EXACT"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "edit_version"),
				resource.TestCheckResourceAttr(resourceName, "fields.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_built_in", "0"),
				resource.TestCheckResourceAttr(resourceName, "is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "lookup_display_name", "TFLookup"),
				resource.TestCheckResourceAttrSet(resourceName, "lookup_reference"),
				resource.TestCheckResourceAttrSet(resourceName, "lookup_reference_string"),
				resource.TestCheckResourceAttr(resourceName, "status_summary.0.status", "successful"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "max_matches", "11"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Required, acctest.Create, LogAnalyticsNamespaceLookupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceLookupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "lookup_name", "TFLookup"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "active_edit_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "categories.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "edit_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fields.#", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_built_in", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_hidden", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lookup_display_name", "TFLookup"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lookup_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lookup_reference"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lookup_reference_string"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status_summary.0.status", "successful"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "Lookup"),
			),
		},
		// verify resource import
		{
			Config:            config + LogAnalyticsNamespaceLookupRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"default_match_value",
				"max_matches",
				"char_encoding",
				"register_lookup_file",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckLogAnalyticsNamespaceLookupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LogAnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_log_analytics_namespace_lookup" {
			noResourceFound = false
			request := oci_log_analytics.GetLookupRequest{}

			if value, ok := rs.Primary.Attributes["lookup_name"]; ok {
				request.LookupName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

			_, err := client.GetLookup(context.Background(), request)

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

func createLookupFile() (string, error) {
	lookupFile, err := ioutil.TempFile(os.TempDir(), "lookup_*.csv")
	if err != nil {
		return "", err
	}

	content := "id,name,city,profitmodel" + "\n" +
		"V001,711 LLC,Hartford,5050" + "\n" +
		"V002,Electra Corp,New Jersey,6040" + "\n" +
		"V003,Mutual Smiles LLC,Greenfield,7030" + "\n" +
		"V004,La Marvella,Hershire,5050" + "\n" +
		"V005,Cuppa Coffee,New Jersey,5050"

	text := []byte(content)
	if _, err = lookupFile.Write(text); err != nil {
		return "", err
	}

	// Close the file
	if err := lookupFile.Close(); err != nil {
		return "", err
	}

	return lookupFile.Name(), nil
}
