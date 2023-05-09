"use strict"
import PostConnection from "./post.js";
import CookieManager from "./cookieManager.js";

async function findToken(account, login, name) {
    const cookie = new CookieManager();
    const token = cookie.getCookie("token");
    if (token === undefined) {
    } else {
        const post = new PostConnection("http://127.0.0.1/api/token", token);
        const resp = await post.SendDataJson();
        if (resp.success === "true") {
            name.textContent = resp.name;
        } else {
            alert("неверные данные");
        }
    }
}

export async function findTokenFull() {
    const cookie = new CookieManager();
    const token = cookie.getCookie("token");
    if (token === undefined) {
    } else {
        const post = new PostConnection("http://127.0.0.1/api/token/full", token);
        const resp = await post.SendDataJson();
        if (resp.success === "true") {
            return resp
        } else {
            alert("неверные данные");
            return undefined
        }
    }
}

export default findToken
