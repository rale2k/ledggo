import hashlib
import random
import sys
import os
import socket
import json
from time import sleep
from flask import Flask, jsonify, request
import requests

app = Flask(__name__)

def get_nodes_info():
    nodes = {}
    for ip_port in app.config['nodes']:
        ip, port = ip_port.split(':')
        url = f"http://{ip}:{port}/info"
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

def select_random_node():
    nodes = app.config['nodes']
    return random.choice(nodes)

def handle_generate_block():
    random_node = select_random_node()
    ip, port = random_node.split(':')
    latest_block_url = f"http://{ip}:8080/blocks/last"
    try:
        response = requests.get(latest_block_url)
        latest_block = response.json()
        latest_block_hash = latest_block['hash'] if response.status_code == 200 else ""
        random_string = ''.join(random.choices('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ', k=10))
        concatenated_string = random_string + latest_block_hash
        block_hash = hashlib.sha256(concatenated_string.encode()).hexdigest()
        new_block = {
            "origin": f"{ip}:8080",
            "hash": block_hash,
            "data": random_string
        }
        post_block_url = f"http://{ip}:8080/blocks"
        response = requests.post(post_block_url, json=new_block)
        if response.status_code == 200:
            return new_block, 200
        else:
            return response.text, 500
    except requests.RequestException as e:
        return {"error": f"Error generating block: {e}"}, 500


def get_latest_block(node):
    ip, port = node.split(':')
    url = f"http://{ip}:{port}/blocks/last"
    try:
        response = requests.get(url)
        if response.status_code == 200:
            return response.json()
        else:
            return "Failed to fetch the latest block."
    except requests.RequestException as e:
        return f"Error getting latest block from {ip}:{port}: {e}"


@app.route('/generate_block', methods=['POST'])
def generate_block():
    result, status_code = handle_generate_block()
    return jsonify(result), status_code

@app.route('/nodes')
def nodes():
    return jsonify(get_nodes_info())

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

@app.route('/latest_blocks')
def latest_blocks():
    nodes = app.config['nodes']
    latest_blocks = {}
    for node in nodes:
        latest_block = get_latest_block(node)
        latest_blocks[node] = latest_block
    return jsonify(latest_blocks)


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
        sleep(0.1)
            
    app.config['nodes'] = previous_ips
    app.run(host='0.0.0.0', port=8000)

if __name__ == "__main__":
    main()
