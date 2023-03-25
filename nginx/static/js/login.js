"use strict"

const url = "http://localhost:10234/";
const loginForm = document.getElementById("loginForm");

loginForm.addEventListener("submit", (e)=>{
    e.preventDefault();
    const data = new FormData(e.target);
    const bodyData = CreateLogJSON(data.get("Username"), data.get("Password"))
    SendLogin(bodyData);
});

function SendLogin(bodyData) {
    sendPost(url + "login", bodyData);
}

async function sendPost(url, object) {
    const resp = await fetch("http://localhost:10234/login", {
        method: "POST",
        mode: 'no-cors',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(object),
    });
    const text = await resp.text();
    alert(text);
}

function CreateLogJSON (username, password) {
    return {
        Username: username,
        Password: password,
    }
}