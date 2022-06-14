package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseCloudDatabaseManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseCloudDatabaseManagement,
		Update:   updateDatabaseCloudDatabaseManagement,
		Read:     readDatabaseCloudDatabaseManagement,
		Delete:   deleteDatabaseCloudDatabaseManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"management_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_end_point_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"credentialdetails": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"user_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"password_secret_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"enable_management": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_key_version_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_backup_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ncharacter_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pdb_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sid_prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_database_point_in_time_recovery_timestamp": {
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
			"vault_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// Computed
		},
	}
}

func createDatabaseCloudDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func updateDatabaseCloudDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func readDatabaseCloudDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseCloudDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseCloudDatabaseManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.Database
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseCloudDatabaseManagementResource-", DatabaseCloudDatabaseManagementResource(), s.D)
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) Create() error {

	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
	}

	if operation {
		// enable operation
		request := oci_database.EnableDatabaseManagementRequest{}

		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			request.DatabaseId = &tmp
		}

		if managementType, ok := s.D.GetOkExists("management_type"); ok {
			request.EnableDatabaseManagementDetails.ManagementType = oci_database.EnableDatabaseManagementDetailsManagementTypeEnum(managementType.(string))
		}

		if privateEndPointId, ok := s.D.GetOkExists("private_end_point_id"); ok {
			tmp := privateEndPointId.(string)
			request.PrivateEndPointId = &tmp
		}

		if serviceName, ok := s.D.GetOkExists("service_name"); ok {
			tmp := serviceName.(string)
			request.ServiceName = &tmp
		}

		if credentialdetails, ok := s.D.GetOkExists("credentialdetails"); ok {
			if tmpList := credentialdetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credentialdetails", 0)
				tmp, err := s.mapToDatabaseManagementCredentialDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.CredentialDetails = &tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res = &response.Database
		return s.getDatabaseFromWorkRequest(workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
	}
	// disable
	request := oci_database.DisableDatabaseManagementRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res = &response.Database
	return s.getDatabaseFromWorkRequest(workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) Update() error {
	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
	}

	if operation {
		// Update operation
		request := oci_database.ModifyDatabaseManagementRequest{}

		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			request.DatabaseId = &tmp
		}

		if managementType, ok := s.D.GetOkExists("management_type"); ok {
			request.ModifyDatabaseManagementDetails.ManagementType = oci_database.ModifyDatabaseManagementDetailsManagementTypeEnum(managementType.(string))
		}

		if privateEndPointId, ok := s.D.GetOkExists("private_end_point_id"); ok {
			tmp := privateEndPointId.(string)
			request.PrivateEndPointId = &tmp
		}

		if serviceName, ok := s.D.GetOkExists("service_name"); ok {
			tmp := serviceName.(string)
			request.ServiceName = &tmp
		}

		if credentialdetails, ok := s.D.GetOkExists("credentialdetails"); ok {
			if tmpList := credentialdetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credentialdetails", 0)
				tmp, err := s.mapToDatabaseManagementCredentialDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.CredentialDetails = &tmp
			}
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		if s.Res.KmsKeyVersionId != nil {
			s.D.Set("kms_key_version_id", *s.Res.KmsKeyVersionId)
		}

		if s.Res.LastBackupTimestamp != nil {
			s.D.Set("last_backup_timestamp", s.Res.LastBackupTimestamp.String())
		}
		response, err := s.Client.ModifyDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res = &response.Database
		return s.getDatabaseFromWorkRequest(workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
	}
	// disable
	request := oci_database.DisableDatabaseManagementRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res = &response.Database
	return s.getDatabaseFromWorkRequest(workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) Delete() error {
	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	if s.Res.VmClusterId != nil {
		s.D.Set("vm_cluster_id", *s.Res.VmClusterId)
	}
	if !operation {
		return nil
	}

	// disable
	request := oci_database.DisableDatabaseManagementRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.Database
	return s.getDatabaseFromWorkRequest(workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) SetData() error {
	return nil
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) mapToDatabaseManagementCredentialDetails(fieldKeyFormat string) (oci_database.DatabaseCredentialDetails, error) {
	result := oci_database.DatabaseCredentialDetails{}

	if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
		tmp := userName.(string)
		result.UserName = &tmp
	}

	if passwordSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_secret_id")); ok {
		tmp := passwordSecretId.(string)
		result.PasswordSecretId = &tmp
	}

	return result, nil
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) getDatabaseFromWorkRequest(workId *string, actionTypeEnum oci_work_requests.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {
	databaseId, err := tfresource.WaitForWorkRequest(s.WorkRequestClient, workId, "database", actionTypeEnum, timeout, s.DisableNotFoundRetries, true)
	log.Printf("[DEBUG] WaitForWorkRequest finished. databaseId: %v err: %v for workId: %v, actionTypeEnum: %v\n", *databaseId, err, *workId, actionTypeEnum)
	if err != nil {
		log.Printf("[ERROR] Database management operation failed, attempting to cancel the workrequest: %v for identifier: %v\n", *workId, databaseId)
		return err
	}

	s.D.SetId(*databaseId)

	return s.Get()
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) Get() error {
	request := oci_database.GetDatabaseRequest{}

	tmp := s.D.Id()
	request.DatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Database
	return nil
}
