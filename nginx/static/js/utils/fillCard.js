import Card from "./createCard.js";

export default function fillCards(data, container, modal, openModal) {
    let i = 0;
    for (const cardData of data) {
        i++;
        if (i > 3) { break };
        const card = Card(cardData, modal ? openModal : null);
        container.append(card);
    }
}