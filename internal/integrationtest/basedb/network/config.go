package network

var (
	NetworkDatasourceBaseConfig = AvailabilityDomainsDatasourceConfig
	NetworkResourceBaseConfig   = VcnResourceConfig + SubnetResourceConfig + InternetGatewayResourceConfig + RouteTableResourceConfig + SecurityListResourceConfig
	BaseConfig                  = NetworkDatasourceBaseConfig + NetworkResourceBaseConfig
)
