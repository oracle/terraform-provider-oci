// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"
)

func DnsTsigKeyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["tsig_key_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DnsTsigKeyResource(), fieldMap, readSingularDnsTsigKey)
}

func readSingularDnsTsigKey(d *schema.ResourceData, m interface{}) error {
	sync := &DnsTsigKeyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsTsigKeyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.GetTsigKeyResponse
}

func (s *DnsTsigKeyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsTsigKeyDataSourceCrud) Get() error {
	request := oci_dns.GetTsigKeyRequest{}

	if tsigKeyId, ok := s.D.GetOkExists("tsig_key_id"); ok {
		tmp := tsigKeyId.(string)
		request.TsigKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.GetTsigKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DnsTsigKeyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Algorithm != nil {
		s.D.Set("algorithm", *s.Res.Algorithm)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Secret != nil {
		s.D.Set("secret", *s.Res.Secret)
	}

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
