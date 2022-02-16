// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityIamWorkRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityIamWorkRequests,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"iam_work_requests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"operation_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"percent_complete": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"resources": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"action_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_accepted": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_finished": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readIdentityIamWorkRequests(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIamWorkRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityIamWorkRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListIamWorkRequestsResponse
}

func (s *IdentityIamWorkRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityIamWorkRequestsDataSourceCrud) Get() error {
	request := oci_identity.ListIamWorkRequestsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if resourceIdentifier, ok := s.D.GetOkExists("resource_identifier"); ok {
		tmp := resourceIdentifier.(string)
		request.ResourceIdentifier = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListIamWorkRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIamWorkRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityIamWorkRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityIamWorkRequestsDataSource-", IdentityIamWorkRequestsDataSource(), s.D))
	iamWorkRequestItems := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		iamWorkRequest := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Id != nil {
			iamWorkRequest["id"] = *r.Id
		}

		iamWorkRequest["operation_type"] = r.OperationType

		if r.PercentComplete != nil {
			iamWorkRequest["percent_complete"] = *r.PercentComplete
		}

		resources := []interface{}{}
		for _, item := range r.Resources {
			resources = append(resources, IamWorkRequestResourceToMap(item))
		}
		iamWorkRequest["resources"] = resources

		iamWorkRequest["status"] = r.Status

		if r.TimeAccepted != nil {
			iamWorkRequest["time_accepted"] = r.TimeAccepted.String()
		}

		if r.TimeFinished != nil {
			iamWorkRequest["time_finished"] = r.TimeFinished.String()
		}

		if r.TimeStarted != nil {
			iamWorkRequest["time_started"] = r.TimeStarted.String()
		}

		iamWorkRequestItems = append(iamWorkRequestItems, iamWorkRequest)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		iamWorkRequestItems = tfresource.ApplyFilters(f.(*schema.Set), iamWorkRequestItems, IdentityIamWorkRequestsDataSource().Schema["iam_work_requests"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("iam_work_requests", iamWorkRequestItems); err != nil {
		return err
	}

	return nil
}
