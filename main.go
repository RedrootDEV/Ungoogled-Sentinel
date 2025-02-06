package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// Configuración desde variables de entorno
var (
	checkInterval  = getEnvAsInt("CHECK_INTERVAL_MINUTES", 30) * 60
	webhookURL     = getEnv("WEBHOOK_URL", "https://discord.com/")
	versionFile    = getEnv("VERSION_FILE", "last_version.txt")
	language       = strings.ToLower(getEnv("LANGUAGE", "en"))
	architecture   = strings.ToLower(getEnv("ARCHITECTURE", "x64"))
	packageType    = strings.ToLower(getEnv("PACKAGE_TYPE", "installer"))
)

// Traducciones
var languageStrings = map[string]map[string]string{
	"es": {
		"title":       "🚀 Nueva versión de Ungoogled Chromium disponible",
		"description": "Se ha detectado una nueva versión para Windows: **%s**",
		"releasePage": "🔗 Página de versiones",
		"download":    "⬇️ Descarga directa",
		"footer":      "Monitor de Ungoogled Chromium",
	},
	"en": {
		"title":       "🚀 New Ungoogled Chromium Version Available",
		"description": "Detected new Windows version: **%s**",
		"releasePage": "🔗 Releases Page",
		"download":    "⬇️ Direct Download",
		"footer":      "Ungoogled Chromium Monitor",
	},
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	var value int
	fmt.Sscanf(valueStr, "%d", &value)
	return value
}

func getLatestUngoogledChromiumVersion() (string, string) {
	url := "https://github.com/ungoogled-software/ungoogled-chromium-windows/releases/latest"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Evita seguir la redirección
		},
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusFound { // 302
		location := resp.Header.Get("Location")
		if location != "" {
			parts := strings.Split(strings.TrimRight(location, "/"), "/")
			version := parts[len(parts)-1]
			return version, location
		}
	}

	return "", ""
}


func generateDownloadLink(version string) string {
	packageMap := map[string]string{
		"installer": fmt.Sprintf("installer_%s.exe", architecture),
		"portable":  fmt.Sprintf("windows_%s.zip", architecture),
	}
	filename := fmt.Sprintf("ungoogled-chromium_%s_%s", version, packageMap[packageType])
	return fmt.Sprintf("https://github.com/ungoogled-software/ungoogled-chromium-windows/releases/download/%s/%s", version, filename)
}

func readLastVersion() string {
	data, err := ioutil.ReadFile(versionFile)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}

func writeLastVersion(version string) {
	_ = ioutil.WriteFile(versionFile, []byte(version), 0644)
}

func sendDiscordNotification(version, githubLink, downloadLink string) {
	lang := languageStrings[language]
	packageLabel := map[string]string{
		"installer": "Installer",
		"portable":  "Portable",
	}[packageType]

	payload := map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":       lang["title"],
				"description": fmt.Sprintf(lang["description"], version),
				"color":       3066993,
				"fields": []map[string]string{
					{"name": lang["releasePage"], "value": fmt.Sprintf("[GitHub Releases](%s)", githubLink), "inline": "true"},
					{"name": lang["download"], "value": fmt.Sprintf("[%s %s](%s)", packageLabel, strings.ToUpper(architecture), downloadLink), "inline": "true"},
				},
				"footer": map[string]string{
					"text": lang["footer"],
					"icon_url": "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png",
				},
			},
		},
	}

	data, _ := json.Marshal(payload)
	resp, err := http.Post(webhookURL, "application/json", strings.NewReader(string(data)))
	if err != nil || resp.StatusCode != 204 {
		fmt.Println("Error al enviar la notificación:", err, "Status:", resp.StatusCode)
	}
}

func main() {
	fmt.Println("Starting monitor...")

	for {
		latestVersion, githubLink := getLatestUngoogledChromiumVersion()

		if latestVersion != "" {
			lastVersion := readLastVersion()

			if latestVersion != lastVersion {
				downloadLink := generateDownloadLink(latestVersion)
				sendDiscordNotification(latestVersion, githubLink, downloadLink)
				writeLastVersion(latestVersion)
			}
		}

		time.Sleep(time.Duration(checkInterval) * time.Second)
	}
}

