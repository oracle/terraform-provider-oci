// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_email "github.com/oracle/oci-go-sdk/v56/email"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SuppressionResourceConfig = SuppressionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_suppression", "test_suppression", acctest.Optional, acctest.Update, suppressionRepresentation)

	suppressionSingularDataSourceRepresentation = map[string]interface{}{
		"suppression_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_email_suppression.test_suppression.id}`},
	}

	suppressionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"email_address":                         acctest.Representation{RepType: acctest.Optional, Create: `johnsmithtester@example.com`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: suppressionDataSourceFilterRepresentation}}
	suppressionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_email_suppression.test_suppression.id}`}},
	}

	suppressionRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"email_address":  acctest.Representation{RepType: acctest.Required, Create: `johnsmithtester@example.com`},
	}

	SuppressionResourceDependencies = ""
)

// issue-routing-tag: email/default
func TestEmailSuppressionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmailSuppressionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_email_suppression.test_suppression"
	datasourceName := "data.oci_email_suppressions.test_suppressions"
	singularDatasourceName := "data.oci_email_suppression.test_suppression"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SuppressionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_email_suppression", "test_suppression", acctest.Required, acctest.Create, suppressionRepresentation), "email", "suppression", t)

	acctest.ResourceTest(t, testAccCheckEmailSuppressionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SuppressionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_suppression", "test_suppression", acctest.Required, acctest.Create, suppressionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				// email address is converted to lower case by the service
				resource.TestCheckResourceAttr(resourceName, "email_address", "johnsmithtester@example.com"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_suppressions", "test_suppressions", acctest.Optional, acctest.Update, suppressionDataSourceRepresentation) +
				compartmentIdVariableStr + SuppressionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_suppression", "test_suppression", acctest.Optional, acctest.Update, suppressionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "email_address", "johnsmithtester@example.com"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "suppressions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "suppressions.0.compartment_id", tenancyId),
				// email address is converted to lower case by the service
				resource.TestCheckResourceAttr(datasourceName, "suppressions.0.email_address", "johnsmithtester@example.com"),
				resource.TestCheckResourceAttrSet(datasourceName, "suppressions.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "suppressions.0.reason"),
				resource.TestCheckResourceAttrSet(datasourceName, "suppressions.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_suppression", "test_suppression", acctest.Required, acctest.Create, suppressionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SuppressionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "suppression_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "email_address", "johnsmithtester@example.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reason"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_suppressed"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + SuppressionResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckEmailSuppressionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).EmailClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_suppression" {
			noResourceFound = false
			request := oci_email.GetSuppressionRequest{}

			tmp := rs.Primary.ID
			request.SuppressionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "email")

			_, err := client.GetSuppression(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("EmailSuppression") {
		resource.AddTestSweepers("EmailSuppression", &resource.Sweeper{
			Name:         "EmailSuppression",
			Dependencies: acctest.DependencyGraph["suppression"],
			F:            sweepEmailSuppressionResource,
		})
	}
}

func sweepEmailSuppressionResource(compartment string) error {
	emailClient := acctest.GetTestClients(&schema.ResourceData{}).EmailClient()
	// EmailSuppressionResource can only run on root compartment
	compartment = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	suppressionIds, err := getSuppressionIds(compartment)
	if err != nil {
		return err
	}
	for _, suppressionId := range suppressionIds {
		if ok := acctest.SweeperDefaultResourceId[suppressionId]; !ok {
			deleteSuppressionRequest := oci_email.DeleteSuppressionRequest{}

			deleteSuppressionRequest.SuppressionId = &suppressionId

			deleteSuppressionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "email")
			_, error := emailClient.DeleteSuppression(context.Background(), deleteSuppressionRequest)
			if error != nil {
				fmt.Printf("Error deleting Suppression %s %s, It is possible that the resource is already deleted. Please verify manually \n", suppressionId, error)
				continue
			}
		}
	}
	return nil
}

func getSuppressionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SuppressionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	emailClient := acctest.GetTestClients(&schema.ResourceData{}).EmailClient()

	listSuppressionsRequest := oci_email.ListSuppressionsRequest{}
	listSuppressionsRequest.CompartmentId = &compartmentId
	listSuppressionsResponse, err := emailClient.ListSuppressions(context.Background(), listSuppressionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Suppression list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, suppression := range listSuppressionsResponse.Items {
		id := *suppression.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SuppressionId", id)
	}
	return resourceIds, nil
}
