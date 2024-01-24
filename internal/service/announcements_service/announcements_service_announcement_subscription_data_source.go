// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package announcements_service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_announcements_service "github.com/oracle/oci-go-sdk/v65/announcementsservice"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AnnouncementsServiceAnnouncementSubscriptionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["announcement_subscription_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AnnouncementsServiceAnnouncementSubscriptionResource(), fieldMap, readSingularAnnouncementsServiceAnnouncementSubscription)
}

func readSingularAnnouncementsServiceAnnouncementSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceAnnouncementSubscriptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnnouncementSubscriptionClient()

	return tfresource.ReadResource(sync)
}

type AnnouncementsServiceAnnouncementSubscriptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_announcements_service.AnnouncementSubscriptionClient
	Res    *oci_announcements_service.GetAnnouncementSubscriptionResponse
}

func (s *AnnouncementsServiceAnnouncementSubscriptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AnnouncementsServiceAnnouncementSubscriptionDataSourceCrud) Get() error {
	request := oci_announcements_service.GetAnnouncementSubscriptionRequest{}

	if announcementSubscriptionId, ok := s.D.GetOkExists("announcement_subscription_id"); ok {
		tmp := announcementSubscriptionId.(string)
		request.AnnouncementSubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "announcements_service")

	response, err := s.Client.GetAnnouncementSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AnnouncementsServiceAnnouncementSubscriptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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
