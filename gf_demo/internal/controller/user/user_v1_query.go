package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf_demo/api/user/v1"
)

func (c *ControllerV1) Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
