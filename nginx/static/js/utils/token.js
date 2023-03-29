"use strict"
import PostConnection from "./post.js";

function findToken() {
    const cookie = new CookieManager();
    const token = cookie.getCookie("token");
    if (token === undefined) {
        alert("no cookie");
    } else {
        console.log(token);
        const post = PostConnection("http://127.0.0.1:10234/token", token);
    }
}