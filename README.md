# ledggo

## Description
Simple HTTP based distributed ledger realized in Golang.

### Upon startup
The application querys each known node for nodes the know and saves them. After, the now extended pool of known nodes are all queried for the longest chain. The longest chain is chosen.

### Submitting a block
Blocks are submitted through the ```POST /blocks``` endpoint, hash and data in JSON form as strings. If the ledger is empty, the hash is the SHA256 hexstring of the data, otherwise, the SHA256 hexstring is generated from ```{the data to submit}+{hash of previous block in chain}```. These conditions are validated upon request, and distribution of the block will only begin on valid blocks. Blocks are sent asynchronously to each known node. If the node submitted is already present, it is not propagated again.

### Peer discovery
When nodes communicate with eachother, they use a http header named ```node-ip``` that contains the location of the sending node in the form ```{ip}:{port}```. Using GIN middleware every request is checked for this header and if it is present, the node is saved to the recipient.

### Problems
* On a single machine(ryzen 7600x) up to 200 nodes work somewhat smoothly. More nodes may work if the transactions are submitted at a slow rate.
* Because of the agressive peer discovery, eventually, most of the time all nodes know all other nodes. With a large network the flooding of new blocks is very severe.
* If a node gets the same block from 2 neighbours at the same time, there may be a race condition where they both check that the block is valid, and they both add it into the chain. Needs a lock or something.

## Prerequisites
- Go 1.22 or higher
- Python 3.6 or higher

## Installation
1. Clone the repository:
    ```shell
    git clone https://github.com/your-username/ledggo.git
    ```

2. Install the required Python packages:
    ```shell
    pip install -r requirements.txt
    ```

## Running a single node
* Available launch options:

    ```shell
    -nodes string
            {ip}:{port} of known nodes to connect to separated by semicolons
    -port int
            Port to run the node on (default 8080)
    ```
   
## Running with the runner
1. Compile the Go program:
    ```shell
    go build -o ledggo
    ```

2. Run the program using the Python runner script:
    ```shell
    python runner.py {number of nodes to run}
    ```

## Ledggo API
**GET /nodes** - Returns a list of all known nodes.

**GET /blocks** - Returns all blocks saved.

**GET /txblocks** - Returns all blocks in transaction.

**GET /blocks/length** - Returns length of the ledger.

**GET /blocks/last** - Returns the latest block saved.

**POST /blocks** - Saves a new block and distributes it to all known nodes.

## Runner API
**GET /nodes** - Returns a list of all nodes in the network and the nodes they know

**GET /blocks/<hash>** - Returns the status of a specific block across all nodes in the network.

**GET /blocks** - Returns all blocks from all nodes in the network.

**GET /blocks** - Returns all blocks in transaction from all nodes in the network.

**GET /latest_blocks** - Returns the latest block from each node in the network..

**POST /generate_block** - Generates a new block and adds it to a randomly selected node in the network.

## DEBUGGING - Create 2 blocks simultaneously
```shell
(
  curl -X POST http://127.0.0.1:8080/blocks -H "Content-Type: application/json" -d '{"data": "VLvvaUvQHt",
  "hash": "4db150153938abef026b0c327d43097f1e5480b6205e6fc1ad8741df061927aa"}' &
  curl -X POST http://127.0.0.1:8080/blocks -H "Content-Type: application/json" -d '{"data": "XAGGuWCvGq",
  "hash": "dc751acf4f4694e08c66e497b9e320d11f0fcdac20d5709f400cad4ace092ce0"}' &
  wait
)
```
```shell
(
  curl -X POST http://127.0.0.1:8000/generate_block  &
  curl -X POST http://127.0.0.1:8000/generate_block  &
  wait
)
```