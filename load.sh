#!/bin/bash
echo "GET ${ADDRESS}/" | vegeta -cpus=1 attack -duration=5s -redirects=0 -rate=50/1s -workers=10 -timeout=10s -insecure | vegeta report
