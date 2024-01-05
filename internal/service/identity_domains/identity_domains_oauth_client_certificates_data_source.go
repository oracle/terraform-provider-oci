// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsOAuthClientCertificatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomainsOAuthClientCertificates,
		Schema: map[string]*schema.Schema{
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oauth_client_certificate_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oauth_client_certificate_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Computed
			"oauth_client_certificates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityDomainsOAuthClientCertificateResource()),
			},
			"items_per_page": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"total_results": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readIdentityDomainsOAuthClientCertificates(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsOAuthClientCertificatesDataSourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

type IdentityDomainsOAuthClientCertificatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.ListOAuthClientCertificatesResponse
}

func (s *IdentityDomainsOAuthClientCertificatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsOAuthClientCertificatesDataSourceCrud) Get() error {
	request := oci_identity_domains.ListOAuthClientCertificatesRequest{}

	if oAuthClientCertificateCount, ok := s.D.GetOkExists("oauth_client_certificate_count"); ok {
		tmp := oAuthClientCertificateCount.(int)
		request.Count = &tmp
	}

	if oAuthClientCertificateFilter, ok := s.D.GetOkExists("oauth_client_certificate_filter"); ok {
		tmp := oAuthClientCertificateFilter.(string)
		request.Filter = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if sortBy, ok := s.D.GetOkExists("sort_by"); ok {
		tmp := sortBy.(string)
		request.SortBy = &tmp
	}

	if sortOrder, ok := s.D.GetOkExists("sort_order"); ok {
		tmp := oci_identity_domains.ListOAuthClientCertificatesSortOrderEnum(sortOrder.(string))
		request.SortOrder = tmp
	}

	if startIndex, ok := s.D.GetOkExists("start_index"); ok {
		tmp := startIndex.(int)
		request.StartIndex = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.ListOAuthClientCertificates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	// IDCS pagination
	startIndex := *response.StartIndex
	for startIndex+*response.ItemsPerPage <= *response.TotalResults {
		startIndex += *response.ItemsPerPage
		request.StartIndex = &startIndex
		listResponse, err := s.Client.ListOAuthClientCertificates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Resources = append(s.Res.Resources, listResponse.Resources...)
	}

	return nil
}

func (s *IdentityDomainsOAuthClientCertificatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsOAuthClientCertificatesDataSource-", IdentityDomainsOAuthClientCertificatesDataSource(), s.D))

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, OAuthClientCertificateToMap(item))
	}
	s.D.Set("oauth_client_certificates", resources)

	if s.Res.ItemsPerPage != nil {
		s.D.Set("items_per_page", *s.Res.ItemsPerPage)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.StartIndex != nil {
		s.D.Set("start_index", *s.Res.StartIndex)
	}

	if s.Res.TotalResults != nil {
		s.D.Set("total_results", *s.Res.TotalResults)
	}

	return nil
}
