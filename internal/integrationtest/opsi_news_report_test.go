// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OpsiNewsReportRequiredOnlyResource = OpsiNewsReportResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Required, acctest.Create, OpsiNewsReportRepresentation)

	OpsiNewsReportResourceConfig = OpsiNewsReportResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Optional, acctest.Update, OpsiNewsReportRepresentation)

	OpsiNewsReportSingularDataSourceRepresentation = map[string]interface{}{
		"news_report_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_news_report.test_news_report.id}`},
	}

	OpsiNewsReportDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"news_report_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_news_report.test_news_report.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`ENABLED`}, Update: []string{`DISABLED`}},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiNewsReportDataSourceFilterRepresentation},
	}

	OpsiNewsReportDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_news_report.test_news_report.id}`}},
	}

	OpsiNewsReportRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"content_types":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiNewsReportContentTypesRepresentation},
		"description":                     acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_DESCRIPTION`, Update: `TF_TEST_REPORT_DESCRIPTION_2`},
		"locale":                          acctest.Representation{RepType: acctest.Required, Create: `EN`},
		"name":                            acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_NAME`, Update: `TF_TEST_REPORT_NAME_2`},
		"news_frequency":                  acctest.Representation{RepType: acctest.Required, Create: `WEEKLY`},
		"ons_topic_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.topic_id}`},
		"are_child_compartments_included": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"day_of_week":                     acctest.Representation{RepType: acctest.Optional, Create: `MONDAY`, Update: `TUESDAY`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                          acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesNewsReportRepresentation},
	}

	OpsiNewsReportContentTypesRepresentation = map[string]interface{}{
		"capacity_planning_resources": acctest.Representation{RepType: acctest.Required, Create: []string{`HOST`, `DATABASE`}, Update: []string{`HOST`, `DATABASE`, `EXADATA`}},
	}

	OpsiNewsReportSqlInsightsFleetAnalysisRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"content_types":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiNewsReportSqlInsightsFleetAnalysisContentTypesRepresentation},
		"description":                     acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_DESCRIPTION`, Update: `TF_TEST_REPORT_DESCRIPTION_2`},
		"locale":                          acctest.Representation{RepType: acctest.Required, Create: `EN`},
		"name":                            acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_NAME`, Update: `TF_TEST_REPORT_NAME_2`},
		"news_frequency":                  acctest.Representation{RepType: acctest.Required, Create: `WEEKLY`},
		"ons_topic_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.topic_id}`},
		"are_child_compartments_included": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"day_of_week":                     acctest.Representation{RepType: acctest.Optional, Create: `MONDAY`, Update: `TUESDAY`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                          acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesNewsReportRepresentation},
	}

	OpsiNewsReportSqlInsightsFleetAnalysisContentTypesRepresentation = map[string]interface{}{
		"sql_insights_fleet_analysis_resources": acctest.Representation{RepType: acctest.Required, Create: []string{`DATABASE`}, Update: []string{`DATABASE`, `EXADATA`}},
	}

	OpsiNewsReportSqlInsightsPerformanceDegradationRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"content_types":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiNewsReportSqlInsightsPerformanceDegradationContentTypesRepresentation},
		"description":                     acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_DESCRIPTION`, Update: `TF_TEST_REPORT_DESCRIPTION_2`},
		"locale":                          acctest.Representation{RepType: acctest.Required, Create: `EN`},
		"name":                            acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_NAME`, Update: `TF_TEST_REPORT_NAME_2`},
		"news_frequency":                  acctest.Representation{RepType: acctest.Required, Create: `WEEKLY`},
		"ons_topic_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.topic_id}`},
		"are_child_compartments_included": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"day_of_week":                     acctest.Representation{RepType: acctest.Optional, Create: `MONDAY`, Update: `TUESDAY`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                          acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesNewsReportRepresentation},
	}

	OpsiNewsReportSqlInsightsPerformanceDegradationContentTypesRepresentation = map[string]interface{}{
		"sql_insights_performance_degradation_resources": acctest.Representation{RepType: acctest.Required, Create: []string{`DATABASE`}, Update: []string{`DATABASE`, `EXADATA`}},
	}

	OpsiNewsReportSqlInsightsPlanChangesRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"content_types":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiNewsReportSqlInsightsPlanChangesContentTypesRepresentation},
		"description":                     acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_DESCRIPTION`, Update: `TF_TEST_REPORT_DESCRIPTION_2`},
		"locale":                          acctest.Representation{RepType: acctest.Required, Create: `EN`},
		"name":                            acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_NAME`, Update: `TF_TEST_REPORT_NAME_2`},
		"news_frequency":                  acctest.Representation{RepType: acctest.Required, Create: `WEEKLY`},
		"ons_topic_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.topic_id}`},
		"are_child_compartments_included": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"day_of_week":                     acctest.Representation{RepType: acctest.Optional, Create: `MONDAY`, Update: `TUESDAY`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                          acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesNewsReportRepresentation},
	}

	OpsiNewsReportSqlInsightsPlanChangesContentTypesRepresentation = map[string]interface{}{
		"sql_insights_plan_changes_resources": acctest.Representation{RepType: acctest.Required, Create: []string{`DATABASE`}, Update: []string{`DATABASE`, `EXADATA`}},
	}

	OpsiNewsReportSqlInsightsTopDatabasesRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"content_types":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiNewsReportSqlInsightsTopDatabasesContentTypesRepresentation},
		"description":                     acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_DESCRIPTION`, Update: `TF_TEST_REPORT_DESCRIPTION_2`},
		"locale":                          acctest.Representation{RepType: acctest.Required, Create: `EN`},
		"name":                            acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_NAME`, Update: `TF_TEST_REPORT_NAME_2`},
		"news_frequency":                  acctest.Representation{RepType: acctest.Required, Create: `WEEKLY`},
		"ons_topic_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.topic_id}`},
		"are_child_compartments_included": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"day_of_week":                     acctest.Representation{RepType: acctest.Optional, Create: `MONDAY`, Update: `TUESDAY`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                          acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesNewsReportRepresentation},
	}

	OpsiNewsReportSqlInsightsTopDatabasesContentTypesRepresentation = map[string]interface{}{
		"sql_insights_top_databases_resources": acctest.Representation{RepType: acctest.Required, Create: []string{`DATABASE`}, Update: []string{`DATABASE`, `EXADATA`}},
	}

	OpsiNewsReportSqlInsightsTopSqlByInsightsRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"content_types":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiNewsReportSqlInsightsTopSqlByInsightsContentTypesRepresentation},
		"description":                     acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_DESCRIPTION`, Update: `TF_TEST_REPORT_DESCRIPTION_2`},
		"locale":                          acctest.Representation{RepType: acctest.Required, Create: `EN`},
		"name":                            acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_NAME`, Update: `TF_TEST_REPORT_NAME_2`},
		"news_frequency":                  acctest.Representation{RepType: acctest.Required, Create: `WEEKLY`},
		"ons_topic_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.topic_id}`},
		"are_child_compartments_included": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"day_of_week":                     acctest.Representation{RepType: acctest.Optional, Create: `MONDAY`, Update: `TUESDAY`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                          acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesNewsReportRepresentation},
	}

	OpsiNewsReportSqlInsightsTopSqlByInsightsContentTypesRepresentation = map[string]interface{}{
		"sql_insights_top_sql_by_insights_resources": acctest.Representation{RepType: acctest.Required, Create: []string{`DATABASE`}, Update: []string{`DATABASE`, `EXADATA`}},
	}

	OpsiNewsReportSqlInsightsTopSqlRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"content_types":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiNewsReportSqlInsightsTopSqlContentTypesRepresentation},
		"description":                     acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_DESCRIPTION`, Update: `TF_TEST_REPORT_DESCRIPTION_2`},
		"locale":                          acctest.Representation{RepType: acctest.Required, Create: `EN`},
		"name":                            acctest.Representation{RepType: acctest.Required, Create: `TF_TEST_REPORT_NAME`, Update: `TF_TEST_REPORT_NAME_2`},
		"news_frequency":                  acctest.Representation{RepType: acctest.Required, Create: `WEEKLY`},
		"ons_topic_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.topic_id}`},
		"are_child_compartments_included": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"day_of_week":                     acctest.Representation{RepType: acctest.Optional, Create: `MONDAY`, Update: `TUESDAY`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"status":                          acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesNewsReportRepresentation},
	}

	OpsiNewsReportSqlInsightsTopSqlContentTypesRepresentation = map[string]interface{}{
		"sql_insights_top_sql_resources": acctest.Representation{RepType: acctest.Required, Create: []string{`DATABASE`}, Update: []string{`DATABASE`, `EXADATA`}},
	}

	ignoreChangesNewsReportRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	OpsiNewsReportResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiNewsReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiNewsReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	onsTopicId := utils.GetEnvSettingWithBlankDefault("topic_id")
	if onsTopicId == "" {
		t.Skip("Provision topic and set topic id to run this test")
	}
	topicIdVariableStr := fmt.Sprintf("variable \"topic_id\" { default = \"%s\" }\n", onsTopicId)

	resourceName := "oci_opsi_news_report.test_news_report"
	datasourceName := "data.oci_opsi_news_reports.test_news_reports"
	singularDatasourceName := "data.oci_opsi_news_report.test_news_report"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+topicIdVariableStr+OpsiNewsReportResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Optional, acctest.Create, OpsiNewsReportRepresentation), "opsi", "newsReport", t)

	acctest.ResourceTest(t, testAccCheckOpsiNewsReportDestroy, []resource.TestStep{
		//Step - Verify Create with Required
		{
			Config: config + compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Required, acctest.Create, OpsiNewsReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_types.0.capacity_planning_resources.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "description", "TF_TEST_REPORT_DESCRIPTION"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_TEST_REPORT_NAME"),
				resource.TestCheckResourceAttr(resourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				//resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpsiNewsReportResourceDependencies,
		},

		//Step - Verify Create with Required for SqlInsightsFleetAnalysis
		{
			Config: config + compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Required, acctest.Create, OpsiNewsReportSqlInsightsFleetAnalysisRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_types.0.sql_insights_fleet_analysis_resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "TF_TEST_REPORT_DESCRIPTION"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_TEST_REPORT_NAME"),
				resource.TestCheckResourceAttr(resourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				//resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpsiNewsReportResourceDependencies,
		},

		//Step - Verify Create with Required for SqlInsightsPerformanceDegradation
		{
			Config: config + compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Required, acctest.Create, OpsiNewsReportSqlInsightsPerformanceDegradationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_types.0.sql_insights_performance_degradation_resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "TF_TEST_REPORT_DESCRIPTION"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_TEST_REPORT_NAME"),
				resource.TestCheckResourceAttr(resourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				//resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpsiNewsReportResourceDependencies,
		},

		//Step - Verify Create with Required for SqlInsightsPlanChanges
		{
			Config: config + compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Required, acctest.Create, OpsiNewsReportSqlInsightsPlanChangesRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_types.0.sql_insights_plan_changes_resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "TF_TEST_REPORT_DESCRIPTION"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_TEST_REPORT_NAME"),
				resource.TestCheckResourceAttr(resourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				//resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpsiNewsReportResourceDependencies,
		},

		//Step - Verify Create with Required for SqlInsightsTopDatabases
		{
			Config: config + compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Required, acctest.Create, OpsiNewsReportSqlInsightsTopDatabasesRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_types.0.sql_insights_top_databases_resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "TF_TEST_REPORT_DESCRIPTION"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_TEST_REPORT_NAME"),
				resource.TestCheckResourceAttr(resourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				//resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpsiNewsReportResourceDependencies,
		},

		//Step - Verify Create with Required for SqlInsightsTopSqlByInsights
		{
			Config: config + compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Required, acctest.Create, OpsiNewsReportSqlInsightsTopSqlByInsightsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_types.0.sql_insights_top_sql_by_insights_resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "TF_TEST_REPORT_DESCRIPTION"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_TEST_REPORT_NAME"),
				resource.TestCheckResourceAttr(resourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				//resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpsiNewsReportResourceDependencies,
		},

		//Step - Verify Create with Required for SqlInsightsTopSql
		{
			Config: config + compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Required, acctest.Create, OpsiNewsReportSqlInsightsTopSqlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_types.0.sql_insights_top_sql_resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "TF_TEST_REPORT_DESCRIPTION"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_TEST_REPORT_NAME"),
				resource.TestCheckResourceAttr(resourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				//resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpsiNewsReportResourceDependencies,
		},

		//Step  - Verify Create with Optionals
		{
			Config: config + compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Optional, acctest.Create, OpsiNewsReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_child_compartments_included", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_types.0.capacity_planning_resources.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "day_of_week", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "description", "TF_TEST_REPORT_DESCRIPTION"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_TEST_REPORT_NAME"),
				resource.TestCheckResourceAttr(resourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

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

		//Step 1 - verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + topicIdVariableStr + compartmentIdUVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OpsiNewsReportRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_child_compartments_included", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_types.0.capacity_planning_resources.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "day_of_week", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "description", "TF_TEST_REPORT_DESCRIPTION"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_TEST_REPORT_NAME"),
				resource.TestCheckResourceAttr(resourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		//Step 2 - verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Optional, acctest.Update, OpsiNewsReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_child_compartments_included", "true"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "content_types.0.capacity_planning_resources.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "day_of_week", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "description", "TF_TEST_REPORT_DESCRIPTION_2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(resourceName, "name", "TF_TEST_REPORT_NAME_2"),
				resource.TestCheckResourceAttr(resourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//Step 3 - verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_news_reports", "test_news_reports", acctest.Optional, acctest.Update, OpsiNewsReportDataSourceRepresentation) +
				compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Optional, acctest.Update, OpsiNewsReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "news_report_id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "status.#", "1"), //status is not a list

				resource.TestCheckResourceAttr(datasourceName, "news_report_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "news_report_collection.0.items.#", "1"),
			),
		},
		//Step 4 - verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_news_report", "test_news_report", acctest.Required, acctest.Create, OpsiNewsReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + topicIdVariableStr + OpsiNewsReportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "news_report_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "are_child_compartments_included", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_types.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_types.0.capacity_planning_resources.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "day_of_week", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "TF_TEST_REPORT_DESCRIPTION_2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locale", "EN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TF_TEST_REPORT_NAME_2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "news_frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		//Step 5 - verify resource import
		{
			Config:                  config + OpsiNewsReportRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOpsiNewsReportDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_news_report" {
			noResourceFound = false
			request := oci_opsi.GetNewsReportRequest{}

			tmp := rs.Primary.ID
			request.NewsReportId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetNewsReport(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpsiNewsReport") {
		resource.AddTestSweepers("OpsiNewsReport", &resource.Sweeper{
			Name:         "OpsiNewsReport",
			Dependencies: acctest.DependencyGraph["newsReport"],
			F:            sweepOpsiNewsReportResource,
		})
	}
}

func sweepOpsiNewsReportResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	newsReportIds, err := getOpsiNewsReportIds(compartment)
	if err != nil {
		return err
	}
	for _, newsReportId := range newsReportIds {
		if ok := acctest.SweeperDefaultResourceId[newsReportId]; !ok {
			deleteNewsReportRequest := oci_opsi.DeleteNewsReportRequest{}

			deleteNewsReportRequest.NewsReportId = &newsReportId

			deleteNewsReportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteNewsReport(context.Background(), deleteNewsReportRequest)
			if error != nil {
				fmt.Printf("Error deleting NewsReport %s %s, It is possible that the resource is already deleted. Please verify manually \n", newsReportId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &newsReportId, OpsiNewsReportSweepWaitCondition, time.Duration(3*time.Minute),
				OpsiNewsReportSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getOpsiNewsReportIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NewsReportId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listNewsReportsRequest := oci_opsi.ListNewsReportsRequest{}
	listNewsReportsRequest.CompartmentId = &compartmentId
	//listNewsReportsRequest.LifecycleState = oci_opsi.ListNewsReportsLifecycleStateActiveNeedsAttention
	listNewsReportsResponse, err := operationsInsightsClient.ListNewsReports(context.Background(), listNewsReportsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NewsReport list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, newsReport := range listNewsReportsResponse.Items {
		id := *newsReport.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NewsReportId", id)
	}
	return resourceIds, nil
}

func OpsiNewsReportSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if newsReportResponse, ok := response.Response.(oci_opsi.GetNewsReportResponse); ok {
		return newsReportResponse.LifecycleState != oci_opsi.LifecycleStateDeleted
	}
	return false
}

func OpsiNewsReportSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetNewsReport(context.Background(), oci_opsi.GetNewsReportRequest{
		NewsReportId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
