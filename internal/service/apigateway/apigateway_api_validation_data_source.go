// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/v56/apigateway"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ApigatewayApiValidationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApigatewayApiValidation,
		Schema: map[string]*schema.Schema{
			"api_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"validations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"msg": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"src": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": {
													Type:     schema.TypeList,
													Computed: true,
													Elem:     schema.TypeFloat,
												},
											},
										},
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"result": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularApigatewayApiValidation(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayApiValidationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiGatewayClient()

	return tfresource.ReadResource(sync)
}

type ApigatewayApiValidationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.ApiGatewayClient
	Res    *oci_apigateway.GetApiValidationsResponse
}

func (s *ApigatewayApiValidationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewayApiValidationDataSourceCrud) Get() error {
	request := oci_apigateway.GetApiValidationsRequest{}

	if apiId, ok := s.D.GetOkExists("api_id"); ok {
		tmp := apiId.(string)
		request.ApiId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apigateway")

	response, err := s.Client.GetApiValidations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApigatewayApiValidationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApigatewayApiValidationDataSource-", ApigatewayApiValidationDataSource(), s.D))

	validations := []interface{}{}
	for _, item := range s.Res.Validations {
		validations = append(validations, ApiValidationDetailsToMap(item))
	}
	s.D.Set("validations", validations)

	return nil
}

func ApiValidationDetailToMap(obj oci_apigateway.ApiValidationDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Msg != nil {
		result["msg"] = string(*obj.Msg)
	}

	result["severity"] = string(obj.Severity)

	src := [][]interface{}{}
	for _, items := range obj.Src {
		temp := []interface{}{}
		for _, item := range items {
			temp = append(temp, item)
		}
		src = append(src, temp)
	}
	result["src"] = src

	return result
}

func ApiValidationDetailsToMap(obj oci_apigateway.ApiValidationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	details := []interface{}{}
	for _, item := range obj.Details {
		details = append(details, ApiValidationDetailToMap(item))
	}
	result["details"] = details

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["result"] = string(obj.Result)

	return result
}
