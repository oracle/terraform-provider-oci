// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v56/mysql"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func MysqlHeatWaveClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["db_system_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MysqlHeatWaveClusterResource(), fieldMap, readSingularMysqlHeatWaveCluster)
}

func readSingularMysqlHeatWaveCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlHeatWaveClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	return tfresource.ReadResource(sync)
}

type MysqlHeatWaveClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.DbSystemClient
	Res    *oci_mysql.GetHeatWaveClusterResponse
}

func (s *MysqlHeatWaveClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlHeatWaveClusterDataSourceCrud) Get() error {
	request := oci_mysql.GetHeatWaveClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.GetHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlHeatWaveClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MysqlHeatWaveClusterDataSource-", MysqlHeatWaveClusterDataSource(), s.D))

	clusterNodes := []interface{}{}
	for _, item := range s.Res.ClusterNodes {
		clusterNodes = append(clusterNodes, HeatWaveNodeToMap(item))
	}
	s.D.Set("cluster_nodes", clusterNodes)

	if s.Res.ClusterSize != nil {
		s.D.Set("cluster_size", *s.Res.ClusterSize)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
