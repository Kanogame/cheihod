"use strict"
import PostConnection from "./post.js";
import CookieManager from "./cookieManager.js";

async function findToken(account, login, name) {
    const cookie = new CookieManager();
    const token = cookie.getCookie("token");
    if (token === undefined) {
        alert("no cookie");
    } else {
        console.log(token);
        account.classList.remove("disabled");
        login.classList.add("disabled");
        const post = new PostConnection("http://127.0.0.1:10234/token", token);
        const resp = await post.SendDataJson();
        name.textContent = resp.name;
    }
}

async function findTokenFull() {
    const cookie = new CookieManager();
    const token = cookie.getCookie("token");
    if (token === undefined) {
        alert("no cookie");
    } else {
        console.log(token);
        const post = new PostConnection("http://127.0.0.1:10234/token/full", token);
        const resp = await post.SendDataJson();
        return resp
    }
}

export default findToken