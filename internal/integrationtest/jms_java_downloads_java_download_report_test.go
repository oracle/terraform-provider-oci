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
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	JmsJavaDownloadsJavaDownloadReportResourceDependencies = DefinedTagsDependencies
	JmsJdReportTimeStartedTime                             = time.Now().AddDate(0, -1, 0).UTC()
	JmsJdReportTimeStarted                                 = JmsJdReportTimeStartedTime.Format(time.RFC3339)

	JmsJdReportTimeEndedTime = time.Now().UTC()
	JmsJdReportTimeEnded     = JmsJdReportTimeEndedTime.Format(time.RFC3339)

	JmsJavaDownloadsJavaDownloadReportSingularDataSourceRepresentation = map[string]interface{}{
		"java_download_report_id": acctest.Representation{
			RepType: acctest.Required,
			Create:  `${oci_jms_java_downloads_java_download_report.test_java_download_report.id}`,
		},
	}

	JmsJavaDownloadsJavaDownloadReportDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"java_download_report_id": acctest.Representation{
			RepType: acctest.Optional,
			Create:  `${oci_jms_java_downloads_java_download_report.test_java_download_report.id}`},
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: JmsJavaDownloadsJavaDownloadReportDataSourceFilterRepresentation}}

	JmsJavaDownloadsJavaDownloadReportDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_jms_java_downloads_java_download_report.test_java_download_report.id}`}},
	}

	JmsJavaDownloadsJavaDownloadReportRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"format":         acctest.Representation{RepType: acctest.Required, Create: `CSV`},
		"defined_tags": acctest.Representation{
			RepType: acctest.Optional,
			Create:  `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`,
			Update:  `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{
			RepType: acctest.Optional,
			Create:  map[string]string{"bar-key": "bar-value"},
			Update:  map[string]string{"bar-key": "updatedValue"}},
		"time_end":   acctest.Representation{RepType: acctest.Optional, Create: JmsJdReportTimeEnded},
		"time_start": acctest.Representation{RepType: acctest.Optional, Create: JmsJdReportTimeStarted},
		"lifecycle": acctest.RepresentationGroup{
			RepType: acctest.Required,
			Group: map[string]interface{}{
				"ignore_changes": acctest.Representation{
					RepType: acctest.Required,
					Create:  []string{`defined_tags`, `system_tags`}},
			},
		},
	}
)

// issue-routing-tag: jms_java_downloads/default
func TestJmsJavaDownloadsJavaDownloadReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsJavaDownloadsJavaDownloadReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_jms_java_downloads_java_download_report.test_java_download_report"
	datasourceName := "data.oci_jms_java_downloads_java_download_reports.test_java_download_reports"
	singularDatasourceName := "data.oci_jms_java_downloads_java_download_report.test_java_download_report"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+
		JmsJavaDownloadsJavaDownloadReportResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap(
			"oci_jms_java_downloads_java_download_report",
			"test_java_download_report",
			acctest.Optional,
			acctest.Create,
			JmsJavaDownloadsJavaDownloadReportRepresentation),
		"jmsjavadownloads",
		"javaDownloadReport",
		t,
	)

	acctest.ResourceTest(t, testAccCheckJmsJavaDownloadsJavaDownloadReportDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config +
				JmsJavaDownloadsJavaDownloadReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_report",
					"test_java_download_report",
					acctest.Optional,
					acctest.Create,
					JmsJavaDownloadsJavaDownloadReportRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "checksum_type"),
				resource.TestCheckResourceAttrSet(resourceName, "checksum_value"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "created_by.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "file_size_in_bytes"),
				resource.TestCheckResourceAttr(resourceName, "format", "CSV"),
				resource.TestCheckResourceAttrSet(resourceName, "sort_by"),
				resource.TestCheckResourceAttrSet(resourceName, "sort_order"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_end", JmsJdReportTimeEnded),
				resource.TestCheckResourceAttr(resourceName, "time_start", JmsJdReportTimeStarted),
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
				JmsJavaDownloadsJavaDownloadReportResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_reports",
					"test_java_download_reports",
					acctest.Optional,
					acctest.Update,
					JmsJavaDownloadsJavaDownloadReportDataSourceRepresentation,
				) +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_report",
					"test_java_download_report",
					acctest.Optional,
					acctest.Create,
					JmsJavaDownloadsJavaDownloadReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(datasourceName, "java_download_report_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "java_download_report_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "java_download_report_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				JmsJavaDownloadsJavaDownloadReportResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_report",
					"test_java_download_report",
					acctest.Required,
					acctest.Create,
					JmsJavaDownloadsJavaDownloadReportSingularDataSourceRepresentation,
				) +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_report",
					"test_java_download_report",
					acctest.Optional,
					acctest.Create,
					JmsJavaDownloadsJavaDownloadReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "java_download_report_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "checksum_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "checksum_value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "created_by.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "file_size_in_bytes"),
				resource.TestCheckResourceAttr(singularDatasourceName, "format", "CSV"),
				resource.TestCheckResourceAttrSet(resourceName, "sort_by"),
				resource.TestCheckResourceAttrSet(resourceName, "sort_order"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_start"),
			),
		},
	})
}

