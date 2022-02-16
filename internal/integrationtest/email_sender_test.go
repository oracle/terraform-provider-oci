// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_email "github.com/oracle/oci-go-sdk/v58/email"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SenderRequiredOnlyResource = SenderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Required, acctest.Create, senderRepresentation)

	SenderResourceConfig = SenderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Update, senderRepresentation)

	senderSingularDataSourceRepresentation = map[string]interface{}{
		"sender_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_email_sender.test_sender.id}`},
	}

	senderDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"domain":         acctest.Representation{RepType: acctest.Optional, Create: `example.com`},
		"email_address":  acctest.Representation{RepType: acctest.Optional, Create: `johnsmithtester@example.com`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: senderDataSourceFilterRepresentation}}
	senderDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_email_sender.test_sender.id}`}},
	}

	senderRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"email_address":  acctest.Representation{RepType: acctest.Required, Create: `johnsmithtester@example.com`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	SenderResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: email/default
func TestEmailSenderResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmailSenderResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_email_sender.test_sender"
	datasourceName := "data.oci_email_senders.test_senders"
	singularDatasourceName := "data.oci_email_sender.test_sender"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SenderResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Create, senderRepresentation), "email", "sender", t)

	acctest.ResourceTest(t, testAccCheckEmailSenderDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SenderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Required, acctest.Create, senderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "email_address", "johnsmithtester@example.com"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SenderResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SenderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Create, senderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "email_address", "johnsmithtester@example.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SenderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(senderRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "email_address", "johnsmithtester@example.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + SenderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Update, senderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "email_address", "johnsmithtester@example.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_senders", "test_senders", acctest.Optional, acctest.Update, senderDataSourceRepresentation) +
				compartmentIdVariableStr + SenderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Update, senderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "email_address", "johnsmithtester@example.com"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "senders.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "senders.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "senders.0.email_address", "johnsmithtester@example.com"),
				resource.TestCheckResourceAttr(datasourceName, "senders.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "senders.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "senders.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "senders.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Required, acctest.Create, senderSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SenderResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sender_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "email_address", "johnsmithtester@example.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_spf"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + SenderResourceConfig,
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

func testAccCheckEmailSenderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).EmailClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_sender" {
			noResourceFound = false
			request := oci_email.GetSenderRequest{}

			tmp := rs.Primary.ID
			request.SenderId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "email")

			response, err := client.GetSender(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_email.SenderLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("EmailSender") {
		resource.AddTestSweepers("EmailSender", &resource.Sweeper{
			Name:         "EmailSender",
			Dependencies: acctest.DependencyGraph["sender"],
			F:            sweepEmailSenderResource,
		})
	}
}

func sweepEmailSenderResource(compartment string) error {
	emailClient := acctest.GetTestClients(&schema.ResourceData{}).EmailClient()
	senderIds, err := getSenderIds(compartment)
	if err != nil {
		return err
	}
	for _, senderId := range senderIds {
		if ok := acctest.SweeperDefaultResourceId[senderId]; !ok {
			deleteSenderRequest := oci_email.DeleteSenderRequest{}

			deleteSenderRequest.SenderId = &senderId

			deleteSenderRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "email")
			_, error := emailClient.DeleteSender(context.Background(), deleteSenderRequest)
			if error != nil {
				fmt.Printf("Error deleting Sender %s %s, It is possible that the resource is already deleted. Please verify manually \n", senderId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &senderId, senderSweepWaitCondition, time.Duration(3*time.Minute),
				senderSweepResponseFetchOperation, "email", true)
		}
	}
	return nil
}

func getSenderIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SenderId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	emailClient := acctest.GetTestClients(&schema.ResourceData{}).EmailClient()

	listSendersRequest := oci_email.ListSendersRequest{}
	listSendersRequest.CompartmentId = &compartmentId
	listSendersRequest.LifecycleState = oci_email.SenderLifecycleStateActive
	listSendersResponse, err := emailClient.ListSenders(context.Background(), listSendersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Sender list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sender := range listSendersResponse.Items {
		id := *sender.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SenderId", id)
	}
	return resourceIds, nil
}

func senderSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if senderResponse, ok := response.Response.(oci_email.GetSenderResponse); ok {
		return senderResponse.LifecycleState != oci_email.SenderLifecycleStateDeleted
	}
	return false
}

func senderSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.EmailClient().GetSender(context.Background(), oci_email.GetSenderRequest{
		SenderId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
