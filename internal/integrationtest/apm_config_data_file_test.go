// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_apm_config "github.com/oracle/oci-go-sdk/v65/apmconfig"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApmConfigDataFileRequiredOnlyResource = ApmConfigDataFileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_data_file", "test_data_file", acctest.Required, acctest.Create, ApmConfigDataFileRepresentation)

	ApmConfigDataFileResourceConfig = ApmConfigDataFileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_data_file", "test_data_file", acctest.Optional, acctest.Update, ApmConfigDataFileRepresentation)

	ApmConfigDataFileSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"apm_type":       acctest.Representation{RepType: acctest.Required, Create: `apmType`},
		"data_file_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_config_data_file.test_data_file.data_file_name}`},
	}

	ApmConfigDataFileDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"apm_type":                  acctest.Representation{RepType: acctest.Optional, Create: `apmType`},
		"metadata":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"testvalue": "test1"}, Update: map[string]string{"testvalue": "test2"}},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `dataFileName`},
		"time_last_modified_after":  acctest.Representation{RepType: acctest.Optional, Create: `2006-01-02T15:04:05Z`},
		"time_last_modified_before": acctest.Representation{RepType: acctest.Optional, Create: `4001-01-02T15:04:05Z`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmConfigDataFileDataSourceFilterRepresentation}}
	ApmConfigDataFileDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apm_config_data_file.test_data_file.data_file_name}`}},
	}

	ApmConfigDataFileRepresentation = map[string]interface{}{
		"content":             acctest.Representation{RepType: acctest.Required, Create: `putDataFileBody`, Update: `putDataFileBody2`},
		"apm_domain_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"apm_type":            acctest.Representation{RepType: acctest.Required, Create: `apmType`},
		"data_file_name":      acctest.Representation{RepType: acctest.Required, Create: `dataFileName`},
		"metadata":            acctest.Representation{RepType: acctest.Required, Create: map[string]string{"testvalue": "test1"}, Update: map[string]string{"testvalue": "test2"}},
		"content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `contentDisposition`, Update: `contentDisposition2`},
		"content_encoding":    acctest.Representation{RepType: acctest.Optional, Create: `contentEncoding`, Update: `contentEncoding2`},
		"content_language":    acctest.Representation{RepType: acctest.Optional, Create: `contentLanguage`, Update: `contentLanguage2`},
		"content_md5":         acctest.Representation{RepType: acctest.Optional, Create: `A/74/26CEk1/SQAvsDf1qQ==`, Update: `NmHTrJ2cNntoNJwYQfoYqA==`},
		"content_type":        acctest.Representation{RepType: acctest.Optional, Create: `text/plain`, Update: `application/octet-stream`},
	}

	ApmConfigDataFileResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation)
)

// issue-routing-tag: apm_config/default
func TestApmConfigDataFileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmConfigDataFileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_config_data_file.test_data_file"
	datasourceName := "data.oci_apm_config_data_files.test_data_files"
	singularDatasourceName := "data.oci_apm_config_data_file.test_data_file"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmConfigDataFileResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_data_file", "test_data_file", acctest.Optional, acctest.Create, ApmConfigDataFileRepresentation), "apmconfig", "dataFile", t)

	acctest.ResourceTest(t, testAccCheckApmConfigDataFileDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ApmConfigDataFileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_config_data_file", "test_data_file", acctest.Required, acctest.Create, ApmConfigDataFileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "apm_type", "apmType"),
				resource.TestCheckResourceAttr(resourceName, "data_file_name", "dataFileName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ApmConfigDataFileResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmConfigDataFileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_config_data_file", "test_data_file", acctest.Optional, acctest.Create, ApmConfigDataFileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content_disposition", "contentDisposition"),
				resource.TestCheckResourceAttr(resourceName, "content_encoding", "contentEncoding"),
				resource.TestCheckResourceAttr(resourceName, "content_language", "contentLanguage"),
				resource.TestCheckResourceAttr(resourceName, "content_md5", "A/74/26CEk1/SQAvsDf1qQ=="),
				resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "apm_type", "apmType"),
				resource.TestCheckResourceAttrSet(resourceName, "data_file_name"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.testvalue", "test1"),

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
			Config: config + compartmentIdVariableStr + ApmConfigDataFileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_config_data_file", "test_data_file", acctest.Optional, acctest.Update, ApmConfigDataFileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content_disposition", "contentDisposition2"),
				resource.TestCheckResourceAttr(resourceName, "content_encoding", "contentEncoding2"),
				resource.TestCheckResourceAttr(resourceName, "content_language", "contentLanguage2"),
				resource.TestCheckResourceAttr(resourceName, "content_md5", "NmHTrJ2cNntoNJwYQfoYqA=="),
				resource.TestCheckResourceAttr(resourceName, "content_type", "application/octet-stream"),
				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "apm_type", "apmType"),
				resource.TestCheckResourceAttrSet(resourceName, "data_file_name"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.testvalue", "test2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_data_files", "test_data_files", acctest.Optional, acctest.Update, ApmConfigDataFileDataSourceRepresentation) +
				compartmentIdVariableStr + ApmConfigDataFileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_config_data_file", "test_data_file", acctest.Optional, acctest.Update, ApmConfigDataFileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "apm_type", "apmType"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.testvalue", "test2"),
				resource.TestCheckResourceAttr(datasourceName, "name", "dataFileName"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_last_modified_after"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_last_modified_before"),

				resource.TestCheckResourceAttr(datasourceName, "items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_data_file", "test_data_file", acctest.Required, acctest.Create, ApmConfigDataFileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmConfigDataFileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "apm_type", "apmType"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_file_name"),
			),
		},
		// verify resource import
		{
			Config:            config + ApmConfigDataFileRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"content_disposition",
				"content_encoding",
				"content_language",
				"content_md5",
				"content_type",
				"content",
				"apm_domain_id",
				"apm_type",
				"metadata",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckApmConfigDataFileDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ConfigClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_config_data_file" {
			noResourceFound = false
			request := oci_apm_config.GetDataFileRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			if value, ok := rs.Primary.Attributes["apm_type"]; ok {
				request.ApmType = &value
			}

			if value, ok := rs.Primary.Attributes["data_file_name"]; ok {
				request.DataFileName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_config")

			_, err := client.GetDataFile(context.Background(), request)

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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ApmConfigDataFile") {
		resource.AddTestSweepers("ApmConfigDataFile", &resource.Sweeper{
			Name:         "ApmConfigDataFile",
			Dependencies: acctest.DependencyGraph["dataFile"],
			F:            sweepApmConfigDataFileResource,
		})
	}
}

func sweepApmConfigDataFileResource(compartment string) error {
	configClient := acctest.GetTestClients(&schema.ResourceData{}).ConfigClient()
	dataFileIds, err := getApmConfigDataFileIds(compartment)
	if err != nil {
		return err
	}
	for _, dataFileId := range dataFileIds {
		if ok := acctest.SweeperDefaultResourceId[dataFileId]; !ok {
			deleteDataFileRequest := oci_apm_config.DeleteDataFileRequest{}

			deleteDataFileRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_config")
			_, error := configClient.DeleteDataFile(context.Background(), deleteDataFileRequest)
			if error != nil {
				fmt.Printf("Error deleting DataFile %s %s, It is possible that the resource is already deleted. Please verify manually \n", dataFileId, error)
				continue
			}
		}
	}
	return nil
}

func getApmConfigDataFileIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DataFileId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	configClient := acctest.GetTestClients(&schema.ResourceData{}).ConfigClient()

	listDataFilesRequest := oci_apm_config.ListDataFilesRequest{}

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for DataFile resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listDataFilesRequest.ApmDomainId = &apmDomainId

		listDataFilesResponse, err := configClient.ListDataFiles(context.Background(), listDataFilesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DataFile list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, dataFile := range listDataFilesResponse.Items {
			id := GetDataFileCompositeId(apmDomainId, *dataFile.ApmType, *dataFile.Name)
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DataFileId", id)
		}

	}
	return resourceIds, nil
}

func GetDataFileCompositeId(apmDomainId string, apmType string, dataFileName string) string {
	apmDomainId = url.PathEscape(apmDomainId)
	apmType = url.PathEscape(apmType)
	dataFileName = url.PathEscape(dataFileName)
	compositeId := "dataFiles/" + dataFileName + "/apmDomainId/" + apmDomainId + "/apmType/" + apmType
	return compositeId
}
