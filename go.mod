module github.com/kulycloud/cli

go 1.15

require (
	github.com/alecthomas/kong v0.2.15
	github.com/kulycloud/api-server v1.0.0
	github.com/kulycloud/common v1.0.0
	github.com/kulycloud/protocol v1.0.0
	gopkg.in/yaml.v2 v2.4.0 // indirect

)

replace github.com/kulycloud/common v1.0.0 => ../common

replace github.com/kulycloud/protocol v1.0.0 => ../protocol

replace github.com/kulycloud/api-server v1.0.0 => ../api-server
