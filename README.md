# 🐍 NetViper

**NetViper** is a high-performance network exploration and connection analysis tool. It combines a fast multi-threaded port scanner with a versatile traffic interceptor, built entirely in Go.

## 🌟 Key Features

### 📡 Port Scanner (Ready)
- **Live Streaming**: Results are instantly pushed to the web interface via WebSockets.
- **Worker Pool Architecture**: Efficient resource management for scanning wide port ranges.
- **Persistent History**: Automated scan result storage using a local SQLite database.

### 🔄 TCP Proxy (Under Development 🛠)
- **Traffic Interception**: Seamlessly redirect data streams between a client and a target server.
- **Real-time Inspection**: View raw byte streams directly in your browser.
- **Bidirectional Logging**: Capture both request and response payloads for deep analysis.

## 🛠 Tech Stack
- **Core:** Go (native net stack, goroutines, channels)
- **Interface:** HTML5, CSS3, JavaScript (Vanilla JS)
- **API:** WebSockets (Gorilla) & REST
- **Database:** SQLite 3 (CGO-free driver)

## 🏗 Project Structure
```text
.
├── cmd/server/       # Application entry point (main.go)
├── internal/
│   ├── scanner/      # Port scanning engine & logic
│   ├── proxy/        # TCP Proxy engine (upcoming)
│   └── repository/   # Data persistence & history management
└── ui/               # Web-interface assets
