## chifra transactions

The `chifra transactions` tool retrieves transactions from the running Ethereum node (using the `--raw` option) or from TrueBlocks (the default). You may provide a transaction `hash`, a `blockNumber.transactionID` pair, or a `blockHash.transactionID` pair (or any combination) to specify the transaction(s).

### Usage

`Usage:`    chifra transactions [-a|-t|-u|-v|-h] &lt;tx_id&gt; [tx_id...]  
`Purpose:`  Retrieve a transaction from the cache or the node.

`Where:`  

| Short Cut | Option | Description |
| -------: | :------- | :------- |
|  | transactions | a space-separated list of one or more transaction identifiers (tx_hash, bn.txID, blk_hash.txID) (required) |
| -a | --articulate | articulate the transactions if an ABI is found for the 'to' address |
| -t | --trace | display the transaction's trace |
| -u | --uniq | display a list of uniq addresses found in this transaction |
| -v | --verbose | set verbose level. Either -v, --verbose or -v:n where 'n' is level |
| -h | --help | display this help screen |

`Notes:`

- `transactions` is one or more space-separated identifiers which may be either a transaction hash, 
  a blockNumber.transactionID pair, or a blockHash.transactionID pair, or any combination.
- This tool checks for valid input syntax, but does not check that the transaction requested exists.
- If the queried node does not store historical state, the results are undefined.

**Source**: [`tools/getTrans`](https://github.com/TrueBlocks/trueblocks-core/tree/master/src/tools/getTrans)
