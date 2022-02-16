// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/v58/email"
)

func EmailDkimDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["dkim_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(EmailDkimResource(), fieldMap, readSingularEmailDkim)
}

func readSingularEmailDkim(d *schema.ResourceData, m interface{}) error {
	sync := &EmailDkimDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

type EmailDkimDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.GetDkimResponse
}

func (s *EmailDkimDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailDkimDataSourceCrud) Get() error {
	request := oci_email.GetDkimRequest{}

	if dkimId, ok := s.D.GetOkExists("dkim_id"); ok {
		tmp := dkimId.(string)
		request.DkimId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

	response, err := s.Client.GetDkim(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *EmailDkimDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CnameRecordValue != nil {
		s.D.Set("cname_record_value", *s.Res.CnameRecordValue)
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

	if s.Res.DnsSubdomainName != nil {
		s.D.Set("dns_subdomain_name", *s.Res.DnsSubdomainName)
	}

	if s.Res.EmailDomainId != nil {
		s.D.Set("email_domain_id", *s.Res.EmailDomainId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TxtRecordValue != nil {
		s.D.Set("txt_record_value", *s.Res.TxtRecordValue)
	}

	return nil
}
