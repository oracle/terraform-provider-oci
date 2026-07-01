package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseCloudDatabaseManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseCloudDatabaseManagementWithContext,
		UpdateContext: updateDatabaseCloudDatabaseManagementWithContext,
		ReadContext:   readDatabaseCloudDatabaseManagementWithContext,
		DeleteContext: deleteDatabaseCloudDatabaseManagementWithContext,
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
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createDatabaseCloudDatabaseManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseCloudDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func updateDatabaseCloudDatabaseManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseCloudDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func readDatabaseCloudDatabaseManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteDatabaseCloudDatabaseManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseCloudDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
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

func (s *DatabaseCloudDatabaseManagementResourceCrud) CreateWithContext(ctx context.Context) error {

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

		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			request.Port = &tmp
		}

		if protocol, ok := s.D.GetOkExists("protocol"); ok {
			request.Protocol = oci_database.EnableDatabaseManagementDetailsProtocolEnum(protocol.(string))
		}

		if role, ok := s.D.GetOkExists("role"); ok {
			request.Role = oci_database.EnableDatabaseManagementDetailsRoleEnum(role.(string))
		}

		if sslSecretId, ok := s.D.GetOkExists("ssl_secret_id"); ok {
			tmp := sslSecretId.(string)
			request.SslSecretId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableDatabaseManagement(ctx, request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res = &response.Database
		return s.getDatabaseFromWorkRequest(ctx, workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
	}
	// disable
	request := oci_database.DisableDatabaseManagementRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableDatabaseManagement(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res = &response.Database
	return s.getDatabaseFromWorkRequest(ctx, workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) UpdateWithContext(ctx context.Context) error {
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

		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			request.Port = &tmp
		}

		if protocol, ok := s.D.GetOkExists("protocol"); ok {
			request.Protocol = oci_database.ModifyDatabaseManagementDetailsProtocolEnum(protocol.(string))
		}

		if role, ok := s.D.GetOkExists("role"); ok {
			request.Role = oci_database.ModifyDatabaseManagementDetailsRoleEnum(role.(string))
		}

		if sslSecretId, ok := s.D.GetOkExists("ssl_secret_id"); ok {
			tmp := sslSecretId.(string)
			request.SslSecretId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.ModifyDatabaseManagement(ctx, request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res = &response.Database
		return s.getDatabaseFromWorkRequest(ctx, workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
	}
	// disable
	request := oci_database.DisableDatabaseManagementRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableDatabaseManagement(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res = &response.Database
	return s.getDatabaseFromWorkRequest(ctx, workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) DeleteWithContext(ctx context.Context) error {
	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
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

	response, err := s.Client.DisableDatabaseManagement(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err := tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.Database
	return s.getDatabaseFromWorkRequest(ctx, workId, oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
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

func (s *DatabaseCloudDatabaseManagementResourceCrud) getDatabaseFromWorkRequest(ctx context.Context, workId *string, actionTypeEnum oci_work_requests.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {
	databaseId, err := tfresource.WaitForWorkRequest(s.WorkRequestClient, workId, "database", actionTypeEnum, timeout, s.DisableNotFoundRetries, true)
	log.Printf("[DEBUG] WaitForWorkRequest finished. databaseId: %v err: %v for workId: %v, actionTypeEnum: %v\n", *databaseId, err, *workId, actionTypeEnum)
	if err != nil {
		log.Printf("[ERROR] Database management operation failed, attempting to cancel the workrequest: %v for identifier: %v\n", *workId, databaseId)
		return err
	}

	s.D.SetId(*databaseId)

	return s.GetWithContext(ctx)
}

func (s *DatabaseCloudDatabaseManagementResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.GetDatabaseRequest{}

	tmp := s.D.Id()
	request.DatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDatabase(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Database
	return nil
}
