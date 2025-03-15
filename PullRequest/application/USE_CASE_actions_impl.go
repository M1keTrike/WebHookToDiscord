package application

import (
	"encoding/json"
	"fmt"
	"log"

	domain "github_wb/domain/value_objects"
)

func ProcessGitHubActions(payload []byte) (int, string) {
	var eventPayload domain.GitHubActionsEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		errorMsg := fmt.Sprintf("Error al procesar payload: %v", err)
		log.Println(errorMsg)
		return 500, errorMsg
	}

	action := eventPayload.Action
	workflowName := eventPayload.Workflow.Name
	conclusion := eventPayload.WorkflowRun.Conclusion
	repoName := eventPayload.Repository.FullName
	runURL := eventPayload.WorkflowRun.HTMLURL

	var successMsg string

	switch action {
	case "requested":
		successMsg = fmt.Sprintf("🚀 **Nuevo Workflow Iniciado**\n🔹 Repositorio: %s\n🔹 Workflow: %s\n🔗 [Ver ejecución](%s)", repoName, workflowName, runURL)
	case "in_progress":
		successMsg = fmt.Sprintf("⏳ **Workflow en Progreso**\n🔹 Repositorio: %s\n🔹 Workflow: %s\n🔗 [Ver ejecución](%s)", repoName, workflowName, runURL)
	case "completed":
		if conclusion == "success" {
			successMsg = fmt.Sprintf("✅ **Workflow Completado Exitosamente**\n🔹 Repositorio: %s\n🔹 Workflow: %s\n🔗 [Ver ejecución](%s)", repoName, workflowName, runURL)
		} else {
			successMsg = fmt.Sprintf("❌ **Workflow Fallido**\n🔹 Repositorio: %s\n🔹 Workflow: %s\n🔗 [Ver ejecución](%s)", repoName, workflowName, runURL)
		}
	default:
		infoMsg := fmt.Sprintf("Evento de GitHub Actions no soportado: %s", action)
		log.Println(infoMsg)
		return 400, infoMsg
	}

	log.Println(successMsg)
	return 200, successMsg
}
