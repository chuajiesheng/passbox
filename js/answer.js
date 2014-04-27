var timer = 30;

function update() {
    if (timer > 0) {
        timer -= 1;
        setHTML("countdown", timer)
    } else {
        setHTML("credential1", Math.floor(Math.random() * 10000000));
        setHTML("credential2", Math.floor(Math.random() * 10000000));
    }
}

setInterval(update, 1000);
