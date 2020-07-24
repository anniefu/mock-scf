#!/bin/bash

./proxy/server &
functions-framework --port 54321 --source python/main.py --target hello
