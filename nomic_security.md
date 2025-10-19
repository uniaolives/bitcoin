# The Security Model of the Nomic Bridge

Nomic has been designed with a strong emphasis on security. The bridge incorporates several features to mitigate the risks associated with cross-chain bridges and to ensure that users' funds are safe.

## Decentralized Custody

The foundation of Nomic's security model is its decentralized custody of the locked Bitcoin. Unlike centralized bridges that rely on a single, trusted entity, Nomic uses a multi-signature wallet to hold the locked Bitcoin. The keys to this wallet are distributed among the Nomic validators, which means that a large number of validators would need to collude in order to steal the funds.

## Emergency Disbursal

The Emergency Disbursal mechanism is a novel security feature that is designed to protect users in the event of a catastrophic failure of the Nomic chain. If the Nomic chain were to halt or become unresponsive, the Emergency Disbursal mechanism would be triggered.

When this happens, the control of the locked Bitcoin is transferred to the validators of the counterparty chains (e.g., Osmosis). The validators of the counterparty chain can then work together to return the Bitcoin to the `nBTC` holders on their respective chains. This ensures that users can always recover their funds, even if the Nomic chain itself is compromised.

## Slashing

Nomic implements a robust slashing mechanism to penalize malicious or negligent validators. If a validator is caught trying to double-sign a transaction or if they are offline for an extended period of time, a portion of their staked tokens will be slashed. This provides a strong economic incentive for validators to act honestly and to maintain the security of the network.

## Circuit Breakers

Nomic has implemented circuit breakers that are designed to prevent a large-scale, rapid draining of the Bitcoin reserves. These circuit breakers enforce rate limits on withdrawals and signatory changes. If the rate of withdrawals or signatory changes exceeds a certain threshold, the circuit breaker will be triggered, and the bridge will be temporarily halted. This gives the validators time to investigate the situation and to take action to prevent any further losses.
