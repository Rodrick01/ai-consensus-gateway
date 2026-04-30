# AI Consensus Gateway

A generic, multi-cloud API Gateway that implements the **Dual-LLM Verification** architecture. 

It exposes a REST API that takes any context and objective, passes it to a "Thinker" LLM to generate a solution, and then passes that solution to an "Auditor" LLM to verify, correct, and validate it before returning the final response.

## Features
* **Agnostic & Generic:** Not tied to YANG or networking. It accepts any text prompt for any use case.
* **Multi-Cloud:** Choose which AI acts as the Thinker and which acts as the Auditor per request. Supported providers:
  * Google Gemini (`gemini`)
  * Anthropic Claude (`claude`)
  * OpenAI GPT-4o (`gpt`)
* **SRE Patterns applied:** Built in Go with strict interface segregation, context timeouts, and structured JSON logging.

## Usage

### 1. Set Environment Variables
Export your API keys:
```bash
export GEMINI_API_KEY="your-key"
export ANTHROPIC_API_KEY="your-key"
export OPENAI_API_KEY="your-key"
```

### 2. Run the Gateway
```bash
go run main.go
```
The server will start on port `8080` (or the port defined in `$PORT`).

### 3. API Usage
Send a `POST` request to `/api/v1/consensus` with a JSON payload:

```bash
curl -X POST http://localhost:8080/api/v1/consensus \
  -H "Content-Type: application/json" \
  -d '{
    "context": "The user submitted a Python script that calculates fibonacci recursively.",
    "objective": "Optimize the script using memoization.",
    "thinker": "claude",
    "auditor": "gemini"
  }'
```

**Response:**
```json
{
  "consensus_result": "def fibonacci(n, memo={}):\n    if n in memo:\n        return memo[n]\n..."
}
```

## How it works

1. The **Thinker** (e.g., Claude) receives the context and objective, and generates a proposed solution.
2. The **Auditor** (e.g., Gemini) receives the original context, the objective, and the Thinker's proposal. It acts as a strict reviewer, correcting hallucinations or errors.
3. The gateway returns the audited consensus result.

## Contributing
This gateway is an open-source contribution to the community. Feel free to fork, add more LLM providers, or integrate eBPF/gNMI sensors in your own forks!
