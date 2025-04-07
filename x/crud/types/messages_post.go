package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePost{}

func NewMsgCreatePost(
	creator string,
	index string,
	title string,
	body string,

) *MsgCreatePost {
	return &MsgCreatePost{
		Creator: creator,
		Index:   index,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgCreatePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePost{}

func NewMsgUpdatePost(
	creator string,
	index string,
	title string,
	body string,

) *MsgUpdatePost {
	return &MsgUpdatePost{
		Creator: creator,
		Index:   index,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgUpdatePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePost{}

func NewMsgDeletePost(
	creator string,
	index string,

) *MsgDeletePost {
	return &MsgDeletePost{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeletePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
