// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	EmailSenderRequiredOnlyResource = EmailSenderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Required, acctest.Create, EmailSenderRepresentation)

	EmailSenderResourceConfig = EmailSenderResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Update, EmailSenderRepresentation)

	EmailEmailSenderSingularDataSourceRepresentation = map[string]interface{}{
		"sender_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_email_sender.test_sender.id}`},
	}

	EmailEmailSenderDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"domain":         acctest.Representation{RepType: acctest.Optional, Create: `example.com`},
		"email_address":  acctest.Representation{RepType: acctest.Optional, Create: `johnsmithtester@example.com`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: EmailSenderDataSourceFilterRepresentation}}
	EmailSenderDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_email_sender.test_sender.id}`}},
	}

	EmailSenderRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"email_address":  acctest.Representation{RepType: acctest.Required, Create: `johnsmithtester@example.com`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	EmailSenderResourceDependencies = DefinedTagsDependencies
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+EmailSenderResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Create, EmailSenderRepresentation), "email", "sender", t)

	acctest.ResourceTest(t, testAccCheckEmailSenderDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + EmailSenderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Required, acctest.Create, EmailSenderRepresentation),
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
			Config: config + compartmentIdVariableStr + EmailSenderResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + EmailSenderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Create, EmailSenderRepresentation),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + EmailSenderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(EmailSenderRepresentation, map[string]interface{}{
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
			Config: config + compartmentIdVariableStr + EmailSenderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Update, EmailSenderRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_senders", "test_senders", acctest.Optional, acctest.Update, EmailEmailSenderDataSourceRepresentation) +
				compartmentIdVariableStr + EmailSenderResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Optional, acctest.Update, EmailSenderRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_email_sender", "test_sender", acctest.Required, acctest.Create, EmailEmailSenderSingularDataSourceRepresentation) +
				compartmentIdVariableStr + EmailSenderResourceConfig,
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
		// verify resource import
		{
			Config:                  config + EmailSenderRequiredOnlyResource,
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
	senderIds, err := getEmailSenderIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &senderId, EmailSenderSweepWaitCondition, time.Duration(3*time.Minute),
				EmailSenderSweepResponseFetchOperation, "email", true)
		}
	}
	return nil
}

func getEmailSenderIds(compartment string) ([]string, error) {
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

func EmailSenderSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if senderResponse, ok := response.Response.(oci_email.GetSenderResponse); ok {
		return senderResponse.LifecycleState != oci_email.SenderLifecycleStateDeleted
	}
	return false
}

func EmailSenderSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.EmailClient().GetSender(context.Background(), oci_email.GetSenderRequest{
		SenderId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
