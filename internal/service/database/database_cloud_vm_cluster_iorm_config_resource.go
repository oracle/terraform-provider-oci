package database

import (
	"context"
	"fmt"

	"terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/tfresource"

	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseCloudVmClusterIormConfigResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseCloudVmClusterIormConfig,
		Read:     readDatabaseCloudVmClusterIormConfig,
		Update:   updateDatabaseCloudVmClusterIormConfig,
		Delete:   deleteDatabaseCloudVmClusterIormConfig,
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
							ValidateFunc: tfresource.ValidateNotEmptyString(),
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
			"cloud_vm_cluster_id": {
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

func createDatabaseCloudVmClusterIormConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudVmClusterIormConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseCloudVmClusterIormConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudVmClusterIormConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func updateDatabaseCloudVmClusterIormConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudVmClusterIormConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseCloudVmClusterIormConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudVmClusterIormConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseCloudVmClusterIormConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.ExadataIormConfig
	DisableNotFoundRetries bool
}

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) ID() string {
	return s.D.Get("cloud_vm_cluster_id").(string)
}

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExadataIormConfigLifecycleStateBootstrapping),
		string(oci_database.ExadataIormConfigLifecycleStateUpdating),
	}
}

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExadataIormConfigLifecycleStateEnabled),
	}
}

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExadataIormConfigLifecycleStateUpdating),
	}
}

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExadataIormConfigLifecycleStateDisabled),
	}
}

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) Create() error {
	request := oci_database.GetCloudVmClusterIormConfigRequest{}

	if cloudVmClusterId, ok := s.D.GetOkExists("cloud_vm_cluster_id"); ok {
		tmp := cloudVmClusterId.(string)
		request.CloudVmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	if _, err := s.Client.GetCloudVmClusterIormConfig(context.Background(), request); err != nil {
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

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) Get() error {
	request := oci_database.GetCloudVmClusterIormConfigRequest{}

	if cloudVmClusterId, ok := s.D.GetOkExists("cloud_vm_cluster_id"); ok {
		tmp := cloudVmClusterId.(string)
		request.CloudVmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetCloudVmClusterIormConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataIormConfig
	return nil
}

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) Update() error {
	request := oci_database.UpdateCloudVmClusterIormConfigRequest{}

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
			request.CloudVmClusterIormConfigUpdateDetails.DbPlans = tmp
		}
	}

	if cloudVmClusterId, ok := s.D.GetOkExists("cloud_vm_cluster_id"); ok {
		tmp := cloudVmClusterId.(string)
		request.CloudVmClusterId = &tmp
	}

	if objective, ok := s.D.GetOkExists("objective"); ok {
		request.CloudVmClusterIormConfigUpdateDetails.Objective = oci_database.ExadataIormConfigUpdateDetailsObjectiveEnum(objective.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateCloudVmClusterIormConfig(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudVmCluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_database.ExadataIormConfigLifecycleStateEnabled }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) SetData() error {

	s.D.SetId(s.D.Get("cloud_vm_cluster_id").(string))
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

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) mapTodbIormConfigUpdateDetail(fieldKeyFormat string) (oci_database.DbIormConfigUpdateDetail, error) {
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

func (s *DatabaseCloudVmClusterIormConfigResourceCrud) Delete() error {
	request := oci_database.UpdateCloudVmClusterIormConfigRequest{}

	if cloudVmClusterId, ok := s.D.GetOkExists("cloud_vm_cluster_id"); ok {
		tmp := cloudVmClusterId.(string)
		request.CloudVmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateCloudVmClusterIormConfig(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudVmCluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.ExadataIormConfig
	return nil
}
