// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
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
	IdentityDomainsNotificationSettingRequiredOnlyResource = IdentityDomainsNotificationSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_notification_setting", "test_notification_setting", acctest.Required, acctest.Create, IdentityDomainsNotificationSettingRepresentation)

	IdentityDomainsNotificationSettingResourceConfig = IdentityDomainsNotificationSettingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_notification_setting", "test_notification_setting", acctest.Optional, acctest.Update, IdentityDomainsNotificationSettingRepresentation)

	IdentityDomainsNotificationSettingSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"notification_setting_id": acctest.Representation{RepType: acctest.Required, Create: `NotificationSettings`},
		"attribute_sets":          acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsNotificationSettingDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsNotificationSettingRepresentation = map[string]interface{}{
		"event_settings":          acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsNotificationSettingEventSettingsRepresentation},
		"from_email_address":      acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsNotificationSettingFromEmailAddressRepresentation},
		"idcs_endpoint":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"notification_enabled":    acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"notification_setting_id": acctest.Representation{RepType: acctest.Required, Create: `NotificationSettings`},
		"schemas":                 acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:NotificationSettings`}},
		"attribute_sets":          acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"external_id":             acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
		"send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"send_notifications_to_secondary_email":                                            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tags":                                                                             acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsNotificationSettingTagsRepresentation},
		"test_mode_enabled":                                                                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"test_recipients":                                                                  acctest.Representation{RepType: acctest.Optional, Create: []string{`testRecipients@email.com`}, Update: []string{`testRecipients2@test.com`}},
	}
	IdentityDomainsNotificationSettingEventSettingsRepresentation = map[string]interface{}{
		"event_id": acctest.Representation{RepType: acctest.Required, Create: `admin.user.create.success`},
		"enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	IdentityDomainsNotificationSettingFromEmailAddressRepresentation = map[string]interface{}{
		"validate":     acctest.Representation{RepType: acctest.Required, Create: `email`, Update: `domain`},
		"value":        acctest.Representation{RepType: acctest.Required, Create: `value@email.com`, Update: `value2@email.com`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}
	IdentityDomainsNotificationSettingTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsNotificationSettingResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsNotificationSettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsNotificationSettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_notification_setting.test_notification_setting"
	datasourceName := "data.oci_identity_domains_notification_settings.test_notification_settings"
	singularDatasourceName := "data.oci_identity_domains_notification_setting.test_notification_setting"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsNotificationSettingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_notification_setting", "test_notification_setting", acctest.Optional, acctest.Create, IdentityDomainsNotificationSettingRepresentation), "identitydomains", "notificationSetting", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsNotificationSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_notification_setting", "test_notification_setting", acctest.Required, acctest.Create, IdentityDomainsNotificationSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "event_settings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "event_settings.0.event_id"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.0.validate", "email"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.0.value", "value@email.com"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "notification_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsNotificationSettingResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsNotificationSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_notification_setting", "test_notification_setting", acctest.Optional, acctest.Create, IdentityDomainsNotificationSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "event_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "event_settings.0.enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "event_settings.0.event_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.0.validate", "email"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.0.value", "value@email.com"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "notification_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email", "false"),
				resource.TestCheckResourceAttr(resourceName, "send_notifications_to_secondary_email", "false"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "test_mode_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "test_recipients.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "notificationSettings", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsNotificationSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_notification_setting", "test_notification_setting", acctest.Optional, acctest.Update, IdentityDomainsNotificationSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "event_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "event_settings.0.enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "event_settings.0.event_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_id"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.0.validate", "domain"),
				resource.TestCheckResourceAttr(resourceName, "from_email_address.0.value", "value2@email.com"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "notification_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_setting_id"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email", "true"),
				resource.TestCheckResourceAttr(resourceName, "send_notifications_to_secondary_email", "true"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "test_mode_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "test_recipients.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_notification_settings", "test_notification_settings", acctest.Optional, acctest.Update, IdentityDomainsNotificationSettingDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsNotificationSettingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_notification_setting", "test_notification_setting", acctest.Optional, acctest.Update, IdentityDomainsNotificationSettingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "notification_settings.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "notification_settings.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_notification_setting", "test_notification_setting", acctest.Required, acctest.Create, IdentityDomainsNotificationSettingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsNotificationSettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notification_setting_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "event_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "event_settings.0.enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "from_email_address.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "from_email_address.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "from_email_address.0.validate", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "from_email_address.0.value", "value2@email.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "send_notifications_to_secondary_email", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_mode_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_recipients.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsNotificationSettingRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_notification_setting", "notificationSettings"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"idcs_last_upgraded_in_release",
				"tags",
				"notification_setting_id",
			},
			ResourceName: resourceName,
		},
	})
}
