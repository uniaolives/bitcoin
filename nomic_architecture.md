# The Architecture of the Nomic Bridge

The Nomic bridge is a sophisticated system that is composed of several key components. These components work together to enable the secure and decentralized transfer of Bitcoin to the Cosmos ecosystem.

## The Nomic Chain

The Nomic chain is a Proof-of-Stake blockchain that is built using the Cosmos SDK. It is the core component of the Nomic bridge and is responsible for:

*   **Decentralized Custody:** The Nomic chain manages the decentralized custody of the locked Bitcoin. The Bitcoin is held in a multi-signature wallet, and the keys are distributed among the Nomic validators.
*   **Minting and Burning nBTC:** The Nomic chain is responsible for minting `nBTC` when Bitcoin is deposited and burning `nBTC` when it is withdrawn.
*   **IBC Compatibility:** The Nomic chain is IBC-compatible, which allows it to seamlessly communicate with other chains in the Cosmos ecosystem.

## Bitcoin Relayers

Bitcoin relayers are off-chain processes that are responsible for monitoring the Bitcoin blockchain and relaying information to the Nomic chain. Their primary responsibilities include:

*   **Monitoring for Deposits:** Bitcoin relayers monitor the Bitcoin blockchain for incoming deposits to the Nomic bridge.
*   **Submitting Deposit Transactions:** Once a deposit has been confirmed, the Bitcoin relayer submits a transaction to the Nomic chain to mint an equivalent amount of `nBTC`.
*   **Broadcasting Bitcoin Transactions:** When a user wants to withdraw their `nBTC` for Bitcoin, the Bitcoin relayer is responsible for broadcasting the withdrawal transaction to the Bitcoin network.

## IBC Relayers

IBC relayers are off-chain processes that are responsible for relaying IBC packets between the Nomic chain and other IBC-compatible chains, such as Osmosis. Their primary responsibilities include:

*   **Monitoring for IBC Packets:** IBC relayers monitor the Nomic chain and the counterparty chain for outgoing IBC packets.
*   **Forwarding IBC Packets:** When an IBC packet is detected, the IBC relayer forwards it to the destination chain.
*   **Maintaining IBC Channels:** IBC relayers are responsible for maintaining the health of the IBC channels between the Nomic chain and the counterparty chains.
