// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_email "github.com/oracle/oci-go-sdk/email"
)

var (
	SenderResourceConfig = SenderResourceDependencies +
		generateResourceFromRepresentationMap("oci_email_sender", "test_sender", Optional, Update, senderRepresentation)

	senderSingularDataSourceRepresentation = map[string]interface{}{
		"sender_id": Representation{repType: Required, create: `${oci_email_sender.test_sender.id}`},
	}

	senderDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"email_address":  Representation{repType: Optional, create: `JohnSmith@example.com`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, senderDataSourceFilterRepresentation}}
	senderDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_email_sender.test_sender.id}`}},
	}

	senderRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"email_address":  Representation{repType: Required, create: `JohnSmith@example.com`},
	}

	SenderResourceDependencies = ""
)

func TestEmailSenderResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_email_sender.test_sender"
	datasourceName := "data.oci_email_senders.test_senders"
	singularDatasourceName := "data.oci_email_sender.test_sender"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckEmailSenderDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + SenderResourceDependencies +
					generateResourceFromRepresentationMap("oci_email_sender", "test_sender", Required, Create, senderRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "email_address", "JohnSmith@example.com"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_email_senders", "test_senders", Optional, Update, senderDataSourceRepresentation) +
					compartmentIdVariableStr + SenderResourceDependencies +
					generateResourceFromRepresentationMap("oci_email_sender", "test_sender", Optional, Update, senderRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "email_address", "JohnSmith@example.com"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "senders.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "senders.0.email_address", "JohnSmith@example.com"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_email_sender", "test_sender", Required, Create, senderSingularDataSourceRepresentation) +
					compartmentIdVariableStr + SenderResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "sender_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "email_address", "JohnSmith@example.com"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_spf", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + SenderResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"compartment_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckEmailSenderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).emailClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_sender" {
			noResourceFound = false
			request := oci_email.GetSenderRequest{}

			tmp := rs.Primary.ID
			request.SenderId = &tmp

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

func initEmailSenderSweeper() {
	resource.AddTestSweepers("EmailSender", &resource.Sweeper{
		Name:         "EmailSender",
		Dependencies: DependencyGraph["sender"],
		F:            sweepEmailSenderResource,
	})
}

func sweepEmailSenderResource(compartment string) error {
	compartmentId := compartment
	emailClient := GetTestClients(&schema.ResourceData{}).emailClient

	listSendersRequest := oci_email.ListSendersRequest{}
	listSendersRequest.CompartmentId = &compartmentId
	listSendersRequest.LifecycleState = oci_email.SenderLifecycleStateActive
	listSendersResponse, err := emailClient.ListSenders(context.Background(), listSendersRequest)

	if err != nil {
		return fmt.Errorf("Error getting Sender list for compartment id : %s , %s \n", compartmentId, err)
	}

	for _, sender := range listSendersResponse.Items {
		if sender.LifecycleState != oci_email.SenderSummaryLifecycleStateDeleted {
			log.Printf("deleting sender %s ", *sender.Id)

			deleteSenderRequest := oci_email.DeleteSenderRequest{}

			deleteSenderRequest.SenderId = sender.Id

			deleteSenderRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "email")
			_, error := emailClient.DeleteSender(context.Background(), deleteSenderRequest)
			if error != nil {
				fmt.Printf("Error deleting Sender %s %s, It is possible that the resource is already deleted. Please verify manually \n", *sender.Id, error)
				continue
			}

			getSenderRequest := oci_email.GetSenderRequest{}

			getSenderRequest.SenderId = sender.Id

			_, error = emailClient.GetSender(context.Background(), getSenderRequest)
			if error != nil {
				fmt.Printf("Error retrieving Sender state %s \n", error)
				continue
			}

			waitTillCondition(testAccProvider, sender.Id, senderSweepWaitCondition, time.Duration(3*time.Minute),
				senderSweepResponseFetchOperation, "email", true)
		}
	}
	return nil
}

func senderSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if senderResponse, ok := response.Response.(oci_email.GetSenderResponse); ok {
		return senderResponse.LifecycleState == oci_email.SenderLifecycleStateDeleted
	}
	return false
}

func senderSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.emailClient.GetSender(context.Background(), oci_email.GetSenderRequest{
		SenderId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
