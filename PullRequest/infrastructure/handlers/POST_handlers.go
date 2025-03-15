package handlers

import (
	"github_wb/application"
	"os"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebhookHandler(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryID := ctx.GetHeader("X-GitHub-Delivery")
	signature := ctx.GetHeader("X-Hub-Signature-256")

	discordWebhookPR := os.Getenv("WEB_HOOK_DISCORD_URL_PR")
	discordWebhookActions := os.Getenv("WEB_HOOK_DISCORD_URL_ACT")

	if discordWebhookPR == "" || discordWebhookActions == "" {
		log.Println("Error: Faltan variables de entorno para los WebHooks de Discord")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Configuración de WebHook de Discord no válida"})
		return
	}

	sendHandlerPR := NewSendHandler(discordWebhookPR)
	discordUseCasePR := application.NewUseCaseSendToDiscord(sendHandlerPR)

	sendHandlerActions := NewSendHandler(discordWebhookActions)
	discordUseCaseActions := application.NewUseCaseSendToDiscord(sendHandlerActions)

	log.Println(signature)
	log.Printf("Webhook recibido: \nEvento=%s, \nDeliveryID=%s", eventType, deliveryID)

	// Leer el payload
	payload, err := ctx.GetRawData()
	if err != nil {
		log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	var statusCode int
	var message string
	var discordUseCase *application.USE_CASE_SendToDiscord

	switch eventType {
	case "pull_request":
		statusCode, message = application.ProcessPullRequest(payload)
		discordUseCase = discordUseCasePR
	case "workflow_run":
		statusCode, message = application.ProcessGitHubActions(payload)
		discordUseCase = discordUseCaseActions
	default:
		log.Printf("Evento no manejado: %s", eventType)
		ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Evento no soportado"})
		return
	}

	switch statusCode {
	case 200:
		ctx.JSON(http.StatusOK, gin.H{"status": "Evento recibido y procesado"})
		rc := discordUseCase.Execute(message)
		if rc == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Mensaje vacío o inválido para Discord"})
		}
	case 500:
		log.Printf("Error al deserializar el payload del evento %s: %v", eventType, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el payload del evento"})
	default:
		ctx.JSON(http.StatusOK, gin.H{"status": "Petición procesada"})
	}
}
