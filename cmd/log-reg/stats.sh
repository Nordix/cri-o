#!/usr/bin/env bash
set -eu

measure() {
local label="$1"
local cmd="$2"
local total=0

echo "== $label =="

sudo crictl stats >/dev/null 2>&1 || true

for i in 1 2 3 4 5; do
    t=$(/usr/bin/time -p sh -c "$cmd" >/dev/null 2>/tmp/time.out; awk '/^real /{print $2}' /tmp/time.out)
    echo "run $i: ${t}s"
    total=$(awk -v a="$total" -v b="$t" 'BEGIN { printf "%.2f", a + b }')
done

avg=$(awk -v t="$total" 'BEGIN { printf "%.2f", t / 5 }')
echo "avg: ${avg}s"
echo
}

measure "default timeout" "sudo crictl stats >/dev/null 2>&1"
measure "10s timeout" "sudo crictl --timeout=10s stats >/dev/null 2>&1"
