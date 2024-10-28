package request

import (
	"ginorm/errors"
)

type LoadPostmanRequest struct {
	CollectionId string `form:"collection_id" json:"nickname" binding:"required,min=1,max=200"`
	ApiKey       string `form:"api_key" json:"api_key" binding:"required,min=1,max=200"`
}

// Valid 验证表单
func (r *LoadPostmanRequest) Valid() *errors.BusinessError {

	return nil
}
