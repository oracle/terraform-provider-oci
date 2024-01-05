// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"
)

func EmailSuppressionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readEmailSuppressions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
				Elem:     tfresource.GetDataSourceItemSchema(EmailSuppressionResource()),
			},
		},
	}
}

func readEmailSuppressions(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSuppressionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

type EmailSuppressionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.ListSuppressionsResponse
}

func (s *EmailSuppressionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailSuppressionsDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

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

func (s *EmailSuppressionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("EmailSuppressionsDataSource-", EmailSuppressionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		suppression := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, EmailSuppressionsDataSource().Schema["suppressions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("suppressions", resources); err != nil {
		return err
	}

	return nil
}
