#!/usr/bin/env bash
# Watch every chapter myAnswer/ dir. On save, print questions/<name>.expected.txt
# when it exists, then run that file with go run (or go test for *_test.go).
# If input/<name>.input.txt exists, pipe it to stdin:
#   cat ../input/<name>.input.txt | go run <name>.go
# Infinite loop until Ctrl+C.
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

stem_of() {
	local stem
	stem="$(basename "$1")"
	stem="${stem%.go}"
	stem="${stem%_test}"
	stem="${stem%.input.txt}"
	printf '%s' "$stem"
}

chapter_dir_of() {
	local dir base
	dir="$(cd "$(dirname "$1")" && pwd)"
	base="$(basename "$dir")"
	case "$base" in
	myAnswer|questions|input)
		cd "$dir/.." && pwd
		;;
	*)
		printf '%s' "$dir"
		;;
	esac
}

expected_file_for() {
	local file="$1"
	printf '%s/questions/%s.expected.txt' "$(chapter_dir_of "$file")" "$(stem_of "$file")"
}

input_file_for() {
	local file="$1"
	local input
	input="$(chapter_dir_of "$file")/input/$(stem_of "$file").input.txt"
	if [[ -f "$input" ]]; then
		printf '%s' "$input"
	fi
}

answer_file_for() {
	local file="$1"
	printf '%s/myAnswer/%s.go' "$(chapter_dir_of "$file")" "$(stem_of "$file")"
}

print_expected() {
	local expected="$1"
	if [[ ! -f "$expected" ]]; then
		return 1
	fi
	banner "expected output"
	cat "$expected"
	echo
	return 0
}

run_go() {
	local file="$1"
	local dir base impl testf expected input
	dir="$(dirname "$file")"
	base="$(basename "$file")"

	if [[ "$base" == *.input.txt ]]; then
		file="$(answer_file_for "$file")"
		dir="$(dirname "$file")"
		base="$(basename "$file")"
		[[ -f "$file" ]] || return
	fi

	stop_run

	banner "$(date '+%H:%M:%S')  $file"

	expected="$(expected_file_for "$file")"
	if print_expected "$expected"; then
		banner "your output"
	fi

	input="$(input_file_for "$file")"
	if [[ -n "$input" ]]; then
		echo "stdin: cat ../input/$(basename "$input") | go run $base"
		echo
	fi

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
			elif [[ -n "$input" ]]; then
				cat "$input" | go run "$base"
			else
				go run "$base"
			fi
		fi
	) &
	GO_PID=$!
}

list_watched() {
	local ch
	for ch in "$ROOT"/[0-9][0-9]-*; do
		[[ -d "$ch/myAnswer" ]] || continue
		find "$ch/myAnswer" -type f -name '*.go' 2>/dev/null
		if [[ -d "$ch/input" ]]; then
			find "$ch/input" -type f -name '*.input.txt' 2>/dev/null
		fi
	done
}

# Seed mtimes so the first loop does not run every file.
while IFS= read -r file; do
	[[ -n "$file" ]] || continue
	mtime_of "$file" >"$(state_path "$file")"
done < <(list_watched)

echo "watch: $ROOT"
echo "watch: saving a .go file in any chapter myAnswer/ runs it"
echo "watch: prints questions/<name>.expected.txt first when that file exists"
echo "watch: if input/<name>.input.txt exists, pipes it: cat ../input/<name>.input.txt | go run ..."
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
	done < <(list_watched)
	sleep 0.4
done
