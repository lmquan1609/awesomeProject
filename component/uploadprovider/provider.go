package uploadprovider

import (
	"awesomeProject/common"
	"context"
)

type UploadProvider interface {
	SaveFileUpload(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
