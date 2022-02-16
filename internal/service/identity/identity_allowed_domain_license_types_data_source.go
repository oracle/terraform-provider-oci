// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"
)

func IdentityAllowedDomainLicenseTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityAllowedDomainLicenseTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"current_license_type_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowed_domain_license_types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readIdentityAllowedDomainLicenseTypes(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAllowedDomainLicenseTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityAllowedDomainLicenseTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListAllowedDomainLicenseTypesResponse
}

func (s *IdentityAllowedDomainLicenseTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityAllowedDomainLicenseTypesDataSourceCrud) Get() error {
	request := oci_identity.ListAllowedDomainLicenseTypesRequest{}

	if currentLicenseTypeName, ok := s.D.GetOkExists("current_license_type_name"); ok {
		tmp := currentLicenseTypeName.(string)
		request.CurrentLicenseTypeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListAllowedDomainLicenseTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityAllowedDomainLicenseTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityAllowedDomainLicenseTypesDataSource-", IdentityAllowedDomainLicenseTypesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		allowedDomainLicenseType := map[string]interface{}{}

		if r.Description != nil {
			allowedDomainLicenseType["description"] = *r.Description
		}

		if r.LicenseType != nil {
			allowedDomainLicenseType["license_type"] = *r.LicenseType
		}

		if r.Name != nil {
			allowedDomainLicenseType["name"] = *r.Name
		}

		resources = append(resources, allowedDomainLicenseType)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityAllowedDomainLicenseTypesDataSource().Schema["allowed_domain_license_types"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("allowed_domain_license_types", resources); err != nil {
		return err
	}

	return nil
}
