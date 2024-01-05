// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package announcements_service

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_announcements_service "github.com/oracle/oci-go-sdk/v65/announcementsservice"
)

func AnnouncementsServiceAnnouncementSubscriptionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAnnouncementsServiceAnnouncementSubscription,
		Read:     readAnnouncementsServiceAnnouncementSubscription,
		Update:   updateAnnouncementsServiceAnnouncementSubscription,
		Delete:   deleteAnnouncementsServiceAnnouncementSubscription,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ons_topic_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"filters": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"preferred_language": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferred_time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAnnouncementsServiceAnnouncementSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceAnnouncementSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnnouncementSubscriptionClient()

	return tfresource.CreateResource(d, sync)
}

func readAnnouncementsServiceAnnouncementSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceAnnouncementSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnnouncementSubscriptionClient()

	return tfresource.ReadResource(sync)
}

func updateAnnouncementsServiceAnnouncementSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceAnnouncementSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnnouncementSubscriptionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAnnouncementsServiceAnnouncementSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceAnnouncementSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnnouncementSubscriptionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AnnouncementsServiceAnnouncementSubscriptionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_announcements_service.AnnouncementSubscriptionClient
	Res                    *oci_announcements_service.AnnouncementSubscription
	DisableNotFoundRetries bool
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_announcements_service.AnnouncementSubscriptionLifecycleStateActive),
	}
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_announcements_service.AnnouncementSubscriptionLifecycleStateDeleted),
	}
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) Create() error {
	request := oci_announcements_service.CreateAnnouncementSubscriptionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if filterGroups, ok := s.D.GetOkExists("filter_groups"); ok {
		if tmpList := filterGroups.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "filter_groups", 0)
			tmp, err := s.mapToFilterGroupDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FilterGroups = FilterGroupDetailsToMap(&tmp)
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if onsTopicId, ok := s.D.GetOkExists("ons_topic_id"); ok {
		tmp := onsTopicId.(string)
		request.OnsTopicId = &tmp
	}

	if preferredLanguage, ok := s.D.GetOkExists("preferred_language"); ok {
		tmp := preferredLanguage.(string)
		request.PreferredLanguage = &tmp
	}

	if preferredTimeZone, ok := s.D.GetOkExists("preferred_time_zone"); ok {
		tmp := preferredTimeZone.(string)
		request.PreferredTimeZone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "announcements_service")

	response, err := s.Client.CreateAnnouncementSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnnouncementSubscription
	return nil
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) Get() error {
	request := oci_announcements_service.GetAnnouncementSubscriptionRequest{}

	tmp := s.D.Id()
	request.AnnouncementSubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "announcements_service")

	response, err := s.Client.GetAnnouncementSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnnouncementSubscription
	return nil
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_announcements_service.UpdateAnnouncementSubscriptionRequest{}

	tmp := s.D.Id()
	request.AnnouncementSubscriptionId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if onsTopicId, ok := s.D.GetOkExists("ons_topic_id"); ok {
		tmp := onsTopicId.(string)
		request.OnsTopicId = &tmp
	}

	if preferredLanguage, ok := s.D.GetOkExists("preferred_language"); ok {
		tmp := preferredLanguage.(string)
		request.PreferredLanguage = &tmp
	}

	if preferredTimeZone, ok := s.D.GetOkExists("preferred_time_zone"); ok {
		tmp := preferredTimeZone.(string)
		request.PreferredTimeZone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "announcements_service")

	response, err := s.Client.UpdateAnnouncementSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnnouncementSubscription
	return nil
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) Delete() error {
	request := oci_announcements_service.DeleteAnnouncementSubscriptionRequest{}

	tmp := s.D.Id()
	request.AnnouncementSubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "announcements_service")

	_, err := s.Client.DeleteAnnouncementSubscription(context.Background(), request)
	return err
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FilterGroups != nil {
		s.D.Set("filter_groups", FilterGroupsToMap(s.Res.FilterGroups))
	} else {
		s.D.Set("filter_groups", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OnsTopicId != nil {
		s.D.Set("ons_topic_id", *s.Res.OnsTopicId)
	}

	if s.Res.PreferredLanguage != nil {
		s.D.Set("preferred_language", *s.Res.PreferredLanguage)
	}

	if s.Res.PreferredTimeZone != nil {
		s.D.Set("preferred_time_zone", *s.Res.PreferredTimeZone)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func FilterGroupDetailsToMap(filterGroupDetails *oci_announcements_service.FilterGroupDetails) map[string]oci_announcements_service.FilterGroupDetails {
	filterGroupDetailsMap := make(map[string]oci_announcements_service.FilterGroupDetails)
	filterGroupDetailsMap["name"] = *filterGroupDetails

	return filterGroupDetailsMap
}

func FilterGroupsToMap(filterGroups map[string]oci_announcements_service.FilterGroup) []interface{} {
	result := []interface{}{}
	for _, item := range filterGroups {
		result = append(result, FilterGroupToMap(item))
	}
	return result
}

func AnnouncementSubscriptionSummaryToMap(obj oci_announcements_service.AnnouncementSubscriptionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.OnsTopicId != nil {
		result["ons_topic_id"] = string(*obj.OnsTopicId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) mapToFilter(fieldKeyFormat string) (oci_announcements_service.Filter, error) {
	result := oci_announcements_service.Filter{}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_announcements_service.FilterTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) mapToFilterGroupDetails(fieldKeyFormat string) (oci_announcements_service.FilterGroupDetails, error) {
	result := oci_announcements_service.FilterGroupDetails{}

	if filters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filters")); ok {
		interfaces := filters.([]interface{})
		tmp := make([]oci_announcements_service.Filter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "filters"), stateDataIndex)
			converted, err := s.mapToFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "filters")) {
			result.Filters = tmp
		}
	}

	return result, nil
}

func FilterGroupToMap(obj oci_announcements_service.FilterGroup) map[string]interface{} {
	result := map[string]interface{}{}

	filters := []interface{}{}
	for _, item := range obj.Filters {
		filters = append(filters, FilterToMap(item))
	}
	result["filters"] = filters

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *AnnouncementsServiceAnnouncementSubscriptionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_announcements_service.ChangeAnnouncementSubscriptionCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AnnouncementSubscriptionId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "announcements_service")

	_, err := s.Client.ChangeAnnouncementSubscriptionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
