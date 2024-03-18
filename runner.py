import sys
import os
import socket

def launch_ledggo(port, previous_ips):
    try:
        previous_ips_str = ' '.join(previous_ips)
        os.execl("./ledggo", "ledggo", f"-port={port}", previous_ips_str)
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

    for _ in range(count):
        os.wait()

if __name__ == "__main__":
    main()