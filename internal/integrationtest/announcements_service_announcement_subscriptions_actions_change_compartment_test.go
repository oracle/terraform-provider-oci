// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentRepresentation = map[string]interface{}{
		"announcement_subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_announcements_service_announcement_subscription.test_announcement_subscription.id}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Required, acctest.Create, AnnouncementsServiceAnnouncementSubscriptionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: announcements_service/default
func TestAnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_announcements_service_announcement_subscriptions_actions_change_compartment.test_announcement_subscriptions_actions_change_compartment"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscriptions_actions_change_compartment", "test_announcement_subscriptions_actions_change_compartment", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentRepresentation, map[string]interface{}{
			"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
		})), "announcementsservice", "announcementSubscriptionsActionsChangeCompartment", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdUVariableStr + compartmentIdVariableStr + AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscriptions_actions_change_compartment", "test_announcement_subscriptions_actions_change_compartment", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentRepresentation, map[string]interface{}{
					"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
				})),
			ExpectNonEmptyPlan: true,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "announcement_subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentIdU, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
