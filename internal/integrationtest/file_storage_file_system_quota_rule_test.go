// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FileStorageFileSystemQuotaRuleRequiredOnlyResource = FileStorageFileSystemQuotaRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system_quota_rule", "test_file_system_quota_rule", acctest.Required, acctest.Create, FileStorageFileSystemQuotaRuleRepresentation)

	FileStorageFileSystemQuotaRuleResourceConfig = FileStorageFileSystemQuotaRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system_quota_rule", "test_file_system_quota_rule", acctest.Optional, acctest.Update, FileStorageFileSystemQuotaRuleRepresentation)

	FileStorageFileSystemQuotaRuleSingularDataSourceRepresentation = map[string]interface{}{
		"file_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"quota_rule_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system_quota_rule.test_file_system_quota_rule.quota_rule_id}`},
	}

	FileStorageFileSystemQuotaRuleDataSourceRepresentation = map[string]interface{}{
		"file_system_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"principal_type":     acctest.Representation{RepType: acctest.Required, Create: `FILE_SYSTEM_LEVEL`},
		"are_violators_only": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: FileStorageFileSystemQuotaRuleDataSourceFilterRepresentation},
	}
	FileStorageFileSystemQuotaRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_file_storage_file_system_quota_rule.test_file_system_quota_rule.quota_rule_id}`}},
	}

	FileStorageFileSystemQuotaRuleRepresentation = map[string]interface{}{
		"file_system_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_file_system.test_file_system.id}`},
		"is_hard_quota":            acctest.Representation{RepType: acctest.Required, Create: `false`},
		"principal_type":           acctest.Representation{RepType: acctest.Required, Create: `FILE_SYSTEM_LEVEL`},
		"quota_limit_in_gigabytes": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}

	FileStorageFileSystemQuotaRuleResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Required, acctest.Create, FileStorageFileSystemRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: file_storage/default
func TestFileStorageFileSystemQuotaRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageFileSystemQuotaRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_file_system_quota_rule.test_file_system_quota_rule"
	datasourceName := "data.oci_file_storage_file_system_quota_rules.test_file_system_quota_rules"
	singularDatasourceName := "data.oci_file_storage_file_system_quota_rule.test_file_system_quota_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FileStorageFileSystemQuotaRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system_quota_rule", "test_file_system_quota_rule", acctest.Optional, acctest.Create, FileStorageFileSystemQuotaRuleRepresentation), "filestorage", "fileSystemQuotaRule", t)

	acctest.ResourceTest(t, testAccCheckFileStorageFileSystemQuotaRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FileStorageFileSystemQuotaRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system_quota_rule", "test_file_system_quota_rule", acctest.Required, acctest.Create, FileStorageFileSystemQuotaRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttr(resourceName, "is_hard_quota", "false"),
				resource.TestCheckResourceAttr(resourceName, "principal_type", "FILE_SYSTEM_LEVEL"),
				resource.TestCheckResourceAttr(resourceName, "quota_limit_in_gigabytes", "10"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FileStorageFileSystemQuotaRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FileStorageFileSystemQuotaRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system_quota_rule", "test_file_system_quota_rule", acctest.Optional, acctest.Create, FileStorageFileSystemQuotaRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_hard_quota", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "principal_id"),
				resource.TestCheckResourceAttr(resourceName, "principal_type", "FILE_SYSTEM_LEVEL"),
				resource.TestCheckResourceAttr(resourceName, "quota_limit_in_gigabytes", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + FileStorageFileSystemQuotaRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system_quota_rule", "test_file_system_quota_rule", acctest.Optional, acctest.Update, FileStorageFileSystemQuotaRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_hard_quota", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "principal_id"),
				resource.TestCheckResourceAttr(resourceName, "principal_type", "FILE_SYSTEM_LEVEL"),
				resource.TestCheckResourceAttr(resourceName, "quota_limit_in_gigabytes", "11"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_file_system_quota_rules", "test_file_system_quota_rules", acctest.Optional, acctest.Update, FileStorageFileSystemQuotaRuleDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageFileSystemQuotaRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system_quota_rule", "test_file_system_quota_rule", acctest.Optional, acctest.Update, FileStorageFileSystemQuotaRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "are_violators_only", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "file_system_id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "principal_id"),
				resource.TestCheckResourceAttr(datasourceName, "principal_type", "FILE_SYSTEM_LEVEL"),

				resource.TestCheckResourceAttr(datasourceName, "quota_rules.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "quota_rules.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "quota_rules.0.file_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "quota_rules.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "quota_rules.0.is_hard_quota", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "quota_rules.0.principal_id"),
				//resource.TestCheckResourceAttr(datasourceName, "quota_rules.0.principal_type", "FILE_SYSTEM_LEVEL"),
				resource.TestCheckResourceAttr(datasourceName, "quota_rules.0.quota_limit_in_gigabytes", "11"),
				resource.TestCheckResourceAttrSet(datasourceName, "quota_rules.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "quota_rules.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_file_system_quota_rule", "test_file_system_quota_rule", acctest.Required, acctest.Create, FileStorageFileSystemQuotaRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageFileSystemQuotaRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "file_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "quota_rule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_hard_quota", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "principal_type", "FILE_SYSTEM_LEVEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "quota_limit_in_gigabytes", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + FileStorageFileSystemQuotaRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFileStorageFileSystemQuotaRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_file_system_quota_rule" {
			noResourceFound = false
			request := oci_file_storage.GetQuotaRuleRequest{}

			if value, ok := rs.Primary.Attributes["file_system_id"]; ok {
				request.FileSystemId = &value
			}

			if value, ok := rs.Primary.Attributes["quota_rule_id"]; ok {
				request.QuotaRuleId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")

			_, err := client.GetQuotaRule(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("FileStorageFileSystemQuotaRule") {
		resource.AddTestSweepers("FileStorageFileSystemQuotaRule", &resource.Sweeper{
			Name:         "FileStorageFileSystemQuotaRule",
			Dependencies: acctest.DependencyGraph["fileSystemQuotaRule"],
			F:            sweepFileStorageFileSystemQuotaRuleResource,
		})
	}
}

func sweepFileStorageFileSystemQuotaRuleResource(compartment string) error {
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()
	fileSystemQuotaRuleIds, err := getFileStorageFileSystemQuotaRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, fileSystemQuotaRuleId := range fileSystemQuotaRuleIds {
		if ok := acctest.SweeperDefaultResourceId[fileSystemQuotaRuleId]; !ok {
			deleteQuotaRuleRequest := oci_file_storage.DeleteQuotaRuleRequest{}

			deleteQuotaRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")
			_, error := fileStorageClient.DeleteQuotaRule(context.Background(), deleteQuotaRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting FileSystemQuotaRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", fileSystemQuotaRuleId, error)
				continue
			}
		}
	}
	return nil
}

func getFileStorageFileSystemQuotaRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FileSystemQuotaRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()

	listQuotaRulesRequest := oci_file_storage.ListQuotaRulesRequest{}

	fileSystemIds, error := getFileStorageFileSystemIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting fileSystemId required for FileSystemQuotaRule resource requests \n")
	}
	for _, fileSystemId := range fileSystemIds {
		listQuotaRulesRequest.FileSystemId = &fileSystemId

		principalTypes := oci_file_storage.GetListQuotaRulesPrincipalTypeEnumValues()
		for _, principalType := range principalTypes {
			listQuotaRulesRequest.PrincipalType = principalType

			listQuotaRulesResponse, err := fileStorageClient.ListQuotaRules(context.Background(), listQuotaRulesRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting FileSystemQuotaRule list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, fileSystemQuotaRule := range listQuotaRulesResponse.Items {
				id := *fileSystemQuotaRule.Id
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FileSystemQuotaRuleId", id)
			}

		}
	}
	return resourceIds, nil
}
