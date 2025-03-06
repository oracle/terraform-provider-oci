package psql

import (
	"fmt"

	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
)

func (s *PsqlBackupResourceCrud) createPsqlSourceRegionClient(region string) error {
	if s.SourceRegionClient == nil {
		sourcePsqlClient, err := oci_psql.NewPostgresqlClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create client for the source region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&sourcePsqlClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionClient = &sourcePsqlClient
	}
	s.SourceRegionClient.SetRegion(region)
	return nil
}
