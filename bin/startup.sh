#!/bin/bash

SERVICE_PATH="./"
PID_DIR="./pid"
LOG_DIR="./logs"

SERVICES="*-api"

# 检查PID目录是否存在，如果不存在则创建
if [ ! -d "$PID_DIR" ]; then
  mkdir -p "$PID_DIR"
fi

if [ ! -d "$LOG_DIR" ]; then
  mkdir -p "$LOG_DIR"
fi

check() {
  ps -ef | grep $1 | grep -v grep >/dev/null
  if [ $? -ne 0 ]; then
    return 1
  else
    return 0
  fi
}

start() {
  app=$(basename $1)
  if [ -f "$PID_DIR/$app.pid" ]; then
    check $(cat "$PID_DIR/$1.pid")
    if [ $? == 0 ]; then
      echo "Service $app Already Running"
      return
    fi
  fi
  nohup $1 >"$LOG_DIR/$app.log" 2>&1 &
  echo $! >"$PID_DIR/$app.pid"
  if [ $? -ne 0 ]; then
    echo "Service $app Start Fail"
  else
    echo "Service $app Background Start"
  fi
}

stop() {
  app=$(basename $1)

  if [ -f "$PID_DIR/$app.pid" ]; then
    check $(cat "$PID_DIR/$1.pid")
    if [ $? != 0 ]; then
      echo "Service $app Not Running."
      return
    fi
    kill $(cat "$PID_DIR/$1.pid")
    rm "$PID_DIR/$app.pid"
    echo "Service $app Stop"
  else
    echo "Service $app Not Running."
  fi
}

pt() {
  for service in $(find -type f -name "$SERVICES"); do
    app=$(basename $service)
    if [ -f "$PID_DIR/$app.pid" ]; then
      check $(cat "$PID_DIR/$service.pid")
      if [ $? == 0 ]; then
        echo "Service $app is Running"
      fi
    fi
  done
}

# 如果没有提供参数，则启动或停止所有服务
if [ $# -eq 0 ]; then
  echo "Start All Services"
  # 启动所有服务
  for service in $(find -type f -name "$SERVICES"); do
    start "$service"
  done
else
  # 遍历所有参数
  while getopts ":ls:q:" opt; do
    case $opt in
    s)
      for service in $(find -type f -name "$OPTARG"); do
        start "$service"
      done
      break
      ;;
    q)
      for service in $(find -type f -name "$OPTARG"); do
        stop "$service"
      done
      break
      ;;
    l)
      pt
      ;;
    \?)
      echo "无效的参数: -$OPTARG"
      exit 1
      ;;
    :)
      if [ "$OPTARG" == "s" ]; then
        ACTION="S"
      else
        ACTION="Q"
      fi
      break
      ;;
    esac
  done

  if [ -n "$ACTION" ]; then
    if [ "$ACTION" == "S" ]; then
      echo "Start All Services"
      for service in $(find -type f -name "$SERVICES"); do
        start "$service"
      done
    else
      echo "Stop All Services"
      for service in $(find -type f -name "$SERVICES"); do
        stop "$service"
      done
    fi
  fi
fi
