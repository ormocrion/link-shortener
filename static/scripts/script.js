function displayTime() {
  let element = document.getElementById("clock");
  let now = new Date();
  element.innerHTML = now.toLocaleTimeString();
  setTimeout(displayTime, 1000);
}

window.onload = displayTime;
