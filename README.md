# Ungoogled Sentinel 🚀

**Ungoogled Sentinel** is a tool that automatically monitors updates for [Ungoogled Chromium for Windows](https://github.com/ungoogled-software/ungoogled-chromium-windows) and sends notifications to a **Discord webhook** with details and download links.

## 📌 Features
- Periodically checks for new versions.
- Supports **x64** and **x86** architectures.
- Sends detailed notifications via **Discord Webhook**.
- Provides direct download links for **Installer** and **Portable** versions.
- Supports **English** and **Spanish** notifications.
- Written in **Go** for high performance and low resource consumption.

## 🐳 Running with Docker Compose

- You can run Ungoogled Sentinel using Docker Compose with the following docker-compose.yml:

   ```sh
	services:
	  ungoogled-sentinel:
		image: redr00t/ungoogled-sentinel:latest
		container_name: ungoogled-sentinel
		environment:
		  - CHECK_INTERVAL_MINUTES=30
		  - WEBHOOK_URL=https://discord.com/api/webhooks/..............
		  - VERSION_FILE=/data/last_version.txt
		  - LANGUAGE=en
		  - ARCHITECTURE=x64
		  - PACKAGE_TYPE=installer
		volumes:
		  - ./Ungoogled-Sentinel/data:/data
		restart: always
   ```

## ⚙️ Installation (Manual)
1. Install Go (1.21+ recommended).
2. Clone the repository:
   ```sh
	git clone https://github.com/RedrootDEV/Ungoogled-Sentinel.git
	cd Ungoogled-Sentinel
   ```
2. Build the application:
   ```sh
   go build -o ungoogled-sentinel main.go
   ```

3. Set environment variables:
- **WEBHOOK_URL** → Your Discord webhook URL.
- **CHECK_INTERVAL_MINUTES** → Interval in minutes (default: 30).
- **VERSION_FILE** → File to store the last detected version (default: last_version.txt).
- **LANGUAGE** → en (English) or es (Spanish) (default: en).
- **ARCHITECTURE** → x64 or x86 (default: x64).
- **PACKAGE_TYPE** → installer or portable (default: installer).

## 🚀 Usage (Manual)

- Run the script manually with:
   ```sh
   ./ungoogled-sentinel
   ```

## 📜 License
- This project is licensed under the [MIT License](https://github.com/RedrootDEV/Ungoogled-Sentinel/blob/main/LICENSE)