"use strict"
import PostConnection from "./post.js";
import CookieManager from "./cookieManager.js";

async function findToken() {
    const cookie = new CookieManager();
    const token = cookie.getCookie("token");
    if (token === undefined) {
        alert("no cookie");
    } else {
        console.log(token);
        const post = new PostConnection("http://127.0.0.1:10234/token", token);
        const resp = await post.SendDataJson();
        console.log(resp);
    }
}

export default findToken