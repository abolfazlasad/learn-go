#!/usr/bin/env bash
# Watch every chapter myAnswer/ dir. On save, run that file with go run
# (or go test for *_test.go). Infinite loop until Ctrl+C.
#
#   ./watch.sh
#   make watch

set -u

ROOT="$(cd "$(dirname "$0")" && pwd)"
STATE="${TMPDIR:-/tmp}/learn-go-watch-$$"
mkdir -p "$STATE"
GO_PID=""

mtime_of() {
	if stat -f '%m' "$1" >/dev/null 2>&1; then
		stat -f '%m' "$1"
	else
		stat -c '%Y' "$1"
	fi
}

state_path() {
	# bash 3.2 has no associative arrays; one file per path
	local hex
	hex="$(printf '%s' "$1" | shasum | awk '{print $1}')"
	printf '%s/%s' "$STATE" "$hex"
}

stop_run() {
	if [[ -n "${GO_PID}" ]] && kill -0 "${GO_PID}" 2>/dev/null; then
		kill "${GO_PID}" 2>/dev/null || true
		wait "${GO_PID}" 2>/dev/null || true
	fi
	GO_PID=""
}

on_exit() {
	stop_run
	rm -rf "$STATE"
}

trap on_exit EXIT
trap 'echo; echo "watch: stopped"; exit 0' INT TERM

banner() {
	local label="  $*"
	local bar
	bar="$(printf '%*s' "${#label}" '' | tr ' ' '=')"
	echo
	echo "$bar"
	echo "$label"
	echo "$bar"
	echo
}

run_go() {
	local file="$1"
	local dir base impl testf
	dir="$(dirname "$file")"
	base="$(basename "$file")"

	stop_run

	banner "$(date '+%H:%M:%S')  $file"

	(
		cd "$dir" || exit 1
		if [[ "$base" == *_test.go ]]; then
			impl="${base%_test.go}.go"
			if [[ -f "$impl" ]]; then
				go test "$impl" "$base"
			else
				go test "$base"
			fi
		else
			testf="${base%.go}_test.go"
			if [[ -f "$testf" ]] && ! grep -q '^package main' "$base"; then
				go test "$base" "$testf"
			else
				go run "$base"
			fi
		fi
	) &
	GO_PID=$!
}

list_answers() {
	local ch
	for ch in "$ROOT"/[0-9][0-9]-*; do
		[[ -d "$ch/myAnswer" ]] || continue
		find "$ch/myAnswer" -type f -name '*.go' 2>/dev/null
	done
}

# Seed mtimes so the first loop does not run every file.
while IFS= read -r file; do
	[[ -n "$file" ]] || continue
	mtime_of "$file" >"$(state_path "$file")"
done < <(list_answers)

echo "watch: $ROOT"
echo "watch: saving a .go file in any chapter myAnswer/ runs it"
echo "watch: Ctrl+C to stop (also stops a running program)"
echo

while true; do
	while IFS= read -r file; do
		[[ -n "$file" ]] || continue
		[[ -f "$file" ]] || continue
		now="$(mtime_of "$file")"
		stamp="$(state_path "$file")"
		prev=""
		[[ -f "$stamp" ]] && prev="$(cat "$stamp")"
		if [[ "$now" != "$prev" ]]; then
			printf '%s' "$now" >"$stamp"
			run_go "$file"
		fi
	done < <(list_answers)
	sleep 0.4
done
