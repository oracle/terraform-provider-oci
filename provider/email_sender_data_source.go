// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/email"

	"github.com/oracle/terraform-provider-oci/crud"
)

func SenderDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularSender,
		Schema: map[string]*schema.Schema{
			"sender_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"email_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_spf": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularSender(d *schema.ResourceData, m interface{}) error {
	sync := &SenderDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return crud.ReadResource(sync)
}

type SenderDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.GetSenderResponse
}

func (s *SenderDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SenderDataSourceCrud) Get() error {
	request := oci_email.GetSenderRequest{}

	if senderId, ok := s.D.GetOkExists("sender_id"); ok {
		tmp := senderId.(string)
		request.SenderId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "email")

	response, err := s.Client.GetSender(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SenderDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.EmailAddress != nil {
		s.D.Set("email_address", *s.Res.EmailAddress)
	}

	if s.Res.IsSpf != nil {
		s.D.Set("is_spf", *s.Res.IsSpf)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return
}
