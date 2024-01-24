// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApigatewaySubscriberDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["subscriber_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApigatewaySubscriberResource(), fieldMap, readSingularApigatewaySubscriber)
}

func readSingularApigatewaySubscriber(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewaySubscriberDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscribersClient()

	return tfresource.ReadResource(sync)
}

type ApigatewaySubscriberDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.SubscribersClient
	Res    *oci_apigateway.GetSubscriberResponse
}

func (s *ApigatewaySubscriberDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewaySubscriberDataSourceCrud) Get() error {
	request := oci_apigateway.GetSubscriberRequest{}

	if subscriberId, ok := s.D.GetOkExists("subscriber_id"); ok {
		tmp := subscriberId.(string)
		request.SubscriberId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apigateway")

	response, err := s.Client.GetSubscriber(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApigatewaySubscriberDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	clients := []interface{}{}
	for _, item := range s.Res.Clients {
		clients = append(clients, ClientToMap(item))
	}
	s.D.Set("clients", clients)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("usage_plans", s.Res.UsagePlans)

	return nil
}
