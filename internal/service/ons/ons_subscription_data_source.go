// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ons

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ons "github.com/oracle/oci-go-sdk/v56/ons"
)

func OnsSubscriptionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["subscription_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OnsSubscriptionResource(), fieldMap, readSingularOnsSubscription)
}

func readSingularOnsSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &OnsSubscriptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NotificationDataPlaneClient()

	return tfresource.ReadResource(sync)
}

type OnsSubscriptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ons.NotificationDataPlaneClient
	Res    *oci_ons.GetSubscriptionResponse
}

func (s *OnsSubscriptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnsSubscriptionDataSourceCrud) Get() error {
	request := oci_ons.GetSubscriptionRequest{}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ons")

	response, err := s.Client.GetSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OnsSubscriptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedTime != nil {
		s.D.Set("created_time", strconv.FormatInt(*s.Res.CreatedTime, 10))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeliverPolicy != nil {
		s.D.Set("delivery_policy", *s.Res.DeliverPolicy)
	}

	if s.Res.Endpoint != nil {
		s.D.Set("endpoint", *s.Res.Endpoint)
	}

	if s.Res.Etag != nil {
		s.D.Set("etag", *s.Res.Etag)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Protocol != nil {
		s.D.Set("protocol", *s.Res.Protocol)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TopicId != nil {
		s.D.Set("topic_id", *s.Res.TopicId)
	}

	return nil
}
