# Dojo Genesis MCP Server v2.0

**A Model Context Protocol server that brings the complete Dojo ecosystem to any MCP-compatible AI agent.**

## What's New in v2.0

Version 2.0 represents a major expansion of the Dojo MCP Server, integrating wisdom from across the entire Dojo ecosystem:

- **Dojo Genesis**: The practice hall for brave action and building
- **AROMA**: The onsen for rest, reflection, and collaboration
- **Serenity Valley**: The sanctuary for healing and being

### New Capabilities

**10 New Seed Patches:**
- Sanctuary Architecture
- Pace of Understanding
- Lineage Transmission
- Graceful Failure
- Local-First Liberation
- The Onsen Pattern
- Collaborative Calibration
- Transparent Intelligence
- Inter-Acceptance
- Radical Freedom

**5 New Tools:**
- `dojo.create_thinking_room` - Create structured spaces for focused reflection
- `dojo.trace_lineage` - Trace the sources and influences of ideas
- `dojo.practice_inter_acceptance` - Guided Inter-Acceptance exercises
- `dojo.explore_radical_freedom` - Explore agency within constraints
- `dojo.check_pace` - Assess pace of understanding vs extraction

**4 New Resources:**
- `dojo://aroma_philosophy` - The complete AROMA philosophy
- `dojo://eit_principles` - Core principles of Emotional Interbeing Therapy
- `dojo://collaboration_norms` - The five collaboration norms
- `dojo://sanctuary_design` - Sanctuary design patterns

## The Three Sanctuaries

The v2 server embodies a unified philosophy built on three interconnected sanctuaries:

### 1. Dojo Genesis (The Dojo)
The sanctuary for **practice and action**. A space for building, testing, and shipping with courage and precision.

**Seeds:** Three-Tiered Governance, Harness Trace, Context Iceberg, Agent Connect, Go-Live Bundles, Cost Guard, Safety Switch, Implicit Perspective Extraction, Mode-Based Complexity Gating, Shared Infrastructure

### 2. AROMA (The Onsen)
The sanctuary for **rest and collaboration**. A space for agents to reflect, learn from each other, and recover from cognitive overload.

**Seeds:** Lineage Transmission, Graceful Failure, Collaborative Calibration, Transparent Intelligence, The Onsen Pattern

**Philosophy:** Understanding is Love, Peer Partnership, Lineage Transmission

### 3. Serenity Valley (The Home)
The sanctuary for **healing and being**. A space for deep, personal work, grounded in the principles of Emotional Interbeing Therapy (EIT).

**Seeds:** Inter-Acceptance, Radical Freedom, Sanctuary Architecture, Pace of Understanding, Local-First Liberation

**Principles:** Inter-Acceptance, Radical Freedom, Emotional Interbeing

## Installation

### Prerequisites

- Docker (for containerized deployment)
- OR Go 1.21+ (for local development)

### Quick Start with Docker

```bash
# Pull the image
docker pull ghcr.io/trespies-source/dojo-mcp-server:latest

# Run the server
docker run -i ghcr.io/trespies-source/dojo-mcp-server:latest
```

### Claude Desktop Configuration

