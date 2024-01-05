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

func ApigatewayApiDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["api_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApigatewayApiResource(), fieldMap, readSingularApigatewayApi)
}

func readSingularApigatewayApi(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayApiDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiGatewayClient()

	return tfresource.ReadResource(sync)
}

type ApigatewayApiDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.ApiGatewayClient
	Res    *oci_apigateway.GetApiResponse
}

func (s *ApigatewayApiDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewayApiDataSourceCrud) Get() error {
	request := oci_apigateway.GetApiRequest{}

	if apiId, ok := s.D.GetOkExists("api_id"); ok {
		tmp := apiId.(string)
		request.ApiId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apigateway")

	response, err := s.Client.GetApi(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApigatewayApiDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	if s.Res.SpecificationType != nil {
		s.D.Set("specification_type", *s.Res.SpecificationType)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	validationResults := []interface{}{}
	for _, item := range s.Res.ValidationResults {
		validationResults = append(validationResults, ApiValidationResultToMap(item))
	}
	s.D.Set("validation_results", validationResults)

	return nil
}
