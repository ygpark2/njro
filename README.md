# njro: Modern Go Microservices Backend

`njro`는 **Ent ORM (`entproto`)**과 **표준 Buf + gRPC**를 결합하여 보일러플레이트를 극소화하고 높은 타입 안전성과 확장성을 제공하는 고성능 마이크로서비스 백엔드 프로젝트입니다.

---

## 🏛️ 아키텍처 개요

본 프로젝트는 서비스의 도메인 특성에 맞추어 최적화된 설계를 적용하고 있습니다:

```
                              ┌────────────────────────────────────────┐
                              │        [agent/] AI Agent & MCP         │
                              │   (Tool Orchestrator & MCP Server)     │
                              └───────────────────┬────────────────────┘
                                                  │ Standard gRPC (Tool Calls)
         ┌───────────────────┬────────────────────┼───────────────────┬───────────────────┐
         ▼                   ▼                    ▼                   ▼                   ▼
  [service/board]      [service/post]      [service/comment]   [service/account]   [service/content]
  (Ent + entproto)    (Ent + entproto)    (Ent + entproto)    (Ent + entproto)    (Ent + entproto)
         │                   │                    │                   │                   │
         └───────────────────┴──────────┬─────────┴───────────────────┴───────────────────┘
                                        ▼
                         ┌─────────────────────────────┐
                         │   RDBMS (PostgreSQL / DB)   │
                         └─────────────────────────────┘

         ┌────────────────────────────────────────┬───────────────────────────────────────┐
         ▼                                        ▼
  [service/emailer]                        [service/search]
  (표준 Buf + gRPC)                       (표준 Buf + gRPC)
```

### 1. 컴포넌트 및 마이크로서비스 목록

| 컴포넌트 / 서비스 | 담당 역할 | 아키텍처 스택 | 기본 포트 / 인터페이스 |
| :--- | :--- | :---: | :---: |
| **`agent/`** | AI 에이전트 자율 실행 및 외부 MCP 도구 제공 | **`MCP JSON-RPC + ReAct Loop`** | `stdio / CLI` |
| **`service/board`** | 게시판 메타데이터 및 카테고리 관리 | **`entproto`** (CRUD 자동 생성) | `:8081` |
| **`service/content`** | 게시글 및 댓글의 본문 텍스트 스토리지 | **`entproto`** (CRUD 자동 생성) | `:8082` |
| **`service/post`** | 게시글 작성, 조회, 추천수, 조회수 등 | **`entproto`** (CRUD 자동 생성) | `:8083` |
| **`service/comment`** | 게시글 댓글 및 대댓글 관리 | **`entproto`** (CRUD 자동 생성) | `:8084` |
| **`service/account`** | 회원 정보, 인증, 프로필 | **`entproto + Bcrypt Hook`** | `:8085` |
| **`service/emailer`** | 이메일 발송 연동 (SMTP/SES/SendGrid) | **표준 `Buf + gRPC`** | `:8086` |
| **`service/search`** | 키워드 인덱싱 및 전문 검색 | **표준 `Buf + gRPC`** | `:8087` |

---

## 🚀 빠른 시작 (Getting Started)

