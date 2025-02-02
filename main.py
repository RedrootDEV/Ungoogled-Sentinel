import requests
import time
import os

# Envs
CHECK_INTERVAL_MINUTES = int(os.getenv("CHECK_INTERVAL_MINUTES", 30))
WEBHOOK_URL = os.getenv("WEBHOOK_URL", "https://discord.com/")
VERSION_FILE = os.getenv("VERSION_FILE", "last_version.txt")
LANGUAGE = os.getenv("LANGUAGE", "en").lower()  # 'es' o 'en'
ARCHITECTURE = os.getenv("ARCHITECTURE", "x64").lower()  # 'x64' o 'x86'
PACKAGE_TYPE = os.getenv("PACKAGE_TYPE", "installer").lower()  # 'installer' o 'portable'

# Language
LANGUAGE_STRINGS = {
    "es": {
        "title": "🚀 Nueva versión de Ungoogled Chromium disponible",
        "description": "Se ha detectado una nueva versión para Windows: **{version}**",
        "release_page": "🔗 Página de versiones",
        "download": "⬇️ Descarga directa",
        "package_types": {
            "installer": "Instalador",
            "portable": "Portable"
        },
        "footer": "Monitor de Ungoogled Chromium"
    },
    "en": {
        "title": "🚀 New Ungoogled Chromium Version Available",
        "description": "Detected new Windows version: **{version}**",
        "release_page": "🔗 Releases Page",
        "download": "⬇️ Direct Download",
        "package_types": {
            "installer": "Installer",
            "portable": "Portable"
        },
        "footer": "Ungoogled Chromium Monitor"
    }
}

def get_latest_ungoogled_chromium_version():
    url = "https://github.com/ungoogled-software/ungoogled-chromium-windows/releases/latest"
    try:
        response = requests.get(url, allow_redirects=False, timeout=10)
        if response.status_code == 302:
            latest_url = response.headers.get("Location")
            if latest_url:
                version = latest_url.rstrip("/").split("/")[-1]
                return version, latest_url
        return None, None
    except requests.RequestException as e:
        print(f"Error en la petición: {e}")
        return None, None

def generate_download_link(version):
    package_map = {
        "installer": f"installer_{ARCHITECTURE}.exe",
        "portable": f"windows_{ARCHITECTURE}.zip"
    }
    
    filename = f"ungoogled-chromium_{version}_{package_map[PACKAGE_TYPE]}"
    return f"https://github.com/ungoogled-software/ungoogled-chromium-windows/releases/download/{version}/{filename}"

def read_last_version():
    if os.path.exists(VERSION_FILE):
        with open(VERSION_FILE, "r") as file:
            return file.read().strip()
    return None

def write_last_version(version):
    with open(VERSION_FILE, "w") as file:
        file.write(version)

def send_discord_notification(version, github_link, download_link):
    lang = LANGUAGE_STRINGS[LANGUAGE]
    package_str = lang["package_types"][PACKAGE_TYPE]
    
    embed = {
        "title": lang["title"],
        "description": lang["description"].format(version=version),
        "color": 3066993,
        "fields": [
            {
                "name": lang["release_page"],
                "value": f"[GitHub Releases]({github_link})",
                "inline": True
            },
            {
                "name": lang["download"],
                "value": f"[{package_str} {ARCHITECTURE.upper()}]({download_link})",
                "inline": True
            }
        ],
        "footer": {
            "text": lang["footer"],
            "icon_url": "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png"
        }
    }
    
    response = requests.post(WEBHOOK_URL, json={"embeds": [embed]})
    if response.status_code != 204:
        print(f"Notification error: {response.status_code}")

def main():
    print("Starting monitor...")
    while True:
        latest_version, github_link = get_latest_ungoogled_chromium_version()
        
        if latest_version:
            last_version = read_last_version()
            if latest_version != last_version:
                download_link = generate_download_link(latest_version)
                send_discord_notification(latest_version, github_link, download_link)
                write_last_version(latest_version)
        
        time.sleep(CHECK_INTERVAL_MINUTES * 60)

if __name__ == "__main__":
    main()