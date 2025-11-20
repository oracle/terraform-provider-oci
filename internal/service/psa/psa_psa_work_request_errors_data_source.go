// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psa "github.com/oracle/oci-go-sdk/v65/psa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsaWorkRequestErrorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readPsaWorkRequestErrors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"work_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"work_request_errors": {
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

func readPsaWorkRequestErrors(d *schema.ResourceData, m interface{}) error {
	sync := &PsaWorkRequestErrorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivateServiceAccessClient()

	return tfresource.ReadResource(sync)
}

type PsaWorkRequestErrorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psa.PrivateServiceAccessClient
	Res    *oci_psa.ListPsaWorkRequestErrorsResponse
}

func (s *PsaWorkRequestErrorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsaWorkRequestErrorsDataSourceCrud) Get() error {
	request := oci_psa.ListPsaWorkRequestErrorsRequest{}

	if workRequestId, ok := s.D.GetOkExists("work_request_id"); ok {
		tmp := workRequestId.(string)
		request.WorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psa")

	response, err := s.Client.ListPsaWorkRequestErrors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PsaWorkRequestErrorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsaWorkRequestErrorsDataSource-", PsaWorkRequestErrorsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		workRequestError := map[string]interface{}{}

		if r.Code != nil {
			workRequestError["code"] = *r.Code
		}

		if r.Message != nil {
			workRequestError["message"] = *r.Message
		}

		if r.Timestamp != nil {
			workRequestError["timestamp"] = r.Timestamp.String()
		}

		resources = append(resources, workRequestError)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, PsaWorkRequestErrorsDataSource().Schema["work_request_errors"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("work_request_errors", resources); err != nil {
		return err
	}

	return nil
}
