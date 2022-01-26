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

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_mysql "github.com/oracle/oci-go-sdk/v56/mysql"
)

func MysqlHeatWaveClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createMysqlHeatWaveCluster,
		Read:   readMysqlHeatWaveCluster,
		Update: updateMysqlHeatWaveCluster,
		Delete: deleteMysqlHeatWaveCluster,
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
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_mysql.HeatWaveClusterLifecycleStateInactive),
					string(oci_mysql.HeatWaveClusterLifecycleStateActive),
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

func createMysqlHeatWaveCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlHeatWaveClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_mysql.HeatWaveClusterLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_mysql.HeatWaveClusterLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	// switch to power off
	if powerOff {
		if err := sync.Stop(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.HeatWaveClusterLifecycleStateInactive)
	}

	return nil
}

func readMysqlHeatWaveCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlHeatWaveClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	return tfresource.ReadResource(sync)
}

func updateMysqlHeatWaveCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlHeatWaveClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_mysql.HeatWaveClusterLifecycleStateActive == oci_mysql.HeatWaveClusterLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_mysql.HeatWaveClusterLifecycleStateInactive == oci_mysql.HeatWaveClusterLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	// switch to power on
	if powerOn {
		if err := sync.Start(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.HeatWaveClusterLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	// switch to power off
	if powerOff {
		if err := sync.Stop(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.HeatWaveClusterLifecycleStateInactive)
	}

	return nil
}

func deleteMysqlHeatWaveCluster(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlHeatWaveClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbSystemClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MysqlHeatWaveClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_mysql.DbSystemClient
	Res                    *oci_mysql.HeatWaveCluster
	DisableNotFoundRetries bool
}

func (s *MysqlHeatWaveClusterResourceCrud) ID() string {
	return getHeatWaveClusterCompositeId(s.D.Get("db_system_id").(string))
}

func (s *MysqlHeatWaveClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateCreating),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateActive),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateDeleting),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateDeleted),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateUpdating),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.HeatWaveClusterLifecycleStateActive),
	}
}

func (s *MysqlHeatWaveClusterResourceCrud) Create() error {
	request := oci_mysql.AddHeatWaveClusterRequest{}

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

	response, err := s.Client.AddHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.HeatWaveCluster
	return nil
}

func (s *MysqlHeatWaveClusterResourceCrud) Get() error {
	request := oci_mysql.GetHeatWaveClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	dbSystemId, err := parseHeatWaveClusterCompositeId(s.D.Id())
	if err == nil {
		request.DbSystemId = &dbSystemId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.HeatWaveCluster
	return nil
}

func (s *MysqlHeatWaveClusterResourceCrud) Update() error {
	request := oci_mysql.UpdateHeatWaveClusterRequest{}

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

	_, err := s.Client.UpdateHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *MysqlHeatWaveClusterResourceCrud) Delete() error {
	request := oci_mysql.DeleteHeatWaveClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteHeatWaveCluster(context.Background(), request)
	return err
}

func (s *MysqlHeatWaveClusterResourceCrud) SetData() error {

	dbSystemId, err := parseHeatWaveClusterCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("db_system_id", &dbSystemId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	clusterNodes := []interface{}{}
	for _, item := range s.Res.ClusterNodes {
		clusterNodes = append(clusterNodes, HeatWaveNodeToMap(item))
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

func getHeatWaveClusterCompositeId(dbSystemId string) string {
	dbSystemId = url.PathEscape(dbSystemId)
	compositeId := "dbSystem/" + dbSystemId + "/heatWaveCluster"
	return compositeId
}

func parseHeatWaveClusterCompositeId(compositeId string) (dbSystemId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("dbSystem/.*/heatWaveCluster", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	dbSystemId, _ = url.PathUnescape(parts[1])

	return
}

func HeatWaveNodeToMap(obj oci_mysql.HeatWaveNode) map[string]interface{} {
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

func (s *MysqlHeatWaveClusterResourceCrud) Stop() error {
	request := oci_mysql.StopHeatWaveClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StopHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.HeatWaveClusterLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *MysqlHeatWaveClusterResourceCrud) Start() error {
	request := oci_mysql.StartHeatWaveClusterRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StartHeatWaveCluster(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.HeatWaveClusterLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}
