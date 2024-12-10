#!/bin/bash


SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)/builds"


start_in_terminal() {
  local command=$1
  if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    gnome-terminal -- bash -c "$command > /dev/null 2>&1; exec bash"
  elif [[ "$OSTYPE" == "darwin"* ]]; then
    osascript -e "tell application \"Terminal\" to do script \"$command\""
  elif [[ "$OSTYPE" == "cygwin" ]] || [[ "$OSTYPE" == "msys" ]]; then
    mintty bash -c "$command; exec bash"
  elif [[ "$OSTYPE" == "win32" ]] || [[ "$OSTYPE" == "wsl"* ]]; then
    wsl bash -c "$command; exec bash"
  else
    echo "Sistema operativo non supportato!"
    exit 1
  fi
}


MAPPER_PATH="$SCRIPT_DIR/mapper"



start_in_terminal "$MAPPER_PATH localhost:50051"
start_in_terminal "$MAPPER_PATH localhost:50052"
start_in_terminal "$MAPPER_PATH localhost:50053"
start_in_terminal "$MAPPER_PATH localhost:50054"



