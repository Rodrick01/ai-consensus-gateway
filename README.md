# AI Consensus Gateway

A professional, multi-cloud API Gateway that implements the **Dual-LLM Verification** architecture. 

It exposes a REST API that takes any context and objective, passes it to a "Thinker" LLM to generate a solution, and then passes that solution to an "Auditor" LLM to verify, correct, and validate it before returning the final response.

## Features
* **Agnostic & Generic:** Not tied to specific use cases. It accepts any text prompt.
* **Multi-Cloud:** Choose which AI acts as the Thinker and which acts as the Auditor per request. Supported providers: Google Gemini (`gemini`), Anthropic Claude (`claude`), OpenAI GPT-4o (`gpt`).
* **SRE Ready:** Built in Go with strict interface segregation, context timeouts, structured JSON logging, and Kubernetes-ready `/health` probes.

## Usage

### 1. Set Environment Variables
```bash
export GEMINI_API_KEY="your-key"
export ANTHROPIC_API_KEY="your-key"
export OPENAI_API_KEY="your-key"
```

### 2. Run the Gateway
```bash
go run main.go
```

### 3. API Endpoints
- `GET /health`: Liveness/Readiness probe.
- `POST /api/v1/consensus`: The main consensus endpoint.
```json
{
  "context": "Raw data goes here",
  "objective": "What you want to achieve",
  "thinker": "claude",
  "auditor": "gemini"
}
```
