"use strict"
import {url} from "../config.json"
import {sendType} from "./post"

var loginForm = document.getElementById("loginForm");

loginForm.addEventListener("submit", (e)=>{
    e.preventDefault();
    const data = new FormData(e.target);
    const bodyData = CreateLogJSON(data.get("Username"), data.get("Password"))
    SendLogin(bodyData);
});

function SendLogin(bodyData) {
    sendType("login");
}