const subscriptions = {};

async function subscribe(eventDescriptor, onmessage) {
    subscriptions[eventDescriptor] = onmessage;
    // Call websocket for this subscription
    window.socket.send(
        JSON.stringify({
            action: "subscribe",
            event: eventDescriptor,
        })
    );
}

async function onmessage(e) {
    let message;
    try {
        message = JSON.parse(e.data);
    } catch {
        console.log("invalid message from socket", message);
        return;
    }
    subscriptions[message["eventDescriptor"]](message);
}

async function main() {
    // Connect to the WebSocket server
    window.socket = new WebSocket(`ws://${window.location.host}/ws`);
    window.socket.onmessage = onmessage;
}

main();
