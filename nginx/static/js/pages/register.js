"use strict"
import PostConnection from "../utils/post.js";
import CookieManager from "../utils/cookieManager.js";

const url = "http://127.0.0.1:10234/reg";
const regForm = document.getElementById("regForm");
const cookie = new CookieManager();

regForm.addEventListener("submit", (e)=>{
    e.preventDefault();
    const data = new FormData(e.target);
    const bodyData = CreateRegJSON(data.get("Username"), data.get("Password"),
    data.get("Email"));
    SendLogin(bodyData);
});

function SendLogin(bodyData) {
    const post = new PostConnection(url, bodyData);
    const resp = post.SendDataText();
}

function CreateRegJSON (username, password, email) {
    return {
        Username: username,
        Password: password,
        Email: email,
    }
}