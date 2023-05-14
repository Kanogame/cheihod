import fillCards from "../utils/fillCard.js";
import PostConnection from "../utils/post.js";

const container = document.getElementById("card-container");
const post = new PostConnection("http://176.65.35.172/api/places/get/30");
const res = await post.SendDataJson();

fillCards(res, container, false)