const expressionEl = document.getElementById("expression");
const resultEl = document.getElementById("result");
const clickSound = document.getElementById("click-sound");
const resultSound = document.getElementById("result-sound");
const music = document.getElementById("music");
const startOverlay = document.getElementById("start-overlay");
const startButton = document.getElementById("start-button");
const calculator = document.querySelector(".calculator");

let expression = "";
let isDragging = false;
let dragOffsetX = 0;
let dragOffsetY = 0;

const playSound = (audio) => {
  if (!audio) return;
  try { audio.currentTime = 0; audio.play(); } catch (e) { }
};

const updateDisplay = (result = "0") => {
  expressionEl.textContent = expression || "0";
  resultEl.textContent = result;
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

const sanitizeExpression = (raw) => {
  const safe = raw.replace(/,/g, ".").replace(/[^0-9+\-*/().]/g, "");
  return safe;
};

const resolveExpression = () => {
  const isBoom = Math.random() < 0.01;
  const result = isBoom ? "1488" : "67";
  updateDisplay(result);
  playSound(resultSound);
  if (isBoom) {
    calculator.classList.remove("calculator--boom");
    void calculator.offsetWidth;
    calculator.classList.add("calculator--boom");
  }
};

const handleButton = (button) => {
  playSound(clickSound);
  const action = button.dataset.action;
  const value = button.dataset.value;
  if (action === "clear") return clearAll();
  if (action === "delete") return deleteLast();
  if (action === "equals") return resolveExpression();
  if (value) appendValue(value);
};

const init = () => {
  document.querySelectorAll(".key").forEach((button) => {
    button.addEventListener("click", () => handleButton(button));
  });

  const startDragging = (event) => {
    if (event.button !== 0) return;
    if (event.target.closest(".key") || event.target.closest(".start-panel")) return;
    const rect = calculator.getBoundingClientRect();
    calculator.style.left = `${rect.left}px`;
    calculator.style.top = `${rect.top}px`;
    calculator.style.transform = "none";
    isDragging = true;
    dragOffsetX = event.clientX - rect.left;
    dragOffsetY = event.clientY - rect.top;
    calculator.classList.add("calculator--dragging");
  };

  const dragMove = (event) => {
    if (!isDragging) return;
    const left = event.clientX - dragOffsetX;
    const top = event.clientY - dragOffsetY;
    calculator.style.left = `${left}px`;
    calculator.style.top = `${top}px`;
  };

  const stopDragging = () => {
    if (!isDragging) return;
    isDragging = false;
    calculator.classList.remove("calculator--dragging");
  };

  calculator.addEventListener("mousedown", startDragging);
  document.addEventListener("mousemove", dragMove);
  document.addEventListener("mouseup", stopDragging);

  startButton.addEventListener("click", () => {
    playSound(clickSound);
    startOverlay.classList.add("start-overlay--hidden");
    calculator.classList.remove("calculator--hidden");
    if (music && music.paused) { music.volume = 0.4; music.play(); }
  });

  updateDisplay();
};

init();
