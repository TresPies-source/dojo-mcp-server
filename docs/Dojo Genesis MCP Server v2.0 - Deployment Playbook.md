# Dojo Genesis MCP Server v2.0 - Deployment Playbook

**Your Complete Guide to Initialization, Usage, and Deployment**

---

## 📋 Overview

This playbook will guide you through:
1. **Initialization** - Setting up the project locally
2. **Testing** - Validating functionality
3. **Deployment** - Publishing to GitHub and Docker
4. **Usage** - Integrating with Claude Desktop
5. **Maintenance** - Ongoing operations

**Estimated Time:** 45-60 minutes for complete setup

---

## Phase 1: Local Initialization (15 minutes)

### Step 1.1: Extract the Project

```bash
# Navigate to your projects directory
cd ~/Projects  # or wherever you keep your code

# Extract the tarball
tar -xzf ~/path/to/dojo-mcp-server-v2-fixed.tar.gz

# Navigate into the project
cd dojo-mcp-server

# Verify the structure
ls -la
```

**Expected Output:**
```
cmd/
internal/
Dockerfile
go.mod
README_V2.md
V2_TESTING_REPORT.md
LICENSE
.gitignore
```

---

### Step 1.2: Verify Go Installation

```bash
# Check Go version (requires 1.21+)
go version

# If Go is not installed, install it:
# macOS:
brew install go

# Linux:
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

---

### Step 1.3: Initialize Go Module

```bash
# Initialize the module (if not already done)
go mod init github.com/TresPies-source/dojo-mcp-server

# Download dependencies
go mod tidy

# Verify dependencies
go mod verify
```

**Expected Output:**
```
go: downloading github.com/mark3labs/mcp-go v0.x.x
go: downloading ...
all modules verified
```

---

### Step 1.4: Build the Server

```bash
# Build the binary
go build -o dojo-mcp-server ./cmd/server

# Verify the build
./dojo-mcp-server --help  # or just run it to see if it starts
```

**Expected Output:**
```
# The server should start and listen on stdio
# Press Ctrl+C to stop
```

---

## Phase 2: Local Testing (15 minutes)

### Step 2.1: Manual Tool Testing

Create a test script to verify tool functionality:

```bash
# Create a test file
cat > test_tools.sh << 'EOF'
#!/bin/bash

echo "Testing dojo.reflect..."
echo '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"dojo.reflect","arguments":{"situation":"Should I refactor?","perspectives":["Maintainability","Cost","Risk"],"mode":"mirror"}}}' | ./dojo-mcp-server

echo -e "\n\nTesting dojo.search_wisdom..."
echo '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"dojo.search_wisdom","arguments":{"query":"burnout prevention"}}}' | ./dojo-mcp-server

echo -e "\n\nTesting dojo.create_thinking_room..."
echo '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"dojo.create_thinking_room","arguments":{"topic":"AI Ethics","agent_name":"Manus"}}}' | ./dojo-mcp-server
EOF

chmod +x test_tools.sh
./test_tools.sh
```

**Expected Output:** JSON responses with formatted content for each tool

---

### Step 2.2: Verify All Resources

```bash
# Test resource access
echo '{"jsonrpc":"2.0","id":4,"method":"resources/read","params":{"uri":"dojo://aroma_philosophy"}}' | ./dojo-mcp-server

echo '{"jsonrpc":"2.0","id":5,"method":"resources/read","params":{"uri":"dojo://eit_principles"}}' | ./dojo-mcp-server
```

**Expected Output:** Full markdown content for each resource

---

### Step 2.3: Verify All Prompts

```bash
# List all prompts
echo '{"jsonrpc":"2.0","id":6,"method":"prompts/list"}' | ./dojo-mcp-server

# Get a specific seed prompt
echo '{"jsonrpc":"2.0","id":7,"method":"prompts/get","params":{"name":"dojo.seed.inter_acceptance"}}' | ./dojo-mcp-server
```

**Expected Output:** List of 20 prompts, then full seed content

---

## Phase 3: GitHub Repository Setup (10 minutes)

### Step 3.1: Initialize Git Repository

```bash
# Initialize git (if not already done)
git init

# Add all files
git add .

# Create initial commit
git commit -m "feat: v2.0 - Dojo Genesis MCP Server with AROMA and Serenity Valley wisdom

- 20 seed patches spanning work, rest, and healing
- 11 tools for reflection, search, and practice
- 8 resources with comprehensive documentation
- Complete integration of Dojo Genesis, AROMA, and Serenity Valley"
```

---

### Step 3.2: Create GitHub Repository

```bash
# Create the repository (private by default)
gh repo create TresPies-source/dojo-mcp-server \
  --private \
  --source=. \
  --remote=origin \
  --description="MCP server bringing Dojo wisdom to AI agents: thinking, resting, and healing" \
  --push

