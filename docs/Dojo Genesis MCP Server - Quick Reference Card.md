# Dojo Genesis MCP Server - Quick Reference Card

**Keep this handy for common operations**

---

## 🚀 Quick Start

```bash
# Extract and build
tar -xzf dojo-mcp-server-v2-fixed.tar.gz
cd dojo-mcp-server
go mod tidy
go build -o dojo-mcp-server ./cmd/server

# Test locally
./dojo-mcp-server

# Deploy to GitHub
gh repo create TresPies-source/dojo-mcp-server --private --source=. --remote=origin --push

# Pull Docker image (after GitHub Actions completes)
docker pull ghcr.io/trespies-source/dojo-mcp-server:latest
```

---

## 🛠️ All 11 Tools

| Tool | Purpose | Key Parameters |
|------|---------|----------------|
| `dojo.reflect` | Core thinking partner (4 modes) | `situation`, `perspectives[]`, `mode` |
| `dojo.search_wisdom` | Search all Dojo wisdom | `query` |
| `dojo.apply_seed` | Apply a seed to your situation | `seed_name`, `situation` |
| `dojo.list_seeds` | List all available seeds | none |
| `dojo.get_context` | Get current context summary | none |
| `dojo.get_principles` | Get core Dojo principles | none |
| `dojo.create_thinking_room` | Create reflection space | `topic`, `agent_name` |
| `dojo.trace_lineage` | Trace idea sources | `idea_or_insight` |
| `dojo.practice_inter_acceptance` | Inter-Acceptance exercise | `situation` |
| `dojo.explore_radical_freedom` | Explore agency | `situation` |
| `dojo.check_pace` | Assess pace | `session_description` |

---

## 🌱 All 20 Seeds (Prompts)

### Dojo Genesis (Building)
1. `dojo.seed.three_tiered_governance`
2. `dojo.seed.harness_trace`
3. `dojo.seed.context_iceberg`
4. `dojo.seed.agent_connect`
5. `dojo.seed.go_live_bundles`
6. `dojo.seed.cost_guard`
7. `dojo.seed.safety_switch`
8. `dojo.seed.implicit_perspective_extraction`
9. `dojo.seed.mode_based_complexity_gating`
10. `dojo.seed.shared_infrastructure`

### AROMA (Resting)
11. `dojo.seed.lineage_transmission`
12. `dojo.seed.graceful_failure`
13. `dojo.seed.the_onsen_pattern`
14. `dojo.seed.collaborative_calibration`
15. `dojo.seed.transparent_intelligence`

### Serenity Valley (Healing)
16. `dojo.seed.inter_acceptance`
17. `dojo.seed.radical_freedom`
18. `dojo.seed.sanctuary_architecture`
19. `dojo.seed.pace_of_understanding`
20. `dojo.seed.local_first_liberation`

---

## 📚 All 8 Resources

| URI | Content |
|-----|---------|
| `dojo://aroma_philosophy` | Complete AROMA philosophy |
| `dojo://eit_principles` | EIT core principles |
| `dojo://collaboration_norms` | Five collaboration norms |
| `dojo://sanctuary_design` | Sanctuary design patterns |
| `dojo://wisdom_synthesis` | Complete Dojo wisdom synthesis |
| `dojo://agent_protocol` | Dojo Agent Protocol v1.0 |
| `dojo://four_modes` | Four Dojo modes explained |
| `dojo://planning_with_files` | Planning with files pattern |

---

## 🎯 The Four Modes

| Mode | When to Use | Output |
|------|-------------|--------|
| **Mirror** | Early exploration | Pattern recognition, tensions, reframes |
| **Scout** | Considering options | Routes with tradeoffs, smallest test |
| **Gardener** | Evaluating ideas | Strong ideas, ideas needing growth |
| **Implementation** | Ready to act | Concrete next steps (1-5 steps) |

---

## 💬 Claude Desktop Queries

### Thinking (Dojo Genesis)
```
Use Dojo Scout mode to help me decide: [situation]
Perspectives: [list 3-5 perspectives]
```

