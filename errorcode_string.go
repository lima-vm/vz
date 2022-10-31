// Code generated by "stringer -type=ErrorCode"; DO NOT EDIT.

package vz

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrorInternal-1]
	_ = x[ErrorInvalidVirtualMachineConfiguration-2]
	_ = x[ErrorInvalidVirtualMachineState-3]
	_ = x[ErrorInvalidVirtualMachineStateTransition-4]
	_ = x[ErrorInvalidDiskImage-5]
	_ = x[ErrorVirtualMachineLimitExceeded-6]
	_ = x[ErrorNetworkError-7]
	_ = x[ErrorOutOfDiskSpace-8]
	_ = x[ErrorOperationCancelled-9]
	_ = x[ErrorNotSupported-10]
	_ = x[ErrorRestoreImageCatalogLoadFailed-10001]
	_ = x[ErrorInvalidRestoreImageCatalog-10002]
	_ = x[ErrorNoSupportedRestoreImagesInCatalog-10003]
	_ = x[ErrorRestoreImageLoadFailed-10004]
	_ = x[ErrorInvalidRestoreImage-10005]
	_ = x[ErrorInstallationRequiresUpdate-10006]
	_ = x[ErrorInstallationFailed-10007]
}

const (
	_ErrorCode_name_0 = "ErrorInternalErrorInvalidVirtualMachineConfigurationErrorInvalidVirtualMachineStateErrorInvalidVirtualMachineStateTransitionErrorInvalidDiskImageErrorVirtualMachineLimitExceededErrorNetworkErrorErrorOutOfDiskSpaceErrorOperationCancelledErrorNotSupported"
	_ErrorCode_name_1 = "ErrorRestoreImageCatalogLoadFailedErrorInvalidRestoreImageCatalogErrorNoSupportedRestoreImagesInCatalogErrorRestoreImageLoadFailedErrorInvalidRestoreImageErrorInstallationRequiresUpdateErrorInstallationFailed"
)

var (
	_ErrorCode_index_0 = [...]uint8{0, 13, 52, 83, 124, 145, 177, 194, 213, 236, 253}
	_ErrorCode_index_1 = [...]uint8{0, 34, 65, 103, 130, 154, 185, 208}
)

func (i ErrorCode) String() string {
	switch {
	case 1 <= i && i <= 10:
		i -= 1
		return _ErrorCode_name_0[_ErrorCode_index_0[i]:_ErrorCode_index_0[i+1]]
	case 10001 <= i && i <= 10007:
		i -= 10001
		return _ErrorCode_name_1[_ErrorCode_index_1[i]:_ErrorCode_index_1[i+1]]
	default:
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
