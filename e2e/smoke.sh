#!/usr/bin/env bash
# Smoke tests the pets API over real HTTP. Usage: smoke.sh [base-url]
set -euo pipefail

BASE="${1:-http://localhost:8080}"
TMP="$(mktemp)"
trap 'rm -f "$TMP"' EXIT

STATUS="" CTYPE="" BODY=""

req() {
  local method="$1" path="$2" body="${3:-}"
  local out
  out="$(curl -sS -o "$TMP" -w '%{http_code} %{content_type}' -X "$method" \
    -H 'Content-Type: application/json' ${body:+--data "$body"} "$BASE$path")"
  STATUS="${out%% *}"
  CTYPE="${out#* }"
  BODY="$(cat "$TMP")"
}

fail() {
  echo "FAIL: $*" >&2
  echo "  status=$STATUS content-type=$CTYPE" >&2
  echo "  body=$BODY" >&2
  exit 1
}

expect_status() { [[ "$STATUS" == "$1" ]] || fail "$2: expected status $1"; }
expect_json() { [[ "$CTYPE" == application/json* ]] || fail "$1: expected JSON content type"; }
expect_jq() { [[ "$(jq -r "$1" <<<"$BODY")" == "$2" ]] || fail "$3: expected $1 == $2"; }

echo "waiting for $BASE ..."
for _ in $(seq 1 30); do
  curl -fsS -o /dev/null "$BASE/pets" && break
  sleep 1
done

req GET /pets
expect_status 200 "list empty"
expect_json "list empty"
expect_jq 'length' 0 "list empty"
echo "ok  list empty"

req POST /pets '{"name":"Rex","species":"dog"}'
expect_status 201 "create"
expect_json "create"
expect_jq '.id' 1 "create"
expect_jq '.name' Rex "create"
echo "ok  create"

req GET /pets/1
expect_status 200 "get"
expect_json "get"
expect_jq '.name' Rex "get"
echo "ok  get"

req PUT /pets/1 '{"name":"Rexy","species":"dog"}'
expect_status 200 "update"
expect_jq '.name' Rexy "update"
echo "ok  update"

req GET /pets
expect_status 200 "list one"
expect_jq 'length' 1 "list one"
echo "ok  list one"

req POST /pets '{"species":"cat"}'
expect_status 422 "reject invalid"
expect_json "reject invalid"
echo "ok  reject invalid"

req DELETE /pets/1
expect_status 204 "delete"
echo "ok  delete"

req GET /pets/1
expect_status 404 "get deleted"
echo "ok  get deleted"

req GET /pets/abc
expect_status 400 "invalid id"
echo "ok  invalid id"

echo "all smoke tests passed"
