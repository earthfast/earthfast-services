package http

import (
	"armada-node/model/dynamic"
	"fmt"
	"net/http"
)

// LogsViewerHandler provides an HTML interface for viewing logs
func LogsViewerHandler(self *dynamic.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the current project ID
		node := self.Get()
		projectID := "None"
		if node != nil && !node.ProjectID.IsZero() {
			projectID = node.ProjectID.Hex()
		}

		// HTML with tabs for different log types
		html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Content Node Logs</title>
    <style>
        body { font-family: monospace; margin: 20px; }
        #logs { background: #f5f5f5; padding: 10px; border-radius: 5px; white-space: pre-wrap; height: 600px; overflow-y: auto; }
        .controls { margin-bottom: 10px; }
        .timestamp { color: #666; }
        .stdout { color: #006400; }
        .stderr { color: #8B0000; }
        .info { color: #0066cc; }
        .warn { color: #cc8400; }
        .error { color: #cc0000; }
        .debug { color: #808080; }
        h2 { margin-top: 0; }
        .tab-buttons { margin-bottom: 10px; }
        .tab-button { padding: 8px 16px; cursor: pointer; background: #e0e0e0; border: none; border-radius: 4px; margin-right: 5px; }
        .tab-button.active { background: #007bff; color: white; }
    </style>
</head>
<body>
    <h1>Content Node Logs</h1>
    
    <div class="tab-buttons">
        <button id="nextjs-tab" class="tab-button active" onclick="switchTab('nextjs')">Next.js Logs</button>
        <button id="server-tab" class="tab-button" onclick="switchTab('server')">Server Logs</button>
    </div>
    
    <div id="nextjs-panel">
        <h2>Project: %s</h2>
        <div class="controls">
            <button onclick="fetchLogs('nextjs')">Refresh Logs</button>
            <button onclick="startAutoRefresh('nextjs')">Auto-Refresh</button>
            <button onclick="stopAutoRefresh()">Stop Refresh</button>
            Lines: <input type="number" id="nextjs-limit" value="100" min="1" max="1000">
        </div>
    </div>
    
    <div id="server-panel" style="display:none">
        <h2>Server Logs</h2>
        <div class="controls">
            <button onclick="fetchLogs('server')">Refresh Logs</button>
            <button onclick="startAutoRefresh('server')">Auto-Refresh</button>
            <button onclick="stopAutoRefresh()">Stop Refresh</button>
            Lines: <input type="number" id="server-limit" value="100" min="1" max="1000">
        </div>
    </div>
    
    <div id="logs">Select a tab and click "Refresh Logs" to view logs.</div>

    <script>
        let refreshInterval;
        let currentTab = 'nextjs';
        
        function switchTab(tab) {
            // Update tab buttons
            document.getElementById('nextjs-tab').classList.toggle('active', tab === 'nextjs');
            document.getElementById('server-tab').classList.toggle('active', tab === 'server');
            
            // Show/hide panels
            document.getElementById('nextjs-panel').style.display = tab === 'nextjs' ? 'block' : 'none';
            document.getElementById('server-panel').style.display = tab === 'server' ? 'block' : 'none';
            
            // Update current tab
            currentTab = tab;
            
            // Stop any existing refresh
            stopAutoRefresh();
            
            // Fetch logs for the selected tab
            fetchLogs(tab);
        }
        
        function fetchLogs(type) {
            const limitId = type + '-limit';
            const limit = document.getElementById(limitId).value;
            
            document.getElementById('logs').innerHTML = 'Loading...';
            
            fetch('/api/logs?type=' + type + '&limit=' + limit)
                .then(response => response.json())
                .then(data => {
                    const logsElement = document.getElementById('logs');
                    logsElement.innerHTML = '';
                    
                    if (data.logs.length === 0) {
                        logsElement.textContent = 'No logs found.';
                        return;
                    }
                    
                    data.logs.forEach(log => {
                        const logLine = document.createElement('div');
                        const timestamp = document.createElement('span');
                        timestamp.className = 'timestamp';
                        timestamp.textContent = '[' + new Date(log.timestamp).toISOString() + '] ';
                        
                        const source = document.createElement('span');
                        source.className = log.source;
                        source.textContent = '[' + log.source + '] ';
                        
                        const message = document.createElement('span');
                        message.textContent = log.message;
                        
                        logLine.appendChild(timestamp);
                        logLine.appendChild(source);
                        logLine.appendChild(message);
                        logsElement.appendChild(logLine);
                    });
                })
                .catch(error => {
                    console.error('Error fetching logs:', error);
                    document.getElementById('logs').textContent = 'Error fetching logs: ' + error.message;
                });
        }
        
        function startAutoRefresh(type) {
            stopAutoRefresh();
            fetchLogs(type);
            refreshInterval = setInterval(() => fetchLogs(type), 5000);
        }
        
        function stopAutoRefresh() {
            if (refreshInterval) {
                clearInterval(refreshInterval);
                refreshInterval = null;
            }
        }
        
        // Initial fetch
        document.addEventListener('DOMContentLoaded', () => {
            fetchLogs('nextjs');
        });
    </script>
</body>
</html>`, projectID)

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	}
}
