// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psa "github.com/oracle/oci-go-sdk/v65/psa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsaPrivateServiceAccessDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["private_service_access_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(PsaPrivateServiceAccessResource(), fieldMap, readSingularPsaPrivateServiceAccesWithContext)
}

func readSingularPsaPrivateServiceAccesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PsaPrivateServiceAccesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivateServiceAccessClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type PsaPrivateServiceAccesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psa.PrivateServiceAccessClient
	Res    *oci_psa.GetPrivateServiceAccessResponse
}

func (s *PsaPrivateServiceAccesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsaPrivateServiceAccesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_psa.GetPrivateServiceAccessRequest{}

	if privateServiceAccessId, ok := s.D.GetOkExists("private_service_access_id"); ok {
		tmp := privateServiceAccessId.(string)
		request.PrivateServiceAccessId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psa")

	response, err := s.Client.GetPrivateServiceAccess(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PsaPrivateServiceAccesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("fqdns", s.Res.Fqdns)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Ipv4Ip != nil {
		s.D.Set("ipv4ip", *s.Res.Ipv4Ip)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.SecurityAttributes != nil {
		s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))
	}

	if s.Res.ServiceId != nil {
		s.D.Set("service_id", *s.Res.ServiceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	if s.Res.VnicId != nil {
		s.D.Set("vnic_id", *s.Res.VnicId)
	}

	return nil
}
