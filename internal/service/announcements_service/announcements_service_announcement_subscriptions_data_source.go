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

func AnnouncementsServiceAnnouncementSubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAnnouncementsServiceAnnouncementSubscriptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"announcement_subscription_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(AnnouncementsServiceAnnouncementSubscriptionResource()),
						},
					},
				},
			},
		},
	}
}

func readAnnouncementsServiceAnnouncementSubscriptions(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceAnnouncementSubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnnouncementSubscriptionClient()

	return tfresource.ReadResource(sync)
}

type AnnouncementsServiceAnnouncementSubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_announcements_service.AnnouncementSubscriptionClient
	Res    *oci_announcements_service.ListAnnouncementSubscriptionsResponse
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsDataSourceCrud) Get() error {
	request := oci_announcements_service.ListAnnouncementSubscriptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_announcements_service.AnnouncementSubscriptionLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "announcements_service")

	response, err := s.Client.ListAnnouncementSubscriptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAnnouncementSubscriptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AnnouncementsServiceAnnouncementSubscriptionsDataSource-", AnnouncementsServiceAnnouncementSubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	announcementSubscription := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AnnouncementSubscriptionSummaryToMap(item))
	}
	announcementSubscription["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AnnouncementsServiceAnnouncementSubscriptionsDataSource().Schema["announcement_subscription_collection"].Elem.(*schema.Resource).Schema)
		announcementSubscription["items"] = items
	}

	resources = append(resources, announcementSubscription)
	if err := s.D.Set("announcement_subscription_collection", resources); err != nil {
		return err
	}

	return nil
}
