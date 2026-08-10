#!/bin/bash
set -e
MODEL=/models/MiniCPM5-1B-Q4_K_M.gguf
URL="https://huggingface.co/openbmb/MiniCPM5-1B-GGUF/resolve/main/MiniCPM5-1B-Q4_K_M.gguf?download=true"

if [ ! -f "$MODEL" ]; then
  echo "[llama] $MODEL not found, downloading from Hugging Face..."
  mkdir -p /models
  curl -L -o "$MODEL" "$URL"
fi

exec /app/llama-server -m "$MODEL" --port 8000 --host 0.0.0.0 -n 1024
