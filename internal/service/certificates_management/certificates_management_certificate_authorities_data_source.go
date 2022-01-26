// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v56/certificatesmanagement"
)

func CertificatesManagementCertificateAuthoritiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCertificatesManagementCertificateAuthorities,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"certificate_authority_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name", "compartment_id", "issuer_certificate_authority_id"},
			},
			"compartment_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"certificate_authority_id"},
			},
			"issuer_certificate_authority_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"certificate_authority_id", "name"},
				RequiredWith:  []string{"compartment_id"},
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"certificate_authority_id", "issuer_certificate_authority_id"},
				RequiredWith:  []string{"compartment_id"},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_authority_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CertificatesManagementCertificateAuthorityResource()),
						},
					},
				},
			},
		},
	}
}

func readCertificatesManagementCertificateAuthorities(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateAuthoritiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

type CertificatesManagementCertificateAuthoritiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.ListCertificateAuthoritiesResponse
}

func (s *CertificatesManagementCertificateAuthoritiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementCertificateAuthoritiesDataSourceCrud) Get() error {
	request := oci_certificates_management.ListCertificateAuthoritiesRequest{}

	if certificateAuthorityId, ok := s.D.GetOkExists("certificate_authority_id"); ok {
		tmp := certificateAuthorityId.(string)
		request.CertificateAuthorityId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if issuerCertificateAuthorityId, ok := s.D.GetOkExists("issuer_certificate_authority_id"); ok {
		tmp := issuerCertificateAuthorityId.(string)
		request.IssuerCertificateAuthorityId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_certificates_management.ListCertificateAuthoritiesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.ListCertificateAuthorities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCertificateAuthorities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CertificatesManagementCertificateAuthoritiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CertificatesManagementCertificateAuthoritiesDataSource-", CertificatesManagementCertificateAuthoritiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	certificateAuthority := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CertificateAuthoritySummaryToMap(item))
	}
	certificateAuthority["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CertificatesManagementCertificateAuthoritiesDataSource().Schema["certificate_authority_collection"].Elem.(*schema.Resource).Schema)
		certificateAuthority["items"] = items
	}

	resources = append(resources, certificateAuthority)
	if err := s.D.Set("certificate_authority_collection", resources); err != nil {
		return err
	}

	return nil
}
