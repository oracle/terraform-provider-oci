// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_ons "github.com/oracle/oci-go-sdk/ons"
)

func SubscriptionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularSubscription,
		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"delivery_policy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"protocol": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &SubscriptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationDataPlaneClient

	return ReadResource(sync)
}

type SubscriptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ons.NotificationDataPlaneClient
	Res    *oci_ons.GetSubscriptionResponse
}

func (s *SubscriptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SubscriptionDataSourceCrud) Get() error {
	request := oci_ons.GetSubscriptionRequest{}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "ons")

	response, err := s.Client.GetSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SubscriptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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

	return nil
}
