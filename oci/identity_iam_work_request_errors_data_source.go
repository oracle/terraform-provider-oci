// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v51/identity"
)

func init() {
	RegisterDatasource("oci_identity_iam_work_request_errors", IdentityIamWorkRequestErrorsDataSource())
}

func IdentityIamWorkRequestErrorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityIamWorkRequestErrors,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"iam_work_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"iam_work_request_errors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timestamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readIdentityIamWorkRequestErrors(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIamWorkRequestErrorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityIamWorkRequestErrorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListIamWorkRequestErrorsResponse
}

func (s *IdentityIamWorkRequestErrorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityIamWorkRequestErrorsDataSourceCrud) Get() error {
	request := oci_identity.ListIamWorkRequestErrorsRequest{}

	if iamWorkRequestId, ok := s.D.GetOkExists("iam_work_request_id"); ok {
		tmp := iamWorkRequestId.(string)
		request.IamWorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "identity")

	response, err := s.Client.ListIamWorkRequestErrors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIamWorkRequestErrors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityIamWorkRequestErrorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("IdentityIamWorkRequestErrorsDataSource-", IdentityIamWorkRequestErrorsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		iamWorkRequestError := map[string]interface{}{}

		if r.Code != nil {
			iamWorkRequestError["code"] = *r.Code
		}

		if r.Message != nil {
			iamWorkRequestError["message"] = *r.Message
		}

		if r.Timestamp != nil {
			iamWorkRequestError["timestamp"] = r.Timestamp.String()
		}

		resources = append(resources, iamWorkRequestError)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityIamWorkRequestErrorsDataSource().Schema["iam_work_request_errors"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("iam_work_request_errors", resources); err != nil {
		return err
	}

	return nil
}
