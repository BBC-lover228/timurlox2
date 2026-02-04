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
  try { audio.currentTime = 0; audio.play(); } catch (e) { }
};

const updateDisplay = () => {
  expressionEl.textContent = expression || "0";
  resultEl.textContent = "0";
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
  try {
    const res = eval(expression || "0");
    resultEl.textContent = String(res);
    playSound(resultSound);
  } catch (e) {
    resultEl.textContent = "Err";
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

  startButton.addEventListener("click", () => {
    playSound(clickSound);
    startOverlay.classList.add("start-overlay--hidden");
    calculator.classList.remove("calculator--hidden");
    if (music && music.paused) { music.volume = 0.4; music.play(); }
  });

  updateDisplay();
};

init();
