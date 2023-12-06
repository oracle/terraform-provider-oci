// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	JmsJdUserOcid                                   = utils.GetEnvSettingWithBlankDefault("user_ocid")
	jmsJdTokenName                                  = fmt.Sprintf("Terraform_Token_%d", time.Now().UnixMilli())
	jmsJdTokenNameUpdate                            = fmt.Sprintf("Terraform_Token_%d", time.Now().AddDate(0, 0, 1).UnixMilli())
	jmsJdTokenCreateExpiry                          = time.Now().AddDate(0, 0, 2).UTC().Format(time.RFC3339)
	jmsJdTokenUpdateExpiry                          = time.Now().AddDate(0, 0, 4).UTC().Format(time.RFC3339)
	JmsJavaDownloadsJavaDownloadTokenResourceConfig = acctest.GenerateResourceFromRepresentationMap(
		"oci_jms_java_downloads_java_download_token",
		"test_java_download_token",
		acctest.Optional,
		acctest.Update,
		JmsJavaDownloadsJavaDownloadTokenRepresentation)

	JmsJavaDownloadsJavaDownloadTokenSingularDataSourceRepresentation = map[string]interface{}{
		"java_download_token_id": acctest.Representation{
			RepType: acctest.Required,
			Create:  `${oci_jms_java_downloads_java_download_token.test_java_download_token.id}`},
	}

	JmsJavaDownloadsJavaDownloadTokenDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: jmsJdTokenName, Update: jmsJdTokenNameUpdate},
		"family_version": acctest.Representation{RepType: acctest.Optional, Create: `11`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_jms_java_downloads_java_download_token.test_java_download_token.id}`},
		"search_by_user": acctest.Representation{RepType: acctest.Optional, Create: JmsJdUserOcid},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"value":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_jms_java_downloads_java_download_token.test_java_download_token.value}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: JmsJavaDownloadsJavaDownloadTokenDataSourceFilterRepresentation}}
	JmsJavaDownloadsJavaDownloadTokenDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_jms_java_downloads_java_download_token.test_java_download_token.id}`}},
	}

	JmsJavaDownloadsJavaDownloadTokenRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `Trial token for script friendly download`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: jmsJdTokenName, Update: jmsJdTokenNameUpdate},
		"java_version":   acctest.Representation{RepType: acctest.Required, Create: `11`},
		"license_type":   acctest.Representation{RepType: acctest.Required, Create: []string{`OTN`}, Update: []string{`OTN`}},
		"time_expires":   acctest.Representation{RepType: acctest.Required, Create: jmsJdTokenCreateExpiry, Update: jmsJdTokenUpdateExpiry},
		"lifecycle": acctest.RepresentationGroup{
			RepType: acctest.Required,
			Group: map[string]interface{}{
				"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
			},
		},
	}
	jmsJdTokenNameOptional                                  = "Optional" + jmsJdTokenName
	JmsJavaDownloadsJavaDownloadTokenOptionalRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `Trial token for script friendly download`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: jmsJdTokenNameOptional, Update: jmsJdTokenNameUpdate},
		"java_version":   acctest.Representation{RepType: acctest.Required, Create: `11`},
		"license_type":   acctest.Representation{RepType: acctest.Required, Create: []string{`OTN`}, Update: []string{`OTN`}},
		"time_expires":   acctest.Representation{RepType: acctest.Required, Create: jmsJdTokenCreateExpiry, Update: jmsJdTokenUpdateExpiry},
		"lifecycle": acctest.RepresentationGroup{
			RepType: acctest.Required,
			Group: map[string]interface{}{
				"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
			},
		},

		"is_default": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
)

// issue-routing-tag: jms_java_downloads/default
func TestJmsJavaDownloadsJavaDownloadTokenResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsJavaDownloadsJavaDownloadTokenResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_jms_java_downloads_java_download_token.test_java_download_token"
	datasourceName := "data.oci_jms_java_downloads_java_download_tokens.test_java_download_tokens"
	singularDatasourceName := "data.oci_jms_java_downloads_java_download_token.test_java_download_token"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+
		acctest.GenerateResourceFromRepresentationMap(
			"oci_jms_java_downloads_java_download_token",
			"test_java_download_token",
			acctest.Optional,
			acctest.Create,
			JmsJavaDownloadsJavaDownloadTokenRepresentation),

		"jmsjavadownloads",
		"javaDownloadToken",
		t,
	)

	acctest.ResourceTest(t, testAccCheckJmsJavaDownloadsJavaDownloadTokenDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_token",
					"test_java_download_token",
					acctest.Required,
					acctest.Create,
					JmsJavaDownloadsJavaDownloadTokenRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Trial token for script friendly download"),
				resource.TestCheckResourceAttr(resourceName, "display_name", jmsJdTokenName),
				resource.TestCheckResourceAttr(resourceName, "java_version", "11"),
				resource.TestCheckResourceAttr(resourceName, "license_type.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", jmsJdTokenCreateExpiry),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config,
		},
		// verify Create with optionals
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_token",
					"test_java_download_token",
					acctest.Optional,
					acctest.Create,
					JmsJavaDownloadsJavaDownloadTokenOptionalRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "created_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "Trial token for script friendly download"),
				resource.TestCheckResourceAttr(resourceName, "display_name", jmsJdTokenNameOptional),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				resource.TestCheckResourceAttr(resourceName, "java_version", "11"),
				resource.TestCheckResourceAttr(resourceName, "license_type.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", jmsJdTokenCreateExpiry),
				resource.TestCheckResourceAttrSet(resourceName, "value"),

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
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_token",
					"test_java_download_token",
					acctest.Optional,
					acctest.Update,
					JmsJavaDownloadsJavaDownloadTokenOptionalRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "created_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", jmsJdTokenNameUpdate),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				resource.TestCheckResourceAttr(resourceName, "java_version", "11"),
				resource.TestCheckResourceAttr(resourceName, "license_type.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", jmsJdTokenUpdateExpiry),
				resource.TestCheckResourceAttrSet(resourceName, "value"),

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
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_tokens",
					"test_java_download_tokens",
					acctest.Optional,
					acctest.Update,
					JmsJavaDownloadsJavaDownloadTokenDataSourceRepresentation,
				) +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_token",
					"test_java_download_token",
					acctest.Optional,
					acctest.Update,
					JmsJavaDownloadsJavaDownloadTokenOptionalRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", jmsJdTokenNameUpdate),
				resource.TestCheckResourceAttr(datasourceName, "family_version", "11"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "search_by_user", JmsJdUserOcid),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "value"),

				resource.TestCheckResourceAttr(datasourceName, "java_download_token_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "java_download_token_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_java_downloads_java_download_token",
					"test_java_download_token",
					acctest.Required,
					acctest.Create,
					JmsJavaDownloadsJavaDownloadTokenSingularDataSourceRepresentation,
				) +
				JmsJavaDownloadsJavaDownloadTokenResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "java_download_token_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", jmsJdTokenNameUpdate),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_default", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "java_version", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "last_updated_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_type.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_expires"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_used"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "value"),
			),
		},
	})
}

