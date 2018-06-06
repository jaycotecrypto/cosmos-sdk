// nolint
package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace sdk.CodespaceType = 4

	CodeInvalidValidator  CodeType = 101
	CodeInvalidDelegation CodeType = 102
	CodeInvalidInput      CodeType = 103
	CodeValidatorJailed   CodeType = 104
	CodeUnauthorized      CodeType = sdk.CodeUnauthorized
	CodeInternal          CodeType = sdk.CodeInternal
	CodeUnknownRequest    CodeType = sdk.CodeUnknownRequest
)

//validator
func ErrNilValidatorAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "validator address is nil")
}
func ErrNoValidatorFound(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidValidator, "validator does not exist for that address")
}
func ErrValidatorAlreadyExists(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidValidator, "validator already exist, cannot re-create validator")
}
func ErrValidatorRevoked(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidValidator, "validator for this address is currently revoked")
}
func ErrBadRemoveValidator(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidValidator, "error removing validator")
}
func ErrDescriptionLength(codespace sdk.CodespaceType, descriptor string, got, max int) sdk.Error {
	msg := fmt.Sprintf("bad description length for %v, got length %v, max is %v", descriptor, got, max)
	return sdk.NewError(codespace, CodeInvalidValidator, msg)
}
func ErrCommissionNegative(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidValidator, "commission must be positive")
}
func ErrCommissionHuge(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidValidator, "commission cannot be more than 100%")
}

// delegation
func ErrNilDelegatorAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "delegator address is nil")
}
func ErrBadDenom(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDelegation, "invalid coin denomination")
}
func ErrBadDelegationAmount(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDelegation, "amount must be > 0")
}
func ErrNoDelegation(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDelegation, "no delegation for this (address, validator) pair")
}
func ErrBadDelegatorAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDelegation, "delegator does not exist for that address")
}
func ErrNoDelegatorForAddress(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDelegation, "delegator does not contain this delegation")
}
func ErrInsufficientShares(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDelegation, "insufficient delegation shares")
}
func ErrDelegationValidatorEmpty(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDelegation, "cannot delegate to an empty validator")
}
func ErrNotEnoughDelegationShares(codespace sdk.CodespaceType, shares string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDelegation, fmt.Sprintf("not enough shares only have %v", shares))
}
func ErrBadSharesAmount(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDelegation, "shares must be > 0")
}
func ErrBadSharesPercent(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDelegation, "shares percent must be >0 and <=1")
}

// messages
func ErrBothShareMsgsGiven(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "both shares amount and shares percent provided")
}
func ErrNeitherShareMsgsGiven(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "neither shares amount nor shares percent provided")
}
func ErrMissingSignature(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidValidator, "missing signature")
}
