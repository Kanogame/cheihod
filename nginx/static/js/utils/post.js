"use strict"

class PostConnection {
    constructor(url, data) {
        this.url = url;
        this.data = data;
    }

    async SendDataText() {
        const resp = await fetch(this.url, {
            method: "POST",
            mode: "cors",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(this.data),
        });
        const text = await resp.text();
        return text
    }

    async SendDataJson() {
        const resp = await fetch(this.url, {
            method: "POST",
            mode: "cors",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(this.data),
        });
        const text = await resp.json();
        return text
    }
}

export default PostConnection;