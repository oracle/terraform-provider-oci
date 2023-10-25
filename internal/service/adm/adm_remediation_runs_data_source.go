// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AdmRemediationRunsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAdmRemediationRuns,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remediation_recipe_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remediation_run_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(AdmRemediationRunResource()),
						},
					},
				},
			},
		},
	}
}

func readAdmRemediationRuns(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRunsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.ReadResource(sync)
}

type AdmRemediationRunsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_adm.ApplicationDependencyManagementClient
	Res    *oci_adm.ListRemediationRunsResponse
}

func (s *AdmRemediationRunsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AdmRemediationRunsDataSourceCrud) Get() error {
	request := oci_adm.ListRemediationRunsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if remediationRecipeId, ok := s.D.GetOkExists("remediation_recipe_id"); ok {
		tmp := remediationRecipeId.(string)
		request.RemediationRecipeId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_adm.RemediationRunLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "adm")

	response, err := s.Client.ListRemediationRuns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRemediationRuns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AdmRemediationRunsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AdmRemediationRunsDataSource-", AdmRemediationRunsDataSource(), s.D))
	resources := []map[string]interface{}{}
	remediationRun := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RemediationRunSummaryToMap(item))
	}
	remediationRun["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AdmRemediationRunsDataSource().Schema["remediation_run_collection"].Elem.(*schema.Resource).Schema)
		remediationRun["items"] = items
	}

	resources = append(resources, remediationRun)
	if err := s.D.Set("remediation_run_collection", resources); err != nil {
		return err
	}

	return nil
}
