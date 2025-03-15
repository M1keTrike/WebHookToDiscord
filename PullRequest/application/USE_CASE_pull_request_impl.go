package application

import (
	"encoding/json"
	"fmt"
	"log"

	domain "github_wb/domain/value_objects"
)

func ProcessPullRequest(payload []byte) (int, string) {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		errorMsg := fmt.Sprintf("Error al procesar payload: %v", err)
		log.Println(errorMsg)
		return 500, errorMsg
	}

	action := eventPayload.Action
	title := eventPayload.PullRequest.Title
	content := eventPayload.PullRequest.Body
	user := eventPayload.PullRequest.User.Login
	prURL := eventPayload.PullRequest.URL

	var successMsg string

	switch action {
	case "created":
		successMsg = fmt.Sprintf("📌 **Nuevo Pull Request Creado**\n🔹 Título: %s\n🔹 Usuario: %s\n🔗 [Ver PR](%s)", title, user, prURL)
	case "opened":
		successMsg = fmt.Sprintf("📂 **Pull Request Abierto**\n🔹 Título: %s\n🔹 Usuario: %s\n🔗 [Ver PR](%s)", title, user, prURL)
	case "closed":
		successMsg = fmt.Sprintf("🚫 **Pull Request Cerrado**\n🔹 Título: %s\n🔹 Usuario: %s\n🔗 [Ver PR](%s)", title, user, prURL)
	case "reopened":
		successMsg = fmt.Sprintf("🔄 **Pull Request Reabierto**\n🔹 Título: %s\n🔹 Usuario: %s\n🔗 [Ver PR](%s)", title, user, prURL)
	case "ready_for_review":
		successMsg = fmt.Sprintf("✅ **Pull Request Listo para Revisión**\n🔹 Título: %s\n🔹 Usuario: %s\n📄 Descripción: %s\n🔗 [Ver PR](%s)", title, user, content, prURL)
	case "merged":
		successMsg = fmt.Sprintf("🎉 **Pull Request Fusionado**\n🔹 Título: %s\n🔹 Usuario: %s\n🔗 [Ver PR](%s)", title, user, prURL)
	default:
		infoMsg := fmt.Sprintf("Evento de Pull Request no soportado: %s", action)
		log.Println(infoMsg)
		return 400, infoMsg
	}

	log.Println(successMsg)
	return 200, successMsg
}
