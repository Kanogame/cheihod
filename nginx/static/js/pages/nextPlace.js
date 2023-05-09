"use strict"

import PostConnection from "../utils/post.js";

const post = new PostConnection("http://127.0.0.1:10234/api/places/get/30");
const res = await post.SendDataJson();

const container = document.getElementById("card-container");
const modalContainer = document.getElementById("modalContainer");

for (const cardData of res) {
    console.log(cardData);
    const card = document.createElement("div");
    card.classList.add("nextup-card");
    const cardName = document.createElement("div");
    cardName.classList.add("nextup-name");
    cardName.textContent = cardData.name;
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
    const button = document.createElement("div");
    button.classList.add("nextup-button");
    button.textContent= "Перейти";
    button.addEventListener("click", () => {
        openModal(cardData);
    })
    card.append(specs);
    card.append(button);
    container.append(card);
}

function openModal(cardData) {
    const modal = document.createElement("div");
    modal.classList.add("modal");
    modalContainer.append(modal);
}