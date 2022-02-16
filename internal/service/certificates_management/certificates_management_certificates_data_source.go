// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_certificates_management "github.com/oracle/oci-go-sdk/v58/certificatesmanagement"
)

func CertificatesManagementCertificatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCertificatesManagementCertificates,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"certificate_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name", "compartment_id", "issuer_certificate_authority_id"},
			},
			"compartment_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"certificate_id"},
			},
			"issuer_certificate_authority_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"certificate_id", "name"},
				RequiredWith:  []string{"compartment_id"},
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"certificate_id", "issuer_certificate_authority_id"},
				RequiredWith:  []string{"compartment_id"},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CertificatesManagementCertificateResource()),
						},
					},
				},
			},
		},
	}
}

func readCertificatesManagementCertificates(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

type CertificatesManagementCertificatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_certificates_management.CertificatesManagementClient
	Res    *oci_certificates_management.ListCertificatesResponse
}

func (s *CertificatesManagementCertificatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CertificatesManagementCertificatesDataSourceCrud) Get() error {
	request := oci_certificates_management.ListCertificatesRequest{}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
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
		request.LifecycleState = oci_certificates_management.ListCertificatesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "certificates_management")

	response, err := s.Client.ListCertificates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCertificates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CertificatesManagementCertificatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CertificatesManagementCertificatesDataSource-", CertificatesManagementCertificatesDataSource(), s.D))
	resources := []map[string]interface{}{}
	certificate := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CertificatesManagementCertificateSummaryToMap(item))
	}
	certificate["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CertificatesManagementCertificatesDataSource().Schema["certificate_collection"].Elem.(*schema.Resource).Schema)
		certificate["items"] = items
	}

	resources = append(resources, certificate)
	if err := s.D.Set("certificate_collection", resources); err != nil {
		return err
	}

	return nil
}
