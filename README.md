# Ungoogled Sentinel 🚀

**Ungoogled Sentinel** is a tool that automatically monitors updates for [Ungoogled Chromium for Windows](https://github.com/ungoogled-software/ungoogled-chromium-windows) and sends notifications to a **Discord webhook** with details and download links.

## 📌 Features
- Periodically checks for new versions.
- Supports **x64** and **x86** architectures.
- Sends detailed notifications via **Discord Webhook**.
- Provides direct download links for **Installer** and **Portable** versions.
- Supports **English** and **Spanish** notifications.

## 🐳 Running with Docker Compose

- You can run Ungoogled Sentinel using Docker Compose with the following docker-compose.yml:

   ```sh
	services:
	  ungoogled-sentinel:
		image: redr00t/ungoogled-sentinel:latest
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
1. Install Python (3.7+ recommended).
2. Install dependencies:
   ```sh
   pip install -r requirements.txt
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
   python main.py
   ```

## 📜 License
- This project is licensed under the MIT License.