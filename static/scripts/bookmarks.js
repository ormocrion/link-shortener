document.addEventListener("DOMContentLoaded", ready);

window.onload = function() {
    document.getElementById('delBtn').addEventListener('click', delBtn)
}

async function ready() {
    let response = await fetch("http://localhost/getBookmarks");

    if (response.ok) {
        let text = await response.text();
        console.log("Запрос отправили");
        showOutput(text);
    } else {
        alert("Ошибка получения ссылок с сервера: " + response.status);
    }
}

function showOutput(text) {
    console.log(text);

    const arr = JSON.parse(text);

    for (let i=0; i < arr.length; i++) {
        let obj = arr[i];
       
        const div = document.createElement("div");

        const tagLink = document.createElement("a");
        tagLink.setAttribute('href', obj.Link);
        tagLink.innerHTML = obj.Link;

        const tagAlias = document.createElement("p");
        const nodeAlias = document.createTextNode("Псевдоним: " + obj.Alias);
        tagAlias.appendChild(nodeAlias);

        const tagDelete = document.createElement("button");
        tagDelete.setAttribute('id', 'delBtn');
        const nodeDelete = document.createTextNode("х");
        tagDelete.appendChild(nodeDelete);
        
        div.appendChild(tagLink);
        div.appendChild(tagAlias);
        div.appendChild(tagDelete);

        const footer = document.querySelector("footer");
        document.body.insertBefore(div, footer);
    }
}

async function execCommand(command) {
    let response = await fetch("http://localhost/service/", {
        method: "POST",
        body: command, 
    });

    if (response.ok) {
        console.log("Отправили запрос на удаление закладки");
    } else {
        alert("Ошибка удаления закладки. Попробуйте позже" + response.status);
    }
}

function delBtn() {
    console.log("Кнопка X нажата");
    execCommand("delete");
}