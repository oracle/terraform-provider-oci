// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApigatewayApiContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApigatewayApiContent,
		Schema: map[string]*schema.Schema{
			"api_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularApigatewayApiContent(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayApiContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiGatewayClient()

	return tfresource.ReadResource(sync)
}

type ApigatewayApiContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.ApiGatewayClient
	Res    *oci_apigateway.GetApiContentResponse
}

func (s *ApigatewayApiContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewayApiContentDataSourceCrud) Get() error {
	request := oci_apigateway.GetApiContentRequest{}

	if apiId, ok := s.D.GetOkExists("api_id"); ok {
		tmp := apiId.(string)
		request.ApiId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apigateway")

	response, err := s.Client.GetApiContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApigatewayApiContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApigatewayApiContentDataSource-", ApigatewayApiContentDataSource(), s.D))

	if s.Res.Content != nil {
		contentReader := s.Res.Content
		contentArray, err := ioutil.ReadAll(contentReader)
		if err != nil {
			log.Printf("Unable to read 'content' from response. Error: %q", err)
			return err
		}
		s.D.Set("content", string(contentArray))
	}
	return nil
}
