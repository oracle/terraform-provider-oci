// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v27/mysql"
)

func init() {
	RegisterDatasource("oci_mysql_analytics_cluster", MysqlAnalyticsClusterDataSource())
}

func MysqlAnalyticsClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["db_system_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(MysqlAnalyticsClusterResource(), fieldMap, readSingularMysqlAnalyticsCluster)
}

func readSingularMysqlAnalyticsCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlAnalyticsClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()

	return ReadResource(sync)
}

type MysqlAnalyticsClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.DbSystemClient
	Res    *oci_mysql.GetAnalyticsClusterResponse
}

func (s *MysqlAnalyticsClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlAnalyticsClusterDataSourceCrud) Get() error {
	request := oci_mysql.GetAnalyticsClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "mysql")

	response, err := s.Client.GetAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlAnalyticsClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	clusterNodes := []interface{}{}
	for _, item := range s.Res.ClusterNodes {
		clusterNodes = append(clusterNodes, AnalyticsClusterNodeToMap(item))
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
