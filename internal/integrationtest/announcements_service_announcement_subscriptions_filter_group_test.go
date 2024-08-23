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
	AnnouncementsServiceAnnouncementSubscriptionsFilterGroupRepresentation = map[string]interface{}{
		"announcement_subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_announcements_service_announcement_subscription.test_announcement_subscription.id}`},
		"filters":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: AnnouncementsServiceAnnouncementSubscriptionsFilterGroupFiltersRepresentation},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: `fg-name`},
	}
	AnnouncementsServiceAnnouncementSubscriptionsFilterGroupFiltersRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `COMPARTMENT_ID`, Update: `PLATFORM_TYPE`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `IAAS`},
	}

	AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Required, acctest.Create, AnnouncementsServiceAnnouncementSubscriptionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: announcements_service/default
func TestAnnouncementsServiceAnnouncementSubscriptionsFilterGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAnnouncementsServiceAnnouncementSubscriptionsFilterGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_announcements_service_announcement_subscriptions_filter_group.test_announcement_subscriptions_filter_group"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscriptions_filter_group", "test_announcement_subscriptions_filter_group", acctest.Required, acctest.Create, AnnouncementsServiceAnnouncementSubscriptionsFilterGroupRepresentation), "announcementsservice", "announcementSubscriptionsFilterGroup", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscriptions_filter_group", "test_announcement_subscriptions_filter_group", acctest.Required, acctest.Create, AnnouncementsServiceAnnouncementSubscriptionsFilterGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "announcement_subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "filters.0.type", "COMPARTMENT_ID"),
				resource.TestCheckResourceAttr(resourceName, "filters.0.value", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "fg-name"),
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
			Config: config + compartmentIdVariableStr + AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscriptions_filter_group", "test_announcement_subscriptions_filter_group", acctest.Optional, acctest.Update, AnnouncementsServiceAnnouncementSubscriptionsFilterGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "announcement_subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "filters.0.type", "PLATFORM_TYPE"),
				resource.TestCheckResourceAttr(resourceName, "filters.0.value", "IAAS"),
				resource.TestCheckResourceAttr(resourceName, "name", "fg-name"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}
