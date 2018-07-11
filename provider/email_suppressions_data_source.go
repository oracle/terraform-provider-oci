// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_email "github.com/oracle/oci-go-sdk/email"

	"time"

	"github.com/oracle/terraform-provider-oci/crud"
)

func SuppressionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSuppressions,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"suppressions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     SuppressionResource(),
			},
		},
	}
}

func readSuppressions(d *schema.ResourceData, m interface{}) error {
	sync := &SuppressionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return crud.ReadResource(sync)
}

type SuppressionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.ListSuppressionsResponse
}

func (s *SuppressionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SuppressionsDataSourceCrud) Get() error {
	request := oci_email.ListSuppressionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if emailAddress, ok := s.D.GetOkExists("email_address"); ok {
		tmp := emailAddress.(string)
		request.EmailAddress = &tmp
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "email")

	response, err := s.Client.ListSuppressions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSuppressions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *SuppressionsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		suppression := map[string]interface{}{}

		if r.EmailAddress != nil {
			suppression["email_address"] = *r.EmailAddress
		}

		if r.Id != nil {
			suppression["id"] = *r.Id
		}

		suppression["reason"] = r.Reason

		if r.TimeCreated != nil {
			suppression["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, suppression)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, SuppressionsDataSource().Schema["suppressions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("suppressions", resources); err != nil {
		panic(err)
	}

	return
}
