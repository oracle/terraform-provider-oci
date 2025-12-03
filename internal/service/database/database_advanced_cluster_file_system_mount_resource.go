package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAdvancedClusterFileSystemMountResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseAdvancedClusterFileSystemMountWithContext,
		ReadContext:   readDatabaseAdvancedClusterFileSystemMountWithContext,
		DeleteContext: deleteDatabaseAdvancedClusterFileSystemMountWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"advanced_cluster_file_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func deleteDatabaseAdvancedClusterFileSystemMountWithContext(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}

func readDatabaseAdvancedClusterFileSystemMountWithContext(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}

func createDatabaseAdvancedClusterFileSystemMountWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseAdvancedClusterFileSystemMountResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func (s *DatabaseAdvancedClusterFileSystemMountResourceCrud) CreateWithContext(ctx context.Context) error {
	if err := s.MountAdvancedClusterFileSystem(ctx); err != nil {
		return err
	}
	return s.GetWithContext(ctx)
}

func (s *DatabaseAdvancedClusterFileSystemMountResourceCrud) MountAdvancedClusterFileSystem(ctx context.Context) error {
	request := oci_database.MountAdvancedClusterFileSystemRequest{}

	if advancedClusterFileSystemId, ok := s.D.GetOkExists("advanced_cluster_file_system_id"); ok {
		tmp := advancedClusterFileSystemId.(string)
		request.AdvancedClusterFileSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.MountAdvancedClusterFileSystem(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "advancedclusterfilesystem", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.GetWithContext(ctx)
}

type DatabaseAdvancedClusterFileSystemMountResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AdvancedClusterFileSystem
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseAdvancedClusterFileSystemMountResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.GetAdvancedClusterFileSystemRequest{}

	if advancedClusterFileSystemId, ok := s.D.GetOk("advanced_cluster_file_system_id"); ok {
		tmp := advancedClusterFileSystemId.(string)
		request.AdvancedClusterFileSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAdvancedClusterFileSystem(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.AdvancedClusterFileSystem
	return nil
}

func (s *DatabaseAdvancedClusterFileSystemMountResourceCrud) SetData() error {
	if s.Res.Id != nil {
		s.D.Set("advanced_cluster_file_system_id", *s.Res.Id)
	}
	return nil
}

func (s *DatabaseAdvancedClusterFileSystemMountResourceCrud) ID() string {
	return *s.Res.Id
}
