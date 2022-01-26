// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v56/databasemanagement"
)

func DatabaseManagementManagedDatabasesChangeDatabaseParameterResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementManagedDatabasesChangeDatabaseParameter,
		Read:     readDatabaseManagementManagedDatabasesChangeDatabaseParameter,
		Delete:   deleteDatabaseManagementManagedDatabasesChangeDatabaseParameter,
		Schema: map[string]*schema.Schema{
			// Required
			"credentials": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"role": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parameters": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"update_comment": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDatabaseManagementManagedDatabasesChangeDatabaseParameter(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementManagedDatabasesChangeDatabaseParameter(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseManagementManagedDatabasesChangeDatabaseParameter(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateDatabaseManagementManagedDatabasesChangeDatabaseParameter(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.UpdateDatabaseParametersResult
	Req                    *oci_database_management.ChangeDatabaseParametersRequest
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceCrud) ID() string {
	return *s.Req.ManagedDatabaseId
}

func (s *DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceCrud) Create() error {
	request := oci_database_management.ChangeDatabaseParametersRequest{}

	if credentials, ok := s.D.GetOkExists("credentials"); ok {
		if tmpList := credentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credentials", 0)
			tmp, err := s.dbmgmt_mapToDatabaseCredentials(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Credentials = &tmp
		}
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]oci_database_management.ChangeDatabaseParameterDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parameters", stateDataIndex)
			converted, err := s.mapToChangeDatabaseParameterDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("parameters") {
			request.Parameters = tmp
		}
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_database_management.ParameterScopeEnum(scope.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.ChangeDatabaseParameters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UpdateDatabaseParametersResult
	s.Req = &request
	return nil
}

func (s *DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceCrud) SetData() error {
	s.D.Set("status", s.Res.Status)

	return nil
}

func (s *DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceCrud) mapToChangeDatabaseParameterDetails(fieldKeyFormat string) (oci_database_management.ChangeDatabaseParameterDetails, error) {
	result := oci_database_management.ChangeDatabaseParameterDetails{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if updateComment, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "update_comment")); ok {
		tmp := updateComment.(string)
		result.UpdateComment = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func ChangeDatabaseParameterDetailsToMap(obj oci_database_management.ChangeDatabaseParameterDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.UpdateComment != nil {
		result["update_comment"] = string(*obj.UpdateComment)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceCrud) dbmgmt_mapToDatabaseCredentials(fieldKeyFormat string) (oci_database_management.DatabaseCredentials, error) {
	result := oci_database_management.DatabaseCredentials{}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		tmp := password.(string)
		result.Password = &tmp
	}

	if role, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "role")); ok {
		result.Role = oci_database_management.DatabaseCredentialsRoleEnum(role.(string))
	}

	if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
		tmp := secretId.(string)
		result.SecretId = &tmp
	}

	if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
		tmp := userName.(string)
		result.UserName = &tmp
	}

	return result, nil
}

func Dbmgmt_DatabaseCredentialsToMap(obj *oci_database_management.DatabaseCredentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Password != nil {
		result["password"] = string(*obj.Password)
	}

	result["role"] = string(obj.Role)

	if obj.SecretId != nil {
		result["secret_id"] = string(*obj.SecretId)
	}

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	return result
}
