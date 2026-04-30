# AI Consensus Gateway (Mākslīgā intelekta konsensa vārteja)

Profesionāla vairāku mākoņu (multi-cloud) API vārteja, kas ievieš **Dual-LLM verifikācijas** arhitektūru.

Tā nodrošina REST API, kas pieņem jebkuru kontekstu un mērķi, nodod to "Domātāja" (Thinker) LLM, lai radītu risinājumu, un pēc tam nodod šo risinājumu "Revidenta" (Auditor) LLM, lai to pārbaudītu, labotu un apstiprinātu pirms galīgās atbildes atgriešanas.

## Funkcijas
* **Agnostisks un ģenērisks:** Nav piesaistīts konkrētiem lietošanas gadījumiem. Pieņem jebkuru teksta uzvedni.
* **Multi-Cloud:** Katram pieprasījumam izvēlieties, kurš AI darbojas kā Domātājs un kurš kā Revidents. Atbalstītie pakalpojumu sniedzēji: Google Gemini (`gemini`), Anthropic Claude (`claude`), OpenAI GPT-4o (`gpt`).
* **Sagatavots SRE:** Izveidots Go valodā ar stingru saskarņu nodalīšanu, konteksta noildzēm (timeouts), strukturētu JSON reģistrēšanu (logging) un Kubernetes gatavām `/health` zondēm.

## Lietošana

### 1. Iestatiet vides mainīgos
```bash
export GEMINI_API_KEY="tava-atslēga"
export ANTHROPIC_API_KEY="tava-atslēga"
export OPENAI_API_KEY="tava-atslēga"
```

### 2. Palaidiet vārteju
```bash
go run main.go
```

### 3. API galapunkti (Endpoints)
- `GET /health`: Dzīvotspējas un gatavības (Liveness/Readiness) zonde.
- `POST /api/v1/consensus`: Galvenais konsensa galapunkts.
```json
{
  "context": "Neapstrādāti dati",
  "objective": "Ko jūs vēlaties sasniegt",
  "thinker": "claude",
  "auditor": "gemini"
}
```
