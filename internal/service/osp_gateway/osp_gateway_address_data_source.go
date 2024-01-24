// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osp_gateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v65/ospgateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OspGatewayAddressDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOspGatewayAddress,
		Schema: map[string]*schema.Schema{
			"address_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"osp_home_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"address_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"city": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"company_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"contributor_class": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"country": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"county": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"department_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"internal_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"job_title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"line1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"line2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"line3": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"line4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"middle_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"municipal_inscription": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_country_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"postal_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"province": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_inscription": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"street_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"street_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularOspGatewayAddress(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewayAddressDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AddressServiceClient()

	return tfresource.ReadResource(sync)
}

type OspGatewayAddressDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osp_gateway.AddressServiceClient
	Res    *oci_osp_gateway.GetAddressResponse
}

func (s *OspGatewayAddressDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OspGatewayAddressDataSourceCrud) Get() error {
	request := oci_osp_gateway.GetAddressRequest{}

	if addressId, ok := s.D.GetOkExists("address_id"); ok {
		tmp := addressId.(string)
		request.AddressId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osp_gateway")

	response, err := s.Client.GetAddress(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OspGatewayAddressDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OspGatewayAddressDataSource-", OspGatewayAddressDataSource(), s.D))

	if s.Res.AddressKey != nil {
		s.D.Set("address_key", *s.Res.AddressKey)
	}

	if s.Res.City != nil {
		s.D.Set("city", *s.Res.City)
	}

	if s.Res.CompanyName != nil {
		s.D.Set("company_name", *s.Res.CompanyName)
	}

	if s.Res.ContributorClass != nil {
		s.D.Set("contributor_class", *s.Res.ContributorClass)
	}

	if s.Res.Country != nil {
		s.D.Set("country", *s.Res.Country)
	}

	if s.Res.County != nil {
		s.D.Set("county", *s.Res.County)
	}

	if s.Res.DepartmentName != nil {
		s.D.Set("department_name", *s.Res.DepartmentName)
	}

	if s.Res.EmailAddress != nil {
		s.D.Set("email_address", *s.Res.EmailAddress)
	}

	if s.Res.FirstName != nil {
		s.D.Set("first_name", *s.Res.FirstName)
	}

	if s.Res.InternalNumber != nil {
		s.D.Set("internal_number", *s.Res.InternalNumber)
	}

	if s.Res.JobTitle != nil {
		s.D.Set("job_title", *s.Res.JobTitle)
	}

	if s.Res.LastName != nil {
		s.D.Set("last_name", *s.Res.LastName)
	}

	if s.Res.Line1 != nil {
		s.D.Set("line1", *s.Res.Line1)
	}

	if s.Res.Line2 != nil {
		s.D.Set("line2", *s.Res.Line2)
	}

	if s.Res.Line3 != nil {
		s.D.Set("line3", *s.Res.Line3)
	}

	if s.Res.Line4 != nil {
		s.D.Set("line4", *s.Res.Line4)
	}

	if s.Res.MiddleName != nil {
		s.D.Set("middle_name", *s.Res.MiddleName)
	}

	if s.Res.MunicipalInscription != nil {
		s.D.Set("municipal_inscription", *s.Res.MunicipalInscription)
	}

	if s.Res.PhoneCountryCode != nil {
		s.D.Set("phone_country_code", *s.Res.PhoneCountryCode)
	}

	if s.Res.PhoneNumber != nil {
		s.D.Set("phone_number", *s.Res.PhoneNumber)
	}

	if s.Res.PostalCode != nil {
		s.D.Set("postal_code", *s.Res.PostalCode)
	}

	if s.Res.Province != nil {
		s.D.Set("province", *s.Res.Province)
	}

	if s.Res.State != nil {
		s.D.Set("state", *s.Res.State)
	}

	if s.Res.StateInscription != nil {
		s.D.Set("state_inscription", *s.Res.StateInscription)
	}

	if s.Res.StreetName != nil {
		s.D.Set("street_name", *s.Res.StreetName)
	}

	if s.Res.StreetNumber != nil {
		s.D.Set("street_number", *s.Res.StreetNumber)
	}

	return nil
}