### 요구사항
* Go 1.24+ (권장: 최신 Go)
* [Buf CLI](https://buf.build) (`brew install bufbuild/buf/buf` 또는 `mise use -g buf`)
* [protoc](https://github.com/protocolbuffers/protobuf) (`brew install protobuf` 또는 `mise use -g protoc`)

### 1. 의존성 다운로드 및 테스트
```bash
# 전체 마이크로서비스 및 설정 테스트 실행
go test -v ./service/board/...
go test -v ./service/content/...
go test -v ./service/post/...
go test -v ./service/comment/...
go test -v ./service/account/...
go test -v ./service/emailer/...
go test -v ./service/search/...
go test -v ./agent/...
go test -v ./pkg/config/...
```

### 2. 개별 서비스 로컬 실행
```bash
# 기본적으로 SQLite 메모리/로컬 DB 모드로 즉시 실행됩니다.
go run ./service/board
go run ./service/post
go run ./service/account
```

---

## 🛠️ 코드 생성 가이드 (Code Generation)

### 1. Ent 기반 서비스 (`board`, `content`, `post`, `comment`, `account`)
Go 코드로 작성된 스키마(`ent/schema/*.go`)를 수정하면 DB 모델과 Protobuf 및 gRPC CRUD 서버 코드가 자동 생성됩니다.

```bash
# 예: service/board 코드 생성
go generate ./service/board/ent/...
cd service/board/ent/proto && go generate ./...
```

### 2. Buf 기반 서비스 (`emailer`, `search`)
루트의 `buf.yaml`과 `buf.gen.yaml`을 통해 별도의 복잡한 protoc 플러그인 설치 없이 표준 BSR 원격 플러그인으로 코드를 생성합니다.

```bash
# Proto 린트 검사
make buf-lint        # 또는 buf lint

# Go Proto 및 gRPC 코드 일괄 생성
make buf-generate    # 또는 buf generate
```

---

## 🐳 Docker 컨테이너 빌드 및 관리

프로젝트 루트의 단일 멀티스테이지 [Dockerfile](file:///Users/youngpark/pjt/projects/njro/Dockerfile)과 [Makefile](file:///Users/youngpark/pjt/projects/njro/Makefile)을 통해 초경량(약 20~23MB) 프로덕션 컨테이너 이미지를 빌드 및 관리할 수 있습니다.

### 1. Makefile을 이용한 빌드 (권장)

#### ① 개별 서비스/에이전트 단축 빌드
```bash
make docker-agent      # AI 에이전트 & MCP 서버 빌드 (약 20.3MB)
make docker-board      # 게시판 서비스 빌드 (njro/board:latest, njro/board:$(VERSION))
make docker-post       # 게시글 서비스 빌드
make docker-comment    # 댓글 서비스 빌드
make docker-account    # 계정 서비스 빌드
make docker-content    # 콘텐츠 서비스 빌드
make docker-emailer    # 이메일러 서비스 빌드
make docker-search     # 검색 서비스 빌드
```

#### ② SERVICE 인자를 지정한 빌드
```bash
make docker SERVICE=agent
make docker SERVICE=board
```

#### ③ 모든 마이크로서비스 및 에이전트 일괄 빌드
```bash
make docker
```

---

### 2. Docker CLI 직접 빌드

Makefile 없이 직접 `docker build`를 실행할 경우 `--build-arg SERVICE=<서비스명>`을 지정합니다:

```bash
# 특정 서비스/에이전트 도커 이미지 빌드
docker build --build-arg SERVICE=agent -t njro/agent:latest .
docker build --build-arg SERVICE=board -t njro/board:latest .
docker build --build-arg SERVICE=post -t njro/post:latest .
docker build --build-arg SERVICE=account -t njro/account:latest .
docker build --build-arg SERVICE=emailer -t njro/emailer:latest .
docker build --build-arg SERVICE=search -t njro/search:latest .

# 컨테이너 실행 예시
docker run --rm -it njro/agent:latest tools
docker run --rm -d -p 8081:8081 --name njro-board njro/board:latest
```

---

### 3. 유용한 Make 타겟
```bash
make test              # 전체 패키지 단위 테스트 실행
make test TARGET=board # 특정 서비스(board) 테스트 실행
make run TARGET=board  # 로컬에서 특정 서비스 즉시 기동
make docker_clean      # Dangling 이미지 및 불필요한 레이어 정리
```

---

## 🔒 보안 및 비즈니스 훅 (Bcrypt Password Hook)
`service/account`의 경우 엔티티 생성/수정 시 평문 비밀번호가 전달되면, [user.go](file:///Users/youngpark/pjt/projects/njro/service/account/ent/schema/user.go)의 Ent Mutation Hook이 자동으로 이를 가로채어 `bcrypt` 단방향 해시로 암호화한 뒤 DB에 안전하게 보관합니다.

---

## 🤖 AI Agent & MCP 레이어 (`agent/`)

`agent/`는 하위 마이크로서비스(`service/*`)들의 순수 gRPC 엔드포인트를 **LLM 도구(Tool) 및 MCP(Model Context Protocol) 규격**으로 연결하여 자율 실행과 외부 AI 클라이언트 연동을 담당하는 최상위 레이어입니다.

### 1. 주요 제공 도구 (MCP Tools)
| 도구 이름 | 연동 서비스 | 설명 |
| :--- | :--- | :--- |
| `search_community` | `service/search` | 게시글/문서 키워드 전문 검색 |
| `list_boards` | `service/board` | 등록된 모든 게시판 목록 및 설명 조회 |
| `get_board` | `service/board` | 특정 게시판 상세 정보 조회 |
| `create_board` | `service/board` | 신규 게시판 생성 |
| `list_posts` | `service/post` | 최근 작성된 게시글 목록 조회 |
| `get_post` | `service/post` | 게시글 메타데이터(제목, 작성자, 게시판) 조회 |
| `create_post` | `service/post` | 신규 게시글 등록 |
| `get_content` | `service/content` | 게시글 또는 댓글의 전문 텍스트(본문) 조회 |
| `send_email` | `service/emailer` | 회원 또는 관리자에게 이메일 알림/보고서 발송 |
| `get_user` | `service/account` | 회원 프로필 및 계정 정보 조회 |

### 2. 내부 디렉토리 구조
* **`agent/client/`**: 7개 백엔드 마이크로서비스(`board`, `post`, `content`, `comment`, `account`, `emailer`, `search`)로의 고성능 gRPC 클라이언트 커넥션 풀 및 UUID 헬퍼 관리.
* **`agent/mcp/`**: Anthropic의 표준 Model Context Protocol (MCP) 규격(2024-11-05)을 준수하는 JSON-RPC 2.0 서버 및 10개 도구 정의/디스패처.
* **`agent/orchestrator/`**: 다단계 ReAct / Function Calling AI 에이전트 루프와 OpenAI 호환 LLM 클라이언트 및 스마트 오프라인 Mock LLM 내장.
* **`agent/moderator/`**: 유해 게시글/스팸을 자동 탐지하여 블라인드 처리하고 작성자에게 경고 메일을 발송하는 운영 에이전트.
* **`agent/curator/`**: 검색 서비스(Bleve)와 연동하여 방대한 커뮤니티 지식을 요약 및 인용 답변하는 RAG 큐레이터 에이전트.
* **`pkg/event/`**: 실시간 백그라운드 처리를 위한 도메인 이벤트 발행/구독(Pub/Sub) 인메모리 이벤트 버스.

### 3. 실행 모드 및 사용법

```bash
# 1. 지원 도구 및 JSON Schema 명세 확인
go run ./agent tools

# 2. 오프라인/온라인 AI 에이전트 자율 실행 (자연어 태스크)
# API 키가 없어도 내장된 Mock LLM이 의도를 분석하여 도구를 실행합니다.
go run ./agent chat "커뮤니티 공지사항 검색해줘"

# 외부 LLM(OpenAI/Claude/Ollama) 연동 시:
OPENAI_API_KEY="sk-..." LLM_MODEL="gpt-4o" go run ./agent chat "자유게시판에 테스트 공지글 올려줘"

# 3. 커뮤니티 콘텐츠 안전성 검사 (Moderator Agent)
go run ./agent moderate <post_id> "게시글 제목" "게시글 본문" "author@example.com"

# 4. RAG 기반 지식 큐레이션 및 질의 응답 (Curator Agent)
go run ./agent curate "MSA 분산 트랜잭션 구현 방법"

# 5. 실시간 도메인 이벤트 수신 및 비동기 모더레이션 대기 (Event-Driven Watcher)
go run ./agent watch

# 6. Model Context Protocol (MCP) 표준 서버 모드 실행 (stdio 기반)
go run ./agent mcp

# 7. 에이전트 및 MCP 프로토콜 단위 테스트 실행
go test -v ./agent/...
go test -v ./pkg/event/...
```

### 4. Claude Desktop 및 Cursor 연동

Claude Desktop(`claude_desktop_config.json`) 또는 Cursor(`.cursor/mcp.json`)에 아래 설정을 등록하면 에디터나 데스크톱 앱에서 바로 우리 백엔드 서비스를 AI 도구로 제어할 수 있습니다:

```json
{
  "mcpServers": {
    "njro-backend": {
      "command": "go",
      "args": ["run", "./agent", "mcp"],
      "env": {
        "BOARD_SERVICE_URL": "localhost:8081",
        "CONTENT_SERVICE_URL": "localhost:8082",
        "POST_SERVICE_URL": "localhost:8083",
        "COMMENT_SERVICE_URL": "localhost:8084",
        "ACCOUNT_SERVICE_URL": "localhost:8085",
        "EMAILER_SERVICE_URL": "localhost:8086",
        "SEARCH_SERVICE_URL": "localhost:8087"
      }
    }
  }
}
```
