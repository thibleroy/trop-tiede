package utils

import "back/lib"

func UpdateError(operation string, dataType string, code int, raw string) lib.IStatus {
	return lib.IStatus{
		Message:    dataType + " " + operation + " error.",
		Code:       code,
		RawMessage: raw,
	}
}

func FindError(dataType string, code int) lib.IStatus {
	return lib.IStatus{
		Message: dataType + " not found.",
		Code:    code,
	}
}

func UpdateSuccess(operation string, dataType string, code int) lib.IStatus {
	return lib.IStatus{
		Message: dataType + " " + operation + " success.",
		Code:    code,
	}
}

func FindSuccess(dataType string, code int) lib.IStatus {
	return lib.IStatus{
		Message: dataType + " found.",
		Code:    code,
	}
}
