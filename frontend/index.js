let connection = null;
const BASE_URL = "ws://localhost:8080/ws";

const lobby = document.getElementById("lobby");
const chat = document.getElementById("chat");

const usernameInput = document.getElementById("username");
const joinButton = document.getElementById("join");

const messageInput = document.getElementById("message");
const sendButton = document.getElementById("send");
const messagesContainer = document.getElementById("messages");

chat.setAttribute("hidden", true);

joinButton.addEventListener("click", () => {
  const username = usernameInput.value;
  if (username) {
    initConnection(username);
  }
});

function initConnection(username) {
  const url = BASE_URL + "?username=" + username;
  connection = new WebSocket(url);

  connection.addEventListener("open", () => {
    console.log("Connected to server");
    lobby.setAttribute("hidden", true);
    chat.removeAttribute("hidden");
  });

  connection.addEventListener("message", (event) => {
    const message = JSON.parse(event.data);
    renderMessage(message);
  });

  sendButton.addEventListener("click", () => {
    const message = messageInput.value;
    if (message) {
      connection.send(message);
      messageInput.value = "";
    }
  });
}

function renderMessage(message) {
  switch (message.data.type) {
    case "chat_message":
      messagesContainer.innerHTML += `<div>${message.sender.username}: ${message.data.content}</div>`;
      break;
    case "client_connected":
      messagesContainer.innerHTML += `<div>${message.sender.username} connected</div>`;
      break;
    case "client_disconnected":
      messagesContainer.innerHTML += `<div>${message.sender.username} disconnected</div>`;
      break;
    default:
      console.log("Unknown message type", message.data.type);
  }
}
