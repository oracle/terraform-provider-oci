// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ons

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ons "github.com/oracle/oci-go-sdk/v65/ons"
)

func OnsNotificationTopicsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOnsNotificationTopics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"notification_topics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(OnsNotificationTopicResource()),
			},
		},
	}
}

func readOnsNotificationTopics(d *schema.ResourceData, m interface{}) error {
	sync := &OnsNotificationTopicsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NotificationControlPlaneClient()

	return tfresource.ReadResource(sync)
}

type OnsNotificationTopicsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ons.NotificationControlPlaneClient
	Res    *oci_ons.ListTopicsResponse
}

func (s *OnsNotificationTopicsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnsNotificationTopicsDataSourceCrud) Get() error {
	request := oci_ons.ListTopicsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ons.NotificationTopicSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ons")

	response, err := s.Client.ListTopics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTopics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OnsNotificationTopicsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OnsNotificationTopicsDataSource-", OnsNotificationTopicsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		notificationTopic := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ApiEndpoint != nil {
			notificationTopic["api_endpoint"] = *r.ApiEndpoint
		}

		if r.DefinedTags != nil {
			notificationTopic["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			notificationTopic["description"] = *r.Description
		}

		if r.Etag != nil {
			notificationTopic["etag"] = *r.Etag
		}

		notificationTopic["freeform_tags"] = r.FreeformTags

		if r.Name != nil {
			notificationTopic["name"] = *r.Name
		}

		if r.ShortTopicId != nil {
			notificationTopic["short_topic_id"] = *r.ShortTopicId
		}

		notificationTopic["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			notificationTopic["time_created"] = r.TimeCreated.String()
		}

		if r.TopicId != nil {
			notificationTopic["topic_id"] = *r.TopicId
		}

		resources = append(resources, notificationTopic)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OnsNotificationTopicsDataSource().Schema["notification_topics"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("notification_topics", resources); err != nil {
		return err
	}

	return nil
}
