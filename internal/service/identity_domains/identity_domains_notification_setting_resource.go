// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsNotificationSettingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsNotificationSetting,
		Read:     readIdentityDomainsNotificationSetting,
		Update:   updateIdentityDomainsNotificationSetting,
		Delete:   deleteIdentityDomainsNotificationSetting,
		Schema: map[string]*schema.Schema{
			// Required
			"event_settings": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"event_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"from_email_address": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"validate": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"validation_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"notification_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"notification_setting_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"send_notifications_to_secondary_email": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"test_mode_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"test_recipients": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_modified_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_upgraded_in_release": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_prevented_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"meta": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityDomainsNotificationSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsNotificationSettingResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.CreateResource(d, sync)
}

func readIdentityDomainsNotificationSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsNotificationSettingResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "notificationSettings")
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

func updateIdentityDomainsNotificationSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsNotificationSettingResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityDomainsNotificationSetting(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityDomainsNotificationSettingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.NotificationSetting
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsNotificationSettingResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsNotificationSettingResourceCrud) Create() error {
	request := oci_identity_domains.PutNotificationSettingRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if eventSettings, ok := s.D.GetOkExists("event_settings"); ok {
		interfaces := eventSettings.([]interface{})
		tmp := make([]oci_identity_domains.NotificationSettingsEventSettings, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "event_settings", stateDataIndex)
			converted, err := s.mapToNotificationSettingsEventSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("event_settings") {
			request.EventSettings = tmp
		}
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if fromEmailAddress, ok := s.D.GetOkExists("from_email_address"); ok {
		if tmpList := fromEmailAddress.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "from_email_address", 0)
			tmp, err := s.mapToNotificationSettingsFromEmailAddress(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FromEmailAddress = &tmp
		}
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if notificationEnabled, ok := s.D.GetOkExists("notification_enabled"); ok {
		tmp := notificationEnabled.(bool)
		request.NotificationEnabled = &tmp
	}

	if notificationSettingId, ok := s.D.GetOkExists("notification_setting_id"); ok {
		tmp := notificationSettingId.(string)
		request.NotificationSettingId = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if sendNotificationToOldAndNewPrimaryEmailsWhenAdminChangesPrimaryEmail, ok := s.D.GetOkExists("send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email"); ok {
		tmp := sendNotificationToOldAndNewPrimaryEmailsWhenAdminChangesPrimaryEmail.(bool)
		request.SendNotificationToOldAndNewPrimaryEmailsWhenAdminChangesPrimaryEmail = &tmp
	}

	if sendNotificationsToSecondaryEmail, ok := s.D.GetOkExists("send_notifications_to_secondary_email"); ok {
		tmp := sendNotificationsToSecondaryEmail.(bool)
		request.SendNotificationsToSecondaryEmail = &tmp
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if testModeEnabled, ok := s.D.GetOkExists("test_mode_enabled"); ok {
		tmp := testModeEnabled.(bool)
		request.TestModeEnabled = &tmp
	}

	if testRecipients, ok := s.D.GetOkExists("test_recipients"); ok {
		interfaces := testRecipients.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("test_recipients") {
			request.TestRecipients = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutNotificationSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotificationSetting
	return nil
}

func (s *IdentityDomainsNotificationSettingResourceCrud) Get() error {
	request := oci_identity_domains.GetNotificationSettingRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	tmp := s.D.Id()
	request.NotificationSettingId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	notificationSettingId, err := parseNotificationSettingCompositeId(s.D.Id())
	if err == nil {
		request.NotificationSettingId = &notificationSettingId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetNotificationSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotificationSetting
	return nil
}

func (s *IdentityDomainsNotificationSettingResourceCrud) Update() error {
	request := oci_identity_domains.PutNotificationSettingRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if eventSettings, ok := s.D.GetOkExists("event_settings"); ok {
		interfaces := eventSettings.([]interface{})
		tmp := make([]oci_identity_domains.NotificationSettingsEventSettings, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "event_settings", stateDataIndex)
			converted, err := s.mapToNotificationSettingsEventSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("event_settings") {
			request.EventSettings = tmp
		}
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if fromEmailAddress, ok := s.D.GetOkExists("from_email_address"); ok {
		if tmpList := fromEmailAddress.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "from_email_address", 0)
			tmp, err := s.mapToNotificationSettingsFromEmailAddress(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FromEmailAddress = &tmp
		}
	}

	tmp := s.D.Id()
	request.Id = &tmp

	if notificationEnabled, ok := s.D.GetOkExists("notification_enabled"); ok {
		tmp := notificationEnabled.(bool)
		request.NotificationEnabled = &tmp
	}

	tmp = s.D.Id()
	request.NotificationSettingId = &tmp

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if sendNotificationToOldAndNewPrimaryEmailsWhenAdminChangesPrimaryEmail, ok := s.D.GetOkExists("send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email"); ok {
		tmp := sendNotificationToOldAndNewPrimaryEmailsWhenAdminChangesPrimaryEmail.(bool)
		request.SendNotificationToOldAndNewPrimaryEmailsWhenAdminChangesPrimaryEmail = &tmp
	}

	if sendNotificationsToSecondaryEmail, ok := s.D.GetOkExists("send_notifications_to_secondary_email"); ok {
		tmp := sendNotificationsToSecondaryEmail.(bool)
		request.SendNotificationsToSecondaryEmail = &tmp
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if testModeEnabled, ok := s.D.GetOkExists("test_mode_enabled"); ok {
		tmp := testModeEnabled.(bool)
		request.TestModeEnabled = &tmp
	}

	if testRecipients, ok := s.D.GetOkExists("test_recipients"); ok {
		interfaces := testRecipients.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("test_recipients") {
			request.TestRecipients = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutNotificationSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotificationSetting
	return nil
}

func (s *IdentityDomainsNotificationSettingResourceCrud) SetData() error {

	notificationSettingId, err := parseNotificationSettingCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(notificationSettingId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	eventSettings := []interface{}{}
	for _, item := range s.Res.EventSettings {
		eventSettings = append(eventSettings, NotificationSettingsEventSettingsToMap(item))
	}
	s.D.Set("event_settings", eventSettings)

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.FromEmailAddress != nil {
		s.D.Set("from_email_address", []interface{}{NotificationSettingsFromEmailAddressToMap(s.Res.FromEmailAddress)})
	} else {
		s.D.Set("from_email_address", nil)
	}

	if s.Res.IdcsCreatedBy != nil {
		s.D.Set("idcs_created_by", []interface{}{idcsCreatedByToMap(s.Res.IdcsCreatedBy)})
	} else {
		s.D.Set("idcs_created_by", nil)
	}

	if s.Res.IdcsLastModifiedBy != nil {
		s.D.Set("idcs_last_modified_by", []interface{}{idcsLastModifiedByToMap(s.Res.IdcsLastModifiedBy)})
	} else {
		s.D.Set("idcs_last_modified_by", nil)
	}

	if s.Res.IdcsLastUpgradedInRelease != nil {
		s.D.Set("idcs_last_upgraded_in_release", *s.Res.IdcsLastUpgradedInRelease)
	}

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.NotificationEnabled != nil {
		s.D.Set("notification_enabled", *s.Res.NotificationEnabled)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.SendNotificationToOldAndNewPrimaryEmailsWhenAdminChangesPrimaryEmail != nil {
		s.D.Set("send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email", *s.Res.SendNotificationToOldAndNewPrimaryEmailsWhenAdminChangesPrimaryEmail)
	}

	if s.Res.SendNotificationsToSecondaryEmail != nil {
		s.D.Set("send_notifications_to_secondary_email", *s.Res.SendNotificationsToSecondaryEmail)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.TestModeEnabled != nil {
		s.D.Set("test_mode_enabled", *s.Res.TestModeEnabled)
	}

	s.D.Set("test_recipients", s.Res.TestRecipients)

	return nil
}

func parseNotificationSettingCompositeId(compositeId string) (notificationSettingId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/notificationSettings/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	notificationSettingId, _ = url.PathUnescape(parts[3])

	return
}

func NotificationSettingToMap(obj oci_identity_domains.NotificationSetting) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	eventSettings := []interface{}{}
	for _, item := range obj.EventSettings {
		eventSettings = append(eventSettings, NotificationSettingsEventSettingsToMap(item))
	}
	result["event_settings"] = eventSettings

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	if obj.FromEmailAddress != nil {
		result["from_email_address"] = []interface{}{NotificationSettingsFromEmailAddressToMap(obj.FromEmailAddress)}
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.IdcsLastUpgradedInRelease != nil {
		result["idcs_last_upgraded_in_release"] = string(*obj.IdcsLastUpgradedInRelease)
	}

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.NotificationEnabled != nil {
		result["notification_enabled"] = bool(*obj.NotificationEnabled)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["schemas"] = obj.Schemas

	if obj.SendNotificationToOldAndNewPrimaryEmailsWhenAdminChangesPrimaryEmail != nil {
		result["send_notification_to_old_and_new_primary_emails_when_admin_changes_primary_email"] = bool(*obj.SendNotificationToOldAndNewPrimaryEmailsWhenAdminChangesPrimaryEmail)
	}

	if obj.SendNotificationsToSecondaryEmail != nil {
		result["send_notifications_to_secondary_email"] = bool(*obj.SendNotificationsToSecondaryEmail)
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.TestModeEnabled != nil {
		result["test_mode_enabled"] = bool(*obj.TestModeEnabled)
	}

	result["test_recipients"] = obj.TestRecipients

	return result
}

func (s *IdentityDomainsNotificationSettingResourceCrud) mapToNotificationSettingsEventSettings(fieldKeyFormat string) (oci_identity_domains.NotificationSettingsEventSettings, error) {
	result := oci_identity_domains.NotificationSettingsEventSettings{}

	if enabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enabled")); ok {
		tmp := enabled.(bool)
		result.Enabled = &tmp
	}

	if eventId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "event_id")); ok {
		tmp := eventId.(string)
		result.EventId = &tmp
	}

	return result, nil
}

func NotificationSettingsEventSettingsToMap(obj oci_identity_domains.NotificationSettingsEventSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Enabled != nil {
		result["enabled"] = bool(*obj.Enabled)
	}

	if obj.EventId != nil {
		result["event_id"] = string(*obj.EventId)
	}

	return result
}

func (s *IdentityDomainsNotificationSettingResourceCrud) mapToNotificationSettingsFromEmailAddress(fieldKeyFormat string) (oci_identity_domains.NotificationSettingsFromEmailAddress, error) {
	result := oci_identity_domains.NotificationSettingsFromEmailAddress{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if validate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validate")); ok {
		result.Validate = oci_identity_domains.NotificationSettingsFromEmailAddressValidateEnum(validate.(string))
	}

	if validationStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validation_status")); ok {
		result.ValidationStatus = oci_identity_domains.NotificationSettingsFromEmailAddressValidationStatusEnum(validationStatus.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func NotificationSettingsFromEmailAddressToMap(obj *oci_identity_domains.NotificationSettingsFromEmailAddress) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["validate"] = string(obj.Validate)

	result["validation_status"] = string(obj.ValidationStatus)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsNotificationSettingResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
	result := oci_identity_domains.Tags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}
