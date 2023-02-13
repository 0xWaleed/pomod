// 🎉

/**
 * @typedef Pomodo
 * @property {string} title
 * @property {boolean} isDone
 * @property {string} type
 * @property {number} minutes
 * @property {number} seconds
 */


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
