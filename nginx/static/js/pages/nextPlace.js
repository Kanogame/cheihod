"use strict"

import PostConnection from "../utils/post.js";
import CookieManager from "../utils/cookieManager.js";
import Card from "../utils/createCard.js";

const post = new PostConnection("http://127.0.0.1:10234/api/places/get/30");
const res = await post.SendDataJson();

const cookie = new CookieManager();
const token = cookie.getCookie("token");

const container = document.getElementById("card-container");
const modalContainer = document.getElementById("modalContainer");

let i = 0;
for (const cardData of res) {
    i++;
    if (i > 3) { break };
    const card = Card(cardData, openModal);
    container.append(card);
}

function openModal(cardData) {
    const modal = document.createElement("div");
    modal.classList.add("modal");
    const card = document.createElement("div");
    card.classList.add("nextup-card");
    card.classList.add("modal-body");
    const cardName = document.createElement("div");
    cardName.classList.add("nextup-name");
    cardName.textContent = "Вы желаете записаться на меропиятие?";
    card.append(cardName);
    
    const specs = document.createElement("div");
    specs.classList.add("nextup-specs");
    const place = document.createElement("div");
    place.innerHTML = `<img src="../sources/svg/location-dot-solid.svg" alt="" class="inline-img">${cardData.place}`;
    const time = document.createElement("div");
    time.innerHTML = `<img src="../sources/svg/clock-solid.svg" alt="" class="inline-img">${cardData.time}`;
    const free = document.createElement("div");
    free.innerHTML = `<img src="../sources/svg/ticket-solid.svg" alt="" class="inline-img">${cardData.capacity}`;
    specs.append(place);
    specs.append(time);
    specs.append(free);
    const button = document.createElement("a");
    button.classList.add("nextup-button");
    button.textContent= "Перейти";
    button.addEventListener("click", async () => {
        const post = new PostConnection("http://127.0.0.1:10234/api/ticket/add", {token: token, placeId: cardData.id});
        const resp = await post.SendDataJson();
        console.log(resp);
    })
    card.append(specs);
    const warn = document.createElement("div");
    warn.textContent = "Запись на мероприятие стоит денег, но пока все бесплатно)";
    card.append(warn);
    card.append(button);
    modal.append(card);
    modalContainer.append(modal);
}