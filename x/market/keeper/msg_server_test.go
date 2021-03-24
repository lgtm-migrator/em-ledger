package keeper

import (
	"context"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/e-money/em-ledger/x/market/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
)

func TestAddLimitOrder(t *testing.T) {
	var (
		ownerAddr = randomAccAddress()
		gotOrder  types.Order
	)

	keeper := marketKeeperMock{}
	svr := NewMsgServerImpl(&keeper)

	specs := map[string]struct {
		req       *types.MsgAddLimitOrder
		mockFn    func(ctx sdk.Context, aggressiveOrder types.Order) (*sdk.Result, error)
		expErr    bool
		expEvents sdk.Events
		expOrder  types.Order
	}{
		"all good": {
			req: &types.MsgAddLimitOrder{
				Owner:         ownerAddr.String(),
				ClientOrderId: "myClientIOrderID",
				TimeInForce:   types.TimeInForce_FillOrKill,
				Source:        sdk.Coin{Denom: "eeur", Amount: sdk.OneInt()},
				Destination:   sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
			},
			mockFn: func(ctx sdk.Context, aggressiveOrder types.Order) (*sdk.Result, error) {
				gotOrder = aggressiveOrder
				return &sdk.Result{
					Events: []abcitypes.Event{{
						Type:       "testing",
						Attributes: []abcitypes.EventAttribute{{Key: []byte("foo"), Value: []byte("bar")}},
					}},
				}, nil
			},
			expEvents: sdk.Events{{
				Type:       "testing",
				Attributes: []abcitypes.EventAttribute{{Key: []byte("foo"), Value: []byte("bar")}},
			}},
			expOrder: types.Order{
				TimeInForce:       types.TimeInForce_FillOrKill,
				Owner:             ownerAddr.String(),
				ClientOrderID:     "myClientIOrderID",
				Source:            sdk.Coin{Denom: "eeur", Amount: sdk.OneInt()},
				SourceRemaining:   sdk.OneInt(),
				SourceFilled:      sdk.ZeroInt(),
				Destination:       sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
				DestinationFilled: sdk.ZeroInt(),
			},
		},
		"owner missing": {
			req: &types.MsgAddLimitOrder{
				ClientOrderId: "myClientIOrderID",
				TimeInForce:   types.TimeInForce_FillOrKill,
				Source:        sdk.Coin{Denom: "eeur", Amount: sdk.OneInt()},
				Destination:   sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
			},
			expErr: true,
		},
		"owner invalid": {
			req: &types.MsgAddLimitOrder{
				Owner:         "invalid",
				ClientOrderId: "myClientIOrderID",
				TimeInForce:   types.TimeInForce_FillOrKill,
				Source:        sdk.Coin{Denom: "eeur", Amount: sdk.OneInt()},
				Destination:   sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
			},
			expErr: true,
		},
		"processing failure": {
			req: &types.MsgAddLimitOrder{
				Owner:         ownerAddr.String(),
				ClientOrderId: "myClientIOrderID",
				TimeInForce:   types.TimeInForce_FillOrKill,
				Source:        sdk.Coin{Denom: "eeur", Amount: sdk.OneInt()},
				Destination:   sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
			},
			mockFn: func(ctx sdk.Context, aggressiveOrder types.Order) (*sdk.Result, error) {
				return nil, errors.New("testing")
			},
			expErr: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			keeper.NewOrderSingleFn = spec.mockFn
			eventManager := sdk.NewEventManager()
			ctx := sdk.Context{}.WithContext(context.Background()).WithEventManager(eventManager)
			_, gotErr := svr.AddLimitOrder(sdk.WrapSDKContext(ctx), spec.req)
			if spec.expErr {
				require.Error(t, gotErr)
				return
			}
			require.NoError(t, gotErr)
			assert.Equal(t, spec.expEvents, eventManager.Events())
			assert.Equal(t, spec.expOrder, gotOrder)
		})
	}
}

