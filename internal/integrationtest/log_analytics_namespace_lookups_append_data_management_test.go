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
	LogAnalyticsNamespaceLookupsAppendDataManagementRepresentation = map[string]interface{}{
		"append_lookup_file": acctest.Representation{RepType: acctest.Required, Create: ``},
		"lookup_name":        acctest.Representation{RepType: acctest.Required, Create: `TFLookup`},
		"namespace":          acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"char_encoding":      acctest.Representation{RepType: acctest.Optional, Create: `UTF-8`},
		"expect":             acctest.Representation{RepType: acctest.Optional, Create: `100-continue`},
		"is_force":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	LogAnalyticsNamespaceLookupsAppendDataManagementResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceLookupsAppendDataManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceLookupsAppendDataManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	lookupFilePath, err := createLookupFile()
	if err != nil {
		t.Fatalf("Unable to create file to upload. Error: %q", err)
	}

	resourceName := "oci_log_analytics_namespace_lookups_append_data_management.test_namespace_lookups_append_data_management"
	parentResourceName := "oci_log_analytics_namespace_lookup.test_namespace_lookup"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsNamespaceLookupsAppendDataManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupRequiredRepresentation, map[string]interface{}{
				"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
			}))+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookups_append_data_management", "test_namespace_lookups_append_data_management", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupsAppendDataManagementRepresentation, map[string]interface{}{
				"append_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
			})), "loganalytics", "namespaceLookupsAppendDataManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create lookup with required attributes
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceLookupsAppendDataManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupRequiredRepresentation, map[string]interface{}{
						"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "lookup_name", "TFLookup"),
				resource.TestCheckResourceAttr(parentResourceName, "type", "Lookup"),
				resource.TestCheckResourceAttr(parentResourceName, "fields.#", "4"),
				resource.TestCheckResourceAttr(parentResourceName, "status_summary.0.status", "successful"),
			),
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceLookupsAppendDataManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupRequiredRepresentation, map[string]interface{}{
						"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookups_append_data_management", "test_namespace_lookups_append_data_management", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupsAppendDataManagementRepresentation, map[string]interface{}{
						"append_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "append_lookup_file", lookupFilePath),
				resource.TestCheckResourceAttr(resourceName, "lookup_name", "TFLookup"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
			),
		},
		// Delete append resource
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceLookupsAppendDataManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupRequiredRepresentation, map[string]interface{}{
						"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "lookup_name", "TFLookup"),
				resource.TestCheckResourceAttr(parentResourceName, "type", "Lookup"),
				resource.TestCheckResourceAttr(parentResourceName, "fields.#", "4"),
				resource.TestCheckResourceAttr(parentResourceName, "status_summary.0.status", "successful"),
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceLookupsAppendDataManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookup", "test_namespace_lookup", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupRequiredRepresentation, map[string]interface{}{
						"register_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_lookups_append_data_management", "test_namespace_lookups_append_data_management", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceLookupsAppendDataManagementRepresentation, map[string]interface{}{
						"append_lookup_file": acctest.Representation{RepType: acctest.Required, Create: lookupFilePath},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "append_lookup_file", lookupFilePath),
				resource.TestCheckResourceAttr(resourceName, "char_encoding", "UTF-8"),
				resource.TestCheckResourceAttr(resourceName, "expect", "100-continue"),
				resource.TestCheckResourceAttr(resourceName, "is_force", "false"),
				resource.TestCheckResourceAttr(resourceName, "lookup_name", "TFLookup"),
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
		// Delete lookup
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceLookupsAppendDataManagementResourceDependencies,
		},
	})
}
