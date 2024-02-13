// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func EmailConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularEmailConfiguration,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"http_submit_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"smtp_submit_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularEmailConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &EmailConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

type EmailConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.GetEmailConfigurationResponse
}

func (s *EmailConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailConfigurationDataSourceCrud) Get() error {
	request := oci_email.GetEmailConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

	response, err := s.Client.GetEmailConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *EmailConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("EmailConfigurationDataSource-", EmailConfigurationDataSource(), s.D))

	if s.Res.HttpSubmitEndpoint != nil {
		s.D.Set("http_submit_endpoint", *s.Res.HttpSubmitEndpoint)
	}

	if s.Res.SmtpSubmitEndpoint != nil {
		s.D.Set("smtp_submit_endpoint", *s.Res.SmtpSubmitEndpoint)
	}

	return nil
}
