package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta `path:"/create" tags:"Create" method:"post" summary:"用户新增接口"`
}

type CreateRes struct {
	g.Meta `json:"g.Meta"`
}

type DeleteReq struct {
	g.Meta `path:"/delete" tags:"Delete" method:"post" summary:"用户删除接口"`
}

type DeleteRes struct {
	g.Meta `json:"g.Meta"`
}

type UpdateReq struct {
	g.Meta `path:"/update" tags:"Update" method:"post" summary:"用户修改接口"`
}

type UpdateRes struct {
	g.Meta `json:"g.Meta"`
}
type QueryReq struct {
	g.Meta `path:"/query" tags:"Update" method:"post" summary:"用户查询接口"`
}

type QueryRes struct {
	g.Meta `json:"g.Meta"`
}
