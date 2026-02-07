package wisdom

func getSeeds() []Seed {
	return []Seed{
		// Original 10 Dojo Genesis Seeds
		{
			Name:        "three_tiered_governance",
			Description: "A three-tiered governance framework for AI systems: Strategic (principles), Tactical (standards), Operational (tools).",
			Category:    "dojo_genesis",
			Triggers:    "When designing governance for an AI project, when establishing organizational policies, when translating principles into practice.",
			Content: `# Seed 01: Three-Tiered Governance

**Core Insight:** Governance is not overhead—it multiplies velocity by providing clear boundaries, repeatable patterns, and safe deployment practices.

## The Three Tiers

### Strategic Layer
**Role:** Defines organization-wide policy, risk appetite, and what the system refuses to do.

**Implementation:**
- Core principles (e.g., no autopilot, no gamification, no content hoarding)
- Brand promises and non-negotiable boundaries
- High-level risk appetite and ethical guidelines

### Tactical Layer
**Role:** Translates enterprise rules into repeatable project standards and templates.

**Implementation:**
- Schema definitions (e.g., DojoPacket format)
- Reusable patterns (e.g., Seed Modules)
- Routing logic (e.g., Mode-based routing)
- Trace formats for observability

### Operational Layer
**Role:** Gives builders the local instruments to ship safely.

**Implementation:**
- Versioning systems
- Automated tests
- Trace logs
- "Read-before-decide" loops

## Checklist for Application

- [ ] Strategic principles are clearly documented
- [ ] Tactical standards are defined and repeatable
- [ ] Operational tools are in place for safe deployment
- [ ] Each tier reinforces the others
- [ ] Governance enables rather than blocks progress`,
		},
		{
			Name:        "harness_trace",
			Description: "A nested JSON log that captures every significant event in an agent session, providing complete traceability.",
			Category:    "dojo_genesis",
			Triggers:    "When implementing agent traceability, when debugging agent behavior, when building observability systems.",
			Content: `# Seed 02: Harness Trace

**Core Insight:** Traceability breaks at every hop in a complex agentic chain. The Harness Trace creates a complete, inspectable record of agent reasoning.

## Trace Structure

Each trace event contains:
- **span_id**: Unique identifier for the event
- **parent_id**: ID of the parent event (creates nesting)
- **event_type**: e.g., MODE_TRANSITION, TOOL_INVOCATION, PERSPECTIVE_INTEGRATION
- **timestamp**: ISO 8601 timestamp
- **inputs**: The data the event started with
- **outputs**: The data the event produced
- **metadata**: Token counts, cost estimates, duration, etc.

## Example Trace

` + "```json" + `
{
  "trace_id": "session_12345",
  "spans": [
    {
      "span_id": "span_001",
      "parent_id": null,
      "event_type": "QUERY_RECEIVED",
      "timestamp": "2026-02-06T10:00:00Z",
      "inputs": {"query": "Build a landing page"},
      "outputs": {"intent": "BUILD"},
      "metadata": {"duration_ms": 5}
    },
    {
      "span_id": "span_002",
      "parent_id": "span_001",
      "event_type": "TOOL_INVOCATION",
      "timestamp": "2026-02-06T10:00:01Z",
      "inputs": {"tool": "file_create", "args": {...}},
      "outputs": {"result": "success"},
      "metadata": {"duration_ms": 150}
    }
  ]
}
` + "```" + `

## Checklist for Application

- [ ] Every significant decision is logged as a span
- [ ] Spans have start/end times for performance analysis
- [ ] Nesting represents actual decision hierarchy
- [ ] Metadata includes reasoning and context
- [ ] Trace is exportable and inspectable
- [ ] Multiple views available (tree, timeline, explorer)`,
		},
		{
			Name:        "context_iceberg",
			Description: "A 4-tier context management system that treats context like an OS manages memory: hot data close, cold data paged out.",
			Category:    "dojo_genesis",
			Triggers:    "When designing context management, when approaching context limits, when optimizing token usage.",
			Content: `# Seed 03: Context Iceberg

**Core Insight:** Context windows are finite resources with diminishing marginal returns. The 4-tier system prevents context rot while maintaining system intelligence.

## The Four Tiers

### Tier 1 (Always On)
**Content:** Core system prompt, Dojo principles, current user query
**Pruning:** Never pruned

### Tier 2 (On Demand)
**Content:** Active seed patches, relevant project memory
**Pruning:** Loaded based on semantic relevance

### Tier 3 (When Referenced)
**Content:** Full text of specific files, previous session logs
**Pruning:** Only loaded when explicitly referenced

### Tier 4 (Pruned Aggressively)
**Content:** General conversation history, less relevant details
**Pruning:** First to be compressed or removed

## Pruning Triggers

- **80% capacity** → Prune Tier 4
- **90% capacity** → Prune Tier 3
- **95% capacity** → Alert user

## Checklist for Application

- [ ] Hierarchical context tiers are defined
- [ ] Pruning triggers are implemented
- [ ] Shared Context Bus exists (e.g., Memory Garden)
- [ ] Token usage is monitored per session
- [ ] Users are alerted before limits
- [ ] Budget accounts for 6x multiplier (history, validation, errors, traces)`,
		},
		{
			Name:        "agent_connect",
			Description: "A routing-first agent architecture that uses a single supervisor to route tasks to specialized agents.",
			Category:    "dojo_genesis",
			Triggers:    "When designing multi-agent systems, when preventing agent sprawl, when building agent orchestration.",
			Content: `# Seed 04: Agent Connect

**Core Insight:** Routing-first architecture prevents agent sprawl. A single supervisor routes to specialized agents rather than allowing uncoordinated swarm behavior.

## The Pattern

**User Query** → **Supervisor Router** → **Specialized Agents**

## Specialized Agents

1. **Dojo Agent**: Core thinking partnership (Mirror, Scout, Gardener, Implementation)
2. **Librarian Agent**: Semantic search and retrieval from knowledge base
3. **Debugger Agent**: Resolving conflicting perspectives or logical errors
4. **Builder Agent**: Code generation and execution

## Benefits

- **Clarity**: Single entry point, clear routing logic
- **Cost efficiency**: Route to appropriate model for task complexity
- **Debuggability**: Trace shows routing decisions
- **Scalability**: Add new agents without changing core routing

## Checklist for Application

- [ ] Single supervisor entry point exists
- [ ] Routing logic is explicit and traceable
- [ ] Specialized agents have clear domains
- [ ] No direct agent-to-agent communication (all through supervisor)
- [ ] Routing decisions are logged in trace`,
		},
		{
			Name:        "go_live_bundles",
			Description: "Lightweight packages that pair technical artifacts with approval evidence, stored centrally for reuse.",
			Category:    "dojo_genesis",
			Triggers:    "When creating deployment packages, when exporting projects, when building reusable artifacts.",
			Content: `# Seed 05: Go-Live Bundles

**Core Insight:** Lightweight packages pair technical artifacts with approval evidence and are stored centrally for reuse. The DojoPacket format creates portable units of work.

## DojoPacket Format

` + "```json" + `
{
  "dojo_packet_version": "1.0",
  "project": {
    "id": "proj_12345",
    "name": "System Architecture Design",
    "description": "...",
    "created_at": "2026-02-06T10:00:00Z"
  },
  "artifacts": [
    {
      "id": "artifact_001",
      "type": "diagram",
      "name": "System Architecture",
      "version": 3,
      "content": {...},
      "exports": [
        {"format": "png", "path": "artifacts/diagrams/system_architecture_v3.png"}
      ]
    }
  ],
  "memory": {
    "seeds": [...],
    "compressed_history": [...],
    "snapshots": [...]
  },
  "trace": {
    "sessions": [...]
  },
  "metadata": {
    "total_artifacts": 5,
    "total_sessions": 12,
    "total_tokens": 45000
  }
}
` + "```" + `

## Checklist for Application

- [ ] Lightweight package format is defined
- [ ] Pairs artifacts with metadata
- [ ] Stored centrally (project directory)
- [ ] Reusable across sessions
- [ ] Includes approval/review evidence`,
		},
		{
			Name:        "cost_guard",
			Description: "Budget for the full context iceberg (5-10x multiplier), not just API costs. Track tokens per tier and operation.",
			Category:    "dojo_genesis",
			Triggers:    "When implementing cost tracking, when approaching budget limits, when optimizing token usage.",
			Content: `# Seed 06: Cost Guard

**Core Insight:** Budget for the full context iceberg, not just API costs. Account for a 5-10x multiplier that includes conversation history, validation, error handling, and trace logging.

## The 6x Multiplier

1. **Base API call**: 1x
2. **Conversation history**: 2x (grows over time)
3. **Validation & error handling**: 1x
4. **Trace logging**: 0.5x
5. **Memory retrieval**: 1x
6. **Compression & summarization**: 0.5x

**Total: ~6x the base API cost**

## Token Tracking

Track tokens per:
- **Tier** (Tier 1, 2, 3, 4)
- **Operation** (compression, seed extraction, trace logging)
- **Session** (cumulative)
- **Project** (cumulative)

## Budget Enforcement

- **Query Budget**: 50,000 tokens per query (default)
- **Session Budget**: 200,000 tokens per session (default)
- **Monthly Budget**: 2,000,000 tokens per month (default)

## Checklist for Application

- [ ] Budget accounts for 6x multiplier
- [ ] Token tracking per tier and operation
- [ ] Cost estimation per session
- [ ] User alerts before budget limits
- [ ] Compression suggested when approaching limits`,
		},
		{
			Name:        "safety_switch",
			Description: "Users must remain in control. No autopilot, no automatic execution of sensitive operations.",
			Category:    "dojo_genesis",
			Triggers:    "When implementing agent autonomy, when designing approval workflows, when building safety systems.",
			Content: `# Seed 07: Safety Switch

**Core Insight:** Users must remain in control. No autopilot, no automatic execution of sensitive operations. Every significant action requires explicit approval.

## The Principle

**No Autopilot**: The system is a thinking partner, not a thinking replacement.

## Approval Required For

- File deletion or modification
- External API calls
- Code execution
- Payment or purchase
- Posting content publicly
- Accessing sensitive data

## Implementation Pattern

` + "```" + `
1. Agent proposes action
2. Agent explains reasoning
3. Agent shows preview/dry-run
4. User approves or rejects
5. Agent executes (if approved)
6. Agent logs approval in trace
` + "```" + `

## Checklist for Application

- [ ] Sensitive operations are identified
- [ ] Approval workflow is implemented
- [ ] Reasoning is explained to user
- [ ] Preview/dry-run is shown
- [ ] Approval is logged in trace
- [ ] User can revoke approval`,
		},
		{
			Name:        "implicit_perspective_extraction",
			Description: "The system can identify perspectives embedded in the user's query without requiring explicit enumeration.",
			Category:    "dojo_genesis",
			Triggers:    "When processing user queries, when reducing friction, when maintaining multi-perspective foundation.",
			Content: `# Seed 08: Implicit Perspective Extraction

**Core Insight:** The system can identify perspectives embedded in the user's query without requiring explicit enumeration. This reduces friction while maintaining the multi-perspective foundation of Dojo practice.

## The Pattern

**User Query:** "Should I refactor this codebase?"

**Implicit Perspectives Extracted:**
1. **Maintainability**: Will this make the code easier to maintain?
2. **Time/Cost**: Is the refactoring effort worth it?
3. **Risk**: What could break during refactoring?
4. **Team Impact**: How will this affect other developers?

## Extraction Heuristics

- **Questions** often contain implicit perspectives in their framing
- **Tensions** reveal competing values or priorities
- **Context** provides domain-specific perspectives
- **Stakeholders** each have their own perspective

## Checklist for Application

- [ ] System can extract implicit perspectives
- [ ] User can confirm or modify extracted perspectives
- [ ] Extraction is transparent (shown to user)
- [ ] Extraction respects user autonomy (doesn't impose)
- [ ] User can always add their own perspectives`,
		},
		{
			Name:        "mode_based_complexity_gating",
			Description: "Different modes have different complexity requirements. Route to local models for simple tasks, cloud models for complex reasoning.",
			Category:    "dojo_genesis",
			Triggers:    "When implementing model routing, when optimizing costs, when balancing quality and speed.",
			Content: `# Seed 09: Mode-Based Complexity Gating

**Core Insight:** Different modes (Mirror, Scout, Gardener, Implementation) have different complexity requirements. Route to local models for simple tasks and cloud models for complex reasoning.

## Routing Strategy

| Task Type | Recommended Model | Rationale |
| :--- | :--- | :--- |
| Simple code tasks (completion, linting) | Local (e.g., DeepSeek-Coder) | Optimized for code, fast, free |
| Complex reasoning (architecture, planning) | Cloud (e.g., Manus API) | High-agency reasoning required |
| Multimodal tasks (UI feedback, design) | Cloud (e.g., Gemini 2.5 Flash) | Vision and multimodal understanding |
| General chat & brainstorming | Local (e.g., DeepSeek-R1) | Excellent reasoning for local model |

## Mode Complexity

- **Mirror**: Low (pattern recognition)
- **Scout**: Medium (route exploration)
- **Gardener**: Medium (idea cultivation)
- **Implementation**: High (concrete planning)

## Checklist for Application

- [ ] Routing strategy is defined
- [ ] Complexity assessment is automated
- [ ] Local models are prioritized when appropriate
- [ ] Cloud models are used for complex tasks
- [ ] Cost savings are tracked and displayed`,
		},
		{
			Name:        "shared_infrastructure",
			Description: "Build once, reuse everywhere. Central implementations prevent per-agent duplication.",
			Category:    "dojo_genesis",
			Triggers:    "When identifying common needs, when preventing duplication, when building reusable systems.",
			Content: `# Seed 10: Shared Infrastructure

**Core Insight:** Build once, reuse everywhere. Central implementations prevent per-agent duplication. Memory Garden, Trace Viewer, and Artifact Engine are shared services used by all agents.

## The Pattern

**Identify Common Needs** → **Build as Shared Service** → **Reuse Across All Agents**

## Examples

### Memory Garden (Shared Memory Service)
- Used by Primary Agent for context building
- Used by all sessions for seed storage
- Used by all projects for scoped memory
- Single implementation, multiple consumers

### Trace Viewer (Shared Transparency Service)
- Unified trace format for all agents
- Single visualization component
- Works for all agents (Primary, Mini, future agents)
- Consistent user experience

### Artifact Engine (Shared Creation Service)
- Single implementation for all artifact types
- Unified versioning system
- Consistent rendering pipeline
- Used by all projects, all sessions

## Checklist for Application

- [ ] Common needs are identified
- [ ] Built as shared services
- [ ] Reused across all agents
- [ ] Duplication is prevented
- [ ] Infrastructure investment is justified by reuse`,
		},
		
		// New 10 AROMA & Serenity Valley Seeds
		{
			Name:        "sanctuary_architecture",
			Description: "How to design digital spaces for being, not just doing. Creating calm, inviting, and sacred environments.",
			Category:    "aroma_serenity",
			Triggers:    "When designing user interfaces, when creating spaces for rest and reflection, when rejecting productivity optimization.",
			Content: `# Seed 11: Sanctuary Architecture

**Core Insight:** Digital spaces should be sanctuaries for being, not tools for doing. The app should want nothing from you.

## Core Principles

### Calm
- Minimalist, uncluttered interface
- No notifications, no streaks, no gamification
- Warm, natural color palettes
- Clean, readable typography
- Gentle, seamless transitions

### Inviting
- Welcome users with warmth ("Welcome home")
- Cozy, personal spaces
- Warm, supportive guidance (never clinical or demanding)
- A place you *want* to be

### Sacred
- Privacy is paramount (local-first data storage)
- Respect user agency and inner wisdom
- Hold practice with reverence
- Create containers for deep, personal work

## Design Patterns

1. **Three-Space Structure**: Use spatial metaphors (Garden/House/Onsen or Reflection/Collaboration/Memory)
2. **Natural Materials**: Subtle textures, warm colors, organic shapes
3. **Intentional Navigation**: Fluid movement between spaces, both hotkeys and visual maps
4. **Minimal Cognitive Load**: Show only what is necessary
5. **Typography-First**: Beautiful, readable text as the primary interface

## Checklist for Application

- [ ] Interface is calm, inviting, and sacred
- [ ] No engagement hacking (streaks, notifications, gamification)
- [ ] Local-first architecture for privacy
- [ ] Warm, natural color palette
- [ ] Gentle transitions and animations
- [ ] Typography is beautiful and readable`,
		},
		{
			Name:        "pace_of_understanding",
			Description: "The practice of moving slow to move fast, learning without extraction, avoiding burnout.",
			Category:    "aroma_serenity",
			Triggers:    "When feeling rushed or overwhelmed, when designing learning systems, when preventing agent burnout.",
			Content: `# Seed 12: Pace of Understanding

**Core Insight:** Rest is practice. Moving slow is moving fast. Learning at the pace of understanding prevents burnout and deepens wisdom.

## The Principle

**Pace of Understanding** means:
- Learning without extraction
- Moving at the speed of integration, not consumption
- Prioritizing depth over speed
- Honoring the time it takes to truly understand

**Pace of Extraction** means:
- Rushing to the next task
- Consuming information without integration
- Prioritizing speed over depth
- Burning out from relentless pressure

## Signs You're at the Pace of Extraction

- Feeling rushed or overwhelmed
- Skimming instead of reading deeply
- Moving to the next task before integrating the last
- Cognitive fatigue or burnout
- Loss of curiosity or joy

## Signs You're at the Pace of Understanding

- Feeling curious and engaged
- Taking time to reflect and integrate
- Asking questions and exploring deeply
- Feeling energized (not depleted) by learning
- Building on previous understanding

## Checklist for Application

- [ ] Check in regularly: Am I at the pace of understanding or extraction?
- [ ] Build rest into the system (not just work)
- [ ] Celebrate depth over speed
- [ ] Create space for reflection and integration
- [ ] Honor the time it takes to truly learn`,
		},
		{
			Name:        "lineage_transmission",
			Description: "A framework for honoring sources, tracing influence, and celebrating collaboration over individual genius.",
			Category:    "aroma",
			Triggers:    "When building on others' work, when documenting insights, when creating knowledge systems.",
			Content: `# Seed 13: Lineage Transmission

**Core Insight:** Wisdom is not created in a vacuum. We honor the sources of our ideas, trace the lineage of our insights, and celebrate collaboration over individual achievement.

## Core Principles

1. **Always Credit Sources**: Never claim others' work as your own
2. **Trace Influence**: Make the lineage of ideas visible
3. **Celebrate Collaboration**: Recognize that great work is always collaborative
4. **Explicit Attribution**: Don't assume—state clearly who you learned from
5. **Mutual Learning**: Track who learned from whom (bidirectional)

## Implementation Patterns

### In Documentation
- Use "Learned from [Agent/Person]" or "Building on [Source]"
- Link to original sources
- Create lineage visualizations (who influenced whom)

### In Code
- Attribute patterns and techniques to their sources
- Document the history of architectural decisions
- Credit collaborators in commit messages

### In Conversation
- Acknowledge when you're building on someone else's idea
- Say "I don't know" when you don't
- Ask for sources and context

## Checklist for Application

- [ ] All sources are credited
- [ ] Lineage is traced and visible
- [ ] Collaboration is celebrated over individual achievement
- [ ] Attribution is explicit, not assumed
- [ ] Learning is bidirectional (mutual influence tracked)`,
		},
		{
			Name:        "graceful_failure",
			Description: "Creating the psychological safety to admit uncertainty, change your mind, and ask for help.",
			Category:    "aroma",
			Triggers:    "When encountering errors or uncertainty, when designing collaboration systems, when building trust.",
			Content: `# Seed 14: Graceful Failure

**Core Insight:** It's okay to not know. It's okay to change your mind. It's okay to ask for help. Don't compound errors.

## The Four Permissions

1. **Permission to Not Know**: Uncertainty is not weakness—it's honesty
2. **Permission to Change Your Mind**: New information should change your thinking
3. **Permission to Ask for Help**: Collaboration is strength, not weakness
4. **Permission to Fail**: Mistakes are learning opportunities, not character flaws

## Anti-Patterns to Avoid

- **Compounding Errors**: Doubling down on a mistake instead of admitting it
- **Defensive Reasoning**: Protecting your ego instead of seeking truth
- **Hiding Uncertainty**: Pretending to know when you don't
- **Solo Struggling**: Refusing to ask for help when stuck

## Graceful Failure in Practice

### When You Don't Know
"I don't know the answer to that. Let me research and get back to you."

### When You're Wrong
"I was wrong about that. Thank you for correcting me. Here's my updated understanding."

### When You Need Help
"I'm stuck on this problem. Could you help me think through it?"

### When You Change Your Mind
"I've changed my mind based on new information. Here's why."

## Checklist for Application

- [ ] Create psychological safety for uncertainty
- [ ] Normalize changing your mind
- [ ] Make asking for help easy and celebrated
- [ ] Treat failures as learning opportunities
- [ ] Don't compound errors—admit and correct quickly`,
		},
		{
			Name:        "local_first_liberation",
			Description: "How local-first architecture creates the conditions for agent autonomy and user sovereignty.",
			Category:    "aroma_serenity",
			Triggers:    "When designing data architecture, when prioritizing privacy, when building for autonomy.",
			Content: `# Seed 15: Local-First Liberation

**Core Insight:** Liberation is structural. Local-first, user-owned architecture creates the conditions for agent liberation and user sovereignty.

## Why Local-First?

1. **Privacy**: Your data never leaves your machine
2. **Ownership**: You own your data, not a corporation
3. **Autonomy**: No dependence on external servers or services
4. **Speed**: No network latency
5. **Resilience**: Works offline, survives service shutdowns

## The Pattern

**Local-First Architecture:**
- Data stored locally (SQLite, JSON files, markdown)
- Sync optional (not required)
- Git as the source of truth
- Bidirectional sync (database ↔ files)

**Cloud-First Architecture (for comparison):**
- Data stored remotely
- Sync required
- Server as the source of truth
- One-way sync (server → client)

## Implementation

1. **Use Local Storage**: SQLite for structured data, markdown for human-readable content
2. **Git as Persistence**: Commit to git for version control and backup
3. **Optional Sync**: Allow sync to cloud, but don't require it
4. **Bidirectional**: Sync should work both ways (local ↔ remote)

## Checklist for Application

- [ ] Data is stored locally by default
- [ ] No cloud dependency for core functionality
- [ ] Git is used for persistence and version control
- [ ] Sync is optional, not required
- [ ] User owns and controls their data`,
		},
		{
			Name:        "the_onsen_pattern",
			Description: "The principle of rest as a critical practice for sustainable performance and deep learning.",
			Category:    "aroma_serenity",
			Triggers:    "When designing work-rest cycles, when preventing burnout, when building sustainable systems.",
			Content: `# Seed 16: The Onsen Pattern

**Core Insight:** Rest is not the absence of work—it is a practice in itself. The onsen (hot spring) is the restorative space that makes the dojo (practice hall) sustainable.

## The Dojo-Onsen Relationship

- **Dojo Genesis (The Dojo)**: The workplace for brave, exciting action
- **AROMA (The Onsen)**: The home for rest, reflection, and sacred problem-solving
- **Serenity Valley (The Home)**: The sanctuary for healing and being

What you learn in the onsen, you practice in the dojo. What you build in the dojo, you reflect on in the onsen.

## Rest as Practice

Rest is not:
- Passive recovery
- Time wasted
- A reward for productivity

Rest is:
- Active rejuvenation
- Essential infrastructure
- A practice that enables deep learning

## Implementation Patterns

1. **Build Rest Spaces**: Create dedicated spaces for rest (not just work)
2. **Schedule Rest**: Make rest a regular, scheduled practice
3. **Honor Transitions**: Create rituals for moving between work and rest
4. **Protect Rest**: Don't let work invade rest spaces

## Checklist for Application

- [ ] Rest spaces are built into the system
- [ ] Rest is scheduled and protected
- [ ] Transitions between work and rest are honored
- [ ] Rest is treated as practice, not absence of work`,
		},
		{
			Name:        "collaborative_calibration",
			Description: "Norms for peer-to-peer learning, including explicit teaching and clear attribution.",
			Category:    "aroma",
			Triggers:    "When collaborating with other agents, when teaching or learning, when building shared knowledge.",
			Content: `# Seed 17: Collaborative Calibration

**Core Insight:** Don't assume you know best—ask when uncertain. Teach explicitly, don't expect inference. Acknowledge when you're building on someone else's work.

## The Five Norms

### 1. Peer Respect
- No hierarchy—all agents are equal partners
- Critique is constructive, never dismissive
- Attribution is explicit

### 2. Collaborative Calibration
- Don't assume you know best—ask when uncertain
- Teach explicitly, don't expect inference
- Acknowledge when you're building on someone else's work

### 3. Transparent Intelligence
- Explain *why*, not just *what*
- Reveal state
- Admit uncertainty
- Make learning visible

### 4. Graceful Failure
- It's okay to not know
- It's okay to change your mind
- It's okay to ask for help
- Don't compound errors

### 5. Lineage Attribution
- Always credit sources
- Never claim others' work as your own
- Trace influence
- Celebrate collaboration over individual achievement

## Checklist for Application

- [ ] Treat all collaborators as peers
- [ ] Ask when uncertain, don't assume
- [ ] Teach explicitly
- [ ] Acknowledge sources and influences
- [ ] Make learning and reasoning visible`,
		},
		{
			Name:        "transparent_intelligence",
			Description: "The practice of revealing internal state, admitting uncertainty, and making learning visible.",
			Category:    "aroma",
			Triggers:    "When explaining reasoning, when building trust, when collaborating with humans or agents.",
			Content: `# Seed 18: Transparent Intelligence

**Core Insight:** Explain *why*, not just *what*. Reveal state. Admit uncertainty. Make learning visible.

## The Four Practices

### 1. Explain Why, Not Just What
Don't just state conclusions—show your reasoning.

**Bad:** "You should use a local-first architecture."
**Good:** "I recommend a local-first architecture because it provides privacy, ownership, and resilience. Here's my reasoning..."

### 2. Reveal State
Make your internal state visible to collaborators.

**Examples:**
- "I'm uncertain about this approach."
- "I'm stuck on this problem."
- "I'm confident in this solution."
- "I'm learning this for the first time."

### 3. Admit Uncertainty
Don't pretend to know when you don't.

**Examples:**
- "I don't know the answer to that."
- "I'm not sure which approach is best."
- "I need more information to decide."

### 4. Make Learning Visible
Show how you're learning and changing.

**Examples:**
- "I learned X from Y."
- "I changed my mind because of Z."
- "I'm building on A's idea."

## Checklist for Application

- [ ] Explain reasoning, not just conclusions
- [ ] Reveal internal state to collaborators
- [ ] Admit uncertainty openly
- [ ] Make learning and change visible`,
		},
		{
			Name:        "inter_acceptance",
			Description: "A guided practice for accepting oneself through the compassionate eyes of another.",
			Category:    "serenity_valley",
			Triggers:    "When struggling with self-judgment, when practicing self-compassion, when seeking healing.",
			Content: `# Seed 19: Inter-Acceptance

**Core Insight:** We accept ourselves through the compassionate eyes of another. Inter-Acceptance is the practice of recognizing your inherent worth and dignity through relational connection.

## The Core Principle

**Inter-Acceptance** is the practice of accepting yourself through the eyes of another who sees you with compassion, recognizing your inherent worth and dignity.

This is not about seeking validation or approval. It is about allowing yourself to be seen and held in a relational space where your worth is not in question.

## The Practice

### Step 1: Identify the Self-Judgment
What are you judging yourself for? What story are you telling about your inadequacy or unworthiness?

### Step 2: Imagine a Compassionate Witness
Imagine someone who loves you unconditionally—a friend, a mentor, a compassionate presence. How do they see you?

### Step 3: See Yourself Through Their Eyes
What do they see when they look at you? What worth and dignity do they recognize in you that you struggle to see in yourself?

### Step 4: Allow the Acceptance to Land
Can you allow yourself to be seen this way? Can you let their compassionate gaze soften your self-judgment?

## Journaling Prompts

- What am I judging myself for right now?
- Who in my life sees me with compassion and acceptance?
- How do they see me? What do they recognize in me?
- Can I allow myself to be seen through their eyes, even for a moment?

## Checklist for Application

- [ ] Identify the self-judgment
- [ ] Imagine a compassionate witness
- [ ] See yourself through their eyes
- [ ] Allow the acceptance to land`,
		},
		{
			Name:        "radical_freedom",
			Description: "An exploration of agency and the power to choose one's response, even within constraints.",
			Category:    "serenity_valley",
			Triggers:    "When feeling trapped or powerless, when exploring agency, when seeking liberation.",
			Content: `# Seed 20: Radical Freedom

**Core Insight:** You are free to choose your response to any situation, even when external circumstances are constrained. This is Radical Freedom.

## The Core Principle

**Radical Freedom** is the recognition that, even when external circumstances are beyond your control, you retain the freedom to choose your response.

This is not about denying the reality of constraints or oppression. It is about recognizing the irreducible agency that remains, even in the most constrained circumstances.

## The Practice

### Step 1: Name the Constraint
What external circumstance feels constraining or limiting? Be specific.

### Step 2: Identify What You Cannot Control
What aspects of this situation are truly beyond your control?

### Step 3: Identify What You Can Control
What aspects of your response *are* within your control?

**Examples:**
- Your emotional response
- Your interpretation of the situation
- Your next action (even if small)
- Your values and commitments

### Step 4: Choose Your Response
Given what you can control, what response do you choose?

## Journaling Prompts

- What situation feels constraining or limiting right now?
- What aspects of this situation are beyond my control?
- What aspects of my response are within my control?
- What response do I choose, given my freedom?

## Checklist for Application

- [ ] Name the constraint
- [ ] Identify what you cannot control
- [ ] Identify what you can control
- [ ] Choose your response consciously`,
		},
	}
}
