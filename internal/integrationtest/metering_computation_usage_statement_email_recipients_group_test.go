// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MeteringComputationUsageStatementEmailRecipientsGroupResourceConfig = MeteringComputationUsageStatementEmailRecipientsGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_statement_email_recipients_group", "test_usage_statement_email_recipients_group", acctest.Optional, acctest.Update, MeteringComputationUsageStatementEmailRecipientsGroupRepresentation)

	MeteringComputationUsageStatementEmailRecipientsGroupSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"email_recipients_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.email_recipient_group_ocid}`},
		"subscription_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
	}

	MeteringComputationUsageStatementEmailRecipientsGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationUsageStatementEmailRecipientsGroupDataSourceFilterRepresentation}}
	MeteringComputationUsageStatementEmailRecipientsGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_metering_computation_usage_statement_email_recipients_group.test_usage_statement_email_recipients_group.id}`}},
	}

	MeteringComputationUsageStatementEmailRecipientsGroupRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"recipients_list": acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationUsageStatementEmailRecipientsGroupRecipientsListRepresentation},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
	}
	MeteringComputationUsageStatementEmailRecipientsGroupRecipientsListRepresentation = map[string]interface{}{
		"email_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.email_id}`},
		"state":      acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`, Update: `INACTIVE`},
		"first_name": acctest.Representation{RepType: acctest.Optional, Create: `firstName`, Update: `firstName2`},
		"last_name":  acctest.Representation{RepType: acctest.Optional, Create: `lastName`, Update: `lastName2`},
	}

	MeteringComputationUsageStatementEmailRecipientsGroupResourceDependencies = ""
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationUsageStatementEmailRecipientsGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationUsageStatementEmailRecipientsGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	emailRecipientGroupOcid := utils.GetEnvSettingWithBlankDefault("email_recipient_group_ocid")
	emailRecipientGroupOcidVariableStr := fmt.Sprintf("variable \"email_recipient_group_ocid\" { default = \"%s\" }\n", emailRecipientGroupOcid)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	emailId := utils.GetEnvSettingWithBlankDefault("email_id")
	emailIdVariableStr := fmt.Sprintf("variable \"email_id\" { default = \"%s\" }\n", emailId)

	resourceName := "oci_metering_computation_usage_statement_email_recipients_group.test_usage_statement_email_recipients_group"
	datasourceName := "data.oci_metering_computation_usage_statement_email_recipients_groups.test_usage_statement_email_recipients_groups"
	singularDatasourceName := "data.oci_metering_computation_usage_statement_email_recipients_group.test_usage_statement_email_recipients_group"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MeteringComputationUsageStatementEmailRecipientsGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_statement_email_recipients_group", "test_usage_statement_email_recipients_group", acctest.Required, acctest.Create, MeteringComputationUsageStatementEmailRecipientsGroupRepresentation), "usageapi", "usageStatementEmailRecipientsGroup", t)

	acctest.ResourceTest(t, testAccCheckMeteringComputationUsageStatementEmailRecipientsGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + MeteringComputationUsageStatementEmailRecipientsGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_statement_email_recipients_group", "test_usage_statement_email_recipients_group", acctest.Required, acctest.Create, MeteringComputationUsageStatementEmailRecipientsGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "recipients_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "recipients_list.0.email_id"),
				resource.TestCheckResourceAttr(resourceName, "recipients_list.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "subscription_id"),

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
			Config: config + compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + MeteringComputationUsageStatementEmailRecipientsGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_statement_email_recipients_group", "test_usage_statement_email_recipients_group", acctest.Optional, acctest.Update, MeteringComputationUsageStatementEmailRecipientsGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "recipients_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "recipients_list.0.email_id"),
				resource.TestCheckResourceAttr(resourceName, "recipients_list.0.first_name", "firstName2"),
				resource.TestCheckResourceAttr(resourceName, "recipients_list.0.last_name", "lastName2"),
				resource.TestCheckResourceAttr(resourceName, "recipients_list.0.state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subscription_id"),

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
			Config: config + emailRecipientGroupOcidVariableStr + subscriptionIdVariableStr + emailIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_usage_statement_email_recipients_groups", "test_usage_statement_email_recipients_groups", acctest.Optional, acctest.Update, MeteringComputationUsageStatementEmailRecipientsGroupDataSourceRepresentation) +
				compartmentIdVariableStr + MeteringComputationUsageStatementEmailRecipientsGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_statement_email_recipients_group", "test_usage_statement_email_recipients_group", acctest.Optional, acctest.Update, MeteringComputationUsageStatementEmailRecipientsGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),

				resource.TestCheckResourceAttr(datasourceName, "email_recipients_group_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "email_recipients_group_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + emailRecipientGroupOcidVariableStr + subscriptionIdVariableStr + emailIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_usage_statement_email_recipients_group", "test_usage_statement_email_recipients_group", acctest.Required, acctest.Create, MeteringComputationUsageStatementEmailRecipientsGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MeteringComputationUsageStatementEmailRecipientsGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "email_recipients_group_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recipients_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recipients_list.0.first_name", "firstName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recipients_list.0.last_name", "lastName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recipients_list.0.state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + MeteringComputationUsageStatementEmailRecipientsGroupResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMeteringComputationUsageStatementEmailRecipientsGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).UsageapiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_metering_computation_usage_statement_email_recipients_group" {
			noResourceFound = false
			request := oci_metering_computation.GetEmailRecipientsGroupRequest{}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			if value, ok := rs.Primary.Attributes["email_recipients_group_id"]; ok {
				request.EmailRecipientsGroupId = &value
			}

			if value, ok := rs.Primary.Attributes["subscription_id"]; ok {
				request.SubscriptionId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "metering_computation")

			_, err := client.GetEmailRecipientsGroup(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("MeteringComputationUsageStatementEmailRecipientsGroup") {
		resource.AddTestSweepers("MeteringComputationUsageStatementEmailRecipientsGroup", &resource.Sweeper{
			Name:         "MeteringComputationUsageStatementEmailRecipientsGroup",
			Dependencies: acctest.DependencyGraph["usageStatementEmailRecipientsGroup"],
			F:            sweepMeteringComputationUsageStatementEmailRecipientsGroupResource,
		})
	}
}

func sweepMeteringComputationUsageStatementEmailRecipientsGroupResource(compartment string) error {
	usageapiClient := acctest.GetTestClients(&schema.ResourceData{}).UsageapiClient()
	usageStatementEmailRecipientsGroupIds, err := getMeteringComputationUsageStatementEmailRecipientsGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, usageStatementEmailRecipientsGroupId := range usageStatementEmailRecipientsGroupIds {
		if ok := acctest.SweeperDefaultResourceId[usageStatementEmailRecipientsGroupId]; !ok {
			deleteEmailRecipientsGroupRequest := oci_metering_computation.DeleteEmailRecipientsGroupRequest{}

			deleteEmailRecipientsGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "metering_computation")
			_, error := usageapiClient.DeleteEmailRecipientsGroup(context.Background(), deleteEmailRecipientsGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting UsageStatementEmailRecipientsGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", usageStatementEmailRecipientsGroupId, error)
				continue
			}
		}
	}
	return nil
}

func getMeteringComputationUsageStatementEmailRecipientsGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "UsageStatementEmailRecipientsGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	usageapiClient := acctest.GetTestClients(&schema.ResourceData{}).UsageapiClient()

	listEmailRecipientsGroupsRequest := oci_metering_computation.ListEmailRecipientsGroupsRequest{}
	listEmailRecipientsGroupsRequest.CompartmentId = &compartmentId

	subscriptionIds, error := getMeteringComputationSubscriptionIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting subscriptionId required for UsageStatementEmailRecipientsGroup resource requests \n")
	}
	for _, subscriptionId := range subscriptionIds {
		listEmailRecipientsGroupsRequest.SubscriptionId = &subscriptionId

		listEmailRecipientsGroupsResponse, err := usageapiClient.ListEmailRecipientsGroups(context.Background(), listEmailRecipientsGroupsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting UsageStatementEmailRecipientsGroup list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, usageStatementEmailRecipientsGroup := range listEmailRecipientsGroupsResponse.Items {
			id := *usageStatementEmailRecipientsGroup.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "UsageStatementEmailRecipientsGroupId", id)
		}

	}
	return resourceIds, nil
}

func getMeteringComputationSubscriptionIds(compartment string) ([]string, error) {
	// subscriptionId is known to the caller, we do not expose this through our APIs
	if compartment == "ocid1.tenancy.region1..aaaaaaaaxfq6ukfwlcojrzbnothpeb6see6r4eqem4ihippjuyvqena7uqwq" {
		subscriptionIds := []string{"10153310"}
		return subscriptionIds, nil
	}
	return nil, fmt.Errorf("Error in getMeteringComputationSubscriptionIds for compartmentId : %s \n", compartment)
}