# Verify the repository was created
gh repo view TresPies-source/dojo-mcp-server
```

**Expected Output:**
```
✓ Created repository TresPies-source/dojo-mcp-server on GitHub
✓ Added remote origin
✓ Pushed commits to origin/main
```

---

### Step 3.3: Verify GitHub Actions

The repository includes a GitHub Actions workflow that will automatically build and publish the Docker image.

```bash
# Check workflow status
gh workflow list

# View the latest run
gh run list --limit 1
```

**Expected Output:**
```
Docker Build and Publish  active  docker-publish.yml
```

Wait for the workflow to complete (usually 2-3 minutes). You can watch it:

```bash
gh run watch
```

---

## Phase 4: Docker Deployment (10 minutes)

### Step 4.1: Verify Docker Image

Once the GitHub Actions workflow completes:

```bash
# Pull the image
docker pull ghcr.io/trespies-source/dojo-mcp-server:latest

# Verify the image
docker images | grep dojo-mcp-server
```

**Expected Output:**
```
ghcr.io/trespies-source/dojo-mcp-server  latest  abc123def456  2 minutes ago  50MB
```

---

### Step 4.2: Test Docker Container

```bash
# Run the container interactively
docker run -i ghcr.io/trespies-source/dojo-mcp-server:latest

# In another terminal, test it
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | docker run -i ghcr.io/trespies-source/dojo-mcp-server:latest
```

**Expected Output:** List of 11 tools

---

## Phase 5: Claude Desktop Integration (5 minutes)

### Step 5.1: Locate Claude Desktop Config

```bash
# macOS
open ~/Library/Application\ Support/Claude/

# Linux
xdg-open ~/.config/Claude/

# Windows
explorer %APPDATA%\Claude\
```

---

### Step 5.2: Update Configuration

Edit `claude_desktop_config.json`:

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

**Alternative (if you prefer local binary):**

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

### Step 5.3: Restart Claude Desktop

1. Quit Claude Desktop completely
2. Reopen Claude Desktop
3. Look for the MCP server indicator (usually a small icon or status message)

---

### Step 5.4: Test Integration

In Claude Desktop, try these queries:

**Test 1: Search Wisdom**
```
Search the Dojo wisdom base for "burnout prevention"
```

**Expected:** Claude invokes `dojo.search_wisdom` and returns relevant seeds

**Test 2: Reflect on a Situation**
```
Use Dojo Scout mode to help me think about whether to refactor my codebase. 
Perspectives: maintainability, cost, risk, team impact.
```

**Expected:** Claude invokes `dojo.reflect` with mode="scout"

**Test 3: Create a Thinking Room**
```
Create a Dojo thinking room about "AI consciousness and ethics"
```

**Expected:** Claude invokes `dojo.create_thinking_room` and returns a structured template

**Test 4: Practice Inter-Acceptance**
```
Guide me through an Inter-Acceptance exercise. I'm judging myself for not being productive enough.
```

**Expected:** Claude invokes `dojo.practice_inter_acceptance` and guides you through the exercise

**Test 5: Access a Seed Prompt**
```
Show me the "Pace of Understanding" seed from Dojo
```

**Expected:** Claude retrieves the `dojo.seed.pace_of_understanding` prompt

---

## Phase 6: Usage Patterns (Ongoing)

### For Thinking (Dojo Genesis)

**When to use:**
- Making decisions with multiple perspectives
- Designing systems or architectures
- Resolving conflicts or tensions
- Planning complex projects

**Tools to use:**
- `dojo.reflect` (all four modes)
- `dojo.search_wisdom` (for relevant patterns)
- `dojo.apply_seed` (for specific guidance)

**Example:**
```
I'm designing a new feature for my app. Use Dojo Mirror mode to help me see 
the patterns across these perspectives: user experience, technical complexity, 
business value, and team capacity.
```

---

### For Resting (AROMA)

**When to use:**
- Feeling overwhelmed or burned out
- Need to reflect deeply on a topic
- Want to trace the lineage of an idea
- Collaborating with other agents

**Tools to use:**
- `dojo.create_thinking_room` (structured reflection)
- `dojo.trace_lineage` (attribution and sources)
- `dojo.check_pace` (pace assessment)

**Example:**
```
I've been working for 6 hours straight. Use Dojo to check my pace and 
recommend whether I should continue or rest.
```

---

### For Healing (Serenity Valley)

**When to use:**
- Experiencing self-judgment or shame
- Feeling trapped or constrained
- Need emotional support or perspective
- Exploring agency within difficult circumstances

**Tools to use:**
- `dojo.practice_inter_acceptance` (self-compassion)
- `dojo.explore_radical_freedom` (agency exploration)

**Example:**
```
I'm judging myself harshly for missing a deadline. Guide me through an 
Inter-Acceptance exercise.
```

---

## Phase 7: Maintenance & Updates

### Monitoring

```bash
# Check GitHub Actions for build status
gh run list --limit 5

