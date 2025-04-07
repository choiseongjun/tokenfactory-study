package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendToken{}

func NewMsgSendToken(creator string, from string, to string, amount uint64) *MsgSendToken {
	return &MsgSendToken{
		Creator: creator,
		From:    from,
		To:      to,
		Amount:  amount,
	}
}

func (msg *MsgSendToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
