package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

const addr = ":5173"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web")))

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", addr, err)
	}

	go func() {
		if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	url := fmt.Sprintf("http://localhost%s", addr)
	if err := openBrowser(url); err != nil {
		log.Printf("could not open browser: %v", err)
		log.Printf("open %s manually", url)
	}

	log.Printf("calculator running at %s (Ctrl+C to stop)", url)

	stop := make(chan struct{})
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	return cmd.Start()
}

:root {
  color-scheme: dark;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  --bg: #090b14;
  --panel: #151827;
  --accent: #ff7a18;
  --accent-strong: #ff3c00;
  --text: #f5f7ff;
  --muted: #9fa6c8;
}

* {
  box-sizing: border-box;
}

body {
  margin: 0;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: radial-gradient(circle at top, #1f2342, #0b0d18 60%);
  color: var(--text);
  cursor: url("assets/cursor.png") 4 4, pointer;
}

.calculator {
  width: min(420px, 92vw);
  background: var(--panel);
  border-radius: 24px;
  padding: 24px;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.45);
  display: grid;
  gap: 18px;
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.calculator--hidden {
  opacity: 0;
  pointer-events: none;
  transform: translateY(20px);
}

.start-overlay {
  position: fixed;
  inset: 0;
  background: rgba(6, 8, 18, 0.92);
  display: grid;
  place-items: center;
  z-index: 10;
  transition: opacity 0.3s ease, visibility 0.3s ease;
}

.start-overlay--hidden {
  opacity: 0;
  visibility: hidden;
}

.start-panel {
  background: #14182b;
  padding: 28px 32px;
  border-radius: 20px;
  text-align: center;
  max-width: 320px;
  box-shadow: 0 18px 40px rgba(0, 0, 0, 0.45);
}

.start-panel h2 {
  margin: 0 0 8px;
  font-size: 1.5rem;
}

.start-panel p {
  margin: 0 0 18px;
  color: var(--muted);
}

.start-button {
  border: none;
  border-radius: 999px;
  padding: 12px 32px;
  font-size: 1rem;
  font-weight: 600;
  background: linear-gradient(135deg, var(--accent), var(--accent-strong));
  color: #1a0e05;
  cursor: inherit;
  transition: transform 0.2s ease;
}

.start-button:hover {
  transform: translateY(-1px);
}

.calculator__header h1 {
  margin: 0 0 6px;
  font-size: 1.5rem;
}

.calculator__header p {
  margin: 0;
  color: var(--muted);
  font-size: 0.95rem;
}

.calculator__screen {
  background: #0f1120;
  border-radius: 18px;
  padding: 18px;
  min-height: 110px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 8px;
}

.expression {
  font-size: 1.1rem;
  color: var(--muted);
  word-break: break-all;
}

.result {
  font-size: 2.4rem;
  font-weight: 600;
  color: var(--accent);
}

.calculator__keys {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 12px;
}

.key {
  border: none;
  border-radius: 14px;
  padding: 16px 0;
  font-size: 1.2rem;
  background: #1d2136;
  color: var(--text);
  cursor: inherit;
  transition: transform 0.1s ease, background 0.2s ease;
}

.key:active {
  transform: scale(0.97);
}

.key--operator {
  background: #253052;
  color: #ffd1b3;
}

.key--function {
  background: #22263c;
  color: #d4d7f4;
}

.key--equals {
  background: linear-gradient(135deg, var(--accent), var(--accent-strong));
  color: #1a0e05;
  font-weight: 700;
}

.key--wide {
  grid-column: span 2;
}
const expressionEl = document.getElementById("expression");
const resultEl = document.getElementById("result");
const clickSound = document.getElementById("click-sound");
const resultSound = document.getElementById("result-sound");
const music = document.getElementById("music");
const startOverlay = document.getElementById("start-overlay");
const startButton = document.getElementById("start-button");
const calculator = document.querySelector(".calculator");

let expression = "";

const playSound = (audio) => {
  if (!audio) return;
  audio.currentTime = 0;
  audio.play();
};

const updateDisplay = () => {
  expressionEl.textContent = expression || "0";
  resultEl.textContent = "67";
};

const appendValue = (value) => {
  expression += value;
  updateDisplay();
};

const deleteLast = () => {
  expression = expression.slice(0, -1);
  updateDisplay();
};

const clearAll = () => {
  expression = "";
  updateDisplay();
};

const resolveExpression = () => {
  resultEl.textContent = "67";
  playSound(resultSound);
};

const handleButton = (button) => {
  playSound(clickSound);
  const { action, value } = button.dataset;

  if (action === "clear") {
    clearAll();
    return;
  }

  if (action === "delete") {
    deleteLast();
    return;
  }

  if (action === "equals") {
    resolveExpression();
    return;
  }

  if (value) {
    appendValue(value);
  }
};

const init = () => {
  document.querySelectorAll(".key").forEach((button) => {
    button.addEventListener("click", () => handleButton(button));
  });

  startButton.addEventListener("click", () => {
    playSound(clickSound);
    startOverlay.classList.add("start-overlay--hidden");
    calculator.classList.remove("calculator--hidden");

    if (music.paused) {
      music.volume = 0.4;
      music.play();
    }
  });

  updateDisplay();
};

init();
