const BASE_API_URL = "/api/v1";

const getTemplateToElement = (tmpl) => {
    const tmplElement = document.createElement("template");
    tmplElement.innerHTML = tmpl.trim();

    return tmplElement.content.firstChild;
};

function showToast(message, type = "success") {
    const toastContainer = document.getElementById("toast-container");

    // Define styles for different types
    const typeStyles = {
        success: "green",
        error: "red",
        warning: "yellow",
    };
    const icon = {
        success: "check",
        error: "times",
        warning: "exclamation-triangle",
    };

    // Create toast element
    const toast = getTemplateToElement(`
        <div class="flex items-center bg-${typeStyles[type]}-100 border border-${typeStyles[type]}-300 rounded-lg shadow-lg p-4">
            <i class="fa fa-${icon[type]} text-${typeStyles[type]}-800" aria-hidden="true"></i>
            <p class="text-sm font-medium text-${typeStyles[type]}-800 ml-2">${message}</p> 
        </div>
    `);

    // Add toast to container
    toastContainer.appendChild(toast);

    // Trigger slide-in animation
    setTimeout(() => {
        toast.classList.remove("translate-x-full", "opacity-0");
        toast.classList.add("translate-x-0", "opacity-100");
    }, 100);

    // Remove toast after 3 seconds
    setTimeout(() => {
        toast.classList.remove("translate-x-0", "opacity-100");
        toast.classList.add("translate-x-full", "opacity-0");
        setTimeout(() => toast.remove(), 300); // Remove element after animation ends
    }, 3000);
}

function capitalize(str) {
    if (!str) return ""; // Handle empty or null strings
    return str.charAt(0).toUpperCase() + str.slice(1);
}

async function getFileContent(url) {
    let res = await fetch(url);
    return await res.text();
}

async function showInitialToast() {
    for (let toast of document.querySelectorAll("input.initial-toast")) {
        showToast(toast.dataset.text, toast.dataset.type);
    }
}

async function main() {
    showInitialToast();
}

main();
