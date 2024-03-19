import sys
import os
import socket
import json
from flask import Flask, jsonify, request
import requests

app = Flask(__name__)

def get_nodes():
    nodes = {}
    for ip_port in app.config['nodes']:
        ip, port = ip_port.split(':')
        url = f"http://{ip}:{port}/nodes"
        try:
            response = requests.get(url)
            if response.status_code == 200:
                nodes[ip_port] = response.json()
        except requests.RequestException as e:
            print(f"Error getting nodes from {ip}:{port}: {e}")
    return nodes

def get_blocks():
    blocks = {}
    for ip_port in app.config['nodes']:
        ip, port = ip_port.split(':')
        url = f"http://{ip}:{port}/blocks"
        try:
            response = requests.get(url)
            if response.status_code == 200:
                blocks[ip_port] = response.json()
        except requests.RequestException as e:
            print(f"Error getting blocks from {ip}:{port}: {e}")
    return blocks

@app.route('/nodes')
def nodes():
    return jsonify(get_nodes())

@app.route('/blocks/<hash>')
def block(hash):
    blocks = {}
    for ip_port in app.config['nodes']:
        ip, port = ip_port.split(':')
        url = f"http://{ip}:{port}/blocks/{hash}"
        try:
            response = requests.get(url)
            blocks[ip_port] = response.status_code == 200
        except requests.RequestException as e:
            print(f"Error getting block {hash} from {ip}:{port}: {e}")
    return jsonify(blocks)

@app.route('/blocks')
def all_blocks():
    return jsonify(get_blocks())

def launch_ledggo(port, previous_ips):
    try:
        previous_ips_str = ';'.join(previous_ips)
        os.execl("./ledggo", "ledggo", f"-port={port}", f"-nodes={previous_ips_str}")
    except OSError as e:
        print(f"Error launching ledggo: {e}")
        sys.exit(1)

def main():
    if len(sys.argv) != 2:
        print("Usage: python \"filename.py\" count")
        sys.exit(1)

    try:
        count = int(sys.argv[1])
    except ValueError:
        print("Count must be a valid integer.")
        sys.exit(1)

    if count <= 0:
        print("Count must be a positive integer.")
        sys.exit(1)

    base_port = 8080
    previous_ips = []

    for i in range(count):
        port = base_port + i
        print(f"Running ledggo instance {i+1} at http://127.0.0.1:{port} with ips {previous_ips[-3:]}")
        pid = os.fork()
        if pid == 0:
            launch_ledggo(port, previous_ips[-3:])
            sys.exit(0)
        elif pid < 0:
            print("Error forking process.")
            sys.exit(1)
        else:
            ip = "127.0.0.1"
            previous_ips.append(f"{ip}:{port}")
            
    app.config['nodes'] = previous_ips
    app.run(host='0.0.0.0', port=8000)

if __name__ == "__main__":
    main()