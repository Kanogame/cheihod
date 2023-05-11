"use strict"
import PostConnection from "../utils/post.js";
import CookieManager from "../utils/cookieManager.js";

const url = "http://176.65.35.172/api/reg";
const regForm = document.getElementById("regForm");
const cookie = new CookieManager();

regForm.addEventListener("submit", async (e)=>{
    e.preventDefault();
    const data = new FormData(e.target);
    const bodyData = CreateRegJSON(data.get("Username"), data.get("Password"),
    data.get("Email"));
    await SendLogin(bodyData);
});

async function SendLogin(bodyData) {
    const post = new PostConnection(url, bodyData);
    const resp = await post.SendDataJson();
    console.log(resp);
    if (resp.success === "true") {
        cookie.setCookie("token", resp.token);
        window.location.href = "../../index.html";
    } else {
        alert("произошла ошибка");
    }
}

function CreateRegJSON (username, password, email) {
    return {
        Username: username,
        Password: password,
        Email: email,
    }
}
