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

func OspGatewayAddressActionVerificationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOspGatewayAddressActionVerification,
		Read:     readOspGatewayAddressActionVerification,
		Delete:   deleteOspGatewayAddressActionVerification,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"osp_home_region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"address_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"city": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"company_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"contributor_class": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"country": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"county": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"department_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"internal_number": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"job_title": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"line1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"line2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"line3": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"line4": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"middle_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"municipal_inscription": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"phone_country_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"postal_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"province": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"state_inscription": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"street_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"street_number": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"address": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

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
				},
			},
			"quality": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"verification_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOspGatewayAddressActionVerification(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewayAddressActionVerificationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AddressServiceClient()

	return tfresource.CreateResource(d, sync)
}

func readOspGatewayAddressActionVerification(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOspGatewayAddressActionVerification(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OspGatewayAddressActionVerificationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_osp_gateway.AddressServiceClient
	Res                    *oci_osp_gateway.VerifyAddressReceipt
	DisableNotFoundRetries bool
}

func (s *OspGatewayAddressActionVerificationResourceCrud) ID() string {
	return "null"
}

func (s *OspGatewayAddressActionVerificationResourceCrud) Create() error {
	request := oci_osp_gateway.VerifyAddressRequest{}

	if addressKey, ok := s.D.GetOkExists("address_key"); ok {
		tmp := addressKey.(string)
		request.AddressKey = &tmp
	}

	if city, ok := s.D.GetOkExists("city"); ok {
		tmp := city.(string)
		request.City = &tmp
	}

	if companyName, ok := s.D.GetOkExists("company_name"); ok {
		tmp := companyName.(string)
		request.CompanyName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if contributorClass, ok := s.D.GetOkExists("contributor_class"); ok {
		tmp := contributorClass.(string)
		request.ContributorClass = &tmp
	}

	if country, ok := s.D.GetOkExists("country"); ok {
		tmp := country.(string)
		request.Country = &tmp
	}

	if county, ok := s.D.GetOkExists("county"); ok {
		tmp := county.(string)
		request.County = &tmp
	}

	if departmentName, ok := s.D.GetOkExists("department_name"); ok {
		tmp := departmentName.(string)
		request.DepartmentName = &tmp
	}

	if emailAddress, ok := s.D.GetOkExists("email_address"); ok {
		tmp := emailAddress.(string)
		request.EmailAddress = &tmp
	}

	if firstName, ok := s.D.GetOkExists("first_name"); ok {
		tmp := firstName.(string)
		request.FirstName = &tmp
	}

	if internalNumber, ok := s.D.GetOkExists("internal_number"); ok {
		tmp := internalNumber.(string)
		request.InternalNumber = &tmp
	}

	if jobTitle, ok := s.D.GetOkExists("job_title"); ok {
		tmp := jobTitle.(string)
		request.JobTitle = &tmp
	}

	if lastName, ok := s.D.GetOkExists("last_name"); ok {
		tmp := lastName.(string)
		request.LastName = &tmp
	}

	if line1, ok := s.D.GetOkExists("line1"); ok {
		tmp := line1.(string)
		request.Line1 = &tmp
	}

	if line2, ok := s.D.GetOkExists("line2"); ok {
		tmp := line2.(string)
		request.Line2 = &tmp
	}

	if line3, ok := s.D.GetOkExists("line3"); ok {
		tmp := line3.(string)
		request.Line3 = &tmp
	}

	if line4, ok := s.D.GetOkExists("line4"); ok {
		tmp := line4.(string)
		request.Line4 = &tmp
	}

	if middleName, ok := s.D.GetOkExists("middle_name"); ok {
		tmp := middleName.(string)
		request.MiddleName = &tmp
	}

	if municipalInscription, ok := s.D.GetOkExists("municipal_inscription"); ok {
		tmp := municipalInscription.(string)
		request.MunicipalInscription = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	if phoneCountryCode, ok := s.D.GetOkExists("phone_country_code"); ok {
		tmp := phoneCountryCode.(string)
		request.PhoneCountryCode = &tmp
	}

	if phoneNumber, ok := s.D.GetOkExists("phone_number"); ok {
		tmp := phoneNumber.(string)
		request.PhoneNumber = &tmp
	}

	if postalCode, ok := s.D.GetOkExists("postal_code"); ok {
		tmp := postalCode.(string)
		request.PostalCode = &tmp
	}

	if province, ok := s.D.GetOkExists("province"); ok {
		tmp := province.(string)
		request.Province = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		tmp := state.(string)
		request.State = &tmp
	}

	if stateInscription, ok := s.D.GetOkExists("state_inscription"); ok {
		tmp := stateInscription.(string)
		request.StateInscription = &tmp
	}

	if streetName, ok := s.D.GetOkExists("street_name"); ok {
		tmp := streetName.(string)
		request.StreetName = &tmp
	}

	if streetNumber, ok := s.D.GetOkExists("street_number"); ok {
		tmp := streetNumber.(string)
		request.StreetNumber = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osp_gateway")

	response, err := s.Client.VerifyAddress(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VerifyAddressReceipt
	return nil
}

func (s *OspGatewayAddressActionVerificationResourceCrud) SetData() error {
	if s.Res.Address != nil {
		s.D.Set("address", []interface{}{AddressToMap(s.Res.Address)})
	} else {
		s.D.Set("address", nil)
	}

	s.D.Set("quality", s.Res.Quality)

	s.D.Set("verification_code", s.Res.VerificationCode)

	return nil
}