func testAccCheckJmsJavaDownloadsJavaDownloadReportDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).JavaDownloadClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_jms_java_downloads_java_download_report" {
			noResourceFound = false
			request := oci_jms_java_downloads.GetJavaDownloadReportRequest{}

			tmp := rs.Primary.ID
			request.JavaDownloadReportId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "jms_java_downloads")

			response, err := client.GetJavaDownloadReport(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_jms_java_downloads.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("JmsJavaDownloadsJavaDownloadReport") {
		resource.AddTestSweepers("JmsJavaDownloadsJavaDownloadReport", &resource.Sweeper{
			Name:         "JmsJavaDownloadsJavaDownloadReport",
			Dependencies: acctest.DependencyGraph["javaDownloadReport"],
			F:            sweepJmsJavaDownloadsJavaDownloadReportResource,
		})
	}
}

func sweepJmsJavaDownloadsJavaDownloadReportResource(compartment string) error {
	javaDownloadClient := acctest.GetTestClients(&schema.ResourceData{}).JavaDownloadClient()
	// JmsJavaDownloadsJavaDownloadReportResource can only run on root compartment
	compartment = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	javaDownloadReportIds, err := getJmsJavaDownloadsJavaDownloadReportIds(compartment)
	if err != nil {
		return err
	}
	for _, javaDownloadReportId := range javaDownloadReportIds {
		if ok := acctest.SweeperDefaultResourceId[javaDownloadReportId]; !ok {
			deleteJavaDownloadReportRequest := oci_jms_java_downloads.DeleteJavaDownloadReportRequest{}

			deleteJavaDownloadReportRequest.JavaDownloadReportId = &javaDownloadReportId

			deleteJavaDownloadReportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "jms_java_downloads")
			_, error := javaDownloadClient.DeleteJavaDownloadReport(context.Background(), deleteJavaDownloadReportRequest)
			if error != nil {
				fmt.Printf("Error deleting JavaDownloadReport %s %s, It is possible that the resource is already deleted. Please verify manually \n", javaDownloadReportId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &javaDownloadReportId, JmsJavaDownloadsJavaDownloadReportSweepWaitCondition, time.Duration(3*time.Minute),
				JmsJavaDownloadsJavaDownloadReportSweepResponseFetchOperation, "jms_java_downloads", true)
		}
	}
	return nil
}

func getJmsJavaDownloadsJavaDownloadReportIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "JavaDownloadReportId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	javaDownloadClient := acctest.GetTestClients(&schema.ResourceData{}).JavaDownloadClient()

	listJavaDownloadReportsRequest := oci_jms_java_downloads.ListJavaDownloadReportsRequest{}
	listJavaDownloadReportsRequest.CompartmentId = &compartmentId
	listJavaDownloadReportsRequest.LifecycleState = oci_jms_java_downloads.ListJavaDownloadReportsLifecycleStateNeedsAttention
	listJavaDownloadReportsResponse, err := javaDownloadClient.ListJavaDownloadReports(context.Background(), listJavaDownloadReportsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting JavaDownloadReport list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, javaDownloadReport := range listJavaDownloadReportsResponse.Items {
		id := *javaDownloadReport.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "JavaDownloadReportId", id)
	}
	return resourceIds, nil
}

func JmsJavaDownloadsJavaDownloadReportSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if javaDownloadReportResponse, ok := response.Response.(oci_jms_java_downloads.GetJavaDownloadReportResponse); ok {
		return javaDownloadReportResponse.LifecycleState != oci_jms_java_downloads.LifecycleStateDeleted
	}
	return false
}

func JmsJavaDownloadsJavaDownloadReportSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.JavaDownloadClient().GetJavaDownloadReport(context.Background(), oci_jms_java_downloads.GetJavaDownloadReportRequest{
		JavaDownloadReportId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
