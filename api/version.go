package api

import (
	"github.com/cpw0321/mammoth/internal"
	"github.com/cpw0321/mammoth/types/responses"
)

// GetVersion 获取版本号
func GetVersion(c *internal.ServiceContext) {
	responses.Success(c, "v1.0.0")
}
