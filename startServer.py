import subprocess
import os
import time
import datetime
import platform

def start_server(session_id):
    print(f"Starting Go server for session {session_id}...")
    run_command = ["go", "run", "main.go"]

    if platform.system() == 'Windows':
        subprocess.Popen(["start", "cmd", "/c"] + run_command, cwd=os.getcwd())
    elif platform.system() == 'Darwin':  # Для macOS
        subprocess.Popen(["open", "-a", "Terminal"] + run_command, cwd=os.getcwd())
    else:
        subprocess.Popen(run_command, cwd=os.getcwd())

# Остальной код остается без изменений...


def check_server(session_id, auto_restart):
    try:
        subprocess.check_output(["curl", "http://localhost:8080"], timeout=5)
        print("Server is running.")
        log_action(session_id, "Server is running.")
    except (subprocess.CalledProcessError, subprocess.TimeoutExpired):
        print("Server is down.")
        log_action(session_id, "Server is down.")
        if auto_restart:
            start_server(session_id)
            log_action(session_id, "Server restarted.")

def log_action(session_id, action):
    with open("logs.txt", "a") as logfile:
        current_time = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        logfile.write(f"{current_time} / id({session_id}) / {action}\n")

if __name__ == "__main__":
    auto_restart = input("Do you want the server to restart automatically? (y/n): ").strip().lower() == "y"
    session_id = input("Enter session ID: ").strip()

    start_server(session_id)
    log_action(session_id, "Server started.")

    try:
        while True:
            check_server(session_id, auto_restart)
            time.sleep(5)
    except KeyboardInterrupt:
        print("Program terminated.")
