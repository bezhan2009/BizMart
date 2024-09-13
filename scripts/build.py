import os
import time
import signal
import subprocess
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler

root_dir = r'C:\Users\Admin\GolandProjects\BizMart'
server_process = None

class FileChangeHandler(FileSystemEventHandler):
    def __init__(self, restart_callback):
        super().__init__()
        self.restart_callback = restart_callback

    def on_modified(self, event):
        # Отслеживаем только изменения в .go файлах
        if event.src_path.endswith('.go'):
            print(f"Detected changes in: {event.src_path}")
            self.restart_callback()

def kill_server():
    """Функция для завершения процесса сервера, если он запущен."""
    global server_process
    if server_process and server_process.poll() is None:
        print(f"Killing server with PID: {server_process.pid}")
        server_process.terminate()
        server_process.wait()
        server_process = None

def start_server():
    """Функция для компиляции и запуска Go сервера."""
    global server_process
    print("Building and starting the server...")

    # Инициализация документации с помощью swag, если есть изменения
    swag_file = os.path.join(root_dir, "docs", "swagger.json")
    if not os.path.exists(swag_file):
        subprocess.run(["swag", "init"], cwd=root_dir)
    else:
        print("Skipping swag init as docs already exist.")

    # Компиляция Go файла
    subprocess.run(["go", "build", "-o", "server.exe", "main.go"], cwd=root_dir)

    # Запуск сервера
    server_process = subprocess.Popen([os.path.join(root_dir, "server.exe")], cwd=root_dir)
    print(f"Server started with PID: {server_process.pid}")

def monitor_changes():
    """Функция для мониторинга изменений с использованием watchdog."""
    event_handler = FileChangeHandler(restart_server)
    observer = Observer()
    observer.schedule(event_handler, path=root_dir, recursive=True)
    observer.start()

    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        observer.stop()
        kill_server()  # Завершаем сервер при Ctrl+C
    observer.join()

def restart_server():
    """Функция для перезапуска сервера."""
    print("Change detected, restarting server...")
    kill_server()
    start_server()

if __name__ == "__main__":
    if not os.path.exists(root_dir):
        print(f"ERROR: Directory {root_dir} does not exist!")
        exit(1)

    def signal_handler(sig, frame):
        """Обработка сигнала Ctrl+C для корректного завершения."""
        print("\nCleaning up...")
        kill_server()
        exit(0)

    # Устанавливаем обработчик Ctrl+C
    signal.signal(signal.SIGINT, signal_handler)

    try:
        start_server()
        monitor_changes()
    except Exception as e:
        print(f"An error occurred: {e}")
        kill_server()
