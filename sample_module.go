// Filename: nbtc_middleware.go
package nbtc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
)

// IBCMiddleware implements the ICS26 callbacks for the nBTC middleware.
// It wraps the standard IBC transfer module to add custom logic for handling
// incoming nBTC tokens.
type IBCMiddleware struct {
	// The underlying IBC application (in this case, the `x/transfer` module).
	app porttypes.IBCModule
	// A keeper for the custom logic (e.g., interacting with a DEX module).
	keeper Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware.
func NewIBCMiddleware(app porttypes.IBCModule, k Keeper) IBCMiddleware {
	return IBCMiddleware{
		app:    app,
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface.
func (im IBCMiddleware) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	// Pass the call to the underlying application.
	return im.app.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, version)
}

// ... other IBCModule interface methods (OnChanOpenTry, OnChanOpenAck, etc.) would be implemented here,
// simply passing the call to the underlying `app`.

// OnRecvPacket is the core of the middleware. It is called when a packet is received.
// It first calls the underlying `x/transfer` module's OnRecvPacket, and if that
// is successful, it executes its custom logic.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	// First, let the underlying `x/transfer` module handle the packet.
	// This will mint the IBC voucher for nBTC.
	ack := im.app.OnRecvPacket(ctx, packet, relayer)

	// If the underlying application returns a failed acknowledgement,
	// we do not proceed with our custom logic.
	if !ack.Success() {
		return ack
	}

	// The packet was successfully processed. Now, we can add our custom logic.
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		// This should not happen if the `x/transfer` module successfully processed it.
		// We'll return the original ack but log an error.
		ctx.Logger().Error("cannot unmarshal ICS-20 packet data", "error", err)
		return ack
	}

	// Check if the received token is nBTC from the Nomic channel.
	// The full denom path would be `port/channel/denom`.
	// For this example, let's assume we have a function to identify the nBTC denom.
	isNBTCToken := im.keeper.IsNBTCDenom(ctx, data.Denom)

	if isNBTCToken {
		// This is where the custom logic for a chain like Osmosis would go.
		// For example, you could automatically deposit the nBTC into a liquidity pool,
		// lend it out on a money market, or perform any other action.

		// Example: Log the reception of nBTC.
		receiver, _ := sdk.AccAddressFromBech32(data.Receiver)
		ctx.Logger().Info("Received nBTC via IBC", "receiver", receiver.String(), "amount", data.Amount)

		// Example: Trigger a custom function in our keeper.
		// err := im.keeper.DepositIntoLiquidityPool(ctx, receiver, data.Amount)
		// if err != nil {
		//     // Handle the error. This is a complex topic. You might want to
		//     // create a new acknowledgement with the error, or you might
		//     // just log it. For simplicity, we'll just log it here.
		//     ctx.Logger().Error("failed to deposit nBTC into liquidity pool", "error", err)
		// }
	}

	return ack
}


// ... other IBCModule interface methods (OnAcknowledgementPacket, OnTimeoutPacket) would also be implemented here.

// NOTE: This is a simplified example. A real-world implementation would require a `Keeper`
// to manage state (e.g., to identify the correct nBTC denom) and would need to be
// properly integrated into the application's module manager and routing.
