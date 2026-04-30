# AI Consensus Gateway

Un API Gateway profesional y multi-nube que implementa la arquitectura de **Verificación Dual-LLM**.

Expone una API REST que toma cualquier contexto y objetivo, se lo pasa a un LLM "Pensador" para generar una solución, y luego le pasa esa solución a un LLM "Auditor" para verificarla, corregirla y validarla antes de devolver la respuesta final.

## Características
* **Agnóstico y Genérico:** No está atado a casos de uso específicos. Acepta cualquier prompt de texto.
* **Multi-Cloud:** Elige qué IA actúa como Pensador y cuál como Auditor por petición. Proveedores soportados: Google Gemini (`gemini`), Anthropic Claude (`claude`), OpenAI GPT-4o (`gpt`).
* **Listo para SRE:** Construido en Go con segregación estricta de interfaces, timeouts de contexto, logs estructurados en JSON y endpoints `/health` listos para Kubernetes.

## Uso

### 1. Variables de Entorno
```bash
export GEMINI_API_KEY="tu-key"
export ANTHROPIC_API_KEY="tu-key"
export OPENAI_API_KEY="tu-key"
```

### 2. Ejecutar el Gateway
```bash
go run main.go
```

### 3. Endpoints de la API
- `GET /health`: Sonda de Liveness/Readiness.
- `POST /api/v1/consensus`: El endpoint principal de consenso.
```json
{
  "context": "Tus datos crudos van aquí",
  "objective": "Lo que quieres lograr",
  "thinker": "claude",
  "auditor": "gemini"
}
```
