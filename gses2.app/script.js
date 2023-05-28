const currentURL = window.location.href;

function updateBTCPrice() {
    fetch(currentURL + "/api/rate")
        .then(response => response.json())
        .then(data => {
            let price = data.rate.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
            if (price.indexOf(".") === -1) {
                price += ".00";
            }
            document.getElementById("btc-price").textContent = "₴ " + price;
        })
        .catch(error => {
            console.error("Failed to update BTC price:", error);
        });
}

updateBTCPrice();

function subscribe() {
    const email = document.getElementById("email-input").value;

    fetch(currentURL + "/api/subscribe", {
        method: "POST",
        body: JSON.stringify({ email: email }),
        headers: {
            "Content-Type": "application/json"
        }
    })
        .then(response => response.json())
        .then(data => {
            toggleAnimation(true)
            console.log("Subscription response:", data);
            document.getElementById("subscription-form").reset();
        })
        .catch(error => {
            toggleAnimation(false)
            console.error("Failed to subscribe:", error);
        });
}

function sendEmails() {
    fetch(currentURL + "/api/sendEmails", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        }
    })
        .then(response => {
            console.log("Emails sent:", response.status);
        })
}

function toggleAnimation(green) {
    const target = document.getElementById("email-input");
    target.classList.remove('green-animation', 'red-animation');

    if (green) {
        target.classList.add('green-animation');
    } else {
        target.classList.add('red-animation');
    }

    const onAnimationEnd = () => {
        target.classList.remove('green-animation', 'red-animation');
        target.removeEventListener('animationend', onAnimationEnd);
    };

    target.addEventListener('animationend', onAnimationEnd);
}