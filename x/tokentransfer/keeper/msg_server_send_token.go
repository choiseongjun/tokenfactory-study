package keeper

import (
	"context"
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"tokenfactory/x/tokentransfer/types"
)

func (k msgServer) SendToken(goCtx context.Context, msg *types.MsgSendToken) (*types.MsgSendTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 보내는 계정 주소 변환
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// 받는 계정 주소 변환
	toAddr, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	// 전송할 코인 생성
	amount := sdk.NewCoins(sdk.NewCoin("token", sdkmath.NewInt(int64(msg.Amount))))

	// 실제 전송 실행
	err = k.bankKeeper.SendCoins(ctx, fromAddr, toAddr, amount)
	if err != nil {
		return nil, err
	}

	// 이벤트 생성
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSendToken,
			sdk.NewAttribute(types.AttributeKeyFrom, msg.From),
			sdk.NewAttribute(types.AttributeKeyTo, msg.To),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", msg.Amount)),
		),
	)

	return &types.MsgSendTokenResponse{}, nil
}