func testAccCheckJmsJavaDownloadsJavaDownloadTokenDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).JavaDownloadClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_jms_java_downloads_java_download_token" {
			noResourceFound = false
			request := oci_jms_java_downloads.GetJavaDownloadTokenRequest{}

			tmp := rs.Primary.ID
			request.JavaDownloadTokenId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "jms_java_downloads")

			response, err := client.GetJavaDownloadToken(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("JmsJavaDownloadsJavaDownloadToken") {
		resource.AddTestSweepers("JmsJavaDownloadsJavaDownloadToken", &resource.Sweeper{
			Name:         "JmsJavaDownloadsJavaDownloadToken",
			Dependencies: acctest.DependencyGraph["javaDownloadToken"],
			F:            sweepJmsJavaDownloadsJavaDownloadTokenResource,
		})
	}
}

func sweepJmsJavaDownloadsJavaDownloadTokenResource(compartment string) error {
	javaDownloadClient := acctest.GetTestClients(&schema.ResourceData{}).JavaDownloadClient()
	// JmsJavaDownloadsJavaDownloadTokenResource can only run on root compartment
	compartment = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	javaDownloadTokenIds, err := getJmsJavaDownloadsJavaDownloadTokenIds(compartment)
	if err != nil {
		return err
	}
	for _, javaDownloadTokenId := range javaDownloadTokenIds {
		if ok := acctest.SweeperDefaultResourceId[javaDownloadTokenId]; !ok {
			deleteJavaDownloadTokenRequest := oci_jms_java_downloads.DeleteJavaDownloadTokenRequest{}

			deleteJavaDownloadTokenRequest.JavaDownloadTokenId = &javaDownloadTokenId

			deleteJavaDownloadTokenRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "jms_java_downloads")
			_, error := javaDownloadClient.DeleteJavaDownloadToken(context.Background(), deleteJavaDownloadTokenRequest)
			if error != nil {
				fmt.Printf("Error deleting JavaDownloadToken %s %s, It is possible that the resource is already deleted. Please verify manually \n", javaDownloadTokenId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &javaDownloadTokenId, JmsJavaDownloadsJavaDownloadTokenSweepWaitCondition, time.Duration(3*time.Minute),
				JmsJavaDownloadsJavaDownloadTokenSweepResponseFetchOperation, "jms_java_downloads", true)
		}
	}
	return nil
}

func getJmsJavaDownloadsJavaDownloadTokenIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "JavaDownloadTokenId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	javaDownloadClient := acctest.GetTestClients(&schema.ResourceData{}).JavaDownloadClient()

	listJavaDownloadTokensRequest := oci_jms_java_downloads.ListJavaDownloadTokensRequest{}
	listJavaDownloadTokensRequest.CompartmentId = &compartmentId
	listJavaDownloadTokensRequest.LifecycleState = oci_jms_java_downloads.ListJavaDownloadTokensLifecycleStateActive
	listJavaDownloadTokensResponse, err := javaDownloadClient.ListJavaDownloadTokens(context.Background(), listJavaDownloadTokensRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting JavaDownloadToken list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, javaDownloadToken := range listJavaDownloadTokensResponse.Items {
		id := *javaDownloadToken.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "JavaDownloadTokenId", id)
	}
	return resourceIds, nil
}

func JmsJavaDownloadsJavaDownloadTokenSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if javaDownloadTokenResponse, ok := response.Response.(oci_jms_java_downloads.GetJavaDownloadTokenResponse); ok {
		return javaDownloadTokenResponse.LifecycleState != oci_jms_java_downloads.LifecycleStateDeleted
	}
	return false
}

func JmsJavaDownloadsJavaDownloadTokenSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.JavaDownloadClient().GetJavaDownloadToken(context.Background(), oci_jms_java_downloads.GetJavaDownloadTokenRequest{
		JavaDownloadTokenId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
