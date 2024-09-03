// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"
)

func GoldenGateDeploymentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateDeployments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"assignable_connection_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assigned_connection_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_sub_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"supported_connection_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deployment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GoldenGateDeploymentResource()),
						},
					},
				},
			},
		},
	}
}

func readGoldenGateDeployments(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateDeploymentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListDeploymentsResponse
}

func (s *GoldenGateDeploymentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentsDataSourceCrud) Get() error {
	request := oci_golden_gate.ListDeploymentsRequest{}

	if assignableConnectionId, ok := s.D.GetOkExists("assignable_connection_id"); ok {
		tmp := assignableConnectionId.(string)
		request.AssignableConnectionId = &tmp
	}

	if assignedConnectionId, ok := s.D.GetOkExists("assigned_connection_id"); ok {
		tmp := assignedConnectionId.(string)
		request.AssignedConnectionId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fqdn, ok := s.D.GetOkExists("fqdn"); ok {
		tmp := fqdn.(string)
		request.Fqdn = &tmp
	}

	if lifecycleSubState, ok := s.D.GetOkExists("lifecycle_sub_state"); ok {
		request.LifecycleSubState = oci_golden_gate.ListDeploymentsLifecycleSubStateEnum(lifecycleSubState.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_golden_gate.ListDeploymentsLifecycleStateEnum(state.(string))
	}

	if supportedConnectionType, ok := s.D.GetOkExists("supported_connection_type"); ok {
		request.SupportedConnectionType = oci_golden_gate.ListDeploymentsSupportedConnectionTypeEnum(supportedConnectionType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListDeployments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDeployments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGateDeploymentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateDeploymentsDataSource-", GoldenGateDeploymentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	deployment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DeploymentSummaryToMap(item))
	}
	deployment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateDeploymentsDataSource().Schema["deployment_collection"].Elem.(*schema.Resource).Schema)
		deployment["items"] = items
	}

	resources = append(resources, deployment)
	if err := s.D.Set("deployment_collection", resources); err != nil {
		return err
	}

	return nil
}