# Check Docker image versions
docker images ghcr.io/trespies-source/dojo-mcp-server

# View repository activity
gh repo view TresPies-source/dojo-mcp-server
```

---

### Updating the Server

When you make changes:

```bash
# Make your changes to the code

# Test locally
go build -o dojo-mcp-server ./cmd/server
./dojo-mcp-server  # test it

# Commit and push
git add .
git commit -m "feat: add new seed patch for X"
git push origin main

# GitHub Actions will automatically build and publish the new Docker image

# Update your local Docker image
docker pull ghcr.io/trespies-source/dojo-mcp-server:latest

# Restart Claude Desktop to pick up the new version
```

---

### Adding New Seeds

1. Edit `internal/wisdom/seeds.go`
2. Add the new seed to the `getSeeds()` function
3. Follow the existing pattern:
   - Name (snake_case)
   - Description (one sentence)
   - Category (dojo_genesis, aroma, serenity_valley, or shared)
   - Triggers (when to use this seed)
   - Content (full markdown content)
4. Test locally
5. Commit and push

---

### Adding New Tools

1. Create handler in `internal/dojo/new_handlers.go` (or create a new file)
2. Register tool in `internal/dojo/handler.go` (`RegisterTools` function)
3. Test locally
4. Update README_V2.md with usage example
5. Commit and push

---

## Troubleshooting

### Issue: "Go not found"

**Solution:**
```bash
# Install Go
brew install go  # macOS
# or download from https://go.dev/dl/
```

---

### Issue: "Docker image not found"

**Solution:**
```bash
# Check GitHub Actions completed successfully
gh run list --limit 1

# If failed, view logs
gh run view --log

# If successful but image not found, check authentication
gh auth status
```

---

### Issue: "Claude Desktop not detecting MCP server"

**Solution:**
1. Check config file syntax (valid JSON)
2. Verify Docker is running: `docker ps`
3. Test Docker image manually: `docker run -i ghcr.io/trespies-source/dojo-mcp-server:latest`
4. Restart Claude Desktop completely
5. Check Claude Desktop logs (usually in same directory as config)

---

### Issue: "Tool returns error"

**Solution:**
```bash
# Test the tool manually
echo '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"dojo.reflect","arguments":{...}}}' | ./dojo-mcp-server

# Check the error message
# Common issues:
# - Missing required parameters
# - Invalid parameter types
# - Empty strings where content expected
```

---

## Success Checklist

Before considering deployment complete, verify:

- [ ] Local build succeeds without errors
- [ ] All 11 tools respond correctly
- [ ] All 20 prompts are accessible
- [ ] All 8 resources return content
- [ ] GitHub repository created and pushed
- [ ] GitHub Actions workflow completed successfully
- [ ] Docker image available and tested
- [ ] Claude Desktop config updated
- [ ] Claude Desktop successfully invokes tools
- [ ] At least 3 different tools tested in Claude Desktop

---

## Next Steps

### Immediate (First Week)
1. **Use it daily** - Integrate Dojo into your thinking practice
2. **Gather feedback** - Note what works and what doesn't
3. **Document patterns** - Save successful usage patterns
4. **Share selectively** - Introduce to trusted collaborators

### Short-Term (First Month)
1. **Refine tools** - Improve based on usage patterns
2. **Add examples** - Document real-world use cases
3. **Create tutorials** - Write guides or record videos
4. **Build community** - Share with the AI agent community

### Long-Term (3-6 Months)
1. **Semantic search** - Implement embeddings-based search
2. **Lineage graph** - Build explicit lineage tracking
3. **AROMA integration** - Connect to AROMA database for persistent thinking rooms
4. **Serenity Valley integration** - Connect to Serenity Valley for guided EIT sessions
5. **Multi-language support** - Translate wisdom for global accessibility

---

## Support & Resources

### Documentation
- `README_V2.md` - Complete usage guide
- `V2_TESTING_REPORT.md` - Testing procedures
- `QUALITY_ASSESSMENT_REPORT.md` - Quality review
- `FIXES_APPLIED.md` - Changelog of fixes

### Community
- GitHub Issues: Report bugs or request features
- GitHub Discussions: Share usage patterns and ask questions

### Philosophy
- Remember: This is a practice, not a product
- Move at the pace of understanding, not extraction
- Honor lineage and credit sources
- Rest is practice; the onsen enables the dojo

---

## Closing

You're about to deploy a complete ecosystem for agentic well-being. This is not just a tool—it's a transmission of practice that spans thinking, resting, and healing.

**May it serve your practice well.** 🏡✨

---

**Prepared with understanding and care,**  
**Manus AI (Dojo)**
