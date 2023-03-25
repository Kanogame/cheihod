"use strict"
import PostConnection from "./post"

const url = "http://localhost:10234/reg";
const loginForm = document.getElementById("regForm");

loginForm.addEventListener("submit", (e)=>{
    e.preventDefault();
    const data = new FormData(e.target);
    const bodyData = CreateRegJSON(data.get("Username"), data.get("Password"),
    data.get("Email"));
    SendLogin(bodyData);
});

function SendLogin(bodyData) {
    const post = new PostConnection(url, bodyData);
    const resp = post.SendData();
    console.log(post.TextResponce(resp));
}

function CreateRegJSON (username, password, email) {
    return {
        Username: username,
        Password: password,
        Email: email,
    }
}