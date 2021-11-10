// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v51/identity"
)

func init() {
	RegisterDatasource("oci_identity_domain", IdentityDomainDataSource())
}

func IdentityDomainDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(IdentityDomainResource(), fieldMap, readSingularIdentityDomain)
}

func readSingularIdentityDomain(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityDomainDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetDomainResponse
}

func (s *IdentityDomainDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainDataSourceCrud) Get() error {
	request := oci_identity.GetDomainRequest{}

	if domainId, ok := s.D.GetOkExists("domain_id"); ok {
		tmp := domainId.(string)
		request.DomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "identity")

	response, err := s.Client.GetDomain(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HomeRegion != nil {
		s.D.Set("home_region", *s.Res.HomeRegion)
	}

	if s.Res.HomeRegionUrl != nil {
		s.D.Set("home_region_url", *s.Res.HomeRegionUrl)
	}

	if s.Res.IsHiddenOnLogin != nil {
		s.D.Set("is_hidden_on_login", *s.Res.IsHiddenOnLogin)
	}

	if s.Res.LicenseType != nil {
		s.D.Set("license_type", *s.Res.LicenseType)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	replicaRegions := []interface{}{}
	for _, item := range s.Res.ReplicaRegions {
		replicaRegions = append(replicaRegions, ReplicatedRegionDetailsToMap(item))
	}
	s.D.Set("replica_regions", replicaRegions)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.Url != nil {
		s.D.Set("url", *s.Res.Url)
	}

	return nil
}
