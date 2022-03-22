// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"

	oci_data_safe "github.com/oracle/oci-go-sdk/v63/datasafe"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AlertRequiredOnlyResource = AlertResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert", "test_alert", acctest.Required, acctest.Create, alertRepresentation)

	AlertResourceConfig = AlertResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert", "test_alert", acctest.Optional, acctest.Update, alertRepresentation)

	alertSingularDataSourceRepresentation = map[string]interface{}{
		"alert_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_alert.test_alert.id}`},
	}

	alertDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"field":                     acctest.Representation{RepType: acctest.Optional, Create: []oci_data_safe.ListAlertsFieldEnum{`severity`}},
		"id":                        acctest.Representation{RepType: acctest.Optional, Create: `${var.alert_id}`},
		"scim_query":                acctest.Representation{RepType: acctest.Optional, Create: `severity eq 'HIGH'`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: alertDataSourceFilterRepresentation},
	}
	alertDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_alert.test_alert.id}`}},
	}

	alertRepresentation = map[string]interface{}{
		"alert_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.alert_id}`},
		"comment":       acctest.Representation{RepType: acctest.Required, Create: `comment`, Update: `comment2`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"status":        acctest.Representation{RepType: acctest.Optional, Create: `OPEN`, Update: `CLOSED`},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreFeatureDetailsSystemTagsChangesRep},
	}

	ignoreFeatureDetailsSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `feature_details`}},
	}
	AlertResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeAlertResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for Alert resource")
	httpreplay.SetScenario("TestDataSafeAlertResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	alertId := utils.GetEnvSettingWithBlankDefault("data_safe_alert_ocid")
	alertIdVariableStr := fmt.Sprintf("variable \"alert_id\" { default = \"%s\" }\n", alertId)

	resourceName := "oci_data_safe_alert.test_alert"
	datasourceName := "data.oci_data_safe_alerts.test_alerts"
	singularDatasourceName := "data.oci_data_safe_alert.test_alert"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+alertIdVariableStr+AlertResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert", "test_alert", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(alertRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "alert", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + alertIdVariableStr + AlertRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + alertIdVariableStr + AlertResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + alertIdVariableStr + AlertResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert", "test_alert", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(alertRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_id"),
				resource.TestCheckResourceAttr(resourceName, "comment", "comment"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "OPEN"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + alertIdVariableStr + AlertResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert", "test_alert", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(alertRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "alert_id"),
				resource.TestCheckResourceAttr(resourceName, "comment", "comment2"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "CLOSED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alerts", "test_alerts", acctest.Optional, acctest.Update, alertDataSourceRepresentation) +
				compartmentIdVariableStr + alertIdVariableStr + AlertResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_alert", "test_alert", acctest.Optional, acctest.Update, alertRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "scim_query", "severity eq 'HIGH'"),
				resource.TestCheckResourceAttr(datasourceName, "alert_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert", "test_alert", acctest.Required, acctest.Create, alertSingularDataSourceRepresentation) +
				compartmentIdVariableStr + alertIdVariableStr + AlertResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "comment", "comment2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operation"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operation_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operation_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "severity"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "CLOSED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_names.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + alertIdVariableStr + AlertResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`alert_id`},
			ResourceName:            resourceName,
		},
	})
}
