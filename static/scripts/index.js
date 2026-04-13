document.addEventListener("submit", sendData);

async function sendData(event) {
  const formData = new FormData(document.querySelector("form"));
  event.preventDefault();
  const link = formData.get("link");

  let response = await fetch("http://localhost/short/", {
    method: "POST",
    body: link,
  });

  if (response.ok) {
    let text = await response.text();

    showOutput(text);
  } else {
    alert("Ошибка HTTP: " + response.status);
  }
}

function showOutput(text) {
  const div = document.getElementById("output").style.display="block";

  const tag = document.getElementById("aliasLink");
  tag.value = text;
}

document.addEventListener("reset", resetOutput);

function resetOutput() {
  const div = document.getElementById("output").style.display="none";
}

window.onload = function() {
  document.getElementById("copy").addEventListener("click", copy);
}

function copy() {
  let copyText = document.getElementById("aliasLink").value;
  navigator.clipboard.writeText(copyText);

  enableCopyMessage();
  setTimeout(disableCopyMessage, 4000);
}

function enableCopyMessage() {
  document.getElementById("hoverCopyMessage").style.display = "block";
}

function disableCopyMessage() {
  document.getElementById("hoverCopyMessage").style.display = "none";
}
