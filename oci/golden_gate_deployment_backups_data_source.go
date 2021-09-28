// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v48/goldengate"
)

func init() {
	RegisterDatasource("oci_golden_gate_deployment_backups", GoldenGateDeploymentBackupsDataSource())
}

func GoldenGateDeploymentBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateDeploymentBackups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deployment_backup_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(GoldenGateDeploymentBackupResource()),
						},
					},
				},
			},
		},
	}
}

func readGoldenGateDeploymentBackups(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).goldenGateClient()

	return ReadResource(sync)
}

type GoldenGateDeploymentBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListDeploymentBackupsResponse
}

func (s *GoldenGateDeploymentBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentBackupsDataSourceCrud) Get() error {
	request := oci_golden_gate.ListDeploymentBackupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_golden_gate.ListDeploymentBackupsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListDeploymentBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDeploymentBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGateDeploymentBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("GoldenGateDeploymentBackupsDataSource-", GoldenGateDeploymentBackupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	deploymentBackup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DeploymentBackupSummaryToMap(item))
	}
	deploymentBackup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateDeploymentBackupsDataSource().Schema["deployment_backup_collection"].Elem.(*schema.Resource).Schema)
		deploymentBackup["items"] = items
	}

	resources = append(resources, deploymentBackup)
	if err := s.D.Set("deployment_backup_collection", resources); err != nil {
		return err
	}

	return nil
}
