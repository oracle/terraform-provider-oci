// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_mysql "github.com/oracle/oci-go-sdk/v56/mysql"
)

func MysqlAnalyticsClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createMysqlAnalyticsCluster,
		Read:   readMysqlAnalyticsCluster,
		Update: updateMysqlAnalyticsCluster,
		Delete: deleteMysqlAnalyticsCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_mysql.AnalyticsClusterLifecycleStateInactive),
					string(oci_mysql.AnalyticsClusterLifecycleStateActive),
				}, true),
			},

			// Computed
			"cluster_nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"node_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMysqlAnalyticsCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlAnalyticsClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_mysql.AnalyticsClusterLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_mysql.AnalyticsClusterLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopAnalyticsCluster(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.AnalyticsClusterLifecycleStateInactive)
	}
	return nil

}

func readMysqlAnalyticsCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlAnalyticsClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	return tfresource.ReadResource(sync)
}

func updateMysqlAnalyticsCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlAnalyticsClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_mysql.AnalyticsClusterLifecycleStateActive == oci_mysql.AnalyticsClusterLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_mysql.AnalyticsClusterLifecycleStateInactive == oci_mysql.AnalyticsClusterLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartAnalyticsCluster(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.AnalyticsClusterLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopAnalyticsCluster(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.AnalyticsClusterLifecycleStateInactive)
	}

	return nil
}

func deleteMysqlAnalyticsCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlAnalyticsClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MysqlAnalyticsClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_mysql.DbSystemClient
	Res                    *oci_mysql.AnalyticsCluster
	DisableNotFoundRetries bool
}

func (s *MysqlAnalyticsClusterResourceCrud) ID() string {
	return getAnalyticsClusterCompositeId(s.D.Get("db_system_id").(string))
}

func (s *MysqlAnalyticsClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateCreating),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateActive),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateDeleting),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateDeleted),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateUpdating),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.AnalyticsClusterLifecycleStateActive),
	}
}

func (s *MysqlAnalyticsClusterResourceCrud) Create() error {
	request := oci_mysql.AddAnalyticsClusterRequest{}

	if clusterSize, ok := s.D.GetOkExists("cluster_size"); ok {
		tmp := clusterSize.(int)
		request.ClusterSize = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.AddAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnalyticsCluster
	return nil
}

func (s *MysqlAnalyticsClusterResourceCrud) Get() error {
	request := oci_mysql.GetAnalyticsClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	dbSystemId, err := parseAnalyticsClusterCompositeId(s.D.Id())
	if err == nil {
		request.DbSystemId = &dbSystemId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnalyticsCluster
	return nil
}

func (s *MysqlAnalyticsClusterResourceCrud) Update() error {
	request := oci_mysql.UpdateAnalyticsClusterRequest{}

	if clusterSize, ok := s.D.GetOkExists("cluster_size"); ok {
		tmp := clusterSize.(int)
		request.ClusterSize = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.UpdateAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *MysqlAnalyticsClusterResourceCrud) Delete() error {
	request := oci_mysql.DeleteAnalyticsClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteAnalyticsCluster(context.Background(), request)
	return err
}

func (s *MysqlAnalyticsClusterResourceCrud) SetData() error {

	dbSystemId, err := parseAnalyticsClusterCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("db_system_id", &dbSystemId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	clusterNodes := []interface{}{}
	for _, item := range s.Res.ClusterNodes {
		clusterNodes = append(clusterNodes, AnalyticsClusterNodeToMap(item))
	}
	s.D.Set("cluster_nodes", clusterNodes)

	if s.Res.ClusterSize != nil {
		s.D.Set("cluster_size", *s.Res.ClusterSize)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
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

func getAnalyticsClusterCompositeId(dbSystemId string) string {
	dbSystemId = url.PathEscape(dbSystemId)
	compositeId := "dbSystem/" + dbSystemId + "/analyticsCluster"
	return compositeId
}

func parseAnalyticsClusterCompositeId(compositeId string) (dbSystemId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("dbSystem/.*/analyticsCluster", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	dbSystemId, _ = url.PathUnescape(parts[1])

	return
}

func (s *MysqlAnalyticsClusterResourceCrud) StartAnalyticsCluster() error {
	request := oci_mysql.StartAnalyticsClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StartAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.AnalyticsClusterLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *MysqlAnalyticsClusterResourceCrud) StopAnalyticsCluster() error {
	request := oci_mysql.StopAnalyticsClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StopAnalyticsCluster(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.AnalyticsClusterLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func AnalyticsClusterNodeToMap(obj oci_mysql.AnalyticsClusterNode) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NodeId != nil {
		result["node_id"] = string(*obj.NodeId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
