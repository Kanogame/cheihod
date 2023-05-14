function Card(cardData, openModal) {
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
    if (cardData.capacity != undefined) {
        free.innerHTML = `<img src="../sources/svg/ticket-solid.svg" alt="" class="inline-img">${cardData.capacity}`;
        specs.append(free);
    }
    specs.append(place);
    specs.append(time);
    const button = document.createElement("div");
    button.classList.add("nextup-button");
    button.textContent= "Перейти";
    button.addEventListener("click", () => {
        openModal(cardData);
    })
    card.append(specs);
    card.append(button);
    return card
}

export default Card;