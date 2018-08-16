# Node Initialization
Currenty setting only allows to emulate consensus on a pre-fixed set of members.
In order to fully emulate consensuses employed by public blockchains, allowing dynamic membership needs to be realized.

## Public/Private Key
_crypto.GenerateKey()_ from _go-ethereum_ is adopted to generated public/private key

## Voting Power
The process of initialization of weight has two steps:

1. Random assignment with the range [0,1)
2. Compute the percentage of each node's weight with respect to the total weight

# Leader Election

1. Generate a random value between [0,1)
2. Add up nodes' weight until accumulated weight is larger than the random value
3. Return the ID of the last node contributing to the accumulated weight

