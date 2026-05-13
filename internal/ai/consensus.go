package ai

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

// ConsensusEngine maneja el debate entre dos LLMs distintos para tomar decisiones seguras.
type ConsensusEngine struct {
	logger  *slog.Logger
	thinker LLMProvider // The primary AI that designs the solution.
	auditor LLMProvider // The strict AI that audits the solution.
}

// NewConsensusEngine initializes and returns a new ConsensusEngine instance.
// It requires a logger and two LLM providers for the Thinker and Auditor roles.
func NewConsensusEngine(logger *slog.Logger, thinker, auditor LLMProvider) *ConsensusEngine {
	return &ConsensusEngine{
		logger:  logger,
		thinker: thinker,
		auditor: auditor,
	}
}

// MaxIterations define el número máximo de veces que Thinker y Auditor iterarán para alcanzar un consenso.
const MaxIterations = 3

// GenerateConsensus ejecuta la arquitectura "Dual LLM Verification" de manera iterativa.
func (ce *ConsensusEngine) GenerateConsensus(ctx context.Context, contextData, objective string) (string, error) {
	ce.logger.InfoContext(ctx, "Iniciando consenso AI Multi-Cloud", 
		slog.String("thinker", ce.thinker.Name()), 
		slog.String("auditor", ce.auditor.Name()),
	)

	var lastProposal string
	var auditorFeedback string

	for i := 1; i <= MaxIterations; i++ {
		ce.logger.InfoContext(ctx, "Iniciando iteración de consenso", slog.Int("iteracion", i))

		// FASE 1: Creación (Thinker)
		var promptThinker string
		if i == 1 {
			promptThinker = fmt.Sprintf(`Analiza el siguiente contexto y resuelve el objetivo solicitado. Devuelve ÚNICAMENTE la respuesta o el código final, sin explicaciones ni formato markdown a menos que se solicite en el objetivo.
Contexto:
%s

Objetivo:
%s`, contextData, objective)
		} else {
			promptThinker = fmt.Sprintf(`Eres la IA principal (Thinker). Has propuesto una solución, pero la IA auditora la ha rechazado con el siguiente feedback.
Corrige tu propuesta según el feedback del auditor. Devuelve ÚNICAMENTE la respuesta o el código final corregido.

Contexto original:
%s

Objetivo original:
%s

Tu propuesta anterior:
%s

Feedback del auditor:
%s`, contextData, objective, lastProposal, auditorFeedback)
		}

		proposedSolution, err := ce.thinker.Ask(ctx, promptThinker)
		if err != nil {
			return "", fmt.Errorf("la IA primaria [%s] falló al pensar en la iteración %d: %w", ce.thinker.Name(), i, err)
		}

		lastProposal = proposedSolution
		ce.logger.DebugContext(ctx, "Propuesta generada por IA primaria", slog.Int("len", len(proposedSolution)), slog.Int("iteracion", i))

		// FASE 2: Auditoría (Auditor)
		promptAuditor := fmt.Sprintf(`Actúa como un auditor experto y estricto. Revisa la siguiente propuesta para cumplir el objetivo dado en base al contexto.
Si la propuesta es correcta, cumple el objetivo y no tiene errores, responde ÚNICAMENTE comenzando con "[APPROVED]" seguido de la propuesta final.
Si la propuesta tiene errores, alucinaciones o rompe buenas prácticas, responde ÚNICAMENTE comenzando con "[REJECTED]" seguido de una breve explicación de qué está mal y cómo la IA principal debe corregirlo.

Contexto original:
%s

Objetivo original:
%s

Propuesta a auditar:
%s`, contextData, objective, proposedSolution)

		auditResponse, err := ce.auditor.Ask(ctx, promptAuditor)
		if err != nil {
			return "", fmt.Errorf("la IA auditora [%s] falló al validar en la iteración %d: %w", ce.auditor.Name(), i, err)
		}

		ce.logger.DebugContext(ctx, "Respuesta del auditor recibida", slog.Int("len", len(auditResponse)), slog.Int("iteracion", i))

		// Analizar la respuesta del Auditor
		auditResponse = strings.TrimSpace(auditResponse)
		if strings.HasPrefix(strings.ToUpper(auditResponse), "[APPROVED]") {
			ce.logger.InfoContext(ctx, "Consenso alcanzado exitosamente", slog.Int("iteraciones_totales", i))
			finalSolution := strings.TrimSpace(auditResponse[len("[APPROVED]"):])
			// En caso de que el auditor haya devuelto solo [APPROVED] sin la propuesta:
			if finalSolution == "" {
				finalSolution = proposedSolution
			}
			return finalSolution, nil
		} else if strings.HasPrefix(strings.ToUpper(auditResponse), "[REJECTED]") {
			auditorFeedback = strings.TrimSpace(auditResponse[len("[REJECTED]"):])
			ce.logger.WarnContext(ctx, "Propuesta rechazada por el auditor, se requiere corrección", slog.Int("iteracion", i))
		} else {
			// El auditor no siguió las instrucciones estrictas. Asumimos rechazo e intentamos extraer el texto como feedback.
			ce.logger.WarnContext(ctx, "El auditor no usó prefijos [APPROVED]/[REJECTED]. Se asume rechazo por seguridad.", slog.Int("iteracion", i))
			auditorFeedback = auditResponse
		}
	}

	ce.logger.ErrorContext(ctx, "No se pudo alcanzar el consenso después de las iteraciones máximas", slog.Int("max_iteraciones", MaxIterations))
	return "", fmt.Errorf("no se pudo alcanzar consenso entre Thinker (%s) y Auditor (%s) despues de %d iteraciones. Último feedback: %s", ce.thinker.Name(), ce.auditor.Name(), MaxIterations, auditorFeedback)
}
