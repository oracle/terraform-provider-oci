package basedb

import (
	"github.com/oracle/terraform-provider-oci/internal/integrationtest/basedb/dbsystems"
	"github.com/oracle/terraform-provider-oci/internal/integrationtest/basedb/network"
	"github.com/oracle/terraform-provider-oci/internal/integrationtest/basedb/tags"
)

var (
	BaseConfig = tags.BaseConfig + network.BaseConfig + dbsystems.BaseConfig
)
