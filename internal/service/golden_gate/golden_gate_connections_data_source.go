// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateConnections,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"assignable_deployment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assignable_deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assigned_deployment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"connection_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"technology_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connection_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GoldenGateConnectionResource()),
						},
					},
				},
			},
		},
	}
}

func readGoldenGateConnections(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateConnectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateConnectionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListConnectionsResponse
}

func (s *GoldenGateConnectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateConnectionsDataSourceCrud) Get() error {
	request := oci_golden_gate.ListConnectionsRequest{}

	if assignableDeploymentId, ok := s.D.GetOkExists("assignable_deployment_id"); ok {
		tmp := assignableDeploymentId.(string)
		request.AssignableDeploymentId = &tmp
	}

	if assignableDeploymentType, ok := s.D.GetOkExists("assignable_deployment_type"); ok {
		request.AssignableDeploymentType = oci_golden_gate.ListConnectionsAssignableDeploymentTypeEnum(assignableDeploymentType.(string))
	}

	if assignedDeploymentId, ok := s.D.GetOkExists("assigned_deployment_id"); ok {
		tmp := assignedDeploymentId.(string)
		request.AssignedDeploymentId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_golden_gate.ConnectionLifecycleStateEnum(state.(string))
	}

	if connectionType, ok := s.D.GetOkExists("connection_type"); ok {
		interfaces := connectionType.([]interface{})
		tmp := make([]oci_golden_gate.ConnectionTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_golden_gate.ConnectionTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("connection_type") {
			request.ConnectionType = tmp
		}
	}

	if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
		interfaces := technologyType.([]interface{})
		tmp := make([]oci_golden_gate.TechnologyTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_golden_gate.TechnologyTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("technology_type") {
			request.TechnologyType = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListConnections(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConnections(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGateConnectionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateConnectionsDataSource-", GoldenGateConnectionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	connection := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ConnectionSummaryToMap(item, true))
	}
	connection["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateConnectionsDataSource().Schema["connection_collection"].Elem.(*schema.Resource).Schema)
		connection["items"] = items
	}

	resources = append(resources, connection)
	if err := s.D.Set("connection_collection", resources); err != nil {
		return err
	}

	return nil
}
