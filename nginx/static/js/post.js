"use strict"

class PostConnection {
    constructor(url, data) {
        this.url = url;
        this.data = data;
    }

    async SendData() {
        const resp = await fetch(this.url, {
            method: "POST",
            mode: 'no-cors',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(this.data),
        });
        const text = await resp.text();
        alert(text);
    }

    async TextResponce(resp) {
        return await resp.text();
    }

    async JsonResponce(resp) {
        return await resp.json();
    }
}

export default PostConnection;