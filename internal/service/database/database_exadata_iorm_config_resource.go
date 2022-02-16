// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"bytes"
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_work_requests "github.com/oracle/oci-go-sdk/v58/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseExadataIormConfigResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExadataIormConfig,
		Read:     readDatabaseExadataIormConfig,
		Update:   updateDatabaseExadataIormConfig,
		Delete:   deleteDatabaseExadataIormConfig,
		Schema: map[string]*schema.Schema{
			// Required
			"db_plans": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Set:      dbPlansHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"db_name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: utils.ValidateNotEmptyString(),
						},
						"share": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 32),
						},

						// Optional

						// Computed
						"flash_cache_limit": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"objective": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_database.ExadataIormConfigUpdateDetailsObjectiveLowLatency),
					string(oci_database.ExadataIormConfigUpdateDetailsObjectiveHighThroughput),
					string(oci_database.ExadataIormConfigUpdateDetailsObjectiveBalanced),
					string(oci_database.ExadataIormConfigUpdateDetailsObjectiveAuto),
				}, true),
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseExadataIormConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataIormConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseExadataIormConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataIormConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func updateDatabaseExadataIormConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataIormConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseExadataIormConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataIormConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExadataIormConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.ExadataIormConfig
	DisableNotFoundRetries bool
}

func (s *DatabaseExadataIormConfigResourceCrud) ID() string {
	return s.D.Get("db_system_id").(string)
}

func (s *DatabaseExadataIormConfigResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExadataIormConfigLifecycleStateBootstrapping),
		string(oci_database.ExadataIormConfigLifecycleStateUpdating),
	}
}

func (s *DatabaseExadataIormConfigResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExadataIormConfigLifecycleStateEnabled),
	}
}

func (s *DatabaseExadataIormConfigResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExadataIormConfigLifecycleStateUpdating),
	}
}

func (s *DatabaseExadataIormConfigResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExadataIormConfigLifecycleStateDisabled),
	}
}

func (s *DatabaseExadataIormConfigResourceCrud) Create() error {
	request := oci_database.GetExadataIormConfigRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	if _, err := s.Client.GetExadataIormConfig(context.Background(), request); err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res.LifecycleState == oci_database.ExadataIormConfigLifecycleStateDisabled || s.Res.LifecycleState == oci_database.ExadataIormConfigLifecycleStateEnabled
	}
	if err := tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate)); err != nil {
		return err
	}

	return s.Update()
}

func (s *DatabaseExadataIormConfigResourceCrud) Get() error {
	request := oci_database.GetExadataIormConfigRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExadataIormConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataIormConfig
	return nil
}

func (s *DatabaseExadataIormConfigResourceCrud) Update() error {
	request := oci_database.UpdateExadataIormConfigRequest{}

	if dbPlans, ok := s.D.GetOkExists("db_plans"); ok {
		set := dbPlans.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_database.DbIormConfigUpdateDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := dbPlansHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_plans", stateDataIndex)
			converted, err := s.mapTodbIormConfigUpdateDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("db_plans") {
			request.DbPlans = tmp
		}
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if objective, ok := s.D.GetOkExists("objective"); ok {
		request.Objective = oci_database.ExadataIormConfigUpdateDetailsObjectiveEnum(objective.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExadataIormConfig(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "dbSystem", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_database.ExadataIormConfigLifecycleStateEnabled }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseExadataIormConfigResourceCrud) SetData() error {

	s.D.SetId(s.D.Get("db_system_id").(string))

	dbPlans := []interface{}{}
	for _, item := range s.Res.DbPlans {
		if configMap := dbIormConfigToMap(item); configMap != nil {
			dbPlans = append(dbPlans, configMap)
		}
	}
	s.D.Set("db_plans", schema.NewSet(dbPlansHashCodeForSets, dbPlans))

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("objective", s.Res.Objective)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}

func (s *DatabaseExadataIormConfigResourceCrud) mapTodbIormConfigUpdateDetail(fieldKeyFormat string) (oci_database.DbIormConfigUpdateDetail, error) {
	result := oci_database.DbIormConfigUpdateDetail{}

	if dbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_name")); ok {
		tmp := dbName.(string)
		result.DbName = &tmp
	}

	if share, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "share")); ok {
		tmp := share.(int)
		result.Share = &tmp
	}

	return result, nil
}

func dbIormConfigToMap(obj oci_database.DbIormConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbName != nil {
		result["db_name"] = string(*obj.DbName)
	} else {
		return nil
	}

	if obj.FlashCacheLimit != nil {
		result["flash_cache_limit"] = string(*obj.FlashCacheLimit)
	} else {
		return nil
	}

	if obj.Share != nil {
		result["share"] = int(*obj.Share)
	} else {
		return nil
	}

	return result
}

func dbPlansHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if dbName, ok := m["db_name"]; ok && dbName != "" {
		buf.WriteString(fmt.Sprintf("%v-", dbName))
	}
	if share, ok := m["share"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", share))
	}
	return hashcode.String(buf.String())
}

func (s *DatabaseExadataIormConfigResourceCrud) Delete() error {
	request := oci_database.UpdateExadataIormConfigRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExadataIormConfig(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "dbSystem", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.ExadataIormConfig
	return nil
}
