//autogenerated:yes
//nolint:revive,lll
package msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type EmergencyStopSrvReq struct {
	msg.Package `ros:"msgs"`

	Emergency uint8
}

type EmergencyStopSrvRes struct {
	msg.Package `ros:"msgs"`
}

type EmergencyStopSrv struct {
	msg.Package `ros:"mower_msgs"`
	EmergencyStopSrvReq
	EmergencyStopSrvRes
}