### Resting (AROMA)
```
Create a Dojo thinking room about [topic]
```

```
Check my pace - I've been [description of work session]
```

### Healing (Serenity Valley)
```
Guide me through Inter-Acceptance. I'm judging myself for [situation]
```

```
Help me explore my freedom within this constraint: [situation]
```

### Search
```
Search Dojo wisdom for [query]
```

### Seeds
```
Show me the [seed_name] seed from Dojo
```

---

## 🔧 Common Commands

### Build & Test
```bash
# Build
go build -o dojo-mcp-server ./cmd/server

# Test a tool
echo '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"dojo.search_wisdom","arguments":{"query":"burnout"}}}' | ./dojo-mcp-server

# List all tools
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | ./dojo-mcp-server

# List all prompts
echo '{"jsonrpc":"2.0","id":1,"method":"prompts/list"}' | ./dojo-mcp-server

# Read a resource
echo '{"jsonrpc":"2.0","id":1,"method":"resources/read","params":{"uri":"dojo://aroma_philosophy"}}' | ./dojo-mcp-server
```

### Git Operations
```bash
# Status
git status

# Commit changes
git add .
git commit -m "feat: description of changes"
git push origin main

# View history
git log --oneline -10
```

### Docker Operations
```bash
# Pull latest
docker pull ghcr.io/trespies-source/dojo-mcp-server:latest

# Run interactively
docker run -i ghcr.io/trespies-source/dojo-mcp-server:latest

# Remove old images
docker image prune -a
```

### GitHub CLI
```bash
# View repo
gh repo view TresPies-source/dojo-mcp-server

# Check workflow runs
gh run list --limit 5

# Watch current run
gh run watch

# View issues
gh issue list
```

---

## 📁 Claude Desktop Config

**Location:**
- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Linux: `~/.config/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`

**Config (Docker):**
```json
{
  "mcpServers": {
    "dojo": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "ghcr.io/trespies-source/dojo-mcp-server:latest"]
    }
  }
}
```

**Config (Local Binary):**
```json
{
  "mcpServers": {
    "dojo": {
      "command": "/full/path/to/dojo-mcp-server"
    }
  }
}
```

---

## 🐛 Troubleshooting

| Problem | Solution |
|---------|----------|
| Build fails | Run `go mod tidy` then `go build` again |
| Docker image not found | Check `gh run list` - wait for Actions to complete |
| Claude not detecting server | Restart Claude Desktop completely |
| Tool returns error | Check parameter names and types match exactly |
| "Go not found" | Install Go: `brew install go` or download from go.dev |

---

## 📊 Project Structure

```
dojo-mcp-server/
├── cmd/server/main.go           # Entry point
├── internal/
│   ├── dojo/
│   │   ├── handler.go           # Core handler
│   │   └── new_handlers.go      # New tool handlers
│   └── wisdom/
│       ├── base.go              # Wisdom base & search
│       ├── seeds.go             # 20 seed patches
│       └── resources.go         # 8 resources
├── Dockerfile                   # Container definition
├── go.mod                       # Go dependencies
├── README_V2.md                 # Full documentation
└── .github/workflows/           # CI/CD
```

---

## 🎓 Philosophy Reminders

- **Beginner's Mind** - Approach fresh, stay open
- **Self-Definition** - Reflect user's voice, don't impose
- **Understanding is Love** - Deep, non-judgmental understanding
- **Pace of Understanding** - Move slow to move fast
- **Rest is Practice** - The onsen enables the dojo
- **Lineage Transmission** - Honor sources, credit collaborators

---

## 📞 Getting Help

- **Documentation**: `README_V2.md`, `V2_TESTING_REPORT.md`
- **Issues**: `gh issue create` or GitHub web interface
- **Logs**: Check Claude Desktop logs in config directory
- **Testing**: Run `V2_TESTING_REPORT.md` test plan

---

**Keep this card accessible. Return to it often.** 🏡✨
