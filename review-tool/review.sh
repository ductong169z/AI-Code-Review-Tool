#!/bin/bash

if [ ! -x "$0" ]; then
  chmod +x "$0"
  echo "✅ Granted execute permission to $0"
fi

echo "🔍 Reviewing staged changes..."

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
"$DIR/reviewer" --review-only

echo "✅ Review completed."
