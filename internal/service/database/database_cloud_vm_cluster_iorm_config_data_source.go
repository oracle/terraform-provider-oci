package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseCloudVmClusterIormConfigDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cloud_vm_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseCloudVmClusterIormConfigResource(), fieldMap, readSingularDatabaseCloudVmClusterIormConfigWithContext)
}

func readSingularDatabaseCloudVmClusterIormConfigWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseCloudVmClusterConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseCloudVmClusterConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetCloudVmClusterIormConfigResponse
}

func (s *DatabaseCloudVmClusterConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseCloudVmClusterConfigDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.GetCloudVmClusterIormConfigRequest{}

	if cloudVmClusterId, ok := s.D.GetOkExists("cloud_vm_cluster_id"); ok {
		tmp := cloudVmClusterId.(string)
		request.CloudVmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetCloudVmClusterIormConfig(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseCloudVmClusterConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseCloudVmClusterIormConfigDataSource-", DatabaseCloudVmClusterIormConfigDataSource(), s.D))

	dbPlans := []interface{}{}
	for _, item := range s.Res.DbPlans {
		if configMap := dbIormConfigToMap(item); configMap != nil {
			dbPlans = append(dbPlans, configMap)
		}
	}
	s.D.Set("db_plans", dbPlans)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("objective", s.Res.Objective)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
