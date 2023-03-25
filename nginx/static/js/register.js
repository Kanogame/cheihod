"use strict"
import PostConnection from "./post.js";

const url = "http://localhost:10234/reg";
const regForm = document.getElementById("regForm");

regForm.addEventListener("submit", (e)=>{
    e.preventDefault();
    const data = new FormData(e.target);
    const bodyData = CreateRegJSON(data.get("Username"), data.get("Password"),
    data.get("Email"));
    SendLogin(bodyData);
});

function SendLogin(bodyData) {
    const post = new PostConnection(url, bodyData);
    const resp = post.SendData();
}

function CreateRegJSON (username, password, email) {
    return {
        Username: username,
        Password: password,
        Email: email,
    }
}