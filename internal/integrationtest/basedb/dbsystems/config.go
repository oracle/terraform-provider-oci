package dbsystems

var (
	DatasourceBaseConfig = DbSystemDatasourceConfig + DbHomesDatasourceConfig + DatabasesDatasourceConfig + DatabaseDatasourceConfig
	ResourceBaseConfig   = DbSystemResourceConfig
	BaseConfig           = DatasourceBaseConfig + ResourceBaseConfig
)
