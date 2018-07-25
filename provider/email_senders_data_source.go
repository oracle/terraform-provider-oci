// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/email"

	"github.com/oracle/terraform-provider-oci/crud"
)

func SendersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSenders,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"senders": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     crud.GetDataSourceItemSchema(SenderResource()),
			},
		},
	}
}

func readSenders(d *schema.ResourceData, m interface{}) error {
	sync := &SendersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return crud.ReadResource(sync)
}

type SendersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.ListSendersResponse
}

func (s *SendersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SendersDataSourceCrud) Get() error {
	request := oci_email.ListSendersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if emailAddress, ok := s.D.GetOkExists("email_address"); ok {
		tmp := emailAddress.(string)
		request.EmailAddress = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_email.SenderLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "email")

	response, err := s.Client.ListSenders(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSenders(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *SendersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		sender := map[string]interface{}{}

		if r.EmailAddress != nil {
			sender["email_address"] = *r.EmailAddress
		}

		if r.Id != nil {
			sender["id"] = *r.Id
		}

		sender["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			sender["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, sender)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, SendersDataSource().Schema["senders"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("senders", resources); err != nil {
		return err
	}

	return nil
}
