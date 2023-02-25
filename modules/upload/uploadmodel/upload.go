package uploadmodel

import "awesomeProject/common"

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err, "file is not image", "ErrFileIsNotImage")
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(err, "cannot save upload file", "ErrCannotSaveFile")
}
