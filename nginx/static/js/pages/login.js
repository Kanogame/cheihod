"use strict"
import PostConnection from "../utils/post.js";
import CookieManager from "../utils/cookieManager.js";


const url = "http://127.0.0.1:10234/login";
const loginForm = document.getElementById("loginForm");
const cookie = new CookieManager();

loginForm.addEventListener("submit", (e)=>{
    e.preventDefault();
    const data = new FormData(e.target);
    const bodyData = CreateLogJSON(data.get("Username"), data.get("Password"))
    SendLogin(bodyData);
});

async function SendLogin(bodyData) {
    const post = new PostConnection(url, bodyData);
    const resp = await post.SendDataJson();
    cookie.setCookie("token", resp.token);
    alert("success");
}

function CreateLogJSON (username, password) {
    return {
        Username: username,
        Password: password,
    }
}