// 🎉


function percentage(partialValue, totalValue) {
  return (100 * partialValue) / totalValue;
}

/**
 * @typedef Pomodo
 * @property {string} title
 * @property {boolean} isDone
 * @property {string} type
 * @property {number} minutes
 * @property {number} seconds
 */


/**
 * @param {Pomodo} instance
 */
function renderTime(instance) {
  const width = percentage(seconds, 60);
  document.body.style.setProperty("--progress-value", `${width}%`);
  document.getElementById("minutes").innerText = instance.minutes.toString();
  document.getElementById("seconds").innerText = instance.seconds.toString();
}

const debugElement = document.getElementById("debug");


let seconds = 60;
let minutes = 25;
setInterval(function () {
  debugElement.innerText = `Seconds: ${seconds}, Minutes: ${minutes}`;

  if (seconds === 0 && minutes === 0) {
    return;
  }

  if (minutes === 0) {

  }

  if (seconds === 0) {
    minutes--;
    seconds = 60;
  }


  seconds--;
  const mock = { seconds, isDone: false, minutes: minutes, title: "Task title", type: "WORK" };
  renderTime(mock);
}, 50);


const pomodElement = document.getElementById("pomod");

const typeElement = pomodElement.querySelector(".type");
const timeElement = pomodElement.querySelector(".time");
const titleElement = pomodElement.querySelector(".title");

const endpoint = "ws://127.0.0.1:8080/tasks/subscribe";

const ws = new WebSocket(endpoint);

ws.onopen = function () {
  console.log("Connected");
};

ws.onmessage = function (e) {
  /**
   * @type {Pomodo}
   */
  const data = JSON.parse(e.data);
  const doneStatus = data.isDone ? "✅" : "⏳️";

  timeElement.innerText = data.isDone ?
    doneStatus :
    `${data.minutes}:${data.seconds} ${doneStatus}`;
  titleElement.innerText = data.title;
  typeElement.innerText = data.type;
};

ws.onclose = function (e) {
  console.log("disconnected");
};
