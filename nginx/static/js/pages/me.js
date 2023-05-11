"use strict"
import { findTokenFull } from "../utils/token.js";
import PostConnection from "../utils/post.js";
import CookieManager from "../utils/cookieManager.js";
import Card from "../utils/createCard.js";

const cookie = new CookieManager();
const token = cookie.getCookie("token");

const name = document.getElementById("me-name");
const email = document.getElementById("me-email");
const ticket = document.getElementById("tickets");
init();

async function init() {
    const data = await findTokenFull();
    name.textContent = data.name;
    email.textContent = data.email;
    const post = new PostConnection("http://176.65.35.172/api/ticket/get", token);
    const res = await post.SendDataJson();
    console.log(res);
    for (const cardData of res) {
        const card = Card(cardData);
        ticket.append(card);
    }
}