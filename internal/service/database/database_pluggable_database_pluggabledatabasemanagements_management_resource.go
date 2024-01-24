// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabasePluggableDatabasePluggabledatabasemanagementsManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabasePluggableDatabasePluggabledatabasemanagementsManagement,
		Read:     readDatabasePluggableDatabasePluggabledatabasemanagementsManagement,
		Update:   updateDatabasePluggableDatabasePluggabledatabasemanagementsManagement,
		Delete:   deleteDatabasePluggableDatabasePluggabledatabasemanagementsManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"pluggable_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_pluggabledatabasemanagement": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"credential_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"password_secret_id": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"private_end_point_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ssl_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_strings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"enable_pluggabledatabasemanagement": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional

						// Computed
						"all_connection_strings": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"pdb_default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pdb_ip_default": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"container_database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_restricted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"open_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pdb_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pluggable_database_management_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"enable_pluggabledatabasemanagement": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional

						// Computed
						"management_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabasePluggableDatabasePluggabledatabasemanagementsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabasePluggabledatabasemanagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabasePluggableDatabasePluggabledatabasemanagementsManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func readDatabasePluggableDatabasePluggabledatabasemanagementsManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateDatabasePluggableDatabasePluggabledatabasemanagementsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabasePluggabledatabasemanagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabasePluggableDatabasePluggabledatabasemanagementsManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabasePluggableDatabasePluggabledatabasemanagementsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabasePluggabledatabasemanagementsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabasePluggableDatabasePluggabledatabasemanagementsManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabasePluggableDatabasePluggabledatabasemanagementsManagementResponse struct {
	enableResponse  *oci_database.EnablePluggableDatabaseManagementResponse
	disableResponse *oci_database.DisablePluggableDatabaseManagementResponse
	modifyResponse  *oci_database.ModifyPluggableDatabaseManagementResponse
}

type DatabasePluggableDatabasePluggabledatabasemanagementsManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *DatabasePluggableDatabasePluggabledatabasemanagementsManagementResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabasePluggableDatabasePluggabledatabasemanagementsManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabasePluggableDatabasePluggabledatabasemanagementsManagementResource-", DatabasePluggableDatabasePluggabledatabasemanagementsManagementResource(), s.D)
}

func (s *DatabasePluggableDatabasePluggabledatabasemanagementsManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_pluggabledatabasemanagement"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database.EnablePluggableDatabaseManagementRequest{}

		if credentialDetails, ok := s.D.GetOkExists("credential_details"); ok {
			if tmpList := credentialDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credential_details", 0)
				tmp, err := s.mapToDatabaseCredentialDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.CredentialDetails = &tmp
			}
		}

		if pluggableDatabaseId, ok := s.D.GetOkExists("pluggable_database_id"); ok {
			tmp := pluggableDatabaseId.(string)
			request.PluggableDatabaseId = &tmp
		}

		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			request.Port = &tmp
		}

		if privateEndPointId, ok := s.D.GetOkExists("private_end_point_id"); ok {
			tmp := privateEndPointId.(string)
			request.PrivateEndPointId = &tmp
		}

		if protocol, ok := s.D.GetOkExists("protocol"); ok {
			request.Protocol = oci_database.EnablePluggableDatabaseManagementDetailsProtocolEnum(protocol.(string))
		}

		if role, ok := s.D.GetOkExists("role"); ok {
			request.Role = oci_database.EnablePluggableDatabaseManagementDetailsRoleEnum(role.(string))
		}

		if serviceName, ok := s.D.GetOkExists("service_name"); ok {
			tmp := serviceName.(string)
			request.ServiceName = &tmp
		}

		if sslSecretId, ok := s.D.GetOkExists("ssl_secret_id"); ok {
			tmp := sslSecretId.(string)
			request.SslSecretId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnablePluggableDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		if workId != nil {
			_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "pluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}

		s.Res.enableResponse = &response
		return nil
	}

	request := oci_database.DisablePluggableDatabaseManagementRequest{}

	if pluggableDatabaseId, ok := s.D.GetOkExists("pluggable_database_id"); ok {
		tmp := pluggableDatabaseId.(string)
		request.PluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisablePluggableDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	if workId != nil {
		_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "pluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *DatabasePluggableDatabasePluggabledatabasemanagementsManagementResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_pluggabledatabasemanagement"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_database.ModifyPluggableDatabaseManagementRequest{}
		if credentialDetails, ok := s.D.GetOkExists("credential_details"); ok {
			if tmpList := credentialDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credential_details", 0)
				tmp, err := s.mapToDatabaseCredentialDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.CredentialDetails = &tmp
			}
		}

		if pluggableDatabaseId, ok := s.D.GetOkExists("pluggable_database_id"); ok {
			tmp := pluggableDatabaseId.(string)
			request.PluggableDatabaseId = &tmp
		}

		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			request.Port = &tmp
		}

		if privateEndPointId, ok := s.D.GetOkExists("private_end_point_id"); ok {
			tmp := privateEndPointId.(string)
			request.PrivateEndPointId = &tmp
		}

		if protocol, ok := s.D.GetOkExists("protocol"); ok {
			request.Protocol = oci_database.ModifyPluggableDatabaseManagementDetailsProtocolEnum(protocol.(string))
		}

		if role, ok := s.D.GetOkExists("role"); ok {
			request.Role = oci_database.ModifyPluggableDatabaseManagementDetailsRoleEnum(role.(string))
		}

		if serviceName, ok := s.D.GetOkExists("service_name"); ok {
			tmp := serviceName.(string)
			request.ServiceName = &tmp
		}

		if sslSecretId, ok := s.D.GetOkExists("ssl_secret_id"); ok {
			tmp := sslSecretId.(string)
			request.SslSecretId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.ModifyPluggableDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId

		if workId != nil {
			_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "pluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}

		s.Res.modifyResponse = &response
		return nil
	}

	request := oci_database.DisablePluggableDatabaseManagementRequest{}

	if pluggableDatabaseId, ok := s.D.GetOkExists("pluggable_database_id"); ok {
		tmp := pluggableDatabaseId.(string)
		request.PluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisablePluggableDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	if workId != nil {
		_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "pluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *DatabasePluggableDatabasePluggabledatabasemanagementsManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_pluggabledatabasemanagement"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_database.DisablePluggableDatabaseManagementRequest{}

	if pluggableDatabaseId, ok := s.D.GetOkExists("pluggable_database_id"); ok {
		tmp := pluggableDatabaseId.(string)
		request.PluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisablePluggableDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	if workId != nil {
		_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "pluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *DatabasePluggableDatabasePluggabledatabasemanagementsManagementResourceCrud) SetData() error {
	return nil
}

func (s *DatabasePluggableDatabasePluggabledatabasemanagementsManagementResourceCrud) mapToDatabaseCredentialDetails(fieldKeyFormat string) (oci_database.DatabaseCredentialDetails, error) {
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
