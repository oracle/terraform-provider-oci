// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityIamWorkRequestLogsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityIamWorkRequestLogs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"iam_work_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"iam_work_request_logs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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

func readIdentityIamWorkRequestLogs(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIamWorkRequestLogsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityIamWorkRequestLogsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListIamWorkRequestLogsResponse
}

func (s *IdentityIamWorkRequestLogsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityIamWorkRequestLogsDataSourceCrud) Get() error {
	request := oci_identity.ListIamWorkRequestLogsRequest{}

	if iamWorkRequestId, ok := s.D.GetOkExists("iam_work_request_id"); ok {
		tmp := iamWorkRequestId.(string)
		request.IamWorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListIamWorkRequestLogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIamWorkRequestLogs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityIamWorkRequestLogsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityIamWorkRequestLogsDataSource-", IdentityIamWorkRequestLogsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		iamWorkRequestLog := map[string]interface{}{}

		if r.Message != nil {
			iamWorkRequestLog["message"] = *r.Message
		}

		if r.Timestamp != nil {
			iamWorkRequestLog["timestamp"] = r.Timestamp.String()
		}

		resources = append(resources, iamWorkRequestLog)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityIamWorkRequestLogsDataSource().Schema["iam_work_request_logs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("iam_work_request_logs", resources); err != nil {
		return err
	}

	return nil
}
