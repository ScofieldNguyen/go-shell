PID=$(pgrep -n -x go-shell)
if [ -z "$PID" ]; then
  echo "Could not find go-shell PID" >&2
  exit 1
fi

echo "Attaching dlv to PID $PID..."
dlv attach "$PID"
