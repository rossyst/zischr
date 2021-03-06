#!/usr/bin/env python3

import serial, time, sys, subprocess, argparse

class Scanner:
    def __init__(self, callback, dev):
        self.dev = dev
        self.callback = callback
        self._connect()

    def _connect(self):
        self.tty = serial.Serial(self.dev, baudrate=9600)

    def _read_scan(self):
        return self.tty.readline().decode("ascii").rstrip("\r\n").lstrip("\n")

    def loop(self):
        while 1:
            try:
                barcode = self._read_scan()
                if barcode == "":
                    self.tty.close()
                    time.sleep(1)
                    self._connect()
                self.callback(barcode)
            except IOError:
                self.tty.close()
                time.sleep(1)
                self._connect()

if __name__ == "__main__":

    parser = argparse.ArgumentParser(description='Scanner')
    parser.add_argument('--id', help='Terminal ID', required=True)
    parser.add_argument('---endpoint', help='backend endpoint', default="localhost:8080")
    parser.add_argument('--bin', help='backend binary', required=True)
    parser.add_argument('--dev', help='Scanner TTY device', default='/dev/ttyUSB0')
    args = parser.parse_args()


    def scan_callback(barcode):
        client_args = [args.bin, "client"]
        client_args += ["--terminalID", args.id]
        client_args += ["--endpoint", args.endpoint]
        if barcode == "13374204":
            client_args += ["scan", "Magic:CashIn" ]
        else:
            client_args += ["scan", "EAN:" + barcode]
        print(client_args)
        subprocess.call(client_args)

    s = Scanner(scan_callback, args.dev)
    s.loop()
