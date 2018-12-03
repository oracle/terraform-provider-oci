// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
	oci_ons "github.com/oracle/oci-go-sdk/ons"
)

var nonPendingState = regexp.MustCompile(`Subscription(.*)is not in pending state\.`)

func ReconfirmationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularReconfirmation,
		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			//Computed
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularReconfirmation(d *schema.ResourceData, m interface{}) error {
	sync := &ReconfirmationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationDataPlaneClient

	return ReadResource(sync)
}

type ReconfirmationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ons.NotificationDataPlaneClient
	Res    *oci_ons.SubscriptionConfirmation
}

func (s *ReconfirmationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ReconfirmationDataSourceCrud) Get() error {
	request := oci_ons.ResendSubscriptionConfirmationRequest{}

	if id, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "ons")

	response, err := s.Client.ResendSubscriptionConfirmation(context.Background(), request)
	if err != nil {
		if nonPendingState.MatchString(err.Error()) {
			s.Res = &oci_ons.SubscriptionConfirmation{}
			s.Res.Id = request.Id
			if url, ok := s.D.GetOkExists("url"); ok {
				tmp := url.(string)
				s.Res.Url = &tmp
			} else {
				tmp := "Subscription is not in pending state."
				s.Res.Url = &tmp
			}
			return nil
		}
		return err
	}

	s.Res = &response.SubscriptionConfirmation
	return nil
}

func (s *ReconfirmationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Url != nil {
		s.D.Set("url", *s.Res.Url)
	}

	return nil
}