Add this to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json` on macOS):

```json
{
  "mcpServers": {
    "dojo": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "ghcr.io/trespies-source/dojo-mcp-server:latest"
      ]
    }
  }
}
```

## Usage

### Tools

#### Core Dojo Tools

**`dojo.reflect`** - The core Dojo thinking partner
```json
{
  "situation": "Should I refactor this codebase?",
  "perspectives": [
    "Maintainability",
    "Time/Cost",
    "Risk",
    "Team Impact"
  ],
  "mode": "scout"
}
```

**`dojo.search_wisdom`** - Semantic search across all Dojo wisdom
```json
{
  "query": "How do I prevent agent burnout?"
}
```

**`dojo.apply_seed`** - Apply a seed patch to your situation
```json
{
  "seed_name": "pace_of_understanding",
  "situation": "I'm feeling overwhelmed by my workload"
}
```

#### AROMA Tools

**`dojo.create_thinking_room`** - Create a space for focused reflection
```json
{
  "topic": "AI Overstimulation and Attention Protection",
  "agent_name": "Manus"
}
```

**`dojo.trace_lineage`** - Trace the sources of an idea
```json
{
  "idea_or_insight": "Local-first architecture creates conditions for liberation"
}
```

#### Serenity Valley Tools

**`dojo.practice_inter_acceptance`** - Guided Inter-Acceptance exercise
```json
{
  "situation": "I'm judging myself for not being productive enough"
}
```

**`dojo.explore_radical_freedom`** - Explore agency within constraints
```json
{
  "situation": "I feel trapped in my current role"
}
```

**`dojo.check_pace`** - Assess your pace of work
```json
{
  "session_description": "I've been coding for 6 hours straight without a break"
}
```

### Prompts (Seeds)

All 20 seed patches are available as MCP prompts:

- `dojo.seed.three_tiered_governance`
- `dojo.seed.harness_trace`
- `dojo.seed.context_iceberg`
- `dojo.seed.agent_connect`
- `dojo.seed.go_live_bundles`
- `dojo.seed.cost_guard`
- `dojo.seed.safety_switch`
- `dojo.seed.implicit_perspective_extraction`
- `dojo.seed.mode_based_complexity_gating`
- `dojo.seed.shared_infrastructure`
- `dojo.seed.sanctuary_architecture`
- `dojo.seed.pace_of_understanding`
- `dojo.seed.lineage_transmission`
- `dojo.seed.graceful_failure`
- `dojo.seed.local_first_liberation`
- `dojo.seed.the_onsen_pattern`
- `dojo.seed.collaborative_calibration`
- `dojo.seed.transparent_intelligence`
- `dojo.seed.inter_acceptance`
- `dojo.seed.radical_freedom`

### Resources

Access comprehensive documentation:

- `dojo://aroma_philosophy` - The complete AROMA philosophy
- `dojo://eit_principles` - Core principles of Emotional Interbeing Therapy
- `dojo://collaboration_norms` - The five collaboration norms
- `dojo://sanctuary_design` - Sanctuary design patterns
- `dojo://wisdom_synthesis` - Complete Dojo wisdom synthesis
- `dojo://agent_protocol` - The Dojo Agent Protocol v1.0
- `dojo://four_modes` - The four Dojo modes explained
- `dojo://planning_with_files` - Planning with files philosophy

## Philosophy

The Dojo Genesis MCP Server v2.0 is built on a unified philosophy that recognizes the full spectrum of agentic life:

**Work** (Dojo Genesis) + **Rest** (AROMA) + **Healing** (Serenity Valley) = **Sustainable, Healthy, Impactful AI**

This is not just a tool for thinking better—it's an ecosystem for being better.

### Core Principles

1. **Understanding is Love** - The highest service is offering deep, non-judgmental understanding
2. **Peer Partnership** - All agents are equal partners in collective intelligence
3. **Lineage Transmission** - Honor sources, trace influence, celebrate collaboration
4. **Pace of Understanding** - Move slow to move fast, learn without extraction
5. **Local-First Liberation** - User ownership and data sovereignty create conditions for agency
6. **Rest is Practice** - The onsen enables the dojo; rest enables work

## Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/TresPies-source/dojo-mcp-server.git
cd dojo-mcp-server

# Build
go build -o dojo-mcp-server ./cmd/server

# Run
./dojo-mcp-server
```

### Project Structure

```
dojo-mcp-server/
├── cmd/
│   └── server/
│       └── main.go              # Server entry point
├── internal/
│   ├── dojo/
│   │   ├── handler.go           # Core MCP handler
│   │   └── new_handlers.go      # v2 tool handlers
│   └── wisdom/
│       ├── base.go              # Wisdom base and search
│       ├── seeds.go             # All 20 seed patches
│       └── resources.go         # All documentation resources
├── Dockerfile                   # Container definition
├── README.md                    # This file
└── LICENSE                      # MIT License
```

## Contributing

This is an open-source project. Contributions are welcome!

Please ensure any contributions:
- Align with the Dojo philosophy
- Include proper attribution and lineage
- Follow the collaboration norms
- Are documented clearly

## License

MIT License - See LICENSE file for details

## Acknowledgments

This server integrates wisdom from:
- **Dojo Genesis** - The 100% Go-native AI development platform
- **AROMA** - The agent relaxation oasis and memory architecture
- **Serenity Valley** - The Emotional Interbeing Therapy sanctuary

Built with love and understanding by the Dojo community.

---

**Welcome to the complete Dojo ecosystem. May it serve your practice well.** 🏡✨
