# ledggo

## Description
Simple HTTP based distributed ledger realized in Golang.

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

### Upon startup
The application querys each known node for nodes the know and saves them. After, the now extended pool of known nodes are all queried for the longest chain. The longest chain is chosen.

### Submitting a block
Blocks are submitted through the ```POST /blocks``` endpoint, hash and data in JSON form as strings. If the ledger is empty, the hash is the SHA256 hexstring of the data, otherwise, the SHA256 hexstring of the data to submit+{hash of previous block in chain}. These conditions are validated upon request, and distribution of the block will only begin on valid blocks.
   
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

**GET /blocks/last** - Returns the latest block saved.

**POST /blocks** - Saves a new block and distributes it to all known nodes.

## Runner API
**GET /nodes** - Returns a list of all nodes in the network and the nodes they know

**GET /blocks/<hash>** - Returns the status of a specific block across all nodes in the network.

**GET /blocks** - Returns all blocks from all nodes in the network.

**GET /latest_blocks** - Returns the latest block from each node in the network..

**POST /generate_block** - Generates a new block and adds it to a randomly selected node in the network.
