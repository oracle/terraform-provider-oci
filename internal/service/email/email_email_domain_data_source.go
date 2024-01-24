// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"
)

func EmailEmailDomainDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["email_domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(EmailEmailDomainResource(), fieldMap, readSingularEmailEmailDomain)
}

func readSingularEmailEmailDomain(d *schema.ResourceData, m interface{}) error {
	sync := &EmailEmailDomainDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

type EmailEmailDomainDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.GetEmailDomainResponse
}

func (s *EmailEmailDomainDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailEmailDomainDataSourceCrud) Get() error {
	request := oci_email.GetEmailDomainRequest{}

	if emailDomainId, ok := s.D.GetOkExists("email_domain_id"); ok {
		tmp := emailDomainId.(string)
		request.EmailDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

	response, err := s.Client.GetEmailDomain(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *EmailEmailDomainDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ActiveDkimId != nil {
		s.D.Set("active_dkim_id", *s.Res.ActiveDkimId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSpf != nil {
		s.D.Set("is_spf", *s.Res.IsSpf)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