func TestAddMarketOrder(t *testing.T) {
	var (
		ownerAddr        = randomAccAddress()
		gotDenom         string
		gotDst           sdk.Coin
		gotMaxSlippage   sdk.Dec
		gotOwner         sdk.AccAddress
		gotTimeInForce   types.TimeInForce
		gotClientOrderId string
	)

	keeper := marketKeeperMock{}
	svr := NewMsgServerImpl(&keeper)

	specs := map[string]struct {
		req       *types.MsgAddMarketOrder
		mockFn    func(ctx sdk.Context, srcDenom string, dst sdk.Coin, maxSlippage sdk.Dec, owner sdk.AccAddress, timeInForce types.TimeInForce, clientOrderId string) (*sdk.Result, error)
		expErr    bool
		expEvents sdk.Events
	}{
		"all good": {
			req: &types.MsgAddMarketOrder{
				Owner:         ownerAddr.String(),
				ClientOrderId: "myClientIOrderID",
				TimeInForce:   types.TimeInForce_FillOrKill,
				Destination:   sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
				MaxSlippage:   sdk.NewDec(10),
			},
			mockFn: func(ctx sdk.Context, srcDenom string, dst sdk.Coin, maxSlippage sdk.Dec, owner sdk.AccAddress, timeInForce types.TimeInForce, clientOrderId string) (*sdk.Result, error) {
				gotDenom, gotDst, gotMaxSlippage, gotOwner, gotTimeInForce, gotClientOrderId = srcDenom, dst, maxSlippage, owner, timeInForce, clientOrderId
				return &sdk.Result{
					Events: []abcitypes.Event{{
						Type:       "testing",
						Attributes: []abcitypes.EventAttribute{{Key: []byte("foo"), Value: []byte("bar")}},
					}},
				}, nil
			},
			expEvents: sdk.Events{{
				Type:       "testing",
				Attributes: []abcitypes.EventAttribute{{Key: []byte("foo"), Value: []byte("bar")}},
			}},
		},
		"owner missing": {
			req: &types.MsgAddMarketOrder{
				ClientOrderId: "myClientIOrderID",
				TimeInForce:   types.TimeInForce_FillOrKill,
				Destination:   sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
				MaxSlippage:   sdk.NewDec(10),
			},
			expErr: true,
		},
		"owner invalid": {
			req: &types.MsgAddMarketOrder{
				Owner:         "invalid",
				ClientOrderId: "myClientIOrderID",
				TimeInForce:   types.TimeInForce_FillOrKill,
				Destination:   sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
				MaxSlippage:   sdk.NewDec(10),
			},
			expErr: true,
		},
		"processing failure": {
			req: &types.MsgAddMarketOrder{
				Owner:         ownerAddr.String(),
				ClientOrderId: "myClientIOrderID",
				TimeInForce:   types.TimeInForce_FillOrKill,
				Destination:   sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
				MaxSlippage:   sdk.NewDec(10),
			},
			mockFn: func(ctx sdk.Context, srcDenom string, dst sdk.Coin, maxSlippage sdk.Dec, owner sdk.AccAddress, timeInForce types.TimeInForce, clientOrderId string) (*sdk.Result, error) {
				return nil, errors.New("testing")
			},
			expErr: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			keeper.NewMarketOrderWithSlippageFn = spec.mockFn
			eventManager := sdk.NewEventManager()
			ctx := sdk.Context{}.WithContext(context.Background()).WithEventManager(eventManager)
			_, gotErr := svr.AddMarketOrder(sdk.WrapSDKContext(ctx), spec.req)
			if spec.expErr {
				require.Error(t, gotErr)
				return
			}
			require.NoError(t, gotErr)
			assert.Equal(t, spec.expEvents, eventManager.Events())
			assert.Equal(t, spec.req.Source, gotDenom)
			assert.Equal(t, spec.req.Destination, gotDst)
			assert.Equal(t, spec.req.MaxSlippage, gotMaxSlippage)
			assert.Equal(t, ownerAddr, gotOwner)
			assert.Equal(t, spec.req.TimeInForce, gotTimeInForce)
			assert.Equal(t, spec.req.ClientOrderId, gotClientOrderId)
		})
	}
}
func TestCancelOrder(t *testing.T) {
	var (
		ownerAddr        = randomAccAddress()
		gotOwner         sdk.AccAddress
		gotClientOrderId string
	)

	keeper := marketKeeperMock{}
	svr := NewMsgServerImpl(&keeper)

	specs := map[string]struct {
		req       *types.MsgCancelOrder
		mockFn    func(ctx sdk.Context, owner sdk.AccAddress, clientOrderId string) (*sdk.Result, error)
		expErr    bool
		expEvents sdk.Events
	}{
		"all good": {
			req: &types.MsgCancelOrder{
				Owner:         ownerAddr.String(),
				ClientOrderId: "myClientIOrderID",
			},
			mockFn: func(ctx sdk.Context, owner sdk.AccAddress, clientOrderId string) (*sdk.Result, error) {
				gotOwner, gotClientOrderId = owner, clientOrderId
				return &sdk.Result{
					Events: []abcitypes.Event{{
						Type:       "testing",
						Attributes: []abcitypes.EventAttribute{{Key: []byte("foo"), Value: []byte("bar")}},
					}},
				}, nil
			},
			expEvents: sdk.Events{{
				Type:       "testing",
				Attributes: []abcitypes.EventAttribute{{Key: []byte("foo"), Value: []byte("bar")}},
			}},
		},
		"owner missing": {
			req: &types.MsgCancelOrder{
				ClientOrderId: "myClientIOrderID",
			},
			expErr: true,
		},
		"owner invalid": {
			req: &types.MsgCancelOrder{
				Owner:         "invalid",
				ClientOrderId: "myClientIOrderID",
			},
			expErr: true,
		},
		"processing failure": {
			req: &types.MsgCancelOrder{
				Owner:         ownerAddr.String(),
				ClientOrderId: "myClientIOrderID",
			},
			mockFn: func(ctx sdk.Context, owner sdk.AccAddress, clientOrderId string) (*sdk.Result, error) {
				return nil, errors.New("testing")
			},
			expErr: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			keeper.CancelOrderFn = spec.mockFn
			eventManager := sdk.NewEventManager()
			ctx := sdk.Context{}.WithContext(context.Background()).WithEventManager(eventManager)
			_, gotErr := svr.CancelOrder(sdk.WrapSDKContext(ctx), spec.req)
			if spec.expErr {
				require.Error(t, gotErr)
				return
			}
			require.NoError(t, gotErr)
			assert.Equal(t, spec.expEvents, eventManager.Events())
			assert.Equal(t, ownerAddr, gotOwner)
			assert.Equal(t, spec.req.ClientOrderId, gotClientOrderId)
		})
	}
}
func TestCancelReplaceLimitOrder(t *testing.T) {
	var (
		ownerAddr            = randomAccAddress()
		gotOrder             types.Order
		gotOrigClientOrderId string
	)

	keeper := marketKeeperMock{}
	svr := NewMsgServerImpl(&keeper)

	specs := map[string]struct {
		req       *types.MsgCancelReplaceLimitOrder
		mockFn    func(ctx sdk.Context, newOrder types.Order, origClientOrderId string) (*sdk.Result, error)
		expErr    bool
		expEvents sdk.Events
		expOrder  types.Order
	}{
		"all good": {
			req: &types.MsgCancelReplaceLimitOrder{
				Owner:             ownerAddr.String(),
				OrigClientOrderId: "origClientID",
				NewClientOrderId:  "myNewClientID",
				Source:            sdk.Coin{Denom: "eeur", Amount: sdk.OneInt()},
				Destination:       sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
			},
			mockFn: func(ctx sdk.Context, newOrder types.Order, origClientOrderId string) (*sdk.Result, error) {
				gotOrder, gotOrigClientOrderId = newOrder, origClientOrderId
				return &sdk.Result{
					Events: []abcitypes.Event{{
						Type:       "testing",
						Attributes: []abcitypes.EventAttribute{{Key: []byte("foo"), Value: []byte("bar")}},
					}},
				}, nil
			},
			expEvents: sdk.Events{{
				Type:       "testing",
				Attributes: []abcitypes.EventAttribute{{Key: []byte("foo"), Value: []byte("bar")}},
			}},
			expOrder: types.Order{
				TimeInForce:       types.TimeInForce_GoodTillCancel,
				Owner:             ownerAddr.String(),
				ClientOrderID:     "myNewClientID",
				Source:            sdk.Coin{Denom: "eeur", Amount: sdk.OneInt()},
				SourceRemaining:   sdk.OneInt(),
				SourceFilled:      sdk.ZeroInt(),
				Destination:       sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
				DestinationFilled: sdk.ZeroInt(),
			},
		},
		"owner missing": {
			req: &types.MsgCancelReplaceLimitOrder{
				OrigClientOrderId: "origClientID",
				NewClientOrderId:  "newClientID",
				Source:            sdk.Coin{Denom: "eeur", Amount: sdk.OneInt()},
				Destination:       sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
			},
			expErr: true,
		},
		"owner invalid": {
			req: &types.MsgCancelReplaceLimitOrder{
				Owner:             "invalid",
				OrigClientOrderId: "origClientID",
				NewClientOrderId:  "newClientID",
				Source:            sdk.Coin{Denom: "eeur", Amount: sdk.OneInt()},
				Destination:       sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
			},
			expErr: true,
		},
		"processing failure": {
			req: &types.MsgCancelReplaceLimitOrder{
				Owner:             ownerAddr.String(),
				OrigClientOrderId: "origClientID",
				NewClientOrderId:  "newClientID",
				Source:            sdk.Coin{Denom: "eeur", Amount: sdk.OneInt()},
				Destination:       sdk.Coin{Denom: "alx", Amount: sdk.OneInt()},
			},
			mockFn: func(ctx sdk.Context, newOrder types.Order, origClientOrderId string) (*sdk.Result, error) {
				return nil, errors.New("testing")
			},
			expErr: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			keeper.CancelReplaceOrderFn = spec.mockFn
			eventManager := sdk.NewEventManager()
			ctx := sdk.Context{}.WithContext(context.Background()).WithEventManager(eventManager)
			_, gotErr := svr.CancelReplaceLimitOrder(sdk.WrapSDKContext(ctx), spec.req)
			if spec.expErr {
				require.Error(t, gotErr)
				return
			}
			require.NoError(t, gotErr)
			assert.Equal(t, spec.expEvents, eventManager.Events())
			assert.Equal(t, spec.expOrder, gotOrder)
			assert.Equal(t, spec.req.OrigClientOrderId, gotOrigClientOrderId)
		})
	}
}

