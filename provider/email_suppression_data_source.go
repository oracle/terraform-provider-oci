// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/email"

	"github.com/oracle/terraform-provider-oci/crud"
)

func SuppressionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularSuppression,
		Schema: map[string]*schema.Schema{
			"suppression_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"email_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reason": {
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

func readSingularSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &SuppressionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return crud.ReadResource(sync)
}

type SuppressionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.GetSuppressionResponse
}

func (s *SuppressionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SuppressionDataSourceCrud) Get() error {
	request := oci_email.GetSuppressionRequest{}

	if suppressionId, ok := s.D.GetOkExists("suppression_id"); ok {
		tmp := suppressionId.(string)
		request.SuppressionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "email")

	response, err := s.Client.GetSuppression(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SuppressionDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	if s.Res.EmailAddress != nil {
		s.D.Set("email_address", *s.Res.EmailAddress)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	s.D.Set("reason", s.Res.Reason)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return
}
