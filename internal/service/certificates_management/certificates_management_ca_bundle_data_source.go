// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v65/certificatesmanagement"
)

func CertificatesManagementCaBundleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ca_bundle_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CertificatesManagementCaBundleResource(), fieldMap, readSingularCertificatesManagementCaBundle)
}

func readSingularCertificatesManagementCaBundle(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCaBundleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

type CertificatesManagementCaBundleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.GetCaBundleResponse
}

func (s *CertificatesManagementCaBundleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementCaBundleDataSourceCrud) Get() error {
	request := oci_certificates_management.GetCaBundleRequest{}

	if caBundleId, ok := s.D.GetOkExists("ca_bundle_id"); ok {
		tmp := caBundleId.(string)
		request.CaBundleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.GetCaBundle(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CertificatesManagementCaBundleDataSourceCrud) SetData() error {
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