type marketKeeperMock struct {
	NewMarketOrderWithSlippageFn func(ctx sdk.Context, srcDenom string, dst sdk.Coin, maxSlippage sdk.Dec, owner sdk.AccAddress, timeInForce types.TimeInForce, clientOrderId string) (*sdk.Result, error)
	NewOrderSingleFn             func(ctx sdk.Context, aggressiveOrder types.Order) (*sdk.Result, error)
	CancelOrderFn                func(ctx sdk.Context, owner sdk.AccAddress, clientOrderId string) (*sdk.Result, error)
	CancelReplaceOrderFn         func(ctx sdk.Context, newOrder types.Order, origClientOrderId string) (*sdk.Result, error)
}

func (m marketKeeperMock) NewMarketOrderWithSlippage(ctx sdk.Context, srcDenom string, dst sdk.Coin, maxSlippage sdk.Dec, owner sdk.AccAddress, timeInForce types.TimeInForce, clientOrderId string) (*sdk.Result, error) {
	if m.NewMarketOrderWithSlippageFn == nil {
		panic("not expected to be called")
	}
	return m.NewMarketOrderWithSlippageFn(ctx, srcDenom, dst, maxSlippage, owner, timeInForce, clientOrderId)
}

func (m marketKeeperMock) NewOrderSingle(ctx sdk.Context, aggressiveOrder types.Order) (*sdk.Result, error) {
	if m.NewOrderSingleFn == nil {
		panic("not expected to be called")
	}
	return m.NewOrderSingleFn(ctx, aggressiveOrder)
}

func (m marketKeeperMock) CancelOrder(ctx sdk.Context, owner sdk.AccAddress, clientOrderId string) (*sdk.Result, error) {
	if m.CancelOrderFn == nil {
		panic("not expected to be called")
	}
	return m.CancelOrderFn(ctx, owner, clientOrderId)
}

func (m marketKeeperMock) CancelReplaceOrder(ctx sdk.Context, newOrder types.Order, origClientOrderId string) (*sdk.Result, error) {
	if m.CancelReplaceOrderFn == nil {
		panic("not expected to be called")
	}
	return m.CancelReplaceOrderFn(ctx, newOrder, origClientOrderId)
}

func randomAccAddress() sdk.AccAddress {
	return rand.Bytes(sdk.AddrLen)
}